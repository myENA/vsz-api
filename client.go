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
	Hostname      string        // REQUIRED address of your VSZ, including port
	Authenticator Authenticator // REQUIRED authentication decorator to use with this client

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
		if debug {
			log.Printf("[request-%d] Error reading response body: %s", request.id, err)
		}
		return httpResponse, nil, err
	} else if debug {
		log.Printf("[request-%d] Response body read and closed", request.id)
	}

	// check for and attempt to handle 401 unauthorized
	if httpResponse.StatusCode == 401 && request.auth {
		if debug {
			log.Printf("[request-%d] Request returned 401, will try invalidate -> refresh loop once", request.id)
		}
		// TODO: is this logic ok...?
		if httpResponse, _, err = c.handleUnauthorized(ctx, request, httpResponse, cas); err != nil {
			if httpResponse != nil {
				httpResponse.Body.Close()
			}
			if debug {
				log.Printf("[request-%d] Saw error after invalidate -> refresh loop: %s", request.id, err)
			}
			return nil, nil, err
		}
	}

	// test for "success"
	if httpResponse.StatusCode == successCode {
		if debug {
			log.Printf("[request-%d] Response code matched expected: %d", request.id, successCode)
		}
		if out != nil {
			if debug {
				log.Printf("[request-%d] Response has model, attempting to unmarshal...", request.id)
			}
			err = json.Unmarshal(buff, out)
			if debug {
				if err != nil {
					log.Printf("[request-%d] Error parsing response body: %s", request.id, err)
				} else {
					log.Printf("[request-%d] Response body parsed successfully", request.id)
				}
			}
		} else if debug {
			log.Printf("[request-%d] Request does not have model, nothing to unmarshal", request.id)
		}
		return httpResponse, buff, err
	} else if debug {
		log.Printf("[request-%d] Response code did not match expected: %d", request.id, successCode)
	}

	if debug {
		log.Printf("[request-%d] Attempting to unmarshal into error...", request.id)
	}

	// finally, try to unmarshal response body into Error type
	err2 := &VSZError{}
	err = json.Unmarshal(buff, err2)
	if err != nil {
		log.Printf("[request-%d] ERROR: Unable to unmarshal response: %s", request.id, err)
		log.Printf("[request-%d] ERROR: Response: %s", request.id, string(buff))
		err = fmt.Errorf("expected \"%d\", saw \"%s\"", successCode, httpResponse.Status)
	} else {
		if debug {
			log.Printf("[request-%d] Error unmarshalled successfully", request.id)
		}
		err = *err2
	}

	if debug {
		log.Printf("[request-%d] Execution complete", request.id)
	}

	return httpResponse, buff, err
}

func (c *Client) handleUnauthorized(ctx context.Context, request *Request, httpResponse *http.Response, cas AuthCAS) (*http.Response, AuthCAS, error) {
	var err error
	// attempt to invalidate
	if cas, err = c.auth.Invalidate(ctx, cas); err != nil {
		if debug {
			log.Printf("[request-%d] Invalidate failed: %s", request.id, err)
		}
		return httpResponse, cas, err
	} else if debug {
		log.Printf("[request-%d] Invalidate succeeded, calling Refresh...", request.id)
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

	if debug {
		log.Printf("[request-%d] Handling...", request.id)
	}

	if httpRequest, err = request.toHTTP(ctx, c.config); err != nil {
		if debug {
			log.Printf("[request-%d] Failed to build *http.Request: %s", request.id, err)
		}
		return nil, cas, err
	}

	// if this api requires an active auth session, try to locate cookie
	if request.auth {
		if debug {
			log.Printf("[request-%d] Auth required, calling Decorate...", request.id)
		}
		if cas, err = c.auth.Decorate(ctx, httpRequest); err != nil {
			if debug {
				log.Printf("[request-%d] Decorate returned error, will try calling refresh: %s", request.id, err)
			}
			if cas, err = c.auth.Refresh(ctx, c, cas); err != nil {
				if debug {
					log.Printf("[request-%d] Refresh returned error, bailing out: %s", request.id, err)
				}
				return nil, cas, err
			} else if cas, err = c.auth.Decorate(ctx, httpRequest); err != nil {
				if debug {
					log.Printf("[request-%d] Post-refresh Decorate returned error, bailing out: %s", request.id, err)
				}
				return nil, cas, err
			}
		} else if debug {
			log.Printf("[request-%d] Decorate succeeded", request.id)
		}
	}

	if debug {
		log.Printf("[request-%d] Creating request context....", request.id)
	}

	if debug {
		log.Printf("[request-%d] Executing...", request.id)
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
		if debug {
			log.Printf("[request-%d] Request failed: %s", request.id, err)
		}
		return cas, nil, err
	}
	if debug {
		log.Printf("[request-%d] Received \"%s\"", request.id, httpResponse.Status)
	}
	return cas, httpResponse, err
}
