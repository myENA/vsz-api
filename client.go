package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	DefaultScheme         = "https"
	DefaultPathPrefix     = "api/public"
	DefaultRequestTimeout = 5 * time.Second

	SessionCookieName = "JSESSIONID"
)

// SessionCookieRefreshFunc will be called in 2 circumstances:
// 1. Initial query failed with 401
// 2. Post-password refresh success
//
// It must attempt to refresh the cookie used during authenticated requests
type SessionCookieRefreshFunc func(client *Client, username, password string) (*http.Cookie, error)

func DefaultCookieRefresher(client *Client, username, password string) (*http.Cookie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.Config().RequestTimeout)
	defer cancel()
	resp, _, err := client.Session().LoginSessionLogonPost(ctx, &LoginSessionLogonPostRequest{
		Username: &username,
		Password: &password,
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unable to refresh cookie, expected 200 OK saw \"%d %s\"", resp.StatusCode, resp.Status)
	}

	cookie := TryExtractSessionCookie(resp)
	if cookie == nil {
		return nil, fmt.Errorf("unable to extract \"%s\" cookie", SessionCookieName)
	}

	return cookie, nil
}

// SessionPasswordRefreshFunc will only ever be called if the initial attempt to refresh the session cookie fails,
// and if implemented should return a new password to use for this user.
type SessionPasswordRefreshFunc func(client *Client, username, password string) (string, error)

func DefaultPasswordRefresher(_ *Client, username, _ string) (string, error) {
	return "", fmt.Errorf("user \"%s\" needs it's password refreshed, please implement a SessionPasswordRefreshFunc", username)
}

type Config struct {
	Address  string // REQUIRED address of your VSZ
	Username string // REQUIRED username to use for this config
	Password string // REQUIRED password to use for this config

	Scheme     string // OPTIONAL defaults to https
	PathPrefix string // OPTIONAL defaults to "api/public"

	RequestTimeout time.Duration // OPTIONAL will default to value of DefaultRequestTimeout

	CookieRefresher   SessionCookieRefreshFunc
	PasswordRefresher SessionPasswordRefreshFunc
}

// DefaultConfig creates a new Config object with a non-pooled client
func DefaultConfig(address, username, password string) *Config {
	return defaultConfig(address, username, password)
}

func defaultConfig(address, username, password string) *Config {
	return &Config{
		Address:           address,
		Username:          username,
		Password:          password,
		Scheme:            DefaultScheme,
		PathPrefix:        DefaultPathPrefix,
		RequestTimeout:    DefaultRequestTimeout,
		CookieRefresher:   DefaultCookieRefresher,
		PasswordRefresher: DefaultPasswordRefresher,
	}
}

type Client struct {
	*bridge
	config *Config
	client *http.Client

	lastUsed time.Time

	unauthProcessor *processor
	authProcessor   *processor

	closed     bool
	closedLock sync.RWMutex

	sessionCookie *http.Cookie

	refreshLock sync.Mutex
}

func NewClient(conf *Config, client *http.Client) *Client {
	def := defaultConfig(conf.Address, conf.Username, conf.Password)
	if conf.Scheme != "" {
		def.Scheme = conf.Scheme
	}
	if conf.PathPrefix != "" {
		def.PathPrefix = conf.PathPrefix
	}
	if conf.RequestTimeout > 0 {
		def.RequestTimeout = conf.RequestTimeout
	}
	if conf.CookieRefresher != nil {
		def.CookieRefresher = conf.CookieRefresher
	}
	if conf.PasswordRefresher != nil {
		def.PasswordRefresher = conf.PasswordRefresher
	}

	if client == nil {
		// shamelessly borrowed from https://github.com/hashicorp/go-cleanhttp/blob/master/cleanhttp.go
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true,
				MaxIdleConnsPerHost:   -1,
			},
		}
	}

	c := &Client{
		config:   conf,
		client:   client,
		lastUsed: time.Now(),
	}

	c.bridge = newBridge(c)

	c.unauthProcessor = newProcessor(context.Background())
	c.authProcessor = newProcessor(context.Background())

	return c
}

func (c *Client) Close() {
	c.closedLock.Lock()
	defer c.closedLock.Unlock()
	c.closed = true
	c.unauthProcessor.cancel()
	c.authProcessor.cancel()
}

func (c *Client) Config() Config {
	return *c.config
}

func (c *Client) LastUsed() time.Time {
	return c.lastUsed
}

// doRequest will attempt to execute a VSZ API request.
//
// Responds with:
// - http response (with closed body) or nil on error
// - body bytes or nil on error
// - any error
func (c *Client) doRequest(request *request, successCode int, out interface{}) (*http.Response, []byte, error) {
	c.closedLock.RLock()
	defer c.closedLock.RUnlock()

	if c.closed {
		return nil, nil, fmt.Errorf(
			"will not process request with path \"%s\", this client has been marked as closed",
			request.uri)
	}

	var err error
	var buff []byte
	var httpResponse *http.Response

	c.lastUsed = time.Now()

	// initial attempt
	httpResponse, err = c.process(request)
	if err != nil {
		return nil, nil, err
	}

	// read everything out of the body and close it.
	buff, err = ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, nil, err
	}

	// if this endpoint requires authentication and received a 401, assume we need to refresh cookie / credentials
	if request.authenticated && httpResponse.StatusCode == 401 {
		if c.config.CookieRefresher != nil {
			c.refreshLock.Lock()
			defer c.refreshLock.Unlock()

			// try to refresh cookie first...
			c.sessionCookie, err = c.config.CookieRefresher(c, c.config.Username, c.config.Password)
			if err != nil {
				if c.config.PasswordRefresher != nil {
					// should indicate to the implementor that this client's credentials are no longer valid
					c.config.Password, err = c.config.PasswordRefresher(c, c.config.Username, c.config.Password)
					if err != nil {
						return nil, nil, err
					}

					// if we were able to refresh the user's credentials, try logging in again...
					c.sessionCookie, err = c.config.CookieRefresher(c, c.config.Username, c.config.Password)
					if err != nil {
						return nil, nil, err
					}
				} else {
					return nil, nil, err
				}
			}

			// if we get this far, assume we were able to refresh the cookie one way or another
			httpResponse, err = c.process(request)
			if err != nil {
				return nil, nil, err
			}

			// read and close response body
			buff, err = ioutil.ReadAll(httpResponse.Body)
			httpResponse.Body.Close()
			if err != nil {
				return httpResponse, nil, err
			}
		} else {
			return httpResponse, nil, errors.New("unauthorized")
		}
	}

	// if success :D
	if successCode == httpResponse.StatusCode {
		if nil != out {
			err = json.Unmarshal(buff, out)
		}
		return httpResponse, buff, err
	}

	// if fail D:
	err2 := &Error{}
	err = json.Unmarshal(buff, err2)
	if err != nil {
		log.Printf("Unable to unmarshal response from call \"%s %s\": %s", request.method, request.uri, err)
		log.Printf("Response: %s", string(buff))
		err = fmt.Errorf("expected \"%d %s\", saw \"%s\"", successCode, http.StatusText(successCode), httpResponse.Status)
	} else {
		err = *err2
	}

	return httpResponse, buff, err
}

func (c *Client) process(request *request) (*http.Response, error) {
	var req *processRequest
	var resp *processResponse

	req = &processRequest{
		request:      request,
		client:       c,
		responseChan: make(chan *processResponse),
	}

	if request.authenticated {
		c.authProcessor.do(req)
	} else {
		c.unauthProcessor.do(req)
	}

	resp = <-req.responseChan

	return resp.http, resp.err
}
