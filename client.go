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
	"sync/atomic"
	"time"
)

const (
	DefaultScheme         = "https"
	DefaultPathPrefix     = "api/public"
	DefaultRequestTimeout = 2500 * time.Millisecond

	SessionCookieName = "JSESSIONID"
)

type Config struct {
	Address  string // REQUIRED address of your VSZ
	Username string // REQUIRED username to use for this config
	Password string // REQUIRED password to use for this config

	Scheme     string // OPTIONAL defaults to https
	PathPrefix string // OPTIONAL defaults to "api/public"

	RequestTimeout time.Duration // OPTIONAL will default to value of DefaultRequestTimeout
}

// DefaultConfig creates a new Config object with a non-pooled client
func DefaultConfig(address, username, password string) *Config {
	return defaultConfig(address, username, password)
}

func defaultConfig(address, username, password string) *Config {
	return &Config{
		Address:        address,
		Username:       username,
		Password:       password,
		Scheme:         DefaultScheme,
		PathPrefix:     DefaultPathPrefix,
		RequestTimeout: DefaultRequestTimeout,
	}
}

type Client struct {
	*bridge
	config *Config
	client *http.Client

	closed     bool
	closedLock sync.RWMutex

	cookie    *http.Cookie
	cookieMu  sync.RWMutex
	cookieCAS uint64
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
		config: conf,
		client: client,
	}

	c.bridge = newBridge(c)

	return c
}

func (c *Client) Close() {
	c.closedLock.Lock()
	defer c.closedLock.Unlock()
	c.closed = true
}

func (c *Client) Config() Config {
	return *c.config
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

	httpRequest, err := request.toHTTP()
	if err != nil {
		return nil, nil, err
	}

	if request.authenticated {
		cookie, err := c.sessionCookie()
		if err != nil {
			return nil, nil, err
		}
		httpRequest.AddCookie(cookie)
	}

	httpResponse, err := c.client.Do(httpRequest)
	if err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		return nil, nil, err
	}

	// read everything out of the body and close it.
	buff, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, nil, err
	}

	// if this endpoint requires authentication and received a 401, assume we need to refresh cookie / credentials
	if request.authenticated && httpResponse.StatusCode == 401 {
		if err := c.refreshCookie(); err != nil {
			return nil, nil, err
		}
		httpRequest, err = request.toHTTP()
		if err != nil {
			return nil, nil, err
		}
		cookie, err := c.sessionCookie()
		if err != nil {
			return nil, nil, err
		}
		httpRequest.AddCookie(cookie)
		httpResponse, err = c.client.Do(httpRequest)
		if err != nil {
			if httpResponse != nil {
				httpResponse.Body.Close()
			}
			return nil, nil, err
		}
		buff, err = ioutil.ReadAll(httpResponse.Body)
		httpResponse.Body.Close()
		if err != nil {
			return httpResponse, nil, err
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

func (c *Client) refreshCookie() error {
	cas := atomic.LoadUint64(&c.cookieCAS)
	c.cookieMu.Lock()
	defer c.cookieMu.Unlock()
	if atomic.LoadUint64(&c.cookieCAS) != cas {
		// should mean that somebody else updated the cookie
		if c.cookie == nil {
			return errors.New("cookie was unable to be refreshed")
		}
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.config.RequestTimeout)
	defer cancel()
	resp, _, err := c.Session().LoginSessionLogonPost(ctx, &LoginSessionLogonPostRequest{
		Username: &c.config.Username,
		Password: &c.config.Password,
	})
	if err != nil {
		c.cookie = nil
		c.cookieCAS = atomic.AddUint64(&c.cookieCAS, 1)
		return err
	}
	cookie := TryExtractSessionCookie(resp)
	if cookie == nil {
		c.cookie = nil
		c.cookieCAS = atomic.AddUint64(&c.cookieCAS, 1)
		return fmt.Errorf("unable to refresh cookie, expected \"200 OK\" saw \"%s\"", resp.Status)
	}
	c.cookie = cookie
	c.cookieCAS = atomic.AddUint64(&c.cookieCAS, 1)
	return nil
}

func (c *Client) sessionCookie() (*http.Cookie, error) {
	c.cookieMu.RLock()
	if c.cookie == nil {
		c.cookieMu.RUnlock()
		if err := c.refreshCookie(); err != nil {
			return nil, err
		} else {
			c.cookieMu.RLock()
		}
	}
	if c.cookie != nil {
		cookie := c.cookie
		c.cookieMu.Unlock()
		return cookie, nil
	}
	c.cookieMu.RUnlock()
	return nil, errors.New("cookie was unable to be fetched")
}
