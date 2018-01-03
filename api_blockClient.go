package api

import (
	"context"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type BlockClient struct {
	client *Client
}
type (
	BlockClientDeleteDelete1RequestIDListSlice []*string

	BlockClientDeleteDelete1Request struct {
		IDList BlockClientDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// BlockClientDeleteDelete1: Delete Block Client List
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BlockClientDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientDeleteDelete1(ctx context.Context, requestBody *BlockClientDeleteDelete1Request) (*http.Response, []byte, error) {
	request := b.client.newRequest(ctx, "DELETE", "/v5_0/blockClient")
	request.body = requestBody
	request.authenticated = true

	return b.client.doRequest(request, 204, nil)
}

type (
	BlockClientCreateBlockClientsPostRequestBlockClientListSlice []*BlockClientCreateBlockClientsPostRequestBlockClientList

	BlockClientCreateBlockClientsPostRequestBlockClientList struct {
		ApMac *string `json:"apMac,omitempty"`
		Mac   *string `json:"mac,omitempty"`
	}

	BlockClientCreateBlockClientsPostRequest struct {
		BlockClientList BlockClientCreateBlockClientsPostRequestBlockClientListSlice `json:"blockClientList,omitempty"`
		Description     *string                                                      `json:"description,omitempty"`
	}

	BlockClientCreateBlockClientsPost201ResponseSlice []*BlockClientCreateBlockClientsPost201Response

	BlockClientCreateBlockClientsPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// BlockClientCreateBlockClientsPost: Create new Block Clients by list
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BlockClientCreateBlockClientsPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - BlockClientCreateBlockClientsPost201ResponseSlice
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientCreateBlockClientsPost(ctx context.Context, requestBody *BlockClientCreateBlockClientsPostRequest) (*http.Response, BlockClientCreateBlockClientsPost201ResponseSlice, error) {
	request := b.client.newRequest(ctx, "POST", "/v5_0/blockClient")
	request.body = requestBody
	request.authenticated = true

	out := make(BlockClientCreateBlockClientsPost201ResponseSlice, 0)
	httpResponse, _, err := b.client.doRequest(request, 201, &(out))
	return httpResponse, out, err
}

type (
	BlockClientCreateBlockClientByApMacPostRequest struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		Mac              *string `json:"mac,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		ZoneID           *string `json:"zoneId,omitempty"`
	}

	BlockClientCreateBlockClientByApMacPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// BlockClientCreateBlockClientByApMacPost: Create a new Block Client by AP MAC
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *BlockClientCreateBlockClientByApMacPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BlockClientCreateBlockClientByApMacPost201Response
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientCreateBlockClientByApMacPost(ctx context.Context, apMac string, requestBody *BlockClientCreateBlockClientByApMacPostRequest) (*http.Response, *BlockClientCreateBlockClientByApMacPost201Response, error) {
	var err error

	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "POST", "/v5_0/blockClient/byApMac/{apMac}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"apMac": apMac,
	}

	out := &BlockClientCreateBlockClientByApMacPost201Response{}
	httpResponse, _, err := b.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	BlockClientRetrieveListGet200ResponseListSlice []*BlockClientRetrieveListGet200ResponseList

	BlockClientRetrieveListGet200ResponseList struct {
		Description      *string `json:"description,omitempty"`
		ID               *string `json:"id,omitempty"`
		Mac              *string `json:"mac,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		ZoneID           *string `json:"zoneId,omitempty"`
	}

	BlockClientRetrieveListGet200Response struct {
		FirstIndex *int                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                          `json:"hasMore,omitempty"`
		List       BlockClientRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                           `json:"totalCount,omitempty"`
	}
)

// BlockClientRetrieveListGet: Retrieve a list of Block Client
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BlockClientRetrieveListGet200Response
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *BlockClientRetrieveListGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "GET", "/v5_0/blockClient/byZone/{zoneId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"zoneId": zoneId,
	}

	out := &BlockClientRetrieveListGet200Response{}
	httpResponse, _, err := b.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	BlockClientRetrieveListPostRequestAttributesSlice []*string

	BlockClientRetrieveListPostRequestExtraFiltersSlice []*BlockClientRetrieveListPostRequestExtraFilters

	BlockClientRetrieveListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	BlockClientRetrieveListPostRequestExtraNotFiltersSlice []*BlockClientRetrieveListPostRequestExtraNotFilters

	BlockClientRetrieveListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	BlockClientRetrieveListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	BlockClientRetrieveListPostRequestFiltersSlice []*BlockClientRetrieveListPostRequestFilters

	BlockClientRetrieveListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	BlockClientRetrieveListPostRequestFullTextSearchFieldsSlice []*string

	BlockClientRetrieveListPostRequestFullTextSearch struct {
		Fields BlockClientRetrieveListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                     `json:"type,omitempty"`
		Value  *string                                                     `json:"value,omitempty"`
	}

	BlockClientRetrieveListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	BlockClientRetrieveListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	BlockClientRetrieveListPostRequestOptions struct {
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
		GuestPassExpiration           *BlockClientRetrieveListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *BlockClientRetrieveListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	BlockClientRetrieveListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	BlockClientRetrieveListPostRequest struct {
		Attributes      BlockClientRetrieveListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    BlockClientRetrieveListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters BlockClientRetrieveListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *BlockClientRetrieveListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         BlockClientRetrieveListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *BlockClientRetrieveListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                   `json:"limit,omitempty"`
		Options         *BlockClientRetrieveListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                   `json:"page,omitempty"`
		Query           *string                                                `json:"query,omitempty"`
		SortInfo        *BlockClientRetrieveListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                   `json:"start,omitempty"`
	}

	BlockClientRetrieveListPost200ResponseListSlice []*BlockClientRetrieveListPost200ResponseList

	BlockClientRetrieveListPost200ResponseList struct {
		Description      *string `json:"description,omitempty"`
		ID               *string `json:"id,omitempty"`
		Mac              *string `json:"mac,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		ZoneID           *string `json:"zoneId,omitempty"`
	}

	BlockClientRetrieveListPost200Response struct {
		FirstIndex *int                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                           `json:"hasMore,omitempty"`
		List       BlockClientRetrieveListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                            `json:"totalCount,omitempty"`
	}
)

// BlockClientRetrieveListPost: Retrieve a list of Block Client
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BlockClientRetrieveListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BlockClientRetrieveListPost200Response
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientRetrieveListPost(ctx context.Context, requestBody *BlockClientRetrieveListPostRequest) (*http.Response, *BlockClientRetrieveListPost200Response, error) {
	request := b.client.newRequest(ctx, "POST", "/v5_0/blockClient/query")
	request.body = requestBody
	request.authenticated = true

	out := &BlockClientRetrieveListPost200Response{}
	httpResponse, _, err := b.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// BlockClientDeleteDelete: Delete a Block Client
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "DELETE", "/v5_0/blockClient/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return b.client.doRequest(request, 204, nil)
}

type (
	BlockClientRetrieveGet200Response struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		Mac              *string `json:"mac,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		ZoneID           *string `json:"zoneId,omitempty"`
	}
)

// BlockClientRetrieveGet: Retrieve a Block Client
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BlockClientRetrieveGet200Response
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientRetrieveGet(ctx context.Context, id string) (*http.Response, *BlockClientRetrieveGet200Response, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "GET", "/v5_0/blockClient/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	out := &BlockClientRetrieveGet200Response{}
	httpResponse, _, err := b.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	BlockClientModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}
)

// BlockClientModifyBasicPatch: Modify a specific Block Client basic
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *BlockClientModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientModifyBasicPatch(ctx context.Context, id string, requestBody *BlockClientModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "PATCH", "/v5_0/blockClient/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return b.client.doRequest(request, 204, nil)
}

type (
	BlockClientModifyBasicPutRequest struct {
		Description *string `json:"description,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}
)

// BlockClientModifyBasicPut: Modify a specific Block Client basic
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
// - requestBody: *BlockClientModifyBasicPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientModifyBasicPut(ctx context.Context, id string, requestBody *BlockClientModifyBasicPutRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "PUT", "/v5_0/blockClient/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}

	return b.client.doRequest(request, 204, nil)
}

type (
	BlockClientCreateBlockClientPostRequest struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		Mac              *string `json:"mac,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		ZoneID           *string `json:"zoneId,omitempty"`
	}

	BlockClientCreateBlockClientPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// BlockClientCreateBlockClientPost: Create a new Block Client
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *BlockClientCreateBlockClientPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BlockClientCreateBlockClientPost201Response
// - error: Error seen or nil if none
func (b *BlockClient) BlockClientCreateBlockClientPost(ctx context.Context, zoneId string, requestBody *BlockClientCreateBlockClientPostRequest) (*http.Response, *BlockClientCreateBlockClientPost201Response, error) {
	var err error

	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}

	request := b.client.newRequest(ctx, "POST", "/v5_0/blockClient/{zoneId}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"zoneId": zoneId,
	}

	out := &BlockClientCreateBlockClientPost201Response{}
	httpResponse, _, err := b.client.doRequest(request, 201, out)
	return httpResponse, out, err
}
