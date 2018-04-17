package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	DefaultScheme     = "https"
	DefaultPort       = 7443
	DefaultPathPrefix = "/api/public"

	// RecommendedMinimumRequestTTL is the minimum recommended timeout value for any given context used during an api
	// request to the VSZ.  This is designed to provide enough time to allow for re-authentication to happen if need be.
	//
	// There is no guard against setting a TTL lower than this value.  It is merely here as a suggestion.
	RecommendedMinimumRequestTTL = 2500 * time.Millisecond
)

type Config struct {
	Hostname string // REQUIRED address of your VSZ, including port

	Port       int    // OPTIONAL defaults to 7443
	Scheme     string // OPTIONAL defaults to https
	PathPrefix string // OPTIONAL defaults to "/api/public"
}

// DefaultConfig creates a new ClientConfig object with a non-pooled client
func DefaultConfig(hostname string) *Config {
	return defaultConfig(hostname)
}

func defaultConfig(address string) *Config {
	return &Config{
		Hostname:   address,
		Scheme:     DefaultScheme,
		Port:       DefaultPort,
		PathPrefix: DefaultPathPrefix,
	}
}

type Client struct {
	*bridge
	config *Config
	client *http.Client
	auth   Authenticator
}

func NewClient(conf *Config, authenticator Authenticator, client *http.Client) (*Client, error) {
	if authenticator == nil {
		return nil, errors.New("authenticator cannot be nil")
	}
	def := defaultConfig(conf.Hostname)
	if conf != nil {
		if conf.Scheme != "" {
			def.Scheme = conf.Scheme
		}
		if conf.Port > 0 {
			def.Port = conf.Port
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
		auth:   authenticator,
	}

	c.bridge = newBridge(c)

	return c, nil
}

func (c *Client) ClientConfig() Config {
	return *c.config
}

func (c *Client) Do(ctx context.Context, request *Request) (*http.Response, error) {
	_, httpResponse, err := c.do(ctx, request)
	return httpResponse, err
}

// Ensure will attempt to execute the request, initiating an Authenticator Invalidate -> Refresh loop if a 401 is seen.
//
// If the "success" status code is seen, it will further attempt to unmarshal the response into a pointer provided
// to "out". Failing that, this method will attempt to unmarshal the seen response into an VSZError struct
func (c *Client) Ensure(ctx context.Context, request *Request, successCode int, out interface{}) (*http.Response, []byte, error) {
	var httpResponse *http.Response
	var cas AuthCAS
	var buff []byte
	var err error

	if cas, httpResponse, err = c.do(ctx, request); err != nil {
		return nil, nil, err
	}

	// read everything out of the body and close it.
	buff, err = ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, nil, err
	}

	// check for and attempt to handle 401 unauthorized
	if httpResponse.StatusCode == 401 && request.auth {
		// TODO: is this logic ok...?
		if httpResponse, _, err = c.handleUnauthorized(ctx, request, httpResponse, cas); err != nil {
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
	}

	// test for "success"
	if httpResponse.StatusCode == successCode {
		if out != nil {
			err = json.Unmarshal(buff, out)
		}
		return httpResponse, buff, err
	}

	// finally, try to unmarshal response body into Error type
	err2 := &VSZError{}
	err = json.Unmarshal(buff, err2)
	if err != nil {
		log.Printf("[request-%d] ERROR: Unable to unmarshal response: %s", request.id, err)
		log.Printf("[request-%d] ERROR: Response: %s", request.id, string(buff))
		err = fmt.Errorf("expected \"%d\", saw \"%s\"", successCode, httpResponse.Status)
	} else {
		err = *err2
	}

	return httpResponse, buff, err
}

func (c *Client) handleUnauthorized(ctx context.Context, request *Request, httpResponse *http.Response, cas AuthCAS) (*http.Response, AuthCAS, error) {
	var err error
	// attempt to invalidate
	if cas, err = c.auth.Invalidate(ctx, cas); err != nil {
		return httpResponse, cas, err
	}
	// close previous attempt's body
	httpResponse.Body.Close()
	return c.tryDo(ctx, request)
}

func (c *Client) tryDo(ctx context.Context, request *Request) (*http.Response, AuthCAS, error) {
	var httpRequest *http.Request
	var httpResponse *http.Response
	var cas AuthCAS
	var err error

	if httpRequest, err = request.toHTTP(ctx, c.config); err != nil {
		return nil, cas, err
	}

	// if this api requires an active auth session, try to locate cookie
	if request.auth {
		if cas, err = c.auth.Decorate(ctx, httpRequest); err != nil {
			if cas, err = c.auth.Refresh(ctx, c, cas); err != nil {
				return nil, cas, err
			} else if cas, err = c.auth.Decorate(ctx, httpRequest); err != nil {
				return nil, cas, err
			}
		}
	}

	httpResponse, err = c.client.Do(httpRequest)
	return httpResponse, cas, err
}

func (c *Client) do(ctx context.Context, request *Request) (AuthCAS, *http.Response, error) {
	if ctx == nil {
		return 0, nil, errors.New("ctx must not be nil")
	}
	var httpResponse *http.Response
	var cas AuthCAS
	var err error
	if httpResponse, cas, err = c.tryDo(ctx, request); err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		return cas, nil, err
	}
	return cas, httpResponse, err
}
