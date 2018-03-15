package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-15T14:33:32-0500
// API Version: v5

type UsersAPI struct {
	client *Client
}
type (
	ScgUserDeleteScgUserDelete1RequestIDListSlice []*string

	ScgUserDeleteScgUserDelete1Request struct {
		IDList ScgUserDeleteScgUserDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// ScgUserDeleteScgUserDelete1: Delete multiple SCG user
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ScgUserDeleteScgUserDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (u *UsersAPI) ScgUserDeleteScgUserDelete1(ctx *UserContext, requestBody *ScgUserDeleteScgUserDelete1Request) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := u.client.newRequest(ctx, "DELETE", "/v5_0/users")
	request.body = requestBody
	request.authenticated = true
	return u.client.doRequest(request, 200, nil)
}

type (
	ScgUserAddScgUserPostRequest struct {
		CompanyName      *string `json:"companyName,omitempty"`
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		Email            *string `json:"email,omitempty"`
		Enabled          *int    `json:"enabled,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Passphrase       *string `json:"passphrase,omitempty"`
		Phone            *string `json:"phone,omitempty"`
		RealName         *string `json:"realName,omitempty"`
		TenantUUID       *string `json:"tenantUUID,omitempty"`
		Title            *string `json:"title,omitempty"`
		UserName         *string `json:"userName,omitempty"`
	}

	ScgUserAddScgUserPost200Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ScgUserAddScgUserPost: Add SCG user
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ScgUserAddScgUserPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserAddScgUserPost200Response
// - error: Error seen or nil if none
func (u *UsersAPI) ScgUserAddScgUserPost(ctx *UserContext, requestBody *ScgUserAddScgUserPostRequest) (*http.Response, *ScgUserAddScgUserPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := u.client.newRequest(ctx, "POST", "/v5_0/users")
	request.body = requestBody
	request.authenticated = true
	out := &ScgUserAddScgUserPost200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserQueryScgUsersPostRequestAttributesSlice []*string

	ScgUserQueryScgUsersPostRequestFiltersSlice []*ScgUserQueryScgUsersPostRequestFilters

	ScgUserQueryScgUsersPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ScgUserQueryScgUsersPostRequestFullTextSearchFieldsSlice []*string

	ScgUserQueryScgUsersPostRequestFullTextSearch struct {
		Fields ScgUserQueryScgUsersPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                  `json:"type,omitempty"`
		Value  *string                                                  `json:"value,omitempty"`
	}

	ScgUserQueryScgUsersPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	ScgUserQueryScgUsersPostRequest struct {
		Attributes     ScgUserQueryScgUsersPostRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        ScgUserQueryScgUsersPostRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *ScgUserQueryScgUsersPostRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                           `json:"limit,omitempty"`
		Options        interface{}                                    `json:"options,omitempty"`
		Page           *int                                           `json:"page,omitempty"`
		SortInfo       *ScgUserQueryScgUsersPostRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                           `json:"start,omitempty"`
	}

	ScgUserQueryScgUsersPost200ResponseListSlice []*ScgUserQueryScgUsersPost200ResponseList

	ScgUserQueryScgUsersPost200ResponseList struct {
		CompanyName      *string `json:"companyName,omitempty"`
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		Email            *string `json:"email,omitempty"`
		Enabled          *int    `json:"enabled,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Passphrase       *string `json:"passphrase,omitempty"`
		Phone            *string `json:"phone,omitempty"`
		RealName         *string `json:"realName,omitempty"`
		TenantUUID       *string `json:"tenantUUID,omitempty"`
		Title            *string `json:"title,omitempty"`
		UserName         *string `json:"userName,omitempty"`
	}

	ScgUserQueryScgUsersPost200Response struct {
		Extra      *json.RawMessage                             `json:"extra,omitempty"`
		FirstIndex *int                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                        `json:"hasMore,omitempty"`
		List       ScgUserQueryScgUsersPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                         `json:"totalCount,omitempty"`
	}
)

// ScgUserQueryScgUsersPost: Query SCG users
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ScgUserQueryScgUsersPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserQueryScgUsersPost200Response
// - error: Error seen or nil if none
func (u *UsersAPI) ScgUserQueryScgUsersPost(ctx *UserContext, requestBody *ScgUserQueryScgUsersPostRequest) (*http.Response, *ScgUserQueryScgUsersPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := u.client.newRequest(ctx, "POST", "/v5_0/users/query")
	request.body = requestBody
	request.authenticated = true
	out := &ScgUserQueryScgUsersPost200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserDeleteScgUserDelete200Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ScgUserDeleteScgUserDelete: Delete SCG user
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - userId (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserDeleteScgUserDelete200Response
// - error: Error seen or nil if none
func (u *UsersAPI) ScgUserDeleteScgUserDelete(ctx *UserContext, userId string) (*http.Response, *ScgUserDeleteScgUserDelete200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(userId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userId\" failed validation check: %s", err)
	}
	request := u.client.newRequest(ctx, "DELETE", "/v5_0/users/{userId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userId": userId,
	}
	out := &ScgUserDeleteScgUserDelete200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserGetScgUserGet200Response struct {
		CompanyName      *string `json:"companyName,omitempty"`
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		Email            *string `json:"email,omitempty"`
		Enabled          *int    `json:"enabled,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Passphrase       *string `json:"passphrase,omitempty"`
		Phone            *string `json:"phone,omitempty"`
		RealName         *string `json:"realName,omitempty"`
		TenantUUID       *string `json:"tenantUUID,omitempty"`
		Title            *string `json:"title,omitempty"`
		UserName         *string `json:"userName,omitempty"`
	}
)

// ScgUserGetScgUserGet: Get SCG user
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - userId (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGetScgUserGet200Response
// - error: Error seen or nil if none
func (u *UsersAPI) ScgUserGetScgUserGet(ctx *UserContext, userId string) (*http.Response, *ScgUserGetScgUserGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(userId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userId\" failed validation check: %s", err)
	}
	request := u.client.newRequest(ctx, "GET", "/v5_0/users/{userId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userId": userId,
	}
	out := &ScgUserGetScgUserGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserUpdateScgUserPatchRequest struct {
		CompanyName      *string `json:"companyName,omitempty"`
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		Email            *string `json:"email,omitempty"`
		Enabled          *int    `json:"enabled,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Passphrase       *string `json:"passphrase,omitempty"`
		Phone            *string `json:"phone,omitempty"`
		RealName         *string `json:"realName,omitempty"`
		TenantUUID       *string `json:"tenantUUID,omitempty"`
		Title            *string `json:"title,omitempty"`
		UserName         *string `json:"userName,omitempty"`
	}

	ScgUserUpdateScgUserPatch200Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ScgUserUpdateScgUserPatch: Update SCG user
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - userId (UUIDv4)
// - requestBody: *ScgUserUpdateScgUserPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserUpdateScgUserPatch200Response
// - error: Error seen or nil if none
func (u *UsersAPI) ScgUserUpdateScgUserPatch(ctx *UserContext, userId string, requestBody *ScgUserUpdateScgUserPatchRequest) (*http.Response, *ScgUserUpdateScgUserPatch200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(userId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userId\" failed validation check: %s", err)
	}
	request := u.client.newRequest(ctx, "PATCH", "/v5_0/users/{userId}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userId": userId,
	}
	out := &ScgUserUpdateScgUserPatch200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
