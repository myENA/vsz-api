package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-04-27T15:10:38-0500
// API Version: v5

type RestartAPI struct {
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
func (r *RestartAPI) AdministrationRestartPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/restart", true)
	return r.client.Ensure(ctx, request, 200, nil)
}
