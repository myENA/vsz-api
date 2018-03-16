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

type SystemAPI struct {
	client *Client
}
type (
	TestAaaServerTestAaaPostRequestAaaServer struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	TestAaaServerTestAaaPostRequest struct {
		AaaServer    *TestAaaServerTestAaaPostRequestAaaServer `json:"aaaServer,omitempty"`
		AaaType      *string                                   `json:"aaaType,omitempty"`
		AuthProtocol *string                                   `json:"authProtocol,omitempty"`
		Password     *string                                   `json:"password,omitempty"`
		ServerType   *string                                   `json:"serverType,omitempty"`
		UserName     *string                                   `json:"userName,omitempty"`
	}

	TestAaaServerTestAaaPost200Response struct {
		PrimaryServer   *string `json:"primaryServer,omitempty"`
		SecondaryServer *string `json:"secondaryServer,omitempty"`
	}
)

// TestAaaServerTestAaaPost: Use this API command to test AAA server.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *TestAaaServerTestAaaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *TestAaaServerTestAaaPost200Response
// - error: Error seen or nil if none
func (s *SystemAPI) TestAaaServerTestAaaPost(ctx context.Context, requestBody *TestAaaServerTestAaaPostRequest) (*http.Response, *TestAaaServerTestAaaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/system/aaa/test", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &TestAaaServerTestAaaPost200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// SystemExecuteApBalancePost: Execute ap balance
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SystemExecuteApBalancePost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/system/ap_balance", true)
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SystemGetApRecoverySettingGet200Response struct {
		Key                *string `json:"key,omitempty"`
		RecoverySsidEnable *bool   `json:"recoverySsidEnable,omitempty"`
		RecoverySsidPskKey *string `json:"recoverySsidPskKey,omitempty"`
	}
)

// SystemGetApRecoverySettingGet: Use this API command to get system level ap recovery ssid information.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemGetApRecoverySettingGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SystemGetApRecoverySettingGet(ctx context.Context) (*http.Response, *SystemGetApRecoverySettingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/system/ap_recovery", true)
	out := &SystemGetApRecoverySettingGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SystemModifyApRecoverySettingPostRequest struct {
		Enabled            *bool   `json:"enabled,omitempty"`
		RecoverySsidEnable *bool   `json:"recoverySsidEnable,omitempty"`
		RecoverySsidPskKey *string `json:"recoverySsidPskKey,omitempty"`
		Rsm                *string `json:"rsm,omitempty"`
	}
)

// SystemModifyApRecoverySettingPost: Use this API command to patch system level ap recovery ssid.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SystemModifyApRecoverySettingPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SystemModifyApRecoverySettingPost(ctx context.Context, requestBody *SystemModifyApRecoverySettingPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/system/ap_recovery", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SystemApModelsGet200Response []*string
)

// SystemApModelsGet: Use this API command to retrieve AP models.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - SystemApModelsGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SystemApModelsGet(ctx context.Context) (*http.Response, SystemApModelsGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/system/apmodels", true)
	out := make(SystemApModelsGet200Response, 0)
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, &(out))
	return httpResponse, out, err
}

type (
	SystemGetApRoutineStatusIntervalSettingGet200Response struct {
		ApRoutineStatusInterval *float64 `json:"apRoutineStatusInterval,omitempty"`
	}
)

// SystemGetApRoutineStatusIntervalSettingGet: Use this API command to get AP routine status interval setting.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemGetApRoutineStatusIntervalSettingGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SystemGetApRoutineStatusIntervalSettingGet(ctx context.Context) (*http.Response, *SystemGetApRoutineStatusIntervalSettingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/system/apRoutineStatusInterval", true)
	out := &SystemGetApRoutineStatusIntervalSettingGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// SystemIncreaseApRoutineStatusIntervalPost: Use this API command to set AP routine status interval setting to 900 seconds.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SystemIncreaseApRoutineStatusIntervalPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/system/apRoutineStatusInterval/slowdown", true)
	return s.client.Ensure(ctx, request, 204, nil)
}

// SystemDecreaseApRoutineStatusIntervalPost: Use this API command to set AP routine status interval setting to 180 seconds.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SystemDecreaseApRoutineStatusIntervalPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/system/apRoutineStatusInterval/speedup", true)
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	CaleaGetCaleaCommonSettingGet200Response struct {
		CaleaServerIP *string `json:"caleaServerIp,omitempty"`
		DcIP          *string `json:"dc_ip,omitempty"`
	}
)

// CaleaGetCaleaCommonSettingGet: Use this API command to get CALEA common setting.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CaleaGetCaleaCommonSettingGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaGetCaleaCommonSettingGet(ctx context.Context) (*http.Response, *CaleaGetCaleaCommonSettingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/system/caleaCommonSetting", true)
	out := &CaleaGetCaleaCommonSettingGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	CaleaSetCaleaCommonSettingPostRequest struct {
		CaleaServerIP *string `json:"caleaServerIp,omitempty"`
		DcIP          *string `json:"dc_ip,omitempty"`
	}
)

// CaleaSetCaleaCommonSettingPost: Use this API command to set CALEA common setting.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *CaleaSetCaleaCommonSettingPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaSetCaleaCommonSettingPost(ctx context.Context, requestBody *CaleaSetCaleaCommonSettingPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/system/caleaCommonSetting", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	CaleaDeleteCaleaUeMacListDelete1RequestMacListSlice []*string

	CaleaDeleteCaleaUeMacListDelete1Request struct {
		MacList CaleaDeleteCaleaUeMacListDelete1RequestMacListSlice `json:"macList,omitempty"`
	}
)

// CaleaDeleteCaleaUeMacListDelete1: Use this API command to delete CALEA UE mac list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *CaleaDeleteCaleaUeMacListDelete1Request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaDeleteCaleaUeMacListDelete1(ctx context.Context, requestBody *CaleaDeleteCaleaUeMacListDelete1Request) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("DELETE", "/v5_0/system/caleaMac", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	CaleaAddCaleaUeMacListPostRequestMacListSlice []*string

	CaleaAddCaleaUeMacListPostRequest struct {
		MacList CaleaAddCaleaUeMacListPostRequestMacListSlice `json:"macList,omitempty"`
	}
)

// CaleaAddCaleaUeMacListPost: Use this API command to add CALEA UE mac list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *CaleaAddCaleaUeMacListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaAddCaleaUeMacListPost(ctx context.Context, requestBody *CaleaAddCaleaUeMacListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/system/caleaMac", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

// CaleaDeleteCaleaUeMacListDelete: Use this API command to delete all CALEA UE mac list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaDeleteCaleaUeMacListDelete(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("DELETE", "/v5_0/system/caleaMacList", true)
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	CaleaGetCaleaUeMacListGet200ResponseListSlice []*string

	CaleaGetCaleaUeMacListGet200Response struct {
		FirstIndex *int                                          `json:"firstIndex,omitempty"`
		HasMore    *bool                                         `json:"hasMore,omitempty"`
		List       CaleaGetCaleaUeMacListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                          `json:"totalCount,omitempty"`
	}
)

// CaleaGetCaleaUeMacListGet: Use this API command to get all CALEA UE mac list.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CaleaGetCaleaUeMacListGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaGetCaleaUeMacListGet(ctx context.Context) (*http.Response, *CaleaGetCaleaUeMacListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/system/caleaMacList", true)
	out := &CaleaGetCaleaUeMacListGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// CaleaUploadCaleaUeMacListCsvFilePost: Use this API command to upload CALEA UE mac list csv file.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) CaleaUploadCaleaUeMacListCsvFilePost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/system/caleaMacList", true)
	return s.client.Ensure(ctx, request, 200, nil)
}

type (
	SystemSystemInventoryGet200ResponseListSlice []*SystemSystemInventoryGet200ResponseList

	SystemSystemInventoryGet200ResponseList struct {
		Clients                     *float64    `json:"clients,omitempty"`
		ConnectedAPs                *float64    `json:"connectedAPs,omitempty"`
		ConnectedDownMeshAPs        *float64    `json:"connectedDownMeshAPs,omitempty"`
		ConnectedMeshAPs            *float64    `json:"connectedMeshAPs,omitempty"`
		ConnectedMeshDisabledAPs    *float64    `json:"connectedMeshDisabledAPs,omitempty"`
		ConnectedRootAPs            *float64    `json:"connectedRootAPs,omitempty"`
		ConnectedeMeshAPs           *float64    `json:"connectedeMeshAPs,omitempty"`
		DisconnectedAPs             *float64    `json:"disconnectedAPs,omitempty"`
		DisconnectedDownMeshAPs     *float64    `json:"disconnectedDownMeshAPs,omitempty"`
		DisconnectedMeshAPs         *float64    `json:"disconnectedMeshAPs,omitempty"`
		DisconnectedMeshDisabledAPs *float64    `json:"disconnectedMeshDisabledAPs,omitempty"`
		DisconnectedRootAPs         *float64    `json:"disconnectedRootAPs,omitempty"`
		DisconnectedeMeshAPs        *float64    `json:"disconnectedeMeshAPs,omitempty"`
		DiscoveryAPs                *float64    `json:"discoveryAPs,omitempty"`
		MeshSSID                    interface{} `json:"meshSSID,omitempty"`
		ProvisionedAPs              *float64    `json:"provisionedAPs,omitempty"`
		RebootingAPs                *float64    `json:"rebootingAPs,omitempty"`
		RebootingDownMeshAPs        *float64    `json:"rebootingDownMeshAPs,omitempty"`
		RebootingMeshAPs            *float64    `json:"rebootingMeshAPs,omitempty"`
		RebootingRootAPs            *float64    `json:"rebootingRootAPs,omitempty"`
		RebootingeMeshAPs           *float64    `json:"rebootingeMeshAPs,omitempty"`
		TotalAPs                    *float64    `json:"totalAPs,omitempty"`
		ZoneID                      *string     `json:"zoneId,omitempty"`
		ZoneName                    *string     `json:"zoneName,omitempty"`
	}

	SystemSystemInventoryGet200Response struct {
		FirstIndex *int                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                        `json:"hasMore,omitempty"`
		List       SystemSystemInventoryGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                         `json:"totalCount,omitempty"`
	}
)

// SystemSystemInventoryGet: Use this API command to retrieve the system inventory with current logon user domain.
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
// - *SystemSystemInventoryGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SystemSystemInventoryGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *SystemSystemInventoryGet200Response, error) {
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
	request := NewRequest("GET", "/v5_0/system/inventory", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &SystemSystemInventoryGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// SystemDisableNbiDelete: Use this API command to disable the user information by Northbound Portal Interface.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - domainId (UUIDv4): Domain ID.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SystemDisableNbiDelete(ctx context.Context, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("DELETE", "/v5_0/system/nbi", true)
	request.SetQueryParameter("domainId", domainId)
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SystemRetrieveNbiGet200Response struct {
		Password *string `json:"password,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}
)

// SystemRetrieveNbiGet: Use this API command to retrieve user information by Northbound Portal Interface.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - domainId (UUIDv4): Domain ID.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemRetrieveNbiGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SystemRetrieveNbiGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *SystemRetrieveNbiGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/system/nbi", true)
	request.SetQueryParameter("domainId", domainId)
	out := &SystemRetrieveNbiGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SystemModifyNbiBasicPatchRequest struct {
		Password *string `json:"password,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}
)

// SystemModifyNbiBasicPatch: Use this API command to modify the user information by Northbound Portal Interface.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SystemModifyNbiBasicPatchRequest
//
// Optional Parameter Map:
// - domainId (UUIDv4): Domain ID.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SystemModifyNbiBasicPatch(ctx context.Context, requestBody *SystemModifyNbiBasicPatchRequest, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("PATCH", "/v5_0/system/nbi", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetQueryParameter("domainId", domainId)
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2AgentSlice []*SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2Agent

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2AgentNotificationTargetSlice []*SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2AgentNotificationTarget

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2Agent struct {
		CommunityName       *string                                                                    `json:"communityName,omitempty"`
		NotificationEnabled *bool                                                                      `json:"notificationEnabled,omitempty"`
		NotificationTarget  SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                                    `json:"notificationType,omitempty"`
		ReadEnabled         *bool                                                                      `json:"readEnabled,omitempty"`
		WriteEnabled        *bool                                                                      `json:"writeEnabled,omitempty"`
	}

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3AgentSlice []*SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3Agent

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3AgentNotificationTargetSlice []*SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3AgentNotificationTarget

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3Agent struct {
		AuthPassword        *string                                                                    `json:"authPassword,omitempty"`
		AuthProtocol        *string                                                                    `json:"authProtocol,omitempty"`
		NotificationEnabled *bool                                                                      `json:"notificationEnabled,omitempty"`
		NotificationTarget  SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                                    `json:"notificationType,omitempty"`
		PrivPassword        *string                                                                    `json:"privPassword,omitempty"`
		PrivProtocol        *string                                                                    `json:"privProtocol,omitempty"`
		ReadEnabled         *bool                                                                      `json:"readEnabled,omitempty"`
		UserName            *string                                                                    `json:"userName,omitempty"`
		WriteEnabled        *bool                                                                      `json:"writeEnabled,omitempty"`
	}

	SnmpAgentRetrieveSnmpAgentGet200Response struct {
		SnmpNotificationEnabled *bool                                                    `json:"snmpNotificationEnabled,omitempty"`
		SnmpV2Agent             SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV2AgentSlice `json:"snmpV2Agent,omitempty"`
		SnmpV3Agent             SnmpAgentRetrieveSnmpAgentGet200ResponseSnmpV3AgentSlice `json:"snmpV3Agent,omitempty"`
	}
)

// SnmpAgentRetrieveSnmpAgentGet: Retrieve SNMP Agent sertting
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SnmpAgentRetrieveSnmpAgentGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SnmpAgentRetrieveSnmpAgentGet(ctx context.Context) (*http.Response, *SnmpAgentRetrieveSnmpAgentGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/system/snmpAgent", true)
	out := &SnmpAgentRetrieveSnmpAgentGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SnmpAgentModifySnmpAgentPutRequestSnmpV2AgentSlice []*SnmpAgentModifySnmpAgentPutRequestSnmpV2Agent

	SnmpAgentModifySnmpAgentPutRequestSnmpV2AgentNotificationTargetSlice []*SnmpAgentModifySnmpAgentPutRequestSnmpV2AgentNotificationTarget

	SnmpAgentModifySnmpAgentPutRequestSnmpV2AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	SnmpAgentModifySnmpAgentPutRequestSnmpV2Agent struct {
		CommunityName       *string                                                              `json:"communityName,omitempty"`
		NotificationEnabled *bool                                                                `json:"notificationEnabled,omitempty"`
		NotificationTarget  SnmpAgentModifySnmpAgentPutRequestSnmpV2AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                              `json:"notificationType,omitempty"`
		ReadEnabled         *bool                                                                `json:"readEnabled,omitempty"`
		WriteEnabled        *bool                                                                `json:"writeEnabled,omitempty"`
	}

	SnmpAgentModifySnmpAgentPutRequestSnmpV3AgentSlice []*SnmpAgentModifySnmpAgentPutRequestSnmpV3Agent

	SnmpAgentModifySnmpAgentPutRequestSnmpV3AgentNotificationTargetSlice []*SnmpAgentModifySnmpAgentPutRequestSnmpV3AgentNotificationTarget

	SnmpAgentModifySnmpAgentPutRequestSnmpV3AgentNotificationTarget struct {
		Address *string `json:"address,omitempty"`
		Port    *int    `json:"port,omitempty"`
	}

	SnmpAgentModifySnmpAgentPutRequestSnmpV3Agent struct {
		AuthPassword        *string                                                              `json:"authPassword,omitempty"`
		AuthProtocol        *string                                                              `json:"authProtocol,omitempty"`
		NotificationEnabled *bool                                                                `json:"notificationEnabled,omitempty"`
		NotificationTarget  SnmpAgentModifySnmpAgentPutRequestSnmpV3AgentNotificationTargetSlice `json:"notificationTarget,omitempty"`
		NotificationType    *string                                                              `json:"notificationType,omitempty"`
		PrivPassword        *string                                                              `json:"privPassword,omitempty"`
		PrivProtocol        *string                                                              `json:"privProtocol,omitempty"`
		ReadEnabled         *bool                                                                `json:"readEnabled,omitempty"`
		UserName            *string                                                              `json:"userName,omitempty"`
		WriteEnabled        *bool                                                                `json:"writeEnabled,omitempty"`
	}

	SnmpAgentModifySnmpAgentPutRequest struct {
		SnmpNotificationEnabled *bool                                              `json:"snmpNotificationEnabled,omitempty"`
		SnmpV2Agent             SnmpAgentModifySnmpAgentPutRequestSnmpV2AgentSlice `json:"snmpV2Agent,omitempty"`
		SnmpV3Agent             SnmpAgentModifySnmpAgentPutRequestSnmpV3AgentSlice `json:"snmpV3Agent,omitempty"`
	}
)

// SnmpAgentModifySnmpAgentPut: Modify syslog server setting
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SnmpAgentModifySnmpAgentPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SnmpAgentModifySnmpAgentPut(ctx context.Context, requestBody *SnmpAgentModifySnmpAgentPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PUT", "/v5_0/system/snmpAgent", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SyslogServerRetrieveSyslogGet200ResponsePrimaryServer struct {
		Host *string `json:"host,omitempty"`
		Port *int    `json:"port,omitempty"`
	}

	SyslogServerRetrieveSyslogGet200ResponsePriority struct {
		Critical      *string `json:"critical,omitempty"`
		Debug         *string `json:"debug,omitempty"`
		Informational *string `json:"informational,omitempty"`
		Major         *string `json:"major,omitempty"`
		Minor         *string `json:"minor,omitempty"`
		Warning       *string `json:"warning,omitempty"`
	}

	SyslogServerRetrieveSyslogGet200ResponseSecondaryServer struct {
		Host           *string `json:"host,omitempty"`
		Port           *int    `json:"port,omitempty"`
		RedundancyMode *string `json:"redundancyMode,omitempty"`
	}

	SyslogServerRetrieveSyslogGet200Response struct {
		AppLogFacility               *string                                                  `json:"appLogFacility,omitempty"`
		AppLogSeverity               *string                                                  `json:"appLogSeverity,omitempty"`
		AuditLogFacility             *string                                                  `json:"auditLogFacility,omitempty"`
		AuditLogSeverity             *string                                                  `json:"auditLogSeverity,omitempty"`
		Enabled                      *bool                                                    `json:"enabled,omitempty"`
		EventLogFacility             *string                                                  `json:"eventLogFacility,omitempty"`
		ForwardUEEventsMsgFormatType *string                                                  `json:"forwardUEEventsMsgFormatType,omitempty"`
		OtherLogSeverity             *string                                                  `json:"otherLogSeverity,omitempty"`
		PrimaryServer                *SyslogServerRetrieveSyslogGet200ResponsePrimaryServer   `json:"primaryServer,omitempty"`
		Priority                     *SyslogServerRetrieveSyslogGet200ResponsePriority        `json:"priority,omitempty"`
		SecondaryServer              *SyslogServerRetrieveSyslogGet200ResponseSecondaryServer `json:"secondaryServer,omitempty"`
	}
)

// SyslogServerRetrieveSyslogGet: Retrieve syslog server sertting
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
// - *SyslogServerRetrieveSyslogGet200Response
// - error: Error seen or nil if none
func (s *SystemAPI) SyslogServerRetrieveSyslogGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *SyslogServerRetrieveSyslogGet200Response, error) {
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
	request := NewRequest("GET", "/v5_0/system/syslog", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &SyslogServerRetrieveSyslogGet200Response{}
	httpResponse, _, err := s.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	SyslogServerModifySyslogPatchRequest struct {
		AppLogFacility               *string `json:"appLogFacility,omitempty"`
		AppLogSeverity               *string `json:"appLogSeverity,omitempty"`
		AuditLogFacility             *string `json:"auditLogFacility,omitempty"`
		AuditLogSeverity             *string `json:"auditLogSeverity,omitempty"`
		Enabled                      *bool   `json:"enabled,omitempty"`
		EventLogFacility             *string `json:"eventLogFacility,omitempty"`
		ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`
		OtherLogSeverity             *string `json:"otherLogSeverity,omitempty"`
	}
)

// SyslogServerModifySyslogPatch: Modify syslog server setting
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SyslogServerModifySyslogPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SyslogServerModifySyslogPatch(ctx context.Context, requestBody *SyslogServerModifySyslogPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/system/syslog", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SyslogServerModifyPrimaryServerPatchRequest struct {
		Host *string `json:"host,omitempty"`
		Port *int    `json:"port,omitempty"`
	}
)

// SyslogServerModifyPrimaryServerPatch: Modify Primary Server of syslog
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SyslogServerModifyPrimaryServerPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SyslogServerModifyPrimaryServerPatch(ctx context.Context, requestBody *SyslogServerModifyPrimaryServerPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/system/syslog/primaryServer", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SyslogServerModifyPriorityPatchRequest struct {
		Critical      *string `json:"critical,omitempty"`
		Debug         *string `json:"debug,omitempty"`
		Informational *string `json:"informational,omitempty"`
		Major         *string `json:"major,omitempty"`
		Minor         *string `json:"minor,omitempty"`
		Warning       *string `json:"warning,omitempty"`
	}
)

// SyslogServerModifyPriorityPatch: Modify Priority of syslog
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SyslogServerModifyPriorityPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SyslogServerModifyPriorityPatch(ctx context.Context, requestBody *SyslogServerModifyPriorityPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/system/syslog/priority", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}

type (
	SyslogServerModifySecondaryServerPatchRequest struct {
		Host           *string `json:"host,omitempty"`
		Port           *int    `json:"port,omitempty"`
		RedundancyMode *string `json:"redundancyMode,omitempty"`
	}
)

// SyslogServerModifySecondaryServerPatch: Modify Secondary Server of syslog
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *SyslogServerModifySecondaryServerPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *SystemAPI) SyslogServerModifySecondaryServerPatch(ctx context.Context, requestBody *SyslogServerModifySecondaryServerPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/system/syslog/secondaryServer", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return s.client.Ensure(ctx, request, 204, nil)
}
