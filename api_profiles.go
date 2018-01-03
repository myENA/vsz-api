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

type Profiles struct {
	client *Client
}
type (
	AccountingProfileDeleteDelete1RequestIDListSlice []*string

	AccountingProfileDeleteDelete1Request struct {
		IDList AccountingProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// AccountingProfileDeleteDelete1: Use this API command to delete a list of accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AccountingProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileDeleteDelete1(ctx context.Context, requestBody *AccountingProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/acct")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	AccountingProfileRetrieveListGet200ResponseListSlice []*AccountingProfileRetrieveListGet200ResponseList

	AccountingProfileRetrieveListGet200ResponseListRealmMappingsSlice []*AccountingProfileRetrieveListGet200ResponseListRealmMappings

	AccountingProfileRetrieveListGet200ResponseListRealmMappings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	AccountingProfileRetrieveListGet200ResponseList struct {
		CreateDateTime   *int                                                              `json:"createDateTime,omitempty"`
		CreatorID        *string                                                           `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                           `json:"creatorUsername,omitempty"`
		Description      *string                                                           `json:"description,omitempty"`
		DomainID         *string                                                           `json:"domainId,omitempty"`
		ID               *string                                                           `json:"id,omitempty"`
		ModifiedDateTime *int                                                              `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                           `json:"modifierId,omitempty"`
		ModifierUsername *string                                                           `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                           `json:"mvnoId,omitempty"`
		Name             *string                                                           `json:"name,omitempty"`
		RealmMappings    AccountingProfileRetrieveListGet200ResponseListRealmMappingsSlice `json:"realmMappings,omitempty"`
	}

	AccountingProfileRetrieveListGet200Response struct {
		Extra      *json.RawMessage                                     `json:"extra,omitempty"`
		FirstIndex *int                                                 `json:"firstIndex,omitempty"`
		HasMore    *bool                                                `json:"hasMore,omitempty"`
		List       AccountingProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                 `json:"totalCount,omitempty"`
	}
)

// AccountingProfileRetrieveListGet: Use this API command to retrieve a list of accounting profiles.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileRetrieveListGet(ctx context.Context) (*http.Response, *AccountingProfileRetrieveListGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/acct")
	request.authenticated = true

	out := &AccountingProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AccountingProfileCreatePostRequestRealmMappingsSlice []*AccountingProfileCreatePostRequestRealmMappings

	AccountingProfileCreatePostRequestRealmMappings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	AccountingProfileCreatePostRequest struct {
		Description   *string                                              `json:"description,omitempty"`
		DomainID      *string                                              `json:"domainId,omitempty"`
		MvnoID        *string                                              `json:"mvnoId,omitempty"`
		Name          *string                                              `json:"name,omitempty"`
		RealmMappings AccountingProfileCreatePostRequestRealmMappingsSlice `json:"realmMappings,omitempty"`
	}

	AccountingProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AccountingProfileCreatePost: Use this API command to create a new accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AccountingProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingProfileCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileCreatePost(ctx context.Context, requestBody *AccountingProfileCreatePostRequest) (*http.Response, *AccountingProfileCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/acct")
	request.body = requestBody
	request.authenticated = true

	out := &AccountingProfileCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AccountingProfileClonePostRequest struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}

	AccountingProfileClonePost200Response struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}
)

// AccountingProfileClonePost: Use this API command to clone an accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *AccountingProfileClonePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingProfileClonePost200Response
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileClonePost(ctx context.Context, id string, requestBody *AccountingProfileClonePostRequest) (*http.Response, *AccountingProfileClonePost200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/acct/clone/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &AccountingProfileClonePost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestAttributesSlice []*string

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraFiltersSlice []*AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraFilters

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraNotFiltersSlice []*AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraNotFilters

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFiltersSlice []*AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFilters

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFullTextSearch struct {
		Fields AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                            `json:"type,omitempty"`
		Value  *string                                                                                            `json:"value,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                              `json:"localUser_userSource,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequest struct {
		Attributes      AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                          `json:"limit,omitempty"`
		Options         *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                          `json:"page,omitempty"`
		Query           *string                                                                                       `json:"query,omitempty"`
		SortInfo        *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                          `json:"start,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseListSlice []*AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseList

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseListRealmMappingsSlice []*AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseListRealmMappings

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseListRealmMappings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseList struct {
		CreateDateTime   *int                                                                                                `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                             `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                             `json:"creatorUsername,omitempty"`
		Description      *string                                                                                             `json:"description,omitempty"`
		DomainID         *string                                                                                             `json:"domainId,omitempty"`
		ID               *string                                                                                             `json:"id,omitempty"`
		ModifiedDateTime *int                                                                                                `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                             `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                             `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                                             `json:"mvnoId,omitempty"`
		Name             *string                                                                                             `json:"name,omitempty"`
		RealmMappings    AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseListRealmMappingsSlice `json:"realmMappings,omitempty"`
	}

	AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                       `json:"extra,omitempty"`
		FirstIndex *int                                                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                  `json:"hasMore,omitempty"`
		List       AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                   `json:"totalCount,omitempty"`
	}
)

// AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost: Use this API command to retrieve a list of accounting profiles by query criteria.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost(ctx context.Context, requestBody *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPostRequest) (*http.Response, *AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/acct/query")
	request.body = requestBody
	request.authenticated = true

	out := &AccountingProfileRetrieveListAccountingProfilesByQueryCritariaPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AccountingProfileDeleteDelete: Use this API command to delete an accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Accounting Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/acct/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	AccountingProfileRetrieveGet200ResponseRealmMappingsSlice []*AccountingProfileRetrieveGet200ResponseRealmMappings

	AccountingProfileRetrieveGet200ResponseRealmMappings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	AccountingProfileRetrieveGet200Response struct {
		CreateDateTime   *int                                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                   `json:"creatorUsername,omitempty"`
		Description      *string                                                   `json:"description,omitempty"`
		DomainID         *string                                                   `json:"domainId,omitempty"`
		ID               *string                                                   `json:"id,omitempty"`
		ModifiedDateTime *int                                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                                   `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                   `json:"mvnoId,omitempty"`
		Name             *string                                                   `json:"name,omitempty"`
		RealmMappings    AccountingProfileRetrieveGet200ResponseRealmMappingsSlice `json:"realmMappings,omitempty"`
	}
)

// AccountingProfileRetrieveGet: Use this API command to retrieve an accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Accounting Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *AccountingProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/acct/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &AccountingProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AccountingProfileModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		MvnoID      *string `json:"mvnoId,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// AccountingProfileModifyBasicPatch: Use this API command to modify the basic information of an accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Accounting Profile ID
// - requestBody: *AccountingProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileModifyBasicPatch(ctx context.Context, id string, requestBody *AccountingProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/acct/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	AccountingProfileModifyRealmMappingsPatchRequestSlice []*AccountingProfileModifyRealmMappingsPatchRequest

	AccountingProfileModifyRealmMappingsPatchRequest struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}
)

// AccountingProfileModifyRealmMappingsPatch: Use this API command to modify accounting service per realm mappings of an accounting profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Accounting Profile ID
// - requestBody: *AccountingProfileModifyRealmMappingsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AccountingProfileModifyRealmMappingsPatch(ctx context.Context, id string, requestBody AccountingProfileModifyRealmMappingsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/acct/{id}/realmMappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	AuthenticationProfileDeleteDelete1RequestIDListSlice []*string

	AuthenticationProfileDeleteDelete1Request struct {
		IDList AuthenticationProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// AuthenticationProfileDeleteDelete1: Use this API command to delete a list of authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AuthenticationProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileDeleteDelete1(ctx context.Context, requestBody *AuthenticationProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/auth")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	AuthenticationProfileRetrieveListGet200ResponseListSlice []*AuthenticationProfileRetrieveListGet200ResponseList

	AuthenticationProfileRetrieveListGet200ResponseListRealmMappingsSlice []*AuthenticationProfileRetrieveListGet200ResponseListRealmMappings

	AuthenticationProfileRetrieveListGet200ResponseListRealmMappings struct {
		AuthorizationMethod *string  `json:"authorizationMethod,omitempty"`
		DynamicVlanID       *float64 `json:"dynamicVlanId,omitempty"`
		HostedAaaEnabled    *bool    `json:"hostedAaaEnabled,omitempty"`
		ID                  *string  `json:"id,omitempty"`
		Name                *string  `json:"name,omitempty"`
		Realm               *string  `json:"realm,omitempty"`
		ServiceType         *string  `json:"serviceType,omitempty"`
	}

	AuthenticationProfileRetrieveListGet200ResponseListTtgCommonSetting struct {
		InterimAcctInterval *float64 `json:"interimAcctInterval,omitempty"`
		MobileCountryCode   *string  `json:"mobileCountryCode,omitempty"`
		MobileNetworkCode   *string  `json:"mobileNetworkCode,omitempty"`
		SessionIdleTimeout  *float64 `json:"sessionIdleTimeout,omitempty"`
		SessionTimeout      *float64 `json:"sessionTimeout,omitempty"`
	}

	AuthenticationProfileRetrieveListGet200ResponseList struct {
		AaaSuppportEnabled *bool                                                                 `json:"aaaSuppportEnabled,omitempty"`
		CreateDateTime     *int                                                                  `json:"createDateTime,omitempty"`
		CreatorID          *string                                                               `json:"creatorId,omitempty"`
		CreatorUsername    *string                                                               `json:"creatorUsername,omitempty"`
		Description        *string                                                               `json:"description,omitempty"`
		DomainID           *string                                                               `json:"domainId,omitempty"`
		GppSuppportEnabled *bool                                                                 `json:"gppSuppportEnabled,omitempty"`
		H20SuppportEnabled *bool                                                                 `json:"h20SuppportEnabled,omitempty"`
		ID                 *string                                                               `json:"id,omitempty"`
		ModifiedDateTime   *int                                                                  `json:"modifiedDateTime,omitempty"`
		ModifierID         *string                                                               `json:"modifierId,omitempty"`
		ModifierUsername   *string                                                               `json:"modifierUsername,omitempty"`
		MvnoID             *string                                                               `json:"mvnoId,omitempty"`
		Name               *string                                                               `json:"name,omitempty"`
		RealmMappings      AuthenticationProfileRetrieveListGet200ResponseListRealmMappingsSlice `json:"realmMappings,omitempty"`
		TtgCommonSetting   *AuthenticationProfileRetrieveListGet200ResponseListTtgCommonSetting  `json:"ttgCommonSetting,omitempty"`
	}

	AuthenticationProfileRetrieveListGet200Response struct {
		Extra      *json.RawMessage                                         `json:"extra,omitempty"`
		FirstIndex *int                                                     `json:"firstIndex,omitempty"`
		HasMore    *bool                                                    `json:"hasMore,omitempty"`
		List       AuthenticationProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                     `json:"totalCount,omitempty"`
	}
)

// AuthenticationProfileRetrieveListGet: Use this API command to retrieve a list of authentication profiles.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileRetrieveListGet(ctx context.Context) (*http.Response, *AuthenticationProfileRetrieveListGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/auth")
	request.authenticated = true

	out := &AuthenticationProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationProfileCreatePostRequestRealmMappingsSlice []*AuthenticationProfileCreatePostRequestRealmMappings

	AuthenticationProfileCreatePostRequestRealmMappings struct {
		AuthorizationMethod *string  `json:"authorizationMethod,omitempty"`
		DynamicVlanID       *float64 `json:"dynamicVlanId,omitempty"`
		HostedAaaEnabled    *bool    `json:"hostedAaaEnabled,omitempty"`
		ID                  *string  `json:"id,omitempty"`
		Name                *string  `json:"name,omitempty"`
		Realm               *string  `json:"realm,omitempty"`
		ServiceType         *string  `json:"serviceType,omitempty"`
	}

	AuthenticationProfileCreatePostRequestTtgCommonSetting struct {
		InterimAcctInterval *float64 `json:"interimAcctInterval,omitempty"`
		MobileCountryCode   *string  `json:"mobileCountryCode,omitempty"`
		MobileNetworkCode   *string  `json:"mobileNetworkCode,omitempty"`
		SessionIdleTimeout  *float64 `json:"sessionIdleTimeout,omitempty"`
		SessionTimeout      *float64 `json:"sessionTimeout,omitempty"`
	}

	AuthenticationProfileCreatePostRequest struct {
		AaaSuppportEnabled *bool                                                    `json:"aaaSuppportEnabled,omitempty"`
		Description        *string                                                  `json:"description,omitempty"`
		DomainID           *string                                                  `json:"domainId,omitempty"`
		GppSuppportEnabled *bool                                                    `json:"gppSuppportEnabled,omitempty"`
		H20SuppportEnabled *bool                                                    `json:"h20SuppportEnabled,omitempty"`
		MvnoID             *string                                                  `json:"mvnoId,omitempty"`
		Name               *string                                                  `json:"name,omitempty"`
		RealmMappings      AuthenticationProfileCreatePostRequestRealmMappingsSlice `json:"realmMappings,omitempty"`
		TtgCommonSetting   *AuthenticationProfileCreatePostRequestTtgCommonSetting  `json:"ttgCommonSetting,omitempty"`
	}

	AuthenticationProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AuthenticationProfileCreatePost: Use this API command to create a new authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AuthenticationProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileCreatePost(ctx context.Context, requestBody *AuthenticationProfileCreatePostRequest) (*http.Response, *AuthenticationProfileCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/auth")
	request.body = requestBody
	request.authenticated = true

	out := &AuthenticationProfileCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AuthenticationProfileRetrieveListAuthorizationProfileGet200ResponseListSlice []*AuthenticationProfileRetrieveListAuthorizationProfileGet200ResponseList

	AuthenticationProfileRetrieveListAuthorizationProfileGet200ResponseList struct {
		ID          *string `json:"id,omitempty"`
		ServiceID   *string `json:"serviceId,omitempty"`
		ServiceName *string `json:"serviceName,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthorizationProfileGet200Response struct {
		FirstIndex *int                                                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                        `json:"hasMore,omitempty"`
		List       AuthenticationProfileRetrieveListAuthorizationProfileGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                         `json:"totalCount,omitempty"`
	}
)

// AuthenticationProfileRetrieveListAuthorizationProfileGet: Use this API command to retrieve a list of authorization profiles.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - xtype (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileRetrieveListAuthorizationProfileGet200Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileRetrieveListAuthorizationProfileGet(ctx context.Context, xtype string) (*http.Response, *AuthenticationProfileRetrieveListAuthorizationProfileGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(xtype)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"type\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/auth/authorizationList")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"type": xtype,
	}

	out := &AuthenticationProfileRetrieveListAuthorizationProfileGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestAttributesSlice []*string

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraFiltersSlice []*AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraFilters

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraNotFiltersSlice []*AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraNotFilters

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFiltersSlice []*AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFilters

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFullTextSearch struct {
		Fields AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                        `json:"type,omitempty"`
		Value  *string                                                                                        `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                            `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                          `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                            `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                            `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                          `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                            `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                            `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                            `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                            `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                            `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                            `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                          `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                            `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                          `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                          `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                          `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                          `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                          `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                            `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                            `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                            `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                          `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                          `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                          `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                          `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                          `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                          `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                          `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                          `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                          `json:"localUser_userSource,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequest struct {
		Attributes      AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                   `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                     `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                                                  `json:"limit,omitempty"`
		Options         *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                      `json:"page,omitempty"`
		Query           *string                                                                                   `json:"query,omitempty"`
		SortInfo        *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                                                  `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200ResponseListSlice []*AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200ResponseList

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200ResponseList struct {
		ID          *string `json:"id,omitempty"`
		ServiceID   *string `json:"serviceId,omitempty"`
		ServiceName *string `json:"serviceName,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200Response struct {
		FirstIndex *int                                                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                              `json:"hasMore,omitempty"`
		List       AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                               `json:"totalCount,omitempty"`
	}
)

// AuthenticationProfileRetrieveListAuthenticationServiceInfoPost: Use this API command to retrieve a list of authentication service.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileRetrieveListAuthenticationServiceInfoPost(ctx context.Context, requestBody *AuthenticationProfileRetrieveListAuthenticationServiceInfoPostRequest) (*http.Response, *AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/auth/authServiceList/query")
	request.body = requestBody
	request.authenticated = true

	out := &AuthenticationProfileRetrieveListAuthenticationServiceInfoPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationProfileClonePostRequest struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}

	AuthenticationProfileClonePost200Response struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}
)

// AuthenticationProfileClonePost: Use this API command to clone an authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *AuthenticationProfileClonePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileClonePost200Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileClonePost(ctx context.Context, id string, requestBody *AuthenticationProfileClonePostRequest) (*http.Response, *AuthenticationProfileClonePost200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/auth/clone/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &AuthenticationProfileClonePost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestAttributesSlice []*string

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraFiltersSlice []*AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraFilters

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraNotFiltersSlice []*AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraNotFilters

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFiltersSlice []*AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFilters

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFullTextSearch struct {
		Fields AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                    `json:"type,omitempty"`
		Value  *string                                                                                                    `json:"value,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                        `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                      `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                        `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                        `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                      `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                        `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                        `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                        `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                        `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                        `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                        `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                      `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                        `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                      `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                      `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                      `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                      `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                      `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                        `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                        `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                        `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                      `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                      `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                      `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                      `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                      `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                      `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                      `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                      `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                      `json:"localUser_userSource,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequest struct {
		Attributes      AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                               `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                 `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                  `json:"limit,omitempty"`
		Options         *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                  `json:"page,omitempty"`
		Query           *string                                                                                               `json:"query,omitempty"`
		SortInfo        *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                  `json:"start,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListSlice []*AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseList

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListRealmMappingsSlice []*AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListRealmMappings

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListRealmMappings struct {
		AuthorizationMethod *string  `json:"authorizationMethod,omitempty"`
		DynamicVlanID       *float64 `json:"dynamicVlanId,omitempty"`
		HostedAaaEnabled    *bool    `json:"hostedAaaEnabled,omitempty"`
		ID                  *string  `json:"id,omitempty"`
		Name                *string  `json:"name,omitempty"`
		Realm               *string  `json:"realm,omitempty"`
		ServiceType         *string  `json:"serviceType,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListTtgCommonSetting struct {
		InterimAcctInterval *float64 `json:"interimAcctInterval,omitempty"`
		MobileCountryCode   *string  `json:"mobileCountryCode,omitempty"`
		MobileNetworkCode   *string  `json:"mobileNetworkCode,omitempty"`
		SessionIdleTimeout  *float64 `json:"sessionIdleTimeout,omitempty"`
		SessionTimeout      *float64 `json:"sessionTimeout,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseList struct {
		AaaSuppportEnabled *bool                                                                                                       `json:"aaaSuppportEnabled,omitempty"`
		CreateDateTime     *int                                                                                                        `json:"createDateTime,omitempty"`
		CreatorID          *string                                                                                                     `json:"creatorId,omitempty"`
		CreatorUsername    *string                                                                                                     `json:"creatorUsername,omitempty"`
		Description        *string                                                                                                     `json:"description,omitempty"`
		DomainID           *string                                                                                                     `json:"domainId,omitempty"`
		GppSuppportEnabled *bool                                                                                                       `json:"gppSuppportEnabled,omitempty"`
		H20SuppportEnabled *bool                                                                                                       `json:"h20SuppportEnabled,omitempty"`
		ID                 *string                                                                                                     `json:"id,omitempty"`
		ModifiedDateTime   *int                                                                                                        `json:"modifiedDateTime,omitempty"`
		ModifierID         *string                                                                                                     `json:"modifierId,omitempty"`
		ModifierUsername   *string                                                                                                     `json:"modifierUsername,omitempty"`
		MvnoID             *string                                                                                                     `json:"mvnoId,omitempty"`
		Name               *string                                                                                                     `json:"name,omitempty"`
		RealmMappings      AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListRealmMappingsSlice `json:"realmMappings,omitempty"`
		TtgCommonSetting   *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListTtgCommonSetting  `json:"ttgCommonSetting,omitempty"`
	}

	AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                               `json:"extra,omitempty"`
		FirstIndex *int                                                                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                          `json:"hasMore,omitempty"`
		List       AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                           `json:"totalCount,omitempty"`
	}
)

// AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost: Use this API command to retrieve a list of authentication profiles by query criteria.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost(ctx context.Context, requestBody *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPostRequest) (*http.Response, *AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/auth/query")
	request.body = requestBody
	request.authenticated = true

	out := &AuthenticationProfileRetrieveListAuthenticationProfilesByQueryCritariaPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AuthenticationProfileDeleteDelete: Use this API command to delete an authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Authentication Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/auth/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	AuthenticationProfileRetrieveGet200ResponseRealmMappingsSlice []*AuthenticationProfileRetrieveGet200ResponseRealmMappings

	AuthenticationProfileRetrieveGet200ResponseRealmMappings struct {
		AuthorizationMethod *string  `json:"authorizationMethod,omitempty"`
		DynamicVlanID       *float64 `json:"dynamicVlanId,omitempty"`
		HostedAaaEnabled    *bool    `json:"hostedAaaEnabled,omitempty"`
		ID                  *string  `json:"id,omitempty"`
		Name                *string  `json:"name,omitempty"`
		Realm               *string  `json:"realm,omitempty"`
		ServiceType         *string  `json:"serviceType,omitempty"`
	}

	AuthenticationProfileRetrieveGet200ResponseTtgCommonSetting struct {
		InterimAcctInterval *float64 `json:"interimAcctInterval,omitempty"`
		MobileCountryCode   *string  `json:"mobileCountryCode,omitempty"`
		MobileNetworkCode   *string  `json:"mobileNetworkCode,omitempty"`
		SessionIdleTimeout  *float64 `json:"sessionIdleTimeout,omitempty"`
		SessionTimeout      *float64 `json:"sessionTimeout,omitempty"`
	}

	AuthenticationProfileRetrieveGet200Response struct {
		AaaSuppportEnabled *bool                                                         `json:"aaaSuppportEnabled,omitempty"`
		CreateDateTime     *int                                                          `json:"createDateTime,omitempty"`
		CreatorID          *string                                                       `json:"creatorId,omitempty"`
		CreatorUsername    *string                                                       `json:"creatorUsername,omitempty"`
		Description        *string                                                       `json:"description,omitempty"`
		DomainID           *string                                                       `json:"domainId,omitempty"`
		GppSuppportEnabled *bool                                                         `json:"gppSuppportEnabled,omitempty"`
		H20SuppportEnabled *bool                                                         `json:"h20SuppportEnabled,omitempty"`
		ID                 *string                                                       `json:"id,omitempty"`
		ModifiedDateTime   *int                                                          `json:"modifiedDateTime,omitempty"`
		ModifierID         *string                                                       `json:"modifierId,omitempty"`
		ModifierUsername   *string                                                       `json:"modifierUsername,omitempty"`
		MvnoID             *string                                                       `json:"mvnoId,omitempty"`
		Name               *string                                                       `json:"name,omitempty"`
		RealmMappings      AuthenticationProfileRetrieveGet200ResponseRealmMappingsSlice `json:"realmMappings,omitempty"`
		TtgCommonSetting   *AuthenticationProfileRetrieveGet200ResponseTtgCommonSetting  `json:"ttgCommonSetting,omitempty"`
	}
)

// AuthenticationProfileRetrieveGet: Use this API command to retrieve an authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Authentication Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *AuthenticationProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/auth/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &AuthenticationProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationProfileModifyBasicPatchRequest struct {
		AaaSuppportEnabled *bool   `json:"aaaSuppportEnabled,omitempty"`
		Description        *string `json:"description,omitempty"`
		DomainID           *string `json:"domainId,omitempty"`
		GppSuppportEnabled *bool   `json:"gppSuppportEnabled,omitempty"`
		H20SuppportEnabled *bool   `json:"h20SuppportEnabled,omitempty"`
		ID                 *string `json:"id,omitempty"`
		MvnoID             *string `json:"mvnoId,omitempty"`
		Name               *string `json:"name,omitempty"`
	}
)

// AuthenticationProfileModifyBasicPatch: Use this API command to modify the basic information of an authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Authentication Profile ID
// - requestBody: *AuthenticationProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileModifyBasicPatch(ctx context.Context, id string, requestBody *AuthenticationProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/auth/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	AuthenticationProfileModifyRealmMappingsPatchRequestSlice []*AuthenticationProfileModifyRealmMappingsPatchRequest

	AuthenticationProfileModifyRealmMappingsPatchRequest struct {
		AuthorizationMethod *string  `json:"authorizationMethod,omitempty"`
		DynamicVlanID       *float64 `json:"dynamicVlanId,omitempty"`
		HostedAaaEnabled    *bool    `json:"hostedAaaEnabled,omitempty"`
		ID                  *string  `json:"id,omitempty"`
		Name                *string  `json:"name,omitempty"`
		Realm               *string  `json:"realm,omitempty"`
		ServiceType         *string  `json:"serviceType,omitempty"`
	}
)

// AuthenticationProfileModifyRealmMappingsPatch: Use this API command to modify realm based authentication service mappings of an authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Authentication Profile ID
// - requestBody: *AuthenticationProfileModifyRealmMappingsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileModifyRealmMappingsPatch(ctx context.Context, id string, requestBody AuthenticationProfileModifyRealmMappingsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/auth/{id}/realmMappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	AuthenticationProfileModify3gppCommonSettingsPatchRequest struct {
		InterimAcctInterval *float64 `json:"interimAcctInterval,omitempty"`
		MobileCountryCode   *string  `json:"mobileCountryCode,omitempty"`
		MobileNetworkCode   *string  `json:"mobileNetworkCode,omitempty"`
		SessionIdleTimeout  *float64 `json:"sessionIdleTimeout,omitempty"`
		SessionTimeout      *float64 `json:"sessionTimeout,omitempty"`
	}
)

// AuthenticationProfileModify3gppCommonSettingsPatch: Use this API command to modify 3GPP common settings of an authentication profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Authentication Profile ID
// - requestBody: *AuthenticationProfileModify3gppCommonSettingsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) AuthenticationProfileModify3gppCommonSettingsPatch(ctx context.Context, id string, requestBody *AuthenticationProfileModify3gppCommonSettingsPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/auth/{id}/ttgCommonSetting")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	BridgeDeleteDelete1RequestIDListSlice []*string

	BridgeDeleteDelete1Request struct {
		IDList BridgeDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// BridgeDeleteDelete1: Use this API command to delete multiple bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BridgeDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) BridgeDeleteDelete1(ctx context.Context, requestBody *BridgeDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/bridge")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	BridgeRetrieveListGet200ResponseListSlice []*BridgeRetrieveListGet200ResponseList

	BridgeRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	BridgeRetrieveListGet200Response struct {
		CreateDateTime   *int                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                   `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                          `json:"extra,omitempty"`
		FirstIndex       *int                                      `json:"firstIndex,omitempty"`
		HasMore          *bool                                     `json:"hasMore,omitempty"`
		List             BridgeRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                   `json:"modifierUsername,omitempty"`
		TotalCount       *int                                      `json:"totalCount,omitempty"`
	}
)

// BridgeRetrieveListGet: Use this API command to retrieve a list of Bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BridgeRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) BridgeRetrieveListGet(ctx context.Context) (*http.Response, *BridgeRetrieveListGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/bridge")
	request.authenticated = true

	out := &BridgeRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	BridgeCreatePostRequestDhcpRelayDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	BridgeCreatePostRequestDhcpRelay struct {
		DhcpOption82     *BridgeCreatePostRequestDhcpRelayDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                         `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                       `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                       `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                         `json:"relayBothEnabled,omitempty"`
	}

	BridgeCreatePostRequest struct {
		Description *string                           `json:"description,omitempty"`
		DhcpRelay   *BridgeCreatePostRequestDhcpRelay `json:"dhcpRelay,omitempty"`
		DomainID    *string                           `json:"domainId,omitempty"`
		ID          *string                           `json:"id,omitempty"`
		Name        *string                           `json:"name,omitempty"`
	}

	BridgeCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// BridgeCreatePost: Use this API command to create Bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BridgeCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BridgeCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) BridgeCreatePost(ctx context.Context, requestBody *BridgeCreatePostRequest) (*http.Response, *BridgeCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/bridge")
	request.body = requestBody
	request.authenticated = true

	out := &BridgeCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	BridgeQueryListPostRequestAttributesSlice []*string

	BridgeQueryListPostRequestFiltersSlice []*BridgeQueryListPostRequestFilters

	BridgeQueryListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	BridgeQueryListPostRequestFullTextSearchFieldsSlice []*string

	BridgeQueryListPostRequestFullTextSearch struct {
		Fields BridgeQueryListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                             `json:"type,omitempty"`
		Value  *string                                             `json:"value,omitempty"`
	}

	BridgeQueryListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	BridgeQueryListPostRequest struct {
		Attributes     BridgeQueryListPostRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        BridgeQueryListPostRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *BridgeQueryListPostRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                      `json:"limit,omitempty"`
		Options        interface{}                               `json:"options,omitempty"`
		Page           *int                                      `json:"page,omitempty"`
		SortInfo       *BridgeQueryListPostRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                      `json:"start,omitempty"`
	}

	BridgeQueryListPost200ResponseListSlice []*BridgeQueryListPost200ResponseList

	BridgeQueryListPost200ResponseListDhcpRelayDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	BridgeQueryListPost200ResponseListDhcpRelay struct {
		DhcpOption82     *BridgeQueryListPost200ResponseListDhcpRelayDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                                    `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                                  `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                                  `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                                    `json:"relayBothEnabled,omitempty"`
	}

	BridgeQueryListPost200ResponseList struct {
		CreateDateTime   *int                                         `json:"createDateTime,omitempty"`
		CreatorID        *string                                      `json:"creatorId,omitempty"`
		CreatorUsername  *string                                      `json:"creatorUsername,omitempty"`
		Description      *string                                      `json:"description,omitempty"`
		DhcpRelay        *BridgeQueryListPost200ResponseListDhcpRelay `json:"dhcpRelay,omitempty"`
		DomainID         *string                                      `json:"domainId,omitempty"`
		ID               *string                                      `json:"id,omitempty"`
		ModifiedDateTime *int                                         `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                      `json:"modifierId,omitempty"`
		ModifierUsername *string                                      `json:"modifierUsername,omitempty"`
		Name             *string                                      `json:"name,omitempty"`
	}

	BridgeQueryListPost200Response struct {
		Extra      *json.RawMessage                        `json:"extra,omitempty"`
		FirstIndex *int                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                   `json:"hasMore,omitempty"`
		List       BridgeQueryListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                    `json:"totalCount,omitempty"`
	}
)

// BridgeQueryListPost: Use this API command to query a list of Bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BridgeQueryListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BridgeQueryListPost200Response
// - error: Error seen or nil if none
func (p *Profiles) BridgeQueryListPost(ctx context.Context, requestBody *BridgeQueryListPostRequest) (*http.Response, *BridgeQueryListPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/bridge/query")
	request.body = requestBody
	request.authenticated = true

	out := &BridgeQueryListPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// BridgeDeleteDelete: Use this API command to delete Bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Bridge ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) BridgeDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/bridge/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	BridgeRetrieveGet200ResponseDhcpRelayDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	BridgeRetrieveGet200ResponseDhcpRelay struct {
		DhcpOption82     *BridgeRetrieveGet200ResponseDhcpRelayDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                              `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                            `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                            `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                              `json:"relayBothEnabled,omitempty"`
	}

	BridgeRetrieveGet200Response struct {
		CreateDateTime   *int                                   `json:"createDateTime,omitempty"`
		CreatorID        *string                                `json:"creatorId,omitempty"`
		CreatorUsername  *string                                `json:"creatorUsername,omitempty"`
		Description      *string                                `json:"description,omitempty"`
		DhcpRelay        *BridgeRetrieveGet200ResponseDhcpRelay `json:"dhcpRelay,omitempty"`
		DomainID         *string                                `json:"domainId,omitempty"`
		ID               *string                                `json:"id,omitempty"`
		ModifiedDateTime *int                                   `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                `json:"modifierId,omitempty"`
		ModifierUsername *string                                `json:"modifierUsername,omitempty"`
		Name             *string                                `json:"name,omitempty"`
	}
)

// BridgeRetrieveGet: Use this API command to retrieve Bridge profile by ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Bridge ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BridgeRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) BridgeRetrieveGet(ctx context.Context, id string) (*http.Response, *BridgeRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/bridge/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &BridgeRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	BridgeModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// BridgeModifyBasicPatch: Use this API command to modify the basic information of Bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Bridge ID
// - requestBody: *BridgeModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) BridgeModifyBasicPatch(ctx context.Context, id string, requestBody *BridgeModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/bridge/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	BridgeModifyDhcpRelayPatchRequestDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	BridgeModifyDhcpRelayPatchRequest struct {
		DhcpOption82     *BridgeModifyDhcpRelayPatchRequestDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                          `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                        `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                        `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                          `json:"relayBothEnabled,omitempty"`
	}
)

// BridgeModifyDhcpRelayPatch: Use this API command to modify DHCP Relay of Bridge profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Bridge ID
// - requestBody: *BridgeModifyDhcpRelayPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) BridgeModifyDhcpRelayPatch(ctx context.Context, id string, requestBody *BridgeModifyDhcpRelayPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/bridge/{id}/dhcpRelay")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	DnsServerManagementDeleteAListOfDnsServerProfileDeleteRequestIDListSlice []*string

	DnsServerManagementDeleteAListOfDnsServerProfileDeleteRequest struct {
		IDList DnsServerManagementDeleteAListOfDnsServerProfileDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// DnsServerManagementDeleteAListOfDnsServerProfileDelete: Use this API command to delete a list of DNS server profile .
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *DnsServerManagementDeleteAListOfDnsServerProfileDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementDeleteAListOfDnsServerProfileDelete(ctx context.Context, requestBody *DnsServerManagementDeleteAListOfDnsServerProfileDeleteRequest) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/dnsserver")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	DnsServerManagementCreatePostRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		MvnoID      *string `json:"mvnoId,omitempty"`
		Name        *string `json:"name,omitempty"`
		PrimaryIP   *string `json:"primaryIp,omitempty"`
		SecondaryIP *string `json:"secondaryIp,omitempty"`
		TertiaryIP  *string `json:"tertiaryIp,omitempty"`
	}

	DnsServerManagementCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// DnsServerManagementCreatePost: Use this API command to create DNS server profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *DnsServerManagementCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DnsServerManagementCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementCreatePost(ctx context.Context, requestBody *DnsServerManagementCreatePostRequest) (*http.Response, *DnsServerManagementCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/dnsserver")
	request.body = requestBody
	request.authenticated = true

	out := &DnsServerManagementCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	DnsServerManagementClonePostRequest struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}

	DnsServerManagementClonePost200Response struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}
)

// DnsServerManagementClonePost: Use this API command to clone an DNS server profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *DnsServerManagementClonePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DnsServerManagementClonePost200Response
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementClonePost(ctx context.Context, id string, requestBody *DnsServerManagementClonePostRequest) (*http.Response, *DnsServerManagementClonePost200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/dnsserver/clone/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &DnsServerManagementClonePost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestAttributesSlice []*string

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraFiltersSlice []*DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraFilters

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraNotFiltersSlice []*DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraNotFilters

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFiltersSlice []*DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFilters

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFullTextSearch struct {
		Fields DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                            `json:"type,omitempty"`
		Value  *string                                                                                            `json:"value,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                              `json:"localUser_userSource,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequest struct {
		Attributes      DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                          `json:"limit,omitempty"`
		Options         *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                          `json:"page,omitempty"`
		Query           *string                                                                                       `json:"query,omitempty"`
		SortInfo        *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                          `json:"start,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200ResponseListSlice []*DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200ResponseList

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200ResponseList struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		MvnoID           *string `json:"mvnoId,omitempty"`
		Name             *string `json:"name,omitempty"`
		PrimaryIP        *string `json:"primaryIp,omitempty"`
		SecondaryIP      *string `json:"secondaryIp,omitempty"`
		TertiaryIP       *string `json:"tertiaryIp,omitempty"`
	}

	DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                       `json:"extra,omitempty"`
		FirstIndex *int                                                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                  `json:"hasMore,omitempty"`
		List       DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                   `json:"totalCount,omitempty"`
	}
)

// DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost: Use this API command to retrieve a list of DNS server profile  by query criteria.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost(ctx context.Context, requestBody *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPostRequest) (*http.Response, *DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/dnsserver/query")
	request.body = requestBody
	request.authenticated = true

	out := &DnsServerManagementRetrieveListDnsServerProfileByQueryCritariaPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// DnsServerManagementDeleteDelete: Use this API command to delete DNS server profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/dnsserver/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	DnsServerManagementRetrieveGet200Response struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		MvnoID           *string `json:"mvnoId,omitempty"`
		Name             *string `json:"name,omitempty"`
		PrimaryIP        *string `json:"primaryIp,omitempty"`
		SecondaryIP      *string `json:"secondaryIp,omitempty"`
		TertiaryIP       *string `json:"tertiaryIp,omitempty"`
	}
)

// DnsServerManagementRetrieveGet: Use this API command to retrieve DNS server profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DnsServerManagementRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementRetrieveGet(ctx context.Context, id string) (*http.Response, *DnsServerManagementRetrieveGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/dnsserver/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &DnsServerManagementRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	DnsServerManagementModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		MvnoID      *string `json:"mvnoId,omitempty"`
		Name        *string `json:"name,omitempty"`
		PrimaryIP   *string `json:"primaryIp,omitempty"`
		SecondaryIP *string `json:"secondaryIp,omitempty"`
		TertiaryIP  *string `json:"tertiaryIp,omitempty"`
	}
)

// DnsServerManagementModifyBasicPatch: Use this API command to modify the basic information of DNS server profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *DnsServerManagementModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) DnsServerManagementModifyBasicPatch(ctx context.Context, id string, requestBody *DnsServerManagementModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/dnsserver/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileDeleteDelete1RequestIDListSlice []*string

	Hotspot20IdentityProviderProfileDeleteDelete1Request struct {
		IDList Hotspot20IdentityProviderProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileDeleteDelete1: Use this API command to delete multiple Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *Hotspot20IdentityProviderProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileDeleteDelete1(ctx context.Context, requestBody *Hotspot20IdentityProviderProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/identityproviders")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseList

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAccountingsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAccountings

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAccountings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAuthenticationsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAuthentications

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAuthentications struct {
		ID          *string  `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Realm       *string  `json:"realm,omitempty"`
		ServiceType *string  `json:"serviceType,omitempty"`
		VlanID      *float64 `json:"vlanId,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListHomeOisSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListHomeOis

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListHomeOis struct {
		Name *string `json:"name,omitempty"`
		Oi   *string `json:"oi,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSU struct {
		CommonLanguageIcon       *string                                                                                                   `json:"commonLanguageIcon,omitempty"`
		OsuNaiRealm              *string                                                                                                   `json:"osuNaiRealm,omitempty"`
		OsuServiceURL            *string                                                                                                   `json:"osuServiceUrl,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuAuthServicesSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuAuthServices

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuAuthServices struct {
		CredentialType *string  `json:"credentialType,omitempty"`
		Expiration     *float64 `json:"expiration,omitempty"`
		ID             *string  `json:"id,omitempty"`
		Name           *string  `json:"name,omitempty"`
		Realm          *string  `json:"realm,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuPortalInternalOSUPortal struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuPortal struct {
		ExternalURL       *string                                                                                                 `json:"externalUrl,omitempty"`
		InternalOSUPortal *Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuPortalInternalOSUPortal `json:"internalOSUPortal,omitempty"`
		Type              *string                                                                                                 `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSU struct {
		Certificate              *Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUCertificate                  `json:"certificate,omitempty"`
		CommonLanguageIcon       *string                                                                                                   `json:"commonLanguageIcon,omitempty"`
		OsuAuthServices          Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuAuthServicesSlice          `json:"osuAuthServices,omitempty"`
		OsuPortal                *Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUOsuPortal                    `json:"osuPortal,omitempty"`
		ProvisioningFormat       *string                                                                                                   `json:"provisioningFormat,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		ProvisioningUpdateType   *string                                                                                                   `json:"provisioningUpdateType,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsu struct {
		ExternalOSU *Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuExternalOSU `json:"externalOSU,omitempty"`
		InternalOSU *Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsuInternalOSU `json:"internalOSU,omitempty"`
		Type        *string                                                                       `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListPlmnsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListPlmns

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListPlmns struct {
		Mcc *string `json:"mcc,omitempty"`
		Mnc *string `json:"mnc,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealms

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethodsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethods

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethodsAuthSettingsSlice []*Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethodsAuthSettings

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethodsAuthSettings struct {
		Info       *string  `json:"info,omitempty"`
		Type       *string  `json:"type,omitempty"`
		VendorID   *float64 `json:"vendorId,omitempty"`
		VendorType *float64 `json:"vendorType,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethods struct {
		AuthSettings Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethodsAuthSettingsSlice `json:"authSettings,omitempty"`
		Type         *string                                                                                         `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealms struct {
		EapMethods Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsEapMethodsSlice `json:"eapMethods,omitempty"`
		Encoding   *string                                                                             `json:"encoding,omitempty"`
		Name       *string                                                                             `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200ResponseList struct {
		Accountings      Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAccountingsSlice     `json:"accountings,omitempty"`
		Authentications  Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListAuthenticationsSlice `json:"authentications,omitempty"`
		CreateDateTime   *int                                                                               `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                            `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                            `json:"creatorUsername,omitempty"`
		Description      *string                                                                            `json:"description,omitempty"`
		DomainID         *string                                                                            `json:"domainId,omitempty"`
		HomeOis          Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListHomeOisSlice         `json:"homeOis,omitempty"`
		ID               *string                                                                            `json:"id,omitempty"`
		ModifiedDateTime *int                                                                               `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                            `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                            `json:"modifierUsername,omitempty"`
		Name             *string                                                                            `json:"name,omitempty"`
		Osu              *Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListOsu                 `json:"osu,omitempty"`
		Plmns            Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListPlmnsSlice           `json:"plmns,omitempty"`
		Realms           Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListRealmsSlice          `json:"realms,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveListGet200Response struct {
		Extra      *json.RawMessage                                                    `json:"extra,omitempty"`
		FirstIndex *int                                                                `json:"firstIndex,omitempty"`
		HasMore    *bool                                                               `json:"hasMore,omitempty"`
		List       Hotspot20IdentityProviderProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                `json:"totalCount,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileRetrieveListGet: Use this API command to retrieve list of Hotspot 2.0 identity providers.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20IdentityProviderProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *Hotspot20IdentityProviderProfileRetrieveListGet200Response, error) {
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

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/hs20/identityproviders")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}

	out := &Hotspot20IdentityProviderProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20IdentityProviderProfileCreatePostRequestAccountingsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestAccountings

	Hotspot20IdentityProviderProfileCreatePostRequestAccountings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		RefBlank    *bool   `json:"refBlank,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestAuthenticationsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestAuthentications

	Hotspot20IdentityProviderProfileCreatePostRequestAuthentications struct {
		ID          *string  `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Realm       *string  `json:"realm,omitempty"`
		RefBlank    *bool    `json:"refBlank,omitempty"`
		ServiceType *string  `json:"serviceType,omitempty"`
		VlanID      *float64 `json:"vlanId,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestHomeOisSlice []*Hotspot20IdentityProviderProfileCreatePostRequestHomeOis

	Hotspot20IdentityProviderProfileCreatePostRequestHomeOis struct {
		Name *string `json:"name,omitempty"`
		Oi   *string `json:"oi,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSU struct {
		CommonLanguageIcon       *string                                                                                      `json:"commonLanguageIcon,omitempty"`
		OsuNaiRealm              *string                                                                                      `json:"osuNaiRealm,omitempty"`
		OsuServiceURL            *string                                                                                      `json:"osuServiceUrl,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuAuthServicesSlice []*Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuAuthServices

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuAuthServices struct {
		CredentialType *string  `json:"credentialType,omitempty"`
		Expiration     *float64 `json:"expiration,omitempty"`
		ID             *string  `json:"id,omitempty"`
		Name           *string  `json:"name,omitempty"`
		Realm          *string  `json:"realm,omitempty"`
		RefBlank       *bool    `json:"refBlank,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuPortalInternalOSUPortal struct {
		ID       *string `json:"id,omitempty"`
		Name     *string `json:"name,omitempty"`
		RefBlank *bool   `json:"refBlank,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuPortal struct {
		ExternalURL       *string                                                                                    `json:"externalUrl,omitempty"`
		InternalOSUPortal *Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuPortalInternalOSUPortal `json:"internalOSUPortal,omitempty"`
		Type              *string                                                                                    `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUWhitelistedDomainsSlice []interface{}

	Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSU struct {
		Certificate              *Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUCertificate                  `json:"certificate,omitempty"`
		CommonLanguageIcon       *string                                                                                      `json:"commonLanguageIcon,omitempty"`
		OsuAuthServices          Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuAuthServicesSlice          `json:"osuAuthServices,omitempty"`
		OsuPortal                *Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUOsuPortal                    `json:"osuPortal,omitempty"`
		ProvisioningFormat       *string                                                                                      `json:"provisioningFormat,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		ProvisioningUpdateType   *string                                                                                      `json:"provisioningUpdateType,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestOsu struct {
		ExternalOSU *Hotspot20IdentityProviderProfileCreatePostRequestOsuExternalOSU `json:"externalOSU,omitempty"`
		InternalOSU *Hotspot20IdentityProviderProfileCreatePostRequestOsuInternalOSU `json:"internalOSU,omitempty"`
		Type        *string                                                          `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestPlmnsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestPlmns

	Hotspot20IdentityProviderProfileCreatePostRequestPlmns struct {
		Mcc *string `json:"mcc,omitempty"`
		Mnc *string `json:"mnc,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestRealmsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestRealms

	Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethodsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethods

	Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethodsAuthSettingsSlice []*Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethodsAuthSettings

	Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethodsAuthSettings struct {
		Info       *string  `json:"info,omitempty"`
		Type       *string  `json:"type,omitempty"`
		VendorID   *float64 `json:"vendorId,omitempty"`
		VendorType *float64 `json:"vendorType,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethods struct {
		AuthSettings Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethodsAuthSettingsSlice `json:"authSettings,omitempty"`
		Type         *string                                                                            `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequestRealms struct {
		EapMethods Hotspot20IdentityProviderProfileCreatePostRequestRealmsEapMethodsSlice `json:"eapMethods,omitempty"`
		Encoding   *string                                                                `json:"encoding,omitempty"`
		Name       *string                                                                `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePostRequest struct {
		Accountings      Hotspot20IdentityProviderProfileCreatePostRequestAccountingsSlice     `json:"accountings,omitempty"`
		Authentications  Hotspot20IdentityProviderProfileCreatePostRequestAuthenticationsSlice `json:"authentications,omitempty"`
		CreateDateTime   *int                                                                  `json:"createDateTime,omitempty"`
		CreatorID        *string                                                               `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                               `json:"creatorUsername,omitempty"`
		Description      *string                                                               `json:"description,omitempty"`
		DomainID         *string                                                               `json:"domainId,omitempty"`
		HomeOis          Hotspot20IdentityProviderProfileCreatePostRequestHomeOisSlice         `json:"homeOis,omitempty"`
		ID               *string                                                               `json:"id,omitempty"`
		ModifiedDateTime *int                                                                  `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                               `json:"modifierId,omitempty"`
		ModifierUsername *string                                                               `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                               `json:"mvnoId,omitempty"`
		Name             *string                                                               `json:"name,omitempty"`
		Osu              *Hotspot20IdentityProviderProfileCreatePostRequestOsu                 `json:"osu,omitempty"`
		Plmns            Hotspot20IdentityProviderProfileCreatePostRequestPlmnsSlice           `json:"plmns,omitempty"`
		Realms           Hotspot20IdentityProviderProfileCreatePostRequestRealmsSlice          `json:"realms,omitempty"`
	}

	Hotspot20IdentityProviderProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileCreatePost: Use this API command to create a new Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *Hotspot20IdentityProviderProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20IdentityProviderProfileCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileCreatePost(ctx context.Context, requestBody *Hotspot20IdentityProviderProfileCreatePostRequest) (*http.Response, *Hotspot20IdentityProviderProfileCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/hs20/identityproviders")
	request.body = requestBody
	request.authenticated = true

	out := &Hotspot20IdentityProviderProfileCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestAttributesSlice []*string

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraFiltersSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraFilters

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraNotFiltersSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraNotFilters

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFiltersSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFilters

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFullTextSearchFieldsSlice []*string

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFullTextSearch struct {
		Fields Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                            `json:"type,omitempty"`
		Value  *string                                                                                            `json:"value,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                              `json:"localUser_userSource,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequest struct {
		Attributes      Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                          `json:"limit,omitempty"`
		Options         *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                          `json:"page,omitempty"`
		Query           *string                                                                                       `json:"query,omitempty"`
		SortInfo        *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                          `json:"start,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseList

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAccountingsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAccountings

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAccountings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAuthenticationsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAuthentications

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAuthentications struct {
		ID          *string  `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Realm       *string  `json:"realm,omitempty"`
		ServiceType *string  `json:"serviceType,omitempty"`
		VlanID      *float64 `json:"vlanId,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListHomeOisSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListHomeOis

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListHomeOis struct {
		Name *string `json:"name,omitempty"`
		Oi   *string `json:"oi,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSU struct {
		CommonLanguageIcon       *string                                                                                                                      `json:"commonLanguageIcon,omitempty"`
		OsuNaiRealm              *string                                                                                                                      `json:"osuNaiRealm,omitempty"`
		OsuServiceURL            *string                                                                                                                      `json:"osuServiceUrl,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuAuthServicesSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuAuthServices

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuAuthServices struct {
		CredentialType *string  `json:"credentialType,omitempty"`
		Expiration     *float64 `json:"expiration,omitempty"`
		ID             *string  `json:"id,omitempty"`
		Name           *string  `json:"name,omitempty"`
		Realm          *string  `json:"realm,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuPortalInternalOSUPortal struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuPortal struct {
		ExternalURL       *string                                                                                                                    `json:"externalUrl,omitempty"`
		InternalOSUPortal *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuPortalInternalOSUPortal `json:"internalOSUPortal,omitempty"`
		Type              *string                                                                                                                    `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSU struct {
		Certificate              *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUCertificate                  `json:"certificate,omitempty"`
		CommonLanguageIcon       *string                                                                                                                      `json:"commonLanguageIcon,omitempty"`
		OsuAuthServices          Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuAuthServicesSlice          `json:"osuAuthServices,omitempty"`
		OsuPortal                *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUOsuPortal                    `json:"osuPortal,omitempty"`
		ProvisioningFormat       *string                                                                                                                      `json:"provisioningFormat,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		ProvisioningUpdateType   *string                                                                                                                      `json:"provisioningUpdateType,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsu struct {
		ExternalOSU *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuExternalOSU `json:"externalOSU,omitempty"`
		InternalOSU *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsuInternalOSU `json:"internalOSU,omitempty"`
		Type        *string                                                                                          `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListPlmnsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListPlmns

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListPlmns struct {
		Mcc *string `json:"mcc,omitempty"`
		Mnc *string `json:"mnc,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealms

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethodsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethods

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethodsAuthSettingsSlice []*Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethodsAuthSettings

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethodsAuthSettings struct {
		Info       *string  `json:"info,omitempty"`
		Type       *string  `json:"type,omitempty"`
		VendorID   *float64 `json:"vendorId,omitempty"`
		VendorType *float64 `json:"vendorType,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethods struct {
		AuthSettings Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethodsAuthSettingsSlice `json:"authSettings,omitempty"`
		Type         *string                                                                                                            `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealms struct {
		EapMethods Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsEapMethodsSlice `json:"eapMethods,omitempty"`
		Encoding   *string                                                                                                `json:"encoding,omitempty"`
		Name       *string                                                                                                `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseList struct {
		Accountings      Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAccountingsSlice     `json:"accountings,omitempty"`
		Authentications  Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListAuthenticationsSlice `json:"authentications,omitempty"`
		CreateDateTime   *int                                                                                                  `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                               `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                               `json:"creatorUsername,omitempty"`
		Description      *string                                                                                               `json:"description,omitempty"`
		DomainID         *string                                                                                               `json:"domainId,omitempty"`
		HomeOis          Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListHomeOisSlice         `json:"homeOis,omitempty"`
		ID               *string                                                                                               `json:"id,omitempty"`
		ModifiedDateTime *int                                                                                                  `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                               `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                               `json:"modifierUsername,omitempty"`
		Name             *string                                                                                               `json:"name,omitempty"`
		Osu              *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListOsu                 `json:"osu,omitempty"`
		Plmns            Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListPlmnsSlice           `json:"plmns,omitempty"`
		Realms           Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListRealmsSlice          `json:"realms,omitempty"`
	}

	Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200Response struct {
		Extra      *json.RawMessage                                                                       `json:"extra,omitempty"`
		FirstIndex *int                                                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                  `json:"hasMore,omitempty"`
		List       Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                   `json:"totalCount,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost: Query hotspot 2.0 identity providers
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost(ctx context.Context, requestBody *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPostRequest) (*http.Response, *Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/hs20/identityproviders/query")
	request.body = requestBody
	request.authenticated = true

	out := &Hotspot20IdentityProviderProfileQueryHotspot20IdentityProviderPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// Hotspot20IdentityProviderProfileDeleteDelete: Use this API command to delete a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/identityproviders/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileRetrieveGet200ResponseAccountingsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseAccountings

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseAccountings struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		RefBlank    *bool   `json:"refBlank,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseAuthenticationsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseAuthentications

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseAuthentications struct {
		ID          *string  `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Realm       *string  `json:"realm,omitempty"`
		RefBlank    *bool    `json:"refBlank,omitempty"`
		ServiceType *string  `json:"serviceType,omitempty"`
		VlanID      *float64 `json:"vlanId,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseHomeOisSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseHomeOis

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseHomeOis struct {
		Name *string `json:"name,omitempty"`
		Oi   *string `json:"oi,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSU struct {
		CommonLanguageIcon       *string                                                                                           `json:"commonLanguageIcon,omitempty"`
		OsuNaiRealm              *string                                                                                           `json:"osuNaiRealm,omitempty"`
		OsuServiceURL            *string                                                                                           `json:"osuServiceUrl,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuAuthServicesSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuAuthServices

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuAuthServices struct {
		CredentialType *string  `json:"credentialType,omitempty"`
		Expiration     *float64 `json:"expiration,omitempty"`
		ID             *string  `json:"id,omitempty"`
		Name           *string  `json:"name,omitempty"`
		Realm          *string  `json:"realm,omitempty"`
		RefBlank       *bool    `json:"refBlank,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuPortalInternalOSUPortal struct {
		ID       *string `json:"id,omitempty"`
		Name     *string `json:"name,omitempty"`
		RefBlank *bool   `json:"refBlank,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuPortal struct {
		ExternalURL       *string                                                                                         `json:"externalUrl,omitempty"`
		InternalOSUPortal *Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuPortalInternalOSUPortal `json:"internalOSUPortal,omitempty"`
		Type              *string                                                                                         `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUWhitelistedDomainsSlice []interface{}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSU struct {
		Certificate              *Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUCertificate                  `json:"certificate,omitempty"`
		CommonLanguageIcon       *string                                                                                           `json:"commonLanguageIcon,omitempty"`
		OsuAuthServices          Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuAuthServicesSlice          `json:"osuAuthServices,omitempty"`
		OsuPortal                *Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUOsuPortal                    `json:"osuPortal,omitempty"`
		ProvisioningFormat       *string                                                                                           `json:"provisioningFormat,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		ProvisioningUpdateType   *string                                                                                           `json:"provisioningUpdateType,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsu struct {
		ExternalOSU *Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuExternalOSU `json:"externalOSU,omitempty"`
		InternalOSU *Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsuInternalOSU `json:"internalOSU,omitempty"`
		Type        *string                                                               `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponsePlmnsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponsePlmns

	Hotspot20IdentityProviderProfileRetrieveGet200ResponsePlmns struct {
		Mcc *string `json:"mcc,omitempty"`
		Mnc *string `json:"mnc,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealms

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethodsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethods

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethodsAuthSettingsSlice []*Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethodsAuthSettings

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethodsAuthSettings struct {
		Info       *string  `json:"info,omitempty"`
		Type       *string  `json:"type,omitempty"`
		VendorID   *float64 `json:"vendorId,omitempty"`
		VendorType *float64 `json:"vendorType,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethods struct {
		AuthSettings Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethodsAuthSettingsSlice `json:"authSettings,omitempty"`
		Type         *string                                                                                 `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealms struct {
		EapMethods Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsEapMethodsSlice `json:"eapMethods,omitempty"`
		Encoding   *string                                                                     `json:"encoding,omitempty"`
		Name       *string                                                                     `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileRetrieveGet200Response struct {
		Accountings      Hotspot20IdentityProviderProfileRetrieveGet200ResponseAccountingsSlice     `json:"accountings,omitempty"`
		Authentications  Hotspot20IdentityProviderProfileRetrieveGet200ResponseAuthenticationsSlice `json:"authentications,omitempty"`
		CreateDateTime   *int                                                                       `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                    `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                    `json:"creatorUsername,omitempty"`
		Description      *string                                                                    `json:"description,omitempty"`
		DomainID         *string                                                                    `json:"domainId,omitempty"`
		HomeOis          Hotspot20IdentityProviderProfileRetrieveGet200ResponseHomeOisSlice         `json:"homeOis,omitempty"`
		ID               *string                                                                    `json:"id,omitempty"`
		ModifiedDateTime *int                                                                       `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                    `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                    `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                    `json:"mvnoId,omitempty"`
		Name             *string                                                                    `json:"name,omitempty"`
		Osu              *Hotspot20IdentityProviderProfileRetrieveGet200ResponseOsu                 `json:"osu,omitempty"`
		Plmns            Hotspot20IdentityProviderProfileRetrieveGet200ResponsePlmnsSlice           `json:"plmns,omitempty"`
		Realms           Hotspot20IdentityProviderProfileRetrieveGet200ResponseRealmsSlice          `json:"realms,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileRetrieveGet: Use this API command to retrieve a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20IdentityProviderProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *Hotspot20IdentityProviderProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/hs20/identityproviders/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &Hotspot20IdentityProviderProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20IdentityProviderProfileModifyBasicPatchRequest struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Name             *string `json:"name,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyBasicPatch: Use this API command to modify the basic information of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyBasicPatch(ctx context.Context, id string, requestBody *Hotspot20IdentityProviderProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

// Hotspot20IdentityProviderProfileDisableAccountingsDelete: Use this API command to disable accountings of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileDisableAccountingsDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/identityproviders/{id}/accountings")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileModifyAccountingsPatchRequestSlice []*Hotspot20IdentityProviderProfileModifyAccountingsPatchRequest

	Hotspot20IdentityProviderProfileModifyAccountingsPatchRequest struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Realm       *string `json:"realm,omitempty"`
		RefBlank    *bool   `json:"refBlank,omitempty"`
		ServiceType *string `json:"serviceType,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyAccountingsPatch: Use this API command to modify accountings of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyAccountingsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyAccountingsPatch(ctx context.Context, id string, requestBody Hotspot20IdentityProviderProfileModifyAccountingsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}/accountings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileModifyAuthenticationsPatchRequestSlice []*Hotspot20IdentityProviderProfileModifyAuthenticationsPatchRequest

	Hotspot20IdentityProviderProfileModifyAuthenticationsPatchRequest struct {
		ID          *string  `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Realm       *string  `json:"realm,omitempty"`
		RefBlank    *bool    `json:"refBlank,omitempty"`
		ServiceType *string  `json:"serviceType,omitempty"`
		VlanID      *float64 `json:"vlanId,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyAuthenticationsPatch: Use this API command to modify authentications of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyAuthenticationsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyAuthenticationsPatch(ctx context.Context, id string, requestBody Hotspot20IdentityProviderProfileModifyAuthenticationsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}/authentications")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileModifyHomeOisPatchRequestSlice []*Hotspot20IdentityProviderProfileModifyHomeOisPatchRequest

	Hotspot20IdentityProviderProfileModifyHomeOisPatchRequest struct {
		Name *string `json:"name,omitempty"`
		Oi   *string `json:"oi,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyHomeOisPatch: Use this API command to modify home OIs of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyHomeOisPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyHomeOisPatch(ctx context.Context, id string, requestBody Hotspot20IdentityProviderProfileModifyHomeOisPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}/homeOis")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

// Hotspot20IdentityProviderProfileDisableOnlineSignupProvisioningDelete: Use this API command to disable online signup & provisioning of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileDisableOnlineSignupProvisioningDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/identityproviders/{id}/osu")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUWhitelistedDomainsSlice []*string

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSU struct {
		CommonLanguageIcon       *string                                                                                                            `json:"commonLanguageIcon,omitempty"`
		OsuNaiRealm              *string                                                                                                            `json:"osuNaiRealm,omitempty"`
		OsuServiceURL            *string                                                                                                            `json:"osuServiceUrl,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuAuthServicesSlice []*Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuAuthServices

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuAuthServices struct {
		CredentialType *string  `json:"credentialType,omitempty"`
		Expiration     *float64 `json:"expiration,omitempty"`
		ID             *string  `json:"id,omitempty"`
		Name           *string  `json:"name,omitempty"`
		Realm          *string  `json:"realm,omitempty"`
		RefBlank       *bool    `json:"refBlank,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuPortalInternalOSUPortal struct {
		ID       *string `json:"id,omitempty"`
		Name     *string `json:"name,omitempty"`
		RefBlank *bool   `json:"refBlank,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuPortal struct {
		ExternalURL       *string                                                                                                          `json:"externalUrl,omitempty"`
		InternalOSUPortal *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuPortalInternalOSUPortal `json:"internalOSUPortal,omitempty"`
		Type              *string                                                                                                          `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUProvisioningProtocalsSlice []*string

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUSubscriptionDescriptionsSlice []*Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUSubscriptionDescriptions

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUSubscriptionDescriptions struct {
		Description *string `json:"description,omitempty"`
		Icon        *string `json:"icon,omitempty"`
		Language    *string `json:"language,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUWhitelistedDomainsSlice []interface{}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSU struct {
		Certificate              *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUCertificate                  `json:"certificate,omitempty"`
		CommonLanguageIcon       *string                                                                                                            `json:"commonLanguageIcon,omitempty"`
		OsuAuthServices          Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuAuthServicesSlice          `json:"osuAuthServices,omitempty"`
		OsuPortal                *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUOsuPortal                    `json:"osuPortal,omitempty"`
		ProvisioningFormat       *string                                                                                                            `json:"provisioningFormat,omitempty"`
		ProvisioningProtocals    Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUProvisioningProtocalsSlice    `json:"provisioningProtocals,omitempty"`
		ProvisioningUpdateType   *string                                                                                                            `json:"provisioningUpdateType,omitempty"`
		SubscriptionDescriptions Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUSubscriptionDescriptionsSlice `json:"subscriptionDescriptions,omitempty"`
		WhitelistedDomains       Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSUWhitelistedDomainsSlice       `json:"whitelistedDomains,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequest struct {
		ExternalOSU *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestExternalOSU `json:"externalOSU,omitempty"`
		InternalOSU *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequestInternalOSU `json:"internalOSU,omitempty"`
		Type        *string                                                                                `json:"type,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatch: Use this API command to modify online signup & provisioning of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatch(ctx context.Context, id string, requestBody *Hotspot20IdentityProviderProfileModifyOnlineSignupProvisioningPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}/osu")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileModifyPlmnsPatchRequestSlice []*Hotspot20IdentityProviderProfileModifyPlmnsPatchRequest

	Hotspot20IdentityProviderProfileModifyPlmnsPatchRequest struct {
		Mcc *string `json:"mcc,omitempty"`
		Mnc *string `json:"mnc,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyPlmnsPatch: Use this API command to modify PLMNs of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyPlmnsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyPlmnsPatch(ctx context.Context, id string, requestBody Hotspot20IdentityProviderProfileModifyPlmnsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}/plmns")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20IdentityProviderProfileModifyRealmsPatchRequestSlice []*Hotspot20IdentityProviderProfileModifyRealmsPatchRequest

	Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethodsSlice []*Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethods

	Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethodsAuthSettingsSlice []*Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethodsAuthSettings

	Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethodsAuthSettings struct {
		Info       *string  `json:"info,omitempty"`
		Type       *string  `json:"type,omitempty"`
		VendorID   *float64 `json:"vendorId,omitempty"`
		VendorType *float64 `json:"vendorType,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethods struct {
		AuthSettings Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethodsAuthSettingsSlice `json:"authSettings,omitempty"`
		Type         *string                                                                             `json:"type,omitempty"`
	}

	Hotspot20IdentityProviderProfileModifyRealmsPatchRequest struct {
		EapMethods Hotspot20IdentityProviderProfileModifyRealmsPatchRequestEapMethodsSlice `json:"eapMethods,omitempty"`
		Encoding   *string                                                                 `json:"encoding,omitempty"`
		Name       *string                                                                 `json:"name,omitempty"`
	}
)

// Hotspot20IdentityProviderProfileModifyRealmsPatch: Use this API command to modify realms of a Hotspot 2.0 identity provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Identity Provider ID
// - requestBody: *Hotspot20IdentityProviderProfileModifyRealmsPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20IdentityProviderProfileModifyRealmsPatch(ctx context.Context, id string, requestBody Hotspot20IdentityProviderProfileModifyRealmsPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/identityproviders/{id}/realms")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20WiFiOperatorProfileDeleteDelete1RequestIDListSlice []*string

	Hotspot20WiFiOperatorProfileDeleteDelete1Request struct {
		IDList Hotspot20WiFiOperatorProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileDeleteDelete1: Use this API command to delete multiple Hotspot 2.0 Wi-Fi operators.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *Hotspot20WiFiOperatorProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileDeleteDelete1(ctx context.Context, requestBody *Hotspot20WiFiOperatorProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/operators")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListSlice []*Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseList

	Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListDomainNamesSlice []*string

	Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListFriendlyNamesSlice []*Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListFriendlyNames

	Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListFriendlyNames struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseList struct {
		Certificate      *Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListCertificate       `json:"certificate,omitempty"`
		CreateDateTime   *int                                                                         `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                      `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                      `json:"creatorUsername,omitempty"`
		Description      *string                                                                      `json:"description,omitempty"`
		DomainID         *string                                                                      `json:"domainId,omitempty"`
		DomainNames      Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListDomainNamesSlice   `json:"domainNames,omitempty"`
		FriendlyNames    Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListFriendlyNamesSlice `json:"friendlyNames,omitempty"`
		ID               *string                                                                      `json:"id,omitempty"`
		ModifiedDateTime *int                                                                         `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                      `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                      `json:"modifierUsername,omitempty"`
		Name             *string                                                                      `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileRetrieveListGet200Response struct {
		Extra      *json.RawMessage                                                `json:"extra,omitempty"`
		FirstIndex *int                                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                                           `json:"hasMore,omitempty"`
		List       Hotspot20WiFiOperatorProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                            `json:"totalCount,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileRetrieveListGet: Use this API command to retrieve list of Hotspot 2.0 Wi-Fi Operators.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WiFiOperatorProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *Hotspot20WiFiOperatorProfileRetrieveListGet200Response, error) {
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

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/hs20/operators")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}

	out := &Hotspot20WiFiOperatorProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20WiFiOperatorProfileCreatePostRequestCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileCreatePostRequestDomainNamesSlice []*string

	Hotspot20WiFiOperatorProfileCreatePostRequestFriendlyNamesSlice []*Hotspot20WiFiOperatorProfileCreatePostRequestFriendlyNames

	Hotspot20WiFiOperatorProfileCreatePostRequestFriendlyNames struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileCreatePostRequest struct {
		Certificate      *Hotspot20WiFiOperatorProfileCreatePostRequestCertificate       `json:"certificate,omitempty"`
		CreateDateTime   *int                                                            `json:"createDateTime,omitempty"`
		CreatorID        *string                                                         `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                         `json:"creatorUsername,omitempty"`
		Description      *string                                                         `json:"description,omitempty"`
		DomainID         *string                                                         `json:"domainId,omitempty"`
		DomainNames      Hotspot20WiFiOperatorProfileCreatePostRequestDomainNamesSlice   `json:"domainNames,omitempty"`
		FriendlyNames    Hotspot20WiFiOperatorProfileCreatePostRequestFriendlyNamesSlice `json:"friendlyNames,omitempty"`
		ID               *string                                                         `json:"id,omitempty"`
		ModifiedDateTime *int                                                            `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                         `json:"modifierId,omitempty"`
		ModifierUsername *string                                                         `json:"modifierUsername,omitempty"`
		Name             *string                                                         `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileCreatePost: Use this API command to create a new Hotspot 2.0 Wi-Fi operator,
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *Hotspot20WiFiOperatorProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WiFiOperatorProfileCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileCreatePost(ctx context.Context, requestBody *Hotspot20WiFiOperatorProfileCreatePostRequest) (*http.Response, *Hotspot20WiFiOperatorProfileCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/hs20/operators")
	request.body = requestBody
	request.authenticated = true

	out := &Hotspot20WiFiOperatorProfileCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestAttributesSlice []*string

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraFiltersSlice []*Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraFilters

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraNotFiltersSlice []*Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraNotFilters

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFiltersSlice []*Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFilters

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFullTextSearchFieldsSlice []*string

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFullTextSearch struct {
		Fields Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                     `json:"type,omitempty"`
		Value  *string                                                                                     `json:"value,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                         `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                       `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                         `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                         `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                       `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                         `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                         `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                         `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                         `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                         `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                         `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                       `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                         `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                       `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                       `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                       `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                       `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                       `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                       `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                       `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                       `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                       `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                       `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                       `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                       `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                       `json:"localUser_userSource,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequest struct {
		Attributes      Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                   `json:"limit,omitempty"`
		Options         *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                   `json:"page,omitempty"`
		Query           *string                                                                                `json:"query,omitempty"`
		SortInfo        *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                   `json:"start,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListSlice []*Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseList

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListDomainNamesSlice []*string

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListFriendlyNamesSlice []*Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListFriendlyNames

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListFriendlyNames struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseList struct {
		Certificate      *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListCertificate       `json:"certificate,omitempty"`
		CreateDateTime   *int                                                                                         `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                      `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                      `json:"creatorUsername,omitempty"`
		Description      *string                                                                                      `json:"description,omitempty"`
		DomainID         *string                                                                                      `json:"domainId,omitempty"`
		DomainNames      Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListDomainNamesSlice   `json:"domainNames,omitempty"`
		FriendlyNames    Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListFriendlyNamesSlice `json:"friendlyNames,omitempty"`
		ID               *string                                                                                      `json:"id,omitempty"`
		ModifiedDateTime *int                                                                                         `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                      `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                      `json:"modifierUsername,omitempty"`
		Name             *string                                                                                      `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200Response struct {
		Extra      *json.RawMessage                                                                `json:"extra,omitempty"`
		FirstIndex *int                                                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                           `json:"hasMore,omitempty"`
		List       Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                            `json:"totalCount,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost: Query hotspot 2.0 Wi-Fi operators
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost(ctx context.Context, requestBody *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPostRequest) (*http.Response, *Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/hs20/operators/query")
	request.body = requestBody
	request.authenticated = true

	out := &Hotspot20WiFiOperatorProfileQueryHotspot20WiFiOperatorsPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// Hotspot20WiFiOperatorProfileDeleteDelete: Use this API command to delete a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/operators/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20WiFiOperatorProfileRetrieveGet200ResponseCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileRetrieveGet200ResponseDomainNamesSlice []*string

	Hotspot20WiFiOperatorProfileRetrieveGet200ResponseFriendlyNamesSlice []*Hotspot20WiFiOperatorProfileRetrieveGet200ResponseFriendlyNames

	Hotspot20WiFiOperatorProfileRetrieveGet200ResponseFriendlyNames struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}

	Hotspot20WiFiOperatorProfileRetrieveGet200Response struct {
		Certificate      *Hotspot20WiFiOperatorProfileRetrieveGet200ResponseCertificate       `json:"certificate,omitempty"`
		CreateDateTime   *int                                                                 `json:"createDateTime,omitempty"`
		CreatorID        *string                                                              `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                              `json:"creatorUsername,omitempty"`
		Description      *string                                                              `json:"description,omitempty"`
		DomainID         *string                                                              `json:"domainId,omitempty"`
		DomainNames      Hotspot20WiFiOperatorProfileRetrieveGet200ResponseDomainNamesSlice   `json:"domainNames,omitempty"`
		FriendlyNames    Hotspot20WiFiOperatorProfileRetrieveGet200ResponseFriendlyNamesSlice `json:"friendlyNames,omitempty"`
		ID               *string                                                              `json:"id,omitempty"`
		ModifiedDateTime *int                                                                 `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                              `json:"modifierId,omitempty"`
		ModifierUsername *string                                                              `json:"modifierUsername,omitempty"`
		Name             *string                                                              `json:"name,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileRetrieveGet: Use this API command to retrieve a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WiFiOperatorProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *Hotspot20WiFiOperatorProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/hs20/operators/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &Hotspot20WiFiOperatorProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20WiFiOperatorProfileModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileModifyBasicPatch: Use this API command to modify the basic information of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
// - requestBody: *Hotspot20WiFiOperatorProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileModifyBasicPatch(ctx context.Context, id string, requestBody *Hotspot20WiFiOperatorProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/operators/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

// Hotspot20WiFiOperatorProfileDisableCertificateDelete: Use this API command to disable certificate of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileDisableCertificateDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/hs20/operators/{id}/certificate")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20WiFiOperatorProfileModifyCertificatePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileModifyCertificatePatch: Use this API command to enable or modify certificate of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
// - requestBody: *Hotspot20WiFiOperatorProfileModifyCertificatePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileModifyCertificatePatch(ctx context.Context, id string, requestBody *Hotspot20WiFiOperatorProfileModifyCertificatePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/operators/{id}/certificate")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20WiFiOperatorProfileModifyDomainNamesPatchRequest []*string
)

// Hotspot20WiFiOperatorProfileModifyDomainNamesPatch: Use this API command to modify domain names of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
// - requestBody: *Hotspot20WiFiOperatorProfileModifyDomainNamesPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileModifyDomainNamesPatch(ctx context.Context, id string, requestBody Hotspot20WiFiOperatorProfileModifyDomainNamesPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/operators/{id}/domainNames")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatchRequestSlice []*Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatchRequest

	Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatchRequest struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}
)

// Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatch: Use this API command to modify friendly names of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): HotSpot 2.0 Wi-Fi Operator Profile ID
// - requestBody: *Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatch(ctx context.Context, id string, requestBody Hotspot20WiFiOperatorProfileModifyFriendlyNamesPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/hs20/operators/{id}/friendlyNames")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	L2ogreDeleteDelete1RequestIDListSlice []*string

	L2ogreDeleteDelete1Request struct {
		IDList L2ogreDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// L2ogreDeleteDelete1: Use this API command to delete multiple L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *L2ogreDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L2ogreDeleteDelete1(ctx context.Context, requestBody *L2ogreDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/l2ogre")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	L2ogreRetrieveListGet200ResponseListSlice []*L2ogreRetrieveListGet200ResponseList

	L2ogreRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	L2ogreRetrieveListGet200Response struct {
		CreateDateTime   *int                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                   `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                          `json:"extra,omitempty"`
		FirstIndex       *int                                      `json:"firstIndex,omitempty"`
		HasMore          *bool                                     `json:"hasMore,omitempty"`
		List             L2ogreRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                   `json:"modifierUsername,omitempty"`
		TotalCount       *int                                      `json:"totalCount,omitempty"`
	}
)

// L2ogreRetrieveListGet: Use this API command to retrieve a list of L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2ogreRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) L2ogreRetrieveListGet(ctx context.Context) (*http.Response, *L2ogreRetrieveListGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/l2ogre")
	request.authenticated = true

	out := &L2ogreRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	L2ogreCreatePostRequestCoreNetworkGateway struct {
		KeepAlivePeriod  *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry   *float64 `json:"keepAliveRetry,omitempty"`
		PrimaryGateway   *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway *string  `json:"secondaryGateway,omitempty"`
		TunnelMTU        *string  `json:"tunnelMTU,omitempty"`
		TunnelMTUSize    *float64 `json:"tunnelMTUSize,omitempty"`
	}

	L2ogreCreatePostRequestDhcpRelayDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	L2ogreCreatePostRequestDhcpRelay struct {
		DhcpOption82     *L2ogreCreatePostRequestDhcpRelayDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                         `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                       `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                       `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                         `json:"relayBothEnabled,omitempty"`
	}

	L2ogreCreatePostRequest struct {
		CoreNetworkGateway *L2ogreCreatePostRequestCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`
		Description        *string                                    `json:"description,omitempty"`
		DhcpRelay          *L2ogreCreatePostRequestDhcpRelay          `json:"dhcpRelay,omitempty"`
		DomainID           *string                                    `json:"domainId,omitempty"`
		ID                 *string                                    `json:"id,omitempty"`
		Name               *string                                    `json:"name,omitempty"`
	}

	L2ogreCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// L2ogreCreatePost: Use this API command to create L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *L2ogreCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2ogreCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) L2ogreCreatePost(ctx context.Context, requestBody *L2ogreCreatePostRequest) (*http.Response, *L2ogreCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/l2ogre")
	request.body = requestBody
	request.authenticated = true

	out := &L2ogreCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	L2ogreQueryListPostRequestAttributesSlice []*string

	L2ogreQueryListPostRequestFiltersSlice []*L2ogreQueryListPostRequestFilters

	L2ogreQueryListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	L2ogreQueryListPostRequestFullTextSearchFieldsSlice []*string

	L2ogreQueryListPostRequestFullTextSearch struct {
		Fields L2ogreQueryListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                             `json:"type,omitempty"`
		Value  *string                                             `json:"value,omitempty"`
	}

	L2ogreQueryListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	L2ogreQueryListPostRequest struct {
		Attributes     L2ogreQueryListPostRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        L2ogreQueryListPostRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *L2ogreQueryListPostRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                      `json:"limit,omitempty"`
		Options        interface{}                               `json:"options,omitempty"`
		Page           *int                                      `json:"page,omitempty"`
		SortInfo       *L2ogreQueryListPostRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                      `json:"start,omitempty"`
	}

	L2ogreQueryListPost200ResponseListSlice []*L2ogreQueryListPost200ResponseList

	L2ogreQueryListPost200ResponseListCoreNetworkGateway struct {
		KeepAlivePeriod  *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry   *float64 `json:"keepAliveRetry,omitempty"`
		PrimaryGateway   *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway *string  `json:"secondaryGateway,omitempty"`
		TunnelMTU        *string  `json:"tunnelMTU,omitempty"`
		TunnelMTUSize    *float64 `json:"tunnelMTUSize,omitempty"`
	}

	L2ogreQueryListPost200ResponseListDhcpRelayDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	L2ogreQueryListPost200ResponseListDhcpRelay struct {
		DhcpOption82     *L2ogreQueryListPost200ResponseListDhcpRelayDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                                    `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                                  `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                                  `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                                    `json:"relayBothEnabled,omitempty"`
	}

	L2ogreQueryListPost200ResponseList struct {
		CoreNetworkGateway *L2ogreQueryListPost200ResponseListCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`
		CreateDateTime     *int                                                  `json:"createDateTime,omitempty"`
		CreatorID          *string                                               `json:"creatorId,omitempty"`
		CreatorUsername    *string                                               `json:"creatorUsername,omitempty"`
		Description        *string                                               `json:"description,omitempty"`
		DhcpRelay          *L2ogreQueryListPost200ResponseListDhcpRelay          `json:"dhcpRelay,omitempty"`
		DomainID           *string                                               `json:"domainId,omitempty"`
		ID                 *string                                               `json:"id,omitempty"`
		ModifiedDateTime   *int                                                  `json:"modifiedDateTime,omitempty"`
		ModifierID         *string                                               `json:"modifierId,omitempty"`
		ModifierUsername   *string                                               `json:"modifierUsername,omitempty"`
		Name               *string                                               `json:"name,omitempty"`
	}

	L2ogreQueryListPost200Response struct {
		Extra      *json.RawMessage                        `json:"extra,omitempty"`
		FirstIndex *int                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                   `json:"hasMore,omitempty"`
		List       L2ogreQueryListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                    `json:"totalCount,omitempty"`
	}
)

// L2ogreQueryListPost: Use this API command to query a list of L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *L2ogreQueryListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2ogreQueryListPost200Response
// - error: Error seen or nil if none
func (p *Profiles) L2ogreQueryListPost(ctx context.Context, requestBody *L2ogreQueryListPostRequest) (*http.Response, *L2ogreQueryListPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/l2ogre/query")
	request.body = requestBody
	request.authenticated = true

	out := &L2ogreQueryListPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// L2ogreDeleteDelete: Use this API command to delete L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L2ogreDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/l2ogre/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	L2ogreRetrieveGet200ResponseCoreNetworkGateway struct {
		KeepAlivePeriod  *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry   *float64 `json:"keepAliveRetry,omitempty"`
		PrimaryGateway   *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway *string  `json:"secondaryGateway,omitempty"`
		TunnelMTU        *string  `json:"tunnelMTU,omitempty"`
		TunnelMTUSize    *float64 `json:"tunnelMTUSize,omitempty"`
	}

	L2ogreRetrieveGet200ResponseDhcpRelayDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	L2ogreRetrieveGet200ResponseDhcpRelay struct {
		DhcpOption82     *L2ogreRetrieveGet200ResponseDhcpRelayDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                              `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                            `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                            `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                              `json:"relayBothEnabled,omitempty"`
	}

	L2ogreRetrieveGet200Response struct {
		CoreNetworkGateway *L2ogreRetrieveGet200ResponseCoreNetworkGateway `json:"coreNetworkGateway,omitempty"`
		CreateDateTime     *int                                            `json:"createDateTime,omitempty"`
		CreatorID          *string                                         `json:"creatorId,omitempty"`
		CreatorUsername    *string                                         `json:"creatorUsername,omitempty"`
		Description        *string                                         `json:"description,omitempty"`
		DhcpRelay          *L2ogreRetrieveGet200ResponseDhcpRelay          `json:"dhcpRelay,omitempty"`
		DomainID           *string                                         `json:"domainId,omitempty"`
		ID                 *string                                         `json:"id,omitempty"`
		ModifiedDateTime   *int                                            `json:"modifiedDateTime,omitempty"`
		ModifierID         *string                                         `json:"modifierId,omitempty"`
		ModifierUsername   *string                                         `json:"modifierUsername,omitempty"`
		Name               *string                                         `json:"name,omitempty"`
	}
)

// L2ogreRetrieveGet: Use this API command to retrieve L2oGRE profile by ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2ogreRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) L2ogreRetrieveGet(ctx context.Context, id string) (*http.Response, *L2ogreRetrieveGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/l2ogre/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &L2ogreRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	L2ogreModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// L2ogreModifyBasicPatch: Use this API command to modify the basic information of L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *L2ogreModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L2ogreModifyBasicPatch(ctx context.Context, id string, requestBody *L2ogreModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/l2ogre/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	L2ogreModifyCoreNetworkGatewayPatchRequest struct {
		KeepAlivePeriod  *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry   *float64 `json:"keepAliveRetry,omitempty"`
		PrimaryGateway   *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway *string  `json:"secondaryGateway,omitempty"`
		TunnelMTU        *string  `json:"tunnelMTU,omitempty"`
		TunnelMTUSize    *float64 `json:"tunnelMTUSize,omitempty"`
	}
)

// L2ogreModifyCoreNetworkGatewayPatch: Use this API command to modify Core Network Gateway of a L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *L2ogreModifyCoreNetworkGatewayPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L2ogreModifyCoreNetworkGatewayPatch(ctx context.Context, id string, requestBody *L2ogreModifyCoreNetworkGatewayPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/l2ogre/{id}/coreNetworkGateway")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	L2ogreModifyDhcpRelayPatchRequestDhcpOption82 struct {
		DhcpOption82Enabled *bool   `json:"dhcpOption82Enabled,omitempty"`
		Subopt1Enabled      *bool   `json:"subopt1Enabled,omitempty"`
		Subopt1Format       *string `json:"subopt1Format,omitempty"`
		Subopt2Enabled      *bool   `json:"subopt2Enabled,omitempty"`
		Subopt2Format       *string `json:"subopt2Format,omitempty"`
		Subopt150Enabled    *bool   `json:"subopt150Enabled,omitempty"`
		Subopt151AreaName   *string `json:"subopt151AreaName,omitempty"`
		Subopt151Enabled    *bool   `json:"subopt151Enabled,omitempty"`
		Subopt151Format     *string `json:"subopt151Format,omitempty"`
	}

	L2ogreModifyDhcpRelayPatchRequest struct {
		DhcpOption82     *L2ogreModifyDhcpRelayPatchRequestDhcpOption82 `json:"dhcpOption82,omitempty"`
		DhcpRelayEnabled *bool                                          `json:"dhcpRelayEnabled,omitempty"`
		DhcpServer1      *string                                        `json:"dhcpServer1,omitempty"`
		DhcpServer2      *string                                        `json:"dhcpServer2,omitempty"`
		RelayBothEnabled *bool                                          `json:"relayBothEnabled,omitempty"`
	}
)

// L2ogreModifyDhcpRelayPatch: Use this API command to modify DHCP Relay of L2oGRE profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *L2ogreModifyDhcpRelayPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L2ogreModifyDhcpRelayPatch(ctx context.Context, id string, requestBody *L2ogreModifyDhcpRelayPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/l2ogre/{id}/dhcpRelay")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	LbsProfileDeleteMultipleLbsProfileDeleteRequestIDListSlice []*string

	LbsProfileDeleteMultipleLbsProfileDeleteRequest struct {
		IDList LbsProfileDeleteMultipleLbsProfileDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// LbsProfileDeleteMultipleLbsProfileDelete: Delete multiple LBS profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *LbsProfileDeleteMultipleLbsProfileDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) LbsProfileDeleteMultipleLbsProfileDelete(ctx context.Context, requestBody *LbsProfileDeleteMultipleLbsProfileDeleteRequest) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/lbs")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 200, nil)
}

type (
	LbsProfileCreateLbsProfilePostRequest struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Name             *string `json:"name,omitempty"`
		Password         *string `json:"password,omitempty"`
		Port             *int    `json:"port,omitempty"`
		URL              *string `json:"url,omitempty"`
		Venue            *string `json:"venue,omitempty"`
	}

	LbsProfileCreateLbsProfilePost200Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// LbsProfileCreateLbsProfilePost: Create LBS profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *LbsProfileCreateLbsProfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *LbsProfileCreateLbsProfilePost200Response
// - error: Error seen or nil if none
func (p *Profiles) LbsProfileCreateLbsProfilePost(ctx context.Context, requestBody *LbsProfileCreateLbsProfilePostRequest) (*http.Response, *LbsProfileCreateLbsProfilePost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/lbs")
	request.body = requestBody
	request.authenticated = true

	out := &LbsProfileCreateLbsProfilePost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	LbsProfileQueryLbsProfilesPostRequestAttributesSlice []*string

	LbsProfileQueryLbsProfilesPostRequestFiltersSlice []*LbsProfileQueryLbsProfilesPostRequestFilters

	LbsProfileQueryLbsProfilesPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	LbsProfileQueryLbsProfilesPostRequestFullTextSearchFieldsSlice []*string

	LbsProfileQueryLbsProfilesPostRequestFullTextSearch struct {
		Fields LbsProfileQueryLbsProfilesPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                        `json:"type,omitempty"`
		Value  *string                                                        `json:"value,omitempty"`
	}

	LbsProfileQueryLbsProfilesPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	LbsProfileQueryLbsProfilesPostRequest struct {
		Attributes     LbsProfileQueryLbsProfilesPostRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        LbsProfileQueryLbsProfilesPostRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *LbsProfileQueryLbsProfilesPostRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                                 `json:"limit,omitempty"`
		Options        interface{}                                          `json:"options,omitempty"`
		Page           *int                                                 `json:"page,omitempty"`
		SortInfo       *LbsProfileQueryLbsProfilesPostRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                                 `json:"start,omitempty"`
	}

	LbsProfileQueryLbsProfilesPost200ResponseListSlice []*LbsProfileQueryLbsProfilesPost200ResponseList

	LbsProfileQueryLbsProfilesPost200ResponseList struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Name             *string `json:"name,omitempty"`
		Password         *string `json:"password,omitempty"`
		Port             *int    `json:"port,omitempty"`
		URL              *string `json:"url,omitempty"`
		Venue            *string `json:"venue,omitempty"`
	}

	LbsProfileQueryLbsProfilesPost200Response struct {
		Extra      *json.RawMessage                                   `json:"extra,omitempty"`
		FirstIndex *int                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                              `json:"hasMore,omitempty"`
		List       LbsProfileQueryLbsProfilesPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                               `json:"totalCount,omitempty"`
	}
)

// LbsProfileQueryLbsProfilesPost: Query LBS profiles
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *LbsProfileQueryLbsProfilesPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *LbsProfileQueryLbsProfilesPost200Response
// - error: Error seen or nil if none
func (p *Profiles) LbsProfileQueryLbsProfilesPost(ctx context.Context, requestBody *LbsProfileQueryLbsProfilesPostRequest) (*http.Response, *LbsProfileQueryLbsProfilesPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/lbs/query")
	request.body = requestBody
	request.authenticated = true

	out := &LbsProfileQueryLbsProfilesPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// LbsProfileDeleteLbsProfileDelete: Delete LBS profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) LbsProfileDeleteLbsProfileDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/lbs/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 200, nil)
}

type (
	LbsProfileCreateLbsProfileGet200Response struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Name             *string `json:"name,omitempty"`
		Password         *string `json:"password,omitempty"`
		Port             *int    `json:"port,omitempty"`
		URL              *string `json:"url,omitempty"`
		Venue            *string `json:"venue,omitempty"`
	}
)

// LbsProfileCreateLbsProfileGet: Retrieve LBS profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *LbsProfileCreateLbsProfileGet200Response
// - error: Error seen or nil if none
func (p *Profiles) LbsProfileCreateLbsProfileGet(ctx context.Context, id string) (*http.Response, *LbsProfileCreateLbsProfileGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/lbs/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &LbsProfileCreateLbsProfileGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	LbsProfileUpdateLbsProfilePatchRequest struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		Name             *string `json:"name,omitempty"`
		Password         *string `json:"password,omitempty"`
		Port             *int    `json:"port,omitempty"`
		URL              *string `json:"url,omitempty"`
		Venue            *string `json:"venue,omitempty"`
	}
)

// LbsProfileUpdateLbsProfilePatch: Update LBS profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *LbsProfileUpdateLbsProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) LbsProfileUpdateLbsProfilePatch(ctx context.Context, id string, requestBody *LbsProfileUpdateLbsProfilePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/lbs/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 200, nil)
}

type (
	IpsecProfileDeleteDelete1RequestIDListSlice []*string

	IpsecProfileDeleteDelete1Request struct {
		IDList IpsecProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// IpsecProfileDeleteDelete1: Delete multiple ipsec
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *IpsecProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileDeleteDelete1(ctx context.Context, requestBody *IpsecProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/tunnel/ipsec")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	IpsecProfileRetrieveListGet200ResponseListSlice []*IpsecProfileRetrieveListGet200ResponseList

	IpsecProfileRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	IpsecProfileRetrieveListGet200Response struct {
		CreateDateTime   *int                                            `json:"createDateTime,omitempty"`
		CreatorID        *string                                         `json:"creatorId,omitempty"`
		CreatorUsername  *string                                         `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                                `json:"extra,omitempty"`
		FirstIndex       *int                                            `json:"firstIndex,omitempty"`
		HasMore          *bool                                           `json:"hasMore,omitempty"`
		List             IpsecProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                            `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                         `json:"modifierId,omitempty"`
		ModifierUsername *string                                         `json:"modifierUsername,omitempty"`
		TotalCount       *int                                            `json:"totalCount,omitempty"`
	}
)

// IpsecProfileRetrieveListGet: Retrieve a list of IPSEC
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IpsecProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *IpsecProfileRetrieveListGet200Response, error) {
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

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/tunnel/ipsec")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}

	out := &IpsecProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	IpsecProfileCreateIpsecPostRequestAdvancedOption struct {
		DhcpOpt43Subcode             *float64 `json:"dhcpOpt43Subcode,omitempty"`
		DpdDelay                     *float64 `json:"dpdDelay,omitempty"`
		EnforceNatt                  *string  `json:"enforceNatt,omitempty"`
		FailoverMode                 *string  `json:"failoverMode,omitempty"`
		FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`
		FailoverRetryInterval        *float64 `json:"failoverRetryInterval,omitempty"`
		FailoverRetryPeriod          *float64 `json:"failoverRetryPeriod,omitempty"`
		IpcompEnable                 *string  `json:"ipcompEnable,omitempty"`
		KeepAliveIntval              *float64 `json:"keepAliveIntval,omitempty"`
		ReplayWindow                 *float64 `json:"replayWindow,omitempty"`
		RetryLimit                   *float64 `json:"retryLimit,omitempty"`
	}

	IpsecProfileCreateIpsecPostRequestCmProtocolOption struct {
		CmpDhcpOpt43Subcode          *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`
		CmpDhcpOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`
		CmpRecipient                 *string  `json:"cmpRecipient,omitempty"`
		CmpServerAddr                *string  `json:"cmpServerAddr,omitempty"`
		CmpServerPath                *string  `json:"cmpServerPath,omitempty"`
	}

	IpsecProfileCreateIpsecPostRequestEspSecurityAssociationEspProposalsSlice []*IpsecProfileCreateIpsecPostRequestEspSecurityAssociationEspProposals

	IpsecProfileCreateIpsecPostRequestEspSecurityAssociationEspProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
	}

	IpsecProfileCreateIpsecPostRequestEspSecurityAssociation struct {
		EspProposalType *string                                                                   `json:"espProposalType,omitempty"`
		EspProposals    IpsecProfileCreateIpsecPostRequestEspSecurityAssociationEspProposalsSlice `json:"espProposals,omitempty"`
	}

	IpsecProfileCreateIpsecPostRequestIkeSecurityAssociationIkeProposalsSlice []*IpsecProfileCreateIpsecPostRequestIkeSecurityAssociationIkeProposals

	IpsecProfileCreateIpsecPostRequestIkeSecurityAssociationIkeProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
		PrfAlg  *string `json:"prfAlg,omitempty"`
	}

	IpsecProfileCreateIpsecPostRequestIkeSecurityAssociation struct {
		IkeProposalType *string                                                                   `json:"ikeProposalType,omitempty"`
		IkeProposals    IpsecProfileCreateIpsecPostRequestIkeSecurityAssociationIkeProposalsSlice `json:"ikeProposals,omitempty"`
	}

	IpsecProfileCreateIpsecPostRequest struct {
		AdvancedOption         *IpsecProfileCreateIpsecPostRequestAdvancedOption         `json:"advancedOption,omitempty"`
		AuthType               *string                                                   `json:"authType,omitempty"`
		CmProtocolOption       *IpsecProfileCreateIpsecPostRequestCmProtocolOption       `json:"cmProtocolOption,omitempty"`
		Description            *string                                                   `json:"description,omitempty"`
		DomainID               *string                                                   `json:"domainId,omitempty"`
		EspRekeyTime           *float64                                                  `json:"espRekeyTime,omitempty"`
		EspRekeyTimeUnit       *string                                                   `json:"espRekeyTimeUnit,omitempty"`
		EspSecurityAssociation *IpsecProfileCreateIpsecPostRequestEspSecurityAssociation `json:"espSecurityAssociation,omitempty"`
		ID                     *string                                                   `json:"id,omitempty"`
		IkeRekeyTime           *float64                                                  `json:"ikeRekeyTime,omitempty"`
		IkeRekeyTimeUnit       *string                                                   `json:"ikeRekeyTimeUnit,omitempty"`
		IkeSecurityAssociation *IpsecProfileCreateIpsecPostRequestIkeSecurityAssociation `json:"ikeSecurityAssociation,omitempty"`
		IPMode                 *string                                                   `json:"ipMode,omitempty"`
		Name                   *string                                                   `json:"name,omitempty"`
		PreSharedKey           *string                                                   `json:"preSharedKey,omitempty"`
		ServerAddr             *string                                                   `json:"serverAddr,omitempty"`
	}

	IpsecProfileCreateIpsecPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// IpsecProfileCreateIpsecPost: Create a new ipsec
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *IpsecProfileCreateIpsecPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IpsecProfileCreateIpsecPost201Response
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileCreateIpsecPost(ctx context.Context, requestBody *IpsecProfileCreateIpsecPostRequest) (*http.Response, *IpsecProfileCreateIpsecPost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/ipsec")
	request.body = requestBody
	request.authenticated = true

	out := &IpsecProfileCreateIpsecPost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	IpsecProfileQueryListPostRequestAttributesSlice []*string

	IpsecProfileQueryListPostRequestExtraFiltersSlice []*IpsecProfileQueryListPostRequestExtraFilters

	IpsecProfileQueryListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	IpsecProfileQueryListPostRequestExtraNotFiltersSlice []*IpsecProfileQueryListPostRequestExtraNotFilters

	IpsecProfileQueryListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	IpsecProfileQueryListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	IpsecProfileQueryListPostRequestFiltersSlice []*IpsecProfileQueryListPostRequestFilters

	IpsecProfileQueryListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	IpsecProfileQueryListPostRequestFullTextSearchFieldsSlice []*string

	IpsecProfileQueryListPostRequestFullTextSearch struct {
		Fields IpsecProfileQueryListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                   `json:"type,omitempty"`
		Value  *string                                                   `json:"value,omitempty"`
	}

	IpsecProfileQueryListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	IpsecProfileQueryListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	IpsecProfileQueryListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                       `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                     `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                       `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                       `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                     `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                       `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                       `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                       `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                       `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                       `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                       `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                     `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                       `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                     `json:"auth_type,omitempty"`
		ForwardingType                *string                                                     `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                     `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                     `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *IpsecProfileQueryListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                     `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                       `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                       `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                       `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *IpsecProfileQueryListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                     `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                     `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                     `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                     `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                     `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                     `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                     `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                     `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                     `json:"localUser_userSource,omitempty"`
	}

	IpsecProfileQueryListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	IpsecProfileQueryListPostRequest struct {
		Attributes      IpsecProfileQueryListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                              `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                `json:"expandDomains,omitempty"`
		ExtraFilters    IpsecProfileQueryListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters IpsecProfileQueryListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *IpsecProfileQueryListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         IpsecProfileQueryListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *IpsecProfileQueryListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                 `json:"limit,omitempty"`
		Options         *IpsecProfileQueryListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                 `json:"page,omitempty"`
		Query           *string                                              `json:"query,omitempty"`
		SortInfo        *IpsecProfileQueryListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                 `json:"start,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseListSlice []*IpsecProfileQueryListPost200ResponseList

	IpsecProfileQueryListPost200ResponseListAdvancedOption struct {
		DhcpOpt43Subcode             *float64 `json:"dhcpOpt43Subcode,omitempty"`
		DpdDelay                     *float64 `json:"dpdDelay,omitempty"`
		EnforceNatt                  *string  `json:"enforceNatt,omitempty"`
		FailoverMode                 *string  `json:"failoverMode,omitempty"`
		FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`
		FailoverRetryInterval        *float64 `json:"failoverRetryInterval,omitempty"`
		FailoverRetryPeriod          *float64 `json:"failoverRetryPeriod,omitempty"`
		IpcompEnable                 *string  `json:"ipcompEnable,omitempty"`
		KeepAliveIntval              *float64 `json:"keepAliveIntval,omitempty"`
		ReplayWindow                 *float64 `json:"replayWindow,omitempty"`
		RetryLimit                   *float64 `json:"retryLimit,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseListCmProtocolOption struct {
		CmpDhcpOpt43Subcode          *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`
		CmpDhcpOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`
		CmpRecipient                 *string  `json:"cmpRecipient,omitempty"`
		CmpServerAddr                *string  `json:"cmpServerAddr,omitempty"`
		CmpServerPath                *string  `json:"cmpServerPath,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseListEspSecurityAssociationEspProposalsSlice []*IpsecProfileQueryListPost200ResponseListEspSecurityAssociationEspProposals

	IpsecProfileQueryListPost200ResponseListEspSecurityAssociationEspProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseListEspSecurityAssociation struct {
		EspProposalType *string                                                                         `json:"espProposalType,omitempty"`
		EspProposals    IpsecProfileQueryListPost200ResponseListEspSecurityAssociationEspProposalsSlice `json:"espProposals,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseListIkeSecurityAssociationIkeProposalsSlice []*IpsecProfileQueryListPost200ResponseListIkeSecurityAssociationIkeProposals

	IpsecProfileQueryListPost200ResponseListIkeSecurityAssociationIkeProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
		PrfAlg  *string `json:"prfAlg,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseListIkeSecurityAssociation struct {
		IkeProposalType *string                                                                         `json:"ikeProposalType,omitempty"`
		IkeProposals    IpsecProfileQueryListPost200ResponseListIkeSecurityAssociationIkeProposalsSlice `json:"ikeProposals,omitempty"`
	}

	IpsecProfileQueryListPost200ResponseList struct {
		AdvancedOption         *IpsecProfileQueryListPost200ResponseListAdvancedOption         `json:"advancedOption,omitempty"`
		AuthType               *string                                                         `json:"authType,omitempty"`
		CmProtocolOption       *IpsecProfileQueryListPost200ResponseListCmProtocolOption       `json:"cmProtocolOption,omitempty"`
		CreateDateTime         *int                                                            `json:"createDateTime,omitempty"`
		CreatorID              *string                                                         `json:"creatorId,omitempty"`
		CreatorUsername        *string                                                         `json:"creatorUsername,omitempty"`
		Description            *string                                                         `json:"description,omitempty"`
		DomainID               *string                                                         `json:"domainId,omitempty"`
		EspRekeyTime           *float64                                                        `json:"espRekeyTime,omitempty"`
		EspRekeyTimeUnit       *string                                                         `json:"espRekeyTimeUnit,omitempty"`
		EspSecurityAssociation *IpsecProfileQueryListPost200ResponseListEspSecurityAssociation `json:"espSecurityAssociation,omitempty"`
		ID                     *string                                                         `json:"id,omitempty"`
		IkeRekeyTime           *float64                                                        `json:"ikeRekeyTime,omitempty"`
		IkeRekeyTimeUnit       *string                                                         `json:"ikeRekeyTimeUnit,omitempty"`
		IkeSecurityAssociation *IpsecProfileQueryListPost200ResponseListIkeSecurityAssociation `json:"ikeSecurityAssociation,omitempty"`
		IPMode                 *string                                                         `json:"ipMode,omitempty"`
		ModifiedDateTime       *int                                                            `json:"modifiedDateTime,omitempty"`
		ModifierID             *string                                                         `json:"modifierId,omitempty"`
		ModifierUsername       *string                                                         `json:"modifierUsername,omitempty"`
		Name                   *string                                                         `json:"name,omitempty"`
		PreSharedKey           *string                                                         `json:"preSharedKey,omitempty"`
		ServerAddr             *string                                                         `json:"serverAddr,omitempty"`
	}

	IpsecProfileQueryListPost200Response struct {
		Extra      *json.RawMessage                              `json:"extra,omitempty"`
		FirstIndex *int                                          `json:"firstIndex,omitempty"`
		HasMore    *bool                                         `json:"hasMore,omitempty"`
		List       IpsecProfileQueryListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                          `json:"totalCount,omitempty"`
	}
)

// IpsecProfileQueryListPost: Query a list of IPSEC
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *IpsecProfileQueryListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IpsecProfileQueryListPost200Response
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileQueryListPost(ctx context.Context, requestBody *IpsecProfileQueryListPostRequest) (*http.Response, *IpsecProfileQueryListPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/ipsec/query")
	request.body = requestBody
	request.authenticated = true

	out := &IpsecProfileQueryListPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// IpsecProfileDeleteDelete: Delete a ipsec
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/tunnel/ipsec/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	IpsecProfileRetrieveGet200ResponseAdvancedOption struct {
		DhcpOpt43Subcode             *float64 `json:"dhcpOpt43Subcode,omitempty"`
		DpdDelay                     *float64 `json:"dpdDelay,omitempty"`
		EnforceNatt                  *string  `json:"enforceNatt,omitempty"`
		FailoverMode                 *string  `json:"failoverMode,omitempty"`
		FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`
		FailoverRetryInterval        *float64 `json:"failoverRetryInterval,omitempty"`
		FailoverRetryPeriod          *float64 `json:"failoverRetryPeriod,omitempty"`
		IpcompEnable                 *string  `json:"ipcompEnable,omitempty"`
		KeepAliveIntval              *float64 `json:"keepAliveIntval,omitempty"`
		ReplayWindow                 *float64 `json:"replayWindow,omitempty"`
		RetryLimit                   *float64 `json:"retryLimit,omitempty"`
	}

	IpsecProfileRetrieveGet200ResponseCmProtocolOption struct {
		CmpDhcpOpt43Subcode          *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`
		CmpDhcpOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`
		CmpRecipient                 *string  `json:"cmpRecipient,omitempty"`
		CmpServerAddr                *string  `json:"cmpServerAddr,omitempty"`
		CmpServerPath                *string  `json:"cmpServerPath,omitempty"`
	}

	IpsecProfileRetrieveGet200ResponseEspSecurityAssociationEspProposalsSlice []*IpsecProfileRetrieveGet200ResponseEspSecurityAssociationEspProposals

	IpsecProfileRetrieveGet200ResponseEspSecurityAssociationEspProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
	}

	IpsecProfileRetrieveGet200ResponseEspSecurityAssociation struct {
		EspProposalType *string                                                                   `json:"espProposalType,omitempty"`
		EspProposals    IpsecProfileRetrieveGet200ResponseEspSecurityAssociationEspProposalsSlice `json:"espProposals,omitempty"`
	}

	IpsecProfileRetrieveGet200ResponseIkeSecurityAssociationIkeProposalsSlice []*IpsecProfileRetrieveGet200ResponseIkeSecurityAssociationIkeProposals

	IpsecProfileRetrieveGet200ResponseIkeSecurityAssociationIkeProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
		PrfAlg  *string `json:"prfAlg,omitempty"`
	}

	IpsecProfileRetrieveGet200ResponseIkeSecurityAssociation struct {
		IkeProposalType *string                                                                   `json:"ikeProposalType,omitempty"`
		IkeProposals    IpsecProfileRetrieveGet200ResponseIkeSecurityAssociationIkeProposalsSlice `json:"ikeProposals,omitempty"`
	}

	IpsecProfileRetrieveGet200Response struct {
		AdvancedOption         *IpsecProfileRetrieveGet200ResponseAdvancedOption         `json:"advancedOption,omitempty"`
		AuthType               *string                                                   `json:"authType,omitempty"`
		CmProtocolOption       *IpsecProfileRetrieveGet200ResponseCmProtocolOption       `json:"cmProtocolOption,omitempty"`
		CreateDateTime         *int                                                      `json:"createDateTime,omitempty"`
		CreatorID              *string                                                   `json:"creatorId,omitempty"`
		CreatorUsername        *string                                                   `json:"creatorUsername,omitempty"`
		Description            *string                                                   `json:"description,omitempty"`
		DomainID               *string                                                   `json:"domainId,omitempty"`
		EspRekeyTime           *float64                                                  `json:"espRekeyTime,omitempty"`
		EspRekeyTimeUnit       *string                                                   `json:"espRekeyTimeUnit,omitempty"`
		EspSecurityAssociation *IpsecProfileRetrieveGet200ResponseEspSecurityAssociation `json:"espSecurityAssociation,omitempty"`
		ID                     *string                                                   `json:"id,omitempty"`
		IkeRekeyTime           *float64                                                  `json:"ikeRekeyTime,omitempty"`
		IkeRekeyTimeUnit       *string                                                   `json:"ikeRekeyTimeUnit,omitempty"`
		IkeSecurityAssociation *IpsecProfileRetrieveGet200ResponseIkeSecurityAssociation `json:"ikeSecurityAssociation,omitempty"`
		IPMode                 *string                                                   `json:"ipMode,omitempty"`
		ModifiedDateTime       *int                                                      `json:"modifiedDateTime,omitempty"`
		ModifierID             *string                                                   `json:"modifierId,omitempty"`
		ModifierUsername       *string                                                   `json:"modifierUsername,omitempty"`
		Name                   *string                                                   `json:"name,omitempty"`
		PreSharedKey           *string                                                   `json:"preSharedKey,omitempty"`
		ServerAddr             *string                                                   `json:"serverAddr,omitempty"`
	}
)

// IpsecProfileRetrieveGet: Retrieve a IPSEC
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IpsecProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *IpsecProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/tunnel/ipsec/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &IpsecProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	IpsecProfileModifyBasicPatchRequest struct {
		AuthType         *string  `json:"authType,omitempty"`
		Description      *string  `json:"description,omitempty"`
		DomainID         *string  `json:"domainId,omitempty"`
		EspRekeyTime     *float64 `json:"espRekeyTime,omitempty"`
		EspRekeyTimeUnit *string  `json:"espRekeyTimeUnit,omitempty"`
		ID               *string  `json:"id,omitempty"`
		IkeRekeyTime     *float64 `json:"ikeRekeyTime,omitempty"`
		IkeRekeyTimeUnit *string  `json:"ikeRekeyTimeUnit,omitempty"`
		IPMode           *string  `json:"ipMode,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PreSharedKey     *string  `json:"preSharedKey,omitempty"`
		ServerAddr       *string  `json:"serverAddr,omitempty"`
	}
)

// IpsecProfileModifyBasicPatch: Modify a specific ipsec basic
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
// - requestBody: *IpsecProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileModifyBasicPatch(ctx context.Context, id string, requestBody *IpsecProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/ipsec/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	IpsecProfileModifyAdvancedoptionPatchRequest struct {
		DhcpOpt43Subcode             *float64 `json:"dhcpOpt43Subcode,omitempty"`
		DpdDelay                     *float64 `json:"dpdDelay,omitempty"`
		EnforceNatt                  *string  `json:"enforceNatt,omitempty"`
		FailoverMode                 *string  `json:"failoverMode,omitempty"`
		FailoverPrimaryCheckInterval *float64 `json:"failoverPrimaryCheckInterval,omitempty"`
		FailoverRetryInterval        *float64 `json:"failoverRetryInterval,omitempty"`
		FailoverRetryPeriod          *float64 `json:"failoverRetryPeriod,omitempty"`
		IpcompEnable                 *string  `json:"ipcompEnable,omitempty"`
		KeepAliveIntval              *float64 `json:"keepAliveIntval,omitempty"`
		ReplayWindow                 *float64 `json:"replayWindow,omitempty"`
		RetryLimit                   *float64 `json:"retryLimit,omitempty"`
	}
)

// IpsecProfileModifyAdvancedoptionPatch: Modify a specific ipsec advancedOption
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
// - requestBody: *IpsecProfileModifyAdvancedoptionPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileModifyAdvancedoptionPatch(ctx context.Context, id string, requestBody *IpsecProfileModifyAdvancedoptionPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/ipsec/{id}/advancedOption")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	IpsecProfileModifyCmprotocoloptionPatchRequest struct {
		CmpDhcpOpt43Subcode          *float64 `json:"cmpDhcpOpt43Subcode,omitempty"`
		CmpDhcpOpt43SubcodeRecipient *float64 `json:"cmpDhcpOpt43SubcodeRecipient,omitempty"`
		CmpRecipient                 *string  `json:"cmpRecipient,omitempty"`
		CmpServerAddr                *string  `json:"cmpServerAddr,omitempty"`
		CmpServerPath                *string  `json:"cmpServerPath,omitempty"`
	}
)

// IpsecProfileModifyCmprotocoloptionPatch: Modify a specific ipsec cmProtocolOption
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
// - requestBody: *IpsecProfileModifyCmprotocoloptionPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileModifyCmprotocoloptionPatch(ctx context.Context, id string, requestBody *IpsecProfileModifyCmprotocoloptionPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/ipsec/{id}/cmProtocolOption")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	IpsecProfileModifyEspsecurityassociationPatchRequestEspProposalsSlice []*IpsecProfileModifyEspsecurityassociationPatchRequestEspProposals

	IpsecProfileModifyEspsecurityassociationPatchRequestEspProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
	}

	IpsecProfileModifyEspsecurityassociationPatchRequest struct {
		EspProposalType *string                                                               `json:"espProposalType,omitempty"`
		EspProposals    IpsecProfileModifyEspsecurityassociationPatchRequestEspProposalsSlice `json:"espProposals,omitempty"`
	}
)

// IpsecProfileModifyEspsecurityassociationPatch: Modify a specific ipsec espSecurityAssociation
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
// - requestBody: *IpsecProfileModifyEspsecurityassociationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileModifyEspsecurityassociationPatch(ctx context.Context, id string, requestBody *IpsecProfileModifyEspsecurityassociationPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/ipsec/{id}/espSecurityAssociation")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	IpsecProfileModifyIkesecurityassociationPatchRequestIkeProposalsSlice []*IpsecProfileModifyIkesecurityassociationPatchRequestIkeProposals

	IpsecProfileModifyIkesecurityassociationPatchRequestIkeProposals struct {
		AuthAlg *string `json:"authAlg,omitempty"`
		DhGroup *string `json:"dhGroup,omitempty"`
		EncAlg  *string `json:"encAlg,omitempty"`
		PrfAlg  *string `json:"prfAlg,omitempty"`
	}

	IpsecProfileModifyIkesecurityassociationPatchRequest struct {
		IkeProposalType *string                                                               `json:"ikeProposalType,omitempty"`
		IkeProposals    IpsecProfileModifyIkesecurityassociationPatchRequestIkeProposalsSlice `json:"ikeProposals,omitempty"`
	}
)

// IpsecProfileModifyIkesecurityassociationPatch: Modify a specific ipsec ikeSecurityAssociation
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): IPSec Tunnel ID
// - requestBody: *IpsecProfileModifyIkesecurityassociationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) IpsecProfileModifyIkesecurityassociationPatch(ctx context.Context, id string, requestBody *IpsecProfileModifyIkesecurityassociationPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/ipsec/{id}/ikeSecurityAssociation")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	L3RoamingModifyL3RoamingBasicConfigurationPatchRequest struct {
		CriteriaType   *string  `json:"criteriaType,omitempty"`
		FeatureEnabled *float64 `json:"featureEnabled,omitempty"`
	}
)

// L3RoamingModifyL3RoamingBasicConfigurationPatch: Use this API command to modify L3 Roaming basic configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *L3RoamingModifyL3RoamingBasicConfigurationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L3RoamingModifyL3RoamingBasicConfigurationPatch(ctx context.Context, requestBody *L3RoamingModifyL3RoamingBasicConfigurationPatchRequest) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/l3Roaming")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	L3RoamingModifyL3RoamingBasicConfigurationPatch1RequestSlice []*L3RoamingModifyL3RoamingBasicConfigurationPatch1Request

	L3RoamingModifyL3RoamingBasicConfigurationPatch1Request struct {
		Activated       *float64 `json:"activated,omitempty"`
		FirmwareVersion *string  `json:"firmwareVersion,omitempty"`
		Key             *string  `json:"key,omitempty"`
		Name            *string  `json:"name,omitempty"`
		Value           *string  `json:"value,omitempty"`
	}
)

// L3RoamingModifyL3RoamingBasicConfigurationPatch1: Use this API command to modify L3 Roaming basic configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *L3RoamingModifyL3RoamingBasicConfigurationPatch1RequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) L3RoamingModifyL3RoamingBasicConfigurationPatch1(ctx context.Context, requestBody L3RoamingModifyL3RoamingBasicConfigurationPatch1RequestSlice) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/l3Roaming/dataPlanes")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	L3RoamingRetriveL3RoamingConfigurationPost200ResponseDataPlanesSlice []*L3RoamingRetriveL3RoamingConfigurationPost200ResponseDataPlanes

	L3RoamingRetriveL3RoamingConfigurationPost200ResponseDataPlanes struct {
		Activated       *float64 `json:"activated,omitempty"`
		FirmwareVersion *string  `json:"firmwareVersion,omitempty"`
		Key             *string  `json:"key,omitempty"`
		Name            *string  `json:"name,omitempty"`
		Value           *string  `json:"value,omitempty"`
	}

	L3RoamingRetriveL3RoamingConfigurationPost200Response struct {
		CriteriaType   *string                                                              `json:"criteriaType,omitempty"`
		DataPlanes     L3RoamingRetriveL3RoamingConfigurationPost200ResponseDataPlanesSlice `json:"dataPlanes,omitempty"`
		FeatureEnabled *float64                                                             `json:"featureEnabled,omitempty"`
	}
)

// L3RoamingRetriveL3RoamingConfigurationPost: Use this API command to retrieve L3 Roaming configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L3RoamingRetriveL3RoamingConfigurationPost200Response
// - error: Error seen or nil if none
func (p *Profiles) L3RoamingRetriveL3RoamingConfigurationPost(ctx context.Context) (*http.Response, *L3RoamingRetriveL3RoamingConfigurationPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/l3Roaming/query")
	request.authenticated = true

	out := &L3RoamingRetriveL3RoamingConfigurationPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusgreTunnelProfileDeleteDelete1RequestIDListSlice []*string

	RuckusgreTunnelProfileDeleteDelete1Request struct {
		IDList RuckusgreTunnelProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// RuckusgreTunnelProfileDeleteDelete1: Use this API command to delete multiple RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusgreTunnelProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileDeleteDelete1(ctx context.Context, requestBody *RuckusgreTunnelProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/tunnel/ruckusgre")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	RuckusgreTunnelProfileRetrieveListGet200ResponseListSlice []*RuckusgreTunnelProfileRetrieveListGet200ResponseList

	RuckusgreTunnelProfileRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusgreTunnelProfileRetrieveListGet200Response struct {
		CreateDateTime   *int                                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                   `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                                          `json:"extra,omitempty"`
		FirstIndex       *int                                                      `json:"firstIndex,omitempty"`
		HasMore          *bool                                                     `json:"hasMore,omitempty"`
		List             RuckusgreTunnelProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                                   `json:"modifierUsername,omitempty"`
		TotalCount       *int                                                      `json:"totalCount,omitempty"`
	}
)

// RuckusgreTunnelProfileRetrieveListGet: Use this API command to retrieve a list of RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusgreTunnelProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileRetrieveListGet(ctx context.Context) (*http.Response, *RuckusgreTunnelProfileRetrieveListGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/tunnel/ruckusgre")
	request.authenticated = true

	out := &RuckusgreTunnelProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusgreTunnelProfileCreatePostRequest struct {
		Description            *string  `json:"description,omitempty"`
		DomainID               *string  `json:"domainId,omitempty"`
		EnableTunnelEncryption *bool    `json:"enableTunnelEncryption,omitempty"`
		ID                     *string  `json:"id,omitempty"`
		Name                   *string  `json:"name,omitempty"`
		TunnelMode             *string  `json:"tunnelMode,omitempty"`
		TunnelMtuAutoEnabled   *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize          *float64 `json:"tunnelMtuSize,omitempty"`
	}

	RuckusgreTunnelProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// RuckusgreTunnelProfileCreatePost: Use this API command to create RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusgreTunnelProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusgreTunnelProfileCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileCreatePost(ctx context.Context, requestBody *RuckusgreTunnelProfileCreatePostRequest) (*http.Response, *RuckusgreTunnelProfileCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/ruckusgre")
	request.body = requestBody
	request.authenticated = true

	out := &RuckusgreTunnelProfileCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	RuckusgreTunnelProfileQueryListPostRequestAttributesSlice []*string

	RuckusgreTunnelProfileQueryListPostRequestExtraFiltersSlice []*RuckusgreTunnelProfileQueryListPostRequestExtraFilters

	RuckusgreTunnelProfileQueryListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestExtraNotFiltersSlice []*RuckusgreTunnelProfileQueryListPostRequestExtraNotFilters

	RuckusgreTunnelProfileQueryListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestFiltersSlice []*RuckusgreTunnelProfileQueryListPostRequestFilters

	RuckusgreTunnelProfileQueryListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestFullTextSearchFieldsSlice []*string

	RuckusgreTunnelProfileQueryListPostRequestFullTextSearch struct {
		Fields RuckusgreTunnelProfileQueryListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                             `json:"type,omitempty"`
		Value  *string                                                             `json:"value,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                 `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                               `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                 `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                 `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                               `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                 `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                 `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                 `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                 `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                 `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                 `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                               `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                 `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                               `json:"auth_type,omitempty"`
		ForwardingType                *string                                                               `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                               `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                               `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *RuckusgreTunnelProfileQueryListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                               `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                 `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                 `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                 `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *RuckusgreTunnelProfileQueryListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                               `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                               `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                               `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                               `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                               `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                               `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                               `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                               `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                               `json:"localUser_userSource,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPostRequest struct {
		Attributes      RuckusgreTunnelProfileQueryListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                        `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                          `json:"expandDomains,omitempty"`
		ExtraFilters    RuckusgreTunnelProfileQueryListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters RuckusgreTunnelProfileQueryListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *RuckusgreTunnelProfileQueryListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         RuckusgreTunnelProfileQueryListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *RuckusgreTunnelProfileQueryListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                           `json:"limit,omitempty"`
		Options         *RuckusgreTunnelProfileQueryListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                           `json:"page,omitempty"`
		Query           *string                                                        `json:"query,omitempty"`
		SortInfo        *RuckusgreTunnelProfileQueryListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                           `json:"start,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPost200ResponseListSlice []*RuckusgreTunnelProfileQueryListPost200ResponseList

	RuckusgreTunnelProfileQueryListPost200ResponseList struct {
		CreateDateTime         *int     `json:"createDateTime,omitempty"`
		CreatorID              *string  `json:"creatorId,omitempty"`
		CreatorUsername        *string  `json:"creatorUsername,omitempty"`
		Description            *string  `json:"description,omitempty"`
		DomainID               *string  `json:"domainId,omitempty"`
		EnableTunnelEncryption *bool    `json:"enableTunnelEncryption,omitempty"`
		ID                     *string  `json:"id,omitempty"`
		ModifiedDateTime       *int     `json:"modifiedDateTime,omitempty"`
		ModifierID             *string  `json:"modifierId,omitempty"`
		ModifierUsername       *string  `json:"modifierUsername,omitempty"`
		Name                   *string  `json:"name,omitempty"`
		TunnelMode             *string  `json:"tunnelMode,omitempty"`
		TunnelMtuAutoEnabled   *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize          *float64 `json:"tunnelMtuSize,omitempty"`
	}

	RuckusgreTunnelProfileQueryListPost200Response struct {
		Extra      *json.RawMessage                                        `json:"extra,omitempty"`
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       RuckusgreTunnelProfileQueryListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// RuckusgreTunnelProfileQueryListPost: Use this API command to query a list of RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusgreTunnelProfileQueryListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusgreTunnelProfileQueryListPost200Response
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileQueryListPost(ctx context.Context, requestBody *RuckusgreTunnelProfileQueryListPostRequest) (*http.Response, *RuckusgreTunnelProfileQueryListPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/ruckusgre/query")
	request.body = requestBody
	request.authenticated = true

	out := &RuckusgreTunnelProfileQueryListPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// RuckusgreTunnelProfileDeleteDelete: Use this API command to delete RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Ruckus GRE Tunnel ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/tunnel/ruckusgre/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	RuckusgreTunnelProfileRetrieveGet200Response struct {
		CreateDateTime         *int     `json:"createDateTime,omitempty"`
		CreatorID              *string  `json:"creatorId,omitempty"`
		CreatorUsername        *string  `json:"creatorUsername,omitempty"`
		Description            *string  `json:"description,omitempty"`
		DomainID               *string  `json:"domainId,omitempty"`
		EnableTunnelEncryption *bool    `json:"enableTunnelEncryption,omitempty"`
		ID                     *string  `json:"id,omitempty"`
		ModifiedDateTime       *int     `json:"modifiedDateTime,omitempty"`
		ModifierID             *string  `json:"modifierId,omitempty"`
		ModifierUsername       *string  `json:"modifierUsername,omitempty"`
		Name                   *string  `json:"name,omitempty"`
		TunnelMode             *string  `json:"tunnelMode,omitempty"`
		TunnelMtuAutoEnabled   *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize          *float64 `json:"tunnelMtuSize,omitempty"`
	}
)

// RuckusgreTunnelProfileRetrieveGet: Use this API command to retrieve RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Ruckus GRE Tunnel ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusgreTunnelProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *RuckusgreTunnelProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/tunnel/ruckusgre/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &RuckusgreTunnelProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusgreTunnelProfileModifyBasicPatchRequest struct {
		Description            *string  `json:"description,omitempty"`
		DomainID               *string  `json:"domainId,omitempty"`
		EnableTunnelEncryption *bool    `json:"enableTunnelEncryption,omitempty"`
		ID                     *string  `json:"id,omitempty"`
		Name                   *string  `json:"name,omitempty"`
		TunnelMode             *string  `json:"tunnelMode,omitempty"`
		TunnelMtuAutoEnabled   *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize          *float64 `json:"tunnelMtuSize,omitempty"`
	}
)

// RuckusgreTunnelProfileModifyBasicPatch: Use this API command to modify the basic information of RuckusGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Ruckus GRE Tunnel ID
// - requestBody: *RuckusgreTunnelProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) RuckusgreTunnelProfileModifyBasicPatch(ctx context.Context, id string, requestBody *RuckusgreTunnelProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/ruckusgre/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	SoftgreTunnelProfileDeleteDelete1RequestIDListSlice []*string

	SoftgreTunnelProfileDeleteDelete1Request struct {
		IDList SoftgreTunnelProfileDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// SoftgreTunnelProfileDeleteDelete1: Use this API command to delete multiple SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SoftgreTunnelProfileDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileDeleteDelete1(ctx context.Context, requestBody *SoftgreTunnelProfileDeleteDelete1Request) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/tunnel/softgre")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	SoftgreTunnelProfileRetrieveListGet200ResponseListSlice []*SoftgreTunnelProfileRetrieveListGet200ResponseList

	SoftgreTunnelProfileRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	SoftgreTunnelProfileRetrieveListGet200Response struct {
		CreateDateTime   *int                                                    `json:"createDateTime,omitempty"`
		CreatorID        *string                                                 `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                 `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                                        `json:"extra,omitempty"`
		FirstIndex       *int                                                    `json:"firstIndex,omitempty"`
		HasMore          *bool                                                   `json:"hasMore,omitempty"`
		List             SoftgreTunnelProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                                    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                 `json:"modifierId,omitempty"`
		ModifierUsername *string                                                 `json:"modifierUsername,omitempty"`
		TotalCount       *int                                                    `json:"totalCount,omitempty"`
	}
)

// SoftgreTunnelProfileRetrieveListGet: Use this API command to retrieve a list of SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SoftgreTunnelProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileRetrieveListGet(ctx context.Context) (*http.Response, *SoftgreTunnelProfileRetrieveListGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/tunnel/softgre")
	request.authenticated = true

	out := &SoftgreTunnelProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SoftgreTunnelProfileCreatePostRequest struct {
		Description             *string  `json:"description,omitempty"`
		DomainID                *string  `json:"domainId,omitempty"`
		ForceDisassociateClient *bool    `json:"forceDisassociateClient,omitempty"`
		ID                      *string  `json:"id,omitempty"`
		IPMode                  *string  `json:"ipMode,omitempty"`
		KeepAlivePeriod         *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry          *float64 `json:"keepAliveRetry,omitempty"`
		Name                    *string  `json:"name,omitempty"`
		PrimaryGateway          *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway        *string  `json:"secondaryGateway,omitempty"`
		TunnelMtuAutoEnabled    *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize           *float64 `json:"tunnelMtuSize,omitempty"`
	}

	SoftgreTunnelProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// SoftgreTunnelProfileCreatePost: Use this API command to create SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SoftgreTunnelProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SoftgreTunnelProfileCreatePost201Response
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileCreatePost(ctx context.Context, requestBody *SoftgreTunnelProfileCreatePostRequest) (*http.Response, *SoftgreTunnelProfileCreatePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/softgre")
	request.body = requestBody
	request.authenticated = true

	out := &SoftgreTunnelProfileCreatePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	SoftgreTunnelProfileQueryListPostRequestAttributesSlice []*string

	SoftgreTunnelProfileQueryListPostRequestExtraFiltersSlice []*SoftgreTunnelProfileQueryListPostRequestExtraFilters

	SoftgreTunnelProfileQueryListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestExtraNotFiltersSlice []*SoftgreTunnelProfileQueryListPostRequestExtraNotFilters

	SoftgreTunnelProfileQueryListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestFiltersSlice []*SoftgreTunnelProfileQueryListPostRequestFilters

	SoftgreTunnelProfileQueryListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestFullTextSearchFieldsSlice []*string

	SoftgreTunnelProfileQueryListPostRequestFullTextSearch struct {
		Fields SoftgreTunnelProfileQueryListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                           `json:"type,omitempty"`
		Value  *string                                                           `json:"value,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                               `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                             `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                               `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                               `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                             `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                               `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                               `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                               `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                               `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                               `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                               `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                             `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                               `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                             `json:"auth_type,omitempty"`
		ForwardingType                *string                                                             `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                             `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                             `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *SoftgreTunnelProfileQueryListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                             `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                               `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                               `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                               `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *SoftgreTunnelProfileQueryListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                             `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                             `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                             `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                             `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                             `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                             `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                             `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                             `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                             `json:"localUser_userSource,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	SoftgreTunnelProfileQueryListPostRequest struct {
		Attributes      SoftgreTunnelProfileQueryListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                      `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                        `json:"expandDomains,omitempty"`
		ExtraFilters    SoftgreTunnelProfileQueryListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters SoftgreTunnelProfileQueryListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *SoftgreTunnelProfileQueryListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         SoftgreTunnelProfileQueryListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *SoftgreTunnelProfileQueryListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                         `json:"limit,omitempty"`
		Options         *SoftgreTunnelProfileQueryListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                         `json:"page,omitempty"`
		Query           *string                                                      `json:"query,omitempty"`
		SortInfo        *SoftgreTunnelProfileQueryListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                         `json:"start,omitempty"`
	}

	SoftgreTunnelProfileQueryListPost200ResponseListSlice []*SoftgreTunnelProfileQueryListPost200ResponseList

	SoftgreTunnelProfileQueryListPost200ResponseList struct {
		CreateDateTime          *int     `json:"createDateTime,omitempty"`
		CreatorID               *string  `json:"creatorId,omitempty"`
		CreatorUsername         *string  `json:"creatorUsername,omitempty"`
		Description             *string  `json:"description,omitempty"`
		DomainID                *string  `json:"domainId,omitempty"`
		ForceDisassociateClient *bool    `json:"forceDisassociateClient,omitempty"`
		ID                      *string  `json:"id,omitempty"`
		IPMode                  *string  `json:"ipMode,omitempty"`
		KeepAlivePeriod         *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry          *float64 `json:"keepAliveRetry,omitempty"`
		ModifiedDateTime        *int     `json:"modifiedDateTime,omitempty"`
		ModifierID              *string  `json:"modifierId,omitempty"`
		ModifierUsername        *string  `json:"modifierUsername,omitempty"`
		Name                    *string  `json:"name,omitempty"`
		PrimaryGateway          *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway        *string  `json:"secondaryGateway,omitempty"`
		TunnelMtuAutoEnabled    *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize           *float64 `json:"tunnelMtuSize,omitempty"`
	}

	SoftgreTunnelProfileQueryListPost200Response struct {
		Extra      *json.RawMessage                                      `json:"extra,omitempty"`
		FirstIndex *int                                                  `json:"firstIndex,omitempty"`
		HasMore    *bool                                                 `json:"hasMore,omitempty"`
		List       SoftgreTunnelProfileQueryListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                  `json:"totalCount,omitempty"`
	}
)

// SoftgreTunnelProfileQueryListPost: Use this API command to query a list of SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SoftgreTunnelProfileQueryListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SoftgreTunnelProfileQueryListPost200Response
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileQueryListPost(ctx context.Context, requestBody *SoftgreTunnelProfileQueryListPostRequest) (*http.Response, *SoftgreTunnelProfileQueryListPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/tunnel/softgre/query")
	request.body = requestBody
	request.authenticated = true

	out := &SoftgreTunnelProfileQueryListPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// SoftgreTunnelProfileDeleteDelete: Use this API command to delete SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Soft GRE Tunnel ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/tunnel/softgre/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	SoftgreTunnelProfileRetrieveGet200Response struct {
		CreateDateTime          *int     `json:"createDateTime,omitempty"`
		CreatorID               *string  `json:"creatorId,omitempty"`
		CreatorUsername         *string  `json:"creatorUsername,omitempty"`
		Description             *string  `json:"description,omitempty"`
		DomainID                *string  `json:"domainId,omitempty"`
		ForceDisassociateClient *bool    `json:"forceDisassociateClient,omitempty"`
		ID                      *string  `json:"id,omitempty"`
		IPMode                  *string  `json:"ipMode,omitempty"`
		KeepAlivePeriod         *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry          *float64 `json:"keepAliveRetry,omitempty"`
		ModifiedDateTime        *int     `json:"modifiedDateTime,omitempty"`
		ModifierID              *string  `json:"modifierId,omitempty"`
		ModifierUsername        *string  `json:"modifierUsername,omitempty"`
		Name                    *string  `json:"name,omitempty"`
		PrimaryGateway          *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway        *string  `json:"secondaryGateway,omitempty"`
		TunnelMtuAutoEnabled    *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize           *float64 `json:"tunnelMtuSize,omitempty"`
	}
)

// SoftgreTunnelProfileRetrieveGet: Use this API command to retrieve SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Soft GRE Tunnel ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SoftgreTunnelProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *SoftgreTunnelProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/tunnel/softgre/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &SoftgreTunnelProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SoftgreTunnelProfileModifyBasicPatchRequest struct {
		Description             *string  `json:"description,omitempty"`
		DomainID                *string  `json:"domainId,omitempty"`
		ForceDisassociateClient *bool    `json:"forceDisassociateClient,omitempty"`
		ID                      *string  `json:"id,omitempty"`
		KeepAlivePeriod         *float64 `json:"keepAlivePeriod,omitempty"`
		KeepAliveRetry          *float64 `json:"keepAliveRetry,omitempty"`
		Name                    *string  `json:"name,omitempty"`
		PrimaryGateway          *string  `json:"primaryGateway,omitempty"`
		SecondaryGateway        *string  `json:"secondaryGateway,omitempty"`
		TunnelMtuAutoEnabled    *string  `json:"tunnelMtuAutoEnabled,omitempty"`
		TunnelMtuSize           *float64 `json:"tunnelMtuSize,omitempty"`
	}
)

// SoftgreTunnelProfileModifyBasicPatch: Use this API command to modify the basic information of SoftGRE tunnel profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Soft GRE Tunnel ID
// - requestBody: *SoftgreTunnelProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) SoftgreTunnelProfileModifyBasicPatch(ctx context.Context, id string, requestBody *SoftgreTunnelProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/tunnel/softgre/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	UserTrafficProfileDeleteDeleteRequestIDListSlice []*string

	UserTrafficProfileDeleteDeleteRequest struct {
		IDList UserTrafficProfileDeleteDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// UserTrafficProfileDeleteDelete: Use this API command to delete a list of traffic profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *UserTrafficProfileDeleteDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileDeleteDelete(ctx context.Context, requestBody *UserTrafficProfileDeleteDeleteRequest) (*http.Response, []byte, error) {
	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/utp")
	request.body = requestBody
	request.authenticated = true

	return p.client.doRequest(request, 204, nil)
}

type (
	UserTrafficProfileRetrieveListUserTrafficProfileGet200ResponseListSlice []*UserTrafficProfileRetrieveListUserTrafficProfileGet200ResponseList

	UserTrafficProfileRetrieveListUserTrafficProfileGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileGet200Response struct {
		CreateDateTime   *int                                                                    `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                 `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                 `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                                                        `json:"extra,omitempty"`
		FirstIndex       *int                                                                    `json:"firstIndex,omitempty"`
		HasMore          *bool                                                                   `json:"hasMore,omitempty"`
		List             UserTrafficProfileRetrieveListUserTrafficProfileGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                                                    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                 `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                 `json:"modifierUsername,omitempty"`
		TotalCount       *int                                                                    `json:"totalCount,omitempty"`
	}
)

// UserTrafficProfileRetrieveListUserTrafficProfileGet: Use this API command to retrieve a list of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *UserTrafficProfileRetrieveListUserTrafficProfileGet200Response
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileRetrieveListUserTrafficProfileGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *UserTrafficProfileRetrieveListUserTrafficProfileGet200Response, error) {
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

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/utp")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}

	out := &UserTrafficProfileRetrieveListUserTrafficProfileGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	UserTrafficProfileCreateUserTrafficProfilePostRequestDownlinkRateLimiting struct {
		DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
		DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileCreateUserTrafficProfilePostRequestIPAclRulesSlice []*UserTrafficProfileCreateUserTrafficProfilePostRequestIPAclRules

	UserTrafficProfileCreateUserTrafficProfilePostRequestIPAclRules struct {
		Action                      *string  `json:"action,omitempty"`
		Description                 *string  `json:"description,omitempty"`
		DestinationIP               *string  `json:"destinationIp,omitempty"`
		DestinationIPMask           *string  `json:"destinationIpMask,omitempty"`
		DestinationMaxPort          *float64 `json:"destinationMaxPort,omitempty"`
		DestinationMinPort          *float64 `json:"destinationMinPort,omitempty"`
		Direction                   *string  `json:"direction,omitempty"`
		DownlinkRateLimitingEnabled *bool    `json:"downlinkRateLimitingEnabled,omitempty"`
		DownlinkRateLimitingMbps    *float64 `json:"downlinkRateLimitingMbps,omitempty"`
		EnableDestinationIPSubnet   *bool    `json:"enableDestinationIpSubnet,omitempty"`
		EnableDestinationPortRange  *bool    `json:"enableDestinationPortRange,omitempty"`
		EnableSourceIPSubnet        *bool    `json:"enableSourceIpSubnet,omitempty"`
		EnableSourcePortRange       *bool    `json:"enableSourcePortRange,omitempty"`
		Priority                    *float64 `json:"priority,omitempty"`
		Protocol                    *string  `json:"protocol,omitempty"`
		SourceIP                    *string  `json:"sourceIp,omitempty"`
		SourceIPMask                *string  `json:"sourceIpMask,omitempty"`
		SourceMaxPort               *float64 `json:"sourceMaxPort,omitempty"`
		SourceMinPort               *float64 `json:"sourceMinPort,omitempty"`
		UplinkRateLimitingEnabled   *bool    `json:"uplinkRateLimitingEnabled,omitempty"`
		UplinkRateLimitingMbps      *float64 `json:"uplinkRateLimitingMbps,omitempty"`
	}

	UserTrafficProfileCreateUserTrafficProfilePostRequestUplinkRateLimiting struct {
		UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
		UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileCreateUserTrafficProfilePostRequest struct {
		AppPolicyID          *string                                                                    `json:"appPolicyId,omitempty"`
		DefaultAction        *string                                                                    `json:"defaultAction,omitempty"`
		Description          *string                                                                    `json:"description,omitempty"`
		DomainID             *string                                                                    `json:"domainId,omitempty"`
		DownlinkRateLimiting *UserTrafficProfileCreateUserTrafficProfilePostRequestDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
		IPAclRules           UserTrafficProfileCreateUserTrafficProfilePostRequestIPAclRulesSlice       `json:"ipAclRules,omitempty"`
		MvnoID               *string                                                                    `json:"mvnoId,omitempty"`
		Name                 *string                                                                    `json:"name,omitempty"`
		UplinkRateLimiting   *UserTrafficProfileCreateUserTrafficProfilePostRequestUplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
	}

	UserTrafficProfileCreateUserTrafficProfilePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// UserTrafficProfileCreateUserTrafficProfilePost: Use this API command to create a new user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *UserTrafficProfileCreateUserTrafficProfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *UserTrafficProfileCreateUserTrafficProfilePost201Response
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileCreateUserTrafficProfilePost(ctx context.Context, requestBody *UserTrafficProfileCreateUserTrafficProfilePostRequest) (*http.Response, *UserTrafficProfileCreateUserTrafficProfilePost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/utp")
	request.body = requestBody
	request.authenticated = true

	out := &UserTrafficProfileCreateUserTrafficProfilePost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	UserTrafficProfileClonePostRequest struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}

	UserTrafficProfileClonePost200Response struct {
		NewID   *string `json:"newId,omitempty"`
		NewName *string `json:"newName,omitempty"`
		OldID   *string `json:"oldId,omitempty"`
		OldName *string `json:"oldName,omitempty"`
	}
)

// UserTrafficProfileClonePost: Use this API command to copy a traffic profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *UserTrafficProfileClonePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *UserTrafficProfileClonePost200Response
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileClonePost(ctx context.Context, id string, requestBody *UserTrafficProfileClonePostRequest) (*http.Response, *UserTrafficProfileClonePost200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/utp/clone/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &UserTrafficProfileClonePost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestAttributesSlice []*string

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraFiltersSlice []*UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraFilters

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraNotFiltersSlice []*UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraNotFilters

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFiltersSlice []*UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFilters

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFullTextSearch struct {
		Fields UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                             `json:"type,omitempty"`
		Value  *string                                                                                             `json:"value,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                 `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                               `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                 `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                 `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                               `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                 `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                 `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                 `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                 `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                 `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                 `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                               `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                 `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                               `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                               `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                               `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                               `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                               `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                 `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                 `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                 `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                               `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                               `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                               `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                               `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                               `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                               `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                               `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                               `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                               `json:"localUser_userSource,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequest struct {
		Attributes      UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                        `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                          `json:"expandDomains,omitempty"`
		ExtraFilters    UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                           `json:"limit,omitempty"`
		Options         *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                           `json:"page,omitempty"`
		Query           *string                                                                                        `json:"query,omitempty"`
		SortInfo        *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                           `json:"start,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListSlice []*UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseList

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListDownlinkRateLimiting struct {
		DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
		DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListIPAclRulesSlice []*UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListIPAclRules

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListIPAclRules struct {
		Action                      *string  `json:"action,omitempty"`
		Description                 *string  `json:"description,omitempty"`
		DestinationIP               *string  `json:"destinationIp,omitempty"`
		DestinationIPMask           *string  `json:"destinationIpMask,omitempty"`
		DestinationMaxPort          *float64 `json:"destinationMaxPort,omitempty"`
		DestinationMinPort          *float64 `json:"destinationMinPort,omitempty"`
		Direction                   *string  `json:"direction,omitempty"`
		DownlinkRateLimitingEnabled *bool    `json:"downlinkRateLimitingEnabled,omitempty"`
		DownlinkRateLimitingMbps    *float64 `json:"downlinkRateLimitingMbps,omitempty"`
		EnableDestinationIPSubnet   *bool    `json:"enableDestinationIpSubnet,omitempty"`
		EnableDestinationPortRange  *bool    `json:"enableDestinationPortRange,omitempty"`
		EnableSourceIPSubnet        *bool    `json:"enableSourceIpSubnet,omitempty"`
		EnableSourcePortRange       *bool    `json:"enableSourcePortRange,omitempty"`
		Priority                    *float64 `json:"priority,omitempty"`
		Protocol                    *string  `json:"protocol,omitempty"`
		SourceIP                    *string  `json:"sourceIp,omitempty"`
		SourceIPMask                *string  `json:"sourceIpMask,omitempty"`
		SourceMaxPort               *float64 `json:"sourceMaxPort,omitempty"`
		SourceMinPort               *float64 `json:"sourceMinPort,omitempty"`
		UplinkRateLimitingEnabled   *bool    `json:"uplinkRateLimitingEnabled,omitempty"`
		UplinkRateLimitingMbps      *float64 `json:"uplinkRateLimitingMbps,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListUplinkRateLimiting struct {
		UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
		UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseList struct {
		AppPolicyID          *string                                                                                                 `json:"appPolicyId,omitempty"`
		CreateDateTime       *int                                                                                                    `json:"createDateTime,omitempty"`
		CreatorID            *string                                                                                                 `json:"creatorId,omitempty"`
		CreatorUsername      *string                                                                                                 `json:"creatorUsername,omitempty"`
		DefaultAction        *string                                                                                                 `json:"defaultAction,omitempty"`
		Description          *string                                                                                                 `json:"description,omitempty"`
		DomainID             *string                                                                                                 `json:"domainId,omitempty"`
		DownlinkRateLimiting *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
		ID                   *string                                                                                                 `json:"id,omitempty"`
		IPAclRules           UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListIPAclRulesSlice       `json:"ipAclRules,omitempty"`
		IsFactoryDefault     *bool                                                                                                   `json:"isFactoryDefault,omitempty"`
		ModifiedDateTime     *int                                                                                                    `json:"modifiedDateTime,omitempty"`
		ModifierID           *string                                                                                                 `json:"modifierId,omitempty"`
		ModifierUsername     *string                                                                                                 `json:"modifierUsername,omitempty"`
		MvnoID               *string                                                                                                 `json:"mvnoId,omitempty"`
		Name                 *string                                                                                                 `json:"name,omitempty"`
		UplinkRateLimiting   *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListUplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
	}

	UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                        `json:"extra,omitempty"`
		FirstIndex *int                                                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                   `json:"hasMore,omitempty"`
		List       UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                    `json:"totalCount,omitempty"`
	}
)

// UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost: Use this API command to retrieve a list of User Traffic Profile by query criteria.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost(ctx context.Context, requestBody *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPostRequest) (*http.Response, *UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/utp/query")
	request.body = requestBody
	request.authenticated = true

	out := &UserTrafficProfileRetrieveListUserTrafficProfileByQueryCritariaPost200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// UserTrafficProfileDeleteUserTrafficProfileDelete: Use this API command to delete an user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileDeleteUserTrafficProfileDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/utp/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseDownlinkRateLimiting struct {
		DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
		DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRulesSlice []*UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRules

	UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRulesDownlinkRateLimiting struct {
		DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
		DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRulesUplinkRateLimiting struct {
		UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
		UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRules struct {
		Action                      *string                                                                                   `json:"action,omitempty"`
		Description                 *string                                                                                   `json:"description,omitempty"`
		DestinationIP               *string                                                                                   `json:"destinationIp,omitempty"`
		DestinationIPMask           *string                                                                                   `json:"destinationIpMask,omitempty"`
		DestinationMaxPort          *float64                                                                                  `json:"destinationMaxPort,omitempty"`
		DestinationMinPort          *float64                                                                                  `json:"destinationMinPort,omitempty"`
		Direction                   *string                                                                                   `json:"direction,omitempty"`
		DownlinkRateLimiting        *UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRulesDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
		DownlinkRateLimitingEnabled *bool                                                                                     `json:"downlinkRateLimitingEnabled,omitempty"`
		DownlinkRateLimitingMbps    *float64                                                                                  `json:"downlinkRateLimitingMbps,omitempty"`
		EnableDestinationIPSubnet   *bool                                                                                     `json:"enableDestinationIpSubnet,omitempty"`
		EnableDestinationPortRange  *bool                                                                                     `json:"enableDestinationPortRange,omitempty"`
		EnableSourceIPSubnet        *bool                                                                                     `json:"enableSourceIpSubnet,omitempty"`
		EnableSourcePortRange       *bool                                                                                     `json:"enableSourcePortRange,omitempty"`
		Priority                    *float64                                                                                  `json:"priority,omitempty"`
		Protocol                    *string                                                                                   `json:"protocol,omitempty"`
		SourceIP                    *string                                                                                   `json:"sourceIp,omitempty"`
		SourceIPMask                *string                                                                                   `json:"sourceIpMask,omitempty"`
		SourceMaxPort               *float64                                                                                  `json:"sourceMaxPort,omitempty"`
		SourceMinPort               *float64                                                                                  `json:"sourceMinPort,omitempty"`
		UplinkRateLimiting          *UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRulesUplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
		UplinkRateLimitingEnabled   *bool                                                                                     `json:"uplinkRateLimitingEnabled,omitempty"`
		UplinkRateLimitingMbps      *float64                                                                                  `json:"uplinkRateLimitingMbps,omitempty"`
	}

	UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseUplinkRateLimiting struct {
		UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
		UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileRetrieveUserTrafficProfileGet200Response struct {
		AppPolicyID          *string                                                                         `json:"appPolicyId,omitempty"`
		CreateDateTime       *int                                                                            `json:"createDateTime,omitempty"`
		CreatorID            *string                                                                         `json:"creatorId,omitempty"`
		CreatorUsername      *string                                                                         `json:"creatorUsername,omitempty"`
		DefaultAction        *string                                                                         `json:"defaultAction,omitempty"`
		Description          *string                                                                         `json:"description,omitempty"`
		DomainID             *string                                                                         `json:"domainId,omitempty"`
		DownlinkRateLimiting *UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
		ID                   *string                                                                         `json:"id,omitempty"`
		IPAclRules           UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseIPAclRulesSlice       `json:"ipAclRules,omitempty"`
		IsFactoryDefault     interface{}                                                                     `json:"isFactoryDefault,omitempty"`
		ModifiedDateTime     *int                                                                            `json:"modifiedDateTime,omitempty"`
		ModifierID           *string                                                                         `json:"modifierId,omitempty"`
		ModifierUsername     *string                                                                         `json:"modifierUsername,omitempty"`
		MvnoID               *string                                                                         `json:"mvnoId,omitempty"`
		Name                 *string                                                                         `json:"name,omitempty"`
		UplinkRateLimiting   *UserTrafficProfileRetrieveUserTrafficProfileGet200ResponseUplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
	}
)

// UserTrafficProfileRetrieveUserTrafficProfileGet: Use this API command to retrieve an user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *UserTrafficProfileRetrieveUserTrafficProfileGet200Response
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileRetrieveUserTrafficProfileGet(ctx context.Context, id string) (*http.Response, *UserTrafficProfileRetrieveUserTrafficProfileGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/utp/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &UserTrafficProfileRetrieveUserTrafficProfileGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	UserTrafficProfileModifyUserTrafficProfilePatchRequest struct {
		AppPolicyID   *string `json:"appPolicyId,omitempty"`
		DefaultAction *string `json:"defaultAction,omitempty"`
		Description   *string `json:"description,omitempty"`
		DomainID      *string `json:"domainId,omitempty"`
		ID            *string `json:"id,omitempty"`
		MvnoID        *string `json:"mvnoId,omitempty"`
		Name          *string `json:"name,omitempty"`
	}
)

// UserTrafficProfileModifyUserTrafficProfilePatch: Use this API command to modify the basic information of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
// - requestBody: *UserTrafficProfileModifyUserTrafficProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileModifyUserTrafficProfilePatch(ctx context.Context, id string, requestBody *UserTrafficProfileModifyUserTrafficProfilePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/utp/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

// UserTrafficProfileDisableDownlinkratelimitingDelete: Use this API command to disable downlink rate limiting of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileDisableDownlinkratelimitingDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/utp/{id}/downlinkRateLimiting")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	UserTrafficProfileModifyDownlinkratelimitingPatchRequest struct {
		DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
		DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
	}
)

// UserTrafficProfileModifyDownlinkratelimitingPatch: Use this API command to modify downlink rate limiting of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
// - requestBody: *UserTrafficProfileModifyDownlinkratelimitingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileModifyDownlinkratelimitingPatch(ctx context.Context, id string, requestBody *UserTrafficProfileModifyDownlinkratelimitingPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/utp/{id}/downlinkRateLimiting")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	UserTrafficProfileModifyIpaclrulesPatchRequestSlice []*UserTrafficProfileModifyIpaclrulesPatchRequest

	UserTrafficProfileModifyIpaclrulesPatchRequestDownlinkRateLimiting struct {
		DownlinkRateLimitingBps     *string `json:"downlinkRateLimitingBps,omitempty"`
		DownlinkRateLimitingEnabled *bool   `json:"downlinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileModifyIpaclrulesPatchRequestUplinkRateLimiting struct {
		UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
		UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
	}

	UserTrafficProfileModifyIpaclrulesPatchRequest struct {
		Action                      *string                                                             `json:"action,omitempty"`
		Description                 *string                                                             `json:"description,omitempty"`
		DestinationIP               *string                                                             `json:"destinationIp,omitempty"`
		DestinationIPMask           *string                                                             `json:"destinationIpMask,omitempty"`
		DestinationMaxPort          *float64                                                            `json:"destinationMaxPort,omitempty"`
		DestinationMinPort          *float64                                                            `json:"destinationMinPort,omitempty"`
		Direction                   *string                                                             `json:"direction,omitempty"`
		DownlinkRateLimiting        *UserTrafficProfileModifyIpaclrulesPatchRequestDownlinkRateLimiting `json:"downlinkRateLimiting,omitempty"`
		DownlinkRateLimitingEnabled *bool                                                               `json:"downlinkRateLimitingEnabled,omitempty"`
		DownlinkRateLimitingMbps    *float64                                                            `json:"downlinkRateLimitingMbps,omitempty"`
		EnableDestinationIPSubnet   *bool                                                               `json:"enableDestinationIpSubnet,omitempty"`
		EnableDestinationPortRange  *bool                                                               `json:"enableDestinationPortRange,omitempty"`
		EnableSourceIPSubnet        *bool                                                               `json:"enableSourceIpSubnet,omitempty"`
		EnableSourcePortRange       *bool                                                               `json:"enableSourcePortRange,omitempty"`
		Priority                    *float64                                                            `json:"priority,omitempty"`
		Protocol                    *string                                                             `json:"protocol,omitempty"`
		SourceIP                    *string                                                             `json:"sourceIp,omitempty"`
		SourceIPMask                *string                                                             `json:"sourceIpMask,omitempty"`
		SourceMaxPort               *float64                                                            `json:"sourceMaxPort,omitempty"`
		SourceMinPort               *float64                                                            `json:"sourceMinPort,omitempty"`
		UplinkRateLimiting          *UserTrafficProfileModifyIpaclrulesPatchRequestUplinkRateLimiting   `json:"uplinkRateLimiting,omitempty"`
		UplinkRateLimitingEnabled   *bool                                                               `json:"uplinkRateLimitingEnabled,omitempty"`
		UplinkRateLimitingMbps      *float64                                                            `json:"uplinkRateLimitingMbps,omitempty"`
	}
)

// UserTrafficProfileModifyIpaclrulesPatch: Use this API command to modify IP ACL rules of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
// - requestBody: *UserTrafficProfileModifyIpaclrulesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileModifyIpaclrulesPatch(ctx context.Context, id string, requestBody UserTrafficProfileModifyIpaclrulesPatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/utp/{id}/ipAclRules")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

// UserTrafficProfileDisableUplinkratelimitingDelete: Use this API command to disable uplink rateLimiting of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileDisableUplinkratelimitingDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/utp/{id}/uplinkRateLimiting")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	UserTrafficProfileModifyUplinkratelimitingPatchRequest struct {
		UplinkRateLimitingBps     *string `json:"uplinkRateLimitingBps,omitempty"`
		UplinkRateLimitingEnabled *bool   `json:"uplinkRateLimitingEnabled,omitempty"`
	}
)

// UserTrafficProfileModifyUplinkratelimitingPatch: Use this API command to modify uplink rate limiting of user traffic profile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): User Traffic Profile ID
// - requestBody: *UserTrafficProfileModifyUplinkratelimitingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) UserTrafficProfileModifyUplinkratelimitingPatch(ctx context.Context, id string, requestBody *UserTrafficProfileModifyUplinkratelimitingPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/utp/{id}/uplinkRateLimiting")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	VdpProfileRetrieveListGet200ResponseListSlice []*VdpProfileRetrieveListGet200ResponseList

	VdpProfileRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	VdpProfileRetrieveListGet200Response struct {
		CreateDateTime   *int                                          `json:"createDateTime,omitempty"`
		CreatorID        *string                                       `json:"creatorId,omitempty"`
		CreatorUsername  *string                                       `json:"creatorUsername,omitempty"`
		Extra            *json.RawMessage                              `json:"extra,omitempty"`
		FirstIndex       *int                                          `json:"firstIndex,omitempty"`
		HasMore          *bool                                         `json:"hasMore,omitempty"`
		List             VdpProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                          `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                       `json:"modifierId,omitempty"`
		ModifierUsername *string                                       `json:"modifierUsername,omitempty"`
		TotalCount       *int                                          `json:"totalCount,omitempty"`
	}
)

// VdpProfileRetrieveListGet: Use this API command to retrieve a list of vdp.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *VdpProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *Profiles) VdpProfileRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *VdpProfileRetrieveListGet200Response, error) {
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

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/vdp")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}

	out := &VdpProfileRetrieveListGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// VdpProfileDeleteDelete: Use this API command to delete an vdp.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) VdpProfileDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/vdp/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	VdpProfileRetrieveGet200Response struct {
		CreateDateTime    *int     `json:"createDateTime,omitempty"`
		CreatorID         *string  `json:"creatorId,omitempty"`
		CreatorUsername   *string  `json:"creatorUsername,omitempty"`
		DataVlan          *float64 `json:"dataVlan,omitempty"`
		ExtIP             *string  `json:"extIp,omitempty"`
		FwVersion         *string  `json:"fwVersion,omitempty"`
		IP                *string  `json:"ip,omitempty"`
		Ipv6              *string  `json:"ipv6,omitempty"`
		IsSupport         *bool    `json:"isSupport,omitempty"`
		LastSeenOn        *string  `json:"lastSeenOn,omitempty"`
		Mac               *string  `json:"mac,omitempty"`
		ManagedBy         *string  `json:"managedBy,omitempty"`
		MgmtExtIP         *string  `json:"mgmtExtIp,omitempty"`
		MgmtIP            *string  `json:"mgmtIp,omitempty"`
		MgmtVlan          *float64 `json:"mgmtVlan,omitempty"`
		Model             *string  `json:"model,omitempty"`
		ModifiedDateTime  *int     `json:"modifiedDateTime,omitempty"`
		ModifierID        *string  `json:"modifierId,omitempty"`
		ModifierUsername  *string  `json:"modifierUsername,omitempty"`
		Name              *string  `json:"name,omitempty"`
		RegistrationState *string  `json:"registrationState,omitempty"`
		SerialNumber      *string  `json:"serialNumber,omitempty"`
		Status            *string  `json:"status,omitempty"`
		Uptime            *string  `json:"uptime,omitempty"`
	}
)

// VdpProfileRetrieveGet: Use this API command to retrieve an vdp.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *VdpProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (p *Profiles) VdpProfileRetrieveGet(ctx context.Context, id string) (*http.Response, *VdpProfileRetrieveGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/vdp/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &VdpProfileRetrieveGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// VdpProfileApprovePut: Use this API command to approve vdp.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) VdpProfileApprovePut(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PUT", "/v5_0/profiles/vdp/{id}/approve")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	ZoneAffinityProfileGetAllZoneAffinityProfilesGet200ResponseListSlice []*ZoneAffinityProfileGetAllZoneAffinityProfilesGet200ResponseList

	ZoneAffinityProfileGetAllZoneAffinityProfilesGet200ResponseListZoneAffinityListSlice []*string

	ZoneAffinityProfileGetAllZoneAffinityProfilesGet200ResponseList struct {
		CreateDateTime   *int                                                                                 `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                              `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                              `json:"creatorUsername,omitempty"`
		Description      *string                                                                              `json:"description,omitempty"`
		ID               *string                                                                              `json:"id,omitempty"`
		ModifiedDateTime *int                                                                                 `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                              `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                              `json:"modifierUsername,omitempty"`
		Name             *string                                                                              `json:"name,omitempty"`
		ZoneAffinityList ZoneAffinityProfileGetAllZoneAffinityProfilesGet200ResponseListZoneAffinityListSlice `json:"zoneAffinityList,omitempty"`
	}

	ZoneAffinityProfileGetAllZoneAffinityProfilesGet200Response struct {
		List ZoneAffinityProfileGetAllZoneAffinityProfilesGet200ResponseListSlice `json:"list,omitempty"`
	}
)

// ZoneAffinityProfileGetAllZoneAffinityProfilesGet: Use this API command to get all zone affinity profiles.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAffinityProfileGetAllZoneAffinityProfilesGet200Response
// - error: Error seen or nil if none
func (p *Profiles) ZoneAffinityProfileGetAllZoneAffinityProfilesGet(ctx context.Context) (*http.Response, *ZoneAffinityProfileGetAllZoneAffinityProfilesGet200Response, error) {
	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/zoneAffinity")
	request.authenticated = true

	out := &ZoneAffinityProfileGetAllZoneAffinityProfilesGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAffinityProfileCreateZoneAffinityProfileSettingPostRequestZoneAffinityListSlice []*string

	ZoneAffinityProfileCreateZoneAffinityProfileSettingPostRequest struct {
		Description      *string                                                                             `json:"description,omitempty"`
		Name             *string                                                                             `json:"name,omitempty"`
		ZoneAffinityList ZoneAffinityProfileCreateZoneAffinityProfileSettingPostRequestZoneAffinityListSlice `json:"zoneAffinityList,omitempty"`
	}

	ZoneAffinityProfileCreateZoneAffinityProfileSettingPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ZoneAffinityProfileCreateZoneAffinityProfileSettingPost: Use this API command to create zone affinity profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ZoneAffinityProfileCreateZoneAffinityProfileSettingPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAffinityProfileCreateZoneAffinityProfileSettingPost201Response
// - error: Error seen or nil if none
func (p *Profiles) ZoneAffinityProfileCreateZoneAffinityProfileSettingPost(ctx context.Context, requestBody *ZoneAffinityProfileCreateZoneAffinityProfileSettingPostRequest) (*http.Response, *ZoneAffinityProfileCreateZoneAffinityProfileSettingPost201Response, error) {
	request := p.client.newRequest(ctx, "POST", "/v5_0/profiles/zoneAffinity")
	request.body = requestBody
	request.authenticated = true

	out := &ZoneAffinityProfileCreateZoneAffinityProfileSettingPost201Response{}
	httpResponse, _, err := p.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// ZoneAffinityProfileDeleteZoneAffinityProfileSettingDelete: Use this API command to delete zone affinity profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) ZoneAffinityProfileDeleteZoneAffinityProfileSettingDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "DELETE", "/v5_0/profiles/zoneAffinity/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}

type (
	ZoneAffinityProfileGetOneZoneAffinityProfileGet200ResponseZoneAffinityListSlice []*string

	ZoneAffinityProfileGetOneZoneAffinityProfileGet200Response struct {
		CreateDateTime   *int                                                                            `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                         `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                         `json:"creatorUsername,omitempty"`
		Description      *string                                                                         `json:"description,omitempty"`
		ID               *string                                                                         `json:"id,omitempty"`
		ModifiedDateTime *int                                                                            `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                         `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                         `json:"modifierUsername,omitempty"`
		Name             *string                                                                         `json:"name,omitempty"`
		ZoneAffinityList ZoneAffinityProfileGetOneZoneAffinityProfileGet200ResponseZoneAffinityListSlice `json:"zoneAffinityList,omitempty"`
	}
)

// ZoneAffinityProfileGetOneZoneAffinityProfileGet: Use this API command to get one zone affinity profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAffinityProfileGetOneZoneAffinityProfileGet200Response
// - error: Error seen or nil if none
func (p *Profiles) ZoneAffinityProfileGetOneZoneAffinityProfileGet(ctx context.Context, id string) (*http.Response, *ZoneAffinityProfileGetOneZoneAffinityProfileGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "GET", "/v5_0/profiles/zoneAffinity/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &ZoneAffinityProfileGetOneZoneAffinityProfileGet200Response{}
	httpResponse, _, err := p.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAffinityProfileModifyZoneAffinityProfileSettingPatchRequest struct {
		Name *string `json:"name,omitempty"`
	}
)

// ZoneAffinityProfileModifyZoneAffinityProfileSettingPatch: Use this API command to modify zone affinity profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *ZoneAffinityProfileModifyZoneAffinityProfileSettingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *Profiles) ZoneAffinityProfileModifyZoneAffinityProfileSettingPatch(ctx context.Context, id string, requestBody *ZoneAffinityProfileModifyZoneAffinityProfileSettingPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := p.client.newRequest(ctx, "PATCH", "/v5_0/profiles/zoneAffinity/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return p.client.doRequest(request, 204, nil)
}
