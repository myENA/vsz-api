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

type ControlPlanes struct {
	client *Client
}
type (
	ControlPlanesRetrieveListGet200ResponseListSlice []*ControlPlanesRetrieveListGet200ResponseList

	ControlPlanesRetrieveListGet200ResponseList struct {
		ClusterIP    *string  `json:"clusterIp,omitempty"`
		ClusterRole  *string  `json:"clusterRole,omitempty"`
		ControlIP    *string  `json:"controlIp,omitempty"`
		Description  *string  `json:"description,omitempty"`
		Firmware     *string  `json:"firmware,omitempty"`
		ManagementIP *string  `json:"managementIp,omitempty"`
		Model        *string  `json:"model,omitempty"`
		Name         *string  `json:"name,omitempty"`
		NumOfAps     *float64 `json:"numOfAps,omitempty"`
		SerialNumber *string  `json:"serialNumber,omitempty"`
		UpTime       *string  `json:"upTime,omitempty"`
		Uptime       *string  `json:"uptime,omitempty"`
	}

	ControlPlanesRetrieveListGet200Response struct {
		FirstIndex *int                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                            `json:"hasMore,omitempty"`
		List       ControlPlanesRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                             `json:"totalCount,omitempty"`
	}
)

// ControlPlanesRetrieveListGet: Use this API command to retrieve the list of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ControlPlanesRetrieveListGet200Response
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesRetrieveListGet(ctx context.Context) (*http.Response, *ControlPlanesRetrieveListGet200Response, error) {
	request := c.client.newRequest(ctx, "GET", "/v5_0/controlPlanes")
	request.authenticated = true

	out := &ControlPlanesRetrieveListGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ControlPlanesModifyIpSupportPatchRequest struct {
		IPMode *string `json:"ipMode,omitempty"`
	}
)

// ControlPlanesModifyIpSupportPatch: Use this API command to modify ip support of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ControlPlanesModifyIpSupportPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpSupportPatch(ctx context.Context, requestBody *ControlPlanesModifyIpSupportPatchRequest) (*http.Response, []byte, error) {
	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/ipSupport")
	request.body = requestBody
	request.authenticated = true

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesRetrieveGet200ResponseIpv4AccessAndCoreSeparation struct {
		DefaultGateway     *string `json:"defaultGateway,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	}

	ControlPlanesRetrieveGet200ResponseIpv4ClusterInterface struct {
		Gateway    *string `json:"gateway,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		IPMode     *string `json:"ipMode,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
	}

	ControlPlanesRetrieveGet200ResponseIpv4ControlInterface struct {
		Gateway    *string `json:"gateway,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		IPMode     *string `json:"ipMode,omitempty"`
		NatIP      *string `json:"natIp,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
	}

	ControlPlanesRetrieveGet200ResponseIpv4ManagementInterface struct {
		Gateway    *string `json:"gateway,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		IPMode     *string `json:"ipMode,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
	}

	ControlPlanesRetrieveGet200ResponseIpv6AccessAndCoreSeparation struct {
		DefaultGateway     *string `json:"defaultGateway,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	}

	ControlPlanesRetrieveGet200ResponseIpv6ControlInterface struct {
		Gateway   *string `json:"gateway,omitempty"`
		IPAddress *string `json:"ipAddress,omitempty"`
		IPMode    *string `json:"ipMode,omitempty"`
	}

	ControlPlanesRetrieveGet200ResponseIpv6ManagementInterface struct {
		Gateway   *string `json:"gateway,omitempty"`
		IPAddress *string `json:"ipAddress,omitempty"`
		IPMode    *string `json:"ipMode,omitempty"`
	}

	ControlPlanesRetrieveGet200Response struct {
		EnableAccessAndCoreSeparation *bool                                                           `json:"enableAccessAndCoreSeparation,omitempty"`
		IPMode                        *string                                                         `json:"ipMode,omitempty"`
		Ipv4AccessAndCoreSeparation   *ControlPlanesRetrieveGet200ResponseIpv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`
		Ipv4ClusterInterface          *ControlPlanesRetrieveGet200ResponseIpv4ClusterInterface        `json:"ipv4ClusterInterface,omitempty"`
		Ipv4ControlInterface          *ControlPlanesRetrieveGet200ResponseIpv4ControlInterface        `json:"ipv4ControlInterface,omitempty"`
		Ipv4ManagementInterface       *ControlPlanesRetrieveGet200ResponseIpv4ManagementInterface     `json:"ipv4ManagementInterface,omitempty"`
		Ipv6AccessAndCoreSeparation   *ControlPlanesRetrieveGet200ResponseIpv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`
		Ipv6ControlInterface          *ControlPlanesRetrieveGet200ResponseIpv6ControlInterface        `json:"ipv6ControlInterface,omitempty"`
		Ipv6ManagementInterface       *ControlPlanesRetrieveGet200ResponseIpv6ManagementInterface     `json:"ipv6ManagementInterface,omitempty"`
	}
)

// ControlPlanesRetrieveGet: Use this API command to retrieve control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ControlPlanesRetrieveGet200Response
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesRetrieveGet(ctx context.Context, bladeUUID string) (*http.Response, *ControlPlanesRetrieveGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "GET", "/v5_0/controlPlanes/{bladeUUID}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	out := &ControlPlanesRetrieveGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ControlPlanesModifyBasicPatchRequest struct {
		EnableAccessAndCoreSeparation *bool `json:"enableAccessAndCoreSeparation,omitempty"`
	}
)

// ControlPlanesModifyBasicPatch: Use this API command to modify the basic information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyBasicPatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyBasicPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv4AccessAndCoreSeparationPatchRequest struct {
		DefaultGateway     *string `json:"defaultGateway,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	}
)

// ControlPlanesModifyIpv4AccessAndCoreSeparationPatch: Use this API command to modify IPv4 access and core information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv4AccessAndCoreSeparationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv4AccessAndCoreSeparationPatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv4AccessAndCoreSeparationPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv4AccessAndCoreSeparation")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv4ClusterInterfacePatchRequest struct {
		Gateway    *string `json:"gateway,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		IPMode     *string `json:"ipMode,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
	}
)

// ControlPlanesModifyIpv4ClusterInterfacePatch: Use this API command to modify IPv4 cluster information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv4ClusterInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv4ClusterInterfacePatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv4ClusterInterfacePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv4ClusterInterface")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv4ControlInterfacePatchRequest struct {
		Gateway    *string `json:"gateway,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		IPMode     *string `json:"ipMode,omitempty"`
		NatIP      *string `json:"natIp,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
	}
)

// ControlPlanesModifyIpv4ControlInterfacePatch: Use this API command to modify IPv4 control information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv4ControlInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv4ControlInterfacePatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv4ControlInterfacePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv4ControlInterface")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv4ManagementInterfacePatchRequest struct {
		Gateway    *string `json:"gateway,omitempty"`
		IPAddress  *string `json:"ipAddress,omitempty"`
		IPMode     *string `json:"ipMode,omitempty"`
		SubnetMask *string `json:"subnetMask,omitempty"`
	}
)

// ControlPlanesModifyIpv4ManagementInterfacePatch: Use this API command to modify IPv4 management information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv4ManagementInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv4ManagementInterfacePatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv4ManagementInterfacePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv4ManagementInterface")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv6AccessAndCoreSeparationPatchRequest struct {
		DefaultGateway     *string `json:"defaultGateway,omitempty"`
		PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
		SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	}
)

// ControlPlanesModifyIpv6AccessAndCoreSeparationPatch: Use this API command to modify IPv6 access and core information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv6AccessAndCoreSeparationPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv6AccessAndCoreSeparationPatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv6AccessAndCoreSeparationPatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv6AccessAndCoreSeparation")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv6ControlInterfacePatchRequest struct {
		Gateway   *string `json:"gateway,omitempty"`
		IPAddress *string `json:"ipAddress,omitempty"`
		IPMode    *string `json:"ipMode,omitempty"`
	}
)

// ControlPlanesModifyIpv6ControlInterfacePatch: Use this API command to modify IPv6 control information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv6ControlInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv6ControlInterfacePatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv6ControlInterfacePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv6ControlInterface")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesModifyIpv6ManagementInterfacePatchRequest struct {
		Gateway   *string `json:"gateway,omitempty"`
		IPAddress *string `json:"ipAddress,omitempty"`
		IPMode    *string `json:"ipMode,omitempty"`
	}
)

// ControlPlanesModifyIpv6ManagementInterfacePatch: Use this API command to modify IPv6 management information of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyIpv6ManagementInterfacePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyIpv6ManagementInterfacePatch(ctx context.Context, bladeUUID string, requestBody *ControlPlanesModifyIpv6ManagementInterfacePatchRequest) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/ipv6ManagementInterface")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

// ControlPlanesDeleteStaticRouteDelete: Use this API command to delete the static route of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesDeleteStaticRouteDelete(ctx context.Context, bladeUUID string) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "DELETE", "/v5_0/controlPlanes/{bladeUUID}/staticRoutes")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}

type (
	ControlPlanesRetrieveStaticRouteGet200ResponseStaticRoutesSlice []*ControlPlanesRetrieveStaticRouteGet200ResponseStaticRoutes

	ControlPlanesRetrieveStaticRouteGet200ResponseStaticRoutes struct {
		Gateway        *string  `json:"gateway,omitempty"`
		InterfaceMode  *string  `json:"interfaceMode,omitempty"`
		Metric         *float64 `json:"metric,omitempty"`
		NetworkAddress *string  `json:"networkAddress,omitempty"`
		SubnetMask     *string  `json:"subnetMask,omitempty"`
	}

	ControlPlanesRetrieveStaticRouteGet200Response struct {
		StaticRoutes ControlPlanesRetrieveStaticRouteGet200ResponseStaticRoutesSlice `json:"staticRoutes,omitempty"`
	}
)

// ControlPlanesRetrieveStaticRouteGet: Use this API command to retrieve static route of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ControlPlanesRetrieveStaticRouteGet200Response
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesRetrieveStaticRouteGet(ctx context.Context, bladeUUID string) (*http.Response, *ControlPlanesRetrieveStaticRouteGet200Response, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "GET", "/v5_0/controlPlanes/{bladeUUID}/staticRoutes")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	out := &ControlPlanesRetrieveStaticRouteGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ControlPlanesModifyStaticRoutePatchRequestSlice []*ControlPlanesModifyStaticRoutePatchRequest

	ControlPlanesModifyStaticRoutePatchRequest struct {
		Gateway        *string  `json:"gateway,omitempty"`
		InterfaceMode  *string  `json:"interfaceMode,omitempty"`
		Metric         *float64 `json:"metric,omitempty"`
		NetworkAddress *string  `json:"networkAddress,omitempty"`
		SubnetMask     *string  `json:"subnetMask,omitempty"`
	}
)

// ControlPlanesModifyStaticRoutePatch: Use this API command to modify the static route of control plane.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - bladeUUID (UUIDv4)
// - requestBody: *ControlPlanesModifyStaticRoutePatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ControlPlanes) ControlPlanesModifyStaticRoutePatch(ctx context.Context, bladeUUID string, requestBody ControlPlanesModifyStaticRoutePatchRequestSlice) (*http.Response, []byte, error) {
	var err error

	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}

	request := c.client.newRequest(ctx, "PATCH", "/v5_0/controlPlanes/{bladeUUID}/staticRoutes")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}

	return c.client.doRequest(request, 204, nil)
}
