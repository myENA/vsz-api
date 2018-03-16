package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type RuckusZonesAPI struct {
	client *Client
}
type (
	RuckusWirelessApZoneRetrieveListGet200ResponseListSlice []*RuckusWirelessApZoneRetrieveListGet200ResponseList

	RuckusWirelessApZoneRetrieveListGet200ResponseList struct {
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		ServiceName *string `json:"serviceName,omitempty"`
	}

	RuckusWirelessApZoneRetrieveListGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       RuckusWirelessApZoneRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveListGet: Use this API command to retrieve the list of Ruckus Wireless AP zones that belong to a domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - domainId (UUIDv4): The domain ID.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *RuckusWirelessApZoneRetrieveListGet200Response, error) {
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
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/rkszones", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("domainId", domainId)
	out := &RuckusWirelessApZoneRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneCreateZonePostRequestLogin struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}

	RuckusWirelessApZoneCreateZonePostRequest struct {
		CountryCode *string                                         `json:"countryCode,omitempty"`
		Description *string                                         `json:"description,omitempty"`
		DomainID    *string                                         `json:"domainId,omitempty"`
		Login       *RuckusWirelessApZoneCreateZonePostRequestLogin `json:"login,omitempty"`
		Name        *string                                         `json:"name,omitempty"`
		Version     *string                                         `json:"version,omitempty"`
	}

	RuckusWirelessApZoneCreateZonePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// RuckusWirelessApZoneCreateZonePost: Use this API command to create a new Ruckus Wireless AP zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusWirelessApZoneCreateZonePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneCreateZonePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneCreateZonePost(ctx context.Context, requestBody *RuckusWirelessApZoneCreateZonePostRequest) (*http.Response, *RuckusWirelessApZoneCreateZonePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/rkszones", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &RuckusWirelessApZoneCreateZonePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	BonjourFencingPolicyDeleteDelete1RequestIDListSlice []*string

	BonjourFencingPolicyDeleteDelete1Request struct {
		IDList BonjourFencingPolicyDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// BonjourFencingPolicyDeleteDelete1: Use this API command to delete Bulk Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BonjourFencingPolicyDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyDeleteDelete1(ctx context.Context, requestBody *BonjourFencingPolicyDeleteDelete1Request) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("DELETE", "/v5_0/rkszones/bonjourFencingPolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return r.client.Ensure(ctx, request, 204, nil)
}

// BonjourFencingPolicyDeleteDelete: Use this API command to delete Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/bonjourFencingPolicy/{id}", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ClientIsolationWhitelistDeleteDelete1RequestIDListSlice []*string

	ClientIsolationWhitelistDeleteDelete1Request struct {
		IDList ClientIsolationWhitelistDeleteDelete1RequestIDListSlice `json:"idList,omitempty"`
	}
)

// ClientIsolationWhitelistDeleteDelete1: Use this API command to delete Bulk Client Isolation Whitelist
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ClientIsolationWhitelistDeleteDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistDeleteDelete1(ctx context.Context, requestBody *ClientIsolationWhitelistDeleteDelete1Request) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("DELETE", "/v5_0/rkszones/clientIsolationWhitelist", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return r.client.Ensure(ctx, request, 204, nil)
}

// ClientIsolationWhitelistDeleteDelete: Delete a Client Isolation Whitelist
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/clientIsolationWhitelist/{id}", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DomainRetrieveListGet200ResponseListSlice []*DomainRetrieveListGet200ResponseList

	DomainRetrieveListGet200ResponseList struct {
		AdministratorCount *int    `json:"administratorCount,omitempty"`
		ApCount            *int    `json:"apCount,omitempty"`
		CreateDatetime     *string `json:"createDatetime,omitempty"`
		CreatedBy          *string `json:"createdBy,omitempty"`
		Description        *string `json:"description,omitempty"`
		DomainType         *string `json:"domainType,omitempty"`
		ID                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		ParentDomainID     *string `json:"parentDomainId,omitempty"`
		SubDomainCount     *int    `json:"subDomainCount,omitempty"`
		ZoneCount          *int    `json:"zoneCount,omitempty"`
	}

	DomainRetrieveListGet200ResponsePropertiesFirstIndex struct {
		Type *string `json:"type,omitempty"`
	}

	DomainRetrieveListGet200ResponsePropertiesHasMore struct {
		Type *string `json:"type,omitempty"`
	}

	DomainRetrieveListGet200ResponsePropertiesListItems struct {
		XRef *string `json:"$ref,omitempty"`
	}

	DomainRetrieveListGet200ResponsePropertiesList struct {
		Items *DomainRetrieveListGet200ResponsePropertiesListItems `json:"items,omitempty"`
		Type  *string                                              `json:"type,omitempty"`
	}

	DomainRetrieveListGet200ResponsePropertiesTotalCount struct {
		Type *string `json:"type,omitempty"`
	}

	DomainRetrieveListGet200ResponseProperties struct {
		FirstIndex *DomainRetrieveListGet200ResponsePropertiesFirstIndex `json:"firstIndex,omitempty"`
		HasMore    *DomainRetrieveListGet200ResponsePropertiesHasMore    `json:"hasMore,omitempty"`
		List       *DomainRetrieveListGet200ResponsePropertiesList       `json:"list,omitempty"`
		TotalCount *DomainRetrieveListGet200ResponsePropertiesTotalCount `json:"totalCount,omitempty"`
	}

	DomainRetrieveListGet200Response struct {
		AdditionalProperties *bool                                       `json:"additionalProperties,omitempty"`
		FirstIndex           *int                                        `json:"firstIndex,omitempty"`
		HasMore              *bool                                       `json:"hasMore,omitempty"`
		List                 DomainRetrieveListGet200ResponseListSlice   `json:"list,omitempty"`
		Properties           *DomainRetrieveListGet200ResponseProperties `json:"properties,omitempty"`
		TotalCount           *int                                        `json:"totalCount,omitempty"`
		Type                 *string                                     `json:"type,omitempty"`
	}
)

// DomainRetrieveListGet: Use this API command to retrieve a list of domain under Administration Domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - recursively (bool): Get domain list recursively.
// - includeSelf (bool): Get domain list include Self.
// - excludeRegularDomain (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DomainRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DomainRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *DomainRetrieveListGet200Response, error) {
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
	recursively, ok := optionalParams["recursively"]
	if ok {
		err = validators.StrIsBool(recursively)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"recursively\" failed validation check: %s", err)
		}
	}
	includeSelf, ok := optionalParams["includeSelf"]
	if ok {
		err = validators.StrIsBool(includeSelf)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"includeSelf\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/rkszones/domains", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("recursively", recursively)
	request.SetQueryParameter("includeSelf", includeSelf)
	request.SetQueryParameter("excludeRegularDomain", optionalParams["excludeRegularDomain"])
	out := &DomainRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DomainCreatePostRequestPropertiesDescription struct {
		XRef        *string `json:"$ref,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	DomainCreatePostRequestPropertiesName struct {
		XRef        *string `json:"$ref,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	DomainCreatePostRequestProperties struct {
		Description *DomainCreatePostRequestPropertiesDescription `json:"description,omitempty"`
		Name        *DomainCreatePostRequestPropertiesName        `json:"name,omitempty"`
	}

	DomainCreatePostRequestRequiredSlice []*string

	DomainCreatePostRequest struct {
		AdditionalProperties *bool                                `json:"additionalProperties,omitempty"`
		Description          *string                              `json:"description,omitempty"`
		DomainType           *string                              `json:"domainType,omitempty"`
		Name                 *string                              `json:"name,omitempty"`
		ParentDomainID       *string                              `json:"parentDomainId,omitempty"`
		Properties           *DomainCreatePostRequestProperties   `json:"properties,omitempty"`
		Required             DomainCreatePostRequestRequiredSlice `json:"required,omitempty"`
		Type                 *string                              `json:"type,omitempty"`
	}

	DomainCreatePost201ResponsePropertiesID struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainCreatePost201ResponseProperties struct {
		ID *DomainCreatePost201ResponsePropertiesID `json:"id,omitempty"`
	}

	DomainCreatePost201Response struct {
		AdditionalProperties *bool                                  `json:"additionalProperties,omitempty"`
		ID                   *string                                `json:"id,omitempty"`
		Properties           *DomainCreatePost201ResponseProperties `json:"properties,omitempty"`
		Type                 *string                                `json:"type,omitempty"`
	}
)

// DomainCreatePost: Use this API command to create new domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *DomainCreatePostRequest
//
// Optional Parameter Map:
// - parentDomainId (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DomainCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DomainCreatePost(ctx context.Context, requestBody *DomainCreatePostRequest, optionalParams map[string]string) (*http.Response, *DomainCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	parentDomainId, ok := optionalParams["parentDomainId"]
	if ok {
		err = validators.StrIsUUIDv4(parentDomainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"parentDomainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("POST", "/v5_0/rkszones/domains", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetQueryParameter("parentDomainId", parentDomainId)
	out := &DomainCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// DomainDeleteDelete: Use this API command to delete domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Domain ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DomainDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/domains/{id}", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DomainRetrieveGet200ResponsePropertiesAdministratorCount struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesApCount struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesCreateDatetime struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesCreatedBy struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesDescription struct {
		XRef        *string `json:"$ref,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesID struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesName struct {
		XRef        *string `json:"$ref,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesParentDomainID struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesSubDomainCount struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponsePropertiesZoneCount struct {
		Description *string `json:"description,omitempty"`
		Type        *string `json:"type,omitempty"`
	}

	DomainRetrieveGet200ResponseProperties struct {
		AdministratorCount *DomainRetrieveGet200ResponsePropertiesAdministratorCount `json:"administratorCount,omitempty"`
		ApCount            *DomainRetrieveGet200ResponsePropertiesApCount            `json:"apCount,omitempty"`
		CreateDatetime     *DomainRetrieveGet200ResponsePropertiesCreateDatetime     `json:"createDatetime,omitempty"`
		CreatedBy          *DomainRetrieveGet200ResponsePropertiesCreatedBy          `json:"createdBy,omitempty"`
		Description        *DomainRetrieveGet200ResponsePropertiesDescription        `json:"description,omitempty"`
		ID                 *DomainRetrieveGet200ResponsePropertiesID                 `json:"id,omitempty"`
		Name               *DomainRetrieveGet200ResponsePropertiesName               `json:"name,omitempty"`
		ParentDomainID     *DomainRetrieveGet200ResponsePropertiesParentDomainID     `json:"parentDomainId,omitempty"`
		SubDomainCount     *DomainRetrieveGet200ResponsePropertiesSubDomainCount     `json:"subDomainCount,omitempty"`
		ZoneCount          *DomainRetrieveGet200ResponsePropertiesZoneCount          `json:"zoneCount,omitempty"`
	}

	DomainRetrieveGet200Response struct {
		AdditionalProperties *bool                                   `json:"additionalProperties,omitempty"`
		AdministratorCount   *int                                    `json:"administratorCount,omitempty"`
		ApCount              *int                                    `json:"apCount,omitempty"`
		CreateDatetime       *string                                 `json:"createDatetime,omitempty"`
		CreatedBy            *string                                 `json:"createdBy,omitempty"`
		Description          *string                                 `json:"description,omitempty"`
		DomainType           *string                                 `json:"domainType,omitempty"`
		ID                   *string                                 `json:"id,omitempty"`
		Name                 *string                                 `json:"name,omitempty"`
		ParentDomainID       *string                                 `json:"parentDomainId,omitempty"`
		Properties           *DomainRetrieveGet200ResponseProperties `json:"properties,omitempty"`
		SubDomainCount       *int                                    `json:"subDomainCount,omitempty"`
		Type                 *string                                 `json:"type,omitempty"`
		ZoneCount            *int                                    `json:"zoneCount,omitempty"`
	}
)

// DomainRetrieveGet: Use this API command to retrieve domain by specified Domain ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Domain ID
//
// Optional Parameter Map:
// - recursively (bool)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DomainRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DomainRetrieveGet(ctx context.Context, id string, optionalParams map[string]string) (*http.Response, *DomainRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	recursively, ok := optionalParams["recursively"]
	if ok {
		err = validators.StrIsBool(recursively)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"recursively\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/rkszones/domains/{id}", true)
	request.SetPathParameter("id", id)
	request.SetQueryParameter("recursively", recursively)
	out := &DomainRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DomainModifyBasicPatchRequest struct {
		Description    *string `json:"description,omitempty"`
		DomainType     *string `json:"domainType,omitempty"`
		Name           *string `json:"name,omitempty"`
		ParentDomainID *string `json:"parentDomainId,omitempty"`
	}
)

// DomainModifyBasicPatch: Use this API command to modify the basic information of domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Domain ID
// - requestBody: *DomainModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DomainModifyBasicPatch(ctx context.Context, id string, requestBody *DomainModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/domains/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DomainRetrieveSubdomainListGet200ResponseListSlice []*DomainRetrieveSubdomainListGet200ResponseList

	DomainRetrieveSubdomainListGet200ResponseList struct {
		AdministratorCount *int    `json:"administratorCount,omitempty"`
		ApCount            *int    `json:"apCount,omitempty"`
		CreateDatetime     *string `json:"createDatetime,omitempty"`
		CreatedBy          *string `json:"createdBy,omitempty"`
		Description        *string `json:"description,omitempty"`
		DomainType         *string `json:"domainType,omitempty"`
		ID                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		ParentDomainID     *string `json:"parentDomainId,omitempty"`
		SubDomainCount     *int    `json:"subDomainCount,omitempty"`
		ZoneCount          *int    `json:"zoneCount,omitempty"`
	}

	DomainRetrieveSubdomainListGet200ResponsePropertiesFirstIndex struct {
		Type *string `json:"type,omitempty"`
	}

	DomainRetrieveSubdomainListGet200ResponsePropertiesHasMore struct {
		Type *string `json:"type,omitempty"`
	}

	DomainRetrieveSubdomainListGet200ResponsePropertiesListItems struct {
		XRef *string `json:"$ref,omitempty"`
	}

	DomainRetrieveSubdomainListGet200ResponsePropertiesList struct {
		Items *DomainRetrieveSubdomainListGet200ResponsePropertiesListItems `json:"items,omitempty"`
		Type  *string                                                       `json:"type,omitempty"`
	}

	DomainRetrieveSubdomainListGet200ResponsePropertiesTotalCount struct {
		Type *string `json:"type,omitempty"`
	}

	DomainRetrieveSubdomainListGet200ResponseProperties struct {
		FirstIndex *DomainRetrieveSubdomainListGet200ResponsePropertiesFirstIndex `json:"firstIndex,omitempty"`
		HasMore    *DomainRetrieveSubdomainListGet200ResponsePropertiesHasMore    `json:"hasMore,omitempty"`
		List       *DomainRetrieveSubdomainListGet200ResponsePropertiesList       `json:"list,omitempty"`
		TotalCount *DomainRetrieveSubdomainListGet200ResponsePropertiesTotalCount `json:"totalCount,omitempty"`
	}

	DomainRetrieveSubdomainListGet200Response struct {
		AdditionalProperties *bool                                                `json:"additionalProperties,omitempty"`
		FirstIndex           *int                                                 `json:"firstIndex,omitempty"`
		HasMore              *bool                                                `json:"hasMore,omitempty"`
		List                 DomainRetrieveSubdomainListGet200ResponseListSlice   `json:"list,omitempty"`
		Properties           *DomainRetrieveSubdomainListGet200ResponseProperties `json:"properties,omitempty"`
		TotalCount           *int                                                 `json:"totalCount,omitempty"`
		Type                 *string                                              `json:"type,omitempty"`
	}
)

// DomainRetrieveSubdomainListGet: Use this API command to retrieve a list of subdomain by specified Domain ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Domain ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - recursively (bool): Get domain list recursively.
// - includeSelf (bool): Get domain list include Self.
// - excludeRegularDomain (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DomainRetrieveSubdomainListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DomainRetrieveSubdomainListGet(ctx context.Context, id string, optionalParams map[string]string) (*http.Response, *DomainRetrieveSubdomainListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
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
	recursively, ok := optionalParams["recursively"]
	if ok {
		err = validators.StrIsBool(recursively)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"recursively\" failed validation check: %s", err)
		}
	}
	includeSelf, ok := optionalParams["includeSelf"]
	if ok {
		err = validators.StrIsBool(includeSelf)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"includeSelf\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/rkszones/domains/{id}/subdomain", true)
	request.SetPathParameter("id", id)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("recursively", recursively)
	request.SetQueryParameter("includeSelf", includeSelf)
	request.SetQueryParameter("excludeRegularDomain", optionalParams["excludeRegularDomain"])
	out := &DomainRetrieveSubdomainListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// DynamicPskDownloadDpskCsvSampleGet: Use this API command to download DPSK CSV sample.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - type (string): DPSK CSV sample type. Valid value is PHASE1 or PHASE2.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskDownloadDpskCsvSampleGet(ctx context.Context, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	xtype, ok := optionalParams["type"]
	if !ok {
		xtype = "PHASE2"
	}
	request := NewRequest("GET", "/v5_0/rkszones/downloadDpskCsvSample", true)
	request.SetQueryParameter("type", xtype)
	return r.client.Ensure(ctx, request, 200, nil)
}

type (
	RuckusWirelessApZoneCreateZoneOfDualPostRequestLogin struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}

	RuckusWirelessApZoneCreateZoneOfDualPostRequest struct {
		CountryCode *string                                               `json:"countryCode,omitempty"`
		Description *string                                               `json:"description,omitempty"`
		DomainID    *string                                               `json:"domainId,omitempty"`
		Login       *RuckusWirelessApZoneCreateZoneOfDualPostRequestLogin `json:"login,omitempty"`
		Name        *string                                               `json:"name,omitempty"`
		Version     *string                                               `json:"version,omitempty"`
	}

	RuckusWirelessApZoneCreateZoneOfDualPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// RuckusWirelessApZoneCreateZoneOfDualPost: Use this API command to create a new Ruckus Wireless AP zone of IPv4/IPv6.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusWirelessApZoneCreateZoneOfDualPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneCreateZoneOfDualPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneCreateZoneOfDualPost(ctx context.Context, requestBody *RuckusWirelessApZoneCreateZoneOfDualPostRequest) (*http.Response, *RuckusWirelessApZoneCreateZoneOfDualPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/rkszones/dual", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &RuckusWirelessApZoneCreateZoneOfDualPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneCreateZoneOfIpv6PostRequestLogin struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}

	RuckusWirelessApZoneCreateZoneOfIpv6PostRequest struct {
		CountryCode *string                                               `json:"countryCode,omitempty"`
		Description *string                                               `json:"description,omitempty"`
		DomainID    *string                                               `json:"domainId,omitempty"`
		Login       *RuckusWirelessApZoneCreateZoneOfIpv6PostRequestLogin `json:"login,omitempty"`
		Name        *string                                               `json:"name,omitempty"`
		Version     *string                                               `json:"version,omitempty"`
	}

	RuckusWirelessApZoneCreateZoneOfIpv6Post201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// RuckusWirelessApZoneCreateZoneOfIpv6Post: Use this API command to create a new Ruckus Wireless AP zone of IPv6.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusWirelessApZoneCreateZoneOfIpv6PostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneCreateZoneOfIpv6Post201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneCreateZoneOfIpv6Post(ctx context.Context, requestBody *RuckusWirelessApZoneCreateZoneOfIpv6PostRequest) (*http.Response, *RuckusWirelessApZoneCreateZoneOfIpv6Post201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/rkszones/ipv6", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &RuckusWirelessApZoneCreateZoneOfIpv6Post201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestAttributesSlice []*string

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraFiltersSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraFilters

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraNotFiltersSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraNotFilters

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFiltersSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFilters

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFullTextSearchFieldsSlice []*string

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFullTextSearch struct {
		Fields RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                 `json:"type,omitempty"`
		Value  *string                                                                                                 `json:"value,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                     `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                   `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                     `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                     `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                   `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                     `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                     `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                     `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                     `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                     `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                     `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                   `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                     `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                   `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                   `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                   `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                   `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                   `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                     `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                     `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                     `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                   `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                   `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                   `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                   `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                   `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                   `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                   `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                   `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                   `json:"localUser_userSource,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequest struct {
		Attributes      RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                            `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                              `json:"expandDomains,omitempty"`
		ExtraFilters    RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                                                           `json:"limit,omitempty"`
		Options         *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                               `json:"page,omitempty"`
		Query           *string                                                                                            `json:"query,omitempty"`
		SortInfo        *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                                                           `json:"start,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseList

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteApsSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteAps

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteAps struct {
		ApGatewayIP     *string `json:"apGatewayIp,omitempty"`
		ApMac           *string `json:"apMac,omitempty"`
		ApName          *string `json:"apName,omitempty"`
		ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
		ApServerIP      *string `json:"apServerIp,omitempty"`
		ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
		ApStatus        *string `json:"apStatus,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteProfilesSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteProfiles

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteProfiles struct {
		Description      *string  `json:"description,omitempty"`
		ID               *string  `json:"id,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PoolEndIP        *string  `json:"poolEndIp,omitempty"`
		PoolStartIP      *string  `json:"poolStartIp,omitempty"`
		PrimaryDNSIP     *string  `json:"primaryDnsIp,omitempty"`
		SecondaryDNSIP   *string  `json:"secondaryDnsIp,omitempty"`
		SubnetMask       *string  `json:"subnetMask,omitempty"`
		SubnetNetworkIP  *string  `json:"subnetNetworkIp,omitempty"`
		VlanID           *float64 `json:"vlanId,omitempty"`
		ZoneID           *string  `json:"zoneId,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseList struct {
		ManualSelect *bool                                                                                                   `json:"manualSelect,omitempty"`
		SiteAps      RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteApsSlice      `json:"siteAps,omitempty"`
		SiteEnabled  *bool                                                                                                   `json:"siteEnabled,omitempty"`
		SiteMode     *string                                                                                                 `json:"siteMode,omitempty"`
		SiteProfiles RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSiteProfilesSlice `json:"siteProfiles,omitempty"`
		ZoneName     *string                                                                                                 `json:"zoneName,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200Response struct {
		FirstIndex *int                                                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                       `json:"hasMore,omitempty"`
		List       RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                        `json:"totalCount,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost: Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost(ctx context.Context, requestBody *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPostRequest) (*http.Response, *RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/rkszones/services/dhcpSiteConfig/query", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &RuckusWirelessApZoneRetrieveDhcpNatServiceConfigurationWithinDomainPost200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// RuckusWirelessApZoneDeleteDelete: Use this API command to delete a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneRetrieveGet200ResponseAltitude struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseApMgmtVlan struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseApRebootTimeout struct {
		GatewayLossTimeoutInSec *float64 `json:"gatewayLossTimeoutInSec,omitempty"`
		ServerLossTimeoutInSec  *float64 `json:"serverLossTimeoutInSec,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseAutoChannelSelection24 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseAutoChannelSelection50 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseBackgroundScanning24 struct {
		FrequencyInSec *float64 `json:"frequencyInSec,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseBackgroundScanning50 struct {
		FrequencyInSec *float64 `json:"frequencyInSec,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseBandBalancing struct {
		Wifi24Percentage *float64 `json:"wifi24Percentage,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseBonjourFencingPolicy struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseClientAdmissionControl24 struct {
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseClientAdmissionControl50 struct {
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseClientLoadBalancing24 struct {
		AdjacentRadioThreshold *float64 `json:"adjacentRadioThreshold,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseClientLoadBalancing50 struct {
		AdjacentRadioThreshold *float64 `json:"adjacentRadioThreshold,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfigSiteApsSlice []*RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfigSiteAps

	RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfigSiteAps struct {
		ApGatewayIP     *string `json:"apGatewayIp,omitempty"`
		ApMac           *string `json:"apMac,omitempty"`
		ApName          *string `json:"apName,omitempty"`
		ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
		ApServerIP      *string `json:"apServerIp,omitempty"`
		ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
		ApStatus        *string `json:"apStatus,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfigSiteProfileIdsSlice []*string

	RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfig struct {
		ManualSelect   *bool                                                                       `json:"manualSelect,omitempty"`
		SiteAps        RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfigSiteApsSlice        `json:"siteAps,omitempty"`
		SiteEnabled    *bool                                                                       `json:"siteEnabled,omitempty"`
		SiteMode       *string                                                                     `json:"siteMode,omitempty"`
		SiteProfileIds RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfigSiteProfileIdsSlice `json:"siteProfileIds,omitempty"`
		ZoneName       *string                                                                     `json:"zoneName,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseIpsecProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseLocationBasedService struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseLogin struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseMesh struct {
		Passphrase *string `json:"passphrase,omitempty"`
		Ssid       *string `json:"ssid,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseNodeAffinityProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseRecoverySsid struct {
		RecoverySsidEnable *float64 `json:"recoverySsidEnable,omitempty"`
		RecoverySsidPskKey *string  `json:"recoverySsidPskKey,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseRogueMaliciousTypesSlice []*string

	RuckusWirelessApZoneRetrieveGet200ResponseRogue struct {
		MaliciousTypes    RuckusWirelessApZoneRetrieveGet200ResponseRogueMaliciousTypesSlice `json:"maliciousTypes,omitempty"`
		ProtectionEnabled *bool                                                              `json:"protectionEnabled,omitempty"`
		ReportType        *string                                                            `json:"reportType,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSmartMonitor struct {
		IntervalInSec  *int `json:"intervalInSec,omitempty"`
		RetryThreshold *int `json:"retryThreshold,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2AgentSlice []*RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2Agent

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2AgentNotificationTargetSlice []*RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2AgentNotificationTarget

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2Agent struct {
		CommunityName       *string                                                                               `json:"communityName,omitempty"`
		NotificationEnabled *bool                                                                                 `json:"notificationEnabled,omitempty"`
		NotificationTarget  RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                                               `json:"notificationType,omitempty"`
		ReadEnabled         *bool                                                                                 `json:"readEnabled,omitempty"`
		WriteEnabled        *bool                                                                                 `json:"writeEnabled,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3AgentSlice []*RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3Agent

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3AgentNotificationTargetSlice []*RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3AgentNotificationTarget

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3Agent struct {
		AuthPassword        *string                                                                               `json:"authPassword,omitempty"`
		AuthProtocol        *string                                                                               `json:"authProtocol,omitempty"`
		NotificationEnabled *bool                                                                                 `json:"notificationEnabled,omitempty"`
		NotificationTarget  RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                                               `json:"notificationType,omitempty"`
		PrivPassword        *string                                                                               `json:"privPassword,omitempty"`
		PrivProtocol        *string                                                                               `json:"privProtocol,omitempty"`
		ReadEnabled         *bool                                                                                 `json:"readEnabled,omitempty"`
		UserName            *string                                                                               `json:"userName,omitempty"`
		WriteEnabled        *bool                                                                                 `json:"writeEnabled,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgent struct {
		ApSnmpEnabled *bool                                                               `json:"apSnmpEnabled,omitempty"`
		SnmpV2Agent   RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV2AgentSlice `json:"snmpV2Agent,omitempty"`
		SnmpV3Agent   RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgentSnmpV3AgentSlice `json:"snmpV3Agent,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseSyslog struct {
		Address  *string `json:"address,omitempty"`
		Facility *string `json:"facility,omitempty"`
		Port     *int    `json:"port,omitempty"`
		Priority *string `json:"priority,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseTimezoneCustomizedTimezoneEnd struct {
		Day   *float64 `json:"day,omitempty"`
		Hour  *float64 `json:"hour,omitempty"`
		Month *float64 `json:"month,omitempty"`
		Week  *float64 `json:"week,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseTimezoneCustomizedTimezoneStart struct {
		Day   *float64 `json:"day,omitempty"`
		Hour  *float64 `json:"hour,omitempty"`
		Month *float64 `json:"month,omitempty"`
		Week  *float64 `json:"week,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseTimezoneCustomizedTimezone struct {
		Abbreviation    *string                                                                    `json:"abbreviation,omitempty"`
		End             *RuckusWirelessApZoneRetrieveGet200ResponseTimezoneCustomizedTimezoneEnd   `json:"end,omitempty"`
		GmtOffset       *float64                                                                   `json:"gmtOffset,omitempty"`
		GmtOffsetMinute *float64                                                                   `json:"gmtOffsetMinute,omitempty"`
		Start           *RuckusWirelessApZoneRetrieveGet200ResponseTimezoneCustomizedTimezoneStart `json:"start,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseTimezone struct {
		CustomizedTimezone *RuckusWirelessApZoneRetrieveGet200ResponseTimezoneCustomizedTimezone `json:"customizedTimezone,omitempty"`
		SystemTimezone     *string                                                               `json:"systemTimezone,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseTunnelProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseUsbSoftwarePackageApplyModelSlice []*string

	RuckusWirelessApZoneRetrieveGet200ResponseUsbSoftwarePackageUsbSoftware struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseUsbSoftwarePackage struct {
		ApplyModel  RuckusWirelessApZoneRetrieveGet200ResponseUsbSoftwarePackageApplyModelSlice `json:"applyModel,omitempty"`
		UsbSoftware *RuckusWirelessApZoneRetrieveGet200ResponseUsbSoftwarePackageUsbSoftware    `json:"usbSoftware,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseVenueProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseWifi24AvailableChannelRangeSlice []*float64

	RuckusWirelessApZoneRetrieveGet200ResponseWifi24ChannelRangeSlice []*float64

	RuckusWirelessApZoneRetrieveGet200ResponseWifi24 struct {
		AvailableChannelRange RuckusWirelessApZoneRetrieveGet200ResponseWifi24AvailableChannelRangeSlice `json:"availableChannelRange,omitempty"`
		Channel               *int                                                                       `json:"channel,omitempty"`
		ChannelRange          RuckusWirelessApZoneRetrieveGet200ResponseWifi24ChannelRangeSlice          `json:"channelRange,omitempty"`
		ChannelWidth          *int                                                                       `json:"channelWidth,omitempty"`
		TxPower               *string                                                                    `json:"txPower,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200ResponseWifi50AvailableIndoorChannelRangeSlice []*float64

	RuckusWirelessApZoneRetrieveGet200ResponseWifi50AvailableOutdoorChannelRangeSlice []*float64

	RuckusWirelessApZoneRetrieveGet200ResponseWifi50IndoorChannelRangeSlice []*float64

	RuckusWirelessApZoneRetrieveGet200ResponseWifi50OutdoorChannelRangeSlice []*float64

	RuckusWirelessApZoneRetrieveGet200ResponseWifi50 struct {
		AvailableIndoorChannelRange  RuckusWirelessApZoneRetrieveGet200ResponseWifi50AvailableIndoorChannelRangeSlice  `json:"availableIndoorChannelRange,omitempty"`
		AvailableOutdoorChannelRange RuckusWirelessApZoneRetrieveGet200ResponseWifi50AvailableOutdoorChannelRangeSlice `json:"availableOutdoorChannelRange,omitempty"`
		ChannelWidth                 *int                                                                              `json:"channelWidth,omitempty"`
		IndoorChannel                *float64                                                                          `json:"indoorChannel,omitempty"`
		IndoorChannelRange           RuckusWirelessApZoneRetrieveGet200ResponseWifi50IndoorChannelRangeSlice           `json:"indoorChannelRange,omitempty"`
		IndoorSecondaryChannel       *float64                                                                          `json:"indoorSecondaryChannel,omitempty"`
		OutdoorChannel               *float64                                                                          `json:"outdoorChannel,omitempty"`
		OutdoorChannelRange          RuckusWirelessApZoneRetrieveGet200ResponseWifi50OutdoorChannelRangeSlice          `json:"outdoorChannelRange,omitempty"`
		OutdoorSecondaryChannel      *float64                                                                          `json:"outdoorSecondaryChannel,omitempty"`
		TxPower                      *string                                                                           `json:"txPower,omitempty"`
	}

	RuckusWirelessApZoneRetrieveGet200Response struct {
		Altitude                    *RuckusWirelessApZoneRetrieveGet200ResponseAltitude                 `json:"altitude,omitempty"`
		ApMgmtVlan                  *RuckusWirelessApZoneRetrieveGet200ResponseApMgmtVlan               `json:"apMgmtVlan,omitempty"`
		ApRebootTimeout             *RuckusWirelessApZoneRetrieveGet200ResponseApRebootTimeout          `json:"apRebootTimeout,omitempty"`
		AutoChannelSelection24      *RuckusWirelessApZoneRetrieveGet200ResponseAutoChannelSelection24   `json:"autoChannelSelection24,omitempty"`
		AutoChannelSelection50      *RuckusWirelessApZoneRetrieveGet200ResponseAutoChannelSelection50   `json:"autoChannelSelection50,omitempty"`
		AwsVenue                    *string                                                             `json:"awsVenue,omitempty"`
		BackgroundScanning24        *RuckusWirelessApZoneRetrieveGet200ResponseBackgroundScanning24     `json:"backgroundScanning24,omitempty"`
		BackgroundScanning50        *RuckusWirelessApZoneRetrieveGet200ResponseBackgroundScanning50     `json:"backgroundScanning50,omitempty"`
		BandBalancing               *RuckusWirelessApZoneRetrieveGet200ResponseBandBalancing            `json:"bandBalancing,omitempty"`
		BonjourFencingPolicy        *RuckusWirelessApZoneRetrieveGet200ResponseBonjourFencingPolicy     `json:"bonjourFencingPolicy,omitempty"`
		BonjourFencingPolicyEnabled *bool                                                               `json:"bonjourFencingPolicyEnabled,omitempty"`
		ChannelEvaluationInterval   *float64                                                            `json:"channelEvaluationInterval,omitempty"`
		ChannelModeEnabled          *bool                                                               `json:"channelModeEnabled,omitempty"`
		ClientAdmissionControl24    *RuckusWirelessApZoneRetrieveGet200ResponseClientAdmissionControl24 `json:"clientAdmissionControl24,omitempty"`
		ClientAdmissionControl50    *RuckusWirelessApZoneRetrieveGet200ResponseClientAdmissionControl50 `json:"clientAdmissionControl50,omitempty"`
		ClientLoadBalancing24       *RuckusWirelessApZoneRetrieveGet200ResponseClientLoadBalancing24    `json:"clientLoadBalancing24,omitempty"`
		ClientLoadBalancing50       *RuckusWirelessApZoneRetrieveGet200ResponseClientLoadBalancing50    `json:"clientLoadBalancing50,omitempty"`
		CountryCode                 *string                                                             `json:"countryCode,omitempty"`
		Description                 *string                                                             `json:"description,omitempty"`
		DfsChannelEnabled           *bool                                                               `json:"dfsChannelEnabled,omitempty"`
		DhcpSiteConfig              *RuckusWirelessApZoneRetrieveGet200ResponseDhcpSiteConfig           `json:"dhcpSiteConfig,omitempty"`
		DomainID                    *string                                                             `json:"domainId,omitempty"`
		DosBarringCheckPeriod       *float64                                                            `json:"dosBarringCheckPeriod,omitempty"`
		DosBarringEnable            *float64                                                            `json:"dosBarringEnable,omitempty"`
		DosBarringPeriod            *float64                                                            `json:"dosBarringPeriod,omitempty"`
		DosBarringThreshold         *float64                                                            `json:"dosBarringThreshold,omitempty"`
		ID                          *string                                                             `json:"id,omitempty"`
		IPMode                      *string                                                             `json:"ipMode,omitempty"`
		IpsecProfile                *RuckusWirelessApZoneRetrieveGet200ResponseIpsecProfile             `json:"ipsecProfile,omitempty"`
		Ipv6TrafficFilterEnabled    *float64                                                            `json:"ipv6TrafficFilterEnabled,omitempty"`
		Latitude                    *float64                                                            `json:"latitude,omitempty"`
		Location                    *string                                                             `json:"location,omitempty"`
		LocationAdditionalInfo      *string                                                             `json:"locationAdditionalInfo,omitempty"`
		LocationBasedService        *RuckusWirelessApZoneRetrieveGet200ResponseLocationBasedService     `json:"locationBasedService,omitempty"`
		Login                       *RuckusWirelessApZoneRetrieveGet200ResponseLogin                    `json:"login,omitempty"`
		Longitude                   *float64                                                            `json:"longitude,omitempty"`
		Mesh                        *RuckusWirelessApZoneRetrieveGet200ResponseMesh                     `json:"mesh,omitempty"`
		Name                        *string                                                             `json:"name,omitempty"`
		NodeAffinityProfile         *RuckusWirelessApZoneRetrieveGet200ResponseNodeAffinityProfile      `json:"nodeAffinityProfile,omitempty"`
		RecoverySsid                *RuckusWirelessApZoneRetrieveGet200ResponseRecoverySsid             `json:"recoverySsid,omitempty"`
		Rogue                       *RuckusWirelessApZoneRetrieveGet200ResponseRogue                    `json:"rogue,omitempty"`
		SmartMonitor                *RuckusWirelessApZoneRetrieveGet200ResponseSmartMonitor             `json:"smartMonitor,omitempty"`
		SnmpAgent                   *RuckusWirelessApZoneRetrieveGet200ResponseSnmpAgent                `json:"snmpAgent,omitempty"`
		Syslog                      *RuckusWirelessApZoneRetrieveGet200ResponseSyslog                   `json:"syslog,omitempty"`
		Timezone                    *RuckusWirelessApZoneRetrieveGet200ResponseTimezone                 `json:"timezone,omitempty"`
		TunnelProfile               *RuckusWirelessApZoneRetrieveGet200ResponseTunnelProfile            `json:"tunnelProfile,omitempty"`
		TunnelType                  *string                                                             `json:"tunnelType,omitempty"`
		UsbSoftwarePackage          *RuckusWirelessApZoneRetrieveGet200ResponseUsbSoftwarePackage       `json:"usbSoftwarePackage,omitempty"`
		VenueProfile                *RuckusWirelessApZoneRetrieveGet200ResponseVenueProfile             `json:"venueProfile,omitempty"`
		Version                     *string                                                             `json:"version,omitempty"`
		VlanOverlappingEnabled      *bool                                                               `json:"vlanOverlappingEnabled,omitempty"`
		Wifi24                      *RuckusWirelessApZoneRetrieveGet200ResponseWifi24                   `json:"wifi24,omitempty"`
		Wifi50                      *RuckusWirelessApZoneRetrieveGet200ResponseWifi50                   `json:"wifi50,omitempty"`
		ZoneAffinityProfileID       *string                                                             `json:"zoneAffinityProfileId,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveGet: Use this API command to retrieve Ruckus Wireless AP zones configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveGet(ctx context.Context, id string) (*http.Response, *RuckusWirelessApZoneRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{id}", true)
	request.SetPathParameter("id", id)
	out := &RuckusWirelessApZoneRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneModifyBasicPatchRequest struct {
		AwsVenue                    *string  `json:"awsVenue,omitempty"`
		BonjourFencingPolicyEnabled *bool    `json:"bonjourFencingPolicyEnabled,omitempty"`
		ChannelEvaluationInterval   *float64 `json:"channelEvaluationInterval,omitempty"`
		ChannelModeEnabled          *bool    `json:"channelModeEnabled,omitempty"`
		CountryCode                 *string  `json:"countryCode,omitempty"`
		Description                 *string  `json:"description,omitempty"`
		DfsChannelEnabled           *bool    `json:"dfsChannelEnabled,omitempty"`
		DomainID                    *string  `json:"domainId,omitempty"`
		DosBarringCheckPeriod       *float64 `json:"dosBarringCheckPeriod,omitempty"`
		DosBarringEnable            *float64 `json:"dosBarringEnable,omitempty"`
		DosBarringPeriod            *float64 `json:"dosBarringPeriod,omitempty"`
		DosBarringThreshold         *float64 `json:"dosBarringThreshold,omitempty"`
		Ipv6TrafficFilterEnabled    *float64 `json:"ipv6TrafficFilterEnabled,omitempty"`
		Latitude                    *float64 `json:"latitude,omitempty"`
		Location                    *string  `json:"location,omitempty"`
		LocationAdditionalInfo      *string  `json:"locationAdditionalInfo,omitempty"`
		Longitude                   *float64 `json:"longitude,omitempty"`
		Name                        *string  `json:"name,omitempty"`
		TunnelType                  *string  `json:"tunnelType,omitempty"`
		VlanOverlappingEnabled      *bool    `json:"vlanOverlappingEnabled,omitempty"`
		ZoneAffinityProfileID       *string  `json:"zoneAffinityProfileId,omitempty"`
	}
)

// RuckusWirelessApZoneModifyBasicPatch: Use this API command to modify the basic information of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyBasicPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableAltitudeDelete: Use this API command to disable altitude configuration of zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableAltitudeDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/altitude", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyAltitudePatchRequest struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}
)

// RuckusWirelessApZoneModifyAltitudePatch: Use this API command to modify the altitude configuration of zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyAltitudePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyAltitudePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyAltitudePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/altitude", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyApManagementVlanPatchRequest struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}
)

// RuckusWirelessApZoneModifyApManagementVlanPatch: Modify AP Management Vlan of a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyApManagementVlanPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyApManagementVlanPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyApManagementVlanPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/apMgmtVlan", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyApRebootTimeoutPatchRequest struct {
		GatewayLossTimeoutInSec *float64 `json:"gatewayLossTimeoutInSec,omitempty"`
		ServerLossTimeoutInSec  *float64 `json:"serverLossTimeoutInSec,omitempty"`
	}
)

// RuckusWirelessApZoneModifyApRebootTimeoutPatch: Use this API command to modify AP reboot timeout for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyApRebootTimeoutPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyApRebootTimeoutPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyApRebootTimeoutPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/apRebootTimeout", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyRadio24gAutoChannelselectmodePatchRequest struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}
)

// RuckusWirelessApZoneModifyRadio24gAutoChannelselectmodePatch: Modify Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC of a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyRadio24gAutoChannelselectmodePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyRadio24gAutoChannelselectmodePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyRadio24gAutoChannelselectmodePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/autoChannelSelection24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyRadio5gAutoChannelselectmodePatchRequest struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}
)

// RuckusWirelessApZoneModifyRadio5gAutoChannelselectmodePatch: Modify Radio 5G Auto ChannelSelectMode and ChannelFly MTBC of a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyRadio5gAutoChannelselectmodePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyRadio5gAutoChannelselectmodePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyRadio5gAutoChannelselectmodePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/autoChannelSelection50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableBackgroundScanning24gDelete: Use this API command to disable background scanning 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableBackgroundScanning24gDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/backgroundScanning24", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyBackgroundScanning24gPatchRequest struct {
		FrequencyInSec *float64 `json:"frequencyInSec,omitempty"`
	}
)

// RuckusWirelessApZoneModifyBackgroundScanning24gPatch: Use this API command to modify the background scanning 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyBackgroundScanning24gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyBackgroundScanning24gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyBackgroundScanning24gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/backgroundScanning24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableBackgroundScanning5gDelete: Use this API command to disable background scanning 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableBackgroundScanning5gDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/backgroundScanning50", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyBackgroundScanning5gPatchRequest struct {
		FrequencyInSec *float64 `json:"frequencyInSec,omitempty"`
	}
)

// RuckusWirelessApZoneModifyBackgroundScanning5gPatch: Use this API command to modify the background scanning 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyBackgroundScanning5gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyBackgroundScanning5gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyBackgroundScanning5gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/backgroundScanning50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableBandBalancingDelete: Use this API command to disable band balancing for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableBandBalancingDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/bandBalancing", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyBandBalancingPatchRequest struct {
		Wifi24Percentage *float64 `json:"wifi24Percentage,omitempty"`
	}
)

// RuckusWirelessApZoneModifyBandBalancingPatch: Use this API command to modify band balancing for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyBandBalancingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyBandBalancingPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyBandBalancingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/bandBalancing", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyBonjourFencingPolicyConfigurationPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// RuckusWirelessApZoneModifyBonjourFencingPolicyConfigurationPatch: Use this API command to modify Bonjour Fencing Policy configuration of Zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyBonjourFencingPolicyConfigurationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyBonjourFencingPolicyConfigurationPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyBonjourFencingPolicyConfigurationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/bonjourFencingPolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableClientAdmissionControl24gDelete: Use this API command to disable client admission control 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableClientAdmissionControl24gDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/clientAdmissionControl24", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyClientAdmissionControl24gPatchRequest struct {
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}
)

// RuckusWirelessApZoneModifyClientAdmissionControl24gPatch: Use this API command to modify the client admission control 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyClientAdmissionControl24gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyClientAdmissionControl24gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyClientAdmissionControl24gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/clientAdmissionControl24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableClientAdmissionControl5gDelete: Use this API command to disable client admission control 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableClientAdmissionControl5gDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/clientAdmissionControl50", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyClientAdmissionControl5gPatchRequest struct {
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}
)

// RuckusWirelessApZoneModifyClientAdmissionControl5gPatch: Use this API command to modify the client admission control 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyClientAdmissionControl5gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyClientAdmissionControl5gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyClientAdmissionControl5gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/clientAdmissionControl50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableClientLoadBalancing24gDelete: Use this API command to disable client load balancing 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableClientLoadBalancing24gDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/clientLoadBalancing24", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyClientLoadBalancing24gPatchRequest struct {
		AdjacentRadioThreshold *float64 `json:"adjacentRadioThreshold,omitempty"`
	}
)

// RuckusWirelessApZoneModifyClientLoadBalancing24gPatch: Use this API command to modify the client load balancing 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyClientLoadBalancing24gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyClientLoadBalancing24gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyClientLoadBalancing24gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/clientLoadBalancing24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableClientLoadBalancing5gDelete: Use this API command to disable client load balancing 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableClientLoadBalancing5gDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/clientLoadBalancing50", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyClientLoadBalancing5gPatchRequest struct {
		AdjacentRadioThreshold *float64 `json:"adjacentRadioThreshold,omitempty"`
	}
)

// RuckusWirelessApZoneModifyClientLoadBalancing5gPatch: Use this API command to modify the client load balancing 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyClientLoadBalancing5gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyClientLoadBalancing5gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyClientLoadBalancing5gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/clientLoadBalancing50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequestSiteApsSlice []*RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequestSiteAps

	RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequestSiteAps struct {
		ApGatewayIP     *string `json:"apGatewayIp,omitempty"`
		ApMac           *string `json:"apMac,omitempty"`
		ApName          *string `json:"apName,omitempty"`
		ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
		ApServerIP      *string `json:"apServerIp,omitempty"`
		ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
		ApStatus        *string `json:"apStatus,omitempty"`
	}

	RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequestSiteProfileIdsSlice []*string

	RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequest struct {
		ManualSelect   *bool                                                                                `json:"manualSelect,omitempty"`
		SiteAps        RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequestSiteApsSlice        `json:"siteAps,omitempty"`
		SiteEnabled    *bool                                                                                `json:"siteEnabled,omitempty"`
		SiteMode       *string                                                                              `json:"siteMode,omitempty"`
		SiteProfileIds RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequestSiteProfileIdsSlice `json:"siteProfileIds,omitempty"`
	}
)

// RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatch: Use this API command to modify DHCP/NAT service configuration of Zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyDhcpNatServiceConfigurationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/dhcpSiteConfig", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyIpsecProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// RuckusWirelessApZoneModifyIpsecProfilePatch: Modify IPsec Profile of a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyIpsecProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyIpsecProfilePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyIpsecProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/ipsecProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableLocationBasedServiceDelete: Use this API command to disable location based service for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableLocationBasedServiceDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/locationBasedService", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyLocationBasedServicePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// RuckusWirelessApZoneModifyLocationBasedServicePatch: Use this API command to modify location based service for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyLocationBasedServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyLocationBasedServicePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyLocationBasedServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/locationBasedService", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyApLogonPatchRequest struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}
)

// RuckusWirelessApZoneModifyApLogonPatch: Use this API command to modify the AP logon information for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyApLogonPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyApLogonPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyApLogonPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/login", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneClearMeshConfigurationDelete: Use this API command to disable mesh networking.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneClearMeshConfigurationDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/mesh", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneRetrieveMeshConfigurationGet200Response struct {
		Passphrase *string `json:"passphrase,omitempty"`
		Ssid       *string `json:"ssid,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveMeshConfigurationGet: Use this API command to retrieve the mesh configuration of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveMeshConfigurationGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveMeshConfigurationGet(ctx context.Context, id string) (*http.Response, *RuckusWirelessApZoneRetrieveMeshConfigurationGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{id}/mesh", true)
	request.SetPathParameter("id", id)
	out := &RuckusWirelessApZoneRetrieveMeshConfigurationGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneModifyMeshConfigurationPatchRequest struct {
		Passphrase *string `json:"passphrase,omitempty"`
		Ssid       *string `json:"ssid,omitempty"`
	}
)

// RuckusWirelessApZoneModifyMeshConfigurationPatch: Use this API command to enable mesh networking or update the mesh configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyMeshConfigurationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyMeshConfigurationPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyMeshConfigurationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/mesh", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyNodeAffinityProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// RuckusWirelessApZoneModifyNodeAffinityProfilePatch: Use this API command to modify node affinity profile for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyNodeAffinityProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyNodeAffinityProfilePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyNodeAffinityProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/nodeAffinityProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableZoneRecoverySsidDelete: Disable recovery ssid setting of a zone(setup a password, or enable/disable)
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableZoneRecoverySsidDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/recoverySsid", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyZoneRecoverySsidPatchRequest struct {
		RecoverySsidEnable *float64 `json:"recoverySsidEnable,omitempty"`
		RecoverySsidPskKey *string  `json:"recoverySsidPskKey,omitempty"`
	}
)

// RuckusWirelessApZoneModifyZoneRecoverySsidPatch: Modify recovery ssid setting of a zone(setup a password, or enable/disable)
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyZoneRecoverySsidPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyZoneRecoverySsidPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyZoneRecoverySsidPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/recoverySsid", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableRogueDelete: Use this API command to disable rogue AP detection for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableRogueDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/rogue", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyRoguePatchRequestMaliciousTypesSlice []*string

	RuckusWirelessApZoneModifyRoguePatchRequest struct {
		MaliciousTypes    RuckusWirelessApZoneModifyRoguePatchRequestMaliciousTypesSlice `json:"maliciousTypes,omitempty"`
		ProtectionEnabled *bool                                                          `json:"protectionEnabled,omitempty"`
		ReportType        *string                                                        `json:"reportType,omitempty"`
	}
)

// RuckusWirelessApZoneModifyRoguePatch: Use this API command to modify the rogue AP detection for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyRoguePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyRoguePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyRoguePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/rogue", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableSmartMonitorDelete: Use this API command to disable smart monitor for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableSmartMonitorDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/smartMonitor", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifySmartMonitorPatchRequest struct {
		IntervalInSec  *int `json:"intervalInSec,omitempty"`
		RetryThreshold *int `json:"retryThreshold,omitempty"`
	}
)

// RuckusWirelessApZoneModifySmartMonitorPatch: Use this API command to modify the smart monitor for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifySmartMonitorPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifySmartMonitorPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifySmartMonitorPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/smartMonitor", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneClearApSnmpOptionsDelete: Use this API command to clear SNMPv2 and SNMPv3 agent that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneClearApSnmpOptionsDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/snmpAgent", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2AgentSlice []*RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2Agent

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2AgentNotificationTargetSlice []*RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2AgentNotificationTarget

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2Agent struct {
		CommunityName       *string                                                                               `json:"communityName,omitempty"`
		NotificationEnabled *bool                                                                                 `json:"notificationEnabled,omitempty"`
		NotificationTarget  RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                                               `json:"notificationType,omitempty"`
		ReadEnabled         *bool                                                                                 `json:"readEnabled,omitempty"`
		WriteEnabled        *bool                                                                                 `json:"writeEnabled,omitempty"`
	}

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3AgentSlice []*RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3Agent

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3AgentNotificationTargetSlice []*RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3AgentNotificationTarget

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3Agent struct {
		AuthPassword        *string                                                                               `json:"authPassword,omitempty"`
		AuthProtocol        *string                                                                               `json:"authProtocol,omitempty"`
		NotificationEnabled *bool                                                                                 `json:"notificationEnabled,omitempty"`
		NotificationTarget  RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                                               `json:"notificationType,omitempty"`
		PrivPassword        *string                                                                               `json:"privPassword,omitempty"`
		PrivProtocol        *string                                                                               `json:"privProtocol,omitempty"`
		ReadEnabled         *bool                                                                                 `json:"readEnabled,omitempty"`
		UserName            *string                                                                               `json:"userName,omitempty"`
		WriteEnabled        *bool                                                                                 `json:"writeEnabled,omitempty"`
	}

	RuckusWirelessApZoneModifyApSnmpOptionsPatchRequest struct {
		ApSnmpEnabled *bool                                                               `json:"apSnmpEnabled,omitempty"`
		SnmpV2Agent   RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV2AgentSlice `json:"snmpV2Agent,omitempty"`
		SnmpV3Agent   RuckusWirelessApZoneModifyApSnmpOptionsPatchRequestSnmpV3AgentSlice `json:"snmpV3Agent,omitempty"`
	}
)

// RuckusWirelessApZoneModifyApSnmpOptionsPatch: Use this API command to modify SNMPv2 and SNMPv3 agent that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyApSnmpOptionsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyApSnmpOptionsPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyApSnmpOptionsPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/snmpAgent", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableSyslogDelete: Use this API command to disable syslog configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableSyslogDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/syslog", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifySyslogPatchRequest struct {
		Address  *string `json:"address,omitempty"`
		Facility *string `json:"facility,omitempty"`
		Port     *int    `json:"port,omitempty"`
		Priority *string `json:"priority,omitempty"`
	}
)

// RuckusWirelessApZoneModifySyslogPatch: Use this API command to modify the syslog configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifySyslogPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifySyslogPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifySyslogPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/syslog", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyTimeZonePatchRequestCustomizedTimezoneEnd struct {
		Day   *float64 `json:"day,omitempty"`
		Hour  *float64 `json:"hour,omitempty"`
		Month *float64 `json:"month,omitempty"`
		Week  *float64 `json:"week,omitempty"`
	}

	RuckusWirelessApZoneModifyTimeZonePatchRequestCustomizedTimezoneStart struct {
		Day   *float64 `json:"day,omitempty"`
		Hour  *float64 `json:"hour,omitempty"`
		Month *float64 `json:"month,omitempty"`
		Week  *float64 `json:"week,omitempty"`
	}

	RuckusWirelessApZoneModifyTimeZonePatchRequestCustomizedTimezone struct {
		Abbreviation    *string                                                                `json:"abbreviation,omitempty"`
		End             *RuckusWirelessApZoneModifyTimeZonePatchRequestCustomizedTimezoneEnd   `json:"end,omitempty"`
		GmtOffset       *float64                                                               `json:"gmtOffset,omitempty"`
		GmtOffsetMinute *float64                                                               `json:"gmtOffsetMinute,omitempty"`
		Start           *RuckusWirelessApZoneModifyTimeZonePatchRequestCustomizedTimezoneStart `json:"start,omitempty"`
	}

	RuckusWirelessApZoneModifyTimeZonePatchRequest struct {
		CustomizedTimezone *RuckusWirelessApZoneModifyTimeZonePatchRequestCustomizedTimezone `json:"customizedTimezone,omitempty"`
		SystemTimezone     *string                                                           `json:"systemTimezone,omitempty"`
	}
)

// RuckusWirelessApZoneModifyTimeZonePatch: Use this API command to modify the time zone of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyTimeZonePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyTimeZonePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyTimeZonePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/timezone", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyTunnelProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// RuckusWirelessApZoneModifyTunnelProfilePatch: Use this API command to change tunnel profile of Zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyTunnelProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyTunnelProfilePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyTunnelProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/tunnelProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneDisableApUsbSoftwarePackageDelete: Disable AP Usb Software Package of a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - applyModel (string): Specify ap models. ex : applyModel=ZF7321U&applyModel=ZF7323
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneDisableApUsbSoftwarePackageDelete(ctx context.Context, id string, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/usbSoftwarePackage", true)
	request.SetPathParameter("id", id)
	request.SetQueryParameter("applyModel", optionalParams["applyModel"])
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequestApplyModelSlice []*string

	RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequestUsbSoftware struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequest struct {
		ApplyModel  RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequestApplyModelSlice `json:"applyModel,omitempty"`
		UsbSoftware *RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequestUsbSoftware    `json:"usbSoftware,omitempty"`
	}
)

// RuckusWirelessApZoneModifyApUsbSoftwarePackagePatch: Modify AP Usb Software Package of a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyApUsbSoftwarePackagePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyApUsbSoftwarePackagePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/usbSoftwarePackage", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// RuckusWirelessApZoneClearHotspot20VenueProfileDelete: Use this API command to clear Hotspot 2.0 venue profile for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneClearHotspot20VenueProfileDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{id}/venueProfile", true)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyHotspot20VenueProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// RuckusWirelessApZoneModifyHotspot20VenueProfilePatch: Use this API command to modify Hotspot 2.0 venue profile for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyHotspot20VenueProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyHotspot20VenueProfilePatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyHotspot20VenueProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/venueProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyRadio24gPatchRequestChannelRangeSlice []*float64

	RuckusWirelessApZoneModifyRadio24gPatchRequest struct {
		Channel      *int                                                            `json:"channel,omitempty"`
		ChannelRange RuckusWirelessApZoneModifyRadio24gPatchRequestChannelRangeSlice `json:"channelRange,omitempty"`
		ChannelWidth *int                                                            `json:"channelWidth,omitempty"`
		TxPower      *string                                                         `json:"txPower,omitempty"`
	}
)

// RuckusWirelessApZoneModifyRadio24gPatch: Use this API command to modify the 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyRadio24gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyRadio24gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyRadio24gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/wifi24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneModifyRadio5gPatchRequestIndoorChannelRangeSlice []*float64

	RuckusWirelessApZoneModifyRadio5gPatchRequestOutdoorChannelRangeSlice []*float64

	RuckusWirelessApZoneModifyRadio5gPatchRequest struct {
		ChannelWidth            *int                                                                  `json:"channelWidth,omitempty"`
		IndoorChannel           *float64                                                              `json:"indoorChannel,omitempty"`
		IndoorChannelRange      RuckusWirelessApZoneModifyRadio5gPatchRequestIndoorChannelRangeSlice  `json:"indoorChannelRange,omitempty"`
		IndoorSecondaryChannel  *float64                                                              `json:"indoorSecondaryChannel,omitempty"`
		OutdoorChannel          *float64                                                              `json:"outdoorChannel,omitempty"`
		OutdoorChannelRange     RuckusWirelessApZoneModifyRadio5gPatchRequestOutdoorChannelRangeSlice `json:"outdoorChannelRange,omitempty"`
		OutdoorSecondaryChannel *float64                                                              `json:"outdoorSecondaryChannel,omitempty"`
		TxPower                 *string                                                               `json:"txPower,omitempty"`
	}
)

// RuckusWirelessApZoneModifyRadio5gPatch: Use this API command to modify the 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneModifyRadio5gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyRadio5gPatch(ctx context.Context, id string, requestBody *RuckusWirelessApZoneModifyRadio5gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{id}/wifi50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveListRadiusAccountingGet200ResponseListSlice []*ZoneAaaRetrieveListRadiusAccountingGet200ResponseList

	ZoneAaaRetrieveListRadiusAccountingGet200ResponseListPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveListRadiusAccountingGet200ResponseListSecondary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveListRadiusAccountingGet200ResponseList struct {
		Description *string                                                         `json:"description,omitempty"`
		ID          *string                                                         `json:"id,omitempty"`
		MvnoID      *string                                                         `json:"mvnoId,omitempty"`
		Name        *string                                                         `json:"name,omitempty"`
		Primary     *ZoneAaaRetrieveListRadiusAccountingGet200ResponseListPrimary   `json:"primary,omitempty"`
		Secondary   *ZoneAaaRetrieveListRadiusAccountingGet200ResponseListSecondary `json:"secondary,omitempty"`
		ZoneID      *string                                                         `json:"zoneId,omitempty"`
	}

	ZoneAaaRetrieveListRadiusAccountingGet200Response struct {
		FirstIndex *int                                                       `json:"firstIndex,omitempty"`
		HasMore    *bool                                                      `json:"hasMore,omitempty"`
		List       ZoneAaaRetrieveListRadiusAccountingGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                       `json:"totalCount,omitempty"`
	}
)

// ZoneAaaRetrieveListRadiusAccountingGet: Use this API command to retrieve a list of radius accounting servers of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveListRadiusAccountingGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveListRadiusAccountingGet(ctx context.Context, zoneId string) (*http.Response, *ZoneAaaRetrieveListRadiusAccountingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/accounting", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaRetrieveListRadiusAccountingGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaCreateRadiusAccountingPostRequestPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaCreateRadiusAccountingPostRequestSecondary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaCreateRadiusAccountingPostRequest struct {
		Description *string                                            `json:"description,omitempty"`
		Name        *string                                            `json:"name,omitempty"`
		Primary     *ZoneAaaCreateRadiusAccountingPostRequestPrimary   `json:"primary,omitempty"`
		Secondary   *ZoneAaaCreateRadiusAccountingPostRequestSecondary `json:"secondary,omitempty"`
	}

	ZoneAaaCreateRadiusAccountingPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ZoneAaaCreateRadiusAccountingPost: Use this API command to create a new radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *ZoneAaaCreateRadiusAccountingPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaCreateRadiusAccountingPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaCreateRadiusAccountingPost(ctx context.Context, zoneId string, requestBody *ZoneAaaCreateRadiusAccountingPostRequest) (*http.Response, *ZoneAaaCreateRadiusAccountingPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/aaa/accounting", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaCreateRadiusAccountingPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ZoneAaaDeleteRadiusAccountingDelete: Use this API command to delete a radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaDeleteRadiusAccountingDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/aaa/accounting/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveRadiusAccountingGet200ResponsePrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveRadiusAccountingGet200ResponseSecondary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveRadiusAccountingGet200Response struct {
		Description *string                                                 `json:"description,omitempty"`
		ID          *string                                                 `json:"id,omitempty"`
		MvnoID      *string                                                 `json:"mvnoId,omitempty"`
		Name        *string                                                 `json:"name,omitempty"`
		Primary     *ZoneAaaRetrieveRadiusAccountingGet200ResponsePrimary   `json:"primary,omitempty"`
		Secondary   *ZoneAaaRetrieveRadiusAccountingGet200ResponseSecondary `json:"secondary,omitempty"`
		ZoneID      *string                                                 `json:"zoneId,omitempty"`
	}
)

// ZoneAaaRetrieveRadiusAccountingGet: Use this API command to retrieve a radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveRadiusAccountingGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveRadiusAccountingGet(ctx context.Context, zoneId string, id string) (*http.Response, *ZoneAaaRetrieveRadiusAccountingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/accounting/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &ZoneAaaRetrieveRadiusAccountingGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaModifyRadiusAccountingPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// ZoneAaaModifyRadiusAccountingPatch: Use this API command to modify the basic information on radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifyRadiusAccountingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifyRadiusAccountingPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifyRadiusAccountingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/accounting/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaModifyPrimaryServerOfRadiusAccountingPatchRequest struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}
)

// ZoneAaaModifyPrimaryServerOfRadiusAccountingPatch: Use this API command to modify primary server on radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifyPrimaryServerOfRadiusAccountingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifyPrimaryServerOfRadiusAccountingPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifyPrimaryServerOfRadiusAccountingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/accounting/{id}/primary", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ZoneAaaDisableSecondaryServerRadiusAccountingDelete: Use this API command to disable secondary server on radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaDisableSecondaryServerRadiusAccountingDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/aaa/accounting/{id}/secondary", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaModifySecondaryServerOfRadiusAccountingPatchRequest struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}
)

// ZoneAaaModifySecondaryServerOfRadiusAccountingPatch: Use this API command to modify secondary server on radius accounting server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifySecondaryServerOfRadiusAccountingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifySecondaryServerOfRadiusAccountingPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifySecondaryServerOfRadiusAccountingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/accounting/{id}/secondary", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveListActivedirectoryGet200ResponseListSlice []*ZoneAaaRetrieveListActivedirectoryGet200ResponseList

	ZoneAaaRetrieveListActivedirectoryGet200ResponseList struct {
		AdminDomainName      *string `json:"adminDomainName,omitempty"`
		Description          *string `json:"description,omitempty"`
		GlobalCatalogEnabled *bool   `json:"globalCatalogEnabled,omitempty"`
		ID                   *string `json:"id,omitempty"`
		IP                   *string `json:"ip,omitempty"`
		MvnoID               *string `json:"mvnoId,omitempty"`
		Name                 *string `json:"name,omitempty"`
		Password             *string `json:"password,omitempty"`
		Port                 *int    `json:"port,omitempty"`
		WindowsDomainName    *string `json:"windowsDomainName,omitempty"`
		ZoneID               *string `json:"zoneId,omitempty"`
	}

	ZoneAaaRetrieveListActivedirectoryGet200Response struct {
		FirstIndex *int                                                      `json:"firstIndex,omitempty"`
		HasMore    *bool                                                     `json:"hasMore,omitempty"`
		List       ZoneAaaRetrieveListActivedirectoryGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                      `json:"totalCount,omitempty"`
	}
)

// ZoneAaaRetrieveListActivedirectoryGet: Use this API command to retrieve a list of active directory servers of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveListActivedirectoryGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveListActivedirectoryGet(ctx context.Context, zoneId string) (*http.Response, *ZoneAaaRetrieveListActivedirectoryGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/ad", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaRetrieveListActivedirectoryGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaCreateActivedirectoryPostRequest struct {
		AdminDomainName      *string `json:"adminDomainName,omitempty"`
		Description          *string `json:"description,omitempty"`
		GlobalCatalogEnabled *bool   `json:"globalCatalogEnabled,omitempty"`
		IP                   *string `json:"ip,omitempty"`
		Name                 *string `json:"name,omitempty"`
		Password             *string `json:"password,omitempty"`
		Port                 *int    `json:"port,omitempty"`
		WindowsDomainName    *string `json:"windowsDomainName,omitempty"`
	}

	ZoneAaaCreateActivedirectoryPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ZoneAaaCreateActivedirectoryPost: Use this API command to create a new active directory server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *ZoneAaaCreateActivedirectoryPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaCreateActivedirectoryPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaCreateActivedirectoryPost(ctx context.Context, zoneId string, requestBody *ZoneAaaCreateActivedirectoryPostRequest) (*http.Response, *ZoneAaaCreateActivedirectoryPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/aaa/ad", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaCreateActivedirectoryPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ZoneAaaDeleteActivedirectoryDelete: Use this API command to delete an active directory server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaDeleteActivedirectoryDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/aaa/ad/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveActivedirectoryGet200Response struct {
		AdminDomainName      *string `json:"adminDomainName,omitempty"`
		Description          *string `json:"description,omitempty"`
		GlobalCatalogEnabled *bool   `json:"globalCatalogEnabled,omitempty"`
		ID                   *string `json:"id,omitempty"`
		IP                   *string `json:"ip,omitempty"`
		MvnoID               *string `json:"mvnoId,omitempty"`
		Name                 *string `json:"name,omitempty"`
		Password             *string `json:"password,omitempty"`
		Port                 *int    `json:"port,omitempty"`
		WindowsDomainName    *string `json:"windowsDomainName,omitempty"`
		ZoneID               *string `json:"zoneId,omitempty"`
	}
)

// ZoneAaaRetrieveActivedirectoryGet: Use this API command to retrieve an active directory server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveActivedirectoryGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveActivedirectoryGet(ctx context.Context, zoneId string, id string) (*http.Response, *ZoneAaaRetrieveActivedirectoryGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/ad/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &ZoneAaaRetrieveActivedirectoryGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaModifyActivedirectoryPatchRequest struct {
		AdminDomainName      *string `json:"adminDomainName,omitempty"`
		Description          *string `json:"description,omitempty"`
		GlobalCatalogEnabled *bool   `json:"globalCatalogEnabled,omitempty"`
		IP                   *string `json:"ip,omitempty"`
		Name                 *string `json:"name,omitempty"`
		Password             *string `json:"password,omitempty"`
		Port                 *int    `json:"port,omitempty"`
		WindowsDomainName    *string `json:"windowsDomainName,omitempty"`
	}
)

// ZoneAaaModifyActivedirectoryPatch: Use this API command to modify the basic information on active directory server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifyActivedirectoryPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifyActivedirectoryPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifyActivedirectoryPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/ad/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveListLdapGet200ResponseListSlice []*ZoneAaaRetrieveListLdapGet200ResponseList

	ZoneAaaRetrieveListLdapGet200ResponseList struct {
		AdminDomainName *string `json:"adminDomainName,omitempty"`
		BaseDomainName  *string `json:"baseDomainName,omitempty"`
		Description     *string `json:"description,omitempty"`
		ID              *string `json:"id,omitempty"`
		IP              *string `json:"ip,omitempty"`
		KeyAttribute    *string `json:"keyAttribute,omitempty"`
		MvnoID          *string `json:"mvnoId,omitempty"`
		Name            *string `json:"name,omitempty"`
		Password        *string `json:"password,omitempty"`
		Port            *int    `json:"port,omitempty"`
		SearchFilter    *string `json:"searchFilter,omitempty"`
		ZoneID          *string `json:"zoneId,omitempty"`
	}

	ZoneAaaRetrieveListLdapGet200Response struct {
		FirstIndex *int                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                          `json:"hasMore,omitempty"`
		List       ZoneAaaRetrieveListLdapGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                           `json:"totalCount,omitempty"`
	}
)

// ZoneAaaRetrieveListLdapGet: Use this API command to retrieve a list of LDAP servers of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveListLdapGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveListLdapGet(ctx context.Context, zoneId string) (*http.Response, *ZoneAaaRetrieveListLdapGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/ldap", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaRetrieveListLdapGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaCreateLdapPostRequest struct {
		AdminDomainName *string `json:"adminDomainName,omitempty"`
		BaseDomainName  *string `json:"baseDomainName,omitempty"`
		Description     *string `json:"description,omitempty"`
		IP              *string `json:"ip,omitempty"`
		KeyAttribute    *string `json:"keyAttribute,omitempty"`
		Name            *string `json:"name,omitempty"`
		Password        *string `json:"password,omitempty"`
		Port            *int    `json:"port,omitempty"`
		SearchFilter    *string `json:"searchFilter,omitempty"`
	}

	ZoneAaaCreateLdapPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ZoneAaaCreateLdapPost: Use this API command to create a new LDAP server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *ZoneAaaCreateLdapPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaCreateLdapPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaCreateLdapPost(ctx context.Context, zoneId string, requestBody *ZoneAaaCreateLdapPostRequest) (*http.Response, *ZoneAaaCreateLdapPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/aaa/ldap", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaCreateLdapPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ZoneAaaDeleteLdapDelete: Use this API command to delete a LDAP server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaDeleteLdapDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/aaa/ldap/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveLdapGet200Response struct {
		AdminDomainName *string `json:"adminDomainName,omitempty"`
		BaseDomainName  *string `json:"baseDomainName,omitempty"`
		Description     *string `json:"description,omitempty"`
		ID              *string `json:"id,omitempty"`
		IP              *string `json:"ip,omitempty"`
		KeyAttribute    *string `json:"keyAttribute,omitempty"`
		MvnoID          *string `json:"mvnoId,omitempty"`
		Name            *string `json:"name,omitempty"`
		Password        *string `json:"password,omitempty"`
		Port            *int    `json:"port,omitempty"`
		SearchFilter    *string `json:"searchFilter,omitempty"`
		ZoneID          *string `json:"zoneId,omitempty"`
	}
)

// ZoneAaaRetrieveLdapGet: Use this API command to retrieve a LDAP server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveLdapGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveLdapGet(ctx context.Context, zoneId string, id string) (*http.Response, *ZoneAaaRetrieveLdapGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/ldap/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &ZoneAaaRetrieveLdapGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaModifyLdapPatchRequest struct {
		AdminDomainName *string `json:"adminDomainName,omitempty"`
		BaseDomainName  *string `json:"baseDomainName,omitempty"`
		Description     *string `json:"description,omitempty"`
		IP              *string `json:"ip,omitempty"`
		KeyAttribute    *string `json:"keyAttribute,omitempty"`
		Name            *string `json:"name,omitempty"`
		Password        *string `json:"password,omitempty"`
		Port            *int    `json:"port,omitempty"`
		SearchFilter    *string `json:"searchFilter,omitempty"`
	}
)

// ZoneAaaModifyLdapPatch: Use this API command to modify the basic information on LDAP server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifyLdapPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifyLdapPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifyLdapPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/ldap/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveListRadiusGet200ResponseListSlice []*ZoneAaaRetrieveListRadiusGet200ResponseList

	ZoneAaaRetrieveListRadiusGet200ResponseListPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveListRadiusGet200ResponseListSecondary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveListRadiusGet200ResponseList struct {
		Description *string                                               `json:"description,omitempty"`
		ID          *string                                               `json:"id,omitempty"`
		MvnoID      *string                                               `json:"mvnoId,omitempty"`
		Name        *string                                               `json:"name,omitempty"`
		Primary     *ZoneAaaRetrieveListRadiusGet200ResponseListPrimary   `json:"primary,omitempty"`
		Secondary   *ZoneAaaRetrieveListRadiusGet200ResponseListSecondary `json:"secondary,omitempty"`
		ZoneID      *string                                               `json:"zoneId,omitempty"`
	}

	ZoneAaaRetrieveListRadiusGet200Response struct {
		FirstIndex *int                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                            `json:"hasMore,omitempty"`
		List       ZoneAaaRetrieveListRadiusGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                             `json:"totalCount,omitempty"`
	}
)

// ZoneAaaRetrieveListRadiusGet: Use this API command to retrieve a list of radius servers of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveListRadiusGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveListRadiusGet(ctx context.Context, zoneId string) (*http.Response, *ZoneAaaRetrieveListRadiusGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/radius", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaRetrieveListRadiusGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaCreateRadiusPostRequestPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaCreateRadiusPostRequestSecondary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaCreateRadiusPostRequest struct {
		Description *string                                  `json:"description,omitempty"`
		Name        *string                                  `json:"name,omitempty"`
		Primary     *ZoneAaaCreateRadiusPostRequestPrimary   `json:"primary,omitempty"`
		Secondary   *ZoneAaaCreateRadiusPostRequestSecondary `json:"secondary,omitempty"`
	}

	ZoneAaaCreateRadiusPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ZoneAaaCreateRadiusPost: Use this API command to create a new radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *ZoneAaaCreateRadiusPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaCreateRadiusPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaCreateRadiusPost(ctx context.Context, zoneId string, requestBody *ZoneAaaCreateRadiusPostRequest) (*http.Response, *ZoneAaaCreateRadiusPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/aaa/radius", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &ZoneAaaCreateRadiusPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ZoneAaaDeleteRadiusDelete: Use this API command to delete a radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaDeleteRadiusDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/aaa/radius/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaRetrieveRadiusGet200ResponsePrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveRadiusGet200ResponseSecondary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	ZoneAaaRetrieveRadiusGet200Response struct {
		Description *string                                       `json:"description,omitempty"`
		ID          *string                                       `json:"id,omitempty"`
		MvnoID      *string                                       `json:"mvnoId,omitempty"`
		Name        *string                                       `json:"name,omitempty"`
		Primary     *ZoneAaaRetrieveRadiusGet200ResponsePrimary   `json:"primary,omitempty"`
		Secondary   *ZoneAaaRetrieveRadiusGet200ResponseSecondary `json:"secondary,omitempty"`
		ZoneID      *string                                       `json:"zoneId,omitempty"`
	}
)

// ZoneAaaRetrieveRadiusGet: Use this API command to retrieve a radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ZoneAaaRetrieveRadiusGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaRetrieveRadiusGet(ctx context.Context, zoneId string, id string) (*http.Response, *ZoneAaaRetrieveRadiusGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/aaa/radius/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &ZoneAaaRetrieveRadiusGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ZoneAaaModifyRadiusPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// ZoneAaaModifyRadiusPatch: Use this API command to modify the basic information on radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifyRadiusPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifyRadiusPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifyRadiusPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/radius/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaModifyPrimaryServerOfRadiusPatchRequest struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}
)

// ZoneAaaModifyPrimaryServerOfRadiusPatch: Use this API command to modify primary server on radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifyPrimaryServerOfRadiusPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifyPrimaryServerOfRadiusPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifyPrimaryServerOfRadiusPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/radius/{id}/primary", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ZoneAaaDisableSecondaryServerRadiusDelete: Use this API command to disable secondary server on radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaDisableSecondaryServerRadiusDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/aaa/radius/{id}/secondary", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ZoneAaaModifySecondaryServerOfRadiusPatchRequest struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}
)

// ZoneAaaModifySecondaryServerOfRadiusPatch: Use this API command to modify secondary server on radius server of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ZoneAaaModifySecondaryServerOfRadiusPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ZoneAaaModifySecondaryServerOfRadiusPatch(ctx context.Context, zoneId string, id string, requestBody *ZoneAaaModifySecondaryServerOfRadiusPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/aaa/radius/{id}/secondary", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseListSlice []*RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseList

	RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseListUnsupportedApModelSummarySlice []*RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseListUnsupportedApModelSummary

	RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseListUnsupportedApModelSummary struct {
		Amount *float64 `json:"amount,omitempty"`
		Model  *string  `json:"model,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseList struct {
		FirmwareVersion           *string                                                                                    `json:"firmwareVersion,omitempty"`
		Supported                 *bool                                                                                      `json:"supported,omitempty"`
		UnsupportedApModelSummary RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseListUnsupportedApModelSummarySlice `json:"unsupportedApModelSummary,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApFirmwareListGet200Response struct {
		FirstIndex *int                                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                                             `json:"hasMore,omitempty"`
		List       RuckusWirelessApZoneRetrieveApFirmwareListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                              `json:"totalCount,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveApFirmwareListGet: Use this API command to retrieve AP Firmware the list that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveApFirmwareListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveApFirmwareListGet(ctx context.Context, zoneId string) (*http.Response, *RuckusWirelessApZoneRetrieveApFirmwareListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apFirmware", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &RuckusWirelessApZoneRetrieveApFirmwareListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneChangeApFirmwarePutRequest struct {
		FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	}
)

// RuckusWirelessApZoneChangeApFirmwarePut: Use this API command to change the AP Firmware that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneChangeApFirmwarePutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneChangeApFirmwarePut(ctx context.Context, zoneId string, requestBody *RuckusWirelessApZoneChangeApFirmwarePutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/rkszones/{zoneId}/apFirmware", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupRetrieveListGet200ResponseListSlice []*ApGroupRetrieveListGet200ResponseList

	ApGroupRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveListGet200Response struct {
		FirstIndex *int                                       `json:"firstIndex,omitempty"`
		HasMore    *bool                                      `json:"hasMore,omitempty"`
		List       ApGroupRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                       `json:"totalCount,omitempty"`
	}
)

// ApGroupRetrieveListGet: Use this API command to retrieve the list of AP groups that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApGroupRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *ApGroupRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apgroups", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &ApGroupRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApGroupCreatePostRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	ApGroupCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ApGroupCreatePost: Use this API command to create new AP group within a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *ApGroupCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApGroupCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupCreatePost(ctx context.Context, zoneId string, requestBody *ApGroupCreatePostRequest) (*http.Response, *ApGroupCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/apgroups", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &ApGroupCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	ApGroupRetrieveGet1200ResponseAltitude struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseApMgmtVlan struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseAutoChannelSelection24 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseAutoChannelSelection50 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseClientAdmissionControl24 struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseClientAdmissionControl50 struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseLocationBasedService struct {
		Enabled *bool   `json:"enabled,omitempty"`
		ID      *string `json:"id,omitempty"`
		Name    *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseMembersSlice []*ApGroupRetrieveGet1200ResponseMembers

	ApGroupRetrieveGet1200ResponseMembers struct {
		ApMac *string `json:"apMac,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseVenueProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseWifi24AvailableChannelRangeSlice []*float64

	ApGroupRetrieveGet1200ResponseWifi24ChannelRangeSlice []*float64

	ApGroupRetrieveGet1200ResponseWifi24 struct {
		AvailableChannelRange ApGroupRetrieveGet1200ResponseWifi24AvailableChannelRangeSlice `json:"availableChannelRange,omitempty"`
		Channel               *int                                                           `json:"channel,omitempty"`
		ChannelRange          ApGroupRetrieveGet1200ResponseWifi24ChannelRangeSlice          `json:"channelRange,omitempty"`
		ChannelWidth          *int                                                           `json:"channelWidth,omitempty"`
		TxPower               *string                                                        `json:"txPower,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseWifi50AvailableIndoorChannelRangeSlice []*float64

	ApGroupRetrieveGet1200ResponseWifi50AvailableOutdoorChannelRangeSlice []*float64

	ApGroupRetrieveGet1200ResponseWifi50IndoorChannelRangeSlice []*float64

	ApGroupRetrieveGet1200ResponseWifi50OutdoorChannelRangeSlice []*float64

	ApGroupRetrieveGet1200ResponseWifi50 struct {
		AvailableIndoorChannelRange  ApGroupRetrieveGet1200ResponseWifi50AvailableIndoorChannelRangeSlice  `json:"availableIndoorChannelRange,omitempty"`
		AvailableOutdoorChannelRange ApGroupRetrieveGet1200ResponseWifi50AvailableOutdoorChannelRangeSlice `json:"availableOutdoorChannelRange,omitempty"`
		ChannelWidth                 *int                                                                  `json:"channelWidth,omitempty"`
		IndoorChannel                *float64                                                              `json:"indoorChannel,omitempty"`
		IndoorChannelRange           ApGroupRetrieveGet1200ResponseWifi50IndoorChannelRangeSlice           `json:"indoorChannelRange,omitempty"`
		IndoorSecondaryChannel       *float64                                                              `json:"indoorSecondaryChannel,omitempty"`
		OutdoorChannel               *float64                                                              `json:"outdoorChannel,omitempty"`
		OutdoorChannelRange          ApGroupRetrieveGet1200ResponseWifi50OutdoorChannelRangeSlice          `json:"outdoorChannelRange,omitempty"`
		OutdoorSecondaryChannel      *float64                                                              `json:"outdoorSecondaryChannel,omitempty"`
		TxPower                      *string                                                               `json:"txPower,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseWlanGroup24 struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet1200ResponseWlanGroup50 struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet1200Response struct {
		Altitude                  *ApGroupRetrieveGet1200ResponseAltitude                 `json:"altitude,omitempty"`
		ApMgmtVlan                *ApGroupRetrieveGet1200ResponseApMgmtVlan               `json:"apMgmtVlan,omitempty"`
		AutoChannelSelection24    *ApGroupRetrieveGet1200ResponseAutoChannelSelection24   `json:"autoChannelSelection24,omitempty"`
		AutoChannelSelection50    *ApGroupRetrieveGet1200ResponseAutoChannelSelection50   `json:"autoChannelSelection50,omitempty"`
		AwsVenue                  *string                                                 `json:"awsVenue,omitempty"`
		ChannelEvaluationInterval *float64                                                `json:"channelEvaluationInterval,omitempty"`
		ClientAdmissionControl24  *ApGroupRetrieveGet1200ResponseClientAdmissionControl24 `json:"clientAdmissionControl24,omitempty"`
		ClientAdmissionControl50  *ApGroupRetrieveGet1200ResponseClientAdmissionControl50 `json:"clientAdmissionControl50,omitempty"`
		Description               *string                                                 `json:"description,omitempty"`
		ID                        *string                                                 `json:"id,omitempty"`
		Latitude                  *float64                                                `json:"latitude,omitempty"`
		Location                  *string                                                 `json:"location,omitempty"`
		LocationAdditionalInfo    *string                                                 `json:"locationAdditionalInfo,omitempty"`
		LocationBasedService      *ApGroupRetrieveGet1200ResponseLocationBasedService     `json:"locationBasedService,omitempty"`
		Longitude                 *float64                                                `json:"longitude,omitempty"`
		Members                   ApGroupRetrieveGet1200ResponseMembersSlice              `json:"members,omitempty"`
		Name                      *string                                                 `json:"name,omitempty"`
		VenueProfile              *ApGroupRetrieveGet1200ResponseVenueProfile             `json:"venueProfile,omitempty"`
		Wifi24                    *ApGroupRetrieveGet1200ResponseWifi24                   `json:"wifi24,omitempty"`
		Wifi50                    *ApGroupRetrieveGet1200ResponseWifi50                   `json:"wifi50,omitempty"`
		WlanGroup24               *ApGroupRetrieveGet1200ResponseWlanGroup24              `json:"wlanGroup24,omitempty"`
		WlanGroup50               *ApGroupRetrieveGet1200ResponseWlanGroup50              `json:"wlanGroup50,omitempty"`
		ZoneID                    *string                                                 `json:"zoneId,omitempty"`
	}
)

// ApGroupRetrieveGet1: Use this API command to retrieve information about default AP group of zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApGroupRetrieveGet1200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupRetrieveGet1(ctx context.Context, zoneId string) (*http.Response, *ApGroupRetrieveGet1200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apgroups/default", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ApGroupRetrieveGet1200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// ApGroupDeleteDelete: Use this API command to delete an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupRetrieveGet200ResponseAltitude struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}

	ApGroupRetrieveGet200ResponseApMgmtVlan struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}

	ApGroupRetrieveGet200ResponseAutoChannelSelection24 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	ApGroupRetrieveGet200ResponseAutoChannelSelection50 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	ApGroupRetrieveGet200ResponseClientAdmissionControl24 struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	ApGroupRetrieveGet200ResponseClientAdmissionControl50 struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	ApGroupRetrieveGet200ResponseLocationBasedService struct {
		Enabled *bool   `json:"enabled,omitempty"`
		ID      *string `json:"id,omitempty"`
		Name    *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet200ResponseMembersSlice []*ApGroupRetrieveGet200ResponseMembers

	ApGroupRetrieveGet200ResponseMembers struct {
		ApMac *string `json:"apMac,omitempty"`
	}

	ApGroupRetrieveGet200ResponseVenueProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet200ResponseWifi24AvailableChannelRangeSlice []*float64

	ApGroupRetrieveGet200ResponseWifi24ChannelRangeSlice []*float64

	ApGroupRetrieveGet200ResponseWifi24 struct {
		AvailableChannelRange ApGroupRetrieveGet200ResponseWifi24AvailableChannelRangeSlice `json:"availableChannelRange,omitempty"`
		Channel               *int                                                          `json:"channel,omitempty"`
		ChannelRange          ApGroupRetrieveGet200ResponseWifi24ChannelRangeSlice          `json:"channelRange,omitempty"`
		ChannelWidth          *int                                                          `json:"channelWidth,omitempty"`
		TxPower               *string                                                       `json:"txPower,omitempty"`
	}

	ApGroupRetrieveGet200ResponseWifi50AvailableIndoorChannelRangeSlice []*float64

	ApGroupRetrieveGet200ResponseWifi50AvailableOutdoorChannelRangeSlice []*float64

	ApGroupRetrieveGet200ResponseWifi50IndoorChannelRangeSlice []*float64

	ApGroupRetrieveGet200ResponseWifi50OutdoorChannelRangeSlice []*float64

	ApGroupRetrieveGet200ResponseWifi50 struct {
		AvailableIndoorChannelRange  ApGroupRetrieveGet200ResponseWifi50AvailableIndoorChannelRangeSlice  `json:"availableIndoorChannelRange,omitempty"`
		AvailableOutdoorChannelRange ApGroupRetrieveGet200ResponseWifi50AvailableOutdoorChannelRangeSlice `json:"availableOutdoorChannelRange,omitempty"`
		ChannelWidth                 *int                                                                 `json:"channelWidth,omitempty"`
		IndoorChannel                *float64                                                             `json:"indoorChannel,omitempty"`
		IndoorChannelRange           ApGroupRetrieveGet200ResponseWifi50IndoorChannelRangeSlice           `json:"indoorChannelRange,omitempty"`
		IndoorSecondaryChannel       *float64                                                             `json:"indoorSecondaryChannel,omitempty"`
		OutdoorChannel               *float64                                                             `json:"outdoorChannel,omitempty"`
		OutdoorChannelRange          ApGroupRetrieveGet200ResponseWifi50OutdoorChannelRangeSlice          `json:"outdoorChannelRange,omitempty"`
		OutdoorSecondaryChannel      *float64                                                             `json:"outdoorSecondaryChannel,omitempty"`
		TxPower                      *string                                                              `json:"txPower,omitempty"`
	}

	ApGroupRetrieveGet200ResponseWlanGroup24 struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet200ResponseWlanGroup50 struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveGet200Response struct {
		Altitude                  *ApGroupRetrieveGet200ResponseAltitude                 `json:"altitude,omitempty"`
		ApMgmtVlan                *ApGroupRetrieveGet200ResponseApMgmtVlan               `json:"apMgmtVlan,omitempty"`
		AutoChannelSelection24    *ApGroupRetrieveGet200ResponseAutoChannelSelection24   `json:"autoChannelSelection24,omitempty"`
		AutoChannelSelection50    *ApGroupRetrieveGet200ResponseAutoChannelSelection50   `json:"autoChannelSelection50,omitempty"`
		AwsVenue                  *string                                                `json:"awsVenue,omitempty"`
		ChannelEvaluationInterval *float64                                               `json:"channelEvaluationInterval,omitempty"`
		ClientAdmissionControl24  *ApGroupRetrieveGet200ResponseClientAdmissionControl24 `json:"clientAdmissionControl24,omitempty"`
		ClientAdmissionControl50  *ApGroupRetrieveGet200ResponseClientAdmissionControl50 `json:"clientAdmissionControl50,omitempty"`
		Description               *string                                                `json:"description,omitempty"`
		ID                        *string                                                `json:"id,omitempty"`
		Latitude                  *float64                                               `json:"latitude,omitempty"`
		Location                  *string                                                `json:"location,omitempty"`
		LocationAdditionalInfo    *string                                                `json:"locationAdditionalInfo,omitempty"`
		LocationBasedService      *ApGroupRetrieveGet200ResponseLocationBasedService     `json:"locationBasedService,omitempty"`
		Longitude                 *float64                                               `json:"longitude,omitempty"`
		Members                   ApGroupRetrieveGet200ResponseMembersSlice              `json:"members,omitempty"`
		Name                      *string                                                `json:"name,omitempty"`
		VenueProfile              *ApGroupRetrieveGet200ResponseVenueProfile             `json:"venueProfile,omitempty"`
		Wifi24                    *ApGroupRetrieveGet200ResponseWifi24                   `json:"wifi24,omitempty"`
		Wifi50                    *ApGroupRetrieveGet200ResponseWifi50                   `json:"wifi50,omitempty"`
		WlanGroup24               *ApGroupRetrieveGet200ResponseWlanGroup24              `json:"wlanGroup24,omitempty"`
		WlanGroup50               *ApGroupRetrieveGet200ResponseWlanGroup50              `json:"wlanGroup50,omitempty"`
		ZoneID                    *string                                                `json:"zoneId,omitempty"`
	}
)

// ApGroupRetrieveGet: Use this API command to retrieve information about an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApGroupRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *ApGroupRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apgroups/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &ApGroupRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApGroupModifyBasicPatchRequest struct {
		AwsVenue                  *string  `json:"awsVenue,omitempty"`
		ChannelEvaluationInterval *float64 `json:"channelEvaluationInterval,omitempty"`
		Description               *string  `json:"description,omitempty"`
		Latitude                  *float64 `json:"latitude,omitempty"`
		Location                  *string  `json:"location,omitempty"`
		LocationAdditionalInfo    *string  `json:"locationAdditionalInfo,omitempty"`
		Longitude                 *float64 `json:"longitude,omitempty"`
		Name                      *string  `json:"name,omitempty"`
	}
)

// ApGroupModifyBasicPatch: Use this API command to modify the basic information of an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableAltitudeOfApgroupDelete: Use this API command to clear the altitude of AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableAltitudeOfApgroupDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/altitude", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyAltitudeOfApgroupPatchRequest struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}
)

// ApGroupModifyAltitudeOfApgroupPatch: Use this API command to modify the altitude of AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyAltitudeOfApgroupPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyAltitudeOfApgroupPatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyAltitudeOfApgroupPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/altitude", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableApManagementVlanOverrideDelete: Disable AP Management Vlan Override of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableApManagementVlanOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/apMgmtVlan", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyApManagementVlanPatchRequest struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}
)

// ApGroupModifyApManagementVlanPatch: Modify AP Management Vlan of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyApManagementVlanPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyApManagementVlanPatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyApManagementVlanPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/apMgmtVlan", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableOverrideApModelDelete: Use this API command to disable AP model specific configuration override zone that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - model (string): Access Point Model
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableOverrideApModelDelete(ctx context.Context, zoneId string, id string, model string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(model)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"model\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/apmodel/{model}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("model", model)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupRetrieveApModelGet200ResponseExternalAntenna24 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	ApGroupRetrieveApModelGet200ResponseExternalAntenna50 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	ApGroupRetrieveApModelGet200ResponseLanPortsSlice []*ApGroupRetrieveApModelGet200ResponseLanPorts

	ApGroupRetrieveApModelGet200ResponseLanPortsEthPortProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupRetrieveApModelGet200ResponseLanPorts struct {
		Enabled        *bool                                                       `json:"enabled,omitempty"`
		EthPortProfile *ApGroupRetrieveApModelGet200ResponseLanPortsEthPortProfile `json:"ethPortProfile,omitempty"`
		PortName       *string                                                     `json:"portName,omitempty"`
	}

	ApGroupRetrieveApModelGet200ResponseLldp struct {
		AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
		Enabled                *bool `json:"enabled,omitempty"`
		HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
		ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
	}

	ApGroupRetrieveApModelGet200Response struct {
		ExternalAntenna24     *ApGroupRetrieveApModelGet200ResponseExternalAntenna24 `json:"externalAntenna24,omitempty"`
		ExternalAntenna50     *ApGroupRetrieveApModelGet200ResponseExternalAntenna50 `json:"externalAntenna50,omitempty"`
		InternalHeaterEnabled *bool                                                  `json:"internalHeaterEnabled,omitempty"`
		LanPorts              ApGroupRetrieveApModelGet200ResponseLanPortsSlice      `json:"lanPorts,omitempty"`
		LedMode               interface{}                                            `json:"ledMode,omitempty"`
		LedStatusEnabled      *bool                                                  `json:"ledStatusEnabled,omitempty"`
		Lldp                  *ApGroupRetrieveApModelGet200ResponseLldp              `json:"lldp,omitempty"`
		PoeModeSetting        *string                                                `json:"poeModeSetting,omitempty"`
		PoeOutPortEnabled     *bool                                                  `json:"poeOutPortEnabled,omitempty"`
		PoeTxChain            *int                                                   `json:"poeTxChain,omitempty"`
		RadioBand             interface{}                                            `json:"radioBand,omitempty"`
		UsbPowerEnable        *bool                                                  `json:"usbPowerEnable,omitempty"`
	}
)

// ApGroupRetrieveApModelGet: Use this API command to retrieve AP model specific configuration override zone that belong to an AP group, empty mean not override zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - model (string): Access Point Model
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApGroupRetrieveApModelGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupRetrieveApModelGet(ctx context.Context, zoneId string, id string, model string) (*http.Response, *ApGroupRetrieveApModelGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(model)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"model\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apgroups/{id}/apmodel/{model}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("model", model)
	out := &ApGroupRetrieveApModelGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApGroupOverrideApModelPutRequestExternalAntenna24 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	ApGroupOverrideApModelPutRequestExternalAntenna50 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	ApGroupOverrideApModelPutRequestLanPortsSlice []*ApGroupOverrideApModelPutRequestLanPorts

	ApGroupOverrideApModelPutRequestLanPortsEthPortProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupOverrideApModelPutRequestLanPorts struct {
		Enabled        *bool                                                   `json:"enabled,omitempty"`
		EthPortProfile *ApGroupOverrideApModelPutRequestLanPortsEthPortProfile `json:"ethPortProfile,omitempty"`
		PortName       *string                                                 `json:"portName,omitempty"`
	}

	ApGroupOverrideApModelPutRequestLldp struct {
		AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
		Enabled                *bool `json:"enabled,omitempty"`
		HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
		ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
	}

	ApGroupOverrideApModelPutRequest struct {
		ExternalAntenna24     *ApGroupOverrideApModelPutRequestExternalAntenna24 `json:"externalAntenna24,omitempty"`
		ExternalAntenna50     *ApGroupOverrideApModelPutRequestExternalAntenna50 `json:"externalAntenna50,omitempty"`
		InternalHeaterEnabled *bool                                              `json:"internalHeaterEnabled,omitempty"`
		LanPorts              ApGroupOverrideApModelPutRequestLanPortsSlice      `json:"lanPorts,omitempty"`
		LedMode               interface{}                                        `json:"ledMode,omitempty"`
		LedStatusEnabled      *bool                                              `json:"ledStatusEnabled,omitempty"`
		Lldp                  *ApGroupOverrideApModelPutRequestLldp              `json:"lldp,omitempty"`
		PoeModeSetting        *string                                            `json:"poeModeSetting,omitempty"`
		PoeOutPortEnabled     *bool                                              `json:"poeOutPortEnabled,omitempty"`
		PoeTxChain            *int                                               `json:"poeTxChain,omitempty"`
		RadioBand             interface{}                                        `json:"radioBand,omitempty"`
		UsbPowerEnable        *bool                                              `json:"usbPowerEnable,omitempty"`
	}
)

// ApGroupOverrideApModelPut: Use this API command to modify AP model specific configuration override zone that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - model (string): Access Point Model
// - requestBody: *ApGroupOverrideApModelPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupOverrideApModelPut(ctx context.Context, zoneId string, id string, model string, requestBody *ApGroupOverrideApModelPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(model)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"model\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/rkszones/{zoneId}/apgroups/{id}/apmodel/{model}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("model", model)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio24gAutoChannelselectmodeOverrideDelete: Disable Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio24gAutoChannelselectmodeOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/autoChannelSelection24", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyRadio24gAutoChannelselectmodeOverridePatchRequest struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}
)

// ApGroupModifyRadio24gAutoChannelselectmodeOverridePatch: Override Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyRadio24gAutoChannelselectmodeOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyRadio24gAutoChannelselectmodeOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyRadio24gAutoChannelselectmodeOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/autoChannelSelection24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gAutoChannelselectmodeOverrideDelete: Disable Radio 5G Auto ChannelSelectMode and ChannelFly MTBC Override of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gAutoChannelselectmodeOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/autoChannelSelection50", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyRadio5gAutoChannelselectmodeOverridePatchRequest struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}
)

// ApGroupModifyRadio5gAutoChannelselectmodeOverridePatch: Override Radio 5G Auto ChannelSelectMode and ChannelFly MTBC of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyRadio5gAutoChannelselectmodeOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyRadio5gAutoChannelselectmodeOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyRadio5gAutoChannelselectmodeOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/autoChannelSelection50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableChannelEvaluationIntervalOverrideDelete: Disable Channel Evaluation Interval Override of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableChannelEvaluationIntervalOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/channelEvaluationInterval", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableClientAdmissionControl24gOverrideDelete: Use this API command to disable client admission control 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableClientAdmissionControl24gOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/clientAdmissionControl24", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyClientAdmissionControl24gOverridePatchRequest struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}
)

// ApGroupModifyClientAdmissionControl24gOverridePatch: Use this API command to modify client admission control 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyClientAdmissionControl24gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyClientAdmissionControl24gOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyClientAdmissionControl24gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/clientAdmissionControl24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableClientAdmissionControl5gOverrideDelete: Use this API command to disable client admission control 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableClientAdmissionControl5gOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/clientAdmissionControl50", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyClientAdmissionControl5gOverridePatchRequest struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}
)

// ApGroupModifyClientAdmissionControl5gOverridePatch: Use this API command to modify client admission control 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyClientAdmissionControl5gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyClientAdmissionControl5gOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyClientAdmissionControl5gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/clientAdmissionControl50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableLocationOverrideDelete: Use this API command to disable location override for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableLocationOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/location", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableLocationAdditionalinfoOverrideDelete: Use this API command to disable location additionalInfo override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableLocationAdditionalinfoOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/locationAdditionalInfo", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableLocationBasedServiceOverrideDelete: Use this API command to disable location based service override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableLocationBasedServiceOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/locationBasedService", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyLocationBasedServiceOverridePatchRequest struct {
		Enabled *bool   `json:"enabled,omitempty"`
		ID      *string `json:"id,omitempty"`
		Name    *string `json:"name,omitempty"`
	}
)

// ApGroupModifyLocationBasedServiceOverridePatch: Use this API command to modify location based service override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyLocationBasedServiceOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyLocationBasedServiceOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyLocationBasedServiceOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/locationBasedService", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupAddMemberListPostRequestMemberListSlice []*ApGroupAddMemberListPostRequestMemberList

	ApGroupAddMemberListPostRequestMemberList struct {
		ApMac *string `json:"apMac,omitempty"`
	}

	ApGroupAddMemberListPostRequest struct {
		MemberList ApGroupAddMemberListPostRequestMemberListSlice `json:"memberList,omitempty"`
	}
)

// ApGroupAddMemberListPost: Add multiple members to an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupAddMemberListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupAddMemberListPost(ctx context.Context, zoneId string, id string, requestBody *ApGroupAddMemberListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/apgroups/{id}/members", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 201, nil)
}

// ApGroupRemoveMemberDelete: Use this API command to remove a member AP from an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupRemoveMemberDelete(ctx context.Context, zoneId string, id string, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/members/{apMac}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("apMac", apMac)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupAddMemberPost: Use this API command to add a member AP to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupAddMemberPost(ctx context.Context, zoneId string, id string, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/apgroups/{id}/members/{apMac}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("apMac", apMac)
	return r.client.Ensure(ctx, request, 201, nil)
}

// ApGroupDisableApUsbSoftwarePackageDelete: Disable AP Usb Software Package of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Optional Parameter Map:
// - applyModel (string): Specify ap models. ex : applyModel=ZF7321U&applyModel=ZF7323
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableApUsbSoftwarePackageDelete(ctx context.Context, zoneId string, id string, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/usbSoftwarePackage", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetQueryParameter("applyModel", optionalParams["applyModel"])
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyApUsbSoftwarePackagePatchRequestApplyModelSlice []*string

	ApGroupModifyApUsbSoftwarePackagePatchRequestUsbSoftware struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApGroupModifyApUsbSoftwarePackagePatchRequest struct {
		ApplyModel  ApGroupModifyApUsbSoftwarePackagePatchRequestApplyModelSlice `json:"applyModel,omitempty"`
		UsbSoftware *ApGroupModifyApUsbSoftwarePackagePatchRequestUsbSoftware    `json:"usbSoftware,omitempty"`
	}
)

// ApGroupModifyApUsbSoftwarePackagePatch: Modify AP Usb Software Package of an AP group
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyApUsbSoftwarePackagePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyApUsbSoftwarePackagePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyApUsbSoftwarePackagePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/usbSoftwarePackage", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApUsbSoftwarePackageGetApgroupAssociateGet200Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApUsbSoftwarePackageGetApgroupAssociateGet: Get APUsbSoftwarePackage associate with APGroup by model name
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - modelName (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApUsbSoftwarePackageGetApgroupAssociateGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApUsbSoftwarePackageGetApgroupAssociateGet(ctx context.Context, zoneId string, id string, modelName string) (*http.Response, *ApUsbSoftwarePackageGetApgroupAssociateGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(modelName)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"modelName\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apgroups/{id}/usbsoftware/{modelName}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("modelName", modelName)
	out := &ApUsbSoftwarePackageGetApgroupAssociateGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// ApGroupClearHotspot20VenueProfileDelete: Use this API command to clear Hotspot 2.0 venue profile for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupClearHotspot20VenueProfileDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/venueProfile", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyHotspot20VenueProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApGroupModifyHotspot20VenueProfilePatch: Use this API command to  modify Hotspot 2.0 venue profile for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyHotspot20VenueProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyHotspot20VenueProfilePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyHotspot20VenueProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/venueProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio24gOverrideDelete: Use this API command to disable 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio24gOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi24", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyRadio24gOverridePatchRequestChannelRangeSlice []*float64

	ApGroupModifyRadio24gOverridePatchRequest struct {
		Channel      *int                                                       `json:"channel,omitempty"`
		ChannelRange ApGroupModifyRadio24gOverridePatchRequestChannelRangeSlice `json:"channelRange,omitempty"`
		ChannelWidth *int                                                       `json:"channelWidth,omitempty"`
		TxPower      *string                                                    `json:"txPower,omitempty"`
	}
)

// ApGroupModifyRadio24gOverridePatch: Use this API command to modify the 2.4GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyRadio24gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyRadio24gOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyRadio24gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio24gChannelOverrideDelete: Use this API command to disable 2.4GHz radio channel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio24gChannelOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi24/channel", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio24gChannelrangeOverrideDelete: Use this API command to disable 2.4GHz radio channelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio24gChannelrangeOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi24/channelRange", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio24gChannelwidthOverrideDelete: Use this API command to disable 2.4GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio24gChannelwidthOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi24/channelWidth", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio24gTxpowerOverrideDelete: Use this API command to disable 2.4GHz radio txPower override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio24gTxpowerOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi24/txPower", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gOverrideDelete: Use this API command to disable 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyRadio5gOverridePatchRequestIndoorChannelRangeSlice []*float64

	ApGroupModifyRadio5gOverridePatchRequestOutdoorChannelRangeSlice []*float64

	ApGroupModifyRadio5gOverridePatchRequest struct {
		ChannelWidth            *int                                                             `json:"channelWidth,omitempty"`
		IndoorChannel           *float64                                                         `json:"indoorChannel,omitempty"`
		IndoorChannelRange      ApGroupModifyRadio5gOverridePatchRequestIndoorChannelRangeSlice  `json:"indoorChannelRange,omitempty"`
		IndoorSecondaryChannel  *float64                                                         `json:"indoorSecondaryChannel,omitempty"`
		OutdoorChannel          *float64                                                         `json:"outdoorChannel,omitempty"`
		OutdoorChannelRange     ApGroupModifyRadio5gOverridePatchRequestOutdoorChannelRangeSlice `json:"outdoorChannelRange,omitempty"`
		OutdoorSecondaryChannel *float64                                                         `json:"outdoorSecondaryChannel,omitempty"`
		TxPower                 *string                                                          `json:"txPower,omitempty"`
	}
)

// ApGroupModifyRadio5gOverridePatch: Use this API command to modify the 5GHz radio configuration override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyRadio5gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyRadio5gOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyRadio5gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gChannelwidthOverrideDelete: Use this API command to disable 5GHz radio channelWidth override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gChannelwidthOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50/channelWidth", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gIndoorchannelOverrideDelete: Use this API command to disable 5GHz radio indoorChannel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gIndoorchannelOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50/indoorChannel", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gIndoorchannelrangeOverrideDelete: Use this API command to disable 5GHz radio indoorChannelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gIndoorchannelrangeOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50/indoorChannelRange", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gOutdoorchannelOverrideDelete: Use this API command to disable 5GHz radio outdoorChannel override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gOutdoorchannelOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50/outdoorChannel", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gOutdoorchannelrangeOverrideDelete: Use this API command to disable 5GHz radio outdoorChannelRange override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gOutdoorchannelrangeOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50/outdoorChannelRange", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableRadio5gTxpowerOverrideDelete: Use this API command to disable 5GHz radio txPower override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableRadio5gTxpowerOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wifi50/txPower", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableWlanGroup24gOverrideDelete: Use this API command to disable WLAN group on 2.4GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableWlanGroup24gOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wlanGroup24", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyWlanGroup24gOverridePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApGroupModifyWlanGroup24gOverridePatch: Use this API command to modify the WLAN group on 2.4GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyWlanGroup24gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyWlanGroup24gOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyWlanGroup24gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wlanGroup24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// ApGroupDisableWlanGroup5gOverrideDelete: Use this API command to disable WLAN group on 5GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupDisableWlanGroup5gOverrideDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wlanGroup50", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApGroupModifyWlanGroup5gOverridePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApGroupModifyWlanGroup5gOverridePatch: Use this API command to modify the WLAN group on 5GHz radio override zone for APs that belong to an AP group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): AP Group ID
// - requestBody: *ApGroupModifyWlanGroup5gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApGroupModifyWlanGroup5gOverridePatch(ctx context.Context, zoneId string, id string, requestBody *ApGroupModifyWlanGroup5gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/apgroups/{id}/wlanGroup50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	RuckusWirelessApZoneRetrieveApModelGet200ResponseExternalAntenna24 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApModelGet200ResponseExternalAntenna50 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApModelGet200ResponseLanPortsSlice []*RuckusWirelessApZoneRetrieveApModelGet200ResponseLanPorts

	RuckusWirelessApZoneRetrieveApModelGet200ResponseLanPortsEthPortProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApModelGet200ResponseLanPorts struct {
		Enabled        *bool                                                                    `json:"enabled,omitempty"`
		EthPortProfile *RuckusWirelessApZoneRetrieveApModelGet200ResponseLanPortsEthPortProfile `json:"ethPortProfile,omitempty"`
		PortName       *string                                                                  `json:"portName,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApModelGet200ResponseLldp struct {
		AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
		Enabled                *bool `json:"enabled,omitempty"`
		HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
		ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
	}

	RuckusWirelessApZoneRetrieveApModelGet200Response struct {
		ExternalAntenna24     *RuckusWirelessApZoneRetrieveApModelGet200ResponseExternalAntenna24 `json:"externalAntenna24,omitempty"`
		ExternalAntenna50     *RuckusWirelessApZoneRetrieveApModelGet200ResponseExternalAntenna50 `json:"externalAntenna50,omitempty"`
		InternalHeaterEnabled *bool                                                               `json:"internalHeaterEnabled,omitempty"`
		LanPorts              RuckusWirelessApZoneRetrieveApModelGet200ResponseLanPortsSlice      `json:"lanPorts,omitempty"`
		LedMode               interface{}                                                         `json:"ledMode,omitempty"`
		LedStatusEnabled      *bool                                                               `json:"ledStatusEnabled,omitempty"`
		Lldp                  *RuckusWirelessApZoneRetrieveApModelGet200ResponseLldp              `json:"lldp,omitempty"`
		PoeModeSetting        *string                                                             `json:"poeModeSetting,omitempty"`
		PoeOutPortEnabled     *bool                                                               `json:"poeOutPortEnabled,omitempty"`
		PoeTxChain            *int                                                                `json:"poeTxChain,omitempty"`
		RadioBand             interface{}                                                         `json:"radioBand,omitempty"`
		UsbPowerEnable        *bool                                                               `json:"usbPowerEnable,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveApModelGet: Use this API command to retrieve AP model specific configuration that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - model (string): Access Point Model
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveApModelGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveApModelGet(ctx context.Context, zoneId string, model string) (*http.Response, *RuckusWirelessApZoneRetrieveApModelGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(model)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"model\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/apmodel/{model}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("model", model)
	out := &RuckusWirelessApZoneRetrieveApModelGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneModifyApModelPutRequestExternalAntenna24 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	RuckusWirelessApZoneModifyApModelPutRequestExternalAntenna50 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	RuckusWirelessApZoneModifyApModelPutRequestLanPortsSlice []*RuckusWirelessApZoneModifyApModelPutRequestLanPorts

	RuckusWirelessApZoneModifyApModelPutRequestLanPortsEthPortProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	RuckusWirelessApZoneModifyApModelPutRequestLanPorts struct {
		Enabled        *bool                                                              `json:"enabled,omitempty"`
		EthPortProfile *RuckusWirelessApZoneModifyApModelPutRequestLanPortsEthPortProfile `json:"ethPortProfile,omitempty"`
		PortName       *string                                                            `json:"portName,omitempty"`
	}

	RuckusWirelessApZoneModifyApModelPutRequestLldp struct {
		AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
		Enabled                *bool `json:"enabled,omitempty"`
		HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
		ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
	}

	RuckusWirelessApZoneModifyApModelPutRequest struct {
		ExternalAntenna24     *RuckusWirelessApZoneModifyApModelPutRequestExternalAntenna24 `json:"externalAntenna24,omitempty"`
		ExternalAntenna50     *RuckusWirelessApZoneModifyApModelPutRequestExternalAntenna50 `json:"externalAntenna50,omitempty"`
		InternalHeaterEnabled *bool                                                         `json:"internalHeaterEnabled,omitempty"`
		LanPorts              RuckusWirelessApZoneModifyApModelPutRequestLanPortsSlice      `json:"lanPorts,omitempty"`
		LedMode               interface{}                                                   `json:"ledMode,omitempty"`
		LedStatusEnabled      *bool                                                         `json:"ledStatusEnabled,omitempty"`
		Lldp                  *RuckusWirelessApZoneModifyApModelPutRequestLldp              `json:"lldp,omitempty"`
		PoeModeSetting        *string                                                       `json:"poeModeSetting,omitempty"`
		PoeOutPortEnabled     *bool                                                         `json:"poeOutPortEnabled,omitempty"`
		PoeTxChain            *int                                                          `json:"poeTxChain,omitempty"`
		RadioBand             interface{}                                                   `json:"radioBand,omitempty"`
		UsbPowerEnable        *bool                                                         `json:"usbPowerEnable,omitempty"`
	}
)

// RuckusWirelessApZoneModifyApModelPut: Use this API command to modify the AP model specific configuration that belong to a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - model (string): Access Point Model
// - requestBody: *RuckusWirelessApZoneModifyApModelPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneModifyApModelPut(ctx context.Context, zoneId string, model string, requestBody *RuckusWirelessApZoneModifyApModelPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(model)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"model\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/rkszones/{zoneId}/apmodel/{model}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("model", model)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourFencingPolicyRetrieveListGet200ResponseListSlice []*BonjourFencingPolicyRetrieveListGet200ResponseList

	BonjourFencingPolicyRetrieveListGet200ResponseListBonjourFencingRuleListSlice []*BonjourFencingPolicyRetrieveListGet200ResponseListBonjourFencingRuleList

	BonjourFencingPolicyRetrieveListGet200ResponseListBonjourFencingRuleList struct {
		ClosestAp    *string `json:"closestAp,omitempty"`
		Description  *string `json:"description,omitempty"`
		DeviceMac    *string `json:"deviceMac,omitempty"`
		DeviceType   *string `json:"deviceType,omitempty"`
		FencingRange *string `json:"fencingRange,omitempty"`
		ServiceType  *string `json:"serviceType,omitempty"`
	}

	BonjourFencingPolicyRetrieveListGet200ResponseList struct {
		BonjourFencingRuleList BonjourFencingPolicyRetrieveListGet200ResponseListBonjourFencingRuleListSlice `json:"bonjourFencingRuleList,omitempty"`
		CreateDateTime         *int                                                                          `json:"createDateTime,omitempty"`
		CreatorID              *string                                                                       `json:"creatorId,omitempty"`
		CreatorUsername        *string                                                                       `json:"creatorUsername,omitempty"`
		Description            *string                                                                       `json:"description,omitempty"`
		ID                     *string                                                                       `json:"id,omitempty"`
		ModifiedDateTime       *int                                                                          `json:"modifiedDateTime,omitempty"`
		ModifierID             *string                                                                       `json:"modifierId,omitempty"`
		ModifierUsername       *string                                                                       `json:"modifierUsername,omitempty"`
		Name                   *string                                                                       `json:"name,omitempty"`
		ZoneID                 *string                                                                       `json:"zoneId,omitempty"`
	}

	BonjourFencingPolicyRetrieveListGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       BonjourFencingPolicyRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// BonjourFencingPolicyRetrieveListGet: Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourFencingPolicyRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *BonjourFencingPolicyRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/bonjourFencingPolicy", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &BonjourFencingPolicyRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	BonjourFencingPolicyCreatePostRequestBonjourFencingRuleListSlice []*BonjourFencingPolicyCreatePostRequestBonjourFencingRuleList

	BonjourFencingPolicyCreatePostRequestBonjourFencingRuleList struct {
		ClosestAp    *string `json:"closestAp,omitempty"`
		Description  *string `json:"description,omitempty"`
		DeviceMac    *string `json:"deviceMac,omitempty"`
		DeviceType   *string `json:"deviceType,omitempty"`
		FencingRange *string `json:"fencingRange,omitempty"`
		ServiceType  *string `json:"serviceType,omitempty"`
	}

	BonjourFencingPolicyCreatePostRequest struct {
		BonjourFencingRuleList BonjourFencingPolicyCreatePostRequestBonjourFencingRuleListSlice `json:"bonjourFencingRuleList,omitempty"`
		Description            *string                                                          `json:"description,omitempty"`
		Name                   *string                                                          `json:"name,omitempty"`
	}

	BonjourFencingPolicyCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// BonjourFencingPolicyCreatePost: Use this API command to create Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *BonjourFencingPolicyCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourFencingPolicyCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyCreatePost(ctx context.Context, zoneId string, requestBody *BonjourFencingPolicyCreatePostRequest) (*http.Response, *BonjourFencingPolicyCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/bonjourFencingPolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &BonjourFencingPolicyCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	BonjourFencingPolicyRetrieveGet200ResponseBonjourFencingRuleListSlice []*BonjourFencingPolicyRetrieveGet200ResponseBonjourFencingRuleList

	BonjourFencingPolicyRetrieveGet200ResponseBonjourFencingRuleList struct {
		ClosestAp    *string `json:"closestAp,omitempty"`
		Description  *string `json:"description,omitempty"`
		DeviceMac    *string `json:"deviceMac,omitempty"`
		DeviceType   *string `json:"deviceType,omitempty"`
		FencingRange *string `json:"fencingRange,omitempty"`
		ServiceType  *string `json:"serviceType,omitempty"`
	}

	BonjourFencingPolicyRetrieveGet200Response struct {
		BonjourFencingRuleList BonjourFencingPolicyRetrieveGet200ResponseBonjourFencingRuleListSlice `json:"bonjourFencingRuleList,omitempty"`
		CreateDateTime         *int                                                                  `json:"createDateTime,omitempty"`
		CreatorID              *string                                                               `json:"creatorId,omitempty"`
		CreatorUsername        *string                                                               `json:"creatorUsername,omitempty"`
		Description            *string                                                               `json:"description,omitempty"`
		ID                     *string                                                               `json:"id,omitempty"`
		ModifiedDateTime       *int                                                                  `json:"modifiedDateTime,omitempty"`
		ModifierID             *string                                                               `json:"modifierId,omitempty"`
		ModifierUsername       *string                                                               `json:"modifierUsername,omitempty"`
		Name                   *string                                                               `json:"name,omitempty"`
		ZoneID                 *string                                                               `json:"zoneId,omitempty"`
	}
)

// BonjourFencingPolicyRetrieveGet: Use this API command to retrieve Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourFencingPolicyRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *BonjourFencingPolicyRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/bonjourFencingPolicy/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &BonjourFencingPolicyRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	BonjourFencingPolicyModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// BonjourFencingPolicyModifyBasicPatch: Use this API command to modify the basic information of Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *BonjourFencingPolicyModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *BonjourFencingPolicyModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/bonjourFencingPolicy/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourFencingPolicyModifyBonjourFencingRuleListPatchRequestSlice []*BonjourFencingPolicyModifyBonjourFencingRuleListPatchRequest

	BonjourFencingPolicyModifyBonjourFencingRuleListPatchRequest struct {
		ClosestAp    *string `json:"closestAp,omitempty"`
		Description  *string `json:"description,omitempty"`
		DeviceMac    *string `json:"deviceMac,omitempty"`
		DeviceType   *string `json:"deviceType,omitempty"`
		FencingRange *string `json:"fencingRange,omitempty"`
		ServiceType  *string `json:"serviceType,omitempty"`
	}
)

// BonjourFencingPolicyModifyBonjourFencingRuleListPatch: Use this API command to modify Bonjour Fencing Rule List.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *BonjourFencingPolicyModifyBonjourFencingRuleListPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourFencingPolicyModifyBonjourFencingRuleListPatch(ctx context.Context, zoneId string, id string, requestBody BonjourFencingPolicyModifyBonjourFencingRuleListPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/bonjourFencingPolicy/{id}/bonjourFencingRuleList", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourGatewayPoliciesModifyEnablePatchRequest struct {
		EnabledBonjourGateway *bool `json:"enabledBonjourGateway,omitempty"`
	}
)

// BonjourGatewayPoliciesModifyEnablePatch: Use this API command to enable/disable bonjour gateway policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *BonjourGatewayPoliciesModifyEnablePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesModifyEnablePatch(ctx context.Context, zoneId string, requestBody *BonjourGatewayPoliciesModifyEnablePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/bounjourGateway/enable", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourGatewayPoliciesRetrieveListGet200ResponseListSlice []*BonjourGatewayPoliciesRetrieveListGet200ResponseList

	BonjourGatewayPoliciesRetrieveListGet200ResponseList struct {
		Description    *string `json:"description,omitempty"`
		ID             *string `json:"id,omitempty"`
		LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
		LastModifiedOn *string `json:"lastModifiedOn,omitempty"`
		Name           *string `json:"name,omitempty"`
	}

	BonjourGatewayPoliciesRetrieveListGet200Response struct {
		FirstIndex *int                                                      `json:"firstIndex,omitempty"`
		HasMore    *bool                                                     `json:"hasMore,omitempty"`
		List       BonjourGatewayPoliciesRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                      `json:"totalCount,omitempty"`
	}
)

// BonjourGatewayPoliciesRetrieveListGet: Use this API command to retrieve a list of bonjour gateway policies.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourGatewayPoliciesRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *BonjourGatewayPoliciesRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/bounjourGateway/policies", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &BonjourGatewayPoliciesRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	BonjourGatewayPoliciesCreatePostRequestBonjourPolicyRuleListSlice []*BonjourGatewayPoliciesCreatePostRequestBonjourPolicyRuleList

	BonjourGatewayPoliciesCreatePostRequestBonjourPolicyRuleList struct {
		BridgeService *string  `json:"bridgeService,omitempty"`
		FromVlan      *float64 `json:"fromVlan,omitempty"`
		Notes         *string  `json:"notes,omitempty"`
		Protocol      *string  `json:"protocol,omitempty"`
		ToVlan        *float64 `json:"toVlan,omitempty"`
	}

	BonjourGatewayPoliciesCreatePostRequest struct {
		BonjourPolicyRuleList BonjourGatewayPoliciesCreatePostRequestBonjourPolicyRuleListSlice `json:"bonjourPolicyRuleList,omitempty"`
		Description           *string                                                           `json:"description,omitempty"`
		Name                  *string                                                           `json:"name,omitempty"`
	}

	BonjourGatewayPoliciesCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// BonjourGatewayPoliciesCreatePost: Use this API command to create bonjour gateway policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *BonjourGatewayPoliciesCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourGatewayPoliciesCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesCreatePost(ctx context.Context, zoneId string, requestBody *BonjourGatewayPoliciesCreatePostRequest) (*http.Response, *BonjourGatewayPoliciesCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/bounjourGateway/policies", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &BonjourGatewayPoliciesCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// BonjourGatewayPoliciesDeleteDelete: Use this API command to delete bonjour gateway policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/bounjourGateway/policies/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourGatewayPoliciesRetrieveGet200ResponseBonjourPolicyRuleListSlice []*BonjourGatewayPoliciesRetrieveGet200ResponseBonjourPolicyRuleList

	BonjourGatewayPoliciesRetrieveGet200ResponseBonjourPolicyRuleList struct {
		BridgeService *string  `json:"bridgeService,omitempty"`
		FromVlan      *float64 `json:"fromVlan,omitempty"`
		Notes         *string  `json:"notes,omitempty"`
		Priority      *float64 `json:"priority,omitempty"`
		Protocol      *string  `json:"protocol,omitempty"`
		ToVlan        *float64 `json:"toVlan,omitempty"`
	}

	BonjourGatewayPoliciesRetrieveGet200Response struct {
		BonjourPolicyRuleList BonjourGatewayPoliciesRetrieveGet200ResponseBonjourPolicyRuleListSlice `json:"bonjourPolicyRuleList,omitempty"`
		Description           *string                                                                `json:"description,omitempty"`
		Name                  *string                                                                `json:"name,omitempty"`
	}
)

// BonjourGatewayPoliciesRetrieveGet: Use this API command to retrieve bonjour gateway policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourGatewayPoliciesRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *BonjourGatewayPoliciesRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/bounjourGateway/policies/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &BonjourGatewayPoliciesRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	BonjourGatewayPoliciesModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// BonjourGatewayPoliciesModifyBasicPatch: Use this API command to modify the basic information of bonjour gateway policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *BonjourGatewayPoliciesModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *BonjourGatewayPoliciesModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/bounjourGateway/policies/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatchRequestSlice []*BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatchRequest

	BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatchRequest struct {
		BridgeService *string  `json:"bridgeService,omitempty"`
		FromVlan      *float64 `json:"fromVlan,omitempty"`
		Notes         *string  `json:"notes,omitempty"`
		Protocol      *string  `json:"protocol,omitempty"`
		ToVlan        *float64 `json:"toVlan,omitempty"`
	}
)

// BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatch: Use this API command to modify bonjour gateway policy rules.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatch(ctx context.Context, zoneId string, id string, requestBody BonjourGatewayPoliciesModifyBonjourGatewayPolicyRuleListPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/bounjourGateway/policies/{id}/bonjourPolicyRuleList", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ClientIsolationWhitelistRetrieveListGet200ResponseListSlice []*ClientIsolationWhitelistRetrieveListGet200ResponseList

	ClientIsolationWhitelistRetrieveListGet200ResponseListWhitelistSlice []*ClientIsolationWhitelistRetrieveListGet200ResponseListWhitelist

	ClientIsolationWhitelistRetrieveListGet200ResponseListWhitelist struct {
		Description *string `json:"description,omitempty"`
		IP          *string `json:"ip,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListGet200ResponseList struct {
		ClientIsolationAutoEnabled *bool                                                                `json:"clientIsolationAutoEnabled,omitempty"`
		CreateDateTime             *int                                                                 `json:"createDateTime,omitempty"`
		CreatorID                  *string                                                              `json:"creatorId,omitempty"`
		CreatorUsername            *string                                                              `json:"creatorUsername,omitempty"`
		Description                *string                                                              `json:"description,omitempty"`
		ID                         *string                                                              `json:"id,omitempty"`
		ModifiedDateTime           *int                                                                 `json:"modifiedDateTime,omitempty"`
		ModifierID                 *string                                                              `json:"modifierId,omitempty"`
		ModifierUsername           *string                                                              `json:"modifierUsername,omitempty"`
		Name                       *string                                                              `json:"name,omitempty"`
		Whitelist                  ClientIsolationWhitelistRetrieveListGet200ResponseListWhitelistSlice `json:"whitelist,omitempty"`
		ZoneID                     *string                                                              `json:"zoneId,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListGet200Response struct {
		CreateDateTime   *int                                                        `json:"createDateTime,omitempty"`
		CreatorID        *string                                                     `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                     `json:"creatorUsername,omitempty"`
		FirstIndex       *int                                                        `json:"firstIndex,omitempty"`
		HasMore          *bool                                                       `json:"hasMore,omitempty"`
		List             ClientIsolationWhitelistRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                                        `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                     `json:"modifierId,omitempty"`
		ModifierUsername *string                                                     `json:"modifierUsername,omitempty"`
		TotalCount       *int                                                        `json:"totalCount,omitempty"`
	}
)

// ClientIsolationWhitelistRetrieveListGet: Retrieve a list of Client Isolation Whitelist
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ClientIsolationWhitelistRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *ClientIsolationWhitelistRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/clientIsolationWhitelist", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ClientIsolationWhitelistRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequestWhitelistSlice []*ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequestWhitelist

	ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequestWhitelist struct {
		Description *string `json:"description,omitempty"`
		IP          *string `json:"ip,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}

	ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequest struct {
		ClientIsolationAutoEnabled *bool                                                                           `json:"clientIsolationAutoEnabled,omitempty"`
		Description                *string                                                                         `json:"description,omitempty"`
		Name                       *string                                                                         `json:"name,omitempty"`
		Whitelist                  ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequestWhitelistSlice `json:"whitelist,omitempty"`
	}

	ClientIsolationWhitelistCreateClientIsolationWhitelistPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ClientIsolationWhitelistCreateClientIsolationWhitelistPost: Create a new ClientIsolationWhitelist
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ClientIsolationWhitelistCreateClientIsolationWhitelistPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistCreateClientIsolationWhitelistPost(ctx context.Context, zoneId string, requestBody *ClientIsolationWhitelistCreateClientIsolationWhitelistPostRequest) (*http.Response, *ClientIsolationWhitelistCreateClientIsolationWhitelistPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/clientIsolationWhitelist", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &ClientIsolationWhitelistCreateClientIsolationWhitelistPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	ClientIsolationWhitelistRetrieveGet200ResponseWhitelistSlice []*ClientIsolationWhitelistRetrieveGet200ResponseWhitelist

	ClientIsolationWhitelistRetrieveGet200ResponseWhitelist struct {
		Description *string `json:"description,omitempty"`
		IP          *string `json:"ip,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}

	ClientIsolationWhitelistRetrieveGet200Response struct {
		ClientIsolationAutoEnabled *bool                                                        `json:"clientIsolationAutoEnabled,omitempty"`
		CreateDateTime             *int                                                         `json:"createDateTime,omitempty"`
		CreatorID                  *string                                                      `json:"creatorId,omitempty"`
		CreatorUsername            *string                                                      `json:"creatorUsername,omitempty"`
		Description                *string                                                      `json:"description,omitempty"`
		ID                         *string                                                      `json:"id,omitempty"`
		ModifiedDateTime           *int                                                         `json:"modifiedDateTime,omitempty"`
		ModifierID                 *string                                                      `json:"modifierId,omitempty"`
		ModifierUsername           *string                                                      `json:"modifierUsername,omitempty"`
		Name                       *string                                                      `json:"name,omitempty"`
		Whitelist                  ClientIsolationWhitelistRetrieveGet200ResponseWhitelistSlice `json:"whitelist,omitempty"`
		ZoneID                     *string                                                      `json:"zoneId,omitempty"`
	}
)

// ClientIsolationWhitelistRetrieveGet: Retrieve an Client Isolation Whitelist
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ClientIsolationWhitelistRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *ClientIsolationWhitelistRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/clientIsolationWhitelist/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &ClientIsolationWhitelistRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ClientIsolationWhitelistModifyBasicPatchRequest struct {
		ClientIsolationAutoEnabled *bool   `json:"clientIsolationAutoEnabled,omitempty"`
		Description                *string `json:"description,omitempty"`
		Name                       *string `json:"name,omitempty"`
	}
)

// ClientIsolationWhitelistModifyBasicPatch: Modify a specific Client Isolation Whitelist basic
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ClientIsolationWhitelistModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *ClientIsolationWhitelistModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/clientIsolationWhitelist/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatchRequestSlice []*ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatchRequest

	ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatchRequest struct {
		Description *string `json:"description,omitempty"`
		IP          *string `json:"ip,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}
)

// ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatch: Modify Client Isolation Whitelist entries configuration
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatch(ctx context.Context, zoneId string, id string, requestBody ClientIsolationWhitelistModifyClientIsolationWhitelistEntriesConfigurationPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/clientIsolationWhitelist/{id}/whitelist", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DynamicPskRetrieveIntervalOfDeleteExpiredDpskGet200Response struct {
		DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
	}
)

// DynamicPskRetrieveIntervalOfDeleteExpiredDpskGet: Use this API command to retrieve interval of delete expired DPSK of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskRetrieveIntervalOfDeleteExpiredDpskGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskRetrieveIntervalOfDeleteExpiredDpskGet(ctx context.Context, zoneId string) (*http.Response, *DynamicPskRetrieveIntervalOfDeleteExpiredDpskGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/deleteExpiredDpsk", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &DynamicPskRetrieveIntervalOfDeleteExpiredDpskGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DynamicPskModifyIntervalOfDeleteExpiredDpskPutRequest struct {
		DeleteExpiredDpsk *string `json:"deleteExpiredDpsk,omitempty"`
	}
)

// DynamicPskModifyIntervalOfDeleteExpiredDpskPut: Use this API command to modify interval of delete expired DPSK of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *DynamicPskModifyIntervalOfDeleteExpiredDpskPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskModifyIntervalOfDeleteExpiredDpskPut(ctx context.Context, zoneId string, requestBody *DynamicPskModifyIntervalOfDeleteExpiredDpskPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/rkszones/{zoneId}/deleteExpiredDpsk", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DevicePolicyRetrieveListDevicePolicyPorfileGet200ResponseListSlice []*DevicePolicyRetrieveListDevicePolicyPorfileGet200ResponseList

	DevicePolicyRetrieveListDevicePolicyPorfileGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	DevicePolicyRetrieveListDevicePolicyPorfileGet200Response struct {
		FirstIndex *int                                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                                              `json:"hasMore,omitempty"`
		List       DevicePolicyRetrieveListDevicePolicyPorfileGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                               `json:"totalCount,omitempty"`
	}
)

// DevicePolicyRetrieveListDevicePolicyPorfileGet: Retrieve a list of Device Policy Porfiles within a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DevicePolicyRetrieveListDevicePolicyPorfileGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DevicePolicyRetrieveListDevicePolicyPorfileGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *DevicePolicyRetrieveListDevicePolicyPorfileGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/devicePolicy", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &DevicePolicyRetrieveListDevicePolicyPorfileGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DevicePolicyCreateDevicePolicyPorfilePostRequest struct {
		DefaultAction *string `json:"defaultAction,omitempty"`
		Description   *string `json:"description,omitempty"`
		Name          *string `json:"name,omitempty"`
	}

	DevicePolicyCreateDevicePolicyPorfilePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// DevicePolicyCreateDevicePolicyPorfilePost: Create a new Device Policy Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *DevicePolicyCreateDevicePolicyPorfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DevicePolicyCreateDevicePolicyPorfilePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DevicePolicyCreateDevicePolicyPorfilePost(ctx context.Context, zoneId string, requestBody *DevicePolicyCreateDevicePolicyPorfilePostRequest) (*http.Response, *DevicePolicyCreateDevicePolicyPorfilePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/devicePolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &DevicePolicyCreateDevicePolicyPorfilePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// DevicePolicyDeleteDevicePolicyPorfileDelete: Delete Device Policy Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DevicePolicyDeleteDevicePolicyPorfileDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/devicePolicy/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DevicePolicyRetrieveDevicePolicyPorfileGet200ResponseRuleSlice []*DevicePolicyRetrieveDevicePolicyPorfileGet200ResponseRule

	DevicePolicyRetrieveDevicePolicyPorfileGet200ResponseRule struct {
		Action      *string  `json:"action,omitempty"`
		Description *string  `json:"description,omitempty"`
		DeviceType  *string  `json:"deviceType,omitempty"`
		Downlink    *float64 `json:"downlink,omitempty"`
		Uplink      *float64 `json:"uplink,omitempty"`
		Vlan        *float64 `json:"vlan,omitempty"`
	}

	DevicePolicyRetrieveDevicePolicyPorfileGet200Response struct {
		DefaultAction *string                                                        `json:"defaultAction,omitempty"`
		Description   *string                                                        `json:"description,omitempty"`
		ID            *string                                                        `json:"id,omitempty"`
		Name          *string                                                        `json:"name,omitempty"`
		Rule          DevicePolicyRetrieveDevicePolicyPorfileGet200ResponseRuleSlice `json:"rule,omitempty"`
	}
)

// DevicePolicyRetrieveDevicePolicyPorfileGet: Retrieve a Device Policy Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DevicePolicyRetrieveDevicePolicyPorfileGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DevicePolicyRetrieveDevicePolicyPorfileGet(ctx context.Context, zoneId string, id string) (*http.Response, *DevicePolicyRetrieveDevicePolicyPorfileGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/devicePolicy/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &DevicePolicyRetrieveDevicePolicyPorfileGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DevicePolicyModifyDevicePolicyPorfilePatchRequest struct {
		DefaultAction *string `json:"defaultAction,omitempty"`
		Description   *string `json:"description,omitempty"`
		Name          *string `json:"name,omitempty"`
	}
)

// DevicePolicyModifyDevicePolicyPorfilePatch: Modify a specific Device Policy Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DevicePolicyModifyDevicePolicyPorfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DevicePolicyModifyDevicePolicyPorfilePatch(ctx context.Context, zoneId string, id string, requestBody *DevicePolicyModifyDevicePolicyPorfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/devicePolicy/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DevicePolicyModifyRuleOfDevicePolicyPorfilePatchRequestSlice []*DevicePolicyModifyRuleOfDevicePolicyPorfilePatchRequest

	DevicePolicyModifyRuleOfDevicePolicyPorfilePatchRequest struct {
		Action      *string  `json:"action,omitempty"`
		Description *string  `json:"description,omitempty"`
		DeviceType  *string  `json:"deviceType,omitempty"`
		Downlink    *float64 `json:"downlink,omitempty"`
		Uplink      *float64 `json:"uplink,omitempty"`
		Vlan        *float64 `json:"vlan,omitempty"`
	}
)

// DevicePolicyModifyRuleOfDevicePolicyPorfilePatch: Modify Rule of Device Policy Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DevicePolicyModifyRuleOfDevicePolicyPorfilePatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DevicePolicyModifyRuleOfDevicePolicyPorfilePatch(ctx context.Context, zoneId string, id string, requestBody DevicePolicyModifyRuleOfDevicePolicyPorfilePatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/devicePolicy/{id}/rule", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DhcpDeleteMultipleDhcpPoolsDeleteRequestIDListSlice []*string

	DhcpDeleteMultipleDhcpPoolsDeleteRequest struct {
		IDList DhcpDeleteMultipleDhcpPoolsDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// DhcpDeleteMultipleDhcpPoolsDelete: Use this API command to delete multiple DHCP Pools.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *DhcpDeleteMultipleDhcpPoolsDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpDeleteMultipleDhcpPoolsDelete(ctx context.Context, zoneId string, requestBody *DhcpDeleteMultipleDhcpPoolsDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DhcpGetDhcpPoolListGet200ResponseListSlice []*DhcpGetDhcpPoolListGet200ResponseList

	DhcpGetDhcpPoolListGet200ResponseList struct {
		Description      *string  `json:"description,omitempty"`
		ID               *string  `json:"id,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PoolEndIP        *string  `json:"poolEndIp,omitempty"`
		PoolStartIP      *string  `json:"poolStartIp,omitempty"`
		PrimaryDNSIP     *string  `json:"primaryDnsIp,omitempty"`
		SecondaryDNSIP   *string  `json:"secondaryDnsIp,omitempty"`
		SubnetMask       *string  `json:"subnetMask,omitempty"`
		SubnetNetworkIP  *string  `json:"subnetNetworkIp,omitempty"`
		VlanID           *float64 `json:"vlanId,omitempty"`
		ZoneID           *string  `json:"zoneId,omitempty"`
	}

	DhcpGetDhcpPoolListGet200Response struct {
		CreateDateTime   *int                                       `json:"createDateTime,omitempty"`
		CreatorID        *string                                    `json:"creatorId,omitempty"`
		CreatorUsername  *string                                    `json:"creatorUsername,omitempty"`
		FirstIndex       *int                                       `json:"firstIndex,omitempty"`
		HasMore          *bool                                      `json:"hasMore,omitempty"`
		List             DhcpGetDhcpPoolListGet200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                       `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                    `json:"modifierId,omitempty"`
		ModifierUsername *string                                    `json:"modifierUsername,omitempty"`
		TotalCount       *int                                       `json:"totalCount,omitempty"`
	}
)

// DhcpGetDhcpPoolListGet: Use this API command to get DHCP Pool list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpGetDhcpPoolListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpGetDhcpPoolListGet(ctx context.Context, zoneId string) (*http.Response, *DhcpGetDhcpPoolListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpProfile", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &DhcpGetDhcpPoolListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DhcpCreateDhcpPoolPostRequest struct {
		Description      *string  `json:"description,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PoolEndIP        *string  `json:"poolEndIp,omitempty"`
		PoolStartIP      *string  `json:"poolStartIp,omitempty"`
		PrimaryDNSIP     *string  `json:"primaryDnsIp,omitempty"`
		SecondaryDNSIP   *string  `json:"secondaryDnsIp,omitempty"`
		SubnetMask       *string  `json:"subnetMask,omitempty"`
		SubnetNetworkIP  *string  `json:"subnetNetworkIp,omitempty"`
		VlanID           *float64 `json:"vlanId,omitempty"`
	}

	DhcpCreateDhcpPoolPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// DhcpCreateDhcpPoolPost: Use this API command to create DHCP Pool.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *DhcpCreateDhcpPoolPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpCreateDhcpPoolPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpCreateDhcpPoolPost(ctx context.Context, zoneId string, requestBody *DhcpCreateDhcpPoolPostRequest) (*http.Response, *DhcpCreateDhcpPoolPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &DhcpCreateDhcpPoolPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// DhcpDeleteDhcpPoolByPoolSIdDelete: Use this API command to delete DHCP Pool by pool's ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpDeleteDhcpPoolByPoolSIdDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpProfile/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DhcpGetDhcpPoolByPoolSIdGet200Response struct {
		Description      *string  `json:"description,omitempty"`
		ID               *string  `json:"id,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PoolEndIP        *string  `json:"poolEndIp,omitempty"`
		PoolStartIP      *string  `json:"poolStartIp,omitempty"`
		PrimaryDNSIP     *string  `json:"primaryDnsIp,omitempty"`
		SecondaryDNSIP   *string  `json:"secondaryDnsIp,omitempty"`
		SubnetMask       *string  `json:"subnetMask,omitempty"`
		SubnetNetworkIP  *string  `json:"subnetNetworkIp,omitempty"`
		VlanID           *float64 `json:"vlanId,omitempty"`
		ZoneID           *string  `json:"zoneId,omitempty"`
	}
)

// DhcpGetDhcpPoolByPoolSIdGet: Use this API command to get DHCP Pool by pool's ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpGetDhcpPoolByPoolSIdGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpGetDhcpPoolByPoolSIdGet(ctx context.Context, zoneId string, id string) (*http.Response, *DhcpGetDhcpPoolByPoolSIdGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpProfile/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &DhcpGetDhcpPoolByPoolSIdGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DhcpModifyDhcpProfileByPoolSIdPatchRequest struct {
		Description      *string  `json:"description,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PoolEndIP        *string  `json:"poolEndIp,omitempty"`
		PoolStartIP      *string  `json:"poolStartIp,omitempty"`
		PrimaryDNSIP     *string  `json:"primaryDnsIp,omitempty"`
		SecondaryDNSIP   *string  `json:"secondaryDnsIp,omitempty"`
		SubnetMask       *string  `json:"subnetMask,omitempty"`
		SubnetNetworkIP  *string  `json:"subnetNetworkIp,omitempty"`
		VlanID           *float64 `json:"vlanId,omitempty"`
	}
)

// DhcpModifyDhcpProfileByPoolSIdPatch: Use this API command to modify DHCP Pool by pool's ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DhcpModifyDhcpProfileByPoolSIdPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpModifyDhcpProfileByPoolSIdPatch(ctx context.Context, zoneId string, id string, requestBody *DhcpModifyDhcpProfileByPoolSIdPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpProfile/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DhcpGetDhcpConfigurationGet200ResponseSiteApsSlice []*DhcpGetDhcpConfigurationGet200ResponseSiteAps

	DhcpGetDhcpConfigurationGet200ResponseSiteAps struct {
		ApGatewayEnabled *bool   `json:"apGatewayEnabled,omitempty"`
		ApGatewayIP      *string `json:"apGatewayIp,omitempty"`
		ApMac            *string `json:"apMac,omitempty"`
		ApName           *string `json:"apName,omitempty"`
		ApServerEnabled  *bool   `json:"apServerEnabled,omitempty"`
		ApServerIP       *string `json:"apServerIp,omitempty"`
		ApServerPrimary  *bool   `json:"apServerPrimary,omitempty"`
		ApStatus         *string `json:"apStatus,omitempty"`
	}

	DhcpGetDhcpConfigurationGet200ResponseSiteProfilesSlice []*string

	DhcpGetDhcpConfigurationGet200Response struct {
		ManualSelect *bool                                                   `json:"manualSelect,omitempty"`
		SiteAps      DhcpGetDhcpConfigurationGet200ResponseSiteApsSlice      `json:"siteAps,omitempty"`
		SiteEnabled  *bool                                                   `json:"siteEnabled,omitempty"`
		SiteMode     *string                                                 `json:"siteMode,omitempty"`
		SiteProfiles DhcpGetDhcpConfigurationGet200ResponseSiteProfilesSlice `json:"siteProfiles,omitempty"`
		ZoneName     *string                                                 `json:"zoneName,omitempty"`
	}
)

// DhcpGetDhcpConfigurationGet: Use this API command to get DHCP Configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpGetDhcpConfigurationGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DhcpGetDhcpConfigurationGet(ctx context.Context, zoneId string) (*http.Response, *DhcpGetDhcpConfigurationGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpSiteConfig", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &DhcpGetDhcpConfigurationGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequestSiteApsSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequestSiteAps

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequestSiteAps struct {
		ApMac           *string `json:"apMac,omitempty"`
		ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
		ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequestSiteProfileIdsSlice []*string

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequest struct {
		ManualSelect   *bool                                                                                `json:"manualSelect,omitempty"`
		SiteAps        RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequestSiteApsSlice        `json:"siteAps,omitempty"`
		SiteEnabled    *bool                                                                                `json:"siteEnabled,omitempty"`
		SiteMode       *string                                                                              `json:"siteMode,omitempty"`
		SiteProfileIds RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequestSiteProfileIdsSlice `json:"siteProfileIds,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200ResponseSiteApsSlice []*RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200ResponseSiteAps

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200ResponseSiteAps struct {
		ApGatewayEnabled *bool   `json:"apGatewayEnabled,omitempty"`
		ApGatewayIP      *string `json:"apGatewayIp,omitempty"`
		ApMac            *string `json:"apMac,omitempty"`
		ApName           *string `json:"apName,omitempty"`
		ApServerEnabled  *bool   `json:"apServerEnabled,omitempty"`
		ApServerIP       *string `json:"apServerIp,omitempty"`
		ApServerPrimary  *bool   `json:"apServerPrimary,omitempty"`
		ApStatus         *string `json:"apStatus,omitempty"`
	}

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200ResponseSiteProfilesSlice []*string

	RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200Response struct {
		ManualSelect *bool                                                                                  `json:"manualSelect,omitempty"`
		SiteAps      RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200ResponseSiteApsSlice      `json:"siteAps,omitempty"`
		SiteEnabled  *bool                                                                                  `json:"siteEnabled,omitempty"`
		SiteMode     *string                                                                                `json:"siteMode,omitempty"`
		SiteProfiles RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200ResponseSiteProfilesSlice `json:"siteProfiles,omitempty"`
		ZoneName     *string                                                                                `json:"zoneName,omitempty"`
	}
)

// RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost: Use this API command to get the DHCP/NAT service IP assignment when selecting with Enable on Multiple APs. In Manually Select AP mode(the manualSelect is true), the body should contain with selected APs(include the siteAps array). Otherwise, it's no need to include the selected APs in Auto Select AP mode(see samples).
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost(ctx context.Context, zoneId string, requestBody *RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPostRequest) (*http.Response, *RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/dhcpSite/dhcpSiteConfig/doAssignIp", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &RuckusWirelessApZoneRetrieveDhcpNatServiceIpAssignmentPost200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DiffservRetrieveListGet200ResponseListSlice []*DiffservRetrieveListGet200ResponseList

	DiffservRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	DiffservRetrieveListGet200Response struct {
		FirstIndex *int                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                       `json:"hasMore,omitempty"`
		List       DiffservRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                        `json:"totalCount,omitempty"`
	}
)

// DiffservRetrieveListGet: Use this API command to retrieve a list of DiffServ profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DiffservRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *DiffservRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/diffserv", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &DiffservRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DiffservCreatePostRequestDownlinkDiffServ struct {
		Downlink       *string `json:"downlink,omitempty"`
		DownlinkEnable *bool   `json:"downlinkEnable,omitempty"`
	}

	DiffservCreatePostRequestPreservedListSlice []*string

	DiffservCreatePostRequestUplinkDiffServ struct {
		Uplink       *string `json:"uplink,omitempty"`
		UplinkEnable *bool   `json:"uplinkEnable,omitempty"`
	}

	DiffservCreatePostRequest struct {
		Description      *string                                     `json:"description,omitempty"`
		DownlinkDiffServ *DiffservCreatePostRequestDownlinkDiffServ  `json:"downlinkDiffServ,omitempty"`
		Name             *string                                     `json:"name,omitempty"`
		PreservedList    DiffservCreatePostRequestPreservedListSlice `json:"preservedList,omitempty"`
		UplinkDiffServ   *DiffservCreatePostRequestUplinkDiffServ    `json:"uplinkDiffServ,omitempty"`
	}

	DiffservCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// DiffservCreatePost: Use this API command to create DiffServ profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *DiffservCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DiffservCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservCreatePost(ctx context.Context, zoneId string, requestBody *DiffservCreatePostRequest) (*http.Response, *DiffservCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/diffserv", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &DiffservCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// DiffservDeleteDelete: Use this API command to delete DiffServ profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/diffserv/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DiffservRetrieveGet200ResponseDownlinkDiffServ struct {
		Downlink       *string `json:"downlink,omitempty"`
		DownlinkEnable *bool   `json:"downlinkEnable,omitempty"`
	}

	DiffservRetrieveGet200ResponsePreservedListSlice []*string

	DiffservRetrieveGet200ResponseUplinkDiffServ struct {
		Uplink       *string `json:"uplink,omitempty"`
		UplinkEnable *bool   `json:"uplinkEnable,omitempty"`
	}

	DiffservRetrieveGet200Response struct {
		Description      *string                                          `json:"description,omitempty"`
		DownlinkDiffServ *DiffservRetrieveGet200ResponseDownlinkDiffServ  `json:"downlinkDiffServ,omitempty"`
		ID               *string                                          `json:"id,omitempty"`
		Name             *string                                          `json:"name,omitempty"`
		PreservedList    DiffservRetrieveGet200ResponsePreservedListSlice `json:"preservedList,omitempty"`
		UplinkDiffServ   *DiffservRetrieveGet200ResponseUplinkDiffServ    `json:"uplinkDiffServ,omitempty"`
	}
)

// DiffservRetrieveGet: Use this API command to retrieve DiffServ profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DiffservRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *DiffservRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/diffserv/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &DiffservRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DiffservModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// DiffservModifyBasicPatch: Use this API command to modify the basic information of DiffServ profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DiffservModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *DiffservModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/diffserv/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DiffservModifyDownlinkDiffservPatchRequest struct {
		Downlink       *string `json:"downlink,omitempty"`
		DownlinkEnable *bool   `json:"downlinkEnable,omitempty"`
	}
)

// DiffservModifyDownlinkDiffservPatch: Use this API command to modify downlink diffserv.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DiffservModifyDownlinkDiffservPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservModifyDownlinkDiffservPatch(ctx context.Context, zoneId string, id string, requestBody *DiffservModifyDownlinkDiffservPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/diffserv/{id}/downlinkDiffServ", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DiffservModifyPreservedListPatchRequest []*string
)

// DiffservModifyPreservedListPatch: Use this API command to modify preserved list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DiffservModifyPreservedListPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservModifyPreservedListPatch(ctx context.Context, zoneId string, id string, requestBody DiffservModifyPreservedListPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/diffserv/{id}/preservedList", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DiffservModifyUplinkDiffservPatchRequest struct {
		Uplink       *string `json:"uplink,omitempty"`
		UplinkEnable *bool   `json:"uplinkEnable,omitempty"`
	}
)

// DiffservModifyUplinkDiffservPatch: Use this API command to modify uplink diffserv.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *DiffservModifyUplinkDiffservPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DiffservModifyUplinkDiffservPatch(ctx context.Context, zoneId string, id string, requestBody *DiffservModifyUplinkDiffservPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/diffserv/{id}/uplinkDiffServ", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DynamicPskRetrieveDpskInfoByZoneGet200ResponseListSlice []*DynamicPskRetrieveDpskInfoByZoneGet200ResponseList

	DynamicPskRetrieveDpskInfoByZoneGet200ResponseList struct {
		CreationDateTime   *float64 `json:"creationDateTime,omitempty"`
		ExpirationDateTime *string  `json:"expirationDateTime,omitempty"`
		GroupDpsk          *bool    `json:"groupDpsk,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		MacAddress         *string  `json:"macAddress,omitempty"`
		Passphrase         *string  `json:"passphrase,omitempty"`
		UserName           *string  `json:"userName,omitempty"`
		UserRoleID         *string  `json:"userRoleId,omitempty"`
		VlanID             *float64 `json:"vlanId,omitempty"`
		WlanID             *string  `json:"wlanId,omitempty"`
	}

	DynamicPskRetrieveDpskInfoByZoneGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       DynamicPskRetrieveDpskInfoByZoneGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// DynamicPskRetrieveDpskInfoByZoneGet: Use this API command to retrieve DPSK info of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskRetrieveDpskInfoByZoneGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskRetrieveDpskInfoByZoneGet(ctx context.Context, zoneId string) (*http.Response, *DynamicPskRetrieveDpskInfoByZoneGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/dpsk", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &DynamicPskRetrieveDpskInfoByZoneGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200ResponseListSlice []*DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200ResponseList

	DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200ResponseList struct {
		Ssid     *string `json:"ssid,omitempty"`
		WlanID   *string `json:"wlanId,omitempty"`
		WlanName *string `json:"wlanName,omitempty"`
	}

	DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200Response struct {
		FirstIndex *int                                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                                              `json:"hasMore,omitempty"`
		List       DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                               `json:"totalCount,omitempty"`
	}
)

// DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet: Use this API command to retrieve DPSK enabled WLAN info of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet(ctx context.Context, zoneId string) (*http.Response, *DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/dpskEnabledWlans", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &DynamicPskRetrieveDpskEnabledWlanInfoByZoneGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20WlanProfileRetrieveListGet200ResponseListSlice []*Hotspot20WlanProfileRetrieveListGet200ResponseList

	Hotspot20WlanProfileRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileRetrieveListGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       Hotspot20WlanProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// Hotspot20WlanProfileRetrieveListGet: Use this API command to retrieve a list of Hotspot 2.0 WLAN profiles of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WlanProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *Hotspot20WlanProfileRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/hs20s", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &Hotspot20WlanProfileRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20WlanProfileCreatePostRequestDefaultIdentityProvider struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileCreatePostRequestIdentityProvidersSlice []*Hotspot20WlanProfileCreatePostRequestIdentityProviders

	Hotspot20WlanProfileCreatePostRequestIdentityProviders struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileCreatePostRequestOperator struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileCreatePostRequestSignupSsid struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileCreatePostRequest struct {
		AccessNetworkType       *string                                                       `json:"accessNetworkType,omitempty"`
		DefaultIdentityProvider *Hotspot20WlanProfileCreatePostRequestDefaultIdentityProvider `json:"defaultIdentityProvider,omitempty"`
		Description             *string                                                       `json:"description,omitempty"`
		IdentityProviders       Hotspot20WlanProfileCreatePostRequestIdentityProvidersSlice   `json:"identityProviders,omitempty"`
		InternetOption          *bool                                                         `json:"internetOption,omitempty"`
		Ipv4AddressType         *string                                                       `json:"ipv4AddressType,omitempty"`
		Ipv6AddressType         *string                                                       `json:"ipv6AddressType,omitempty"`
		Name                    *string                                                       `json:"name,omitempty"`
		Operator                *Hotspot20WlanProfileCreatePostRequestOperator                `json:"operator,omitempty"`
		SignupSsid              *Hotspot20WlanProfileCreatePostRequestSignupSsid              `json:"signupSsid,omitempty"`
	}

	Hotspot20WlanProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// Hotspot20WlanProfileCreatePost: Use this API command to create a new Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *Hotspot20WlanProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WlanProfileCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileCreatePost(ctx context.Context, zoneId string, requestBody *Hotspot20WlanProfileCreatePostRequest) (*http.Response, *Hotspot20WlanProfileCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/hs20s", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &Hotspot20WlanProfileCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// Hotspot20WlanProfileDeleteDelete: Use this API command to delete a Hotspot 2.0 WLAN Profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/hs20s/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileRetrieveGet200ResponseConnectionCapabilitiesSlice []*Hotspot20WlanProfileRetrieveGet200ResponseConnectionCapabilities

	Hotspot20WlanProfileRetrieveGet200ResponseConnectionCapabilities struct {
		PortNumber     *float64 `json:"portNumber,omitempty"`
		ProtocolName   *string  `json:"protocolName,omitempty"`
		ProtocolNumber *float64 `json:"protocolNumber,omitempty"`
		Status         *string  `json:"status,omitempty"`
	}

	Hotspot20WlanProfileRetrieveGet200ResponseCustomConnectionCapabilitiesSlice []*Hotspot20WlanProfileRetrieveGet200ResponseCustomConnectionCapabilities

	Hotspot20WlanProfileRetrieveGet200ResponseCustomConnectionCapabilities struct {
		PortNumber     *float64 `json:"portNumber,omitempty"`
		ProtocolName   *string  `json:"protocolName,omitempty"`
		ProtocolNumber *float64 `json:"protocolNumber,omitempty"`
		Status         *string  `json:"status,omitempty"`
	}

	Hotspot20WlanProfileRetrieveGet200ResponseDefaultIdentityProvider struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileRetrieveGet200ResponseIdentityProvidersSlice []*Hotspot20WlanProfileRetrieveGet200ResponseIdentityProviders

	Hotspot20WlanProfileRetrieveGet200ResponseIdentityProviders struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileRetrieveGet200ResponseOperator struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileRetrieveGet200ResponseSignupSsid struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20WlanProfileRetrieveGet200Response struct {
		AccessNetworkType            *string                                                                     `json:"accessNetworkType,omitempty"`
		ConnectionCapabilities       Hotspot20WlanProfileRetrieveGet200ResponseConnectionCapabilitiesSlice       `json:"connectionCapabilities,omitempty"`
		CustomConnectionCapabilities Hotspot20WlanProfileRetrieveGet200ResponseCustomConnectionCapabilitiesSlice `json:"customConnectionCapabilities,omitempty"`
		DefaultIdentityProvider      *Hotspot20WlanProfileRetrieveGet200ResponseDefaultIdentityProvider          `json:"defaultIdentityProvider,omitempty"`
		Description                  *string                                                                     `json:"description,omitempty"`
		ID                           *string                                                                     `json:"id,omitempty"`
		IdentityProviders            Hotspot20WlanProfileRetrieveGet200ResponseIdentityProvidersSlice            `json:"identityProviders,omitempty"`
		InternetOption               *bool                                                                       `json:"internetOption,omitempty"`
		Ipv4AddressType              *string                                                                     `json:"ipv4AddressType,omitempty"`
		Ipv6AddressType              *string                                                                     `json:"ipv6AddressType,omitempty"`
		Name                         *string                                                                     `json:"name,omitempty"`
		Operator                     *Hotspot20WlanProfileRetrieveGet200ResponseOperator                         `json:"operator,omitempty"`
		SignupSsid                   *Hotspot20WlanProfileRetrieveGet200ResponseSignupSsid                       `json:"signupSsid,omitempty"`
		ZoneID                       *string                                                                     `json:"zoneId,omitempty"`
	}
)

// Hotspot20WlanProfileRetrieveGet: Use this API command to retrieve a Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20WlanProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *Hotspot20WlanProfileRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/hs20s/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &Hotspot20WlanProfileRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20WlanProfileModifyBasicPatchRequest struct {
		AccessNetworkType *string `json:"accessNetworkType,omitempty"`
		Description       *string `json:"description,omitempty"`
		InternetOption    *bool   `json:"internetOption,omitempty"`
		Ipv4AddressType   *string `json:"ipv4AddressType,omitempty"`
		Ipv6AddressType   *string `json:"ipv6AddressType,omitempty"`
		Name              *string `json:"name,omitempty"`
	}
)

// Hotspot20WlanProfileModifyBasicPatch: Use this API command to modify the basic information on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *Hotspot20WlanProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatchRequestSlice []*Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatchRequest

	Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatchRequest struct {
		PortNumber     *float64 `json:"portNumber,omitempty"`
		ProtocolName   *string  `json:"protocolName,omitempty"`
		ProtocolNumber *float64 `json:"protocolNumber,omitempty"`
		Status         *string  `json:"status,omitempty"`
	}
)

// Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatch: Use this API command to modify connection capabilities on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatch(ctx context.Context, zoneId string, id string, requestBody Hotspot20WlanProfileModifyDefaultConnectionCapabilitiesPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}/connectionCapabilities", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatchRequestSlice []*Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatchRequest

	Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatchRequest struct {
		PortNumber     *float64 `json:"portNumber,omitempty"`
		ProtocolName   *string  `json:"protocolName,omitempty"`
		ProtocolNumber *float64 `json:"protocolNumber,omitempty"`
		Status         *string  `json:"status,omitempty"`
	}
)

// Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatch: Use this API command to modify custom connection capabilities on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatch(ctx context.Context, zoneId string, id string, requestBody Hotspot20WlanProfileModifyCustomConnectionCapabilitiesPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}/customConnectionCapabilities", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileModifyDefaultIdentityProviderPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// Hotspot20WlanProfileModifyDefaultIdentityProviderPatch: Use this API command to modify the default identity provider profile on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifyDefaultIdentityProviderPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifyDefaultIdentityProviderPatch(ctx context.Context, zoneId string, id string, requestBody *Hotspot20WlanProfileModifyDefaultIdentityProviderPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}/defaultIdentityProvider", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileModifyIdentityProvidersPatchRequestSlice []*Hotspot20WlanProfileModifyIdentityProvidersPatchRequest

	Hotspot20WlanProfileModifyIdentityProvidersPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// Hotspot20WlanProfileModifyIdentityProvidersPatch: Use this API command to modify identity provider profiles on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifyIdentityProvidersPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifyIdentityProvidersPatch(ctx context.Context, zoneId string, id string, requestBody Hotspot20WlanProfileModifyIdentityProvidersPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}/identityProviders", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileModifyOperatorPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// Hotspot20WlanProfileModifyOperatorPatch: Use this API command to modify operator profile on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifyOperatorPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifyOperatorPatch(ctx context.Context, zoneId string, id string, requestBody *Hotspot20WlanProfileModifyOperatorPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}/operator", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20WlanProfileModifySignupSsidPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// Hotspot20WlanProfileModifySignupSsidPatch: Use this API command to modify signup SSID on Hotspot 2.0 WLAN profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Profile ID
// - requestBody: *Hotspot20WlanProfileModifySignupSsidPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20WlanProfileModifySignupSsidPatch(ctx context.Context, zoneId string, id string, requestBody *Hotspot20WlanProfileModifySignupSsidPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20s/{id}/signupSsid", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20VenueProfileRetrieveListGet200ResponseListSlice []*Hotspot20VenueProfileRetrieveListGet200ResponseList

	Hotspot20VenueProfileRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	Hotspot20VenueProfileRetrieveListGet200Response struct {
		FirstIndex *int                                                     `json:"firstIndex,omitempty"`
		HasMore    *bool                                                    `json:"hasMore,omitempty"`
		List       Hotspot20VenueProfileRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                     `json:"totalCount,omitempty"`
	}
)

// Hotspot20VenueProfileRetrieveListGet: Use this API command to retrieve a list of Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20VenueProfileRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20VenueProfileRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *Hotspot20VenueProfileRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/hs20/venues", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &Hotspot20VenueProfileRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20VenueProfileCreatePostRequestVenueNamesSlice []*Hotspot20VenueProfileCreatePostRequestVenueNames

	Hotspot20VenueProfileCreatePostRequestVenueNames struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}

	Hotspot20VenueProfileCreatePostRequest struct {
		Description         *string                                               `json:"description,omitempty"`
		DownlinkSpeedInKbps *float64                                              `json:"downlinkSpeedInKbps,omitempty"`
		Group               *string                                               `json:"group,omitempty"`
		Name                *string                                               `json:"name,omitempty"`
		Type                *string                                               `json:"type,omitempty"`
		UplinkSpeedInKbps   *float64                                              `json:"uplinkSpeedInKbps,omitempty"`
		VenueNames          Hotspot20VenueProfileCreatePostRequestVenueNamesSlice `json:"venueNames,omitempty"`
	}

	Hotspot20VenueProfileCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// Hotspot20VenueProfileCreatePost: Use this API command to create a new Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *Hotspot20VenueProfileCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20VenueProfileCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20VenueProfileCreatePost(ctx context.Context, zoneId string, requestBody *Hotspot20VenueProfileCreatePostRequest) (*http.Response, *Hotspot20VenueProfileCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/hs20/venues", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &Hotspot20VenueProfileCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// Hotspot20VenueProfileDeleteDelete: Use this API command to delete Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Venue Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20VenueProfileDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/hs20/venues/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20VenueProfileRetrieveGet200ResponseVenueNamesSlice []*Hotspot20VenueProfileRetrieveGet200ResponseVenueNames

	Hotspot20VenueProfileRetrieveGet200ResponseVenueNames struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}

	Hotspot20VenueProfileRetrieveGet200Response struct {
		Description         *string                                                    `json:"description,omitempty"`
		DownlinkSpeedInKbps *float64                                                   `json:"downlinkSpeedInKbps,omitempty"`
		Group               *string                                                    `json:"group,omitempty"`
		ID                  *string                                                    `json:"id,omitempty"`
		Name                *string                                                    `json:"name,omitempty"`
		Type                *string                                                    `json:"type,omitempty"`
		UplinkSpeedInKbps   *float64                                                   `json:"uplinkSpeedInKbps,omitempty"`
		VenueNames          Hotspot20VenueProfileRetrieveGet200ResponseVenueNamesSlice `json:"venueNames,omitempty"`
		ZoneID              *string                                                    `json:"zoneId,omitempty"`
	}
)

// Hotspot20VenueProfileRetrieveGet: Use this API command to retrieve a Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Venue Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *Hotspot20VenueProfileRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20VenueProfileRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *Hotspot20VenueProfileRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/hs20/venues/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &Hotspot20VenueProfileRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	Hotspot20VenueProfileModifyBasicPatchRequest struct {
		Description         *string  `json:"description,omitempty"`
		DownlinkSpeedInKbps *float64 `json:"downlinkSpeedInKbps,omitempty"`
		Group               *string  `json:"group,omitempty"`
		Name                *string  `json:"name,omitempty"`
		Type                *string  `json:"type,omitempty"`
		UplinkSpeedInKbps   *float64 `json:"uplinkSpeedInKbps,omitempty"`
	}
)

// Hotspot20VenueProfileModifyBasicPatch: Use this API command to modify the basic information on Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Venue Profile ID
// - requestBody: *Hotspot20VenueProfileModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20VenueProfileModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *Hotspot20VenueProfileModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20/venues/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	Hotspot20VenueProfileModifyVenueNamesPatchRequestSlice []*Hotspot20VenueProfileModifyVenueNamesPatchRequest

	Hotspot20VenueProfileModifyVenueNamesPatchRequest struct {
		Language *string `json:"language,omitempty"`
		Name     *string `json:"name,omitempty"`
	}
)

// Hotspot20VenueProfileModifyVenueNamesPatch: Use this API command to modify the venue names on Hotspot 2.0 venue profile of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot 2.0 Venue Profile ID
// - requestBody: *Hotspot20VenueProfileModifyVenueNamesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) Hotspot20VenueProfileModifyVenueNamesPatch(ctx context.Context, zoneId string, id string, requestBody Hotspot20VenueProfileModifyVenueNamesPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/hs20/venues/{id}/venueNames", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	L2AccessControlRetrieveListGet200ResponseListSlice []*L2AccessControlRetrieveListGet200ResponseList

	L2AccessControlRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	L2AccessControlRetrieveListGet200Response struct {
		FirstIndex *int                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                              `json:"hasMore,omitempty"`
		List       L2AccessControlRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                               `json:"totalCount,omitempty"`
	}
)

// L2AccessControlRetrieveListGet: Retrieve a list of L2 Access Control
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2AccessControlRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) L2AccessControlRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *L2AccessControlRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/l2ACL", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &L2AccessControlRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	L2AccessControlCreateL2AccessControlPostRequestRuleMacsSlice []*string

	L2AccessControlCreateL2AccessControlPostRequest struct {
		Description *string                                                      `json:"description,omitempty"`
		Name        *string                                                      `json:"name,omitempty"`
		Restriction *string                                                      `json:"restriction,omitempty"`
		RuleMacs    L2AccessControlCreateL2AccessControlPostRequestRuleMacsSlice `json:"ruleMacs,omitempty"`
	}

	L2AccessControlCreateL2AccessControlPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// L2AccessControlCreateL2AccessControlPost: Create a new L2 Access Control
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *L2AccessControlCreateL2AccessControlPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2AccessControlCreateL2AccessControlPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) L2AccessControlCreateL2AccessControlPost(ctx context.Context, zoneId string, requestBody *L2AccessControlCreateL2AccessControlPostRequest) (*http.Response, *L2AccessControlCreateL2AccessControlPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/l2ACL", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &L2AccessControlCreateL2AccessControlPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// L2AccessControlDeleteDelete: Delete an L2 Access Control
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): L2 ACL ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) L2AccessControlDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/l2ACL/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	L2AccessControlRetrieveGet200ResponseRuleMacsSlice []*string

	L2AccessControlRetrieveGet200Response struct {
		Description *string                                            `json:"description,omitempty"`
		ID          *string                                            `json:"id,omitempty"`
		Name        *string                                            `json:"name,omitempty"`
		Restriction *string                                            `json:"restriction,omitempty"`
		RuleMacs    L2AccessControlRetrieveGet200ResponseRuleMacsSlice `json:"ruleMacs,omitempty"`
		ZoneID      *string                                            `json:"zoneId,omitempty"`
	}
)

// L2AccessControlRetrieveGet: Retrieve an L2 Access Control
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): L2 ACL ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *L2AccessControlRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) L2AccessControlRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *L2AccessControlRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/l2ACL/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &L2AccessControlRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	L2AccessControlModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
		Restriction *string `json:"restriction,omitempty"`
	}
)

// L2AccessControlModifyBasicPatch: Modify a specific L2 Access Control basic
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): L2 ACL ID
// - requestBody: *L2AccessControlModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) L2AccessControlModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *L2AccessControlModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/l2ACL/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	L2AccessControlModifyRuleMacsPatchRequest []*string
)

// L2AccessControlModifyRuleMacsPatch: Modify a specific L2 Access Control Rule Macs
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): L2 ACL ID
// - requestBody: *L2AccessControlModifyRuleMacsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) L2AccessControlModifyRuleMacsPatch(ctx context.Context, zoneId string, id string, requestBody L2AccessControlModifyRuleMacsPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/l2ACL/{id}/ruleMacs", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	GuestAccessRetrieveListGet200ResponseListSlice []*GuestAccessRetrieveListGet200ResponseList

	GuestAccessRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	GuestAccessRetrieveListGet200Response struct {
		FirstIndex *int                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                          `json:"hasMore,omitempty"`
		List       GuestAccessRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                           `json:"totalCount,omitempty"`
	}
)

// GuestAccessRetrieveListGet: Use this API command to retrieve a list of guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *GuestAccessRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *GuestAccessRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/guest", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &GuestAccessRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	GuestAccessCreatePostRequestPortalCustomization struct {
		Language                   *string `json:"language,omitempty"`
		Logo                       *string `json:"logo,omitempty"`
		TermsAndConditionsRequired *bool   `json:"termsAndConditionsRequired,omitempty"`
		TermsAndConditionsText     *string `json:"termsAndConditionsText,omitempty"`
		Title                      *string `json:"title,omitempty"`
	}

	GuestAccessCreatePostRequestRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	GuestAccessCreatePostRequestSmsGateway struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	GuestAccessCreatePostRequestUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	GuestAccessCreatePostRequest struct {
		Description         *string                                          `json:"description,omitempty"`
		Name                *string                                          `json:"name,omitempty"`
		PortalCustomization *GuestAccessCreatePostRequestPortalCustomization `json:"portalCustomization,omitempty"`
		Redirect            *GuestAccessCreatePostRequestRedirect            `json:"redirect,omitempty"`
		SmsGateway          *GuestAccessCreatePostRequestSmsGateway          `json:"smsGateway,omitempty"`
		UserSession         *GuestAccessCreatePostRequestUserSession         `json:"userSession,omitempty"`
	}

	GuestAccessCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// GuestAccessCreatePost: Use this API command to create new guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *GuestAccessCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *GuestAccessCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessCreatePost(ctx context.Context, zoneId string, requestBody *GuestAccessCreatePostRequest) (*http.Response, *GuestAccessCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/portals/guest", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &GuestAccessCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// GuestAccessDeleteDelete: Use this API command to delete guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/guest/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	GuestAccessRetrieveGet200ResponsePortalCustomization struct {
		Language                   *string `json:"language,omitempty"`
		Logo                       *string `json:"logo,omitempty"`
		TermsAndConditionsRequired *bool   `json:"termsAndConditionsRequired,omitempty"`
		TermsAndConditionsText     *string `json:"termsAndConditionsText,omitempty"`
		Title                      *string `json:"title,omitempty"`
	}

	GuestAccessRetrieveGet200ResponseRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	GuestAccessRetrieveGet200ResponseSmsGateway struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	GuestAccessRetrieveGet200ResponseUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	GuestAccessRetrieveGet200Response struct {
		Description         *string                                               `json:"description,omitempty"`
		ID                  *string                                               `json:"id,omitempty"`
		Name                *string                                               `json:"name,omitempty"`
		PortalCustomization *GuestAccessRetrieveGet200ResponsePortalCustomization `json:"portalCustomization,omitempty"`
		Redirect            *GuestAccessRetrieveGet200ResponseRedirect            `json:"redirect,omitempty"`
		SmsGateway          *GuestAccessRetrieveGet200ResponseSmsGateway          `json:"smsGateway,omitempty"`
		UserSession         *GuestAccessRetrieveGet200ResponseUserSession         `json:"userSession,omitempty"`
		ZoneID              *string                                               `json:"zoneId,omitempty"`
	}
)

// GuestAccessRetrieveGet: Use this API command to retrieve guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *GuestAccessRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *GuestAccessRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/guest/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &GuestAccessRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	GuestAccessModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// GuestAccessModifyBasicPatch: Use this API command to modify the basic information on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *GuestAccessModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *GuestAccessModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/guest/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	GuestAccessModifyPortalCustomizationPatchRequest struct {
		Language                   *string `json:"language,omitempty"`
		Logo                       *string `json:"logo,omitempty"`
		TermsAndConditionsRequired *bool   `json:"termsAndConditionsRequired,omitempty"`
		TermsAndConditionsText     *string `json:"termsAndConditionsText,omitempty"`
		Title                      *string `json:"title,omitempty"`
	}
)

// GuestAccessModifyPortalCustomizationPatch: Use this API command to modify the portal customization on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *GuestAccessModifyPortalCustomizationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessModifyPortalCustomizationPatch(ctx context.Context, zoneId string, id string, requestBody *GuestAccessModifyPortalCustomizationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/guest/{id}/portalCustomization", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// GuestAccessRedirectToUrlUserVisitDelete: Use this API command to set redirect to the URL that user intends to visit on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessRedirectToUrlUserVisitDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/guest/{id}/redirect", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	GuestAccessModifyRedirectPatchRequest struct {
		URL *string `json:"url,omitempty"`
	}
)

// GuestAccessModifyRedirectPatch: Use this API command to modify the redirect information on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *GuestAccessModifyRedirectPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessModifyRedirectPatch(ctx context.Context, zoneId string, id string, requestBody *GuestAccessModifyRedirectPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/guest/{id}/redirect", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// GuestAccessDisableSmsGatewayDelete: Use this API command to disable SMS gateway on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessDisableSmsGatewayDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/guest/{id}/smsGateway", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	GuestAccessModifySmsGatewayPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// GuestAccessModifySmsGatewayPatch: Use this API command to modify SMS gateway on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *GuestAccessModifySmsGatewayPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessModifySmsGatewayPatch(ctx context.Context, zoneId string, id string, requestBody *GuestAccessModifySmsGatewayPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/guest/{id}/smsGateway", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	GuestAccessModifyUserSessionPatchRequest struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}
)

// GuestAccessModifyUserSessionPatch: Use this API command to modify the user session on guest access of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *GuestAccessModifyUserSessionPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) GuestAccessModifyUserSessionPatch(ctx context.Context, zoneId string, id string, requestBody *GuestAccessModifyUserSessionPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/guest/{id}/userSession", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	HotspotServiceRetrieveListGet200ResponseListSlice []*HotspotServiceRetrieveListGet200ResponseList

	HotspotServiceRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	HotspotServiceRetrieveListGet200Response struct {
		FirstIndex *int                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                             `json:"hasMore,omitempty"`
		List       HotspotServiceRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                              `json:"totalCount,omitempty"`
	}
)

// HotspotServiceRetrieveListGet: Use this API command to retrieve a list of Hotspot(WISPr) of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *HotspotServiceRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *HotspotServiceRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/hotspot", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &HotspotServiceRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	HotspotServiceCreateExternalPostRequestLocation struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	HotspotServiceCreateExternalPostRequestRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	HotspotServiceCreateExternalPostRequestUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	HotspotServiceCreateExternalPostRequestWalledGardensSlice []*string

	HotspotServiceCreateExternalPostRequest struct {
		Description        *string                                                   `json:"description,omitempty"`
		Location           *HotspotServiceCreateExternalPostRequestLocation          `json:"location,omitempty"`
		MacAddressFormat   *float64                                                  `json:"macAddressFormat,omitempty"`
		Name               *string                                                   `json:"name,omitempty"`
		PortalURL          *string                                                   `json:"portalUrl,omitempty"`
		Redirect           *HotspotServiceCreateExternalPostRequestRedirect          `json:"redirect,omitempty"`
		SmartClientSupport *string                                                   `json:"smartClientSupport,omitempty"`
		UserSession        *HotspotServiceCreateExternalPostRequestUserSession       `json:"userSession,omitempty"`
		WalledGardens      HotspotServiceCreateExternalPostRequestWalledGardensSlice `json:"walledGardens,omitempty"`
	}

	HotspotServiceCreateExternalPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// HotspotServiceCreateExternalPost: Use this API command to create a new Hotspot(WISPr) with external logon URL of a zone.MacAddressFormat
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *HotspotServiceCreateExternalPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *HotspotServiceCreateExternalPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceCreateExternalPost(ctx context.Context, zoneId string, requestBody *HotspotServiceCreateExternalPostRequest) (*http.Response, *HotspotServiceCreateExternalPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/portals/hotspot/external", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &HotspotServiceCreateExternalPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	HotspotServiceCreateInternalPostRequestLocation struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	HotspotServiceCreateInternalPostRequestRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	HotspotServiceCreateInternalPostRequestUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	HotspotServiceCreateInternalPostRequestWalledGardensSlice []*string

	HotspotServiceCreateInternalPostRequest struct {
		Description        *string                                                   `json:"description,omitempty"`
		Location           *HotspotServiceCreateInternalPostRequestLocation          `json:"location,omitempty"`
		MacAddressFormat   *float64                                                  `json:"macAddressFormat,omitempty"`
		Name               *string                                                   `json:"name,omitempty"`
		Redirect           *HotspotServiceCreateInternalPostRequestRedirect          `json:"redirect,omitempty"`
		SmartClientSupport *string                                                   `json:"smartClientSupport,omitempty"`
		UserSession        *HotspotServiceCreateInternalPostRequestUserSession       `json:"userSession,omitempty"`
		WalledGardens      HotspotServiceCreateInternalPostRequestWalledGardensSlice `json:"walledGardens,omitempty"`
	}

	HotspotServiceCreateInternalPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// HotspotServiceCreateInternalPost: Use this API command to create a new Hotspot(WISPr) with internal logon URL of a zone.MacAddressFormat
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *HotspotServiceCreateInternalPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *HotspotServiceCreateInternalPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceCreateInternalPost(ctx context.Context, zoneId string, requestBody *HotspotServiceCreateInternalPostRequest) (*http.Response, *HotspotServiceCreateInternalPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/portals/hotspot/internal", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &HotspotServiceCreateInternalPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	HotspotServiceCreateSmartClientOnlyPostRequestLocation struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	HotspotServiceCreateSmartClientOnlyPostRequestRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	HotspotServiceCreateSmartClientOnlyPostRequestUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	HotspotServiceCreateSmartClientOnlyPostRequestWalledGardensSlice []*string

	HotspotServiceCreateSmartClientOnlyPostRequest struct {
		Description      *string                                                          `json:"description,omitempty"`
		Location         *HotspotServiceCreateSmartClientOnlyPostRequestLocation          `json:"location,omitempty"`
		MacAddressFormat *float64                                                         `json:"macAddressFormat,omitempty"`
		Name             *string                                                          `json:"name,omitempty"`
		Redirect         *HotspotServiceCreateSmartClientOnlyPostRequestRedirect          `json:"redirect,omitempty"`
		SmartClientInfo  *string                                                          `json:"smartClientInfo,omitempty"`
		UserSession      *HotspotServiceCreateSmartClientOnlyPostRequestUserSession       `json:"userSession,omitempty"`
		WalledGardens    HotspotServiceCreateSmartClientOnlyPostRequestWalledGardensSlice `json:"walledGardens,omitempty"`
	}

	HotspotServiceCreateSmartClientOnlyPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// HotspotServiceCreateSmartClientOnlyPost: Use this API command to create a new Hotspot(WISPr) with smart client only of a zone.MacAddressFormat
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *HotspotServiceCreateSmartClientOnlyPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *HotspotServiceCreateSmartClientOnlyPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceCreateSmartClientOnlyPost(ctx context.Context, zoneId string, requestBody *HotspotServiceCreateSmartClientOnlyPostRequest) (*http.Response, *HotspotServiceCreateSmartClientOnlyPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/portals/hotspot/smartClientOnly", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &HotspotServiceCreateSmartClientOnlyPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// HotspotServiceDeleteDelete: Use this API command to delete a Hotspot(WISPr) of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	HotspotServiceRetrieveGet200ResponseLocation struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	HotspotServiceRetrieveGet200ResponsePortalCustomization struct {
		Language                   *string `json:"language,omitempty"`
		Logo                       *string `json:"logo,omitempty"`
		TermsAndConditionsRequired *bool   `json:"termsAndConditionsRequired,omitempty"`
		TermsAndConditionsText     *string `json:"termsAndConditionsText,omitempty"`
		Title                      *string `json:"title,omitempty"`
	}

	HotspotServiceRetrieveGet200ResponseRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	HotspotServiceRetrieveGet200ResponseUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	HotspotServiceRetrieveGet200ResponseWalledGardensSlice []*string

	HotspotServiceRetrieveGet200Response struct {
		Description         *string                                                  `json:"description,omitempty"`
		ID                  *string                                                  `json:"id,omitempty"`
		Location            *HotspotServiceRetrieveGet200ResponseLocation            `json:"location,omitempty"`
		MacAddressFormat    *float64                                                 `json:"macAddressFormat,omitempty"`
		Name                *string                                                  `json:"name,omitempty"`
		PortalCustomization *HotspotServiceRetrieveGet200ResponsePortalCustomization `json:"portalCustomization,omitempty"`
		PortalType          *string                                                  `json:"portalType,omitempty"`
		PortalURL           *string                                                  `json:"portalUrl,omitempty"`
		Redirect            *HotspotServiceRetrieveGet200ResponseRedirect            `json:"redirect,omitempty"`
		SmartClientInfo     *string                                                  `json:"smartClientInfo,omitempty"`
		SmartClientSupport  *string                                                  `json:"smartClientSupport,omitempty"`
		UserSession         *HotspotServiceRetrieveGet200ResponseUserSession         `json:"userSession,omitempty"`
		WalledGardens       HotspotServiceRetrieveGet200ResponseWalledGardensSlice   `json:"walledGardens,omitempty"`
		ZoneID              *string                                                  `json:"zoneId,omitempty"`
	}
)

// HotspotServiceRetrieveGet: Use this API command to retrieve a Hotspot(WISPr) of zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *HotspotServiceRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *HotspotServiceRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &HotspotServiceRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	HotspotServiceModifyBasicPatchRequest struct {
		Description        *string  `json:"description,omitempty"`
		MacAddressFormat   *float64 `json:"macAddressFormat,omitempty"`
		Name               *string  `json:"name,omitempty"`
		PortalURL          *string  `json:"portalUrl,omitempty"`
		SmartClientInfo    *string  `json:"smartClientInfo,omitempty"`
		SmartClientSupport *string  `json:"smartClientSupport,omitempty"`
	}
)

// HotspotServiceModifyBasicPatch: Use this API command to modify the basic information on Hotspot(WISPr) of a zone.MacAddressFormat
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
// - requestBody: *HotspotServiceModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *HotspotServiceModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	HotspotServiceModifyLocationPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// HotspotServiceModifyLocationPatch: Use this API command to modify the location information on Hotspot(WISPr) of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
// - requestBody: *HotspotServiceModifyLocationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceModifyLocationPatch(ctx context.Context, zoneId string, id string, requestBody *HotspotServiceModifyLocationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}/location", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	HotspotServiceModifyPortalCustomizationPatchRequest struct {
		Language                   *string `json:"language,omitempty"`
		Logo                       *string `json:"logo,omitempty"`
		TermsAndConditionsRequired *bool   `json:"termsAndConditionsRequired,omitempty"`
		TermsAndConditionsText     *string `json:"termsAndConditionsText,omitempty"`
		Title                      *string `json:"title,omitempty"`
	}
)

// HotspotServiceModifyPortalCustomizationPatch: Use this API command to modify portal customization on Hotspot(WISPr) of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
// - requestBody: *HotspotServiceModifyPortalCustomizationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceModifyPortalCustomizationPatch(ctx context.Context, zoneId string, id string, requestBody *HotspotServiceModifyPortalCustomizationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}/portalCustomization", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	HotspotServiceModifyRedirectPatchRequest struct {
		URL *string `json:"url,omitempty"`
	}
)

// HotspotServiceModifyRedirectPatch: Use this API command to modify the redirect information on Hotspot(WISPr) of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
// - requestBody: *HotspotServiceModifyRedirectPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceModifyRedirectPatch(ctx context.Context, zoneId string, id string, requestBody *HotspotServiceModifyRedirectPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}/redirect", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	HotspotServiceModifyWalledGardensPatchRequest []*string
)

// HotspotServiceModifyWalledGardensPatch: Use this API command to modify walled gardens on Hotspot(WISPr) of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): HotSpot Profile ID
// - requestBody: *HotspotServiceModifyWalledGardensPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) HotspotServiceModifyWalledGardensPatch(ctx context.Context, zoneId string, id string, requestBody HotspotServiceModifyWalledGardensPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/hotspot/{id}/walledGardens", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WebAuthenticationRetrieveListGet200ResponseListSlice []*WebAuthenticationRetrieveListGet200ResponseList

	WebAuthenticationRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WebAuthenticationRetrieveListGet200Response struct {
		FirstIndex *int                                                 `json:"firstIndex,omitempty"`
		HasMore    *bool                                                `json:"hasMore,omitempty"`
		List       WebAuthenticationRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                 `json:"totalCount,omitempty"`
	}
)

// WebAuthenticationRetrieveListGet: Use this API command to retrieve a list of web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WebAuthenticationRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationRetrieveListGet(ctx context.Context, zoneId string) (*http.Response, *WebAuthenticationRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/webauth", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &WebAuthenticationRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WebAuthenticationCreatePostRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	WebAuthenticationCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WebAuthenticationCreatePost: Use this API command to create a new web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WebAuthenticationCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WebAuthenticationCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationCreatePost(ctx context.Context, zoneId string, requestBody *WebAuthenticationCreatePostRequest) (*http.Response, *WebAuthenticationCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/portals/webauth", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WebAuthenticationCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// WebAuthenticationDeleteDelete: Use this API command to delete an web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/webauth/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WebAuthenticationRetrieveGet200ResponseRedirect struct {
		URL *string `json:"url,omitempty"`
	}

	WebAuthenticationRetrieveGet200ResponseUserSession struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}

	WebAuthenticationRetrieveGet200Response struct {
		Description    *string                                             `json:"description,omitempty"`
		ID             *string                                             `json:"id,omitempty"`
		Name           *string                                             `json:"name,omitempty"`
		PortalLanguage *string                                             `json:"portalLanguage,omitempty"`
		Redirect       *WebAuthenticationRetrieveGet200ResponseRedirect    `json:"redirect,omitempty"`
		UserSession    *WebAuthenticationRetrieveGet200ResponseUserSession `json:"userSession,omitempty"`
		ZoneID         *string                                             `json:"zoneId,omitempty"`
	}
)

// WebAuthenticationRetrieveGet: Use this API command to retrieve a web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WebAuthenticationRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *WebAuthenticationRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/webauth/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &WebAuthenticationRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WebAuthenticationModifyBasicPatchRequest struct {
		Description    *string `json:"description,omitempty"`
		Name           *string `json:"name,omitempty"`
		PortalLanguage *string `json:"portalLanguage,omitempty"`
	}
)

// WebAuthenticationModifyBasicPatch: Use this API command to modify the basic information on web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *WebAuthenticationModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *WebAuthenticationModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/webauth/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WebAuthenticationRedirectToUrlUserVisitDelete: Use this API command to set redirect to the URL that user intends to visit on web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationRedirectToUrlUserVisitDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/webauth/{id}/redirect", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WebAuthenticationModifyRedirectPatchRequest struct {
		URL *string `json:"url,omitempty"`
	}
)

// WebAuthenticationModifyRedirectPatch: Use this API command to modify the redirect information on web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *WebAuthenticationModifyRedirectPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationModifyRedirectPatch(ctx context.Context, zoneId string, id string, requestBody *WebAuthenticationModifyRedirectPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/webauth/{id}/redirect", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WebAuthenticationModifyUsersessionPatchRequest struct {
		GracePeriodInMin *float64 `json:"gracePeriodInMin,omitempty"`
		TimeoutInMin     *float64 `json:"timeoutInMin,omitempty"`
	}
)

// WebAuthenticationModifyUsersessionPatch: Use this API command to modify the user session on web authentication of a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *WebAuthenticationModifyUsersessionPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WebAuthenticationModifyUsersessionPatch(ctx context.Context, zoneId string, id string, requestBody *WebAuthenticationModifyUsersessionPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/webauth/{id}/userSession", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WechatRetrieveListGet200ResponseListSlice []*WechatRetrieveListGet200ResponseList

	WechatRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WechatRetrieveListGet200Response struct {
		FirstIndex *int                                      `json:"firstIndex,omitempty"`
		HasMore    *bool                                     `json:"hasMore,omitempty"`
		List       WechatRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                      `json:"totalCount,omitempty"`
	}
)

// WechatRetrieveListGet: Use this API command to retrieve a list of wechat portal.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WechatRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *WechatRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/wechat", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &WechatRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WechatCreatePostRequestDnatPortMappingSlice []*WechatCreatePostRequestDnatPortMapping

	WechatCreatePostRequestDnatPortMapping struct {
		DestPort   *float64 `json:"destPort,omitempty"`
		SourcePort *float64 `json:"sourcePort,omitempty"`
	}

	WechatCreatePostRequestWhiteListSlice []*string

	WechatCreatePostRequest struct {
		AuthURL         *string                                     `json:"authUrl,omitempty"`
		BlackList       *string                                     `json:"blackList,omitempty"`
		Description     *string                                     `json:"description,omitempty"`
		DnatDestination *string                                     `json:"dnatDestination,omitempty"`
		DnatPortMapping WechatCreatePostRequestDnatPortMappingSlice `json:"dnatPortMapping,omitempty"`
		GracePeriod     *float64                                    `json:"gracePeriod,omitempty"`
		Name            *string                                     `json:"name,omitempty"`
		WhiteList       WechatCreatePostRequestWhiteListSlice       `json:"whiteList,omitempty"`
	}

	WechatCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WechatCreatePost: Use this API command to create wechat portal.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WechatCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WechatCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatCreatePost(ctx context.Context, zoneId string, requestBody *WechatCreatePostRequest) (*http.Response, *WechatCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/portals/wechat", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WechatCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// WechatDeleteDelete: Use this API command to delete wechat portal.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/portals/wechat/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WechatRetrieveGet200ResponseDnatPortMappingSlice []*WechatRetrieveGet200ResponseDnatPortMapping

	WechatRetrieveGet200ResponseDnatPortMapping struct {
		DestPort   *float64 `json:"destPort,omitempty"`
		SourcePort *float64 `json:"sourcePort,omitempty"`
	}

	WechatRetrieveGet200ResponseWhiteListSlice []*string

	WechatRetrieveGet200Response struct {
		AuthURL         *string                                          `json:"authUrl,omitempty"`
		BlackList       *string                                          `json:"blackList,omitempty"`
		Description     *string                                          `json:"description,omitempty"`
		DnatDestination *string                                          `json:"dnatDestination,omitempty"`
		DnatPortMapping WechatRetrieveGet200ResponseDnatPortMappingSlice `json:"dnatPortMapping,omitempty"`
		GracePeriod     *float64                                         `json:"gracePeriod,omitempty"`
		Name            *string                                          `json:"name,omitempty"`
		WhiteList       WechatRetrieveGet200ResponseWhiteListSlice       `json:"whiteList,omitempty"`
	}
)

// WechatRetrieveGet: Use this API command to retrieve wechat portal by ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WechatRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *WechatRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/portals/wechat/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &WechatRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WechatModifyBasicPatchRequest struct {
		AuthURL         *string  `json:"authUrl,omitempty"`
		BlackList       *string  `json:"blackList,omitempty"`
		Description     *string  `json:"description,omitempty"`
		DnatDestination *string  `json:"dnatDestination,omitempty"`
		GracePeriod     *float64 `json:"gracePeriod,omitempty"`
		Name            *string  `json:"name,omitempty"`
	}
)

// WechatModifyBasicPatch: Use this API command to modify the basic information of wechat portal.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *WechatModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *WechatModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/wechat/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WechatModifyDnatPortMappingPatchRequestSlice []*WechatModifyDnatPortMappingPatchRequest

	WechatModifyDnatPortMappingPatchRequest struct {
		DestPort   *float64 `json:"destPort,omitempty"`
		SourcePort *float64 `json:"sourcePort,omitempty"`
	}
)

// WechatModifyDnatPortMappingPatch: Use this API command to modify DNAT port mapping of a wechat profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *WechatModifyDnatPortMappingPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatModifyDnatPortMappingPatch(ctx context.Context, zoneId string, id string, requestBody WechatModifyDnatPortMappingPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/wechat/{id}/dnatPortMapping", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WechatModifyWhitelistPatchRequest []*string
)

// WechatModifyWhitelistPatch: Use this API command to modify whiteList of a wechat profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *WechatModifyWhitelistPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WechatModifyWhitelistPatch(ctx context.Context, zoneId string, id string, requestBody WechatModifyWhitelistPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/portals/wechat/{id}/whiteList", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	EthernetPortProfileRetrieveListEthernetPortPorfileGet200ResponseListSlice []*EthernetPortProfileRetrieveListEthernetPortPorfileGet200ResponseList

	EthernetPortProfileRetrieveListEthernetPortPorfileGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileRetrieveListEthernetPortPorfileGet200Response struct {
		FirstIndex *int                                                                      `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                     `json:"hasMore,omitempty"`
		List       EthernetPortProfileRetrieveListEthernetPortPorfileGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                      `json:"totalCount,omitempty"`
	}
)

// EthernetPortProfileRetrieveListEthernetPortPorfileGet: Retrieve a list of Ethernet Port Porfiles within a zone
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *EthernetPortProfileRetrieveListEthernetPortPorfileGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) EthernetPortProfileRetrieveListEthernetPortPorfileGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *EthernetPortProfileRetrieveListEthernetPortPorfileGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/profile/ethernetPort", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &EthernetPortProfileRetrieveListEthernetPortPorfileGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAccountingServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAccounting struct {
		EnableUseSCGasProxy *bool                                                                                       `json:"enableUseSCGasProxy,omitempty"`
		Server              *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAccountingServer `json:"server,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAuthenticationServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAuthentication struct {
		EnableUseSCGasProxy *bool                                                                                           `json:"enableUseSCGasProxy,omitempty"`
		Server              *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAuthenticationServer `json:"server,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticator struct {
		Accounting           *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAccounting     `json:"accounting,omitempty"`
		Authentication       *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticatorAuthentication `json:"authentication,omitempty"`
		DisabledAccounting   *bool                                                                                     `json:"disabledAccounting,omitempty"`
		MacAuthByPassEnabled *bool                                                                                     `json:"macAuthByPassEnabled,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XSupplicant struct {
		Password *string `json:"password,omitempty"`
		Type     *string `json:"type,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021X struct {
		Authenticator *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XAuthenticator `json:"authenticator,omitempty"`
		Supplicant    *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021XSupplicant    `json:"supplicant,omitempty"`
		Type          *string                                                                     `json:"type,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePostRequest struct {
		X8021X             *EthernetPortProfileCreateEthernetPortPorfilePostRequestX8021X `json:"_8021X,omitempty"`
		Description        *string                                                        `json:"description,omitempty"`
		DynamicVlanEnabled *bool                                                          `json:"dynamicVlanEnabled,omitempty"`
		GuestVlan          *float64                                                       `json:"guestVlan,omitempty"`
		Name               *string                                                        `json:"name,omitempty"`
		TunnelEnabled      *bool                                                          `json:"tunnelEnabled,omitempty"`
		Type               *string                                                        `json:"type,omitempty"`
		UntagID            *float64                                                       `json:"untagId,omitempty"`
		VlanMembers        *string                                                        `json:"vlanMembers,omitempty"`
	}

	EthernetPortProfileCreateEthernetPortPorfilePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// EthernetPortProfileCreateEthernetPortPorfilePost: Create a new Ethernet Port Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *EthernetPortProfileCreateEthernetPortPorfilePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *EthernetPortProfileCreateEthernetPortPorfilePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) EthernetPortProfileCreateEthernetPortPorfilePost(ctx context.Context, zoneId string, requestBody *EthernetPortProfileCreateEthernetPortPorfilePostRequest) (*http.Response, *EthernetPortProfileCreateEthernetPortPorfilePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/profile/ethernetPort", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &EthernetPortProfileCreateEthernetPortPorfilePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// EthernetPortProfileDeleteEthernetPortPorfileDelete: Delete Ethernet Port Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) EthernetPortProfileDeleteEthernetPortPorfileDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/profile/ethernetPort/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAccountingServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAccounting struct {
		EnableUseSCGasProxy *bool                                                                                            `json:"enableUseSCGasProxy,omitempty"`
		Server              *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAccountingServer `json:"server,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAuthenticationServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAuthentication struct {
		EnableUseSCGasProxy *bool                                                                                                `json:"enableUseSCGasProxy,omitempty"`
		Server              *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAuthenticationServer `json:"server,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticator struct {
		Accounting           *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAccounting     `json:"accounting,omitempty"`
		Authentication       *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticatorAuthentication `json:"authentication,omitempty"`
		DisabledAccounting   *bool                                                                                          `json:"disabledAccounting,omitempty"`
		MacAuthByPassEnabled *bool                                                                                          `json:"macAuthByPassEnabled,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XSupplicant struct {
		Password *string `json:"password,omitempty"`
		Type     *string `json:"type,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021X struct {
		Authenticator *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XAuthenticator `json:"authenticator,omitempty"`
		Supplicant    *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021XSupplicant    `json:"supplicant,omitempty"`
		Type          *string                                                                          `json:"type,omitempty"`
	}

	EthernetPortProfileRetrieveEthernetPortPorfileGet200Response struct {
		X8021X             *EthernetPortProfileRetrieveEthernetPortPorfileGet200ResponseX8021X `json:"_8021X,omitempty"`
		Description        *string                                                             `json:"description,omitempty"`
		DynamicVlanEnabled *bool                                                               `json:"dynamicVlanEnabled,omitempty"`
		GuestVlan          *float64                                                            `json:"guestVlan,omitempty"`
		ID                 *string                                                             `json:"id,omitempty"`
		Name               *string                                                             `json:"name,omitempty"`
		TunnelEnabled      *bool                                                               `json:"tunnelEnabled,omitempty"`
		Type               *string                                                             `json:"type,omitempty"`
		UntagID            *float64                                                            `json:"untagId,omitempty"`
		VlanMembers        *string                                                             `json:"vlanMembers,omitempty"`
	}
)

// EthernetPortProfileRetrieveEthernetPortPorfileGet: Retrieve a Ethernet Port Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *EthernetPortProfileRetrieveEthernetPortPorfileGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) EthernetPortProfileRetrieveEthernetPortPorfileGet(ctx context.Context, zoneId string, id string) (*http.Response, *EthernetPortProfileRetrieveEthernetPortPorfileGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/profile/ethernetPort/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &EthernetPortProfileRetrieveEthernetPortPorfileGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	EthernetPortProfileModifyEthernetPortPorfilePatchRequest struct {
		Description        *string  `json:"description,omitempty"`
		DynamicVlanEnabled *bool    `json:"dynamicVlanEnabled,omitempty"`
		GuestVlan          *float64 `json:"guestVlan,omitempty"`
		Name               *string  `json:"name,omitempty"`
		TunnelEnabled      *bool    `json:"tunnelEnabled,omitempty"`
		UntagID            *float64 `json:"untagId,omitempty"`
		VlanMembers        *string  `json:"vlanMembers,omitempty"`
	}
)

// EthernetPortProfileModifyEthernetPortPorfilePatch: Modify a specific Ethernet Port Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *EthernetPortProfileModifyEthernetPortPorfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) EthernetPortProfileModifyEthernetPortPorfilePatch(ctx context.Context, zoneId string, id string, requestBody *EthernetPortProfileModifyEthernetPortPorfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/profile/ethernetPort/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAccountingServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAccounting struct {
		EnableUseSCGasProxy *bool                                                                                         `json:"enableUseSCGasProxy,omitempty"`
		Server              *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAccountingServer `json:"server,omitempty"`
	}

	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAuthenticationServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAuthentication struct {
		EnableUseSCGasProxy *bool                                                                                             `json:"enableUseSCGasProxy,omitempty"`
		Server              *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAuthenticationServer `json:"server,omitempty"`
	}

	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticator struct {
		Accounting           *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAccounting     `json:"accounting,omitempty"`
		Authentication       *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticatorAuthentication `json:"authentication,omitempty"`
		DisabledAccounting   *bool                                                                                       `json:"disabledAccounting,omitempty"`
		MacAuthByPassEnabled *bool                                                                                       `json:"macAuthByPassEnabled,omitempty"`
	}

	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestSupplicant struct {
		Password *string `json:"password,omitempty"`
		Type     *string `json:"type,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}

	EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequest struct {
		Authenticator *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestAuthenticator `json:"authenticator,omitempty"`
		Supplicant    *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequestSupplicant    `json:"supplicant,omitempty"`
		Type          *string                                                                       `json:"type,omitempty"`
	}
)

// EthernetPortProfileModify8021xOfEthernetPortPorfilePatch: Modify _8021X of Ethernet Port Porfile
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
// - requestBody: *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) EthernetPortProfileModify8021xOfEthernetPortPorfilePatch(ctx context.Context, zoneId string, id string, requestBody *EthernetPortProfileModify8021xOfEthernetPortPorfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/profile/ethernetPort/{id}/_8021X", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApUsbSoftwarePackageRetrieveListGet200ResponseListSlice []*ApUsbSoftwarePackageRetrieveListGet200ResponseList

	ApUsbSoftwarePackageRetrieveListGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApUsbSoftwarePackageRetrieveListGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       ApUsbSoftwarePackageRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// ApUsbSoftwarePackageRetrieveListGet: Retrieve a list of AP Usb Software Package
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApUsbSoftwarePackageRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApUsbSoftwarePackageRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *ApUsbSoftwarePackageRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/usbsoftware", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &ApUsbSoftwarePackageRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApUsbSoftwarePackageUploadFilePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ApUsbSoftwarePackageUploadFilePost: Create new AP Usb Software Package by upload file
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApUsbSoftwarePackageUploadFilePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApUsbSoftwarePackageUploadFilePost(ctx context.Context, zoneId string) (*http.Response, *ApUsbSoftwarePackageUploadFilePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/usbsoftware", true)
	request.SetPathParameter("zoneId", zoneId)
	out := &ApUsbSoftwarePackageUploadFilePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ApUsbSoftwarePackageDeleteDelete: Delete specified AP Usb Software Package
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApUsbSoftwarePackageDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/usbsoftware/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	ApUsbSoftwarePackageGetZoneAssociateGet200Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApUsbSoftwarePackageGetZoneAssociateGet: Get APUsbSoftwarePackage associate with zone by model name
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - modelName (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApUsbSoftwarePackageGetZoneAssociateGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) ApUsbSoftwarePackageGetZoneAssociateGet(ctx context.Context, zoneId string, modelName string) (*http.Response, *ApUsbSoftwarePackageGetZoneAssociateGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(modelName)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"modelName\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/usbsoftware/{modelName}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("modelName", modelName)
	out := &ApUsbSoftwarePackageGetZoneAssociateGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanGroupRetrieveListGet200ResponseListSlice []*WlanGroupRetrieveListGet200ResponseList

	WlanGroupRetrieveListGet200ResponseListMembersSlice []*WlanGroupRetrieveListGet200ResponseListMembers

	WlanGroupRetrieveListGet200ResponseListMembersVlanPooling struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanGroupRetrieveListGet200ResponseListMembers struct {
		AccessVlan  *float64                                                   `json:"accessVlan,omitempty"`
		ID          *string                                                    `json:"id,omitempty"`
		NasID       *string                                                    `json:"nasId,omitempty"`
		VlanPooling *WlanGroupRetrieveListGet200ResponseListMembersVlanPooling `json:"vlanPooling,omitempty"`
	}

	WlanGroupRetrieveListGet200ResponseList struct {
		CreateDateTime   *int                                                `json:"createDateTime,omitempty"`
		CreatorID        *string                                             `json:"creatorId,omitempty"`
		CreatorUsername  *string                                             `json:"creatorUsername,omitempty"`
		Description      *string                                             `json:"description,omitempty"`
		ID               *string                                             `json:"id,omitempty"`
		Members          WlanGroupRetrieveListGet200ResponseListMembersSlice `json:"members,omitempty"`
		ModifiedDateTime *int                                                `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                             `json:"modifierId,omitempty"`
		ModifierUsername *string                                             `json:"modifierUsername,omitempty"`
		Name             *string                                             `json:"name,omitempty"`
		ZoneID           *string                                             `json:"zoneId,omitempty"`
	}

	WlanGroupRetrieveListGet200Response struct {
		FirstIndex *int                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                        `json:"hasMore,omitempty"`
		List       WlanGroupRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                         `json:"totalCount,omitempty"`
	}
)

// WlanGroupRetrieveListGet: Use this API command to retrieve the list of WLAN groups within a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanGroupRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *WlanGroupRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlangroups", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &WlanGroupRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanGroupCreatePostRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	WlanGroupCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanGroupCreatePost: Use this API command to create a new WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanGroupCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanGroupCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupCreatePost(ctx context.Context, zoneId string, requestBody *WlanGroupCreatePostRequest) (*http.Response, *WlanGroupCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlangroups", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanGroupCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// WlanGroupDeleteDelete: Use this API command to delete a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlangroups/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanGroupRetrieveGet200ResponseMembersSlice []*WlanGroupRetrieveGet200ResponseMembers

	WlanGroupRetrieveGet200ResponseMembersVlanPooling struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanGroupRetrieveGet200ResponseMembers struct {
		AccessVlan  *float64                                           `json:"accessVlan,omitempty"`
		ID          *string                                            `json:"id,omitempty"`
		NasID       *string                                            `json:"nasId,omitempty"`
		VlanPooling *WlanGroupRetrieveGet200ResponseMembersVlanPooling `json:"vlanPooling,omitempty"`
	}

	WlanGroupRetrieveGet200Response struct {
		CreateDateTime   *int                                        `json:"createDateTime,omitempty"`
		CreatorID        *string                                     `json:"creatorId,omitempty"`
		CreatorUsername  *string                                     `json:"creatorUsername,omitempty"`
		Description      *string                                     `json:"description,omitempty"`
		ID               *string                                     `json:"id,omitempty"`
		Members          WlanGroupRetrieveGet200ResponseMembersSlice `json:"members,omitempty"`
		ModifiedDateTime *int                                        `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                     `json:"modifierId,omitempty"`
		ModifierUsername *string                                     `json:"modifierUsername,omitempty"`
		Name             *string                                     `json:"name,omitempty"`
		ZoneID           *string                                     `json:"zoneId,omitempty"`
	}
)

// WlanGroupRetrieveGet: Use this API command to retrieve the WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanGroupRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *WlanGroupRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlangroups/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &WlanGroupRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanGroupModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// WlanGroupModifyBasicPatch: Use this API command to modify the basic information of a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - requestBody: *WlanGroupModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *WlanGroupModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlangroups/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanGroupAddMemberPostRequestVlanPooling struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanGroupAddMemberPostRequest struct {
		AccessVlan  *float64                                  `json:"accessVlan,omitempty"`
		ID          *string                                   `json:"id,omitempty"`
		NasID       *string                                   `json:"nasId,omitempty"`
		VlanPooling *WlanGroupAddMemberPostRequestVlanPooling `json:"vlanPooling,omitempty"`
	}
)

// WlanGroupAddMemberPost: Use this API command to add a member to a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - requestBody: *WlanGroupAddMemberPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupAddMemberPost(ctx context.Context, zoneId string, id string, requestBody *WlanGroupAddMemberPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlangroups/{id}/members", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 201, nil)
}

// WlanGroupRemoveMemberDelete: Use this API command to remove a member from a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - memberId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupRemoveMemberDelete(ctx context.Context, zoneId string, id string, memberId string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(memberId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"memberId\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlangroups/{id}/members/{memberId}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("memberId", memberId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanGroupModifyMemberPatchRequest struct {
		AccessVlan *float64 `json:"accessVlan,omitempty"`
		NasID      *string  `json:"nasId,omitempty"`
	}
)

// WlanGroupModifyMemberPatch: Use this API command to modify a member of a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - memberId (string)
// - requestBody: *WlanGroupModifyMemberPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupModifyMemberPatch(ctx context.Context, zoneId string, id string, memberId string, requestBody *WlanGroupModifyMemberPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(memberId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"memberId\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlangroups/{id}/members/{memberId}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("memberId", memberId)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanGroupDisableMemberNasOverrideDelete: Use this API command to disable a member NAS-ID override of a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - memberId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupDisableMemberNasOverrideDelete(ctx context.Context, zoneId string, id string, memberId string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(memberId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"memberId\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlangroups/{id}/members/{memberId}/nasId", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("memberId", memberId)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanGroupDisableMemberVlanOverrideDelete: Use this API command to disable a member VLAN override of a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - memberId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupDisableMemberVlanOverrideDelete(ctx context.Context, zoneId string, id string, memberId string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(memberId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"memberId\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlangroups/{id}/members/{memberId}/vlanOverride", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("memberId", memberId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanGroupModifyMemberVlanPoolingPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanGroupModifyMemberVlanPoolingPatch: Use this API command to modify a member's VLAN pooling of a WLAN group.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Group ID
// - memberId (string)
// - requestBody: *WlanGroupModifyMemberVlanPoolingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanGroupModifyMemberVlanPoolingPatch(ctx context.Context, zoneId string, id string, memberId string, requestBody *WlanGroupModifyMemberVlanPoolingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(memberId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"memberId\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlangroups/{id}/members/{memberId}/vlanPooling", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("memberId", memberId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanRetrieveListGet200ResponseListSlice []*WlanRetrieveListGet200ResponseList

	WlanRetrieveListGet200ResponseList struct {
		ID     *string `json:"id,omitempty"`
		MvnoID *string `json:"mvnoId,omitempty"`
		Name   *string `json:"name,omitempty"`
		Ssid   *string `json:"ssid,omitempty"`
		ZoneID *string `json:"zoneId,omitempty"`
	}

	WlanRetrieveListGet200Response struct {
		FirstIndex *int                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                   `json:"hasMore,omitempty"`
		List       WlanRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                    `json:"totalCount,omitempty"`
	}
)

// WlanRetrieveListGet: Use this API command to retrieve a list of WLANs within a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *WlanRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlans", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &WlanRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanCreateStandardOpenPostRequest struct {
		Description *string `json:"description,omitempty"`
		Hessid      *string `json:"hessid,omitempty"`
		Name        *string `json:"name,omitempty"`
		Ssid        *string `json:"ssid,omitempty"`
	}

	WlanCreateStandardOpenPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateStandardOpenPost: Use this API command to create a new standard, open and non-tunneled basic WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateStandardOpenPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateStandardOpenPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateStandardOpenPost(ctx context.Context, zoneId string, requestBody *WlanCreateStandardOpenPostRequest) (*http.Response, *WlanCreateStandardOpenPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateStandardOpenPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanSchedulerRetrieveListGet200ResponseListSlice []*WlanSchedulerRetrieveListGet200ResponseList

	WlanSchedulerRetrieveListGet200ResponseListFriSlice []*string

	WlanSchedulerRetrieveListGet200ResponseListMonSlice []*string

	WlanSchedulerRetrieveListGet200ResponseListSatSlice []*string

	WlanSchedulerRetrieveListGet200ResponseListSunSlice []*string

	WlanSchedulerRetrieveListGet200ResponseListThuSlice []*string

	WlanSchedulerRetrieveListGet200ResponseListTueSlice []*string

	WlanSchedulerRetrieveListGet200ResponseListWedSlice []*string

	WlanSchedulerRetrieveListGet200ResponseList struct {
		Description *string                                             `json:"description,omitempty"`
		Fri         WlanSchedulerRetrieveListGet200ResponseListFriSlice `json:"fri,omitempty"`
		ID          *string                                             `json:"id,omitempty"`
		Mon         WlanSchedulerRetrieveListGet200ResponseListMonSlice `json:"mon,omitempty"`
		Name        *string                                             `json:"name,omitempty"`
		Sat         WlanSchedulerRetrieveListGet200ResponseListSatSlice `json:"sat,omitempty"`
		Sun         WlanSchedulerRetrieveListGet200ResponseListSunSlice `json:"sun,omitempty"`
		Thu         WlanSchedulerRetrieveListGet200ResponseListThuSlice `json:"thu,omitempty"`
		Tue         WlanSchedulerRetrieveListGet200ResponseListTueSlice `json:"tue,omitempty"`
		Wed         WlanSchedulerRetrieveListGet200ResponseListWedSlice `json:"wed,omitempty"`
		ZoneID      *string                                             `json:"zoneId,omitempty"`
	}

	WlanSchedulerRetrieveListGet200Response struct {
		FirstIndex *int                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                            `json:"hasMore,omitempty"`
		List       WlanSchedulerRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                             `json:"totalCount,omitempty"`
	}
)

// WlanSchedulerRetrieveListGet: Use this API command to retrieve the list of WLAN schedule from a zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanSchedulerRetrieveListGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanSchedulerRetrieveListGet(ctx context.Context, zoneId string, optionalParams map[string]string) (*http.Response, *WlanSchedulerRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
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
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlanSchedulers", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &WlanSchedulerRetrieveListGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanSchedulerCreatePostRequestFriSlice []*string

	WlanSchedulerCreatePostRequestMonSlice []*string

	WlanSchedulerCreatePostRequestSatSlice []*string

	WlanSchedulerCreatePostRequestSunSlice []*string

	WlanSchedulerCreatePostRequestThuSlice []*string

	WlanSchedulerCreatePostRequestTueSlice []*string

	WlanSchedulerCreatePostRequestWedSlice []*string

	WlanSchedulerCreatePostRequest struct {
		Description *string                                `json:"description,omitempty"`
		Fri         WlanSchedulerCreatePostRequestFriSlice `json:"fri,omitempty"`
		Mon         WlanSchedulerCreatePostRequestMonSlice `json:"mon,omitempty"`
		Name        *string                                `json:"name,omitempty"`
		Sat         WlanSchedulerCreatePostRequestSatSlice `json:"sat,omitempty"`
		Sun         WlanSchedulerCreatePostRequestSunSlice `json:"sun,omitempty"`
		Thu         WlanSchedulerCreatePostRequestThuSlice `json:"thu,omitempty"`
		Tue         WlanSchedulerCreatePostRequestTueSlice `json:"tue,omitempty"`
		Wed         WlanSchedulerCreatePostRequestWedSlice `json:"wed,omitempty"`
	}

	WlanSchedulerCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanSchedulerCreatePost: Use this API command to create a new WLAN schedule.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanSchedulerCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanSchedulerCreatePost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanSchedulerCreatePost(ctx context.Context, zoneId string, requestBody *WlanSchedulerCreatePostRequest) (*http.Response, *WlanSchedulerCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlanSchedulers", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanSchedulerCreatePost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// WlanSchedulerDeleteDelete: Use this API command to delete a WLAN schedule.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Schedule ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanSchedulerDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlanSchedulers/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanSchedulerRetrieveGet200ResponseFriSlice []*string

	WlanSchedulerRetrieveGet200ResponseMonSlice []*string

	WlanSchedulerRetrieveGet200ResponseSatSlice []*string

	WlanSchedulerRetrieveGet200ResponseSunSlice []*string

	WlanSchedulerRetrieveGet200ResponseThuSlice []*string

	WlanSchedulerRetrieveGet200ResponseTueSlice []*string

	WlanSchedulerRetrieveGet200ResponseWedSlice []*string

	WlanSchedulerRetrieveGet200Response struct {
		Description *string                                     `json:"description,omitempty"`
		Fri         WlanSchedulerRetrieveGet200ResponseFriSlice `json:"fri,omitempty"`
		ID          *string                                     `json:"id,omitempty"`
		Mon         WlanSchedulerRetrieveGet200ResponseMonSlice `json:"mon,omitempty"`
		Name        *string                                     `json:"name,omitempty"`
		Sat         WlanSchedulerRetrieveGet200ResponseSatSlice `json:"sat,omitempty"`
		Sun         WlanSchedulerRetrieveGet200ResponseSunSlice `json:"sun,omitempty"`
		Thu         WlanSchedulerRetrieveGet200ResponseThuSlice `json:"thu,omitempty"`
		Tue         WlanSchedulerRetrieveGet200ResponseTueSlice `json:"tue,omitempty"`
		Wed         WlanSchedulerRetrieveGet200ResponseWedSlice `json:"wed,omitempty"`
		ZoneID      *string                                     `json:"zoneId,omitempty"`
	}
)

// WlanSchedulerRetrieveGet: Use this API command to retrieve a WLAN schedule.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Schedule ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanSchedulerRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanSchedulerRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *WlanSchedulerRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlanSchedulers/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &WlanSchedulerRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanSchedulerModifyPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// WlanSchedulerModifyPatch: Use this API command to modify a WLAN schedule.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (UUIDv4): WLAN Schedule ID
// - requestBody: *WlanSchedulerModifyPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanSchedulerModifyPatch(ctx context.Context, zoneId string, id string, requestBody *WlanSchedulerModifyPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlanSchedulers/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanCreateGuestAccessPostRequestAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanCreateGuestAccessPostRequestPortalServiceProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanCreateGuestAccessPostRequest struct {
		AuthServiceOrProfile *WlanCreateGuestAccessPostRequestAuthServiceOrProfile `json:"authServiceOrProfile,omitempty"`
		Description          *string                                               `json:"description,omitempty"`
		Hessid               *string                                               `json:"hessid,omitempty"`
		Name                 *string                                               `json:"name,omitempty"`
		PortalServiceProfile *WlanCreateGuestAccessPostRequestPortalServiceProfile `json:"portalServiceProfile,omitempty"`
		Ssid                 *string                                               `json:"ssid,omitempty"`
	}

	WlanCreateGuestAccessPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateGuestAccessPost: Use this API command to create a new guest access WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateGuestAccessPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateGuestAccessPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateGuestAccessPost(ctx context.Context, zoneId string, requestBody *WlanCreateGuestAccessPostRequest) (*http.Response, *WlanCreateGuestAccessPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/guest", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateGuestAccessPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateHotspot20PostRequestHotspot20Profile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanCreateHotspot20PostRequest struct {
		Description      *string                                         `json:"description,omitempty"`
		Hessid           *string                                         `json:"hessid,omitempty"`
		Hotspot20Profile *WlanCreateHotspot20PostRequestHotspot20Profile `json:"hotspot20Profile,omitempty"`
		Name             *string                                         `json:"name,omitempty"`
		Ssid             *string                                         `json:"ssid,omitempty"`
	}

	WlanCreateHotspot20Post201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateHotspot20Post: Use this API command to create a new Hotspot 2.0 access WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateHotspot20PostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateHotspot20Post201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateHotspot20Post(ctx context.Context, zoneId string, requestBody *WlanCreateHotspot20PostRequest) (*http.Response, *WlanCreateHotspot20Post201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/hotspot20", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateHotspot20Post201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateHotspot20OsenPostRequest struct {
		Description *string `json:"description,omitempty"`
		Hessid      *string `json:"hessid,omitempty"`
		Name        *string `json:"name,omitempty"`
		Ssid        *string `json:"ssid,omitempty"`
	}

	WlanCreateHotspot20OsenPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateHotspot20OsenPost: Use this API command to create a new Hotspot 2.0 Secure Online Signup WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateHotspot20OsenPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateHotspot20OsenPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateHotspot20OsenPost(ctx context.Context, zoneId string, requestBody *WlanCreateHotspot20OsenPostRequest) (*http.Response, *WlanCreateHotspot20OsenPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/hotspot20osen", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateHotspot20OsenPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreate8021xPostRequestAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanCreate8021xPostRequest struct {
		AuthServiceOrProfile *WlanCreate8021xPostRequestAuthServiceOrProfile `json:"authServiceOrProfile,omitempty"`
		Description          *string                                         `json:"description,omitempty"`
		Hessid               *string                                         `json:"hessid,omitempty"`
		Name                 *string                                         `json:"name,omitempty"`
		Ssid                 *string                                         `json:"ssid,omitempty"`
	}

	WlanCreate8021xPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreate8021xPost: Use this API command to create a new standard, 802.1X and non-tunneled WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreate8021xPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreate8021xPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreate8021xPost(ctx context.Context, zoneId string, requestBody *WlanCreate8021xPostRequest) (*http.Response, *WlanCreate8021xPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/standard80211", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreate8021xPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateMacAuthPostRequestAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanCreateMacAuthPostRequest struct {
		AuthServiceOrProfile *WlanCreateMacAuthPostRequestAuthServiceOrProfile `json:"authServiceOrProfile,omitempty"`
		Description          *string                                           `json:"description,omitempty"`
		Hessid               *string                                           `json:"hessid,omitempty"`
		Name                 *string                                           `json:"name,omitempty"`
		Ssid                 *string                                           `json:"ssid,omitempty"`
	}

	WlanCreateMacAuthPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateMacAuthPost: Use this API command to create a new standard, MAC auth and non-tunneled WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateMacAuthPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateMacAuthPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateMacAuthPost(ctx context.Context, zoneId string, requestBody *WlanCreateMacAuthPostRequest) (*http.Response, *WlanCreateMacAuthPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/standardmac", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateMacAuthPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateWebAuthPostRequestAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanCreateWebAuthPostRequestPortalServiceProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanCreateWebAuthPostRequest struct {
		AuthServiceOrProfile *WlanCreateWebAuthPostRequestAuthServiceOrProfile `json:"authServiceOrProfile,omitempty"`
		Description          *string                                           `json:"description,omitempty"`
		Hessid               *string                                           `json:"hessid,omitempty"`
		Name                 *string                                           `json:"name,omitempty"`
		PortalServiceProfile *WlanCreateWebAuthPostRequestPortalServiceProfile `json:"portalServiceProfile,omitempty"`
		Ssid                 *string                                           `json:"ssid,omitempty"`
	}

	WlanCreateWebAuthPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateWebAuthPost: Use this API command to creates new web authentication WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateWebAuthPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateWebAuthPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateWebAuthPost(ctx context.Context, zoneId string, requestBody *WlanCreateWebAuthPostRequest) (*http.Response, *WlanCreateWebAuthPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/webauth", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateWebAuthPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateWechatPostRequestPortalServiceProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanCreateWechatPostRequest struct {
		Description          *string                                          `json:"description,omitempty"`
		Hessid               *string                                          `json:"hessid,omitempty"`
		Name                 *string                                          `json:"name,omitempty"`
		PortalServiceProfile *WlanCreateWechatPostRequestPortalServiceProfile `json:"portalServiceProfile,omitempty"`
		Ssid                 *string                                          `json:"ssid,omitempty"`
	}

	WlanCreateWechatPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateWechatPost: Use this API command to create a new wechat WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateWechatPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateWechatPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateWechatPost(ctx context.Context, zoneId string, requestBody *WlanCreateWechatPostRequest) (*http.Response, *WlanCreateWechatPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/wechat", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateWechatPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateHotspotPostRequestAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanCreateHotspotPostRequestPortalServiceProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanCreateHotspotPostRequest struct {
		AuthServiceOrProfile *WlanCreateHotspotPostRequestAuthServiceOrProfile `json:"authServiceOrProfile,omitempty"`
		Description          *string                                           `json:"description,omitempty"`
		Hessid               *string                                           `json:"hessid,omitempty"`
		Name                 *string                                           `json:"name,omitempty"`
		PortalServiceProfile *WlanCreateHotspotPostRequestPortalServiceProfile `json:"portalServiceProfile,omitempty"`
		Ssid                 *string                                           `json:"ssid,omitempty"`
	}

	WlanCreateHotspotPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateHotspotPost: Use this API command to create new hotspot (WISPr) WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateHotspotPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateHotspotPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateHotspotPost(ctx context.Context, zoneId string, requestBody *WlanCreateHotspotPostRequest) (*http.Response, *WlanCreateHotspotPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/wispr", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateHotspotPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	WlanCreateHotspotMacBypassPostRequestAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanCreateHotspotMacBypassPostRequestPortalServiceProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanCreateHotspotMacBypassPostRequest struct {
		AuthServiceOrProfile *WlanCreateHotspotMacBypassPostRequestAuthServiceOrProfile `json:"authServiceOrProfile,omitempty"`
		Description          *string                                                    `json:"description,omitempty"`
		Hessid               *string                                                    `json:"hessid,omitempty"`
		Name                 *string                                                    `json:"name,omitempty"`
		PortalServiceProfile *WlanCreateHotspotMacBypassPostRequestPortalServiceProfile `json:"portalServiceProfile,omitempty"`
		Ssid                 *string                                                    `json:"ssid,omitempty"`
	}

	WlanCreateHotspotMacBypassPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// WlanCreateHotspotMacBypassPost: Use this API command to create a new hotspot (WISPr) with MAC bypass WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - requestBody: *WlanCreateHotspotMacBypassPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanCreateHotspotMacBypassPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanCreateHotspotMacBypassPost(ctx context.Context, zoneId string, requestBody *WlanCreateHotspotMacBypassPostRequest) (*http.Response, *WlanCreateHotspotMacBypassPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/wisprmac", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	out := &WlanCreateHotspotMacBypassPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// WlanDeleteDelete: Use this API command to delete a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDeleteDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanRetrieveGet200ResponseAccountingServiceOrProfile struct {
		AccountingDelayEnabled *bool    `json:"accountingDelayEnabled,omitempty"`
		ID                     *string  `json:"id,omitempty"`
		InterimUpdateMin       *float64 `json:"interimUpdateMin,omitempty"`
		Name                   *string  `json:"name,omitempty"`
		ThroughController      *bool    `json:"throughController,omitempty"`
	}

	WlanRetrieveGet200ResponseAdvancedOptionsClientIsolationWhitelist struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseAdvancedOptions struct {
		AvcEnabled                  *bool                                                              `json:"avcEnabled,omitempty"`
		BandBalancing               *string                                                            `json:"bandBalancing,omitempty"`
		BssMinRateMbps              *string                                                            `json:"bssMinRateMbps,omitempty"`
		ClientFingerprintingEnabled *bool                                                              `json:"clientFingerprintingEnabled,omitempty"`
		ClientIdleTimeoutSec        *float64                                                           `json:"clientIdleTimeoutSec,omitempty"`
		ClientIsolationEnabled      *bool                                                              `json:"clientIsolationEnabled,omitempty"`
		ClientIsolationWhitelist    *WlanRetrieveGet200ResponseAdvancedOptionsClientIsolationWhitelist `json:"clientIsolationWhitelist,omitempty"`
		ClientLoadBalancingEnabled  *bool                                                              `json:"clientLoadBalancingEnabled,omitempty"`
		DgafEnabled                 *bool                                                              `json:"dgafEnabled,omitempty"`
		Dhcp82Format                *string                                                            `json:"dhcp82Format,omitempty"`
		DhcpOption82Enabled         *bool                                                              `json:"dhcpOption82Enabled,omitempty"`
		DownlinkEnabled             *bool                                                              `json:"downlinkEnabled,omitempty"`
		DownlinkRate                *float64                                                           `json:"downlinkRate,omitempty"`
		DownlinkRateLimiting        *string                                                            `json:"downlinkRateLimiting,omitempty"`
		ForceClientDHCPTimeoutSec   *float64                                                           `json:"forceClientDHCPTimeoutSec,omitempty"`
		ForceDHCPEnabled            *bool                                                              `json:"forceDHCPEnabled,omitempty"`
		HideSsidEnabled             *bool                                                              `json:"hideSsidEnabled,omitempty"`
		MaxClientsPerRadio          *float64                                                           `json:"maxClientsPerRadio,omitempty"`
		MgmtTxRateMbps              *string                                                            `json:"mgmtTxRateMbps,omitempty"`
		OfdmOnlyEnabled             *bool                                                              `json:"ofdmOnlyEnabled,omitempty"`
		OkcEnabled                  *bool                                                              `json:"okcEnabled,omitempty"`
		PmkCachingEnabled           *bool                                                              `json:"pmkCachingEnabled,omitempty"`
		Priority                    *string                                                            `json:"priority,omitempty"`
		ProxyARPEnabled             *bool                                                              `json:"proxyARPEnabled,omitempty"`
		Support8021DEnabled         *bool                                                              `json:"support80211dEnabled,omitempty"`
		Support8021KEnabled         *bool                                                              `json:"support80211kEnabled,omitempty"`
		UnauthClientStatsEnabled    *bool                                                              `json:"unauthClientStatsEnabled,omitempty"`
		UplinkEnabled               *bool                                                              `json:"uplinkEnabled,omitempty"`
		UplinkRate                  *float64                                                           `json:"uplinkRate,omitempty"`
		UplinkRateLimiting          *string                                                            `json:"uplinkRateLimiting,omitempty"`
	}

	WlanRetrieveGet200ResponseAuthServiceOrProfile struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}

	WlanRetrieveGet200ResponseCoreTunnelProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Type *string `json:"type,omitempty"`
	}

	WlanRetrieveGet200ResponseDefaultUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseDevicePolicy struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseDiffServProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseDNSServerProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseDpsk struct {
		DpskEnabled  *bool    `json:"dpskEnabled,omitempty"`
		DpskFromType *string  `json:"dpskFromType,omitempty"`
		DpskType     *string  `json:"dpskType,omitempty"`
		Expiration   *string  `json:"expiration,omitempty"`
		Length       *float64 `json:"length,omitempty"`
	}

	WlanRetrieveGet200ResponseEncryption struct {
		Algorithm           *string  `json:"algorithm,omitempty"`
		KeyInHex            *string  `json:"keyInHex,omitempty"`
		KeyIndex            *float64 `json:"keyIndex,omitempty"`
		Method              *string  `json:"method,omitempty"`
		Mfp                 *string  `json:"mfp,omitempty"`
		MobilityDomainID    *float64 `json:"mobilityDomainId,omitempty"`
		Passphrase          *string  `json:"passphrase,omitempty"`
		Support8021REnabled *bool    `json:"support80211rEnabled,omitempty"`
	}

	WlanRetrieveGet200ResponseHotspot20Profile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseL2ACL struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseMacAuth struct {
		CustomizedPassword *string `json:"customizedPassword,omitempty"`
		MacAuthMacFormat   *string `json:"macAuthMacFormat,omitempty"`
	}

	WlanRetrieveGet200ResponsePortalServiceProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseQosMapsSlice []*WlanRetrieveGet200ResponseQosMaps

	WlanRetrieveGet200ResponseQosMapsExceptsSlice []*float64

	WlanRetrieveGet200ResponseQosMaps struct {
		Enable   *bool                                         `json:"enable,omitempty"`
		Excepts  WlanRetrieveGet200ResponseQosMapsExceptsSlice `json:"excepts,omitempty"`
		High     *int                                          `json:"high,omitempty"`
		Low      *int                                          `json:"low,omitempty"`
		Priority *int                                          `json:"priority,omitempty"`
	}

	WlanRetrieveGet200ResponseRadiusOptions struct {
		CalledStaIDType        *string  `json:"calledStaIdType,omitempty"`
		CustomizedNasID        *string  `json:"customizedNasId,omitempty"`
		NasIDType              *string  `json:"nasIdType,omitempty"`
		NasIPType              *string  `json:"nasIpType,omitempty"`
		NasIPUserDefined       *string  `json:"nasIpUserDefined,omitempty"`
		NasMaxRetry            *float64 `json:"nasMaxRetry,omitempty"`
		NasReconnectPrimaryMin *float64 `json:"nasReconnectPrimaryMin,omitempty"`
		NasRequestTimeoutSec   *float64 `json:"nasRequestTimeoutSec,omitempty"`
	}

	WlanRetrieveGet200ResponseSchedule struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Type *string `json:"type,omitempty"`
	}

	WlanRetrieveGet200ResponseVlanVlanPooling struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseVlanVlanpooling struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanRetrieveGet200ResponseVlan struct {
		AaaVlanOverride *bool                                      `json:"aaaVlanOverride,omitempty"`
		AccessVlan      *float64                                   `json:"accessVlan,omitempty"`
		CoreQinQEnabled *bool                                      `json:"coreQinQEnabled,omitempty"`
		CoreSVlan       *float64                                   `json:"coreSVlan,omitempty"`
		VlanPooling     *WlanRetrieveGet200ResponseVlanVlanPooling `json:"vlanPooling,omitempty"`
		Vlanpooling     *WlanRetrieveGet200ResponseVlanVlanpooling `json:"vlanpooling,omitempty"`
	}

	WlanRetrieveGet200Response struct {
		AccessTunnelType           *string                                               `json:"accessTunnelType,omitempty"`
		AccountingServiceOrProfile *WlanRetrieveGet200ResponseAccountingServiceOrProfile `json:"accountingServiceOrProfile,omitempty"`
		AdvancedOptions            *WlanRetrieveGet200ResponseAdvancedOptions            `json:"advancedOptions,omitempty"`
		AuthServiceOrProfile       *WlanRetrieveGet200ResponseAuthServiceOrProfile       `json:"authServiceOrProfile,omitempty"`
		AwsExtNasIPEnable          *bool                                                 `json:"awsExtNasIPEnable,omitempty"`
		AwsVenueEnable             *bool                                                 `json:"awsVenueEnable,omitempty"`
		BypassCNA                  *bool                                                 `json:"bypassCNA,omitempty"`
		CaleaEnabled               *bool                                                 `json:"caleaEnabled,omitempty"`
		CoreTunnelProfile          *WlanRetrieveGet200ResponseCoreTunnelProfile          `json:"coreTunnelProfile,omitempty"`
		DefaultUserTrafficProfile  *WlanRetrieveGet200ResponseDefaultUserTrafficProfile  `json:"defaultUserTrafficProfile,omitempty"`
		Description                *string                                               `json:"description,omitempty"`
		DevicePolicy               *WlanRetrieveGet200ResponseDevicePolicy               `json:"devicePolicy,omitempty"`
		DiffServProfile            *WlanRetrieveGet200ResponseDiffServProfile            `json:"diffServProfile,omitempty"`
		DNSServerProfile           *WlanRetrieveGet200ResponseDNSServerProfile           `json:"dnsServerProfile,omitempty"`
		DpTunnelNatEnabled         *bool                                                 `json:"dpTunnelNatEnabled,omitempty"`
		Dpsk                       *WlanRetrieveGet200ResponseDpsk                       `json:"dpsk,omitempty"`
		Encryption                 *WlanRetrieveGet200ResponseEncryption                 `json:"encryption,omitempty"`
		Hessid                     *string                                               `json:"hessid,omitempty"`
		Hotspot20Profile           *WlanRetrieveGet200ResponseHotspot20Profile           `json:"hotspot20Profile,omitempty"`
		ID                         *string                                               `json:"id,omitempty"`
		L2ACL                      *WlanRetrieveGet200ResponseL2ACL                      `json:"l2ACL,omitempty"`
		MacAuth                    *WlanRetrieveGet200ResponseMacAuth                    `json:"macAuth,omitempty"`
		Name                       *string                                               `json:"name,omitempty"`
		OperatorRealm              *string                                               `json:"operatorRealm,omitempty"`
		PortalServiceProfile       *WlanRetrieveGet200ResponsePortalServiceProfile       `json:"portalServiceProfile,omitempty"`
		QosMaps                    WlanRetrieveGet200ResponseQosMapsSlice                `json:"qosMaps,omitempty"`
		RadiusOptions              *WlanRetrieveGet200ResponseRadiusOptions              `json:"radiusOptions,omitempty"`
		Schedule                   *WlanRetrieveGet200ResponseSchedule                   `json:"schedule,omitempty"`
		Ssid                       *string                                               `json:"ssid,omitempty"`
		Type                       *string                                               `json:"type,omitempty"`
		Vlan                       *WlanRetrieveGet200ResponseVlan                       `json:"vlan,omitempty"`
		ZoneID                     *string                                               `json:"zoneId,omitempty"`
	}
)

// WlanRetrieveGet: Use this API command to retrieve a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WlanRetrieveGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanRetrieveGet(ctx context.Context, zoneId string, id string) (*http.Response, *WlanRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlans/{id}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &WlanRetrieveGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanModifyBasicPatchRequest struct {
		AccessTunnelType   *string `json:"accessTunnelType,omitempty"`
		AwsExtNasIPEnable  *bool   `json:"awsExtNasIPEnable,omitempty"`
		AwsVenueEnable     *bool   `json:"awsVenueEnable,omitempty"`
		BypassCNA          *bool   `json:"bypassCNA,omitempty"`
		CaleaEnabled       *bool   `json:"caleaEnabled,omitempty"`
		Description        *string `json:"description,omitempty"`
		DpTunnelNatEnabled *bool   `json:"dpTunnelNatEnabled,omitempty"`
		Hessid             *string `json:"hessid,omitempty"`
		Name               *string `json:"name,omitempty"`
		OperatorRealm      *string `json:"operatorRealm,omitempty"`
		Ssid               *string `json:"ssid,omitempty"`
	}
)

// WlanModifyBasicPatch: Use this API command to modify the basic information of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyBasicPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanDisableAccountingDelete: Use this API command to disable the accounting of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDisableAccountingDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}/accountingServiceOrProfile", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyAccountingPatchRequest struct {
		AccountingDelayEnabled *bool    `json:"accountingDelayEnabled,omitempty"`
		ID                     *string  `json:"id,omitempty"`
		InterimUpdateMin       *float64 `json:"interimUpdateMin,omitempty"`
		Name                   *string  `json:"name,omitempty"`
		ThroughController      *bool    `json:"throughController,omitempty"`
	}
)

// WlanModifyAccountingPatch: Use this API command to modify the accounting settings of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyAccountingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyAccountingPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyAccountingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/accountingServiceOrProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyAdvancedOptionsPatchRequestClientIsolationWhitelist struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanModifyAdvancedOptionsPatchRequest struct {
		AvcEnabled                  *bool                                                          `json:"avcEnabled,omitempty"`
		BandBalancing               *string                                                        `json:"bandBalancing,omitempty"`
		BssMinRateMbps              *string                                                        `json:"bssMinRateMbps,omitempty"`
		ClientFingerprintingEnabled *bool                                                          `json:"clientFingerprintingEnabled,omitempty"`
		ClientIdleTimeoutSec        *float64                                                       `json:"clientIdleTimeoutSec,omitempty"`
		ClientIsolationEnabled      *bool                                                          `json:"clientIsolationEnabled,omitempty"`
		ClientIsolationWhitelist    *WlanModifyAdvancedOptionsPatchRequestClientIsolationWhitelist `json:"clientIsolationWhitelist,omitempty"`
		ClientLoadBalancingEnabled  *bool                                                          `json:"clientLoadBalancingEnabled,omitempty"`
		DgafEnabled                 *bool                                                          `json:"dgafEnabled,omitempty"`
		Dhcp82Format                *string                                                        `json:"dhcp82Format,omitempty"`
		DhcpOption82Enabled         *bool                                                          `json:"dhcpOption82Enabled,omitempty"`
		DownlinkEnabled             *bool                                                          `json:"downlinkEnabled,omitempty"`
		DownlinkRate                *float64                                                       `json:"downlinkRate,omitempty"`
		ForceClientDHCPTimeoutSec   *float64                                                       `json:"forceClientDHCPTimeoutSec,omitempty"`
		HideSsidEnabled             *bool                                                          `json:"hideSsidEnabled,omitempty"`
		MaxClientsPerRadio          *float64                                                       `json:"maxClientsPerRadio,omitempty"`
		MgmtTxRateMbps              *string                                                        `json:"mgmtTxRateMbps,omitempty"`
		OfdmOnlyEnabled             *bool                                                          `json:"ofdmOnlyEnabled,omitempty"`
		OkcEnabled                  *bool                                                          `json:"okcEnabled,omitempty"`
		PmkCachingEnabled           *bool                                                          `json:"pmkCachingEnabled,omitempty"`
		Priority                    *string                                                        `json:"priority,omitempty"`
		ProxyARPEnabled             *bool                                                          `json:"proxyARPEnabled,omitempty"`
		Support8021DEnabled         *bool                                                          `json:"support80211dEnabled,omitempty"`
		Support8021KEnabled         *bool                                                          `json:"support80211kEnabled,omitempty"`
		UnauthClientStatsEnabled    *bool                                                          `json:"unauthClientStatsEnabled,omitempty"`
		UplinkEnabled               *bool                                                          `json:"uplinkEnabled,omitempty"`
		UplinkRate                  *float64                                                       `json:"uplinkRate,omitempty"`
	}
)

// WlanModifyAdvancedOptionsPatch: Use this API command to modify the advanced settings of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyAdvancedOptionsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyAdvancedOptionsPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyAdvancedOptionsPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/advancedOptions", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyAuthenticationPatchRequest struct {
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		ThroughController       *bool   `json:"throughController,omitempty"`
	}
)

// WlanModifyAuthenticationPatch: Use this API command to modify the authentication method of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyAuthenticationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyAuthenticationPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyAuthenticationPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/authServiceOrProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	CaleaSetCaleaConfigForWlanPatchRequest struct {
		AwsExtNasIPEnable  *bool `json:"awsExtNasIPEnable,omitempty"`
		AwsVenueEnable     *bool `json:"awsVenueEnable,omitempty"`
		CaleaEnabled       *bool `json:"caleaEnabled,omitempty"`
		DpTunnelNatEnabled *bool `json:"dpTunnelNatEnabled,omitempty"`
	}
)

// CaleaSetCaleaConfigForWlanPatch: Use this API command to set CALEA config for wlan.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *CaleaSetCaleaConfigForWlanPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) CaleaSetCaleaConfigForWlanPatch(ctx context.Context, zoneId string, id string, requestBody *CaleaSetCaleaConfigForWlanPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/calea", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyCoreTunnelPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Type *string `json:"type,omitempty"`
	}
)

// WlanModifyCoreTunnelPatch: Use this API command to modify the core tunnel configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyCoreTunnelPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyCoreTunnelPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyCoreTunnelPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/coreTunnelProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyUserTrafficProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyUserTrafficProfilePatch: Use this API command to modify the user traffic profile configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyUserTrafficProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyUserTrafficProfilePatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyUserTrafficProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/defaultUserTrafficProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanDisableDevicePolicyDelete: Use this API command to disable the device policy of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDisableDevicePolicyDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}/devicePolicy", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyDevicePolicyPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyDevicePolicyPatch: Use this API command to modify the device policy of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyDevicePolicyPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyDevicePolicyPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyDevicePolicyPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/devicePolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanDisableDiffservProfileDelete: Use this API command to disable the DiffServ profile of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDisableDiffservProfileDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}/diffServProfile", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyDiffservProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyDiffservProfilePatch: Use this API command to modify the DiffServ profile of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyDiffservProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyDiffservProfilePatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyDiffservProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/diffServProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanDisableDnsServerProfileDelete: Use this API command to disable DNS server profile of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDisableDnsServerProfileDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}/dnsServerProfile", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyDnsServerProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyDnsServerProfilePatch: Use this API command to modify DNS server profile of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyDnsServerProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyDnsServerProfilePatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyDnsServerProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/dnsServerProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DynamicPskRetrieveDpskInfoByWlanGet200ResponseListSlice []*DynamicPskRetrieveDpskInfoByWlanGet200ResponseList

	DynamicPskRetrieveDpskInfoByWlanGet200ResponseList struct {
		CreationDateTime   *float64 `json:"creationDateTime,omitempty"`
		ExpirationDateTime *string  `json:"expirationDateTime,omitempty"`
		GroupDpsk          *bool    `json:"groupDpsk,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		MacAddress         *string  `json:"macAddress,omitempty"`
		Passphrase         *string  `json:"passphrase,omitempty"`
		UserName           *string  `json:"userName,omitempty"`
		UserRoleID         *string  `json:"userRoleId,omitempty"`
		VlanID             *float64 `json:"vlanId,omitempty"`
		WlanID             *string  `json:"wlanId,omitempty"`
	}

	DynamicPskRetrieveDpskInfoByWlanGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       DynamicPskRetrieveDpskInfoByWlanGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// DynamicPskRetrieveDpskInfoByWlanGet: Use this API command to retrieve DPSK info of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskRetrieveDpskInfoByWlanGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskRetrieveDpskInfoByWlanGet(ctx context.Context, zoneId string, id string) (*http.Response, *DynamicPskRetrieveDpskInfoByWlanGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlans/{id}/dpsk", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &DynamicPskRetrieveDpskInfoByWlanGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WlanModifyDpskSettingPatchRequest struct {
		DpskEnabled  *bool    `json:"dpskEnabled,omitempty"`
		DpskFromType *string  `json:"dpskFromType,omitempty"`
		DpskType     *string  `json:"dpskType,omitempty"`
		Expiration   *string  `json:"expiration,omitempty"`
		Length       *float64 `json:"length,omitempty"`
	}
)

// WlanModifyDpskSettingPatch: Use this API command to modify DPSK setting of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyDpskSettingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyDpskSettingPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyDpskSettingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/dpsk", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	DynamicPskDeleteDpskPostRequestIDListSlice []*string

	DynamicPskDeleteDpskPostRequest struct {
		IDList DynamicPskDeleteDpskPostRequestIDListSlice `json:"idList,omitempty"`
	}

	DynamicPskDeleteDpskPost200Response struct {
		ResultCount *float64 `json:"resultCount,omitempty"`
	}
)

// DynamicPskDeleteDpskPost: Use this API command to delete DPSKs of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *DynamicPskDeleteDpskPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskDeleteDpskPost200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskDeleteDpskPost(ctx context.Context, zoneId string, id string, requestBody *DynamicPskDeleteDpskPostRequest) (*http.Response, *DynamicPskDeleteDpskPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/{id}/dpsk", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &DynamicPskDeleteDpskPost200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DynamicPskBatchGenerateDpsksPostRequestPassphraseListSlice []*string

	DynamicPskBatchGenerateDpsksPostRequest struct {
		Amount         *float64                                                   `json:"amount,omitempty"`
		GroupDpsk      *bool                                                      `json:"groupDpsk,omitempty"`
		PassphraseList DynamicPskBatchGenerateDpsksPostRequestPassphraseListSlice `json:"passphraseList,omitempty"`
		UserName       *string                                                    `json:"userName,omitempty"`
		UserRoleID     *string                                                    `json:"userRoleId,omitempty"`
		VlanID         *float64                                                   `json:"vlanId,omitempty"`
	}

	DynamicPskBatchGenerateDpsksPost201ResponseDpskInfoListSlice []*DynamicPskBatchGenerateDpsksPost201ResponseDpskInfoList

	DynamicPskBatchGenerateDpsksPost201ResponseDpskInfoList struct {
		CreationDateTime   interface{} `json:"creationDateTime,omitempty"`
		ExpirationDateTime *string     `json:"expirationDateTime,omitempty"`
		GroupDpsk          *bool       `json:"groupDpsk,omitempty"`
		ID                 *string     `json:"id,omitempty"`
		MacAddress         *string     `json:"macAddress,omitempty"`
		Passphrase         *string     `json:"passphrase,omitempty"`
		UserName           *string     `json:"userName,omitempty"`
		UserRoleID         *string     `json:"userRoleId,omitempty"`
		VlanID             *float64    `json:"vlanId,omitempty"`
		WlanID             *string     `json:"wlanId,omitempty"`
	}

	DynamicPskBatchGenerateDpsksPost201Response struct {
		DpskInfoList DynamicPskBatchGenerateDpsksPost201ResponseDpskInfoListSlice `json:"dpskInfoList,omitempty"`
		ResultCount  *float64                                                     `json:"resultCount,omitempty"`
	}
)

// DynamicPskBatchGenerateDpsksPost: Use this API command to batch generate DPSKs of a WLAN. You can either specify passphrases or not. If the amount is bigger than 1, system will generate usernames with index. e.g. student-1, student-2, …etc.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *DynamicPskBatchGenerateDpsksPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskBatchGenerateDpsksPost201Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskBatchGenerateDpsksPost(ctx context.Context, zoneId string, id string, requestBody *DynamicPskBatchGenerateDpsksPostRequest) (*http.Response, *DynamicPskBatchGenerateDpsksPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/{id}/dpsk/batchGenUnbound", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	out := &DynamicPskBatchGenerateDpsksPost201Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	DynamicPskRetrieveDpskInfoByIdGet200ResponseListSlice []*DynamicPskRetrieveDpskInfoByIdGet200ResponseList

	DynamicPskRetrieveDpskInfoByIdGet200ResponseList struct {
		CreationDateTime   *float64 `json:"creationDateTime,omitempty"`
		ExpirationDateTime *string  `json:"expirationDateTime,omitempty"`
		GroupDpsk          *bool    `json:"groupDpsk,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		MacAddress         *string  `json:"macAddress,omitempty"`
		Passphrase         *string  `json:"passphrase,omitempty"`
		UserName           *string  `json:"userName,omitempty"`
		UserRoleID         *string  `json:"userRoleId,omitempty"`
		VlanID             *float64 `json:"vlanId,omitempty"`
		WlanID             *string  `json:"wlanId,omitempty"`
	}

	DynamicPskRetrieveDpskInfoByIdGet200Response struct {
		FirstIndex *int                                                  `json:"firstIndex,omitempty"`
		HasMore    *bool                                                 `json:"hasMore,omitempty"`
		List       DynamicPskRetrieveDpskInfoByIdGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                  `json:"totalCount,omitempty"`
	}
)

// DynamicPskRetrieveDpskInfoByIdGet: Use this API command to retrieve DPSK info.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - dpskId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DynamicPskRetrieveDpskInfoByIdGet200Response
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskRetrieveDpskInfoByIdGet(ctx context.Context, zoneId string, id string, dpskId string) (*http.Response, *DynamicPskRetrieveDpskInfoByIdGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(dpskId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"dpskId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/rkszones/{zoneId}/wlans/{id}/dpsk/{dpskId}", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("dpskId", dpskId)
	out := &DynamicPskRetrieveDpskInfoByIdGet200Response{}
	httpResponse, _, err := r.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DynamicPskUpdateDpskInfoByIdPatchRequest struct {
		UserName *string `json:"userName,omitempty"`
	}
)

// DynamicPskUpdateDpskInfoByIdPatch: Use this API command to update DPSK info.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - dpskId (string)
// - requestBody: *DynamicPskUpdateDpskInfoByIdPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) DynamicPskUpdateDpskInfoByIdPatch(ctx context.Context, zoneId string, id string, dpskId string, requestBody *DynamicPskUpdateDpskInfoByIdPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(dpskId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"dpskId\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/dpsk/{dpskId}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	request.SetPathParameter("dpskId", dpskId)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyEncryptionPatchRequest struct {
		Algorithm           *string  `json:"algorithm,omitempty"`
		KeyInHex            *string  `json:"keyInHex,omitempty"`
		KeyIndex            *float64 `json:"keyIndex,omitempty"`
		Method              *string  `json:"method,omitempty"`
		Mfp                 *string  `json:"mfp,omitempty"`
		MobilityDomainID    *float64 `json:"mobilityDomainId,omitempty"`
		Passphrase          *string  `json:"passphrase,omitempty"`
		Support8021REnabled *bool    `json:"support80211rEnabled,omitempty"`
	}
)

// WlanModifyEncryptionPatch: Use this API command to modify the encryption settings of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyEncryptionPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyEncryptionPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyEncryptionPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/encryption", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyHotspot20ProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyHotspot20ProfilePatch: Use this API command to modify the Hotspot 2.0 profile configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyHotspot20ProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyHotspot20ProfilePatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyHotspot20ProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/hotspot20Profile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanDisableLayer2AclDelete: Use this API command to disable the layer 2 access control list (ACL) configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDisableLayer2AclDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}/l2ACL", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyLayer2AclPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyLayer2AclPatch: Use this API command to modify the layer 2 access control list (ACL) configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyLayer2AclPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyLayer2AclPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyLayer2AclPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/l2ACL", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyMacAuthPatchRequest struct {
		CustomizedPassword *string `json:"customizedPassword,omitempty"`
		MacAuthMacFormat   *string `json:"macAuthMacFormat,omitempty"`
	}
)

// WlanModifyMacAuthPatch: Use this API command to modify the MAC authentication settings of a WLAN. macAuthMacFormat : Open(aabbccddeeff), 802.1X(AA-BB-CC-DD-EE-FF), UpperColon (AA:BB:CC:DD:EE:FF), Upper (AABBCCDDEEFF), LowerDash (aa-bb-cc-dd-ee-ff) and  LowerColon (aa:bb:cc:dd:ee:ff)
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyMacAuthPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyMacAuthPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyMacAuthPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/macAuth", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyPortalProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// WlanModifyPortalProfilePatch: Use this API command to modify the portal configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyPortalProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyPortalProfilePatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyPortalProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/portalServiceProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanDisableQosMapSetDelete: Use this API command to disable Qos Map Set of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanDisableQosMapSetDelete(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/rkszones/{zoneId}/wlans/{id}/qosMaps", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyQosMapSetPatchRequestSlice []*WlanModifyQosMapSetPatchRequest

	WlanModifyQosMapSetPatchRequestExceptsSlice []*float64

	WlanModifyQosMapSetPatchRequest struct {
		Enable   *bool                                       `json:"enable,omitempty"`
		Excepts  WlanModifyQosMapSetPatchRequestExceptsSlice `json:"excepts,omitempty"`
		High     *float64                                    `json:"high,omitempty"`
		Low      *float64                                    `json:"low,omitempty"`
		Priority *float64                                    `json:"priority,omitempty"`
	}
)

// WlanModifyQosMapSetPatch: Use this API command to modify Qos Map Set of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyQosMapSetPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyQosMapSetPatch(ctx context.Context, zoneId string, id string, requestBody WlanModifyQosMapSetPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/qosMaps", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

// WlanEnableQosMapSetPost: Use this API command to enable Qos Map Set of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanEnableQosMapSetPost(ctx context.Context, zoneId string, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/rkszones/{zoneId}/wlans/{id}/qosMaps", true)
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 201, nil)
}

type (
	WlanModifyRadiusOptionsPatchRequest struct {
		CalledStaIDType        *string  `json:"calledStaIdType,omitempty"`
		CustomizedNasID        *string  `json:"customizedNasId,omitempty"`
		NasIDType              *string  `json:"nasIdType,omitempty"`
		NasIPType              *string  `json:"nasIpType,omitempty"`
		NasIPUserDefined       *string  `json:"nasIpUserDefined,omitempty"`
		NasMaxRetry            *float64 `json:"nasMaxRetry,omitempty"`
		NasReconnectPrimaryMin *float64 `json:"nasReconnectPrimaryMin,omitempty"`
		NasRequestTimeoutSec   *float64 `json:"nasRequestTimeoutSec,omitempty"`
	}
)

// WlanModifyRadiusOptionsPatch: Use this API command to modify the RADIUS settings of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyRadiusOptionsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyRadiusOptionsPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyRadiusOptionsPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/radiusOptions", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifySchedulePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Type *string `json:"type,omitempty"`
	}
)

// WlanModifySchedulePatch: Use this API command to modify the schedule configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifySchedulePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifySchedulePatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifySchedulePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/schedule", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}

type (
	WlanModifyVlanPatchRequestVlanPooling struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	WlanModifyVlanPatchRequest struct {
		AaaVlanOverride *bool                                  `json:"aaaVlanOverride,omitempty"`
		AccessVlan      *float64                               `json:"accessVlan,omitempty"`
		CoreQinQEnabled *bool                                  `json:"coreQinQEnabled,omitempty"`
		CoreSVlan       *float64                               `json:"coreSVlan,omitempty"`
		VlanPooling     *WlanModifyVlanPatchRequestVlanPooling `json:"vlanPooling,omitempty"`
		Vlanpooling     interface{}                            `json:"vlanpooling,omitempty"`
	}
)

// WlanModifyVlanPatch: Use this API command to modify the VLAN configuration of a WLAN.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - zoneId (UUIDv4): Zone ID
// - id (integer): WLAN ID
// - requestBody: *WlanModifyVlanPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (r *RuckusZonesAPI) WlanModifyVlanPatch(ctx context.Context, zoneId string, id string, requestBody *WlanModifyVlanPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(zoneId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
	}
	err = validators.StrIsPositiveInt(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/rkszones/{zoneId}/wlans/{id}/vlan", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("zoneId", zoneId)
	request.SetPathParameter("id", id)
	return r.client.Ensure(ctx, request, 204, nil)
}
