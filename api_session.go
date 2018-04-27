package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-04-27T15:10:38-0500
// API Version: v5

type SessionAPI struct {
	client *Client
}

// LoginSessionLogoffDelete: Use this API command to log off of the controller.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SessionAPI) LoginSessionLogoffDelete(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("DELETE", "/v5_0/session", true)
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	LoginSessionRetrieveGet200ResponseAPIVersionsSlice []string

	LoginSessionRetrieveGet200Response struct {
		AdminID           *string                                            `json:"adminId,omitempty"`
		AdminRoleID       *string                                            `json:"adminRoleId,omitempty"`
		APIVersions       LoginSessionRetrieveGet200ResponseAPIVersionsSlice `json:"apiVersions,omitempty"`
		ClientIP          *string                                            `json:"clientIp,omitempty"`
		CpID              *string                                            `json:"cpId,omitempty"`
		DomainID          *string                                            `json:"domainId,omitempty"`
		MvnoID            *string                                            `json:"mvnoId,omitempty"`
		TimeZoneUtcOffset *float64                                           `json:"timeZoneUtcOffset,omitempty"`
	}
)

// LoginSessionRetrieveGet: Use this API command to retrieve information about the current logon session.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *LoginSessionRetrieveGet200Response
// - error: Error seen or nil if none
func (s *SessionAPI) LoginSessionRetrieveGet(ctx context.Context) (*http.Response, *LoginSessionRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/session", true)
	out := &LoginSessionRetrieveGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	LoginSessionLogonPostRequestAPIVersionsSlice []string

	LoginSessionLogonPostRequest struct {
		APIVersions       LoginSessionLogonPostRequestAPIVersionsSlice `json:"apiVersions,omitempty"`
		DomainName        *string                                      `json:"domainName,omitempty"`
		Password          *string                                      `json:"password,omitempty"`
		TimeZoneUtcOffset *string                                      `json:"timeZoneUtcOffset,omitempty"`
		Username          *string                                      `json:"username,omitempty"`
	}

	LoginSessionLogonPost200Response struct {
		APIVersion        *string `json:"apiVersion,omitempty"`
		ControllerVersion *string `json:"controllerVersion,omitempty"`
	}
)

// LoginSessionLogonPost: Use this API command to log on to the controller and acquire a valid logon session.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *LoginSessionLogonPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *LoginSessionLogonPost200Response
// - error: Error seen or nil if none
func (s *SessionAPI) LoginSessionLogonPost(ctx context.Context, requestBody *LoginSessionLogonPostRequest) (*http.Response, *LoginSessionLogonPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/session", false)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &LoginSessionLogonPost200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}
