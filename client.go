package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultScheme     = "https"
	DefaultPathPrefix = "api/public"

	SessionCookieName = "JSESSIONID"
)

type Config struct {
	Address       string        // REQUIRED address of your VSZ, including port
	Authenticator Authenticator // REQUIRED authentication decorator to use with this client

	Scheme     string // OPTIONAL defaults to https
	PathPrefix string // OPTIONAL defaults to "api/public"
}

// DefaultConfig creates a new ClientConfig object with a non-pooled client
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
	auth   Authenticator
}

func NewClient(conf *Config, authenticator Authenticator, client *http.Client) (*Client, error) {
	if authenticator == nil {
		return nil, errors.New("authenticator cannot be nil")
	}
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
		auth:   authenticator,
	}

	c.bridge = newBridge(c)

	return c, nil
}

func (c *Client) ClientConfig() Config {
	return *c.config
}

func (c *Client) Do(ctx context.Context, request *Request) (*http.Response, error) {
	var httpResponse *http.Response
	var ctxTTL time.Duration
	var err error
	if ctxTime, ok := ctx.Deadline(); ok && ctx.Err() == nil {
		ctxTTL = ctxTime.Sub(time.Now())
	}
	if httpResponse, _, err = c.tryDo(ctx, ctxTTL, request, 0); err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		if debug {
			log.Printf("[request-%d] Request failed: %s", request.id, err)
		}
		return nil, err
	}
	if debug {
		log.Printf("[request-%d] Received \"%s\"", request.id, httpResponse.Status)
	}
	return httpResponse, err
}

// Ensure will attempt to execute the request, further attempting to unmarshal the response into a pointer provided
// to "out" given that the response status code matches that passed as "successCode". Failing that, this method
// will attempt to unmarshal the seen response into an Error struct
func (c *Client) Ensure(ctx context.Context, request *Request, successCode int, out interface{}) (*http.Response, []byte, error) {
	var httpResponse *http.Response
	var buff []byte
	var err error

	if httpResponse, err = c.Do(ctx, request); err != nil {
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

	// if success :D
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

	// if fail D:
	err2 := &Error{}
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

// doTest will attempt to execute the http request, testing the response to determine if we saw a malformation
func (c *Client) doTest(req *http.Request) (*http.Response, bool, error) {
	resp, err := c.client.Do(req)
	return resp,
		resp == nil && err != nil && (err == context.DeadlineExceeded || strings.HasPrefix(err.Error(), "malformed ")),
		err
}

func (c *Client) tryDo(ctx context.Context, ctxTTL time.Duration, request *Request, tries int) (*http.Response, AuthCAS, error) {
	var httpRequest *http.Request
	var httpResponse *http.Response
	var malformed bool
	var cas AuthCAS
	var cancel context.CancelFunc
	var err error

	// TODO: this is pretty inefficient, think of something better.
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	if debug {
		log.Printf("[request-%d] Attempt %d", request.id, tries+1)
	}

	if ctx == nil {
		if ctxTTL == 0 {
			log.Printf("[request-%d] WARNING: Provided context had infinite TTL, using infinite TTL for retry attempt...", request.id)
			ctx = context.Background()
		} else {
			if debug {
				log.Printf("[request-%d] Using %s ttl on retry attempt", request.id, ctxTTL)
			}
			ctx, cancel = context.WithTimeout(context.Background(), ctxTTL)
		}
	}

	if httpRequest, err = request.toHTTP(ctx, c.config.Scheme, c.config.Address, c.config.PathPrefix); err != nil {
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
		if cas, err = c.auth.Decorate(httpRequest); err != nil {
			if debug {
				log.Printf("[request-%d] Decorate returned error, will try calling refresh: %s", request.id, err)
			}
			if cas, err = c.auth.Refresh(c, cas); err != nil {
				if debug {
					log.Printf("[request-%d] Refresh returned error, bailing out: %s", request.id, err)
				}
				return nil, cas, err
			} else if cas, err = c.auth.Decorate(httpRequest); err != nil {
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
	if httpResponse, malformed, err = c.doTest(httpRequest); malformed {
		if tries < 1 {
			log.Printf("[request-%d] ERROR: Saw \"%s\", which may indicate a malformed response.  Trying again...", request.id, err)
			return c.tryDo(nil, ctxTTL, request, tries+1)
		}
		log.Printf("[request-%d] ERROR: Saw \"%s\", which may indicate a malformed response.  We have tried %d times, will not try again", request.id, err, tries)
	}

	return httpResponse, cas, err
}
