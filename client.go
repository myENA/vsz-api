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
	DefaultRequestTimeout = 5 * time.Second

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
func (c *Client) doTest(req *http.Request) (*http.Response, bool, error) {
	resp, err := c.client.Do(req)
	return resp,
		resp == nil && err != nil && (err == context.DeadlineExceeded || strings.HasPrefix(err.Error(), "malformed ")),
		err
}

func (c *Client) doTry(request *request, tries int) (*http.Response, CookieCAS, error) {
	var httpRequest *http.Request
	var httpResponse *http.Response
	var cookie *http.Cookie
	var malformed bool
	var cas CookieCAS
	var ctx context.Context
	var cancel context.CancelFunc
	var err error

	if debug {
		log.Printf("[request-%d] Attempt %d", request.id, tries+1)
	}

	if httpRequest, err = request.toHTTP(); err != nil {
		if debug {
			log.Printf("[request-%d] Failed to build *http.Request: %s", request.id, err)
		}
		return nil, cas, err
	}

	// if this api requires an active auth session, try to locate cookie
	if request.authenticated {
		if debug {
			log.Printf("[request-%d] Auth required", request.id)
		}
		if cookie, cas = request.ctx.Cookie(); cookie == nil {
			if debug {
				log.Printf("[request-%d] Cookie not found, attempting to refresh...", request.id)
			}
			// if cookie is not set, attempt to refresh
			if _, err = request.ctx.RefreshCookie(c, cas); err != nil {
				if debug {
					log.Printf("[request-%d] Cookie unable to be refreshed: %s", request.id, err)
				}
				// if refresh fails, bail out
				return nil, cas, err
			}
			// attempt to locate cookie one more time after refresh
			if cookie, cas = request.ctx.Cookie(); cookie == nil {
				if debug {
					log.Printf("[request-%d] Unable to refresh cookie")
				}
				// if cookie still nil, bail out
				return nil, cas, fmt.Errorf("unable to locate cookie for user %s", request.ctx.Username())
			} else if debug {
				log.Printf("[request-%d] Cookie refreshed", request.id)
			}
		}
		if debug {
			log.Printf("[request-%d] Adding session cookie to header...", request.id)
		}
		// otherwise add cookie to request
		httpRequest.AddCookie(cookie)
	}

	if debug {
		log.Printf("[request-%d] Creating request context....", request.id)
	}
	ctx, cancel = request.ctx.RequestContext()
	defer cancel()
	httpRequest = httpRequest.WithContext(ctx)

	if debug {
		log.Printf("[request-%d] Executing...", request.id)
	}
	if httpResponse, malformed, err = c.doTest(httpRequest); malformed {
		if tries < 1 {
			log.Printf("[request-%d] ERROR: Saw \"%s\", which may indicate a malformed response.  Trying again...", request.id, err)
			return c.doTry(request, tries+1)
		}
		log.Printf("[request-%d] ERROR: Saw \"%s\", which may indicate a malformed response.  We have tried %d times, will not try again", request.id, err, tries)
	}

	return httpResponse, cas, err
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

	if httpResponse, cas, err = c.doTry(request, 0); err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		if debug {
			log.Printf("[request-%d] Request failed: %s", request.id, err)
		}
		return nil, nil, err
	}

	if debug {
		log.Printf("[request-%d] Received \"%s\"", request.id, httpResponse.Status)
	}

	// if this endpoint requires authentication and received a 401, try to refresh cookie once.
	if request.authenticated && httpResponse.StatusCode == 401 {
		if debug {
			log.Printf("[request-%d] Saw 401 response, will attempt to refresh cookie...", request.id)
		}
		// close previous attempt's body
		httpResponse.Body.Close()
		if _, err = request.ctx.RefreshCookie(c, cas); err != nil {
			log.Printf("[request-%d] Failed to refresh cookie, bailing out: %s", request.id, err)
			// if fail, bail out
			return nil, nil, err
		}
		// try one more time to locate, this time we don't care about cas as this is the last time we'll try in this routine.
		if cookie, _ = request.ctx.Cookie(); cookie == nil {
			if debug {
				log.Printf("[request-%d] Unable to refresh cookie, bailing out", request.id)
			}
			// if cookie still nil, bail out
			return nil, nil, fmt.Errorf("unable to refresh cookie for user %s", request.ctx.Username())
		} else if debug {
			log.Printf("[request-%d] Cookie refreshed successfully, attempting again...", request.id)
		}
		// try once final time to execute the request...
		if httpResponse, _, err = c.doTry(request, 0); err != nil {
			if httpResponse != nil {
				httpResponse.Body.Close()
			}
			if debug {
				log.Printf("[request-%d] 2nd request attempt failed, bailing out: %s", request.id, err)
			}
			return nil, nil, err
		} else if debug {
			log.Printf("[request-%d] 2nd request attempt received \"%s\"", request.id, httpResponse.Status)
		}
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
