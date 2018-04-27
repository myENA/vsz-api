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

type QueryAPI struct {
	client *Client
}
type (
	QueryWithFilterQueryApsRequestAttributesSlice []string

	QueryWithFilterQueryApsRequestFiltersSlice []*QueryWithFilterQueryApsRequestFilters

	QueryWithFilterQueryApsRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryApsRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryApsRequestFullTextSearch struct {
		Fields QueryWithFilterQueryApsRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                 `json:"type,omitempty"`
		Value  *string                                                 `json:"value,omitempty"`
	}

	QueryWithFilterQueryApsRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryApsRequest struct {
		Attributes     QueryWithFilterQueryApsRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        QueryWithFilterQueryApsRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *QueryWithFilterQueryApsRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                          `json:"limit,omitempty"`
		Options        interface{}                                   `json:"options,omitempty"`
		Page           *int                                          `json:"page,omitempty"`
		SortInfo       *QueryWithFilterQueryApsRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                          `json:"start,omitempty"`
	}

	QueryWithFilterQueryAps200ResponseListSlice []*QueryWithFilterQueryAps200ResponseList

	QueryWithFilterQueryAps200ResponseList struct {
		AdministrativeState            *string     `json:"administrativeState,omitempty"`
		Airtime5G                      interface{} `json:"airtime5G,omitempty"`
		Airtime24G                     interface{} `json:"airtime24G,omitempty"`
		Alerts                         *float64    `json:"alerts,omitempty"`
		ApGroupID                      *string     `json:"apGroupId,omitempty"`
		ApGroupName                    *string     `json:"apGroupName,omitempty"`
		ApMac                          *string     `json:"apMac,omitempty"`
		CableModemResetSupported       *bool       `json:"cableModemResetSupported,omitempty"`
		CableModemSupported            *bool       `json:"cableModemSupported,omitempty"`
		Capacity                       interface{} `json:"capacity,omitempty"`
		Capacity24G                    interface{} `json:"capacity24G,omitempty"`
		Capacity50G                    interface{} `json:"capacity50G,omitempty"`
		Channel5G                      interface{} `json:"channel5G,omitempty"`
		Channel24G                     interface{} `json:"channel24G,omitempty"`
		ConfigOverride                 *bool       `json:"configOverride,omitempty"`
		ConfigurationStatus            *string     `json:"configurationStatus,omitempty"`
		ConnectionFailure              interface{} `json:"connectionFailure,omitempty"`
		ConnectionStatus               *string     `json:"connectionStatus,omitempty"`
		ControlBladeID                 *string     `json:"controlBladeId,omitempty"`
		ControlBladeName               *string     `json:"controlBladeName,omitempty"`
		CrashDump                      interface{} `json:"crashDump,omitempty"`
		Description                    *string     `json:"description,omitempty"`
		DeviceGps                      *string     `json:"deviceGps,omitempty"`
		DeviceName                     *string     `json:"deviceName,omitempty"`
		DomainID                       *string     `json:"domainId,omitempty"`
		DomainName                     interface{} `json:"domainName,omitempty"`
		DpIP                           interface{} `json:"dpIp,omitempty"`
		EnabledBonjourGateway          *bool       `json:"enabledBonjourGateway,omitempty"`
		ExtIP                          *string     `json:"extIp,omitempty"`
		ExtPort                        *string     `json:"extPort,omitempty"`
		FirmwareVersion                *string     `json:"firmwareVersion,omitempty"`
		IndoorMapID                    interface{} `json:"indoorMapId,omitempty"`
		IndoorMapLocation              interface{} `json:"indoorMapLocation,omitempty"`
		IndoorMapName                  interface{} `json:"indoorMapName,omitempty"`
		IndoorMapXy                    interface{} `json:"indoorMapXy,omitempty"`
		IP                             *string     `json:"ip,omitempty"`
		Ipv6Address                    interface{} `json:"ipv6Address,omitempty"`
		IsAirtimeUtilization24GFlagged *bool       `json:"isAirtimeUtilization24GFlagged,omitempty"`
		IsAirtimeUtilization50GFlagged *bool       `json:"isAirtimeUtilization50GFlagged,omitempty"`
		IsCapacity24GFlagged           *bool       `json:"isCapacity24GFlagged,omitempty"`
		IsCapacity50GFlagged           *bool       `json:"isCapacity50GFlagged,omitempty"`
		IsConnectionFailure24GFlagged  *bool       `json:"isConnectionFailure24GFlagged,omitempty"`
		IsConnectionFailure50GFlagged  *bool       `json:"isConnectionFailure50GFlagged,omitempty"`
		IsConnectionFailureFlagged     *bool       `json:"isConnectionFailureFlagged,omitempty"`
		IsConnectionTotalCountFlagged  interface{} `json:"isConnectionTotalCountFlagged,omitempty"`
		IsCriticalAp                   *bool       `json:"isCriticalAp,omitempty"`
		IsLatency24GFlagged            *bool       `json:"isLatency24GFlagged,omitempty"`
		IsLatency50GFlagged            *bool       `json:"isLatency50GFlagged,omitempty"`
		IsOverallHealthStatusFlagged   *bool       `json:"isOverallHealthStatusFlagged,omitempty"`
		LastSeen                       *float64    `json:"lastSeen,omitempty"`
		Latency24G                     interface{} `json:"latency24G,omitempty"`
		Latency50G                     interface{} `json:"latency50G,omitempty"`
		LbsStatus                      *string     `json:"lbsStatus,omitempty"`
		Location                       *string     `json:"location,omitempty"`
		ManagementVlan                 interface{} `json:"managementVlan,omitempty"`
		MeshMode                       *string     `json:"meshMode,omitempty"`
		MeshRole                       *string     `json:"meshRole,omitempty"`
		Model                          *string     `json:"model,omitempty"`
		Noise5G                        interface{} `json:"noise5G,omitempty"`
		Noise24G                       interface{} `json:"noise24G,omitempty"`
		NumClients                     *float64    `json:"numClients,omitempty"`
		NumClients5G                   *float64    `json:"numClients5G,omitempty"`
		NumClients24G                  *float64    `json:"numClients24G,omitempty"`
		ProvisionMethod                *string     `json:"provisionMethod,omitempty"`
		ProvisionStage                 interface{} `json:"provisionStage,omitempty"`
		RegistrationState              *string     `json:"registrationState,omitempty"`
		RegistrationTime               *float64    `json:"registrationTime,omitempty"`
		Retry5G                        interface{} `json:"retry5G,omitempty"`
		Retry24G                       interface{} `json:"retry24G,omitempty"`
		Rx                             interface{} `json:"rx,omitempty"`
		Serial                         *string     `json:"serial,omitempty"`
		Status                         *string     `json:"status,omitempty"`
		SwapInMac                      interface{} `json:"swapInMac,omitempty"`
		SwapOutMac                     interface{} `json:"swapOutMac,omitempty"`
		Tx                             interface{} `json:"tx,omitempty"`
		TxRx                           interface{} `json:"txRx,omitempty"`
		WlanGroup24ID                  *string     `json:"wlanGroup24Id,omitempty"`
		WlanGroup24Name                *string     `json:"wlanGroup24Name,omitempty"`
		WlanGroup50ID                  *string     `json:"wlanGroup50Id,omitempty"`
		WlanGroup50Name                *string     `json:"wlanGroup50Name,omitempty"`
		ZoneFirmwareVersion            *string     `json:"zoneFirmwareVersion,omitempty"`
		ZoneID                         *string     `json:"zoneId,omitempty"`
		ZoneName                       *string     `json:"zoneName,omitempty"`
	}

	QueryWithFilterQueryAps200Response struct {
		FirstIndex *int                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                       `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryAps200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                        `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryAps: Query aps with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryApsRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryAps200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryAps(ctx context.Context, requestBody *QueryWithFilterQueryApsRequest) (*http.Response, *QueryWithFilterQueryAps200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/ap", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryAps200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestAttributesSlice []string

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraFiltersSlice []*ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraFilters

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraNotFiltersSlice []*ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraNotFilters

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFiltersSlice []*ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFilters

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFullTextSearchFieldsSlice []string

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFullTextSearch struct {
		Fields ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                       `json:"type,omitempty"`
		Value  *string                                                                                       `json:"value,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                           `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                         `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                           `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                           `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                         `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                           `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                           `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                           `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                           `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                           `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                           `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                         `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                           `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                         `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                         `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                         `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                         `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                         `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                           `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                           `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                           `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                         `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                         `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                         `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                         `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                         `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                         `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                         `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                         `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                         `json:"localUser_userSource,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequest struct {
		Attributes      ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                  `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                    `json:"expandDomains,omitempty"`
		ExtraFilters    ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                                                 `json:"limit,omitempty"`
		Options         *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                     `json:"page,omitempty"`
		Query           *string                                                                                  `json:"query,omitempty"`
		SortInfo        *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                                                 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseListSlice []*ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseList

	ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseListApplicationRulesSlice []*ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseListApplicationRules

	ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseListApplicationRules struct {
		AppID              *string  `json:"appId,omitempty"`
		AppName            *string  `json:"appName,omitempty"`
		ApplicationType    *string  `json:"applicationType,omitempty"`
		CatID              *string  `json:"catId,omitempty"`
		CatName            *string  `json:"catName,omitempty"`
		ClassificationType *string  `json:"classificationType,omitempty"`
		Downlink           *float64 `json:"downlink,omitempty"`
		MarkingPriority    *string  `json:"markingPriority,omitempty"`
		MarkingType        *string  `json:"markingType,omitempty"`
		Priority           *float64 `json:"priority,omitempty"`
		RuleType           *string  `json:"ruleType,omitempty"`
		Uplink             *float64 `json:"uplink,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseList struct {
		ApplicationRules ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseListApplicationRulesSlice `json:"applicationRules,omitempty"`
		CreateDateTime   *int                                                                                              `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                           `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                           `json:"creatorUsername,omitempty"`
		Description      *string                                                                                           `json:"description,omitempty"`
		DomainID         *string                                                                                           `json:"domainId,omitempty"`
		ID               *string                                                                                           `json:"id,omitempty"`
		ModifiedDateTime *int                                                                                              `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                           `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                           `json:"modifierUsername,omitempty"`
		Name             *string                                                                                           `json:"name,omitempty"`
		TenantID         *string                                                                                           `json:"tenantId,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveListPost200Response struct {
		Extra      *json.RawMessage                                                                  `json:"extra,omitempty"`
		FirstIndex *int                                                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                             `json:"hasMore,omitempty"`
		List       ApplicationVisibilityControlApplicationPolicyRetrieveListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                              `json:"totalCount,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationPolicyRetrieveListPost: Use this API command to retrieve a list of AVC Application Policy profiles.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationPolicyRetrieveListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) ApplicationVisibilityControlApplicationPolicyRetrieveListPost(ctx context.Context, requestBody *ApplicationVisibilityControlApplicationPolicyRetrieveListPostRequest) (*http.Response, *ApplicationVisibilityControlApplicationPolicyRetrieveListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/applicationPolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &ApplicationVisibilityControlApplicationPolicyRetrieveListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryClientsPostRequestAttributesSlice []string

	QueryWithFilterQueryClientsPostRequestExtraFiltersSlice []*QueryWithFilterQueryClientsPostRequestExtraFilters

	QueryWithFilterQueryClientsPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryClientsPostRequestExtraNotFilters

	QueryWithFilterQueryClientsPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestFiltersSlice []*QueryWithFilterQueryClientsPostRequestFilters

	QueryWithFilterQueryClientsPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryClientsPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryClientsPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                         `json:"type,omitempty"`
		Value  *string                                                         `json:"value,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                             `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                           `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                             `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                             `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                           `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                             `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                             `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                             `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                             `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                             `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                             `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                           `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                             `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                           `json:"auth_type,omitempty"`
		ForwardingType                *string                                                           `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                           `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                           `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterQueryClientsPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                           `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                             `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                             `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                             `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryClientsPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                           `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                           `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                           `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                           `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                           `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                           `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                           `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                           `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                           `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryClientsPostRequest struct {
		Attributes      QueryWithFilterQueryClientsPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                    `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                      `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryClientsPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryClientsPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryClientsPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryClientsPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryClientsPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                       `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryClientsPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                       `json:"page,omitempty"`
		Query           *string                                                    `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryClientsPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                       `json:"start,omitempty"`
	}

	QueryWithFilterQueryClientsPost200ResponseListSlice []*QueryWithFilterQueryClientsPost200ResponseList

	QueryWithFilterQueryClientsPost200ResponseList struct {
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

	QueryWithFilterQueryClientsPost200Response struct {
		Extra      *json.RawMessage                                    `json:"extra,omitempty"`
		FirstIndex *int                                                `json:"firstIndex,omitempty"`
		HasMore    *bool                                               `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryClientsPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryClientsPost: Query clients with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryClientsPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryClientsPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryClientsPost(ctx context.Context, requestBody *QueryWithFilterQueryClientsPostRequest) (*http.Response, *QueryWithFilterQueryClientsPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/client", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryClientsPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryDpsksPostRequestAttributesSlice []string

	QueryWithFilterQueryDpsksPostRequestExtraFiltersSlice []*QueryWithFilterQueryDpsksPostRequestExtraFilters

	QueryWithFilterQueryDpsksPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryDpsksPostRequestExtraNotFilters

	QueryWithFilterQueryDpsksPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestFiltersSlice []*QueryWithFilterQueryDpsksPostRequestFilters

	QueryWithFilterQueryDpsksPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryDpsksPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryDpsksPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                       `json:"type,omitempty"`
		Value  *string                                                       `json:"value,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                           `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                         `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                           `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                           `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                         `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                           `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                           `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                           `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                           `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                           `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                           `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                         `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                           `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                         `json:"auth_type,omitempty"`
		ForwardingType                *string                                                         `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                         `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                         `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterQueryDpsksPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                         `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                           `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                           `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                           `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryDpsksPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                         `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                         `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                         `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                         `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                         `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                         `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                         `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                         `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                         `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryDpsksPostRequest struct {
		Attributes      QueryWithFilterQueryDpsksPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                  `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                    `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryDpsksPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryDpsksPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryDpsksPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryDpsksPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryDpsksPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                     `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryDpsksPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                     `json:"page,omitempty"`
		Query           *string                                                  `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryDpsksPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                     `json:"start,omitempty"`
	}

	QueryWithFilterQueryDpsksPost200ResponseListSlice []*QueryWithFilterQueryDpsksPost200ResponseList

	QueryWithFilterQueryDpsksPost200ResponseList struct {
		CreateDateTime      *int     `json:"createDateTime,omitempty"`
		CreationDateTime    *float64 `json:"creationDateTime,omitempty"`
		DomainID            *string  `json:"domainId,omitempty"`
		ExpirationStartTime *float64 `json:"expirationStartTime,omitempty"`
		ExpirationTime      *float64 `json:"expirationTime,omitempty"`
		Expired             *bool    `json:"expired,omitempty"`
		Group               *bool    `json:"group,omitempty"`
		Key                 *string  `json:"key,omitempty"`
		TenantID            *string  `json:"tenantId,omitempty"`
		TTL                 *float64 `json:"ttl,omitempty"`
		UeMac               *string  `json:"ueMac,omitempty"`
		UserName            *string  `json:"userName,omitempty"`
		UserRoleID          *string  `json:"userRoleId,omitempty"`
		VlanID              *float64 `json:"vlanId,omitempty"`
		WlanID              *string  `json:"wlanId,omitempty"`
		ZoneID              *string  `json:"zoneId,omitempty"`
	}

	QueryWithFilterQueryDpsksPost200Response struct {
		Extra      *json.RawMessage                                  `json:"extra,omitempty"`
		FirstIndex *int                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                             `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryDpsksPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                              `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryDpsksPost: Query DPSKs with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryDpsksPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryDpsksPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryDpsksPost(ctx context.Context, requestBody *QueryWithFilterQueryDpsksPostRequest) (*http.Response, *QueryWithFilterQueryDpsksPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/dpsk", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryDpsksPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	WirelessClientHistoricalClientPostRequestAttributesSlice []string

	WirelessClientHistoricalClientPostRequestExtraFiltersSlice []*WirelessClientHistoricalClientPostRequestExtraFilters

	WirelessClientHistoricalClientPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestExtraNotFiltersSlice []*WirelessClientHistoricalClientPostRequestExtraNotFilters

	WirelessClientHistoricalClientPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestFiltersSlice []*WirelessClientHistoricalClientPostRequestFilters

	WirelessClientHistoricalClientPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestFullTextSearchFieldsSlice []string

	WirelessClientHistoricalClientPostRequestFullTextSearch struct {
		Fields WirelessClientHistoricalClientPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                            `json:"type,omitempty"`
		Value  *string                                                            `json:"value,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *WirelessClientHistoricalClientPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *WirelessClientHistoricalClientPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                              `json:"localUser_userSource,omitempty"`
	}

	WirelessClientHistoricalClientPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	WirelessClientHistoricalClientPostRequest struct {
		Attributes      WirelessClientHistoricalClientPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    WirelessClientHistoricalClientPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters WirelessClientHistoricalClientPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *WirelessClientHistoricalClientPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         WirelessClientHistoricalClientPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *WirelessClientHistoricalClientPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                      `json:"limit,omitempty"`
		Options         *WirelessClientHistoricalClientPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                          `json:"page,omitempty"`
		Query           *string                                                       `json:"query,omitempty"`
		SortInfo        *WirelessClientHistoricalClientPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                      `json:"start,omitempty"`
	}

	WirelessClientHistoricalClientPost200ResponseListSlice []*WirelessClientHistoricalClientPost200ResponseList

	WirelessClientHistoricalClientPost200ResponseList struct {
		ApMac            *string  `json:"apMac,omitempty"`
		ClientMac        *string  `json:"clientMac,omitempty"`
		CoreNetworkType  *string  `json:"coreNetworkType,omitempty"`
		IPAddress        *string  `json:"ipAddress,omitempty"`
		Ipv6Address      *string  `json:"ipv6Address,omitempty"`
		MvnoID           *string  `json:"mvnoId,omitempty"`
		MvnoName         *string  `json:"mvnoName,omitempty"`
		RxBytes          *int     `json:"rxBytes,omitempty"`
		RxDrops          *int     `json:"rxDrops,omitempty"`
		RxFrames         *int     `json:"rxFrames,omitempty"`
		SessionEndTime   *float64 `json:"sessionEndTime,omitempty"`
		SessionStartTime *float64 `json:"sessionStartTime,omitempty"`
		Ssid             *string  `json:"ssid,omitempty"`
		TxBytes          *int     `json:"txBytes,omitempty"`
		TxDrops          *int     `json:"txDrops,omitempty"`
		TxFrames         *int     `json:"txFrames,omitempty"`
	}

	WirelessClientHistoricalClientPost200Response struct {
		Extra      *json.RawMessage                                       `json:"extra,omitempty"`
		FirstIndex *int                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                  `json:"hasMore,omitempty"`
		List       WirelessClientHistoricalClientPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                   `json:"totalCount,omitempty"`
	}
)

// WirelessClientHistoricalClientPost: Use this API command to retrive historical client.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *WirelessClientHistoricalClientPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WirelessClientHistoricalClientPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) WirelessClientHistoricalClientPost(ctx context.Context, requestBody *WirelessClientHistoricalClientPostRequest) (*http.Response, *WirelessClientHistoricalClientPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/historicalclient", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &WirelessClientHistoricalClientPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterRetrieveIndoormapListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveIndoormapListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveIndoormapListPostRequestExtraFilters

	QueryWithFilterRetrieveIndoormapListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveIndoormapListPostRequestExtraNotFilters

	QueryWithFilterRetrieveIndoormapListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestFiltersSlice []*QueryWithFilterRetrieveIndoormapListPostRequestFilters

	QueryWithFilterRetrieveIndoormapListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveIndoormapListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveIndoormapListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                  `json:"type,omitempty"`
		Value  *string                                                                  `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                      `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                    `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                      `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                      `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                    `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                      `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                      `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                      `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                      `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                      `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                      `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                    `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                      `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                    `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                    `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                    `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                    `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveIndoormapListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                    `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                      `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                      `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                      `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveIndoormapListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                    `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                    `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                    `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                    `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                    `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                    `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                    `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                    `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                    `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPostRequest struct {
		Attributes      QueryWithFilterRetrieveIndoormapListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                             `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                               `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveIndoormapListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveIndoormapListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveIndoormapListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveIndoormapListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveIndoormapListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveIndoormapListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                `json:"page,omitempty"`
		Query           *string                                                             `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveIndoormapListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPost200ResponseListSlice []*QueryWithFilterRetrieveIndoormapListPost200ResponseList

	QueryWithFilterRetrieveIndoormapListPost200ResponseListScaleA struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPost200ResponseListScaleB struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPost200ResponseListScale struct {
		A        *QueryWithFilterRetrieveIndoormapListPost200ResponseListScaleA `json:"a,omitempty"`
		B        *QueryWithFilterRetrieveIndoormapListPost200ResponseListScaleB `json:"b,omitempty"`
		Distance *float64                                                       `json:"distance,omitempty"`
		Unit     *string                                                        `json:"unit,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPost200ResponseList struct {
		Address       *string                                                       `json:"address,omitempty"`
		ApCount       *int                                                          `json:"apCount,omitempty"`
		ApGroupID     *string                                                       `json:"apGroupId,omitempty"`
		Description   *string                                                       `json:"description,omitempty"`
		DomainID      *string                                                       `json:"domainId,omitempty"`
		GroupType     *string                                                       `json:"groupType,omitempty"`
		ID            *string                                                       `json:"id,omitempty"`
		ImageData     *string                                                       `json:"imageData,omitempty"`
		ImageFileName *string                                                       `json:"imageFileName,omitempty"`
		Key           *string                                                       `json:"key,omitempty"`
		Latitude      *float64                                                      `json:"latitude,omitempty"`
		Longitude     *float64                                                      `json:"longitude,omitempty"`
		Name          *string                                                       `json:"name,omitempty"`
		Scale         *QueryWithFilterRetrieveIndoormapListPost200ResponseListScale `json:"scale,omitempty"`
		TenantID      *string                                                       `json:"tenantId,omitempty"`
		ZoneID        *string                                                       `json:"zoneId,omitempty"`
	}

	QueryWithFilterRetrieveIndoormapListPost200Response struct {
		Extra      *json.RawMessage                                             `json:"extra,omitempty"`
		FirstIndex *int                                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                                        `json:"hasMore,omitempty"`
		List       QueryWithFilterRetrieveIndoormapListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                         `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterRetrieveIndoormapListPost: Query indoorMap with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveIndoormapListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterRetrieveIndoormapListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveIndoormapListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveIndoormapListPostRequest) (*http.Response, *QueryWithFilterRetrieveIndoormapListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/indoorMap", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterRetrieveIndoormapListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryMeshNeighborApListPostRequestAttributesSlice []string

	QueryWithFilterQueryMeshNeighborApListPostRequestFiltersSlice []*QueryWithFilterQueryMeshNeighborApListPostRequestFilters

	QueryWithFilterQueryMeshNeighborApListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryMeshNeighborApListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryMeshNeighborApListPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryMeshNeighborApListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                    `json:"type,omitempty"`
		Value  *string                                                                    `json:"value,omitempty"`
	}

	QueryWithFilterQueryMeshNeighborApListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryMeshNeighborApListPostRequest struct {
		Attributes     QueryWithFilterQueryMeshNeighborApListPostRequestAttributesSlice `json:"attributes,omitempty"`
		Filters        QueryWithFilterQueryMeshNeighborApListPostRequestFiltersSlice    `json:"filters,omitempty"`
		FullTextSearch *QueryWithFilterQueryMeshNeighborApListPostRequestFullTextSearch `json:"fullTextSearch,omitempty"`
		Limit          *int                                                             `json:"limit,omitempty"`
		Options        interface{}                                                      `json:"options,omitempty"`
		Page           *int                                                             `json:"page,omitempty"`
		SortInfo       *QueryWithFilterQueryMeshNeighborApListPostRequestSortInfo       `json:"sortInfo,omitempty"`
		Start          *int                                                             `json:"start,omitempty"`
	}

	QueryWithFilterQueryMeshNeighborApListPost200ResponseListSlice []*QueryWithFilterQueryMeshNeighborApListPost200ResponseList

	QueryWithFilterQueryMeshNeighborApListPost200ResponseList struct {
		ApFirmware        *string  `json:"apFirmware,omitempty"`
		ApMac             *string  `json:"apMac,omitempty"`
		ApModel           *string  `json:"apModel,omitempty"`
		ApName            *string  `json:"apName,omitempty"`
		Channel           *string  `json:"channel,omitempty"`
		ConnectionStatus  *string  `json:"connectionStatus,omitempty"`
		ExternalIPAddress *string  `json:"externalIPAddress,omitempty"`
		IPAddress         *string  `json:"ipAddress,omitempty"`
		Rssi              *float64 `json:"rssi,omitempty"`
		ZoneName          *string  `json:"zoneName,omitempty"`
	}

	QueryWithFilterQueryMeshNeighborApListPost200Response struct {
		Extra             *json.RawMessage                                               `json:"extra,omitempty"`
		FirstIndex        *int                                                           `json:"firstIndex,omitempty"`
		HasMore           *bool                                                          `json:"hasMore,omitempty"`
		List              QueryWithFilterQueryMeshNeighborApListPost200ResponseListSlice `json:"list,omitempty"`
		RawDataTotalCount *int                                                           `json:"rawDataTotalCount,omitempty"`
		TotalCount        *int                                                           `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryMeshNeighborApListPost: Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *QueryWithFilterQueryMeshNeighborApListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryMeshNeighborApListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryMeshNeighborApListPost(ctx context.Context, apMac string, requestBody *QueryWithFilterQueryMeshNeighborApListPostRequest) (*http.Response, *QueryWithFilterQueryMeshNeighborApListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/query/mesh/{apMac}/neighbor", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	out := &QueryWithFilterQueryMeshNeighborApListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryRogueApListPostRequestAttributesSlice []string

	QueryWithFilterQueryRogueApListPostRequestExtraFiltersSlice []*QueryWithFilterQueryRogueApListPostRequestExtraFilters

	QueryWithFilterQueryRogueApListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryRogueApListPostRequestExtraNotFilters

	QueryWithFilterQueryRogueApListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestFiltersSlice []*QueryWithFilterQueryRogueApListPostRequestFilters

	QueryWithFilterQueryRogueApListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryRogueApListPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryRogueApListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                             `json:"type,omitempty"`
		Value  *string                                                             `json:"value,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequestOptions struct {
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
		GuestPassExpiration           *QueryWithFilterQueryRogueApListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                               `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                 `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                 `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                 `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryRogueApListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	QueryWithFilterQueryRogueApListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryRogueApListPostRequest struct {
		Attributes      QueryWithFilterQueryRogueApListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                        `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                          `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryRogueApListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryRogueApListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryRogueApListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryRogueApListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryRogueApListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                           `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryRogueApListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                           `json:"page,omitempty"`
		Query           *string                                                        `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryRogueApListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                           `json:"start,omitempty"`
	}

	QueryWithFilterQueryRogueApListPost200ResponseListSlice []*QueryWithFilterQueryRogueApListPost200ResponseList

	QueryWithFilterQueryRogueApListPost200ResponseListDetectedByAPSlice []*QueryWithFilterQueryRogueApListPost200ResponseListDetectedByAP

	QueryWithFilterQueryRogueApListPost200ResponseListDetectedByAP struct {
		ApMac        *string  `json:"apMac,omitempty"`
		ApName       *string  `json:"apName,omitempty"`
		LastDetected *float64 `json:"lastDetected,omitempty"`
		Rssi         *string  `json:"rssi,omitempty"`
		ZoneName     *string  `json:"zoneName,omitempty"`
	}

	QueryWithFilterQueryRogueApListPost200ResponseList struct {
		Channel      *float64                                                            `json:"channel,omitempty"`
		DetectedByAP QueryWithFilterQueryRogueApListPost200ResponseListDetectedByAPSlice `json:"detectedByAP,omitempty"`
		Encryption   *string                                                             `json:"encryption,omitempty"`
		LastDetected *float64                                                            `json:"lastDetected,omitempty"`
		Radio        *string                                                             `json:"radio,omitempty"`
		RogueAPMac   *string                                                             `json:"rogueAPMac,omitempty"`
		RogueMac     *string                                                             `json:"rogueMac,omitempty"`
		Ssid         *string                                                             `json:"ssid,omitempty"`
		Type         *string                                                             `json:"type,omitempty"`
	}

	QueryWithFilterQueryRogueApListPost200Response struct {
		Extra             *json.RawMessage                                        `json:"extra,omitempty"`
		FirstIndex        *int                                                    `json:"firstIndex,omitempty"`
		HasMore           *bool                                                   `json:"hasMore,omitempty"`
		List              QueryWithFilterQueryRogueApListPost200ResponseListSlice `json:"list,omitempty"`
		RawDataTotalCount *int                                                    `json:"rawDataTotalCount,omitempty"`
		TotalCount        *int                                                    `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryRogueApListPost: Use this API command to retrieve a list of rogue access points.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryRogueApListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryRogueApListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryRogueApListPost(ctx context.Context, requestBody *QueryWithFilterQueryRogueApListPostRequest) (*http.Response, *QueryWithFilterQueryRogueApListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/roguesInfoList", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryRogueApListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryAaaserversPostRequestAttributesSlice []string

	QueryWithFilterQueryAaaserversPostRequestExtraFiltersSlice []*QueryWithFilterQueryAaaserversPostRequestExtraFilters

	QueryWithFilterQueryAaaserversPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryAaaserversPostRequestExtraNotFilters

	QueryWithFilterQueryAaaserversPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestFiltersSlice []*QueryWithFilterQueryAaaserversPostRequestFilters

	QueryWithFilterQueryAaaserversPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryAaaserversPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryAaaserversPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                            `json:"type,omitempty"`
		Value  *string                                                            `json:"value,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterQueryAaaserversPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryAaaserversPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                              `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryAaaserversPostRequest struct {
		Attributes      QueryWithFilterQueryAaaserversPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryAaaserversPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryAaaserversPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryAaaserversPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryAaaserversPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryAaaserversPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                          `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryAaaserversPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                          `json:"page,omitempty"`
		Query           *string                                                       `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryAaaserversPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                          `json:"start,omitempty"`
	}

	QueryWithFilterQueryAaaserversPost200ResponseListSlice []*QueryWithFilterQueryAaaserversPost200ResponseList

	QueryWithFilterQueryAaaserversPost200ResponseList struct {
		AdminDomainName       *string  `json:"adminDomainName,omitempty"`
		AuthType              *string  `json:"authType,omitempty"`
		CreateOn              *float64 `json:"createOn,omitempty"`
		CreatorUUID           *string  `json:"creatorUUID,omitempty"`
		Description           *string  `json:"description,omitempty"`
		DomainID              *string  `json:"domainId,omitempty"`
		DomainName            *string  `json:"domainName,omitempty"`
		EnableSecondaryRadius *float64 `json:"enableSecondaryRadius,omitempty"`
		GlobalCatalog         *bool    `json:"globalCatalog,omitempty"`
		ID                    *string  `json:"id,omitempty"`
		IP                    *string  `json:"ip,omitempty"`
		Ipv6                  *string  `json:"ipv6,omitempty"`
		IsConflict            *float64 `json:"isConflict,omitempty"`
		Key                   *string  `json:"key,omitempty"`
		ModifiedDateTime      *int     `json:"modifiedDateTime,omitempty"`
		ModifierUsername      *string  `json:"modifierUsername,omitempty"`
		Name                  *string  `json:"name,omitempty"`
		Port                  *int     `json:"port,omitempty"`
		RadiusIP              *string  `json:"radiusIP,omitempty"`
		RadiusIPv6            *string  `json:"radiusIPv6,omitempty"`
		RadiusPort            *float64 `json:"radiusPort,omitempty"`
		RadiusRealm           *string  `json:"radiusRealm,omitempty"`
		SecondaryRadiusIP     *string  `json:"secondaryRadiusIP,omitempty"`
		SecondaryRadiusIPv6   *string  `json:"secondaryRadiusIPv6,omitempty"`
		SecondaryRadiusPort   *float64 `json:"secondaryRadiusPort,omitempty"`
		TacacsService         *string  `json:"tacacsService,omitempty"`
		TenantUUID            *string  `json:"tenantUUID,omitempty"`
		Type                  *string  `json:"type,omitempty"`
		WindowsDomainName     *string  `json:"windowsDomainName,omitempty"`
		ZoneUUID              *string  `json:"zoneUUID,omitempty"`
	}

	QueryWithFilterQueryAaaserversPost200Response struct {
		Extra      *json.RawMessage                                       `json:"extra,omitempty"`
		FirstIndex *int                                                   `json:"firstIndex,omitempty"`
		HasMore    *bool                                                  `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryAaaserversPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                   `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryAaaserversPost: Query AAAServers with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryAaaserversPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryAaaserversPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryAaaserversPost(ctx context.Context, requestBody *QueryWithFilterQueryAaaserversPostRequest) (*http.Response, *QueryWithFilterQueryAaaserversPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/aaaServer", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryAaaserversPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryAccountingAaaserversPostRequestAttributesSlice []string

	QueryWithFilterQueryAccountingAaaserversPostRequestExtraFiltersSlice []*QueryWithFilterQueryAccountingAaaserversPostRequestExtraFilters

	QueryWithFilterQueryAccountingAaaserversPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryAccountingAaaserversPostRequestExtraNotFilters

	QueryWithFilterQueryAccountingAaaserversPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestFiltersSlice []*QueryWithFilterQueryAccountingAaaserversPostRequestFilters

	QueryWithFilterQueryAccountingAaaserversPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryAccountingAaaserversPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryAccountingAaaserversPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                      `json:"type,omitempty"`
		Value  *string                                                                      `json:"value,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                          `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                        `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                          `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                          `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                        `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                          `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                          `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                          `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                          `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                          `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                          `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                        `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                          `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                        `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                        `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                        `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                        `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterQueryAccountingAaaserversPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                        `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                          `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                          `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                          `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryAccountingAaaserversPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                        `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                        `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                        `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                        `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                        `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                        `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                        `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                        `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                        `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPostRequest struct {
		Attributes      QueryWithFilterQueryAccountingAaaserversPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                 `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                   `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryAccountingAaaserversPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryAccountingAaaserversPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryAccountingAaaserversPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryAccountingAaaserversPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryAccountingAaaserversPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                    `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryAccountingAaaserversPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                    `json:"page,omitempty"`
		Query           *string                                                                 `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryAccountingAaaserversPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                    `json:"start,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPost200ResponseListSlice []*QueryWithFilterQueryAccountingAaaserversPost200ResponseList

	QueryWithFilterQueryAccountingAaaserversPost200ResponseList struct {
		AdminDomainName       *string  `json:"adminDomainName,omitempty"`
		AuthType              *string  `json:"authType,omitempty"`
		CreateOn              *float64 `json:"createOn,omitempty"`
		CreatorUUID           *string  `json:"creatorUUID,omitempty"`
		Description           *string  `json:"description,omitempty"`
		DomainID              *string  `json:"domainId,omitempty"`
		DomainName            *string  `json:"domainName,omitempty"`
		EnableSecondaryRadius *float64 `json:"enableSecondaryRadius,omitempty"`
		GlobalCatalog         *bool    `json:"globalCatalog,omitempty"`
		ID                    *string  `json:"id,omitempty"`
		IP                    *string  `json:"ip,omitempty"`
		Ipv6                  *string  `json:"ipv6,omitempty"`
		IsConflict            *float64 `json:"isConflict,omitempty"`
		Key                   *string  `json:"key,omitempty"`
		ModifiedDateTime      *int     `json:"modifiedDateTime,omitempty"`
		ModifierUsername      *string  `json:"modifierUsername,omitempty"`
		Name                  *string  `json:"name,omitempty"`
		Port                  *int     `json:"port,omitempty"`
		RadiusIP              *string  `json:"radiusIP,omitempty"`
		RadiusIPv6            *string  `json:"radiusIPv6,omitempty"`
		RadiusPort            *float64 `json:"radiusPort,omitempty"`
		RadiusRealm           *string  `json:"radiusRealm,omitempty"`
		SecondaryRadiusIP     *string  `json:"secondaryRadiusIP,omitempty"`
		SecondaryRadiusIPv6   *string  `json:"secondaryRadiusIPv6,omitempty"`
		SecondaryRadiusPort   *float64 `json:"secondaryRadiusPort,omitempty"`
		TacacsService         *string  `json:"tacacsService,omitempty"`
		TenantUUID            *string  `json:"tenantUUID,omitempty"`
		Type                  *string  `json:"type,omitempty"`
		WindowsDomainName     *string  `json:"windowsDomainName,omitempty"`
		ZoneUUID              *string  `json:"zoneUUID,omitempty"`
	}

	QueryWithFilterQueryAccountingAaaserversPost200Response struct {
		Extra      *json.RawMessage                                                 `json:"extra,omitempty"`
		FirstIndex *int                                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                                            `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryAccountingAaaserversPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                             `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryAccountingAaaserversPost: Query Accounting AAAServers with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryAccountingAaaserversPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryAccountingAaaserversPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryAccountingAaaserversPost(ctx context.Context, requestBody *QueryWithFilterQueryAccountingAaaserversPostRequest) (*http.Response, *QueryWithFilterQueryAccountingAaaserversPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/aaaServer/acct", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryAccountingAaaserversPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryAuthenticationAaaserversPostRequestAttributesSlice []string

	QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraFiltersSlice []*QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraFilters

	QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraNotFilters

	QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestFiltersSlice []*QueryWithFilterQueryAuthenticationAaaserversPostRequestFilters

	QueryWithFilterQueryAuthenticationAaaserversPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryAuthenticationAaaserversPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryAuthenticationAaaserversPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                          `json:"type,omitempty"`
		Value  *string                                                                          `json:"value,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                              `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                            `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                              `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                              `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                            `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                              `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                              `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                              `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                              `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                              `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                              `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                            `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                              `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                            `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                            `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                            `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                            `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterQueryAuthenticationAaaserversPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                            `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                              `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                              `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                              `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryAuthenticationAaaserversPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                            `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                            `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                            `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                            `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                            `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                            `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                            `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                            `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                            `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPostRequest struct {
		Attributes      QueryWithFilterQueryAuthenticationAaaserversPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                     `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                       `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryAuthenticationAaaserversPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryAuthenticationAaaserversPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryAuthenticationAaaserversPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                        `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryAuthenticationAaaserversPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                        `json:"page,omitempty"`
		Query           *string                                                                     `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryAuthenticationAaaserversPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                        `json:"start,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPost200ResponseListSlice []*QueryWithFilterQueryAuthenticationAaaserversPost200ResponseList

	QueryWithFilterQueryAuthenticationAaaserversPost200ResponseList struct {
		AdminDomainName       *string  `json:"adminDomainName,omitempty"`
		AuthType              *string  `json:"authType,omitempty"`
		CreateOn              *float64 `json:"createOn,omitempty"`
		CreatorUUID           *string  `json:"creatorUUID,omitempty"`
		Description           *string  `json:"description,omitempty"`
		DomainID              *string  `json:"domainId,omitempty"`
		DomainName            *string  `json:"domainName,omitempty"`
		EnableSecondaryRadius *float64 `json:"enableSecondaryRadius,omitempty"`
		GlobalCatalog         *bool    `json:"globalCatalog,omitempty"`
		ID                    *string  `json:"id,omitempty"`
		IP                    *string  `json:"ip,omitempty"`
		Ipv6                  *string  `json:"ipv6,omitempty"`
		IsConflict            *float64 `json:"isConflict,omitempty"`
		Key                   *string  `json:"key,omitempty"`
		ModifiedDateTime      *int     `json:"modifiedDateTime,omitempty"`
		ModifierUsername      *string  `json:"modifierUsername,omitempty"`
		Name                  *string  `json:"name,omitempty"`
		Port                  *int     `json:"port,omitempty"`
		RadiusIP              *string  `json:"radiusIP,omitempty"`
		RadiusIPv6            *string  `json:"radiusIPv6,omitempty"`
		RadiusPort            *float64 `json:"radiusPort,omitempty"`
		RadiusRealm           *string  `json:"radiusRealm,omitempty"`
		SecondaryRadiusIP     *string  `json:"secondaryRadiusIP,omitempty"`
		SecondaryRadiusIPv6   *string  `json:"secondaryRadiusIPv6,omitempty"`
		SecondaryRadiusPort   *float64 `json:"secondaryRadiusPort,omitempty"`
		TacacsService         *string  `json:"tacacsService,omitempty"`
		TenantUUID            *string  `json:"tenantUUID,omitempty"`
		Type                  *string  `json:"type,omitempty"`
		WindowsDomainName     *string  `json:"windowsDomainName,omitempty"`
		ZoneUUID              *string  `json:"zoneUUID,omitempty"`
	}

	QueryWithFilterQueryAuthenticationAaaserversPost200Response struct {
		Extra      *json.RawMessage                                                     `json:"extra,omitempty"`
		FirstIndex *int                                                                 `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryAuthenticationAaaserversPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                 `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryAuthenticationAaaserversPost: Query Authentication AAAServers with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryAuthenticationAaaserversPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryAuthenticationAaaserversPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryAuthenticationAaaserversPost(ctx context.Context, requestBody *QueryWithFilterQueryAuthenticationAaaserversPostRequest) (*http.Response, *QueryWithFilterQueryAuthenticationAaaserversPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/aaaServer/auth", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryAuthenticationAaaserversPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFilters

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                    `json:"type,omitempty"`
		Value  *string                                                                                    `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                        `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                      `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                        `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                        `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                      `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                        `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                        `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                        `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                        `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                        `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                        `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                      `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                        `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                      `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                      `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                      `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                      `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                      `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                        `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                        `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                        `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                      `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                      `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                      `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                      `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                      `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                      `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                      `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                      `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                      `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                               `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                 `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                  `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                  `json:"page,omitempty"`
		Query           *string                                                                               `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                  `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPost: Query AP USBSoftware Package Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveApUsbsoftwarePackageProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/apUsbSoftwarePackage", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	BonjourFencingPolicyRetrieveListPostRequestAttributesSlice []string

	BonjourFencingPolicyRetrieveListPostRequestExtraFiltersSlice []*BonjourFencingPolicyRetrieveListPostRequestExtraFilters

	BonjourFencingPolicyRetrieveListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestExtraNotFiltersSlice []*BonjourFencingPolicyRetrieveListPostRequestExtraNotFilters

	BonjourFencingPolicyRetrieveListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestFiltersSlice []*BonjourFencingPolicyRetrieveListPostRequestFilters

	BonjourFencingPolicyRetrieveListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestFullTextSearchFieldsSlice []string

	BonjourFencingPolicyRetrieveListPostRequestFullTextSearch struct {
		Fields BonjourFencingPolicyRetrieveListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                              `json:"type,omitempty"`
		Value  *string                                                              `json:"value,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                  `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                  `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                  `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                  `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                  `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                  `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                  `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                  `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                  `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                  `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *BonjourFencingPolicyRetrieveListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                  `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                  `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                  `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *BonjourFencingPolicyRetrieveListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                `json:"localUser_userSource,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPostRequest struct {
		Attributes      BonjourFencingPolicyRetrieveListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                         `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                           `json:"expandDomains,omitempty"`
		ExtraFilters    BonjourFencingPolicyRetrieveListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters BonjourFencingPolicyRetrieveListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *BonjourFencingPolicyRetrieveListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         BonjourFencingPolicyRetrieveListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *BonjourFencingPolicyRetrieveListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                        `json:"limit,omitempty"`
		Options         *BonjourFencingPolicyRetrieveListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                            `json:"page,omitempty"`
		Query           *string                                                         `json:"query,omitempty"`
		SortInfo        *BonjourFencingPolicyRetrieveListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                        `json:"start,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPost200ResponseListSlice []*BonjourFencingPolicyRetrieveListPost200ResponseList

	BonjourFencingPolicyRetrieveListPost200ResponseListBonjourFencingRuleListSlice []*BonjourFencingPolicyRetrieveListPost200ResponseListBonjourFencingRuleList

	BonjourFencingPolicyRetrieveListPost200ResponseListBonjourFencingRuleList struct {
		ClosestAp    *string `json:"closestAp,omitempty"`
		Description  *string `json:"description,omitempty"`
		DeviceMac    *string `json:"deviceMac,omitempty"`
		DeviceType   *string `json:"deviceType,omitempty"`
		FencingRange *string `json:"fencingRange,omitempty"`
		ServiceType  *string `json:"serviceType,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPost200ResponseList struct {
		BonjourFencingRuleList BonjourFencingPolicyRetrieveListPost200ResponseListBonjourFencingRuleListSlice `json:"bonjourFencingRuleList,omitempty"`
		CreateDateTime         *int                                                                           `json:"createDateTime,omitempty"`
		CreatorID              *string                                                                        `json:"creatorId,omitempty"`
		CreatorUsername        *string                                                                        `json:"creatorUsername,omitempty"`
		Description            *string                                                                        `json:"description,omitempty"`
		ID                     *string                                                                        `json:"id,omitempty"`
		ModifiedDateTime       *int                                                                           `json:"modifiedDateTime,omitempty"`
		ModifierID             *string                                                                        `json:"modifierId,omitempty"`
		ModifierUsername       *string                                                                        `json:"modifierUsername,omitempty"`
		Name                   *string                                                                        `json:"name,omitempty"`
		ZoneID                 *string                                                                        `json:"zoneId,omitempty"`
	}

	BonjourFencingPolicyRetrieveListPost200Response struct {
		FirstIndex *int                                                     `json:"firstIndex,omitempty"`
		HasMore    *bool                                                    `json:"hasMore,omitempty"`
		List       BonjourFencingPolicyRetrieveListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                     `json:"totalCount,omitempty"`
	}
)

// BonjourFencingPolicyRetrieveListPost: Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *BonjourFencingPolicyRetrieveListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourFencingPolicyRetrieveListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) BonjourFencingPolicyRetrieveListPost(ctx context.Context, requestBody *BonjourFencingPolicyRetrieveListPostRequest) (*http.Response, *BonjourFencingPolicyRetrieveListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/bonjourFencingPolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &BonjourFencingPolicyRetrieveListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFilters

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                             `json:"type,omitempty"`
		Value  *string                                                                             `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                 `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                               `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                 `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                 `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                               `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                 `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                 `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                 `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                 `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                 `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                 `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                               `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                 `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                               `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                               `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                               `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                               `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                               `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                 `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                 `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                 `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                               `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                               `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                               `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                               `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                               `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                               `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                               `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                               `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                               `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveBonjourpolicyProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                        `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                          `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                           `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                           `json:"page,omitempty"`
		Query           *string                                                                        `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                           `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveBonjourpolicyProfileListPost: Query bonjourPolicy Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveBonjourpolicyProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveBonjourpolicyProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/bonjourPolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	ClientIsolationWhitelistRetrieveListPostRequestAttributesSlice []string

	ClientIsolationWhitelistRetrieveListPostRequestExtraFiltersSlice []*ClientIsolationWhitelistRetrieveListPostRequestExtraFilters

	ClientIsolationWhitelistRetrieveListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestExtraNotFiltersSlice []*ClientIsolationWhitelistRetrieveListPostRequestExtraNotFilters

	ClientIsolationWhitelistRetrieveListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestFiltersSlice []*ClientIsolationWhitelistRetrieveListPostRequestFilters

	ClientIsolationWhitelistRetrieveListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestFullTextSearchFieldsSlice []string

	ClientIsolationWhitelistRetrieveListPostRequestFullTextSearch struct {
		Fields ClientIsolationWhitelistRetrieveListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                  `json:"type,omitempty"`
		Value  *string                                                                  `json:"value,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                      `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                    `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                      `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                      `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                    `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                      `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                      `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                      `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                      `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                      `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                      `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                    `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                      `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                    `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                    `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                    `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                    `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *ClientIsolationWhitelistRetrieveListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                    `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                      `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                      `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                      `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *ClientIsolationWhitelistRetrieveListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                    `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                    `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                    `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                    `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                    `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                    `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                    `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                    `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                    `json:"localUser_userSource,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPostRequest struct {
		Attributes      ClientIsolationWhitelistRetrieveListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                             `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                               `json:"expandDomains,omitempty"`
		ExtraFilters    ClientIsolationWhitelistRetrieveListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters ClientIsolationWhitelistRetrieveListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *ClientIsolationWhitelistRetrieveListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         ClientIsolationWhitelistRetrieveListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *ClientIsolationWhitelistRetrieveListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                            `json:"limit,omitempty"`
		Options         *ClientIsolationWhitelistRetrieveListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                `json:"page,omitempty"`
		Query           *string                                                             `json:"query,omitempty"`
		SortInfo        *ClientIsolationWhitelistRetrieveListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                            `json:"start,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPost200ResponseListSlice []*ClientIsolationWhitelistRetrieveListPost200ResponseList

	ClientIsolationWhitelistRetrieveListPost200ResponseListWhitelistSlice []*ClientIsolationWhitelistRetrieveListPost200ResponseListWhitelist

	ClientIsolationWhitelistRetrieveListPost200ResponseListWhitelist struct {
		Description *string `json:"description,omitempty"`
		IP          *string `json:"ip,omitempty"`
		Mac         *string `json:"mac,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPost200ResponseList struct {
		ClientIsolationAutoEnabled *bool                                                                 `json:"clientIsolationAutoEnabled,omitempty"`
		CreateDateTime             *int                                                                  `json:"createDateTime,omitempty"`
		CreatorID                  *string                                                               `json:"creatorId,omitempty"`
		CreatorUsername            *string                                                               `json:"creatorUsername,omitempty"`
		Description                *string                                                               `json:"description,omitempty"`
		ID                         *string                                                               `json:"id,omitempty"`
		ModifiedDateTime           *int                                                                  `json:"modifiedDateTime,omitempty"`
		ModifierID                 *string                                                               `json:"modifierId,omitempty"`
		ModifierUsername           *string                                                               `json:"modifierUsername,omitempty"`
		Name                       *string                                                               `json:"name,omitempty"`
		Whitelist                  ClientIsolationWhitelistRetrieveListPost200ResponseListWhitelistSlice `json:"whitelist,omitempty"`
		ZoneID                     *string                                                               `json:"zoneId,omitempty"`
	}

	ClientIsolationWhitelistRetrieveListPost200Response struct {
		CreateDateTime   *int                                                         `json:"createDateTime,omitempty"`
		CreatorID        *string                                                      `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                      `json:"creatorUsername,omitempty"`
		FirstIndex       *int                                                         `json:"firstIndex,omitempty"`
		HasMore          *bool                                                        `json:"hasMore,omitempty"`
		List             ClientIsolationWhitelistRetrieveListPost200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                                         `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                      `json:"modifierId,omitempty"`
		ModifierUsername *string                                                      `json:"modifierUsername,omitempty"`
		TotalCount       *int                                                         `json:"totalCount,omitempty"`
	}
)

// ClientIsolationWhitelistRetrieveListPost: Retrieve a list of Client Isolation Whitelist
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ClientIsolationWhitelistRetrieveListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ClientIsolationWhitelistRetrieveListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) ClientIsolationWhitelistRetrieveListPost(ctx context.Context, requestBody *ClientIsolationWhitelistRetrieveListPostRequest) (*http.Response, *ClientIsolationWhitelistRetrieveListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/clientIsolationWhitelist", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &ClientIsolationWhitelistRetrieveListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFilters

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                            `json:"type,omitempty"`
		Value  *string                                                                            `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveDevicePolicyProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveDevicePolicyProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                              `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveDevicePolicyProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveDevicePolicyProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveDevicePolicyProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveDevicePolicyProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                          `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveDevicePolicyProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                          `json:"page,omitempty"`
		Query           *string                                                                       `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveDevicePolicyProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                          `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveDevicePolicyProfileListPost: Query Device Policy Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveDevicePolicyProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveDevicePolicyProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveDevicePolicyProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/devicePolicy", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveDhcpProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveDhcpProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveDhcpProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveDhcpProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveDhcpProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveDhcpProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveDhcpProfileListPostRequestFilters

	QueryWithFilterRetrieveDhcpProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveDhcpProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveDhcpProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                    `json:"type,omitempty"`
		Value  *string                                                                    `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                        `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                      `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                        `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                        `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                      `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                        `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                        `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                        `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                        `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                        `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                        `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                      `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                        `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                      `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                      `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                      `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                      `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveDhcpProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                      `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                        `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                        `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                        `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveDhcpProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                      `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                      `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                      `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                      `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                      `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                      `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                      `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                      `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                      `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveDhcpProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                               `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                 `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveDhcpProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveDhcpProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveDhcpProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveDhcpProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveDhcpProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                  `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveDhcpProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                  `json:"page,omitempty"`
		Query           *string                                                               `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveDhcpProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                  `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDhcpProfileListPost200ResponseListSlice []*QueryWithFilterRetrieveDhcpProfileListPost200ResponseList

	QueryWithFilterRetrieveDhcpProfileListPost200ResponseList struct {
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

	QueryWithFilterRetrieveDhcpProfileListPost200Response struct {
		CreateDateTime   *int                                                           `json:"createDateTime,omitempty"`
		CreatorID        *string                                                        `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                        `json:"creatorUsername,omitempty"`
		FirstIndex       *int                                                           `json:"firstIndex,omitempty"`
		HasMore          *bool                                                          `json:"hasMore,omitempty"`
		List             QueryWithFilterRetrieveDhcpProfileListPost200ResponseListSlice `json:"list,omitempty"`
		ModifiedDateTime *int                                                           `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                        `json:"modifierId,omitempty"`
		ModifierUsername *string                                                        `json:"modifierUsername,omitempty"`
		TotalCount       *int                                                           `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterRetrieveDhcpProfileListPost: Query DHCP Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveDhcpProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterRetrieveDhcpProfileListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveDhcpProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveDhcpProfileListPostRequest) (*http.Response, *QueryWithFilterRetrieveDhcpProfileListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/dhcpProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterRetrieveDhcpProfileListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterRetrieveDscpProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveDscpProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveDscpProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveDscpProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveDscpProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveDscpProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveDscpProfileListPostRequestFilters

	QueryWithFilterRetrieveDscpProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveDscpProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveDscpProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                    `json:"type,omitempty"`
		Value  *string                                                                    `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                        `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                      `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                        `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                        `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                      `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                        `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                        `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                        `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                        `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                        `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                        `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                      `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                        `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                      `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                      `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                      `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                      `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveDscpProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                      `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                        `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                        `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                        `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveDscpProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                      `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                      `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                      `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                      `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                      `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                      `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                      `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                      `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                      `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveDscpProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveDscpProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                               `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                 `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveDscpProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveDscpProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveDscpProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveDscpProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveDscpProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                  `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveDscpProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                  `json:"page,omitempty"`
		Query           *string                                                               `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveDscpProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                  `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveDscpProfileListPost: Query DSCP Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveDscpProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveDscpProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveDscpProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/dscpProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveEthernetPortProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveEthernetPortProfileListPostRequestFilters

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveEthernetPortProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                            `json:"type,omitempty"`
		Value  *string                                                                            `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                              `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                              `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                              `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                              `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                              `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                              `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                              `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveEthernetPortProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveEthernetPortProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                              `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                              `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                              `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                              `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                              `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                              `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                              `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                              `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                              `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveEthernetPortProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveEthernetPortProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveEthernetPortProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveEthernetPortProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveEthernetPortProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                          `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveEthernetPortProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                          `json:"page,omitempty"`
		Query           *string                                                                       `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveEthernetPortProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                          `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveEthernetPortProfileListPost: Query Ethernet Port Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveEthernetPortProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveEthernetPortProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveEthernetPortProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/ethernetPortProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveGuessAccessProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveGuessAccessProfileListPostRequestFilters

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveGuessAccessProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                           `json:"type,omitempty"`
		Value  *string                                                                           `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                               `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                             `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                               `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                               `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                             `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                               `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                               `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                               `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                               `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                               `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                               `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                             `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                               `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                             `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                             `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                             `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                             `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveGuessAccessProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                             `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                               `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                               `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                               `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveGuessAccessProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                             `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                             `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                             `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                             `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                             `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                             `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                             `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                             `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                             `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveGuessAccessProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveGuessAccessProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                      `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                        `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveGuessAccessProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveGuessAccessProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveGuessAccessProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                         `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveGuessAccessProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                         `json:"page,omitempty"`
		Query           *string                                                                      `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveGuessAccessProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                         `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveGuessAccessProfileListPost: Query Guess Access Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveGuessAccessProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveGuessAccessProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveGuessAccessProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/guessAccess", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveHotspotProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveHotspotProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveHotspotProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveHotspotProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveHotspotProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveHotspotProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveHotspotProfileListPostRequestFilters

	QueryWithFilterRetrieveHotspotProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveHotspotProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveHotspotProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                       `json:"type,omitempty"`
		Value  *string                                                                       `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                           `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                         `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                           `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                           `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                         `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                           `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                           `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                           `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                           `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                           `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                           `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                         `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                           `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                         `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                         `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                         `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                         `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveHotspotProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                         `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                           `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                           `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                           `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveHotspotProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                         `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                         `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                         `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                         `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                         `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                         `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                         `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                         `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                         `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveHotspotProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveHotspotProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                  `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                    `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveHotspotProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveHotspotProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveHotspotProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveHotspotProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveHotspotProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                     `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveHotspotProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                     `json:"page,omitempty"`
		Query           *string                                                                  `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveHotspotProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                     `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveHotspotProfileListPost: Query Hotspot Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveHotspotProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveHotspotProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveHotspotProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/hotspot", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveHotspot20ProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveHotspot20ProfileListPostRequestFilters

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveHotspot20ProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                         `json:"type,omitempty"`
		Value  *string                                                                         `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                             `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                           `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                             `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                             `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                           `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                             `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                             `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                             `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                             `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                             `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                             `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                           `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                             `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                           `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                           `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                           `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                           `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveHotspot20ProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                           `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                             `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                             `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                             `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveHotspot20ProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                           `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                           `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                           `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                           `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                           `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                           `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                           `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                           `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                           `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveHotspot20ProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveHotspot20ProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                    `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                      `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveHotspot20ProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveHotspot20ProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveHotspot20ProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                       `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveHotspot20ProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                       `json:"page,omitempty"`
		Query           *string                                                                    `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveHotspot20ProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                       `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveHotspot20ProfileListPost: Query Hotspot20 Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveHotspot20ProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveHotspot20ProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveHotspot20ProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/hotspot20Profile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFilters

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                               `json:"type,omitempty"`
		Value  *string                                                                               `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                   `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                 `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                   `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                   `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                 `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                   `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                   `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                   `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                   `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                   `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                   `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                 `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                   `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                 `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                 `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                 `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                 `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                 `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                   `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                   `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                   `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                 `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                 `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                 `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                 `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                 `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                 `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                 `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                 `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                 `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                          `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                            `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                             `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                             `json:"page,omitempty"`
		Query           *string                                                                          `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                             `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveL2AccesscontrolProfileListPost: Query L2 AccessControl Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveL2AccesscontrolProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveL2AccesscontrolProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/l2AccessControl", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveVenueProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveVenueProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveVenueProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveVenueProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveVenueProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveVenueProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveVenueProfileListPostRequestFilters

	QueryWithFilterRetrieveVenueProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveVenueProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveVenueProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                     `json:"type,omitempty"`
		Value  *string                                                                     `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequestOptions struct {
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
		GuestPassExpiration           *QueryWithFilterRetrieveVenueProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveVenueProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	QueryWithFilterRetrieveVenueProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveVenueProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveVenueProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveVenueProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveVenueProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveVenueProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveVenueProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveVenueProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                   `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveVenueProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                   `json:"page,omitempty"`
		Query           *string                                                                `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveVenueProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                   `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveVenueProfileListPost: Query Venue Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveVenueProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveVenueProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveVenueProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/venueProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFilters

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                           `json:"type,omitempty"`
		Value  *string                                                                           `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                               `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                             `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                               `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                               `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                             `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                               `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                               `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                               `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                               `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                               `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                               `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                             `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                               `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                             `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                             `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                             `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                             `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveVlanPoolingProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                             `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                               `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                               `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                               `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveVlanPoolingProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                             `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                             `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                             `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                             `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                             `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                             `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                             `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                             `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                             `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveVlanPoolingProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveVlanPoolingProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                      `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                        `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveVlanPoolingProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveVlanPoolingProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                         `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveVlanPoolingProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                         `json:"page,omitempty"`
		Query           *string                                                                      `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveVlanPoolingProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                         `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveVlanPoolingProfileListPost: Query Vlan Pooling Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveVlanPoolingProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveVlanPoolingProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveVlanPoolingProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/vlanPooling", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFilters

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                 `json:"type,omitempty"`
		Value  *string                                                                                 `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                     `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                   `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                     `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                     `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                   `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                     `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                     `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                     `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                     `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                     `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                     `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                   `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                     `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                   `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                   `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                   `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                   `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                   `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                     `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                     `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                     `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                   `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                   `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                   `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                   `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                   `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                   `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                   `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                   `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                   `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveWebAuthenticationProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                            `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                              `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                               `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                               `json:"page,omitempty"`
		Query           *string                                                                            `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                               `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveWebAuthenticationProfileListPost: Query Web Authentications with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveWebAuthenticationProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveWebAuthenticationProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/webAuthentication", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveWechatProfileListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveWechatProfileListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveWechatProfileListPostRequestExtraFilters

	QueryWithFilterRetrieveWechatProfileListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveWechatProfileListPostRequestExtraNotFilters

	QueryWithFilterRetrieveWechatProfileListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestFiltersSlice []*QueryWithFilterRetrieveWechatProfileListPostRequestFilters

	QueryWithFilterRetrieveWechatProfileListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveWechatProfileListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveWechatProfileListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                      `json:"type,omitempty"`
		Value  *string                                                                      `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                          `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                        `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                          `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                          `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                        `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                          `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                          `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                          `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                          `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                          `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                          `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                        `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                          `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                        `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                        `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                        `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                        `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveWechatProfileListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                        `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                          `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                          `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                          `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveWechatProfileListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                        `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                        `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                        `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                        `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                        `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                        `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                        `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                        `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                        `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveWechatProfileListPostRequest struct {
		Attributes      QueryWithFilterRetrieveWechatProfileListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                 `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                   `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveWechatProfileListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveWechatProfileListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveWechatProfileListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveWechatProfileListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveWechatProfileListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                    `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveWechatProfileListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                    `json:"page,omitempty"`
		Query           *string                                                                 `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveWechatProfileListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                    `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveWechatProfileListPost: Query Wechat Profiles with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveWechatProfileListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveWechatProfileListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveWechatProfileListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/wechatProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveWlanSchedulerListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraFilters

	QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraNotFilters

	QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestFiltersSlice []*QueryWithFilterRetrieveWlanSchedulerListPostRequestFilters

	QueryWithFilterRetrieveWlanSchedulerListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveWlanSchedulerListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveWlanSchedulerListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                      `json:"type,omitempty"`
		Value  *string                                                                      `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                          `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                        `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                          `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                          `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                        `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                          `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                          `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                          `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                          `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                          `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                          `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                        `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                          `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                        `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                        `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                        `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                        `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterRetrieveWlanSchedulerListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                        `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                          `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                          `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                          `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveWlanSchedulerListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                        `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                        `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                        `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                        `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                        `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                        `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                        `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                        `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                        `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveWlanSchedulerListPostRequest struct {
		Attributes      QueryWithFilterRetrieveWlanSchedulerListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                 `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                   `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveWlanSchedulerListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveWlanSchedulerListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveWlanSchedulerListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                    `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveWlanSchedulerListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                    `json:"page,omitempty"`
		Query           *string                                                                 `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveWlanSchedulerListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                    `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveWlanSchedulerListPost: Query Wlan Schedulers with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveWlanSchedulerListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveWlanSchedulerListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveWlanSchedulerListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/services/wlanScheduler", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	QueryWithFilterRetrieveUserListPostRequestAttributesSlice []string

	QueryWithFilterRetrieveUserListPostRequestExtraFiltersSlice []*QueryWithFilterRetrieveUserListPostRequestExtraFilters

	QueryWithFilterRetrieveUserListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestExtraNotFiltersSlice []*QueryWithFilterRetrieveUserListPostRequestExtraNotFilters

	QueryWithFilterRetrieveUserListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestFiltersSlice []*QueryWithFilterRetrieveUserListPostRequestFilters

	QueryWithFilterRetrieveUserListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterRetrieveUserListPostRequestFullTextSearch struct {
		Fields QueryWithFilterRetrieveUserListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                             `json:"type,omitempty"`
		Value  *string                                                             `json:"value,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequestOptions struct {
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
		GuestPassExpiration           *QueryWithFilterRetrieveUserListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                               `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                 `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                 `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                 `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterRetrieveUserListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	QueryWithFilterRetrieveUserListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterRetrieveUserListPostRequest struct {
		Attributes      QueryWithFilterRetrieveUserListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                        `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                          `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterRetrieveUserListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterRetrieveUserListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterRetrieveUserListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterRetrieveUserListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterRetrieveUserListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                           `json:"limit,omitempty"`
		Options         *QueryWithFilterRetrieveUserListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                           `json:"page,omitempty"`
		Query           *string                                                        `json:"query,omitempty"`
		SortInfo        *QueryWithFilterRetrieveUserListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                           `json:"start,omitempty"`
	}
)

// QueryWithFilterRetrieveUserListPost: Query users with specified filters.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterRetrieveUserListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterRetrieveUserListPost(ctx context.Context, requestBody *QueryWithFilterRetrieveUserListPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/user", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return q.client.Ensure(ctx, request, 200, nil)
}

type (
	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestAttributesSlice []string

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraFiltersSlice []*ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraFilters

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraNotFiltersSlice []*ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraNotFilters

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFiltersSlice []*ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFilters

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFullTextSearchFieldsSlice []string

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFullTextSearch struct {
		Fields ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                 `json:"type,omitempty"`
		Value  *string                                                                                 `json:"value,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                     `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                   `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                     `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                     `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                   `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                     `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                     `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                     `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                     `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                     `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                     `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                   `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                     `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                   `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                   `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                   `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                   `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *ApplicationVisibilityControlUserDefinedRetrieveListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                   `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                     `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                     `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                     `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *ApplicationVisibilityControlUserDefinedRetrieveListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                   `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                   `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                   `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                   `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                   `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                   `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                   `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                   `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                   `json:"localUser_userSource,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPostRequest struct {
		Attributes      ApplicationVisibilityControlUserDefinedRetrieveListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                            `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                              `json:"expandDomains,omitempty"`
		ExtraFilters    ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *ApplicationVisibilityControlUserDefinedRetrieveListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *ApplicationVisibilityControlUserDefinedRetrieveListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                                           `json:"limit,omitempty"`
		Options         *ApplicationVisibilityControlUserDefinedRetrieveListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                               `json:"page,omitempty"`
		Query           *string                                                                            `json:"query,omitempty"`
		SortInfo        *ApplicationVisibilityControlUserDefinedRetrieveListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                                           `json:"start,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPost200ResponseListSlice []*ApplicationVisibilityControlUserDefinedRetrieveListPost200ResponseList

	ApplicationVisibilityControlUserDefinedRetrieveListPost200ResponseList struct {
		AppID            *float64 `json:"appId,omitempty"`
		CreateDateTime   *int     `json:"createDateTime,omitempty"`
		CreatorID        *string  `json:"creatorId,omitempty"`
		CreatorUsername  *string  `json:"creatorUsername,omitempty"`
		DestIP           *string  `json:"destIp,omitempty"`
		DestPort         *float64 `json:"destPort,omitempty"`
		DomainID         *string  `json:"domainId,omitempty"`
		ID               *string  `json:"id,omitempty"`
		ModifiedDateTime *int     `json:"modifiedDateTime,omitempty"`
		ModifierID       *string  `json:"modifierId,omitempty"`
		ModifierUsername *string  `json:"modifierUsername,omitempty"`
		Name             *string  `json:"name,omitempty"`
		Netmask          *string  `json:"netmask,omitempty"`
		Protocol         *string  `json:"protocol,omitempty"`
		TenantID         *string  `json:"tenantId,omitempty"`
		Type             *string  `json:"type,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedRetrieveListPost200Response struct {
		Extra      *json.RawMessage                                                            `json:"extra,omitempty"`
		FirstIndex *int                                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                       `json:"hasMore,omitempty"`
		List       ApplicationVisibilityControlUserDefinedRetrieveListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                        `json:"totalCount,omitempty"`
	}
)

// ApplicationVisibilityControlUserDefinedRetrieveListPost: Use this API command to retrieve a list of AVC User Defined profiles.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ApplicationVisibilityControlUserDefinedRetrieveListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlUserDefinedRetrieveListPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) ApplicationVisibilityControlUserDefinedRetrieveListPost(ctx context.Context, requestBody *ApplicationVisibilityControlUserDefinedRetrieveListPostRequest) (*http.Response, *ApplicationVisibilityControlUserDefinedRetrieveListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/userDefined", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &ApplicationVisibilityControlUserDefinedRetrieveListPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	QueryWithFilterQueryWlansPostRequestAttributesSlice []string

	QueryWithFilterQueryWlansPostRequestExtraFiltersSlice []*QueryWithFilterQueryWlansPostRequestExtraFilters

	QueryWithFilterQueryWlansPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestExtraNotFiltersSlice []*QueryWithFilterQueryWlansPostRequestExtraNotFilters

	QueryWithFilterQueryWlansPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestFiltersSlice []*QueryWithFilterQueryWlansPostRequestFilters

	QueryWithFilterQueryWlansPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestFullTextSearchFieldsSlice []string

	QueryWithFilterQueryWlansPostRequestFullTextSearch struct {
		Fields QueryWithFilterQueryWlansPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                       `json:"type,omitempty"`
		Value  *string                                                       `json:"value,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                           `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                         `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                           `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                           `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                         `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                           `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                           `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                           `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                           `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                           `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                           `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                         `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                           `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                         `json:"auth_type,omitempty"`
		ForwardingType                *string                                                         `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                         `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                         `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *QueryWithFilterQueryWlansPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                         `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                           `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                           `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                           `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *QueryWithFilterQueryWlansPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                         `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                         `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                         `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                         `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                         `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                         `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                         `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                         `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                         `json:"localUser_userSource,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	QueryWithFilterQueryWlansPostRequest struct {
		Attributes      QueryWithFilterQueryWlansPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                  `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                    `json:"expandDomains,omitempty"`
		ExtraFilters    QueryWithFilterQueryWlansPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters QueryWithFilterQueryWlansPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *QueryWithFilterQueryWlansPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         QueryWithFilterQueryWlansPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *QueryWithFilterQueryWlansPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                     `json:"limit,omitempty"`
		Options         *QueryWithFilterQueryWlansPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                     `json:"page,omitempty"`
		Query           *string                                                  `json:"query,omitempty"`
		SortInfo        *QueryWithFilterQueryWlansPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                     `json:"start,omitempty"`
	}

	QueryWithFilterQueryWlansPost200ResponseListSlice []*QueryWithFilterQueryWlansPost200ResponseList

	QueryWithFilterQueryWlansPost200ResponseList struct {
		Alerts                *float64 `json:"alerts,omitempty"`
		ApplicationVisibility *string  `json:"applicationVisibility,omitempty"`
		AuthMethod            *string  `json:"authMethod,omitempty"`
		Clients               *float64 `json:"clients,omitempty"`
		Description           *string  `json:"description,omitempty"`
		Enability1K           *string  `json:"enability11k,omitempty"`
		Enability1R           *string  `json:"enability11r,omitempty"`
		EncryptionMethod      *string  `json:"encryptionMethod,omitempty"`
		Name                  *string  `json:"name,omitempty"`
		Ssid                  *string  `json:"ssid,omitempty"`
		Status                *string  `json:"status,omitempty"`
		TenantDomainName      *string  `json:"tenantDomainName,omitempty"`
		TenantID              *string  `json:"tenantId,omitempty"`
		Traffic               *float64 `json:"traffic,omitempty"`
		TrafficDownlink       *float64 `json:"trafficDownlink,omitempty"`
		TrafficUplink         *float64 `json:"trafficUplink,omitempty"`
		Tunneled              *string  `json:"tunneled,omitempty"`
		Utp                   *string  `json:"utp,omitempty"`
		Vlan                  *float64 `json:"vlan,omitempty"`
		WepEncryptionStrength *float64 `json:"wepEncryptionStrength,omitempty"`
		WlanID                *string  `json:"wlanId,omitempty"`
		WpaVersion            *string  `json:"wpaVersion,omitempty"`
		ZeroITEnabled         *string  `json:"zeroITEnabled,omitempty"`
		ZeroITOnboard         *string  `json:"zeroITOnboard,omitempty"`
	}

	QueryWithFilterQueryWlansPost200Response struct {
		Extra      *json.RawMessage                                  `json:"extra,omitempty"`
		FirstIndex *int                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                             `json:"hasMore,omitempty"`
		List       QueryWithFilterQueryWlansPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                              `json:"totalCount,omitempty"`
	}
)

// QueryWithFilterQueryWlansPost: Query WLANs with specified filters
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *QueryWithFilterQueryWlansPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *QueryWithFilterQueryWlansPost200Response
// - error: Error seen or nil if none
func (q *QueryAPI) QueryWithFilterQueryWlansPost(ctx context.Context, requestBody *QueryWithFilterQueryWlansPostRequest) (*http.Response, *QueryWithFilterQueryWlansPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/query/wlan", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &QueryWithFilterQueryWlansPost200Response{}
	httpResponse, _, err := q.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}
