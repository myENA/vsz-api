package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"sync/atomic"
)

var requestID uint64

type Request struct {
	id            uint64
	method        string
	uri           string
	authenticated bool

	queryParameters   map[string]string
	queryParametersMu sync.RWMutex

	pathParameters    map[string]string
	routeParametersMu sync.RWMutex

	headers   url.Values
	headersMu sync.RWMutex

	cookies   []*http.Cookie
	cookiesMu sync.RWMutex

	body   []byte
	bodyMu sync.RWMutex
}

func NewRequest(method, uri string, authenticated bool) *Request {
	r := &Request{
		id:              atomic.AddUint64(&requestID, 1),
		method:          method,
		uri:             uri,
		authenticated:   authenticated,
		queryParameters: make(map[string]string),
		pathParameters:  make(map[string]string),
		headers:         make(url.Values),
		cookies:         make([]*http.Cookie, 0),
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
	return r.authenticated
}

func (r *Request) SetHeaders(headers url.Values) {
	var l int
	r.headersMu.Lock()
	r.headers = make(url.Values, len(headers))
	for name, values := range headers {
		l = len(values)
		r.headers[name] = make([]string, l, l)
		copy(r.headers[name], values)
	}
	r.headersMu.Unlock()
}

func (r *Request) AddHeader(name, value string) {
	r.headersMu.Lock()
	r.headers.Add(name, value)
	r.headersMu.Unlock()
}

// SetHeader will attempt to overwrite an existing header with the same name, simply adding it if one is not found.
func (r *Request) SetHeader(name, value string) {
	r.headersMu.Lock()
	r.headers.Set(name, value)
	r.headersMu.Unlock()
}

// RemoveHeader will attempt to remove a header from this request, returning the value being removed.
func (r *Request) RemoveHeader(name string) (string, bool) {
	r.headersMu.Lock()
	current := r.headers.Get(name)
	if current == "" {
		r.headersMu.Unlock()
		return "", false
	}
	r.headers.Del(name)
	r.headersMu.Unlock()
	return current, true
}

// Headers will return a copy of current headers on this request
func (r *Request) Headers() url.Values {
	var l int
	r.headersMu.RLock()
	l = len(r.headers)
	if l == 0 {
		r.headersMu.RUnlock()
		return nil
	}
	headers := make(url.Values, l)
	for name, values := range r.headers {
		l = len(values)
		headers[name] = make([]string, l, l)
		copy(headers[name], values)
	}
	r.headersMu.RUnlock()
	return headers
}

func (r *Request) SetCookies(cookies []*http.Cookie) {
	r.cookiesMu.Lock()
	l := len(cookies)
	r.cookies = make([]*http.Cookie, l, l)
	copy(r.cookies, cookies)
	r.cookiesMu.Unlock()
}

func (r *Request) AddCookie(cookie *http.Cookie) {
	r.cookiesMu.Lock()
	r.cookies = append(r.cookies, cookie)
	r.cookiesMu.Unlock()
}

// SetCookie will attempt to locate and overwrite a cookie with the same name, simply appending it to the list if one is
// not found
func (r *Request) SetCookie(cookie *http.Cookie) {
	r.cookiesMu.Lock()
	for i, cc := range r.cookies {
		if cc.Name == cookie.Name {
			r.cookies[i] = cookie
			r.cookiesMu.Unlock()
			return
		}
	}
	r.cookies = append(r.cookies, cookie)
	r.cookiesMu.Unlock()
}

func (r *Request) RemoveCookie(name string) (*http.Cookie, bool) {
	r.cookiesMu.Lock()
	for _, cookie := range r.cookies {
		if cookie.Name == name {
			r.cookiesMu.Unlock()
			return cookie, true
		}
	}
	r.cookiesMu.Unlock()
	return nil, false
}

// Cookies will return a copy of the list of cookies to be used with this request
// NOTE: The cookies are pointers.  Be aware.
func (r *Request) Cookies() []*http.Cookie {
	r.cookiesMu.RLock()
	l := len(r.cookies)
	if l == 0 {
		r.cookiesMu.RUnlock()
		return nil
	}
	cookies := make([]*http.Cookie, l, l)
	copy(cookies, r.cookies)
	r.cookiesMu.RUnlock()
	return cookies
}

func (r *Request) SetQueryParameters(params map[string]string) {
	r.queryParametersMu.Lock()
	r.queryParameters = make(map[string]string, len(params))
	for k, v := range params {
		r.queryParameters[k] = v
	}
	r.queryParametersMu.Unlock()
}

func (r *Request) SetQueryParameter(param, value string) {
	r.queryParametersMu.Lock()
	r.queryParameters[param] = value
	r.queryParametersMu.Unlock()
}

func (r *Request) QueryParameters() map[string]string {
	r.queryParametersMu.RLock()
	params := make(map[string]string, len(r.queryParameters))
	for k, v := range r.queryParameters {
		params[k] = v
	}
	r.queryParametersMu.RUnlock()
	return params
}

func (r *Request) SetPathParameters(params map[string]string) {
	r.routeParametersMu.Lock()
	r.pathParameters = make(map[string]string, len(params))
	for k, v := range params {
		r.pathParameters[k] = v
	}
	r.routeParametersMu.Unlock()
}

func (r *Request) SetPathParameter(param, value string) {
	r.routeParametersMu.Lock()
	r.pathParameters[param] = value
	r.routeParametersMu.Unlock()
}

func (r *Request) PathParameters() map[string]string {
	r.routeParametersMu.RLock()
	params := make(map[string]string, len(r.pathParameters))
	for k, v := range r.pathParameters {
		params[k] = v
	}
	r.routeParametersMu.RUnlock()
	return params
}

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

func (r *Request) SetBodyModel(model interface{}) error {
	if model == nil {
		r.SetBody(nil)
		return nil
	}
	b, err := json.Marshal(model)
	if err != nil {
		return err
	}
	r.SetBody(b)
	return nil
}

func (r *Request) Body() []byte {
	r.bodyMu.RLock()
	l := len(r.body)
	if l == 0 {
		r.bodyMu.RUnlock()
		return nil
	}
	tmp := make([]byte, l, l)
	copy(tmp, r.body)
	r.bodyMu.RUnlock()
	return tmp
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
