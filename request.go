package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"sync/atomic"
)

var requestID uint64

type request struct {
	id     uint64
	config *Config
	method string
	uri    string
	ctx    *UserContext

	authenticated   bool
	pathParameters  map[string]string
	queryParameters map[string]string
	body            interface{}
}

func (c *Client) newRequest(ctx *UserContext, method, uri string) *request {
	r := &request{
		id:     atomic.AddUint64(&requestID, 1),
		config: c.config,
		uri:    uri,
		ctx:    ctx,
		method: method,
	}
	return r
}

func (r *request) compileURL() string {
	url := fmt.Sprintf(
		"%s://%s",
		r.config.Scheme,
		path.Join(r.config.Address, r.config.PathPrefix, r.uri))

	if nil == r.pathParameters && nil == r.queryParameters {
		return url
	}

	if len(r.pathParameters) > 0 {
		for name, value := range r.pathParameters {
			url = strings.Replace(url, fmt.Sprintf("{%s}", name), value, 1)
		}
	}

	if len(r.queryParameters) > 0 {
		url = fmt.Sprintf("%s%s", url, buildQueryParamString(r.queryParameters))
	}

	return url
}

// toHTTP will attempt to construct an executable http.request
func (r *request) toHTTP() (*http.Request, error) {
	var err error
	var httpRequest *http.Request

	body := r.body
	method := r.method
	compiledURL := r.compileURL()

	// if debug mode is enabled, prepare a big'ol log statement.
	if debug {
		logMsg := fmt.Sprintf("[request-%d] Preparing request \"%s %s\"", r.id, method, compiledURL)

		if nil == body {
			logMsg = fmt.Sprintf("%s without body", logMsg)
		} else {
			b, err := json.Marshal(body)
			if nil != err {
				logMsg = fmt.Sprintf("%s with unknown body (%s)", logMsg, err.Error())
			} else {
				logMsg = fmt.Sprintf("%s with body: %s", logMsg, string(b))
			}
		}

		log.Print(logMsg)
	}

	if nil == body {
		httpRequest, err = http.NewRequest(method, compiledURL, nil)
	} else {
		buff := &bytes.Buffer{}
		err = json.NewEncoder(buff).Encode(body)
		if nil != err {
			return nil, err
		}
		httpRequest, err = http.NewRequest(method, compiledURL, buff)
	}

	if nil != err {
		return nil, err
	}

	if nil != body {
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	httpRequest.Header.Set("Accept", "application/json")

	return httpRequest, nil
}
