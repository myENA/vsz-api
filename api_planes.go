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

type PlanesAPI struct {
	client *Client
}
type (
	DataPlaneRetrieveListGet200ResponseListSlice []*DataPlaneRetrieveListGet200ResponseList

	DataPlaneRetrieveListGet200ResponseList struct {
		BladeName    *string  `json:"bladeName,omitempty"`
		DpStatus     *string  `json:"dpStatus,omitempty"`
		FwVersion    *string  `json:"fwVersion,omitempty"`
		GreTunnels   *float64 `json:"greTunnels,omitempty"`
		ID           *string  `json:"id,omitempty"`
		IP           *string  `json:"ip,omitempty"`
		Ipv6         *string  `json:"ipv6,omitempty"`
		LastSeen     *string  `json:"lastSeen,omitempty"`
		Mac          *string  `json:"mac,omitempty"`
		ManagedBy    *string  `json:"managedBy,omitempty"`
		Model        *string  `json:"model,omitempty"`
		SerialNumber *string  `json:"serialNumber,omitempty"`
		Uptime       *string  `json:"uptime,omitempty"`
	}

	DataPlaneRetrieveListGet200Response struct {
		FirstIndex *int                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                        `json:"hasMore,omitempty"`
		List       DataPlaneRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                         `json:"totalCount,omitempty"`
	}
)

// DataPlaneRetrieveListGet: Use this API command to retrieve a list of data planes.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DataPlaneRetrieveListGet200Response
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneRetrieveListGet(ctx context.Context) (*http.Response, *DataPlaneRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/planes", true)
	out := &DataPlaneRetrieveListGet200Response{}
	httpResponse, _, err := p.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DataPlaneGetDpMeshTunnelSettingGet200Response struct {
		Encrypted *bool `json:"encrypted,omitempty"`
	}
)

// DataPlaneGetDpMeshTunnelSettingGet: Use this API command to get DP mesh tunnel setting.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DataPlaneGetDpMeshTunnelSettingGet200Response
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneGetDpMeshTunnelSettingGet(ctx context.Context) (*http.Response, *DataPlaneGetDpMeshTunnelSettingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/planes/dpTunnel/setting", true)
	out := &DataPlaneGetDpMeshTunnelSettingGet200Response{}
	httpResponse, _, err := p.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DataPlaneUpdateDpMeshTunnelSettingPutRequest struct {
		Encrypted *bool `json:"encrypted,omitempty"`
	}
)

// DataPlaneUpdateDpMeshTunnelSettingPut: Use this API command to update DP mesh tunnel setting.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *DataPlaneUpdateDpMeshTunnelSettingPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneUpdateDpMeshTunnelSettingPut(ctx context.Context, requestBody *DataPlaneUpdateDpMeshTunnelSettingPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PUT", "/v5_0/planes/dpTunnel/setting", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return p.client.Ensure(ctx, request, 204, nil)
}

type (
	DataPlaneRetrieveGet200ResponseIpv6PrimaryInterface struct {
		Gateway            *string `json:"gateway,omitempty"`
		IPAddress          *string `json:"ipAddress,omitempty"`
		IPMode             *string `json:"ipMode,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	}

	DataPlaneRetrieveGet200ResponsePrimaryInterface struct {
		Gateway            *string `json:"gateway,omitempty"`
		IPAddress          *string `json:"ipAddress,omitempty"`
		IPMode             *string `json:"ipMode,omitempty"`
		NatIP              *string `json:"natIp,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
		SubnetMask         *string `json:"subnetMask,omitempty"`
		Vlan               *string `json:"vlan,omitempty"`
	}

	DataPlaneRetrieveGet200ResponseSecondaryInterface struct {
		IPAddress  *string `json:"ipAddress,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
		Vlan       *string `json:"vlan,omitempty"`
	}

	DataPlaneRetrieveGet200ResponseStaticRouteSlice []*DataPlaneRetrieveGet200ResponseStaticRoute

	DataPlaneRetrieveGet200ResponseStaticRoute struct {
		Gateway        *string `json:"gateway,omitempty"`
		NetworkAddress *string `json:"networkAddress,omitempty"`
		SubnetMask     *string `json:"subnetMask,omitempty"`
	}

	DataPlaneRetrieveGet200Response struct {
		InterfaceMode        *string                                              `json:"interfaceMode,omitempty"`
		Ipv6PrimaryInterface *DataPlaneRetrieveGet200ResponseIpv6PrimaryInterface `json:"ipv6PrimaryInterface,omitempty"`
		PrimaryInterface     *DataPlaneRetrieveGet200ResponsePrimaryInterface     `json:"primaryInterface,omitempty"`
		SecondaryInterface   *DataPlaneRetrieveGet200ResponseSecondaryInterface   `json:"secondaryInterface,omitempty"`
		StaticRoute          DataPlaneRetrieveGet200ResponseStaticRouteSlice      `json:"staticRoute,omitempty"`
	}
)

// DataPlaneRetrieveGet: Use this API command to retrieve data plane by id.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DataPlaneRetrieveGet200Response
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneRetrieveGet(ctx context.Context, bladeUUID string) (*http.Response, *DataPlaneRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/planes/{bladeUUID}", true)
	request.SetPathParameter("bladeUUID", bladeUUID)
	out := &DataPlaneRetrieveGet200Response{}
	httpResponse, _, err := p.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	DataPlaneModifyBasicPatchRequest struct {
		InterfaceMode *string `json:"interfaceMode,omitempty"`
	}
)

// DataPlaneModifyBasicPatch: Use this API command to modify the basic information of data plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *DataPlaneModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneModifyBasicPatch(ctx context.Context, bladeUUID string, requestBody *DataPlaneModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/planes/{bladeUUID}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("bladeUUID", bladeUUID)
	return p.client.Ensure(ctx, request, 204, nil)
}

type (
	DataPlaneModifyIpv6PrimaryInterfacePatchRequest struct {
		Gateway            *string `json:"gateway,omitempty"`
		IPAddress          *string `json:"ipAddress,omitempty"`
		IPMode             *string `json:"ipMode,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	}
)

// DataPlaneModifyIpv6PrimaryInterfacePatch: Use this API command to modify ipv6 Primary Interface of data plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *DataPlaneModifyIpv6PrimaryInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneModifyIpv6PrimaryInterfacePatch(ctx context.Context, bladeUUID string, requestBody *DataPlaneModifyIpv6PrimaryInterfacePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/planes/{bladeUUID}/ipv6PrimaryInterface", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("bladeUUID", bladeUUID)
	return p.client.Ensure(ctx, request, 204, nil)
}

type (
	DataPlaneModifyPrimaryInterfacePatchRequest struct {
		Gateway            *string `json:"gateway,omitempty"`
		IPAddress          *string `json:"ipAddress,omitempty"`
		IPMode             *string `json:"ipMode,omitempty"`
		NatIP              *string `json:"natIp,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
		SubnetMask         *string `json:"subnetMask,omitempty"`
		Vlan               *string `json:"vlan,omitempty"`
	}
)

// DataPlaneModifyPrimaryInterfacePatch: Use this API command to modify primary interface of data plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *DataPlaneModifyPrimaryInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneModifyPrimaryInterfacePatch(ctx context.Context, bladeUUID string, requestBody *DataPlaneModifyPrimaryInterfacePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/planes/{bladeUUID}/primaryInterface", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("bladeUUID", bladeUUID)
	return p.client.Ensure(ctx, request, 204, nil)
}

type (
	DataPlaneModifySecondaryInterfacePatchRequest struct {
		IPAddress  *string `json:"ipAddress,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
		Vlan       *string `json:"vlan,omitempty"`
	}
)

// DataPlaneModifySecondaryInterfacePatch: Use this API command to modify secondary interface of data plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *DataPlaneModifySecondaryInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneModifySecondaryInterfacePatch(ctx context.Context, bladeUUID string, requestBody *DataPlaneModifySecondaryInterfacePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/planes/{bladeUUID}/secondaryInterface", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("bladeUUID", bladeUUID)
	return p.client.Ensure(ctx, request, 204, nil)
}

// DataPlaneDeleteStaticRouteDelete: Use this API command to delete static route.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneDeleteStaticRouteDelete(ctx context.Context, bladeUUID string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/planes/{bladeUUID}/staticRoute", true)
	request.SetPathParameter("bladeUUID", bladeUUID)
	return p.client.Ensure(ctx, request, 204, nil)
}

type (
	DataPlaneModifyStaticRoutePatchRequestSlice []*DataPlaneModifyStaticRoutePatchRequest

	DataPlaneModifyStaticRoutePatchRequest struct {
		Gateway        *string `json:"gateway,omitempty"`
		NetworkAddress *string `json:"networkAddress,omitempty"`
		SubnetMask     *string `json:"subnetMask,omitempty"`
	}
)

// DataPlaneModifyStaticRoutePatch: Use this API command to modify static route of data plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *DataPlaneModifyStaticRoutePatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (p *PlanesAPI) DataPlaneModifyStaticRoutePatch(ctx context.Context, bladeUUID string, requestBody DataPlaneModifyStaticRoutePatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/planes/{bladeUUID}/staticRoute", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("bladeUUID", bladeUUID)
	return p.client.Ensure(ctx, request, 204, nil)
}
