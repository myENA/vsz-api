package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type SMSGatewayAPI struct {
	client *Client
}
type (
	SmsGatewayGetSmsGatewayGet200Response struct {
		AccountSid *string `json:"accountSid,omitempty"`
		AuthToken  *string `json:"authToken,omitempty"`
		DomainID   *string `json:"domainId,omitempty"`
		Enabled    *int    `json:"enabled,omitempty"`
		From       *string `json:"from,omitempty"`
		ID         *string `json:"id,omitempty"`
		ServerName *string `json:"serverName,omitempty"`
		ServerType *string `json:"serverType,omitempty"`
	}
)

// SmsGatewayGetSmsGatewayGet: Get SMS gateway
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - domainId (UUIDv4): MSP/Partner domain id of the SMS gateway
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SmsGatewayGetSmsGatewayGet200Response
// - error: Error seen or nil if none
func (s *SMSGatewayAPI) SmsGatewayGetSmsGatewayGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *SmsGatewayGetSmsGatewayGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/smsGateway", true)
	request.SetQueryParameter("domainId", domainId)
	out := &SmsGatewayGetSmsGatewayGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SmsGatewayUpdateSmsGatewayPatchRequest struct {
		AccountSid *string `json:"accountSid,omitempty"`
		AuthToken  *string `json:"authToken,omitempty"`
		DomainID   *string `json:"domainId,omitempty"`
		Enabled    *int    `json:"enabled,omitempty"`
		From       *string `json:"from,omitempty"`
		ID         *string `json:"id,omitempty"`
		ServerName *string `json:"serverName,omitempty"`
		ServerType *string `json:"serverType,omitempty"`
	}
)

// SmsGatewayUpdateSmsGatewayPatch: Update SMS gateway
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SmsGatewayUpdateSmsGatewayPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SMSGatewayAPI) SmsGatewayUpdateSmsGatewayPatch(ctx context.Context, requestBody *SmsGatewayUpdateSmsGatewayPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/smsGateway", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SmsGatewayCreateSmsGatewayPost200ResponseListSlice []*SmsGatewayCreateSmsGatewayPost200ResponseList

	SmsGatewayCreateSmsGatewayPost200ResponseList struct {
		AccountSid *string `json:"accountSid,omitempty"`
		AuthToken  *string `json:"authToken,omitempty"`
		DomainID   *string `json:"domainId,omitempty"`
		Enabled    *int    `json:"enabled,omitempty"`
		From       *string `json:"from,omitempty"`
		ID         *string `json:"id,omitempty"`
		ServerName *string `json:"serverName,omitempty"`
		ServerType *string `json:"serverType,omitempty"`
	}

	SmsGatewayCreateSmsGatewayPost200Response struct {
		Extra      *json.RawMessage                                   `json:"extra,omitempty"`
		FirstIndex *int                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                              `json:"hasMore,omitempty"`
		List       SmsGatewayCreateSmsGatewayPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                               `json:"totalCount,omitempty"`
	}
)

// SmsGatewayCreateSmsGatewayPost: Create SMS gateway
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SmsGatewayCreateSmsGatewayPost200Response
// - error: Error seen or nil if none
func (s *SMSGatewayAPI) SmsGatewayCreateSmsGatewayPost(ctx context.Context) (*http.Response, *SmsGatewayCreateSmsGatewayPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/smsGateway", true)
	out := &SmsGatewayCreateSmsGatewayPost200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}
