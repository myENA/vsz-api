package api

import (
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-15T14:33:32-0500
// API Version: v5

type RogueAPI struct {
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
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *MarkRogueGetKnownRogueGet200Response
// - error: Error seen or nil if none
func (r *RogueAPI) MarkRogueGetKnownRogueGet(ctx *UserContext) (*http.Response, *MarkRogueGetKnownRogueGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
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
// - ctx (*UserContext): Context to use for this request
// - requestBody: *MarkRogueMarkKnownRoguePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RogueAPI) MarkRogueMarkKnownRoguePost(ctx *UserContext, requestBody *MarkRogueMarkKnownRoguePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
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
// - ctx (*UserContext): Context to use for this request
// - requestBody: *MarkRogueUnmarkRoguePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RogueAPI) MarkRogueUnmarkRoguePost(ctx *UserContext, requestBody *MarkRogueUnmarkRoguePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := r.client.newRequest(ctx, "POST", "/v5_0/rogue/unMark")
	request.body = requestBody
	request.authenticated = true
	return r.client.doRequest(request, 204, nil)
}
