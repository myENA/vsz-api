package api

import (
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-15T14:33:32-0500
// API Version: v5

type ShutdownAPI struct {
	client *Client
}

// AdministrationShutdownPost: Use this API command to shut down the controller.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ShutdownAPI) AdministrationShutdownPost(ctx *UserContext) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/shutdown")
	request.authenticated = true
	return s.client.doRequest(request, 200, nil)
}
