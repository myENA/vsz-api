package api

import (
	"context"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type LWAPP2SCG struct {
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
func (l *LWAPP2SCG) LwappToScgModifyBasicPatch(ctx context.Context, requestBody *LwappToScgModifyBasicPatchRequest) (*http.Response, []byte, error) {
	request := l.client.newRequest(ctx, "PATCH", "/v5_0/lwapp2scg")
	request.body = requestBody
	request.authenticated = true

	return l.client.doRequest(request, 204, nil)
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
func (l *LWAPP2SCG) LwappToScgModifyAplistPatch(ctx context.Context, requestBody LwappToScgModifyAplistPatchRequest) (*http.Response, []byte, error) {
	request := l.client.newRequest(ctx, "PATCH", "/v5_0/lwapp2scg/apList")
	request.body = requestBody
	request.authenticated = true

	return l.client.doRequest(request, 204, nil)
}
