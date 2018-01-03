package api

import (
	"context"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type Rogue struct {
	client *Client
}
type (
	MarkRogueGetKnownRogueGet200Response struct {
		RogueMacList *string `json:"rogueMacList,omitempty"`
	}
)

// MarkRogueGetKnownRogueGet: Get Known Rogue AP list
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *MarkRogueGetKnownRogueGet200Response
// - error: Error seen or nil if none
func (r *Rogue) MarkRogueGetKnownRogueGet(ctx context.Context) (*http.Response, *MarkRogueGetKnownRogueGet200Response, error) {
	request := r.client.newRequest(ctx, "GET", "/v5_0/rogue/markKnown")
	request.authenticated = true

	out := &MarkRogueGetKnownRogueGet200Response{}
	httpResponse, _, err := r.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	MarkRogueMarkKnownRoguePostRequest struct {
		RogueMacList *string `json:"rogueMacList,omitempty"`
	}
)

// MarkRogueMarkKnownRoguePost: Mark a rogue AP as know
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *MarkRogueMarkKnownRoguePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *Rogue) MarkRogueMarkKnownRoguePost(ctx context.Context, requestBody *MarkRogueMarkKnownRoguePostRequest) (*http.Response, []byte, error) {
	request := r.client.newRequest(ctx, "POST", "/v5_0/rogue/markKnown")
	request.body = requestBody
	request.authenticated = true

	return r.client.doRequest(request, 204, nil)
}

type (
	MarkRogueUnmarkRoguePostRequest struct {
		RogueMacList *string `json:"rogueMacList,omitempty"`
	}
)

// MarkRogueUnmarkRoguePost: Unmark a rogue AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *MarkRogueUnmarkRoguePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *Rogue) MarkRogueUnmarkRoguePost(ctx context.Context, requestBody *MarkRogueUnmarkRoguePostRequest) (*http.Response, []byte, error) {
	request := r.client.newRequest(ctx, "POST", "/v5_0/rogue/unMark")
	request.body = requestBody
	request.authenticated = true

	return r.client.doRequest(request, 204, nil)
}
