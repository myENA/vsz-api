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
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type ClientsAPI struct {
	client *Client
}
type (
	WirelessClientBulkDeauthClientPostRequestClientListSlice []*WirelessClientBulkDeauthClientPostRequestClientList

	WirelessClientBulkDeauthClientPostRequestClientList struct {
		ApMac *string `json:"apMac,omitempty"`
		Mac   *string `json:"mac,omitempty"`
	}

	WirelessClientBulkDeauthClientPostRequest struct {
		ClientList WirelessClientBulkDeauthClientPostRequestClientListSlice `json:"clientList,omitempty"`
	}
)

// WirelessClientBulkDeauthClientPost: Use this API command to bulk deauth client.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *WirelessClientBulkDeauthClientPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClientsAPI) WirelessClientBulkDeauthClientPost(ctx context.Context, requestBody *WirelessClientBulkDeauthClientPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/clients/bulkDeauth", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return c.client.Ensure(ctx, request, 204, nil)
}

type (
	WirelessClientBulkDisconnectClientPostRequestClientListSlice []*WirelessClientBulkDisconnectClientPostRequestClientList

	WirelessClientBulkDisconnectClientPostRequestClientList struct {
		ApMac *string `json:"apMac,omitempty"`
		Mac   *string `json:"mac,omitempty"`
	}

	WirelessClientBulkDisconnectClientPostRequest struct {
		ClientList WirelessClientBulkDisconnectClientPostRequestClientListSlice `json:"clientList,omitempty"`
	}
)

// WirelessClientBulkDisconnectClientPost: Use this API command to bulk disconnect client.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *WirelessClientBulkDisconnectClientPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClientsAPI) WirelessClientBulkDisconnectClientPost(ctx context.Context, requestBody *WirelessClientBulkDisconnectClientPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/clients/bulkDisconnect", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return c.client.Ensure(ctx, request, 204, nil)
}

type (
	WirelessClientQueryClientByWlanNamePostRequestAttributesSlice []*string

	WirelessClientQueryClientByWlanNamePostRequestExtraFiltersSlice []*WirelessClientQueryClientByWlanNamePostRequestExtraFilters

	WirelessClientQueryClientByWlanNamePostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestExtraNotFiltersSlice []*WirelessClientQueryClientByWlanNamePostRequestExtraNotFilters

	WirelessClientQueryClientByWlanNamePostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestFiltersSlice []*WirelessClientQueryClientByWlanNamePostRequestFilters

	WirelessClientQueryClientByWlanNamePostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestFullTextSearchFieldsSlice []*string

	WirelessClientQueryClientByWlanNamePostRequestFullTextSearch struct {
		Fields WirelessClientQueryClientByWlanNamePostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                 `json:"type,omitempty"`
		Value  *string                                                                 `json:"value,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                     `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                   `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                     `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                     `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                   `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                     `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                     `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                     `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                     `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                     `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                     `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                   `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                     `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                   `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                   `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                   `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                   `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *WirelessClientQueryClientByWlanNamePostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                   `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                     `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                     `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                     `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *WirelessClientQueryClientByWlanNamePostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                   `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                   `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                   `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                   `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                   `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                   `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                   `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                   `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                   `json:"localUser_userSource,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePostRequest struct {
		Attributes      WirelessClientQueryClientByWlanNamePostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                            `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                              `json:"expandDomains,omitempty"`
		ExtraFilters    WirelessClientQueryClientByWlanNamePostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters WirelessClientQueryClientByWlanNamePostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *WirelessClientQueryClientByWlanNamePostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         WirelessClientQueryClientByWlanNamePostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *WirelessClientQueryClientByWlanNamePostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                               `json:"limit,omitempty"`
		Options         *WirelessClientQueryClientByWlanNamePostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                               `json:"page,omitempty"`
		Query           *string                                                            `json:"query,omitempty"`
		SortInfo        *WirelessClientQueryClientByWlanNamePostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                               `json:"start,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePost200ResponseListSlice []*WirelessClientQueryClientByWlanNamePost200ResponseList

	WirelessClientQueryClientByWlanNamePost200ResponseList struct {
		Alerts           *float64 `json:"alerts,omitempty"`
		ApMac            *string  `json:"apMac,omitempty"`
		ApName           *string  `json:"apName,omitempty"`
		AuthMethod       *string  `json:"authMethod,omitempty"`
		AuthStatus       *string  `json:"authStatus,omitempty"`
		Channel          *float64 `json:"channel,omitempty"`
		ClientMac        *string  `json:"clientMac,omitempty"`
		ControlPlaneName *string  `json:"controlPlaneName,omitempty"`
		DataPlaneName    *string  `json:"dataPlaneName,omitempty"`
		Downlink         *float64 `json:"downlink,omitempty"`
		DownlinkRate     *float64 `json:"downlinkRate,omitempty"`
		EncryptionMethod *string  `json:"encryptionMethod,omitempty"`
		Hostname         *string  `json:"hostname,omitempty"`
		IPAddress        *string  `json:"ipAddress,omitempty"`
		Ipv6Address      *string  `json:"ipv6Address,omitempty"`
		OsType           *string  `json:"osType,omitempty"`
		RadioType        *string  `json:"radioType,omitempty"`
		Role             *string  `json:"role,omitempty"`
		Rssi             *float64 `json:"rssi,omitempty"`
		RxBytes          *int     `json:"rxBytes,omitempty"`
		RxFrames         *int     `json:"rxFrames,omitempty"`
		SessionStartTime *float64 `json:"sessionStartTime,omitempty"`
		Snr              *float64 `json:"snr,omitempty"`
		Speedflex        *float64 `json:"speedflex,omitempty"`
		Ssid             *string  `json:"ssid,omitempty"`
		Status           *string  `json:"status,omitempty"`
		Traffic          *float64 `json:"traffic,omitempty"`
		TxBytes          *int     `json:"txBytes,omitempty"`
		TxDropDataFrames *int     `json:"txDropDataFrames,omitempty"`
		TxFrames         *int     `json:"txFrames,omitempty"`
		TxRxBytes        *int     `json:"txRxBytes,omitempty"`
		Uplink           *float64 `json:"uplink,omitempty"`
		UplinkRate       *float64 `json:"uplinkRate,omitempty"`
		UserName         *string  `json:"userName,omitempty"`
		Vlan             *float64 `json:"vlan,omitempty"`
		WlanType         *string  `json:"wlanType,omitempty"`
		ZoneID           *string  `json:"zoneId,omitempty"`
		ZoneVersion      *string  `json:"zoneVersion,omitempty"`
	}

	WirelessClientQueryClientByWlanNamePost200Response struct {
		Extra      *json.RawMessage                                            `json:"extra,omitempty"`
		FirstIndex *int                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                       `json:"hasMore,omitempty"`
		List       WirelessClientQueryClientByWlanNamePost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                        `json:"totalCount,omitempty"`
	}
)

// WirelessClientQueryClientByWlanNamePost: Use this API command to query client by wlan name.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - wlanname (string)
// - requestBody: *WirelessClientQueryClientByWlanNamePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WirelessClientQueryClientByWlanNamePost200Response
// - error: Error seen or nil if none
func (c *ClientsAPI) WirelessClientQueryClientByWlanNamePost(ctx context.Context, wlanname string, requestBody *WirelessClientQueryClientByWlanNamePostRequest) (*http.Response, *WirelessClientQueryClientByWlanNamePost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(wlanname)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"wlanname\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/clients/byWlanName/{wlanname}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("wlanname", wlanname)
	out := &WirelessClientQueryClientByWlanNamePost200Response{}
	httpResponse, _, err := c.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WirelessClientDeauthClientPostRequest struct {
		ApMac *string `json:"apMac,omitempty"`
		Mac   *string `json:"mac,omitempty"`
	}
)

// WirelessClientDeauthClientPost: Use this API command to deauth client.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *WirelessClientDeauthClientPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClientsAPI) WirelessClientDeauthClientPost(ctx context.Context, requestBody *WirelessClientDeauthClientPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/clients/deauth", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return c.client.Ensure(ctx, request, 204, nil)
}

type (
	WirelessClientDisconnectClientPostRequest struct {
		ApMac *string `json:"apMac,omitempty"`
		Mac   *string `json:"mac,omitempty"`
	}
)

// WirelessClientDisconnectClientPost: Use this API command to disconnect client.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *WirelessClientDisconnectClientPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClientsAPI) WirelessClientDisconnectClientPost(ctx context.Context, requestBody *WirelessClientDisconnectClientPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/clients/disconnect", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return c.client.Ensure(ctx, request, 204, nil)
}
