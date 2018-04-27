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

type MapsAPI struct {
	client *Client
}
type (
	IndoormapGetIndoorMapListGet200ResponseListSlice []*IndoormapGetIndoorMapListGet200ResponseList

	IndoormapGetIndoorMapListGet200ResponseListScaleA struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapGetIndoorMapListGet200ResponseListScaleB struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapGetIndoorMapListGet200ResponseListScale struct {
		A        *IndoormapGetIndoorMapListGet200ResponseListScaleA `json:"a,omitempty"`
		B        *IndoormapGetIndoorMapListGet200ResponseListScaleB `json:"b,omitempty"`
		Distance *float64                                           `json:"distance,omitempty"`
		Unit     *string                                            `json:"unit,omitempty"`
	}

	IndoormapGetIndoorMapListGet200ResponseList struct {
		Address       *string                                           `json:"address,omitempty"`
		ApGroupID     *string                                           `json:"apGroupId,omitempty"`
		Description   *string                                           `json:"description,omitempty"`
		DomainID      *string                                           `json:"domainId,omitempty"`
		GroupType     *string                                           `json:"groupType,omitempty"`
		ID            *string                                           `json:"id,omitempty"`
		ImageData     *string                                           `json:"imageData,omitempty"`
		ImageFileName *string                                           `json:"imageFileName,omitempty"`
		Key           *string                                           `json:"key,omitempty"`
		Latitude      *float64                                          `json:"latitude,omitempty"`
		Longitude     *float64                                          `json:"longitude,omitempty"`
		Name          *string                                           `json:"name,omitempty"`
		Orientation   *string                                           `json:"orientation,omitempty"`
		Scale         *IndoormapGetIndoorMapListGet200ResponseListScale `json:"scale,omitempty"`
		TenantID      *string                                           `json:"tenantId,omitempty"`
		ZoneID        *string                                           `json:"zoneId,omitempty"`
	}

	IndoormapGetIndoorMapListGet200Response struct {
		Extra      *json.RawMessage                                 `json:"extra,omitempty"`
		FirstIndex *int                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                            `json:"hasMore,omitempty"`
		List       IndoormapGetIndoorMapListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                             `json:"totalCount,omitempty"`
	}
)

// IndoormapGetIndoorMapListGet: Use this API command to get indoor map list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - groupId (string)
// - groupType (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapGetIndoorMapListGet200Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapGetIndoorMapListGet(ctx context.Context, groupId string, groupType string) (*http.Response, *IndoormapGetIndoorMapListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(groupId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"groupId\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(groupType)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"groupType\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/maps", true)
	request.SetQueryParameter("groupId", groupId)
	request.SetQueryParameter("groupType", groupType)
	out := &IndoormapGetIndoorMapListGet200Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	IndoormapCreateIndoormapPostRequestScaleA struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapCreateIndoormapPostRequestScaleB struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapCreateIndoormapPostRequestScale struct {
		A        *IndoormapCreateIndoormapPostRequestScaleA `json:"a,omitempty"`
		B        *IndoormapCreateIndoormapPostRequestScaleB `json:"b,omitempty"`
		Distance *float64                                   `json:"distance,omitempty"`
		Unit     *string                                    `json:"unit,omitempty"`
	}

	IndoormapCreateIndoormapPostRequest struct {
		Address       *string                                   `json:"address,omitempty"`
		ApGroupID     *string                                   `json:"apGroupId,omitempty"`
		Description   *string                                   `json:"description,omitempty"`
		DomainID      *string                                   `json:"domainId,omitempty"`
		GroupType     *string                                   `json:"groupType,omitempty"`
		ID            *string                                   `json:"id,omitempty"`
		ImageData     *string                                   `json:"imageData,omitempty"`
		ImageFileName *string                                   `json:"imageFileName,omitempty"`
		InputItem     *string                                   `json:"inputItem,omitempty"`
		Key           *string                                   `json:"key,omitempty"`
		Latitude      *float64                                  `json:"latitude,omitempty"`
		Longitude     *float64                                  `json:"longitude,omitempty"`
		Name          *string                                   `json:"name,omitempty"`
		Orientation   *string                                   `json:"orientation,omitempty"`
		Scale         *IndoormapCreateIndoormapPostRequestScale `json:"scale,omitempty"`
		TenantID      *string                                   `json:"tenantId,omitempty"`
		ZoneID        *string                                   `json:"zoneId,omitempty"`
	}

	IndoormapCreateIndoormapPost201Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// IndoormapCreateIndoormapPost: Use this API command to create indoorMap
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *IndoormapCreateIndoormapPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapCreateIndoormapPost201Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapCreateIndoormapPost(ctx context.Context, requestBody *IndoormapCreateIndoormapPostRequest) (*http.Response, *IndoormapCreateIndoormapPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/maps", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &IndoormapCreateIndoormapPost201Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

type (
	IndoormapQueryIndoormapPostRequestAttributesSlice []string

	IndoormapQueryIndoormapPostRequestFiltersSlice []*IndoormapQueryIndoormapPostRequestFilters

	IndoormapQueryIndoormapPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	IndoormapQueryIndoormapPostRequestFullTextSearchFieldsSlice []string

	IndoormapQueryIndoormapPostRequestFullTextSearch struct {
		Fields IndoormapQueryIndoormapPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                     `json:"type,omitempty"`
		Value  *string                                                     `json:"value,omitempty"`
	}

	IndoormapQueryIndoormapPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	IndoormapQueryIndoormapPostRequest struct {
		Attributes     IndoormapQueryIndoormapPostRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        IndoormapQueryIndoormapPostRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *IndoormapQueryIndoormapPostRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                              `json:"limit,omitempty"`
		Options        interface{}                                       `json:"options,omitempty"`
		Page           *int                                              `json:"page,omitempty"`
		SortInfo       *IndoormapQueryIndoormapPostRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                              `json:"start,omitempty"`
	}

	IndoormapQueryIndoormapPost200ResponseListSlice []*IndoormapQueryIndoormapPost200ResponseList

	IndoormapQueryIndoormapPost200ResponseListScaleA struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapQueryIndoormapPost200ResponseListScaleB struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapQueryIndoormapPost200ResponseListScale struct {
		A        *IndoormapQueryIndoormapPost200ResponseListScaleA `json:"a,omitempty"`
		B        *IndoormapQueryIndoormapPost200ResponseListScaleB `json:"b,omitempty"`
		Distance *float64                                          `json:"distance,omitempty"`
		Unit     *string                                           `json:"unit,omitempty"`
	}

	IndoormapQueryIndoormapPost200ResponseList struct {
		Address       *string                                          `json:"address,omitempty"`
		ApGroupID     *string                                          `json:"apGroupId,omitempty"`
		Description   *string                                          `json:"description,omitempty"`
		DomainID      *string                                          `json:"domainId,omitempty"`
		GroupType     *string                                          `json:"groupType,omitempty"`
		ID            *string                                          `json:"id,omitempty"`
		ImageData     *string                                          `json:"imageData,omitempty"`
		ImageFileName *string                                          `json:"imageFileName,omitempty"`
		Key           *string                                          `json:"key,omitempty"`
		Latitude      *float64                                         `json:"latitude,omitempty"`
		Longitude     *float64                                         `json:"longitude,omitempty"`
		Name          *string                                          `json:"name,omitempty"`
		Orientation   *string                                          `json:"orientation,omitempty"`
		Scale         *IndoormapQueryIndoormapPost200ResponseListScale `json:"scale,omitempty"`
		TenantID      *string                                          `json:"tenantId,omitempty"`
		ZoneID        *string                                          `json:"zoneId,omitempty"`
	}

	IndoormapQueryIndoormapPost200Response struct {
		Extra      *json.RawMessage                                `json:"extra,omitempty"`
		FirstIndex *int                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                           `json:"hasMore,omitempty"`
		List       IndoormapQueryIndoormapPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                            `json:"totalCount,omitempty"`
	}
)

// IndoormapQueryIndoormapPost: Use this API command to query indoorMap
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *IndoormapQueryIndoormapPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapQueryIndoormapPost200Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapQueryIndoormapPost(ctx context.Context, requestBody *IndoormapQueryIndoormapPostRequest) (*http.Response, *IndoormapQueryIndoormapPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/maps/query", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &IndoormapQueryIndoormapPost200Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	IndoormapDeleteIndoormapDelete204Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// IndoormapDeleteIndoormapDelete: Use this API command to delete indoor map.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - indoorMapId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapDeleteIndoormapDelete204Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapDeleteIndoormapDelete(ctx context.Context, indoorMapId string) (*http.Response, *IndoormapDeleteIndoormapDelete204Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(indoorMapId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"indoorMapId\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/maps/{indoorMapId}", true)
	request.SetPathParameter("indoorMapId", indoorMapId)
	out := &IndoormapDeleteIndoormapDelete204Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 204, out)
	return httpResponse, out, err
}

type (
	IndoormapGetIndoormapGet200ResponseScaleA struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapGetIndoormapGet200ResponseScaleB struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapGetIndoormapGet200ResponseScale struct {
		A        *IndoormapGetIndoormapGet200ResponseScaleA `json:"a,omitempty"`
		B        *IndoormapGetIndoormapGet200ResponseScaleB `json:"b,omitempty"`
		Distance *float64                                   `json:"distance,omitempty"`
		Unit     *string                                    `json:"unit,omitempty"`
	}

	IndoormapGetIndoormapGet200Response struct {
		Address       *string                                   `json:"address,omitempty"`
		ApGroupID     *string                                   `json:"apGroupId,omitempty"`
		Description   *string                                   `json:"description,omitempty"`
		DomainID      *string                                   `json:"domainId,omitempty"`
		GroupType     *string                                   `json:"groupType,omitempty"`
		ID            *string                                   `json:"id,omitempty"`
		ImageData     *string                                   `json:"imageData,omitempty"`
		ImageFileName *string                                   `json:"imageFileName,omitempty"`
		InputItem     *string                                   `json:"inputItem,omitempty"`
		Key           *string                                   `json:"key,omitempty"`
		Latitude      *float64                                  `json:"latitude,omitempty"`
		Longitude     *float64                                  `json:"longitude,omitempty"`
		Name          *string                                   `json:"name,omitempty"`
		Orientation   *string                                   `json:"orientation,omitempty"`
		Scale         *IndoormapGetIndoormapGet200ResponseScale `json:"scale,omitempty"`
		TenantID      *string                                   `json:"tenantId,omitempty"`
		ZoneID        *string                                   `json:"zoneId,omitempty"`
	}
)

// IndoormapGetIndoormapGet: Use this API command to get indoor maps.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - indoorMapId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapGetIndoormapGet200Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapGetIndoormapGet(ctx context.Context, indoorMapId string) (*http.Response, *IndoormapGetIndoormapGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(indoorMapId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"indoorMapId\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/maps/{indoorMapId}", true)
	request.SetPathParameter("indoorMapId", indoorMapId)
	out := &IndoormapGetIndoormapGet200Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	IndoormapUpdateIndoorMapPatchRequest struct {
		Address       *string  `json:"address,omitempty"`
		ApGroupID     *string  `json:"apGroupId,omitempty"`
		Description   *string  `json:"description,omitempty"`
		DomainID      *string  `json:"domainId,omitempty"`
		GroupType     *string  `json:"groupType,omitempty"`
		ID            *string  `json:"id,omitempty"`
		ImageData     *string  `json:"imageData,omitempty"`
		ImageFileName *string  `json:"imageFileName,omitempty"`
		Latitude      *float64 `json:"latitude,omitempty"`
		Longitude     *float64 `json:"longitude,omitempty"`
		Name          *string  `json:"name,omitempty"`
		Orientation   *string  `json:"orientation,omitempty"`
		TenantID      *string  `json:"tenantId,omitempty"`
		ZoneID        *string  `json:"zoneId,omitempty"`
	}

	IndoormapUpdateIndoorMapPatch204Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// IndoormapUpdateIndoorMapPatch: Use this API command to update specific indoor map.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - indoorMapId (string)
// - requestBody: *IndoormapUpdateIndoorMapPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapUpdateIndoorMapPatch204Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapUpdateIndoorMapPatch(ctx context.Context, indoorMapId string, requestBody *IndoormapUpdateIndoorMapPatchRequest) (*http.Response, *IndoormapUpdateIndoorMapPatch204Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(indoorMapId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"indoorMapId\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/maps/{indoorMapId}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("indoorMapId", indoorMapId)
	out := &IndoormapUpdateIndoorMapPatch204Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 204, out)
	return httpResponse, out, err
}

type (
	IndoormapPutApsInIndoorMapPutRequestSlice []*IndoormapPutApsInIndoorMapPutRequest

	IndoormapPutApsInIndoorMapPutRequestIndoorMapXy struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapPutApsInIndoorMapPutRequest struct {
		IndoorMapXy *IndoormapPutApsInIndoorMapPutRequestIndoorMapXy `json:"indoorMapXy,omitempty"`
		Mac         *string                                          `json:"mac,omitempty"`
	}

	IndoormapPutApsInIndoorMapPut204Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// IndoormapPutApsInIndoorMapPut: Use this API command to put Aps in indoor map.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - indoorMapId (string)
// - requestBody: *IndoormapPutApsInIndoorMapPutRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *IndoormapPutApsInIndoorMapPut204Response
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapPutApsInIndoorMapPut(ctx context.Context, indoorMapId string, requestBody IndoormapPutApsInIndoorMapPutRequestSlice) (*http.Response, *IndoormapPutApsInIndoorMapPut204Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(indoorMapId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"indoorMapId\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/maps/{indoorMapId}/aps", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("indoorMapId", indoorMapId)
	out := &IndoormapPutApsInIndoorMapPut204Response{}
	httpResponse, _, err := m.client.Ensure(ctx, request, 204, out)
	return httpResponse, out, err
}

type (
	IndoormapUpdateIndoorMapPatch1RequestA struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapUpdateIndoorMapPatch1RequestB struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	IndoormapUpdateIndoorMapPatch1Request struct {
		A        *IndoormapUpdateIndoorMapPatch1RequestA `json:"a,omitempty"`
		B        *IndoormapUpdateIndoorMapPatch1RequestB `json:"b,omitempty"`
		Distance *float64                                `json:"distance,omitempty"`
		Unit     *string                                 `json:"unit,omitempty"`
	}
)

// IndoormapUpdateIndoorMapPatch1: Use this API command to update the scale of specific indoor map.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - indoorMapId (string)
// - requestBody: *IndoormapUpdateIndoorMapPatch1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (m *MapsAPI) IndoormapUpdateIndoorMapPatch1(ctx context.Context, indoorMapId string, requestBody *IndoormapUpdateIndoorMapPatch1Request) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(indoorMapId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"indoorMapId\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/maps/{indoorMapId}/scale", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("indoorMapId", indoorMapId)
	return m.client.Ensure(ctx, request, 204, nil)
}
