package api

import (
	"context"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type Restart struct {
	client *Client
}

// AdministrationRestartPost: Use this API command to restart the controller.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *Restart) AdministrationRestartPost(ctx context.Context) (*http.Response, []byte, error) {
	request := r.client.newRequest(ctx, "POST", "/v5_0/restart")
	request.authenticated = true

	return r.client.doRequest(request, 200, nil)
}
