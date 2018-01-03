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

type VLANPoolings struct {
	client *Client
}
type (
	VlanpoolingBulkDeleteVlanPoolingDeleteRequestIDListSlice []*string

	VlanpoolingBulkDeleteVlanPoolingDeleteRequest struct {
		IDList VlanpoolingBulkDeleteVlanPoolingDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// VlanpoolingBulkDeleteVlanPoolingDelete: Use this API command to bulk delete VLAN pooling.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *VlanpoolingBulkDeleteVlanPoolingDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (v *VLANPoolings) VlanpoolingBulkDeleteVlanPoolingDelete(ctx context.Context, requestBody *VlanpoolingBulkDeleteVlanPoolingDeleteRequest) (*http.Response, []byte, error) {
	request := v.client.newRequest(ctx, "DELETE", "/v5_0/vlanpoolings")
	request.body = requestBody
	request.authenticated = true

	return v.client.doRequest(request, 204, nil)
}

type (
	VlanpoolingCreatePostRequest struct {
		Algo        *string `json:"algo,omitempty"`
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		Name        *string `json:"name,omitempty"`
		Pool        *string `json:"pool,omitempty"`
	}

	VlanpoolingCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// VlanpoolingCreatePost: Use this API command to create new VLAN pooling.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *VlanpoolingCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *VlanpoolingCreatePost201Response
// - error: Error seen or nil if none
func (v *VLANPoolings) VlanpoolingCreatePost(ctx context.Context, requestBody *VlanpoolingCreatePostRequest) (*http.Response, *VlanpoolingCreatePost201Response, error) {
	request := v.client.newRequest(ctx, "POST", "/v5_0/vlanpoolings")
	request.body = requestBody
	request.authenticated = true

	out := &VlanpoolingCreatePost201Response{}
	httpResponse, _, err := v.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	VlanpoolingRetrieveListPostRequestAttributesSlice []*string

	VlanpoolingRetrieveListPostRequestExtraFiltersSlice []*VlanpoolingRetrieveListPostRequestExtraFilters

	VlanpoolingRetrieveListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestExtraNotFiltersSlice []*VlanpoolingRetrieveListPostRequestExtraNotFilters

	VlanpoolingRetrieveListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestFiltersSlice []*VlanpoolingRetrieveListPostRequestFilters

	VlanpoolingRetrieveListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestFullTextSearchFieldsSlice []*string

	VlanpoolingRetrieveListPostRequestFullTextSearch struct {
		Fields VlanpoolingRetrieveListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                     `json:"type,omitempty"`
		Value  *string                                                     `json:"value,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                         `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                       `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                         `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                         `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                       `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                         `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                         `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                         `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                         `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                         `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                         `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                       `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                         `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                       `json:"auth_type,omitempty"`
		ForwardingType                *string                                                       `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                       `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                       `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *VlanpoolingRetrieveListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *VlanpoolingRetrieveListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                       `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                       `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                       `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                       `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                       `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                       `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                       `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                       `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                       `json:"localUser_userSource,omitempty"`
	}

	VlanpoolingRetrieveListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	VlanpoolingRetrieveListPostRequest struct {
		Attributes      VlanpoolingRetrieveListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    VlanpoolingRetrieveListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters VlanpoolingRetrieveListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *VlanpoolingRetrieveListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         VlanpoolingRetrieveListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *VlanpoolingRetrieveListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                               `json:"limit,omitempty"`
		Options         *VlanpoolingRetrieveListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                   `json:"page,omitempty"`
		Query           *string                                                `json:"query,omitempty"`
		SortInfo        *VlanpoolingRetrieveListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                               `json:"start,omitempty"`
	}

	VlanpoolingRetrieveListPost200ResponseListSlice []*VlanpoolingRetrieveListPost200ResponseList

	VlanpoolingRetrieveListPost200ResponseList struct {
		Algo        *string `json:"algo,omitempty"`
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Pool        *string `json:"pool,omitempty"`
	}

	VlanpoolingRetrieveListPost200Response struct {
		Extra      *json.RawMessage                                `json:"extra,omitempty"`
		FirstIndex *int                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                           `json:"hasMore,omitempty"`
		List       VlanpoolingRetrieveListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                            `json:"totalCount,omitempty"`
	}
)

// VlanpoolingRetrieveListPost: Use this API command to retrieve a list of VLAN poolings.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *VlanpoolingRetrieveListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *VlanpoolingRetrieveListPost200Response
// - error: Error seen or nil if none
func (v *VLANPoolings) VlanpoolingRetrieveListPost(ctx context.Context, requestBody *VlanpoolingRetrieveListPostRequest) (*http.Response, *VlanpoolingRetrieveListPost200Response, error) {
	request := v.client.newRequest(ctx, "POST", "/v5_0/vlanpoolings/query")
	request.body = requestBody
	request.authenticated = true

	out := &VlanpoolingRetrieveListPost200Response{}
	httpResponse, _, err := v.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// VlanpoolingDeleteDelete: Use this API command to delete VLAN pooling
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (v *VLANPoolings) VlanpoolingDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := v.client.newRequest(ctx, "DELETE", "/v5_0/vlanpoolings/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return v.client.doRequest(request, 204, nil)
}

type (
	VlanpoolingRetrieveGet200Response struct {
		Algo        *string `json:"algo,omitempty"`
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Pool        *string `json:"pool,omitempty"`
	}
)

// VlanpoolingRetrieveGet: Use this API command to retrieve VLAN pooling.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *VlanpoolingRetrieveGet200Response
// - error: Error seen or nil if none
func (v *VLANPoolings) VlanpoolingRetrieveGet(ctx context.Context, id string) (*http.Response, *VlanpoolingRetrieveGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := v.client.newRequest(ctx, "GET", "/v5_0/vlanpoolings/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &VlanpoolingRetrieveGet200Response{}
	httpResponse, _, err := v.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	VlanpoolingModifyBasicPatchRequest struct {
		Algo        *string `json:"algo,omitempty"`
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		Name        *string `json:"name,omitempty"`
		Pool        *string `json:"pool,omitempty"`
	}
)

// VlanpoolingModifyBasicPatch: Use this API command to modify the basic information on VLAN pooling
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *VlanpoolingModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (v *VLANPoolings) VlanpoolingModifyBasicPatch(ctx context.Context, id string, requestBody *VlanpoolingModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := v.client.newRequest(ctx, "PATCH", "/v5_0/vlanpoolings/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return v.client.doRequest(request, 204, nil)
}
