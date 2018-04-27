package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-04-27T15:10:38-0500
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
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *MarkRogueGetKnownRogueGet200Response
// - error: Error seen or nil if none
func (r *RogueAPI) MarkRogueGetKnownRogueGet(ctx context.Context) (*http.Response, *MarkRogueGetKnownRogueGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/rogue/markKnown", true)
	out := &MarkRogueGetKnownRogueGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
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
func (r *RogueAPI) MarkRogueMarkKnownRoguePost(ctx context.Context, requestBody *MarkRogueMarkKnownRoguePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/rogue/markKnown", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return r.client.Ensure(ctx, request, 204, nil)
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
func (r *RogueAPI) MarkRogueUnmarkRoguePost(ctx context.Context, requestBody *MarkRogueUnmarkRoguePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/rogue/unMark", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return r.client.Ensure(ctx, request, 204, nil)
}
