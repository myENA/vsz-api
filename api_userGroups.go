package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type UserGroups struct {
	client *Client
}
type (
	ScgUserGroupDeleteScgUserGroupDelete1RequestIDListSlice []*string

	ScgUserGroupDeleteScgUserGroupDelete1Request struct {
		IDList ScgUserGroupDeleteScgUserGroupDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// ScgUserGroupDeleteScgUserGroupDelete1: Delete multiple SCG user group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ScgUserGroupDeleteScgUserGroupDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupDeleteScgUserGroupDelete1(ctx context.Context, requestBody *ScgUserGroupDeleteScgUserGroupDelete1Request) (*http.Response, []byte, error) {
	request := u.client.newRequest(ctx, "DELETE", "/v5_0/userGroups")
	request.body = requestBody
	request.authenticated = true

	return u.client.doRequest(request, 200, nil)
}

type (
	ScgUserGroupAddScgUserGroupPostRequestPermissionsSlice []*ScgUserGroupAddScgUserGroupPostRequestPermissions

	ScgUserGroupAddScgUserGroupPostRequestPermissionsIdsSlice []*string

	ScgUserGroupAddScgUserGroupPostRequestPermissions struct {
		Access   *string                                                   `json:"access,omitempty"`
		Display  *string                                                   `json:"display,omitempty"`
		Ids      ScgUserGroupAddScgUserGroupPostRequestPermissionsIdsSlice `json:"ids,omitempty"`
		Resource *string                                                   `json:"resource,omitempty"`
	}

	ScgUserGroupAddScgUserGroupPostRequestResourceGroupsSlice []*ScgUserGroupAddScgUserGroupPostRequestResourceGroups

	ScgUserGroupAddScgUserGroupPostRequestResourceGroups struct {
		ID   *string `json:"id,omitempty"`
		Type *string `json:"type,omitempty"`
	}

	ScgUserGroupAddScgUserGroupPostRequestUsersSlice []*ScgUserGroupAddScgUserGroupPostRequestUsers

	ScgUserGroupAddScgUserGroupPostRequestUsers struct {
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

	ScgUserGroupAddScgUserGroupPostRequest struct {
		CreateDateTime   *int                                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                   `json:"creatorUsername,omitempty"`
		Description      *string                                                   `json:"description,omitempty"`
		DomainID         *string                                                   `json:"domainId,omitempty"`
		ID               *string                                                   `json:"id,omitempty"`
		ModifiedDateTime *int                                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                                   `json:"modifierUsername,omitempty"`
		Name             *string                                                   `json:"name,omitempty"`
		Permissions      ScgUserGroupAddScgUserGroupPostRequestPermissionsSlice    `json:"permissions,omitempty"`
		ResourceGroups   ScgUserGroupAddScgUserGroupPostRequestResourceGroupsSlice `json:"resourceGroups,omitempty"`
		Role             *string                                                   `json:"role,omitempty"`
		TenantID         *string                                                   `json:"tenantId,omitempty"`
		Users            ScgUserGroupAddScgUserGroupPostRequestUsersSlice          `json:"users,omitempty"`
	}

	ScgUserGroupAddScgUserGroupPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ScgUserGroupAddScgUserGroupPost: Add SCG user group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ScgUserGroupAddScgUserGroupPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupAddScgUserGroupPost201Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupAddScgUserGroupPost(ctx context.Context, requestBody *ScgUserGroupAddScgUserGroupPostRequest) (*http.Response, *ScgUserGroupAddScgUserGroupPost201Response, error) {
	request := u.client.newRequest(ctx, "POST", "/v5_0/userGroups")
	request.body = requestBody
	request.authenticated = true

	out := &ScgUserGroupAddScgUserGroupPost201Response{}
	httpResponse, _, err := u.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListSlice []*ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseList

	ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListItemsSlice []*ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListItems

	ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListItems struct {
		Access   *string `json:"access,omitempty"`
		Display  *string `json:"display,omitempty"`
		Resource *string `json:"resource,omitempty"`
	}

	ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListItemsDescriptionSlice []*string

	ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseList struct {
		Access           *string                                                                                `json:"access,omitempty"`
		Display          *string                                                                                `json:"display,omitempty"`
		Items            ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListItemsSlice            `json:"items,omitempty"`
		ItemsDescription ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListItemsDescriptionSlice `json:"itemsDescription,omitempty"`
		Resource         *string                                                                                `json:"resource,omitempty"`
	}

	ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200Response struct {
		Extra      *json.RawMessage                                                       `json:"extra,omitempty"`
		FirstIndex *int                                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                  `json:"hasMore,omitempty"`
		List       ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                   `json:"totalCount,omitempty"`
	}
)

// ScgUserGroupGetPermittedCategoriesOfCurrentUserGet: Get permitted categories of current user
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupGetPermittedCategoriesOfCurrentUserGet(ctx context.Context) (*http.Response, *ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200Response, error) {
	request := u.client.newRequest(ctx, "GET", "/v5_0/userGroups/currentUser/permissionCategories")
	request.authenticated = true

	out := &ScgUserGroupGetPermittedCategoriesOfCurrentUserGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupQueryUserGroupsPostRequestAttributesSlice []*string

	ScgUserGroupQueryUserGroupsPostRequestExtraFiltersSlice []*ScgUserGroupQueryUserGroupsPostRequestExtraFilters

	ScgUserGroupQueryUserGroupsPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestExtraNotFiltersSlice []*ScgUserGroupQueryUserGroupsPostRequestExtraNotFilters

	ScgUserGroupQueryUserGroupsPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestFiltersSlice []*ScgUserGroupQueryUserGroupsPostRequestFilters

	ScgUserGroupQueryUserGroupsPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestFullTextSearchFieldsSlice []*string

	ScgUserGroupQueryUserGroupsPostRequestFullTextSearch struct {
		Fields ScgUserGroupQueryUserGroupsPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                         `json:"type,omitempty"`
		Value  *string                                                         `json:"value,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                             `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                           `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                             `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                             `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                           `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                             `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                             `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                             `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                             `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                             `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                             `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                           `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                             `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                           `json:"auth_type,omitempty"`
		ForwardingType                *string                                                           `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                           `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                           `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *ScgUserGroupQueryUserGroupsPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                           `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                             `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                             `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                             `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *ScgUserGroupQueryUserGroupsPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                           `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                           `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                           `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                           `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                           `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                           `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                           `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                           `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                           `json:"localUser_userSource,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPostRequest struct {
		Attributes      ScgUserGroupQueryUserGroupsPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                    `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                      `json:"expandDomains,omitempty"`
		ExtraFilters    ScgUserGroupQueryUserGroupsPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters ScgUserGroupQueryUserGroupsPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *ScgUserGroupQueryUserGroupsPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         ScgUserGroupQueryUserGroupsPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *ScgUserGroupQueryUserGroupsPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                       `json:"limit,omitempty"`
		Options         *ScgUserGroupQueryUserGroupsPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                       `json:"page,omitempty"`
		Query           *string                                                    `json:"query,omitempty"`
		SortInfo        *ScgUserGroupQueryUserGroupsPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                       `json:"start,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPost200ResponseListSlice []*ScgUserGroupQueryUserGroupsPost200ResponseList

	ScgUserGroupQueryUserGroupsPost200ResponseListPermissionsSlice []*ScgUserGroupQueryUserGroupsPost200ResponseListPermissions

	ScgUserGroupQueryUserGroupsPost200ResponseListPermissionsIdsSlice []*string

	ScgUserGroupQueryUserGroupsPost200ResponseListPermissions struct {
		Access   *string                                                           `json:"access,omitempty"`
		Display  *string                                                           `json:"display,omitempty"`
		Ids      ScgUserGroupQueryUserGroupsPost200ResponseListPermissionsIdsSlice `json:"ids,omitempty"`
		Resource *string                                                           `json:"resource,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPost200ResponseListResourceGroupsSlice []*ScgUserGroupQueryUserGroupsPost200ResponseListResourceGroups

	ScgUserGroupQueryUserGroupsPost200ResponseListResourceGroups struct {
		ID   *string `json:"id,omitempty"`
		Type *string `json:"type,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPost200ResponseListUsersSlice []*ScgUserGroupQueryUserGroupsPost200ResponseListUsers

	ScgUserGroupQueryUserGroupsPost200ResponseListUsers struct {
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

	ScgUserGroupQueryUserGroupsPost200ResponseList struct {
		CreateDateTime   *int                                                              `json:"createDateTime,omitempty"`
		CreatorID        *string                                                           `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                           `json:"creatorUsername,omitempty"`
		Description      *string                                                           `json:"description,omitempty"`
		DomainID         *string                                                           `json:"domainId,omitempty"`
		ID               *string                                                           `json:"id,omitempty"`
		ModifiedDateTime *int                                                              `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                           `json:"modifierId,omitempty"`
		ModifierUsername *string                                                           `json:"modifierUsername,omitempty"`
		Name             *string                                                           `json:"name,omitempty"`
		Permissions      ScgUserGroupQueryUserGroupsPost200ResponseListPermissionsSlice    `json:"permissions,omitempty"`
		ResourceGroups   ScgUserGroupQueryUserGroupsPost200ResponseListResourceGroupsSlice `json:"resourceGroups,omitempty"`
		Role             *string                                                           `json:"role,omitempty"`
		TenantID         *string                                                           `json:"tenantId,omitempty"`
		Users            ScgUserGroupQueryUserGroupsPost200ResponseListUsersSlice          `json:"users,omitempty"`
	}

	ScgUserGroupQueryUserGroupsPost200Response struct {
		Extra      *json.RawMessage                                    `json:"extra,omitempty"`
		FirstIndex *int                                                `json:"firstIndex,omitempty"`
		HasMore    *bool                                               `json:"hasMore,omitempty"`
		List       ScgUserGroupQueryUserGroupsPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                `json:"totalCount,omitempty"`
	}
)

// ScgUserGroupQueryUserGroupsPost: Query user groups
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ScgUserGroupQueryUserGroupsPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupQueryUserGroupsPost200Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupQueryUserGroupsPost(ctx context.Context, requestBody *ScgUserGroupQueryUserGroupsPostRequest) (*http.Response, *ScgUserGroupQueryUserGroupsPost200Response, error) {
	request := u.client.newRequest(ctx, "POST", "/v5_0/userGroups/query")
	request.body = requestBody
	request.authenticated = true

	out := &ScgUserGroupQueryUserGroupsPost200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupGetPreDefinedRolesGet200ResponseListSlice []*ScgUserGroupGetPreDefinedRolesGet200ResponseList

	ScgUserGroupGetPreDefinedRolesGet200ResponseList struct {
		Label *string `json:"label,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	ScgUserGroupGetPreDefinedRolesGet200Response struct {
		FirstIndex *int                                                  `json:"firstIndex,omitempty"`
		HasMore    *bool                                                 `json:"hasMore,omitempty"`
		List       ScgUserGroupGetPreDefinedRolesGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                  `json:"totalCount,omitempty"`
	}
)

// ScgUserGroupGetPreDefinedRolesGet: Get pre-defined roles
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupGetPreDefinedRolesGet200Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupGetPreDefinedRolesGet(ctx context.Context) (*http.Response, *ScgUserGroupGetPreDefinedRolesGet200Response, error) {
	request := u.client.newRequest(ctx, "GET", "/v5_0/userGroups/roles")
	request.authenticated = true

	out := &ScgUserGroupGetPreDefinedRolesGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListSlice []*ScgUserGroupGetPermissionItemsOfRoleGet200ResponseList

	ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListItemsSlice []*ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListItems

	ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListItems struct {
		Access   *string `json:"access,omitempty"`
		Display  *string `json:"display,omitempty"`
		Resource *string `json:"resource,omitempty"`
	}

	ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListItemsDescriptionSlice []*string

	ScgUserGroupGetPermissionItemsOfRoleGet200ResponseList struct {
		Access           *string                                                                     `json:"access,omitempty"`
		Display          *string                                                                     `json:"display,omitempty"`
		Items            ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListItemsSlice            `json:"items,omitempty"`
		ItemsDescription ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListItemsDescriptionSlice `json:"itemsDescription,omitempty"`
		Resource         *string                                                                     `json:"resource,omitempty"`
	}

	ScgUserGroupGetPermissionItemsOfRoleGet200Response struct {
		Extra      *json.RawMessage                                            `json:"extra,omitempty"`
		FirstIndex *int                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                       `json:"hasMore,omitempty"`
		List       ScgUserGroupGetPermissionItemsOfRoleGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                        `json:"totalCount,omitempty"`
	}
)

// ScgUserGroupGetPermissionItemsOfRoleGet: Get permission items of role
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - role (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupGetPermissionItemsOfRoleGet200Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupGetPermissionItemsOfRoleGet(ctx context.Context, role string) (*http.Response, *ScgUserGroupGetPermissionItemsOfRoleGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(role)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"role\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "GET", "/v5_0/userGroups/roles/{role}/permissions")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"role": role,
	}

	out := &ScgUserGroupGetPermissionItemsOfRoleGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupDeleteScgUserGroupDelete204Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ScgUserGroupDeleteScgUserGroupDelete: Delete SCG user group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - userGroupId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupDeleteScgUserGroupDelete204Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupDeleteScgUserGroupDelete(ctx context.Context, userGroupId string) (*http.Response, *ScgUserGroupDeleteScgUserGroupDelete204Response, error) {
	var err error

	err = validators.StrNotEmpty(userGroupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userGroupId\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "DELETE", "/v5_0/userGroups/{userGroupId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userGroupId": userGroupId,
	}

	out := &ScgUserGroupDeleteScgUserGroupDelete204Response{}
	httpResponse, _, err := u.client.doRequest(request, 204, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupGetScgUserGroupGet200ResponsePermissionsSlice []*ScgUserGroupGetScgUserGroupGet200ResponsePermissions

	ScgUserGroupGetScgUserGroupGet200ResponsePermissionsIdsSlice []*string

	ScgUserGroupGetScgUserGroupGet200ResponsePermissions struct {
		Access   *string                                                      `json:"access,omitempty"`
		Display  *string                                                      `json:"display,omitempty"`
		Ids      ScgUserGroupGetScgUserGroupGet200ResponsePermissionsIdsSlice `json:"ids,omitempty"`
		Resource *string                                                      `json:"resource,omitempty"`
	}

	ScgUserGroupGetScgUserGroupGet200ResponseResourceGroupsSlice []*ScgUserGroupGetScgUserGroupGet200ResponseResourceGroups

	ScgUserGroupGetScgUserGroupGet200ResponseResourceGroups struct {
		ID   *string `json:"id,omitempty"`
		Type *string `json:"type,omitempty"`
	}

	ScgUserGroupGetScgUserGroupGet200ResponseUsersSlice []*ScgUserGroupGetScgUserGroupGet200ResponseUsers

	ScgUserGroupGetScgUserGroupGet200ResponseUsers struct {
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

	ScgUserGroupGetScgUserGroupGet200Response struct {
		CreateDateTime   *int                                                         `json:"createDateTime,omitempty"`
		CreatorID        *string                                                      `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                      `json:"creatorUsername,omitempty"`
		Description      *string                                                      `json:"description,omitempty"`
		DomainID         *string                                                      `json:"domainId,omitempty"`
		ID               *string                                                      `json:"id,omitempty"`
		ModifiedDateTime *int                                                         `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                      `json:"modifierId,omitempty"`
		ModifierUsername *string                                                      `json:"modifierUsername,omitempty"`
		Name             *string                                                      `json:"name,omitempty"`
		Permissions      ScgUserGroupGetScgUserGroupGet200ResponsePermissionsSlice    `json:"permissions,omitempty"`
		ResourceGroups   ScgUserGroupGetScgUserGroupGet200ResponseResourceGroupsSlice `json:"resourceGroups,omitempty"`
		Role             *string                                                      `json:"role,omitempty"`
		TenantID         *string                                                      `json:"tenantId,omitempty"`
		Users            ScgUserGroupGetScgUserGroupGet200ResponseUsersSlice          `json:"users,omitempty"`
	}
)

// ScgUserGroupGetScgUserGroupGet: Get SCG user group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - userGroupId (string)
//
// Optional Parameter Map:
// - includeUsers (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ScgUserGroupGetScgUserGroupGet200Response
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupGetScgUserGroupGet(ctx context.Context, userGroupId string, optionalParams map[string]string) (*http.Response, *ScgUserGroupGetScgUserGroupGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(userGroupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userGroupId\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "GET", "/v5_0/userGroups/{userGroupId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userGroupId": userGroupId,
	}
	request.queryParameters = map[string]string{
		"includeUsers": optionalParams["includeUsers"],
	}

	out := &ScgUserGroupGetScgUserGroupGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ScgUserGroupUpdateUserGroupsPatchRequestPermissionsSlice []*ScgUserGroupUpdateUserGroupsPatchRequestPermissions

	ScgUserGroupUpdateUserGroupsPatchRequestPermissions struct {
		Access   *string `json:"access,omitempty"`
		Display  *string `json:"display,omitempty"`
		Resource *string `json:"resource,omitempty"`
	}

	ScgUserGroupUpdateUserGroupsPatchRequestResourceGroupsSlice []*ScgUserGroupUpdateUserGroupsPatchRequestResourceGroups

	ScgUserGroupUpdateUserGroupsPatchRequestResourceGroups struct {
		ID   *string `json:"id,omitempty"`
		Type *string `json:"type,omitempty"`
	}

	ScgUserGroupUpdateUserGroupsPatchRequestUsersSlice []*ScgUserGroupUpdateUserGroupsPatchRequestUsers

	ScgUserGroupUpdateUserGroupsPatchRequestUsers struct {
		CompanyName interface{} `json:"companyName,omitempty"`
		DomainID    *string     `json:"domainId,omitempty"`
		Email       *string     `json:"email,omitempty"`
		Enabled     *int        `json:"enabled,omitempty"`
		ID          *string     `json:"id,omitempty"`
		Passphrase  interface{} `json:"passphrase,omitempty"`
		Phone       interface{} `json:"phone,omitempty"`
		RealName    *string     `json:"realName,omitempty"`
		TenantUUID  *string     `json:"tenantUUID,omitempty"`
		Title       interface{} `json:"title,omitempty"`
		UserName    *string     `json:"userName,omitempty"`
	}

	ScgUserGroupUpdateUserGroupsPatchRequest struct {
		Description    *string                                                     `json:"description,omitempty"`
		ID             *string                                                     `json:"id,omitempty"`
		Name           *string                                                     `json:"name,omitempty"`
		Permissions    ScgUserGroupUpdateUserGroupsPatchRequestPermissionsSlice    `json:"permissions,omitempty"`
		ResourceGroups ScgUserGroupUpdateUserGroupsPatchRequestResourceGroupsSlice `json:"resourceGroups,omitempty"`
		Role           *string                                                     `json:"role,omitempty"`
		Users          ScgUserGroupUpdateUserGroupsPatchRequestUsersSlice          `json:"users,omitempty"`
	}
)

// ScgUserGroupUpdateUserGroupsPatch: Update user groups
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - userGroupId (string)
// - requestBody: *ScgUserGroupUpdateUserGroupsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupUpdateUserGroupsPatch(ctx context.Context, userGroupId string, requestBody *ScgUserGroupUpdateUserGroupsPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(userGroupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userGroupId\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "PATCH", "/v5_0/userGroups/{userGroupId}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userGroupId": userGroupId,
	}

	return u.client.doRequest(request, 204, nil)
}

type (
	ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequestSlice []*ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequest

	ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequestIdsSlice []*string

	ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequest struct {
		Access   *string                                                       `json:"access,omitempty"`
		Display  *string                                                       `json:"display,omitempty"`
		Ids      ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequestIdsSlice `json:"ids,omitempty"`
		Resource *string                                                       `json:"resource,omitempty"`
	}
)

// ScgUserGroupUpdatePermissionsOfUserGroupsPatch: Update permissions of user groups
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - userGroupId (string)
// - requestBody: *ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupUpdatePermissionsOfUserGroupsPatch(ctx context.Context, userGroupId string, requestBody ScgUserGroupUpdatePermissionsOfUserGroupsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(userGroupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userGroupId\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "PATCH", "/v5_0/userGroups/{userGroupId}/permissions")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userGroupId": userGroupId,
	}

	return u.client.doRequest(request, 204, nil)
}

type (
	ScgUserGroupUpdatePermissionScopeOfUserGroupsPatchRequestSlice []*ScgUserGroupUpdatePermissionScopeOfUserGroupsPatchRequest

	ScgUserGroupUpdatePermissionScopeOfUserGroupsPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Type *string `json:"type,omitempty"`
	}
)

// ScgUserGroupUpdatePermissionScopeOfUserGroupsPatch: Update permission scope of user groups
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - userGroupId (string)
// - requestBody: *ScgUserGroupUpdatePermissionScopeOfUserGroupsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupUpdatePermissionScopeOfUserGroupsPatch(ctx context.Context, userGroupId string, requestBody ScgUserGroupUpdatePermissionScopeOfUserGroupsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(userGroupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userGroupId\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "PATCH", "/v5_0/userGroups/{userGroupId}/resourceGroups")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userGroupId": userGroupId,
	}

	return u.client.doRequest(request, 204, nil)
}

type (
	ScgUserGroupUpdateUserListOfUserGroupsPatchRequestSlice []*ScgUserGroupUpdateUserListOfUserGroupsPatchRequest

	ScgUserGroupUpdateUserListOfUserGroupsPatchRequest struct {
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

// ScgUserGroupUpdateUserListOfUserGroupsPatch: Update user list of user groups
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - userGroupId (string)
// - requestBody: *ScgUserGroupUpdateUserListOfUserGroupsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (u *UserGroups) ScgUserGroupUpdateUserListOfUserGroupsPatch(ctx context.Context, userGroupId string, requestBody ScgUserGroupUpdateUserListOfUserGroupsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(userGroupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"userGroupId\" failed validation check: %s", err)
	}

	request := u.client.newRequest(ctx, "PATCH", "/v5_0/userGroups/{userGroupId}/users")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"userGroupId": userGroupId,
	}

	return u.client.doRequest(request, 204, nil)
}
