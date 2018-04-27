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
// Generation Date: 2018-04-27T15:10:38-0500
// API Version: v5

type SCIAPI struct {
	client *Client
}
type (
	SciModifyScienabledSettingPatchRequest struct {
		SciEnabled *bool `json:"sciEnabled,omitempty"`
	}
)

// SciModifyScienabledSettingPatch: Use this API command to modify SCI settings is enabled or not.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SciModifyScienabledSettingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SCIAPI) SciModifyScienabledSettingPatch(ctx context.Context, requestBody *SciModifyScienabledSettingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/sci/sciEnabled", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SciRetrieveSciAcceptedEventCodesGet200ResponseListSlice []*SciRetrieveSciAcceptedEventCodesGet200ResponseList

	SciRetrieveSciAcceptedEventCodesGet200ResponseList struct {
		Code *float64 `json:"code,omitempty"`
		Type *string  `json:"type,omitempty"`
	}

	SciRetrieveSciAcceptedEventCodesGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       SciRetrieveSciAcceptedEventCodesGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// SciRetrieveSciAcceptedEventCodesGet: Use this API command to retrieve SciAcceptedEventCodes.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SciRetrieveSciAcceptedEventCodesGet200Response
// - error: Error seen or nil if none
func (s *SCIAPI) SciRetrieveSciAcceptedEventCodesGet(ctx context.Context) (*http.Response, *SciRetrieveSciAcceptedEventCodesGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/sci/sciEventCode", true)
	out := &SciRetrieveSciAcceptedEventCodesGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SciModifySciAcceptedEventCodesPostRequestSciAcceptedEventCodesSlice []float64

	SciModifySciAcceptedEventCodesPostRequest struct {
		SciAcceptedEventCodes SciModifySciAcceptedEventCodesPostRequestSciAcceptedEventCodesSlice `json:"sciAcceptedEventCodes,omitempty"`
	}
)

// SciModifySciAcceptedEventCodesPost: Use this API command to modify SciAcceptedEventCodes.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SciModifySciAcceptedEventCodesPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SCIAPI) SciModifySciAcceptedEventCodesPost(ctx context.Context, requestBody *SciModifySciAcceptedEventCodesPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/sci/sciEventCode", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SciDeleteSciProfileListDeleteRequestListSlice []*SciDeleteSciProfileListDeleteRequestList

	SciDeleteSciProfileListDeleteRequestList struct {
		ID *string `json:"id,omitempty"`
	}

	SciDeleteSciProfileListDeleteRequest struct {
		List SciDeleteSciProfileListDeleteRequestListSlice `json:"list,omitempty"`
	}
)

// SciDeleteSciProfileListDelete: Use this API command to delete sciProfile list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SciDeleteSciProfileListDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SCIAPI) SciDeleteSciProfileListDelete(ctx context.Context, requestBody *SciDeleteSciProfileListDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("DELETE", "/v5_0/sci/sciProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SciRetrieveSciProfileListGet200ResponseListSlice []*SciRetrieveSciProfileListGet200ResponseList

	SciRetrieveSciProfileListGet200ResponseList struct {
		ID            *string  `json:"id,omitempty"`
		SciPassword   *string  `json:"sciPassword,omitempty"`
		SciPriority   *float64 `json:"sciPriority,omitempty"`
		SciProfile    *string  `json:"sciProfile,omitempty"`
		SciServerHost *string  `json:"sciServerHost,omitempty"`
		SciServerPort *string  `json:"sciServerPort,omitempty"`
		SciSystemID   *string  `json:"sciSystemId,omitempty"`
		SciUser       *string  `json:"sciUser,omitempty"`
	}

	SciRetrieveSciProfileListGet200Response struct {
		Extra *json.RawMessage                                 `json:"extra,omitempty"`
		List  SciRetrieveSciProfileListGet200ResponseListSlice `json:"list,omitempty"`
	}
)

// SciRetrieveSciProfileListGet: Use this API command to retrieve sciProfile list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SciRetrieveSciProfileListGet200Response
// - error: Error seen or nil if none
func (s *SCIAPI) SciRetrieveSciProfileListGet(ctx context.Context) (*http.Response, *SciRetrieveSciProfileListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/sci/sciProfile", true)
	out := &SciRetrieveSciProfileListGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SciCreateSciProfilePostRequest struct {
		SciPassword   *string `json:"sciPassword,omitempty"`
		SciProfile    *string `json:"sciProfile,omitempty"`
		SciServerHost *string `json:"sciServerHost,omitempty"`
		SciServerPort *string `json:"sciServerPort,omitempty"`
		SciSystemID   *string `json:"sciSystemId,omitempty"`
		SciUser       *string `json:"sciUser,omitempty"`
	}

	SciCreateSciProfilePost200Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// SciCreateSciProfilePost: Use this API command to create sciProfile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SciCreateSciProfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SciCreateSciProfilePost200Response
// - error: Error seen or nil if none
func (s *SCIAPI) SciCreateSciProfilePost(ctx context.Context, requestBody *SciCreateSciProfilePostRequest) (*http.Response, *SciCreateSciProfilePost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/sci/sciProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &SciCreateSciProfilePost200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SciModifySciPrioritiesPostRequestListSlice []*SciModifySciPrioritiesPostRequestList

	SciModifySciPrioritiesPostRequestList struct {
		ID          *string  `json:"id,omitempty"`
		SciPriority *float64 `json:"sciPriority,omitempty"`
		SciProfile  *string  `json:"sciProfile,omitempty"`
	}

	SciModifySciPrioritiesPostRequest struct {
		List SciModifySciPrioritiesPostRequestListSlice `json:"list,omitempty"`
	}
)

// SciModifySciPrioritiesPost: Use this API command to modify sciPriorities.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SciModifySciPrioritiesPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SCIAPI) SciModifySciPrioritiesPost(ctx context.Context, requestBody *SciModifySciPrioritiesPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/sci/sciProfile/sciPriority", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 200, nil)
}

// SciDeleteSciProfileDelete: Use this API command to delete sciProfile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SCIAPI) SciDeleteSciProfileDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/sci/sciProfile/{id}", true)
	request.SetPathParameter("id", id)
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SciRetrieveSciProfileGet200Response struct {
		ID            *string  `json:"id,omitempty"`
		SciPassword   *string  `json:"sciPassword,omitempty"`
		SciPriority   *float64 `json:"sciPriority,omitempty"`
		SciProfile    *string  `json:"sciProfile,omitempty"`
		SciServerHost *string  `json:"sciServerHost,omitempty"`
		SciServerPort *string  `json:"sciServerPort,omitempty"`
		SciSystemID   *string  `json:"sciSystemId,omitempty"`
		SciUser       *string  `json:"sciUser,omitempty"`
	}
)

// SciRetrieveSciProfileGet: Use this API command to retrieve sciProfile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SciRetrieveSciProfileGet200Response
// - error: Error seen or nil if none
func (s *SCIAPI) SciRetrieveSciProfileGet(ctx context.Context, id string) (*http.Response, *SciRetrieveSciProfileGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/sci/sciProfile/{id}", true)
	request.SetPathParameter("id", id)
	out := &SciRetrieveSciProfileGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SciModifySciProfilePatchRequest struct {
		ID            *string `json:"id,omitempty"`
		SciPassword   *string `json:"sciPassword,omitempty"`
		SciProfile    *string `json:"sciProfile,omitempty"`
		SciServerHost *string `json:"sciServerHost,omitempty"`
		SciServerPort *string `json:"sciServerPort,omitempty"`
		SciSystemID   *string `json:"sciSystemId,omitempty"`
		SciUser       *string `json:"sciUser,omitempty"`
	}
)

// SciModifySciProfilePatch: Use this API command to modify sciProfile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *SciModifySciProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SCIAPI) SciModifySciProfilePatch(ctx context.Context, id string, requestBody *SciModifySciProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/sci/sciProfile/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return s.client.Ensure(ctx, request, 200, nil)
}
