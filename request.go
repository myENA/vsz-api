package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

var requestID uint64

type request struct {
	id     uint64
	config *Config
	method string
	uri    string
	ctx    context.Context

	authenticated   bool
	pathParameters  map[string]string
	queryParameters map[string]string
	body            interface{}
}

func (c *Client) newRequest(ctx context.Context, method, uri string) *request {
	r := &request{
		id:     atomic.AddUint64(&requestID, 1),
		config: c.config,
		uri:    uri,
		ctx:    ctx,
		method: method,
	}
	return r
}

// toHTTP will attempt to construct an executable http.request
func (r *request) toHTTP() (*http.Request, error) {
	var err error
	var httpRequest *http.Request

	body := r.body
	method := r.method
	compiledURL := compileRequestURLString(r.config.Scheme, r.config.Address, r.config.PathPrefix, r.uri, r.pathParameters, r.queryParameters)

	if debug {
		logMsg := fmt.Sprintf("Preparing request \"%s %s\"", method, compiledURL)

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
