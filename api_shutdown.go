package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type ShutdownAPI struct {
	client *Client
}

// AdministrationShutdownPost: Use this API command to shut down the controller.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ShutdownAPI) AdministrationShutdownPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/shutdown", true)
	return s.client.Ensure(ctx, request, 200, nil)
}
