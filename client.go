package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	DefaultScheme         = "https"
	DefaultPathPrefix     = "api/public"
	DefaultRequestTimeout = 2500 * time.Millisecond

	SessionCookieName = "JSESSIONID"
)

type Config struct {
	Address string // REQUIRED address of your VSZ

	Scheme     string // OPTIONAL defaults to https
	PathPrefix string // OPTIONAL defaults to "api/public"
}

// DefaultConfig creates a new Config object with a non-pooled client
func DefaultConfig(address string) *Config {
	return defaultConfig(address)
}

func defaultConfig(address string) *Config {
	return &Config{
		Address:    address,
		Scheme:     DefaultScheme,
		PathPrefix: DefaultPathPrefix,
	}
}

type Client struct {
	*bridge
	config *Config
	client *http.Client

	cookie    *http.Cookie
	cookieMu  sync.RWMutex
	cookieCAS uint64
}

func NewClient(conf *Config, client *http.Client) *Client {
	def := defaultConfig(conf.Address)
	if conf != nil {
		if conf.Scheme != "" {
			def.Scheme = conf.Scheme
		}
		if conf.PathPrefix != "" {
			def.PathPrefix = conf.PathPrefix
		}
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

func (c *Client) Config() Config {
	return *c.config
}

// doTest will attempt to execute the http request, testing the response to determine if we saw a malformation
func (c *Client) doTest(req *http.Request) (*http.Response, error, bool) {
	resp, err := c.client.Do(req)
	return resp, err, resp == nil && err != nil && (err == context.DeadlineExceeded || strings.HasPrefix(err.Error(), "malformed "))
}

func (c *Client) doTry(request *request, tries int) (*http.Response, error) {
	var httpRequest *http.Request
	var httpResponse *http.Response
	var cookie *http.Cookie
	var malformed bool
	var cas CookieCAS
	var err error

	if httpRequest, err = request.toHTTP(); err != nil {
		return nil, err
	}

	// if this api requires an active auth session, try to locate cookie
	if request.authenticated {
		if cookie, cas = request.ctx.Cookie(); cookie == nil {
			// if cookie is not set, attempt to refresh
			if _, err = request.ctx.RefreshCookie(c, cas); err != nil {
				// if refresh fails, bail out
				return nil, err
			}
		}
		// attempt to locate cookie one more time after refresh
		if cookie, cas = request.ctx.Cookie(); cookie == nil {
			// if cookie still nil, bail out
			return nil, fmt.Errorf("unable to locate cookie for user %s", request.ctx.Username())
		}
		// otherwise add cookie to request
		httpRequest.AddCookie(cookie)
	}

	if httpResponse, err, malformed = c.doTest(httpRequest); malformed {
		if tries < 1 {
			log.Printf("ERROR: Query \"%s\" returned \"%s\", which may indicate a malformed response.  Trying again...", request.uri, err)
			return c.doTry(request, tries+1)
		}
		log.Printf("ERROR: Query \"%s\" returned \"%s\", which may indicate a malformed response.  We have tried %d times, will not try again", request.uri, err, tries)
	}

	return httpResponse, err
}

// doRequest will attempt to execute a VSZ API request.
//
// Responds with:
// - http response (with closed body) or nil on error
// - body bytes or nil on error
// - any error
func (c *Client) doRequest(request *request, successCode int, out interface{}) (*http.Response, []byte, error) {
	var httpResponse *http.Response
	var cookie *http.Cookie
	var cas CookieCAS
	var buff []byte
	var err error

	if httpResponse, err = c.doTry(request, 0); err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		return nil, nil, err
	}

	// read everything out of the body and close it.
	buff, err = ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, nil, err
	}

	// if this endpoint requires authentication and received a 401, try to refresh cookie once.
	if request.authenticated && httpResponse.StatusCode == 401 {
		if _, err = request.ctx.RefreshCookie(c, cas); err != nil {
			// if fail, bail out
			return nil, nil, err
		}
		// try one more time to locate, this time we don't care about cas as this is the last time we'll try in this routine.
		if cookie, _ = request.ctx.Cookie(); cookie == nil {
			// if cookie still nil, bail out
			return nil, nil, fmt.Errorf("unable to refresh cookief or user %s", request.ctx.Username())
		}
		// try once final time to execute the request...
		if httpResponse, err = c.doTry(request, 0); err != nil {
			if httpResponse != nil {
				httpResponse.Body.Close()
			}
			return nil, nil, err
		}
		// from this point on we only care if the response code matches the provided "success" one, we will not attempt
		// to refresh cookie again in this routine.
		buff, err = ioutil.ReadAll(httpResponse.Body)
		httpResponse.Body.Close()
		if err != nil {
			return httpResponse, nil, err
		}
	}

	// if success :D
	if successCode == httpResponse.StatusCode {
		if out != nil {
			err = json.Unmarshal(buff, out)
		}
		return httpResponse, buff, err
	}

	// if fail D:
	err2 := &Error{}
	err = json.Unmarshal(buff, err2)
	if err != nil {
		log.Printf("ERROR: Unable to unmarshal response from call \"%s %s\": %s", request.method, request.uri, err)
		log.Printf("ERROR: Response: %s", string(buff))
		err = fmt.Errorf("expected \"%d %s\", saw \"%s\"", successCode, http.StatusText(successCode), httpResponse.Status)
	} else {
		err = *err2
	}

	return httpResponse, buff, err
}
