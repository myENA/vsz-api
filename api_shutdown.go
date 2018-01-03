package api

import (
	"context"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type Shutdown struct {
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
func (s *Shutdown) AdministrationShutdownPost(ctx context.Context) (*http.Response, []byte, error) {
	request := s.client.newRequest(ctx, "POST", "/v5_0/shutdown")
	request.authenticated = true

	return s.client.doRequest(request, 200, nil)
}
