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

type PrecedenceAPI struct {
	client *Client
}
type (
	PrecedenceProfileBulkDeletePrecedenceProfileDeleteRequestIDListSlice []*string

	PrecedenceProfileBulkDeletePrecedenceProfileDeleteRequest struct {
		IDList PrecedenceProfileBulkDeletePrecedenceProfileDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// PrecedenceProfileBulkDeletePrecedenceProfileDelete: Use this API command to Bulk Delete Precedence Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *PrecedenceProfileBulkDeletePrecedenceProfileDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileBulkDeletePrecedenceProfileDelete(ctx *UserContext, requestBody *PrecedenceProfileBulkDeletePrecedenceProfileDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/precedence")
	request.body = requestBody
	request.authenticated = true
	return p.client.doRequest(request, 204, nil)
}

type (
	PrecedenceProfileGetPrecedenceProfileListGet200ResponseListSlice []*PrecedenceProfileGetPrecedenceProfileListGet200ResponseList

	PrecedenceProfileGetPrecedenceProfileListGet200ResponseListRateLimitingPrecedenceSlice []*PrecedenceProfileGetPrecedenceProfileListGet200ResponseListRateLimitingPrecedence

	PrecedenceProfileGetPrecedenceProfileListGet200ResponseListRateLimitingPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileGetPrecedenceProfileListGet200ResponseListVlanPrecedenceSlice []*PrecedenceProfileGetPrecedenceProfileListGet200ResponseListVlanPrecedence

	PrecedenceProfileGetPrecedenceProfileListGet200ResponseListVlanPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileGetPrecedenceProfileListGet200ResponseList struct {
		DomainID               *string                                                                                `json:"domainId,omitempty"`
		ID                     *string                                                                                `json:"id,omitempty"`
		Name                   *string                                                                                `json:"name,omitempty"`
		RateLimitingPrecedence PrecedenceProfileGetPrecedenceProfileListGet200ResponseListRateLimitingPrecedenceSlice `json:"rateLimitingPrecedence,omitempty"`
		VlanPrecedence         PrecedenceProfileGetPrecedenceProfileListGet200ResponseListVlanPrecedenceSlice         `json:"vlanPrecedence,omitempty"`
	}

	PrecedenceProfileGetPrecedenceProfileListGet200Response struct {
		Extra      *json.RawMessage                                                 `json:"extra,omitempty"`
		FirstIndex *int                                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                                            `json:"hasMore,omitempty"`
		List       PrecedenceProfileGetPrecedenceProfileListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                             `json:"totalCount,omitempty"`
	}
)

// PrecedenceProfileGetPrecedenceProfileListGet: Use this API command to Get Precedence Profile list.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *PrecedenceProfileGetPrecedenceProfileListGet200Response
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileGetPrecedenceProfileListGet(ctx *UserContext, optionalParams map[string]string) (*http.Response, *PrecedenceProfileGetPrecedenceProfileListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	index, ok := optionalParams["index"]
	if ok {
		err = validators.StrIsPositiveInt(index)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"index\" failed validation check: %s", err)
		}
	} else {
		index = "0"
	}
	listSize, ok := optionalParams["listSize"]
	if ok {
		err = validators.StrIsPositiveInt(listSize)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"listSize\" failed validation check: %s", err)
		}
	} else {
		listSize = "100"
	}
	request := p.client.newRequest(ctx, "GET", "/v5_0/precedence")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}
	out := &PrecedenceProfileGetPrecedenceProfileListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	PrecedenceProfileCreatePrecedenceProfilePostRequestRateLimitingPrecedenceSlice []*PrecedenceProfileCreatePrecedenceProfilePostRequestRateLimitingPrecedence

	PrecedenceProfileCreatePrecedenceProfilePostRequestRateLimitingPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileCreatePrecedenceProfilePostRequestVlanPrecedenceSlice []*PrecedenceProfileCreatePrecedenceProfilePostRequestVlanPrecedence

	PrecedenceProfileCreatePrecedenceProfilePostRequestVlanPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileCreatePrecedenceProfilePostRequest struct {
		DomainID               *string                                                                        `json:"domainId,omitempty"`
		Name                   *string                                                                        `json:"name,omitempty"`
		RateLimitingPrecedence PrecedenceProfileCreatePrecedenceProfilePostRequestRateLimitingPrecedenceSlice `json:"rateLimitingPrecedence,omitempty"`
		VlanPrecedence         PrecedenceProfileCreatePrecedenceProfilePostRequestVlanPrecedenceSlice         `json:"vlanPrecedence,omitempty"`
	}

	PrecedenceProfileCreatePrecedenceProfilePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// PrecedenceProfileCreatePrecedenceProfilePost: Use this API command to create Precedence Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *PrecedenceProfileCreatePrecedenceProfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *PrecedenceProfileCreatePrecedenceProfilePost201Response
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileCreatePrecedenceProfilePost(ctx *UserContext, requestBody *PrecedenceProfileCreatePrecedenceProfilePostRequest) (*http.Response, *PrecedenceProfileCreatePrecedenceProfilePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := p.client.newRequest(ctx, "POST", "/v5_0/precedence")
	request.body = requestBody
	request.authenticated = true
	out := &PrecedenceProfileCreatePrecedenceProfilePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	PrecedenceProfileQueryPrecedenceProfilePostRequestAttributesSlice []*string

	PrecedenceProfileQueryPrecedenceProfilePostRequestExtraFiltersSlice []*PrecedenceProfileQueryPrecedenceProfilePostRequestExtraFilters

	PrecedenceProfileQueryPrecedenceProfilePostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestExtraNotFiltersSlice []*PrecedenceProfileQueryPrecedenceProfilePostRequestExtraNotFilters

	PrecedenceProfileQueryPrecedenceProfilePostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestFiltersSlice []*PrecedenceProfileQueryPrecedenceProfilePostRequestFilters

	PrecedenceProfileQueryPrecedenceProfilePostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestFullTextSearchFieldsSlice []*string

	PrecedenceProfileQueryPrecedenceProfilePostRequestFullTextSearch struct {
		Fields PrecedenceProfileQueryPrecedenceProfilePostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                     `json:"type,omitempty"`
		Value  *string                                                                     `json:"value,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                         `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                       `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                         `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                         `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                       `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                         `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                         `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                         `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                         `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                         `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                         `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                       `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                         `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                       `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                       `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                       `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                       `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *PrecedenceProfileQueryPrecedenceProfilePostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *PrecedenceProfileQueryPrecedenceProfilePostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                       `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                       `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                       `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                       `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                       `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                       `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                       `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                       `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                       `json:"localUser_userSource,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePostRequest struct {
		Attributes      PrecedenceProfileQueryPrecedenceProfilePostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    PrecedenceProfileQueryPrecedenceProfilePostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters PrecedenceProfileQueryPrecedenceProfilePostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *PrecedenceProfileQueryPrecedenceProfilePostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         PrecedenceProfileQueryPrecedenceProfilePostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *PrecedenceProfileQueryPrecedenceProfilePostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                   `json:"limit,omitempty"`
		Options         *PrecedenceProfileQueryPrecedenceProfilePostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                   `json:"page,omitempty"`
		Query           *string                                                                `json:"query,omitempty"`
		SortInfo        *PrecedenceProfileQueryPrecedenceProfilePostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                   `json:"start,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePost200ResponseListSlice []*PrecedenceProfileQueryPrecedenceProfilePost200ResponseList

	PrecedenceProfileQueryPrecedenceProfilePost200ResponseListRateLimitingPrecedenceSlice []*PrecedenceProfileQueryPrecedenceProfilePost200ResponseListRateLimitingPrecedence

	PrecedenceProfileQueryPrecedenceProfilePost200ResponseListRateLimitingPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePost200ResponseListVlanPrecedenceSlice []*PrecedenceProfileQueryPrecedenceProfilePost200ResponseListVlanPrecedence

	PrecedenceProfileQueryPrecedenceProfilePost200ResponseListVlanPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePost200ResponseList struct {
		DomainID               *string                                                                               `json:"domainId,omitempty"`
		ID                     *string                                                                               `json:"id,omitempty"`
		Name                   *string                                                                               `json:"name,omitempty"`
		RateLimitingPrecedence PrecedenceProfileQueryPrecedenceProfilePost200ResponseListRateLimitingPrecedenceSlice `json:"rateLimitingPrecedence,omitempty"`
		VlanPrecedence         PrecedenceProfileQueryPrecedenceProfilePost200ResponseListVlanPrecedenceSlice         `json:"vlanPrecedence,omitempty"`
	}

	PrecedenceProfileQueryPrecedenceProfilePost200Response struct {
		Extra      *json.RawMessage                                                `json:"extra,omitempty"`
		FirstIndex *int                                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                                           `json:"hasMore,omitempty"`
		List       PrecedenceProfileQueryPrecedenceProfilePost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                            `json:"totalCount,omitempty"`
	}
)

// PrecedenceProfileQueryPrecedenceProfilePost: Use this API command to query Precedence Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *PrecedenceProfileQueryPrecedenceProfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *PrecedenceProfileQueryPrecedenceProfilePost200Response
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileQueryPrecedenceProfilePost(ctx *UserContext, requestBody *PrecedenceProfileQueryPrecedenceProfilePostRequest) (*http.Response, *PrecedenceProfileQueryPrecedenceProfilePost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := p.client.newRequest(ctx, "POST", "/v5_0/precedence/query")
	request.body = requestBody
	request.authenticated = true
	out := &PrecedenceProfileQueryPrecedenceProfilePost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// PrecedenceProfileDeletePrecedenceProfileByProfileSIdDelete: Use this API command to Delete Precedence Profile by profile's ID.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileDeletePrecedenceProfileByProfileSIdDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/precedence/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return p.client.doRequest(request, 204, nil)
}

type (
	PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseRateLimitingPrecedenceSlice []*PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseRateLimitingPrecedence

	PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseRateLimitingPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseVlanPrecedenceSlice []*PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseVlanPrecedence

	PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseVlanPrecedence struct {
		Name     *string  `json:"name,omitempty"`
		Priority *float64 `json:"priority,omitempty"`
	}

	PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200Response struct {
		DomainID               *string                                                                                    `json:"domainId,omitempty"`
		Name                   *string                                                                                    `json:"name,omitempty"`
		RateLimitingPrecedence PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseRateLimitingPrecedenceSlice `json:"rateLimitingPrecedence,omitempty"`
		VlanPrecedence         PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200ResponseVlanPrecedenceSlice         `json:"vlanPrecedence,omitempty"`
	}
)

// PrecedenceProfileGetPrecedenceProfileByProfileSIdGet: Use this API command to Get Precedence Profile by profile's ID.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200Response
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileGetPrecedenceProfileByProfileSIdGet(ctx *UserContext, id string) (*http.Response, *PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := p.client.newRequest(ctx, "GET", "/v5_0/precedence/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &PrecedenceProfileGetPrecedenceProfileByProfileSIdGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	PrecedenceProfileModifyPrecedenceProfileByProfileSIdPatchRequest struct {
		DomainID *string `json:"domainId,omitempty"`
		Name     *string `json:"name,omitempty"`
	}
)

// PrecedenceProfileModifyPrecedenceProfileByProfileSIdPatch: Use this API command to Modify Precedence Profile by profile's ID.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *PrecedenceProfileModifyPrecedenceProfileByProfileSIdPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PrecedenceAPI) PrecedenceProfileModifyPrecedenceProfileByProfileSIdPatch(ctx *UserContext, id string, requestBody *PrecedenceProfileModifyPrecedenceProfileByProfileSIdPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := p.client.newRequest(ctx, "PATCH", "/v5_0/precedence/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return p.client.doRequest(request, 200, nil)
}
