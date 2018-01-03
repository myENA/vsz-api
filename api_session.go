package api

import (
	"context"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type Session struct {
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
func (s *Session) LoginSessionLogoffDelete(ctx context.Context) (*http.Response, []byte, error) {
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/session")
	request.authenticated = true

	return s.client.doRequest(request, 200, nil)
}

type (
	LoginSessionRetrieveGet200ResponseAPIVersionsSlice []*string

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
func (s *Session) LoginSessionRetrieveGet(ctx context.Context) (*http.Response, *LoginSessionRetrieveGet200Response, error) {
	request := s.client.newRequest(ctx, "GET", "/v5_0/session")
	request.authenticated = true

	out := &LoginSessionRetrieveGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	LoginSessionLogonPostRequestAPIVersionsSlice []*string

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
func (s *Session) LoginSessionLogonPost(ctx context.Context, requestBody *LoginSessionLogonPostRequest) (*http.Response, *LoginSessionLogonPost200Response, error) {
	request := s.client.newRequest(ctx, "POST", "/v5_0/session")
	request.body = requestBody

	out := &LoginSessionLogonPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
