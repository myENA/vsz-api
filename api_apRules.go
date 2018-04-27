package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-04-27T15:10:38-0500
// API Version: v5

type APRulesAPI struct {
	client *Client
}
type (
	ApRegistrationRulesRetrieveListGet200ResponseListSlice []*ApRegistrationRulesRetrieveListGet200ResponseList

	ApRegistrationRulesRetrieveListGet200ResponseList struct {
		Description *string  `json:"description,omitempty"`
		ID          *string  `json:"id,omitempty"`
		Priority    *float64 `json:"priority,omitempty"`
	}

	ApRegistrationRulesRetrieveListGet200Response struct {
		FirstIndex *int                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                  `json:"hasMore,omitempty"`
		List       ApRegistrationRulesRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                   `json:"totalCount,omitempty"`
	}
)

// ApRegistrationRulesRetrieveListGet: Use this API command to retrieve a list of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApRegistrationRulesRetrieveListGet200Response
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesRetrieveListGet(ctx context.Context) (*http.Response, *ApRegistrationRulesRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/apRules", true)
	out := &ApRegistrationRulesRetrieveListGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApRegistrationRulesCreatePostRequestGpsCoordinates struct {
		Distance  *float64 `json:"distance,omitempty"`
		Latitude  *float64 `json:"latitude,omitempty"`
		Longitude *float64 `json:"longitude,omitempty"`
	}

	ApRegistrationRulesCreatePostRequestIPAddressRange struct {
		FromIP *string `json:"fromIp,omitempty"`
		ToIP   *string `json:"toIp,omitempty"`
	}

	ApRegistrationRulesCreatePostRequestMobilityZone struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApRegistrationRulesCreatePostRequestSubnet struct {
		NetworkAddress *string `json:"networkAddress,omitempty"`
		SubnetMask     *string `json:"subnetMask,omitempty"`
	}

	ApRegistrationRulesCreatePostRequest struct {
		Description    *string                                             `json:"description,omitempty"`
		GpsCoordinates *ApRegistrationRulesCreatePostRequestGpsCoordinates `json:"gpsCoordinates,omitempty"`
		IPAddressRange *ApRegistrationRulesCreatePostRequestIPAddressRange `json:"ipAddressRange,omitempty"`
		MobilityZone   *ApRegistrationRulesCreatePostRequestMobilityZone   `json:"mobilityZone,omitempty"`
		ProvisionTag   interface{}                                         `json:"provisionTag,omitempty"`
		Subnet         *ApRegistrationRulesCreatePostRequestSubnet         `json:"subnet,omitempty"`
		Type           *string                                             `json:"type,omitempty"`
	}

	ApRegistrationRulesCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ApRegistrationRulesCreatePost: Use this API command to create AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ApRegistrationRulesCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApRegistrationRulesCreatePost201Response
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesCreatePost(ctx context.Context, requestBody *ApRegistrationRulesCreatePostRequest) (*http.Response, *ApRegistrationRulesCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/apRules", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &ApRegistrationRulesCreatePost201Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ApRegistrationRulesMovePriorityDownGet: Use this API command to move Priority Down of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesMovePriorityDownGet(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/apRules/priorityDown/{id}", true)
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

// ApRegistrationRulesMovePriorityUpGet: Use this API command to move Priority Up of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesMovePriorityUpGet(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/apRules/priorityUp/{id}", true)
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

// ApRegistrationRulesDeleteDelete: Use this API command to delete AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/apRules/{id}", true)
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	ApRegistrationRulesRetrieveGet200ResponseGpsCoordinates struct {
		Distance  *float64 `json:"distance,omitempty"`
		Latitude  *float64 `json:"latitude,omitempty"`
		Longitude *float64 `json:"longitude,omitempty"`
	}

	ApRegistrationRulesRetrieveGet200ResponseIPAddressRange struct {
		FromIP *string `json:"fromIp,omitempty"`
		ToIP   *string `json:"toIp,omitempty"`
	}

	ApRegistrationRulesRetrieveGet200ResponseMobilityZone struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApRegistrationRulesRetrieveGet200ResponseSubnet struct {
		NetworkAddress *string `json:"networkAddress,omitempty"`
		SubnetMask     *string `json:"subnetMask,omitempty"`
	}

	ApRegistrationRulesRetrieveGet200Response struct {
		Description    *string                                                  `json:"description,omitempty"`
		GpsCoordinates *ApRegistrationRulesRetrieveGet200ResponseGpsCoordinates `json:"gpsCoordinates,omitempty"`
		ID             *string                                                  `json:"id,omitempty"`
		IPAddressRange *ApRegistrationRulesRetrieveGet200ResponseIPAddressRange `json:"ipAddressRange,omitempty"`
		MobilityZone   *ApRegistrationRulesRetrieveGet200ResponseMobilityZone   `json:"mobilityZone,omitempty"`
		Priority       *float64                                                 `json:"priority,omitempty"`
		ProvisionTag   interface{}                                              `json:"provisionTag,omitempty"`
		Subnet         *ApRegistrationRulesRetrieveGet200ResponseSubnet         `json:"subnet,omitempty"`
		Type           *string                                                  `json:"type,omitempty"`
	}
)

// ApRegistrationRulesRetrieveGet: Use this API command to retrieve AP Registration Rules profile by ID.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApRegistrationRulesRetrieveGet200Response
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesRetrieveGet(ctx context.Context, id string) (*http.Response, *ApRegistrationRulesRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/apRules/{id}", true)
	request.SetPathParameter("id", id)
	out := &ApRegistrationRulesRetrieveGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApRegistrationRulesModifyBasicPatchRequest struct {
		Description  *string `json:"description,omitempty"`
		ProvisionTag *string `json:"provisionTag,omitempty"`
		Type         *string `json:"type,omitempty"`
	}
)

// ApRegistrationRulesModifyBasicPatch: Use this API command to modify the basic information of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
// - requestBody: *ApRegistrationRulesModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesModifyBasicPatch(ctx context.Context, id string, requestBody *ApRegistrationRulesModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/apRules/{id}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	ApRegistrationRulesModifyGpscoordinatesPatchRequest struct {
		Distance  *float64 `json:"distance,omitempty"`
		Latitude  *float64 `json:"latitude,omitempty"`
		Longitude *float64 `json:"longitude,omitempty"`
	}
)

// ApRegistrationRulesModifyGpscoordinatesPatch: Use this API command to modify GPSCoordinates of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
// - requestBody: *ApRegistrationRulesModifyGpscoordinatesPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesModifyGpscoordinatesPatch(ctx context.Context, id string, requestBody *ApRegistrationRulesModifyGpscoordinatesPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/apRules/{id}/gpsCoordinates", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	ApRegistrationRulesModifyIpAddressRangePatchRequest struct {
		FromIP *string `json:"fromIp,omitempty"`
		ToIP   *string `json:"toIp,omitempty"`
	}
)

// ApRegistrationRulesModifyIpAddressRangePatch: Use this API command to modify IP Address Range of a AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
// - requestBody: *ApRegistrationRulesModifyIpAddressRangePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesModifyIpAddressRangePatch(ctx context.Context, id string, requestBody *ApRegistrationRulesModifyIpAddressRangePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/apRules/{id}/ipAddressRange", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	ApRegistrationRulesModifyMobilityzonePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApRegistrationRulesModifyMobilityzonePatch: Use this API command to modify  mobilityZone of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
// - requestBody: *ApRegistrationRulesModifyMobilityzonePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesModifyMobilityzonePatch(ctx context.Context, id string, requestBody *ApRegistrationRulesModifyMobilityzonePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/apRules/{id}/mobilityZone", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	ApRegistrationRulesModifySubnetPatchRequest struct {
		NetworkAddress *string `json:"networkAddress,omitempty"`
		SubnetMask     *string `json:"subnetMask,omitempty"`
	}
)

// ApRegistrationRulesModifySubnetPatch: Use this API command to modify subnet of AP Registration Rules profile.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): AP Rule ID
// - requestBody: *ApRegistrationRulesModifySubnetPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APRulesAPI) ApRegistrationRulesModifySubnetPatch(ctx context.Context, id string, requestBody *ApRegistrationRulesModifySubnetPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/apRules/{id}/subnet", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("id", id)
	return a.client.Ensure(ctx, request, 204, nil)
}
