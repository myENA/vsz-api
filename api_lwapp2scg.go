package api

import (
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-15T14:33:32-0500
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
// - ctx (*UserContext): Context to use for this request
// - requestBody: *LwappToScgModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (l *LWAPP2SCGAPI) LwappToScgModifyBasicPatch(ctx *UserContext, requestBody *LwappToScgModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
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
// - ctx (*UserContext): Context to use for this request
// - requestBody: *LwappToScgModifyAplistPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (l *LWAPP2SCGAPI) LwappToScgModifyAplistPatch(ctx *UserContext, requestBody LwappToScgModifyAplistPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := l.client.newRequest(ctx, "PATCH", "/v5_0/lwapp2scg/apList")
	request.body = requestBody
	request.authenticated = true
	return l.client.doRequest(request, 204, nil)
}
