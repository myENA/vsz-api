package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"sync"
	"sync/atomic"
)

var requestID uint64

type Request struct {
	id     uint64
	method string
	uri    string
	auth   bool

	pathParams   map[string]string
	pathParamsMu sync.RWMutex

	queryParams   map[string]string
	queryParamsMu sync.RWMutex

	body   []byte
	bodyMu sync.RWMutex
}

func NewRequest(method, uri string, authenticated bool) *Request {
	r := &Request{
		id:          atomic.AddUint64(&requestID, 1),
		method:      method,
		uri:         uri,
		auth:        authenticated,
		pathParams:  make(map[string]string),
		queryParams: make(map[string]string),
	}
	return r
}

func (r *Request) ID() uint64 {
	return r.id
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) URI() string {
	return r.uri
}

func (r *Request) Authenticated() bool {
	return r.auth
}

// PathParameters will return a clone of the current list of path parameters
func (r *Request) PathParameters() map[string]string {
	r.pathParamsMu.RLock()
	l := len(r.pathParams)
	if l == 0 {
		r.pathParamsMu.RUnlock()
		return make(map[string]string)
	}
	tmp := make(map[string]string, len(r.pathParams))
	for k, v := range r.pathParams {
		tmp[k] = v
	}
	r.pathParamsMu.RUnlock()
	return tmp
}

// SetPathParameter will add / overwrite a single path parameter to the provided value
func (r *Request) SetPathParameter(parameter, value string) {
	r.pathParamsMu.Lock()
	r.pathParams[parameter] = value
	r.pathParamsMu.Unlock()
}

// SetPathParameters will replace the existing map of path parameters with whatever you pass in.
func (r *Request) SetPathParameters(parameters map[string]string) {
	r.pathParamsMu.Lock()
	r.pathParams = parameters
	r.pathParamsMu.Unlock()
}

// QueryParameters will return a clone of the current list of query parameters
func (r *Request) QueryParameters() map[string]string {
	r.queryParamsMu.RLock()
	l := len(r.queryParams)
	if l == 0 {
		r.queryParamsMu.RUnlock()
		return make(map[string]string)
	}
	tmp := make(map[string]string, l)
	for k, v := range r.queryParams {
		tmp[k] = v
	}
	r.queryParamsMu.RUnlock()
	return tmp
}

// SetQueryParameter will add / overwrite a single query parameter with the provided value
func (r *Request) SetQueryParameter(parameter, value string) {
	r.queryParamsMu.Lock()
	r.queryParams[parameter] = value
	r.queryParamsMu.Unlock()
}

// SetQueryParameters will replace the existing query parameter map with whatever you pass in
func (r *Request) SetQueryParameters(parameters map[string]string) {
	r.queryParamsMu.Lock()
	r.queryParams = parameters
	r.queryParamsMu.Unlock()
}

// Body will return a clone of the current request body
func (r *Request) Body() []byte {
	r.bodyMu.RLock()
	l := len(r.body)
	if l == 0 {
		r.bodyMu.RUnlock()
		return make([]byte, 0)
	}
	tmp := make([]byte, l, l)
	copy(tmp, r.body)
	r.bodyMu.RUnlock()
	return tmp
}

// SetBody will copy the contents of the provided body to this request
func (r *Request) SetBody(body []byte) {
	r.bodyMu.Lock()
	l := len(body)
	if l == 0 {
		r.body = nil
		r.bodyMu.Unlock()
		return
	}
	r.body = make([]byte, l, l)
	copy(r.body, body)
	r.bodyMu.Unlock()
}

// SetBodyModel will attempt to json.Marshal the input to use as the body of this request.  The body will only be set
// if the marshalling is successful.
func (r *Request) SetBodyModel(model interface{}) error {
	r.bodyMu.Lock()
	if model == nil {
		r.body = nil
		r.bodyMu.Unlock()
		return nil
	}
	b, err := json.Marshal(model)
	if err != nil {
		r.bodyMu.Unlock()
		return err
	}
	r.body = b
	r.bodyMu.Unlock()
	return nil
}

func (r *Request) compileURI() string {
	pathParams := r.PathParameters()
	queryParams := r.QueryParameters()
	uri := r.uri
	if len(pathParams) > 0 {
		for k, v := range pathParams {
			uri = strings.Replace(uri, fmt.Sprintf("{%s}", k), v, 1)
		}
	}
	if len(queryParams) > 0 {
		uri = fmt.Sprintf("%s%s", uri, buildQueryParamString(queryParams))
	}
	return uri
}

// toHTTP will attempt to construct an executable http.request
func (r *Request) toHTTP(ctx context.Context, config *Config) (*http.Request, error) {
	var err error
	var httpRequest *http.Request

	body := r.Body()
	bodyLen := len(body)
	compiledURL := fmt.Sprintf("%s://%s:%d%s", config.Scheme, config.Hostname, config.Port, path.Join(config.PathPrefix, r.compileURI()))

	// if debug mode is enabled, prepare a big'ol log statement.
	if debug {
		logMsg := fmt.Sprintf("[request-%d] Preparing request \"%s %s\"", r.id, r.method, compiledURL)

		if bodyLen == 0 {
			logMsg = fmt.Sprintf("%s without body", logMsg)
		} else {
			logMsg = fmt.Sprintf("%s with body: %s", logMsg, string(body))
		}

		log.Print(logMsg)
	}

	if bodyLen == 0 {
		httpRequest, err = http.NewRequest(r.method, compiledURL, nil)
	} else {
		httpRequest, err = http.NewRequest(r.method, compiledURL, bytes.NewBuffer(body))
	}

	if err != nil {
		return nil, err
	}

	if bodyLen != 0 {
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	httpRequest.Header.Set("Accept", "application/json")

	return httpRequest.WithContext(ctx), nil
}
