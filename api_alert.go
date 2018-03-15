package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-15T14:33:32-0500
// API Version: v5

type AlertAPI struct {
	client *Client
}
type (
	EventAndAlarmAcknowledgeAlarmsPutRequestIDListSlice []*string

	EventAndAlarmAcknowledgeAlarmsPutRequest struct {
		IDList EventAndAlarmAcknowledgeAlarmsPutRequestIDListSlice `json:"idList,omitempty"`
	}
)

// EventAndAlarmAcknowledgeAlarmsPut: Acknowledge multiple Alarms with provided alarmIDs.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *EventAndAlarmAcknowledgeAlarmsPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AlertAPI) EventAndAlarmAcknowledgeAlarmsPut(ctx *UserContext, requestBody *EventAndAlarmAcknowledgeAlarmsPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "PUT", "/v5_0/alert/alarm/ack")
	request.body = requestBody
	request.authenticated = true
	return a.client.doRequest(request, 204, nil)
}

type (
	EventAndAlarmClearAlarmsPutRequestIDListSlice []*string

	EventAndAlarmClearAlarmsPutRequest struct {
		Comment *string                                       `json:"comment,omitempty"`
		IDList  EventAndAlarmClearAlarmsPutRequestIDListSlice `json:"idList,omitempty"`
	}
)

// EventAndAlarmClearAlarmsPut: Clear multiple Alarms with provided alarmIDs.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *EventAndAlarmClearAlarmsPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AlertAPI) EventAndAlarmClearAlarmsPut(ctx *UserContext, requestBody *EventAndAlarmClearAlarmsPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "PUT", "/v5_0/alert/alarm/clear")
	request.body = requestBody
	request.authenticated = true
	return a.client.doRequest(request, 204, nil)
}

type (
	EventAndAlarmRetrieveAlarmListPostRequestAttributesSlice []*string

	EventAndAlarmRetrieveAlarmListPostRequestExtraFiltersSlice []*EventAndAlarmRetrieveAlarmListPostRequestExtraFilters

	EventAndAlarmRetrieveAlarmListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestExtraNotFiltersSlice []*EventAndAlarmRetrieveAlarmListPostRequestExtraNotFilters

	EventAndAlarmRetrieveAlarmListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestFiltersSlice []*EventAndAlarmRetrieveAlarmListPostRequestFilters

	EventAndAlarmRetrieveAlarmListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestFullTextSearchFieldsSlice []*string

	EventAndAlarmRetrieveAlarmListPostRequestFullTextSearch struct {
		Fields EventAndAlarmRetrieveAlarmListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                            `json:"type,omitempty"`
		Value  *string                                                            `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequestOptions struct {
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
		GuestPassExpiration           *EventAndAlarmRetrieveAlarmListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *EventAndAlarmRetrieveAlarmListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	EventAndAlarmRetrieveAlarmListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPostRequest struct {
		Attributes      EventAndAlarmRetrieveAlarmListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    EventAndAlarmRetrieveAlarmListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters EventAndAlarmRetrieveAlarmListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *EventAndAlarmRetrieveAlarmListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         EventAndAlarmRetrieveAlarmListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *EventAndAlarmRetrieveAlarmListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                      `json:"limit,omitempty"`
		Options         *EventAndAlarmRetrieveAlarmListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                          `json:"page,omitempty"`
		Query           *string                                                       `json:"query,omitempty"`
		SortInfo        *EventAndAlarmRetrieveAlarmListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                      `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPost200ResponseListSlice []*EventAndAlarmRetrieveAlarmListPost200ResponseList

	EventAndAlarmRetrieveAlarmListPost200ResponseList struct {
		AckTime       *float64 `json:"ackTime,omitempty"`
		AckUser       *string  `json:"ackUser,omitempty"`
		Acknowledged  *string  `json:"acknowledged,omitempty"`
		Activity      *string  `json:"activity,omitempty"`
		AlarmCode     *float64 `json:"alarmCode,omitempty"`
		AlarmState    *string  `json:"alarmState,omitempty"`
		AlarmType     *string  `json:"alarmType,omitempty"`
		Category      *string  `json:"category,omitempty"`
		ClearComment  *string  `json:"clearComment,omitempty"`
		ClearTime     *float64 `json:"clearTime,omitempty"`
		ClearUser     *string  `json:"clearUser,omitempty"`
		ID            *string  `json:"id,omitempty"`
		InsertionTime *float64 `json:"insertionTime,omitempty"`
		Severity      *string  `json:"severity,omitempty"`
	}

	EventAndAlarmRetrieveAlarmListPost200Response struct {
		Extra             *json.RawMessage                                       `json:"extra,omitempty"`
		FirstIndex        *int                                                   `json:"firstIndex,omitempty"`
		HasMore           *bool                                                  `json:"hasMore,omitempty"`
		List              EventAndAlarmRetrieveAlarmListPost200ResponseListSlice `json:"list,omitempty"`
		RawDataTotalCount *int                                                   `json:"rawDataTotalCount,omitempty"`
		TotalCount        *int                                                   `json:"totalCount,omitempty"`
	}
)

// EventAndAlarmRetrieveAlarmListPost: Query Alarms with specified filters.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *EventAndAlarmRetrieveAlarmListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *EventAndAlarmRetrieveAlarmListPost200Response
// - error: Error seen or nil if none
func (a *AlertAPI) EventAndAlarmRetrieveAlarmListPost(ctx *UserContext, requestBody *EventAndAlarmRetrieveAlarmListPostRequest) (*http.Response, *EventAndAlarmRetrieveAlarmListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "POST", "/v5_0/alert/alarm/list")
	request.body = requestBody
	request.authenticated = true
	out := &EventAndAlarmRetrieveAlarmListPost200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// EventAndAlarmAcknowledgeAlarmPut: Acknowledge a single Alarm with provided alarmID.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - alarmID (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AlertAPI) EventAndAlarmAcknowledgeAlarmPut(ctx *UserContext, alarmID string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(alarmID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"alarmID\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "PUT", "/v5_0/alert/alarm/{alarmID}/ack")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"alarmID": alarmID,
	}
	return a.client.doRequest(request, 200, nil)
}

// EventAndAlarmClearAlarmPut: Clear a single Alarm with provided alarmID.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - alarmID (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AlertAPI) EventAndAlarmClearAlarmPut(ctx *UserContext, alarmID string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(alarmID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"alarmID\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "PUT", "/v5_0/alert/alarm/{alarmID}/clear")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"alarmID": alarmID,
	}
	return a.client.doRequest(request, 200, nil)
}

type (
	EventAndAlarmRetrieveEventListPostRequestAttributesSlice []*string

	EventAndAlarmRetrieveEventListPostRequestExtraFiltersSlice []*EventAndAlarmRetrieveEventListPostRequestExtraFilters

	EventAndAlarmRetrieveEventListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestExtraNotFiltersSlice []*EventAndAlarmRetrieveEventListPostRequestExtraNotFilters

	EventAndAlarmRetrieveEventListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestFiltersSlice []*EventAndAlarmRetrieveEventListPostRequestFilters

	EventAndAlarmRetrieveEventListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestFullTextSearchFieldsSlice []*string

	EventAndAlarmRetrieveEventListPostRequestFullTextSearch struct {
		Fields EventAndAlarmRetrieveEventListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                            `json:"type,omitempty"`
		Value  *string                                                            `json:"value,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequestOptions struct {
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
		GuestPassExpiration           *EventAndAlarmRetrieveEventListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                              `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *EventAndAlarmRetrieveEventListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	EventAndAlarmRetrieveEventListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	EventAndAlarmRetrieveEventListPostRequest struct {
		Attributes      EventAndAlarmRetrieveEventListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                       `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                         `json:"expandDomains,omitempty"`
		ExtraFilters    EventAndAlarmRetrieveEventListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters EventAndAlarmRetrieveEventListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *EventAndAlarmRetrieveEventListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         EventAndAlarmRetrieveEventListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *EventAndAlarmRetrieveEventListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                      `json:"limit,omitempty"`
		Options         *EventAndAlarmRetrieveEventListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                          `json:"page,omitempty"`
		Query           *string                                                       `json:"query,omitempty"`
		SortInfo        *EventAndAlarmRetrieveEventListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                      `json:"start,omitempty"`
	}

	EventAndAlarmRetrieveEventListPost200ResponseListSlice []*EventAndAlarmRetrieveEventListPost200ResponseList

	EventAndAlarmRetrieveEventListPost200ResponseList struct {
		Activity      *string  `json:"activity,omitempty"`
		Category      *string  `json:"category,omitempty"`
		EventCode     *float64 `json:"eventCode,omitempty"`
		EventType     *string  `json:"eventType,omitempty"`
		ID            *string  `json:"id,omitempty"`
		InsertionTime *float64 `json:"insertionTime,omitempty"`
		Severity      *string  `json:"severity,omitempty"`
	}

	EventAndAlarmRetrieveEventListPost200Response struct {
		Extra             *json.RawMessage                                       `json:"extra,omitempty"`
		FirstIndex        *int                                                   `json:"firstIndex,omitempty"`
		HasMore           *bool                                                  `json:"hasMore,omitempty"`
		List              EventAndAlarmRetrieveEventListPost200ResponseListSlice `json:"list,omitempty"`
		RawDataTotalCount *int                                                   `json:"rawDataTotalCount,omitempty"`
		TotalCount        *int                                                   `json:"totalCount,omitempty"`
	}
)

// EventAndAlarmRetrieveEventListPost: Query Events with specified filters.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *EventAndAlarmRetrieveEventListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *EventAndAlarmRetrieveEventListPost200Response
// - error: Error seen or nil if none
func (a *AlertAPI) EventAndAlarmRetrieveEventListPost(ctx *UserContext, requestBody *EventAndAlarmRetrieveEventListPostRequest) (*http.Response, *EventAndAlarmRetrieveEventListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "POST", "/v5_0/alert/event/list")
	request.body = requestBody
	request.authenticated = true
	out := &EventAndAlarmRetrieveEventListPost200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
