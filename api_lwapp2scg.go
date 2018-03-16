package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type LWAPP2SCGAPI struct {
	client *Client
}
type (
	LwappToScgModifyBasicPatchRequest struct {
		NatIPTranslation *bool    `json:"natIpTranslation,omitempty"`
		PasvMaxPort      *float64 `json:"pasvMaxPort,omitempty"`
		PasvMinPort      *float64 `json:"pasvMinPort,omitempty"`
		Policy           *string  `json:"policy,omitempty"`
	}
)

// LwappToScgModifyBasicPatch: Use this API command to modify the basic information of the Lwapp Config.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *LwappToScgModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (l *LWAPP2SCGAPI) LwappToScgModifyBasicPatch(ctx context.Context, requestBody *LwappToScgModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/lwapp2scg", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return l.client.Ensure(ctx, request, 204, nil)
}

type (
	LwappToScgModifyAplistPatchRequest []*string
)

// LwappToScgModifyAplistPatch: Use this API command to modify the apList of the Lwapp Config.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *LwappToScgModifyAplistPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (l *LWAPP2SCGAPI) LwappToScgModifyAplistPatch(ctx context.Context, requestBody LwappToScgModifyAplistPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/lwapp2scg/apList", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return l.client.Ensure(ctx, request, 204, nil)
}
