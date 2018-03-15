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

type ServicesAPI struct {
	client *Client
}
type (
	AccountingServiceDeleteAListOfAccountingServiceDeleteRequestIDListSlice []*string

	AccountingServiceDeleteAListOfAccountingServiceDeleteRequest struct {
		IDList AccountingServiceDeleteAListOfAccountingServiceDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// AccountingServiceDeleteAListOfAccountingServiceDelete: Use this API command to delete a list of accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AccountingServiceDeleteAListOfAccountingServiceDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceDeleteAListOfAccountingServiceDelete(ctx *UserContext, requestBody *AccountingServiceDeleteAListOfAccountingServiceDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/acct")
	request.body = requestBody
	request.authenticated = true
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestAttributesSlice []*string

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraFiltersSlice []*AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraFilters

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraNotFilters

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFiltersSlice []*AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFilters

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                     `json:"type,omitempty"`
		Value  *string                                                                                                     `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                         `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                       `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                         `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                         `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                       `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                         `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                         `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                         `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                         `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                         `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                         `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                       `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                         `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                       `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                       `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                       `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                       `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                       `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                       `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                       `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                       `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                       `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                       `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                       `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                       `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                       `json:"localUser_userSource,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequest struct {
		Attributes      AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                   `json:"limit,omitempty"`
		Options         *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                   `json:"page,omitempty"`
		Query           *string                                                                                                `json:"query,omitempty"`
		SortInfo        *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                   `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200ResponseListSlice []*AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200ResponseList

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200ResponseList struct {
		CreateDateTime   *int    `json:"createDateTime,omitempty"`
		CreatorID        *string `json:"creatorId,omitempty"`
		CreatorUsername  *string `json:"creatorUsername,omitempty"`
		Description      *string `json:"description,omitempty"`
		DomainID         *string `json:"domainId,omitempty"`
		ID               *string `json:"id,omitempty"`
		ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
		ModifierID       *string `json:"modifierId,omitempty"`
		ModifierUsername *string `json:"modifierUsername,omitempty"`
		MvnoID           *string `json:"mvnoId,omitempty"`
		Name             *string `json:"name,omitempty"`
		Protocol         *string `json:"protocol,omitempty"`
		Type             *string `json:"type,omitempty"`
	}

	AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                                `json:"extra,omitempty"`
		FirstIndex *int                                                                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                           `json:"hasMore,omitempty"`
		List       AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                            `json:"totalCount,omitempty"`
	}
)

// AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost: Use this API command to retrieve a list of accounting services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost(ctx *UserContext, requestBody *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPostRequest) (*http.Response, *AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/acct/query")
	request.body = requestBody
	request.authenticated = true
	out := &AccountingServiceRetrieveListAllTypesOfAccountingServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListSlice []*AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseList

	AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseList struct {
		CreateDateTime    *int                                                                                     `json:"createDateTime,omitempty"`
		CreatorID         *string                                                                                  `json:"creatorId,omitempty"`
		CreatorUsername   *string                                                                                  `json:"creatorUsername,omitempty"`
		Description       *string                                                                                  `json:"description,omitempty"`
		DomainID          *string                                                                                  `json:"domainId,omitempty"`
		HealthCheckPolicy *AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                *string                                                                                  `json:"id,omitempty"`
		ModifiedDateTime  *int                                                                                     `json:"modifiedDateTime,omitempty"`
		ModifierID        *string                                                                                  `json:"modifierId,omitempty"`
		ModifierUsername  *string                                                                                  `json:"modifierUsername,omitempty"`
		MvnoID            *string                                                                                  `json:"mvnoId,omitempty"`
		Name              *string                                                                                  `json:"name,omitempty"`
		Primary           *AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListPrimary           `json:"primary,omitempty"`
		Protocol          *string                                                                                  `json:"protocol,omitempty"`
		RateLimiting      *AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary         *AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListSecondary         `json:"secondary,omitempty"`
		Type              *string                                                                                  `json:"type,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceGet200Response struct {
		Extra      *json.RawMessage                                                            `json:"extra,omitempty"`
		FirstIndex *int                                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                       `json:"hasMore,omitempty"`
		List       AccountingServiceRetrieveListRadiusAccountingServiceGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                        `json:"totalCount,omitempty"`
	}
)

// AccountingServiceRetrieveListRadiusAccountingServiceGet: Use this API command to retrieve a list of RADIUS accounting services.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingServiceRetrieveListRadiusAccountingServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceRetrieveListRadiusAccountingServiceGet(ctx *UserContext) (*http.Response, *AccountingServiceRetrieveListRadiusAccountingServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/acct/radius")
	request.authenticated = true
	out := &AccountingServiceRetrieveListRadiusAccountingServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AccountingServiceCreateRadiusAccountingServicePostRequestHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AccountingServiceCreateRadiusAccountingServicePostRequestPrimary struct {
		DomainID     *string `json:"domainId,omitempty"`
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceCreateRadiusAccountingServicePostRequestRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AccountingServiceCreateRadiusAccountingServicePostRequestSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceCreateRadiusAccountingServicePostRequest struct {
		Description       *string                                                                     `json:"description,omitempty"`
		DomainID          *string                                                                     `json:"domainId,omitempty"`
		HealthCheckPolicy *AccountingServiceCreateRadiusAccountingServicePostRequestHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		Name              *string                                                                     `json:"name,omitempty"`
		Primary           *AccountingServiceCreateRadiusAccountingServicePostRequestPrimary           `json:"primary,omitempty"`
		Protocol          *string                                                                     `json:"protocol,omitempty"`
		RateLimiting      *AccountingServiceCreateRadiusAccountingServicePostRequestRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary         *AccountingServiceCreateRadiusAccountingServicePostRequestSecondary         `json:"secondary,omitempty"`
		Type              *string                                                                     `json:"type,omitempty"`
	}

	AccountingServiceCreateRadiusAccountingServicePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AccountingServiceCreateRadiusAccountingServicePost: Use this API command to create a new RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AccountingServiceCreateRadiusAccountingServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingServiceCreateRadiusAccountingServicePost201Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceCreateRadiusAccountingServicePost(ctx *UserContext, requestBody *AccountingServiceCreateRadiusAccountingServicePostRequest) (*http.Response, *AccountingServiceCreateRadiusAccountingServicePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/acct/radius")
	request.body = requestBody
	request.authenticated = true
	out := &AccountingServiceCreateRadiusAccountingServicePost201Response{}
	httpResponse, _, err := s.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestAttributesSlice []*string

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraFiltersSlice []*AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraFilters

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraNotFilters

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFiltersSlice []*AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFilters

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                 `json:"type,omitempty"`
		Value  *string                                                                                                 `json:"value,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestOptions struct {
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
		GuestPassExpiration           *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                   `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                     `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                     `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                     `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequest struct {
		Attributes      AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                            `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                              `json:"expandDomains,omitempty"`
		ExtraFilters    AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                               `json:"limit,omitempty"`
		Options         *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                               `json:"page,omitempty"`
		Query           *string                                                                                            `json:"query,omitempty"`
		SortInfo        *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                               `json:"start,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListSlice []*AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseList

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseList struct {
		CreateDateTime    *int                                                                                                     `json:"createDateTime,omitempty"`
		CreatorID         *string                                                                                                  `json:"creatorId,omitempty"`
		CreatorUsername   *string                                                                                                  `json:"creatorUsername,omitempty"`
		Description       *string                                                                                                  `json:"description,omitempty"`
		DomainID          *string                                                                                                  `json:"domainId,omitempty"`
		HealthCheckPolicy *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                *string                                                                                                  `json:"id,omitempty"`
		ModifiedDateTime  *int                                                                                                     `json:"modifiedDateTime,omitempty"`
		ModifierID        *string                                                                                                  `json:"modifierId,omitempty"`
		ModifierUsername  *string                                                                                                  `json:"modifierUsername,omitempty"`
		MvnoID            *string                                                                                                  `json:"mvnoId,omitempty"`
		Name              *string                                                                                                  `json:"name,omitempty"`
		Primary           *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListPrimary           `json:"primary,omitempty"`
		Protocol          *string                                                                                                  `json:"protocol,omitempty"`
		RateLimiting      *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary         *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListSecondary         `json:"secondary,omitempty"`
		Type              *string                                                                                                  `json:"type,omitempty"`
	}

	AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                            `json:"extra,omitempty"`
		FirstIndex *int                                                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                       `json:"hasMore,omitempty"`
		List       AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                        `json:"totalCount,omitempty"`
	}
)

// AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost: Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost(ctx *UserContext, requestBody *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPostRequest) (*http.Response, *AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/acct/radius/query")
	request.body = requestBody
	request.authenticated = true
	out := &AccountingServiceRetrieveListRadiusAccountingServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AccountingServiceDeleteRadiusAccountingServiceDelete: Use this API command to delete a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceDeleteRadiusAccountingServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/acct/radius/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceRetrieveRadiusAccountingServiceGet200ResponseHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AccountingServiceRetrieveRadiusAccountingServiceGet200ResponsePrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceRetrieveRadiusAccountingServiceGet200ResponseRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AccountingServiceRetrieveRadiusAccountingServiceGet200ResponseSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AccountingServiceRetrieveRadiusAccountingServiceGet200Response struct {
		CreateDateTime    *int                                                                             `json:"createDateTime,omitempty"`
		CreatorID         *string                                                                          `json:"creatorId,omitempty"`
		CreatorUsername   *string                                                                          `json:"creatorUsername,omitempty"`
		Description       *string                                                                          `json:"description,omitempty"`
		DomainID          *string                                                                          `json:"domainId,omitempty"`
		HealthCheckPolicy *AccountingServiceRetrieveRadiusAccountingServiceGet200ResponseHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                *string                                                                          `json:"id,omitempty"`
		ModifiedDateTime  *int                                                                             `json:"modifiedDateTime,omitempty"`
		ModifierID        *string                                                                          `json:"modifierId,omitempty"`
		ModifierUsername  *string                                                                          `json:"modifierUsername,omitempty"`
		MvnoID            *string                                                                          `json:"mvnoId,omitempty"`
		Name              *string                                                                          `json:"name,omitempty"`
		Primary           *AccountingServiceRetrieveRadiusAccountingServiceGet200ResponsePrimary           `json:"primary,omitempty"`
		Protocol          *string                                                                          `json:"protocol,omitempty"`
		RateLimiting      *AccountingServiceRetrieveRadiusAccountingServiceGet200ResponseRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary         *AccountingServiceRetrieveRadiusAccountingServiceGet200ResponseSecondary         `json:"secondary,omitempty"`
		Type              *string                                                                          `json:"type,omitempty"`
	}
)

// AccountingServiceRetrieveRadiusAccountingServiceGet: Use this API command to retrieve a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccountingServiceRetrieveRadiusAccountingServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceRetrieveRadiusAccountingServiceGet(ctx *UserContext, id string) (*http.Response, *AccountingServiceRetrieveRadiusAccountingServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/acct/radius/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AccountingServiceRetrieveRadiusAccountingServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AccountingServiceModifyRadiusAccountingServicePatchRequest struct {
		Description *string `json:"description,omitempty"`
		DomainID    *string `json:"domainId,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		Protocol    *string `json:"protocol,omitempty"`
		Type        *string `json:"type,omitempty"`
	}
)

// AccountingServiceModifyRadiusAccountingServicePatch: Use this API command to modify the basic information of a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AccountingServiceModifyRadiusAccountingServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceModifyRadiusAccountingServicePatch(ctx *UserContext, id string, requestBody *AccountingServiceModifyRadiusAccountingServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/acct/radius/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceModifyHealthCheckPolicyOfAccountingServicePatchRequest struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}
)

// AccountingServiceModifyHealthCheckPolicyOfAccountingServicePatch: Use this API command to modify health check policy of a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AccountingServiceModifyHealthCheckPolicyOfAccountingServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceModifyHealthCheckPolicyOfAccountingServicePatch(ctx *UserContext, id string, requestBody *AccountingServiceModifyHealthCheckPolicyOfAccountingServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/acct/radius/{id}/healthCheckPolicy")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceModifyPrimaryRadiusServerOfAccountingServicePatchRequest struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}
)

// AccountingServiceModifyPrimaryRadiusServerOfAccountingServicePatch: Use this API command to modify primary RADIUS server of a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AccountingServiceModifyPrimaryRadiusServerOfAccountingServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceModifyPrimaryRadiusServerOfAccountingServicePatch(ctx *UserContext, id string, requestBody *AccountingServiceModifyPrimaryRadiusServerOfAccountingServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/acct/radius/{id}/primary")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceModifyRateLimitingOfAccountingServicePatchRequest struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}
)

// AccountingServiceModifyRateLimitingOfAccountingServicePatch: Use this API command to modify rate limiting of a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AccountingServiceModifyRateLimitingOfAccountingServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceModifyRateLimitingOfAccountingServicePatch(ctx *UserContext, id string, requestBody *AccountingServiceModifyRateLimitingOfAccountingServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/acct/radius/{id}/rateLimiting")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

// AccountingServiceDisableSecondaryRadiusServerOfAccountingServiceDelete: Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceDisableSecondaryRadiusServerOfAccountingServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/acct/radius/{id}/secondary")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceModifySecondaryRadiusServerOfAccountingServicePatchRequest struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}
)

// AccountingServiceModifySecondaryRadiusServerOfAccountingServicePatch: Use this API command to modify secondary RADIUS server of a RADIUS accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AccountingServiceModifySecondaryRadiusServerOfAccountingServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceModifySecondaryRadiusServerOfAccountingServicePatch(ctx *UserContext, id string, requestBody *AccountingServiceModifySecondaryRadiusServerOfAccountingServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/acct/radius/{id}/secondary")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AccountingServiceTestSpecificAccountingServicePostRequestLoginRequest struct {
		Password          *string `json:"password,omitempty"`
		Protocol          *string `json:"protocol,omitempty"`
		TimeZoneUtcOffset *string `json:"timeZoneUtcOffset,omitempty"`
		UserName          *string `json:"userName,omitempty"`
	}

	AccountingServiceTestSpecificAccountingServicePostRequest struct {
		ID           *string                                                                `json:"id,omitempty"`
		LoginRequest *AccountingServiceTestSpecificAccountingServicePostRequestLoginRequest `json:"loginRequest,omitempty"`
	}
)

// AccountingServiceTestSpecificAccountingServicePost: Use this API command to test an accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AccountingServiceTestSpecificAccountingServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceTestSpecificAccountingServicePost(ctx *UserContext, id string, requestBody *AccountingServiceTestSpecificAccountingServicePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/acct/test/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

// AccountingServiceDeleteAccountingServiceDelete: Use this API command to delete an accounting service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AccountingServiceDeleteAccountingServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/acct/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceDeleteAListOfAuthenticationServiceDeleteRequestIDListSlice []*string

	AuthenticationServiceDeleteAListOfAuthenticationServiceDeleteRequest struct {
		IDList AuthenticationServiceDeleteAListOfAuthenticationServiceDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// AuthenticationServiceDeleteAListOfAuthenticationServiceDelete: Use this API command to delete a list of authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceDeleteAListOfAuthenticationServiceDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDeleteAListOfAuthenticationServiceDelete(ctx *UserContext, requestBody *AuthenticationServiceDeleteAListOfAuthenticationServiceDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth")
	request.body = requestBody
	request.authenticated = true
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListSlice []*AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseList

	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappings

	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                                    `json:"id,omitempty"`
		Name               *string                                                                                                                    `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappings struct {
		GroupAttr *string                                                                                                  `json:"groupAttr,omitempty"`
		ID        *string                                                                                                  `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseList struct {
		AdminDomainName      *string                                                                                              `json:"adminDomainName,omitempty"`
		CreateDateTime       *int                                                                                                 `json:"createDateTime,omitempty"`
		CreatorID            *string                                                                                              `json:"creatorId,omitempty"`
		CreatorUsername      *string                                                                                              `json:"creatorUsername,omitempty"`
		Description          *string                                                                                              `json:"description,omitempty"`
		DomainID             *string                                                                                              `json:"domainId,omitempty"`
		FriendlyName         *string                                                                                              `json:"friendlyName,omitempty"`
		GlobalCatalogEnabled *bool                                                                                                `json:"globalCatalogEnabled,omitempty"`
		ID                   *string                                                                                              `json:"id,omitempty"`
		IP                   *string                                                                                              `json:"ip,omitempty"`
		Mappings             AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime     *int                                                                                                 `json:"modifiedDateTime,omitempty"`
		ModifierID           *string                                                                                              `json:"modifierId,omitempty"`
		ModifierUsername     *string                                                                                              `json:"modifierUsername,omitempty"`
		MvnoID               *string                                                                                              `json:"mvnoId,omitempty"`
		Name                 *string                                                                                              `json:"name,omitempty"`
		Password             *string                                                                                              `json:"password,omitempty"`
		Port                 *int                                                                                                 `json:"port,omitempty"`
		Protocol             *string                                                                                              `json:"protocol,omitempty"`
		TLSEnabled           *bool                                                                                                `json:"tlsEnabled,omitempty"`
		Type                 *string                                                                                              `json:"type,omitempty"`
		WindowsDomainName    *string                                                                                              `json:"windowsDomainName,omitempty"`
	}

	AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200Response struct {
		Extra      *json.RawMessage                                                                             `json:"extra,omitempty"`
		FirstIndex *int                                                                                         `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                        `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                         `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet: Use this API command to retrieve a list of active directory authentication services.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet(ctx *UserContext) (*http.Response, *AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/ad")
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListActiveDirectoryAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequestMappingsSlice []*AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequestMappings

	AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequestMappingsUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequestMappings struct {
		GroupAttr *string                                                                                     `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequestMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequest struct {
		AdminDomainName      *string                                                                                 `json:"adminDomainName,omitempty"`
		Description          *string                                                                                 `json:"description,omitempty"`
		DomainID             *string                                                                                 `json:"domainId,omitempty"`
		FriendlyName         *string                                                                                 `json:"friendlyName,omitempty"`
		GlobalCatalogEnabled *bool                                                                                   `json:"globalCatalogEnabled,omitempty"`
		ID                   *string                                                                                 `json:"id,omitempty"`
		IP                   *string                                                                                 `json:"ip,omitempty"`
		Mappings             AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequestMappingsSlice `json:"mappings,omitempty"`
		Name                 *string                                                                                 `json:"name,omitempty"`
		Password             *string                                                                                 `json:"password,omitempty"`
		Port                 *int                                                                                    `json:"port,omitempty"`
		TLSEnabled           *bool                                                                                   `json:"tlsEnabled,omitempty"`
		Type                 *string                                                                                 `json:"type,omitempty"`
		WindowsDomainName    *string                                                                                 `json:"windowsDomainName,omitempty"`
	}

	AuthenticationServiceCreateActiveDirectoryAuthenticationServicePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AuthenticationServiceCreateActiveDirectoryAuthenticationServicePost: Use this API command to create a new active directory authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceCreateActiveDirectoryAuthenticationServicePost201Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceCreateActiveDirectoryAuthenticationServicePost(ctx *UserContext, requestBody *AuthenticationServiceCreateActiveDirectoryAuthenticationServicePostRequest) (*http.Response, *AuthenticationServiceCreateActiveDirectoryAuthenticationServicePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/ad")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceCreateActiveDirectoryAuthenticationServicePost201Response{}
	httpResponse, _, err := s.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestAttributesSlice []*string

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice []*AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraFilters

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFiltersSlice []*AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFilters

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                     `json:"type,omitempty"`
		Value  *string                                                                                                     `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                         `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                       `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                         `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                         `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                       `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                         `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                         `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                         `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                         `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                         `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                         `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                       `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                         `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                       `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                       `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                       `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                       `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                       `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                         `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                         `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                         `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                       `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                       `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                       `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                       `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                       `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                       `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                       `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                       `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                       `json:"localUser_userSource,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequest struct {
		Attributes      AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                                `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                  `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                   `json:"limit,omitempty"`
		Options         *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                   `json:"page,omitempty"`
		Query           *string                                                                                                `json:"query,omitempty"`
		SortInfo        *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                   `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListSlice []*AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseList

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappings

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                                       `json:"id,omitempty"`
		Name               *string                                                                                                                       `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappings struct {
		GroupAttr *string                                                                                                     `json:"groupAttr,omitempty"`
		ID        *string                                                                                                     `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseList struct {
		AdminDomainName      *string                                                                                                 `json:"adminDomainName,omitempty"`
		CreateDateTime       *int                                                                                                    `json:"createDateTime,omitempty"`
		CreatorID            *string                                                                                                 `json:"creatorId,omitempty"`
		CreatorUsername      *string                                                                                                 `json:"creatorUsername,omitempty"`
		Description          *string                                                                                                 `json:"description,omitempty"`
		DomainID             *string                                                                                                 `json:"domainId,omitempty"`
		FriendlyName         *string                                                                                                 `json:"friendlyName,omitempty"`
		GlobalCatalogEnabled *bool                                                                                                   `json:"globalCatalogEnabled,omitempty"`
		ID                   *string                                                                                                 `json:"id,omitempty"`
		IP                   *string                                                                                                 `json:"ip,omitempty"`
		Mappings             AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime     *int                                                                                                    `json:"modifiedDateTime,omitempty"`
		ModifierID           *string                                                                                                 `json:"modifierId,omitempty"`
		ModifierUsername     *string                                                                                                 `json:"modifierUsername,omitempty"`
		MvnoID               *string                                                                                                 `json:"mvnoId,omitempty"`
		Name                 *string                                                                                                 `json:"name,omitempty"`
		Password             *string                                                                                                 `json:"password,omitempty"`
		Port                 *int                                                                                                    `json:"port,omitempty"`
		Protocol             *string                                                                                                 `json:"protocol,omitempty"`
		TLSEnabled           *bool                                                                                                   `json:"tlsEnabled,omitempty"`
		Type                 *string                                                                                                 `json:"type,omitempty"`
		WindowsDomainName    *string                                                                                                 `json:"windowsDomainName,omitempty"`
	}

	AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                                `json:"extra,omitempty"`
		FirstIndex *int                                                                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                           `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                            `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost: Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost(ctx *UserContext, requestBody *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPostRequest) (*http.Response, *AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/ad/query")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListAdAuthenticationServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AuthenticationServiceDeleteActiveDirectoryAuthenticationServiceDelete: Use this API command to delete an active directory authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): Active Directory Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDeleteActiveDirectoryAuthenticationServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth/ad/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappingsSlice []*AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappings

	AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappingsUserRole struct {
		ID                 *string                                                                                                            `json:"id,omitempty"`
		Name               *string                                                                                                            `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappings struct {
		GroupAttr *string                                                                                          `json:"groupAttr,omitempty"`
		ID        *string                                                                                          `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200Response struct {
		AdminDomainName      *string                                                                                      `json:"adminDomainName,omitempty"`
		CreateDateTime       *int                                                                                         `json:"createDateTime,omitempty"`
		CreatorID            *string                                                                                      `json:"creatorId,omitempty"`
		CreatorUsername      *string                                                                                      `json:"creatorUsername,omitempty"`
		Description          *string                                                                                      `json:"description,omitempty"`
		DomainID             *string                                                                                      `json:"domainId,omitempty"`
		FriendlyName         *string                                                                                      `json:"friendlyName,omitempty"`
		GlobalCatalogEnabled *bool                                                                                        `json:"globalCatalogEnabled,omitempty"`
		ID                   *string                                                                                      `json:"id,omitempty"`
		IP                   *string                                                                                      `json:"ip,omitempty"`
		Mappings             AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200ResponseMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime     *int                                                                                         `json:"modifiedDateTime,omitempty"`
		ModifierID           *string                                                                                      `json:"modifierId,omitempty"`
		ModifierUsername     *string                                                                                      `json:"modifierUsername,omitempty"`
		MvnoID               *string                                                                                      `json:"mvnoId,omitempty"`
		Name                 *string                                                                                      `json:"name,omitempty"`
		Password             *string                                                                                      `json:"password,omitempty"`
		Port                 *int                                                                                         `json:"port,omitempty"`
		Protocol             *string                                                                                      `json:"protocol,omitempty"`
		TLSEnabled           *bool                                                                                        `json:"tlsEnabled,omitempty"`
		Type                 *string                                                                                      `json:"type,omitempty"`
		WindowsDomainName    *string                                                                                      `json:"windowsDomainName,omitempty"`
	}
)

// AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet: Use this API command to retrieve an active directory authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): Active Directory Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet(ctx *UserContext, id string) (*http.Response, *AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/ad/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AuthenticationServiceRetrieveActiveDirectoryAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceModifyActiveDirectoryAuthenticationServicePatchRequest struct {
		AdminDomainName      *string `json:"adminDomainName,omitempty"`
		Description          *string `json:"description,omitempty"`
		DomainID             *string `json:"domainId,omitempty"`
		FriendlyName         *string `json:"friendlyName,omitempty"`
		GlobalCatalogEnabled *bool   `json:"globalCatalogEnabled,omitempty"`
		ID                   *string `json:"id,omitempty"`
		IP                   *string `json:"ip,omitempty"`
		Name                 *string `json:"name,omitempty"`
		Password             *string `json:"password,omitempty"`
		Port                 *int    `json:"port,omitempty"`
		TLSEnabled           *bool   `json:"tlsEnabled,omitempty"`
		Type                 *string `json:"type,omitempty"`
		WindowsDomainName    *string `json:"windowsDomainName,omitempty"`
	}
)

// AuthenticationServiceModifyActiveDirectoryAuthenticationServicePatch: Use this API command to modify the basic information of an active directory authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): Active Directory Service ID
// - requestBody: *AuthenticationServiceModifyActiveDirectoryAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyActiveDirectoryAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyActiveDirectoryAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/ad/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequestSlice []*AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequest

	AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequestUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequest struct {
		GroupAttr *string                                                                                    `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequestUserRole `json:"userRole,omitempty"`
	}
)

// AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatch: Use this API command to modify user traffic profile mapping of an active directory authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): Active Directory Service ID
// - requestBody: *AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatch(ctx *UserContext, id string, requestBody AuthenticationServiceModifyUserTrafficProfileMappingOfActiveDirectoryPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/ad/{id}/mappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappingsSlice []*AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappings

	AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappingsUserRole struct {
		ID                 *string                                                                                                  `json:"id,omitempty"`
		Name               *string                                                                                                  `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappings struct {
		GroupAttr *string                                                                                `json:"groupAttr,omitempty"`
		ID        *string                                                                                `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveGuestAuthenticationServiceGet200Response struct {
		CreateDateTime   *int                                                                               `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                            `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                            `json:"creatorUsername,omitempty"`
		Description      *string                                                                            `json:"description,omitempty"`
		DomainID         *string                                                                            `json:"domainId,omitempty"`
		FriendlyName     *string                                                                            `json:"friendlyName,omitempty"`
		ID               *string                                                                            `json:"id,omitempty"`
		Mappings         AuthenticationServiceRetrieveGuestAuthenticationServiceGet200ResponseMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime *int                                                                               `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                            `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                            `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                            `json:"mvnoId,omitempty"`
		Name             *string                                                                            `json:"name,omitempty"`
		Protocol         *string                                                                            `json:"protocol,omitempty"`
		Type             *string                                                                            `json:"type,omitempty"`
	}
)

// AuthenticationServiceRetrieveGuestAuthenticationServiceGet: Use this API command to retrieve a Guest authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveGuestAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveGuestAuthenticationServiceGet(ctx *UserContext, id string) (*http.Response, *AuthenticationServiceRetrieveGuestAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/guest/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AuthenticationServiceRetrieveGuestAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseList

	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappings

	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                         `json:"id,omitempty"`
		Name               *string                                                                                                         `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappings struct {
		GroupAttr *string                                                                                       `json:"groupAttr,omitempty"`
		ID        *string                                                                                       `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseList struct {
		AdminDomainName  *string                                                                                   `json:"adminDomainName,omitempty"`
		BaseDomainName   *string                                                                                   `json:"baseDomainName,omitempty"`
		CreateDateTime   *int                                                                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                   `json:"creatorUsername,omitempty"`
		Description      *string                                                                                   `json:"description,omitempty"`
		DomainID         *string                                                                                   `json:"domainId,omitempty"`
		FriendlyName     *string                                                                                   `json:"friendlyName,omitempty"`
		ID               *string                                                                                   `json:"id,omitempty"`
		IP               *string                                                                                   `json:"ip,omitempty"`
		KeyAttribute     *string                                                                                   `json:"keyAttribute,omitempty"`
		Mappings         AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime *int                                                                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                   `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                                   `json:"mvnoId,omitempty"`
		Name             *string                                                                                   `json:"name,omitempty"`
		Password         *string                                                                                   `json:"password,omitempty"`
		Port             *int                                                                                      `json:"port,omitempty"`
		Protocol         *string                                                                                   `json:"protocol,omitempty"`
		SearchFilter     *string                                                                                   `json:"searchFilter,omitempty"`
		TLSEnabled       *bool                                                                                     `json:"tlsEnabled,omitempty"`
		Type             *string                                                                                   `json:"type,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200Response struct {
		Extra      *json.RawMessage                                                                  `json:"extra,omitempty"`
		FirstIndex *int                                                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                             `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                              `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListLdapAuthenticationServiceGet: Use this API command to retrieve a list of LDAP authentication services.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListLdapAuthenticationServiceGet(ctx *UserContext) (*http.Response, *AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/ldap")
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListLdapAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceCreateLdapAuthenticationServicePostRequestMappingsSlice []*AuthenticationServiceCreateLdapAuthenticationServicePostRequestMappings

	AuthenticationServiceCreateLdapAuthenticationServicePostRequestMappingsUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceCreateLdapAuthenticationServicePostRequestMappings struct {
		GroupAttr *string                                                                          `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceCreateLdapAuthenticationServicePostRequestMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceCreateLdapAuthenticationServicePostRequest struct {
		AdminDomainName *string                                                                      `json:"adminDomainName,omitempty"`
		BaseDomainName  *string                                                                      `json:"baseDomainName,omitempty"`
		Description     *string                                                                      `json:"description,omitempty"`
		DomainID        *string                                                                      `json:"domainId,omitempty"`
		FriendlyName    *string                                                                      `json:"friendlyName,omitempty"`
		ID              *string                                                                      `json:"id,omitempty"`
		IP              *string                                                                      `json:"ip,omitempty"`
		KeyAttribute    *string                                                                      `json:"keyAttribute,omitempty"`
		Mappings        AuthenticationServiceCreateLdapAuthenticationServicePostRequestMappingsSlice `json:"mappings,omitempty"`
		Name            *string                                                                      `json:"name,omitempty"`
		Password        *string                                                                      `json:"password,omitempty"`
		Port            *int                                                                         `json:"port,omitempty"`
		SearchFilter    *string                                                                      `json:"searchFilter,omitempty"`
		TLSEnabled      *bool                                                                        `json:"tlsEnabled,omitempty"`
		Type            *string                                                                      `json:"type,omitempty"`
	}

	AuthenticationServiceCreateLdapAuthenticationServicePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AuthenticationServiceCreateLdapAuthenticationServicePost: Use this API command to create a new LDAP authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceCreateLdapAuthenticationServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceCreateLdapAuthenticationServicePost201Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceCreateLdapAuthenticationServicePost(ctx *UserContext, requestBody *AuthenticationServiceCreateLdapAuthenticationServicePostRequest) (*http.Response, *AuthenticationServiceCreateLdapAuthenticationServicePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/ldap")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceCreateLdapAuthenticationServicePost201Response{}
	httpResponse, _, err := s.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestAttributesSlice []*string

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraFilters

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFiltersSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFilters

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                       `json:"type,omitempty"`
		Value  *string                                                                                                       `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                           `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                         `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                           `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                           `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                         `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                           `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                           `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                           `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                           `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                           `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                           `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                         `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                           `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                         `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                         `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                         `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                         `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                         `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                           `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                           `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                           `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                         `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                         `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                         `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                         `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                         `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                         `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                         `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                         `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                         `json:"localUser_userSource,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequest struct {
		Attributes      AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                                  `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                    `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                     `json:"limit,omitempty"`
		Options         *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                     `json:"page,omitempty"`
		Query           *string                                                                                                  `json:"query,omitempty"`
		SortInfo        *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                     `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseList

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappings

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                                         `json:"id,omitempty"`
		Name               *string                                                                                                                         `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappings struct {
		GroupAttr *string                                                                                                       `json:"groupAttr,omitempty"`
		ID        *string                                                                                                       `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseList struct {
		AdminDomainName  *string                                                                                                   `json:"adminDomainName,omitempty"`
		BaseDomainName   *string                                                                                                   `json:"baseDomainName,omitempty"`
		CreateDateTime   *int                                                                                                      `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                                   `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                                   `json:"creatorUsername,omitempty"`
		Description      *string                                                                                                   `json:"description,omitempty"`
		DomainID         *string                                                                                                   `json:"domainId,omitempty"`
		FriendlyName     *string                                                                                                   `json:"friendlyName,omitempty"`
		ID               *string                                                                                                   `json:"id,omitempty"`
		IP               *string                                                                                                   `json:"ip,omitempty"`
		KeyAttribute     *string                                                                                                   `json:"keyAttribute,omitempty"`
		Mappings         AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime *int                                                                                                      `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                                   `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                                   `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                                                   `json:"mvnoId,omitempty"`
		Name             *string                                                                                                   `json:"name,omitempty"`
		Password         *string                                                                                                   `json:"password,omitempty"`
		Port             *int                                                                                                      `json:"port,omitempty"`
		Protocol         *string                                                                                                   `json:"protocol,omitempty"`
		SearchFilter     *string                                                                                                   `json:"searchFilter,omitempty"`
		TLSEnabled       *bool                                                                                                     `json:"tlsEnabled,omitempty"`
		Type             *string                                                                                                   `json:"type,omitempty"`
	}

	AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                                  `json:"extra,omitempty"`
		FirstIndex *int                                                                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                             `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                              `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost: Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost(ctx *UserContext, requestBody *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPostRequest) (*http.Response, *AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/ldap/query")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListLdapAuthenticationServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AuthenticationServiceDeleteLdapAuthenticationServiceDelete: Use this API command to delete a LDAP authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): LDAP Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDeleteLdapAuthenticationServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth/ldap/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappingsSlice []*AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappings

	AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappingsUserRole struct {
		ID                 *string                                                                                                 `json:"id,omitempty"`
		Name               *string                                                                                                 `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappings struct {
		GroupAttr *string                                                                               `json:"groupAttr,omitempty"`
		ID        *string                                                                               `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveLdapAuthenticationServiceGet200Response struct {
		AdminDomainName  *string                                                                           `json:"adminDomainName,omitempty"`
		BaseDomainName   *string                                                                           `json:"baseDomainName,omitempty"`
		CreateDateTime   *int                                                                              `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                           `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                           `json:"creatorUsername,omitempty"`
		Description      *string                                                                           `json:"description,omitempty"`
		DomainID         *string                                                                           `json:"domainId,omitempty"`
		FriendlyName     *string                                                                           `json:"friendlyName,omitempty"`
		ID               *string                                                                           `json:"id,omitempty"`
		IP               *string                                                                           `json:"ip,omitempty"`
		KeyAttribute     *string                                                                           `json:"keyAttribute,omitempty"`
		Mappings         AuthenticationServiceRetrieveLdapAuthenticationServiceGet200ResponseMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime *int                                                                              `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                           `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                           `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                           `json:"mvnoId,omitempty"`
		Name             *string                                                                           `json:"name,omitempty"`
		Password         *string                                                                           `json:"password,omitempty"`
		Port             *int                                                                              `json:"port,omitempty"`
		Protocol         *string                                                                           `json:"protocol,omitempty"`
		SearchFilter     *string                                                                           `json:"searchFilter,omitempty"`
		TLSEnabled       *bool                                                                             `json:"tlsEnabled,omitempty"`
		Type             *string                                                                           `json:"type,omitempty"`
	}
)

// AuthenticationServiceRetrieveLdapAuthenticationServiceGet: Use this API command to retrieve a LDAP authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): LDAP Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveLdapAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveLdapAuthenticationServiceGet(ctx *UserContext, id string) (*http.Response, *AuthenticationServiceRetrieveLdapAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/ldap/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AuthenticationServiceRetrieveLdapAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceModifyLdapAuthenticationServicePatchRequest struct {
		AdminDomainName *string `json:"adminDomainName,omitempty"`
		BaseDomainName  *string `json:"baseDomainName,omitempty"`
		Description     *string `json:"description,omitempty"`
		DomainID        *string `json:"domainId,omitempty"`
		FriendlyName    *string `json:"friendlyName,omitempty"`
		ID              *string `json:"id,omitempty"`
		IP              *string `json:"ip,omitempty"`
		KeyAttribute    *string `json:"keyAttribute,omitempty"`
		Name            *string `json:"name,omitempty"`
		Password        *string `json:"password,omitempty"`
		Port            *int    `json:"port,omitempty"`
		SearchFilter    *string `json:"searchFilter,omitempty"`
		TLSEnabled      *bool   `json:"tlsEnabled,omitempty"`
		Type            *string `json:"type,omitempty"`
	}
)

// AuthenticationServiceModifyLdapAuthenticationServicePatch: Use this API command to modify the basic information of a LDAP authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): LDAP Service ID
// - requestBody: *AuthenticationServiceModifyLdapAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyLdapAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyLdapAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/ldap/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequestSlice []*AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequest

	AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequestUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequest struct {
		GroupAttr *string                                                                         `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequestUserRole `json:"userRole,omitempty"`
	}
)

// AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatch: Use this API command to modify user traffic profile mapping of a LDAP authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): LDAP Service ID
// - requestBody: *AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatch(ctx *UserContext, id string, requestBody AuthenticationServiceModifyUserTrafficProfileMappingOfLdapPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/ldap/{id}/mappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappingsSlice []*AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappings

	AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappingsUserRole struct {
		ID                 *string                                                                                                    `json:"id,omitempty"`
		Name               *string                                                                                                    `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappings struct {
		GroupAttr *string                                                                                  `json:"groupAttr,omitempty"`
		ID        *string                                                                                  `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200Response struct {
		CreateDateTime   *int                                                                                 `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                              `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                              `json:"creatorUsername,omitempty"`
		Description      *string                                                                              `json:"description,omitempty"`
		DomainID         *string                                                                              `json:"domainId,omitempty"`
		FriendlyName     *string                                                                              `json:"friendlyName,omitempty"`
		ID               *string                                                                              `json:"id,omitempty"`
		Mappings         AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200ResponseMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime *int                                                                                 `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                              `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                              `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                              `json:"mvnoId,omitempty"`
		Name             *string                                                                              `json:"name,omitempty"`
		Protocol         *string                                                                              `json:"protocol,omitempty"`
		Type             *string                                                                              `json:"type,omitempty"`
	}
)

// AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet: Use this API command to retrieve a LocalDB authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet(ctx *UserContext, id string) (*http.Response, *AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/local_db/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AuthenticationServiceRetrieveLocaldbAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceUpdateLocaldbAuthenticationServicePatchRequest struct {
		Description  *string `json:"description,omitempty"`
		DomainID     *string `json:"domainId,omitempty"`
		FriendlyName *string `json:"friendlyName,omitempty"`
		ID           *string `json:"id,omitempty"`
		MvnoID       *string `json:"mvnoId,omitempty"`
		Name         *string `json:"name,omitempty"`
		Protocol     *string `json:"protocol,omitempty"`
		Type         *string `json:"type,omitempty"`
	}
)

// AuthenticationServiceUpdateLocaldbAuthenticationServicePatch: Use this API command to update LocalDB authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AuthenticationServiceUpdateLocaldbAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceUpdateLocaldbAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceUpdateLocaldbAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/local_db/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequestSlice []*AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequest

	AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequestUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequest struct {
		GroupAttr *string                                                                                        `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequestUserRole `json:"userRole,omitempty"`
	}
)

// AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatch: Use this API command to update user role mappings of LocalDB authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatch(ctx *UserContext, id string, requestBody AuthenticationServiceUpdateUserRoleMappingsOfLocaldbAuthenticationServicePatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/local_db/{id}/mappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseList

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListDefaultWhitelistedDomainsSlice []*string

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappings

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                          `json:"id,omitempty"`
		Name               *string                                                                                                          `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappings struct {
		GroupAttr *string                                                                                        `json:"groupAttr,omitempty"`
		ID        *string                                                                                        `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListWhitelistedDomainsSlice []*string

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseList struct {
		ApplicationID             *string                                                                                                     `json:"applicationId,omitempty"`
		ApplicationSecret         *string                                                                                                     `json:"applicationSecret,omitempty"`
		CreateDateTime            *int                                                                                                        `json:"createDateTime,omitempty"`
		CreatorID                 *string                                                                                                     `json:"creatorId,omitempty"`
		CreatorUsername           *string                                                                                                     `json:"creatorUsername,omitempty"`
		DefaultWhitelistedDomains AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListDefaultWhitelistedDomainsSlice `json:"defaultWhitelistedDomains,omitempty"`
		Description               *string                                                                                                     `json:"description,omitempty"`
		DomainID                  *string                                                                                                     `json:"domainId,omitempty"`
		FriendlyName              *string                                                                                                     `json:"friendlyName,omitempty"`
		ID                        *string                                                                                                     `json:"id,omitempty"`
		Mappings                  AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListMappingsSlice                  `json:"mappings,omitempty"`
		ModifiedDateTime          *int                                                                                                        `json:"modifiedDateTime,omitempty"`
		ModifierID                *string                                                                                                     `json:"modifierId,omitempty"`
		ModifierUsername          *string                                                                                                     `json:"modifierUsername,omitempty"`
		MvnoID                    *string                                                                                                     `json:"mvnoId,omitempty"`
		Name                      *string                                                                                                     `json:"name,omitempty"`
		Protocol                  *string                                                                                                     `json:"protocol,omitempty"`
		Type                      *string                                                                                                     `json:"type,omitempty"`
		WhitelistedDomains        AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListWhitelistedDomainsSlice        `json:"whitelistedDomains,omitempty"`
		WillCollectEmail          *bool                                                                                                       `json:"willCollectEmail,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200Response struct {
		Extra      *json.RawMessage                                                                   `json:"extra,omitempty"`
		FirstIndex *int                                                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                              `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                               `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListOauthAuthenticationServiceGet: Use this API command to retrieve a list of OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListOauthAuthenticationServiceGet(ctx *UserContext) (*http.Response, *AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/oauth")
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListOauthAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceCreateOauthAuthenticationServicePostRequestMappingsSlice []*AuthenticationServiceCreateOauthAuthenticationServicePostRequestMappings

	AuthenticationServiceCreateOauthAuthenticationServicePostRequestMappingsUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceCreateOauthAuthenticationServicePostRequestMappings struct {
		GroupAttr *string                                                                           `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceCreateOauthAuthenticationServicePostRequestMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceCreateOauthAuthenticationServicePostRequestWhitelistedDomainsSlice []*string

	AuthenticationServiceCreateOauthAuthenticationServicePostRequest struct {
		ApplicationID      *string                                                                                 `json:"applicationId,omitempty"`
		ApplicationSecret  *string                                                                                 `json:"applicationSecret,omitempty"`
		Description        *string                                                                                 `json:"description,omitempty"`
		DomainID           *string                                                                                 `json:"domainId,omitempty"`
		FriendlyName       *string                                                                                 `json:"friendlyName,omitempty"`
		ID                 *string                                                                                 `json:"id,omitempty"`
		Mappings           AuthenticationServiceCreateOauthAuthenticationServicePostRequestMappingsSlice           `json:"mappings,omitempty"`
		Name               *string                                                                                 `json:"name,omitempty"`
		Protocol           *string                                                                                 `json:"protocol,omitempty"`
		Type               *string                                                                                 `json:"type,omitempty"`
		WhitelistedDomains AuthenticationServiceCreateOauthAuthenticationServicePostRequestWhitelistedDomainsSlice `json:"whitelistedDomains,omitempty"`
		WillCollectEmail   *bool                                                                                   `json:"willCollectEmail,omitempty"`
	}

	AuthenticationServiceCreateOauthAuthenticationServicePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AuthenticationServiceCreateOauthAuthenticationServicePost: Use this API command to create a new OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceCreateOauthAuthenticationServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceCreateOauthAuthenticationServicePost201Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceCreateOauthAuthenticationServicePost(ctx *UserContext, requestBody *AuthenticationServiceCreateOauthAuthenticationServicePostRequest) (*http.Response, *AuthenticationServiceCreateOauthAuthenticationServicePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/oauth")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceCreateOauthAuthenticationServicePost201Response{}
	httpResponse, _, err := s.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestAttributesSlice []*string

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraFilters

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFiltersSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFilters

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                        `json:"type,omitempty"`
		Value  *string                                                                                                        `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                            `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                          `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                            `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                            `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                          `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                            `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                            `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                            `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                            `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                            `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                            `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                          `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                            `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                          `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                          `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                          `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                          `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                          `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                            `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                            `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                            `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                          `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                          `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                          `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                          `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                          `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                          `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                          `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                          `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                          `json:"localUser_userSource,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequest struct {
		Attributes      AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                                   `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                     `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                      `json:"limit,omitempty"`
		Options         *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                      `json:"page,omitempty"`
		Query           *string                                                                                                   `json:"query,omitempty"`
		SortInfo        *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                      `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseList

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListDefaultWhitelistedDomainsSlice []*string

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappings

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                                          `json:"id,omitempty"`
		Name               *string                                                                                                                          `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappings struct {
		GroupAttr *string                                                                                                        `json:"groupAttr,omitempty"`
		ID        *string                                                                                                        `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListWhitelistedDomainsSlice []*string

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseList struct {
		ApplicationID             *string                                                                                                                     `json:"applicationId,omitempty"`
		ApplicationSecret         *string                                                                                                                     `json:"applicationSecret,omitempty"`
		CreateDateTime            *int                                                                                                                        `json:"createDateTime,omitempty"`
		CreatorID                 *string                                                                                                                     `json:"creatorId,omitempty"`
		CreatorUsername           *string                                                                                                                     `json:"creatorUsername,omitempty"`
		DefaultWhitelistedDomains AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListDefaultWhitelistedDomainsSlice `json:"defaultWhitelistedDomains,omitempty"`
		Description               *string                                                                                                                     `json:"description,omitempty"`
		DomainID                  *string                                                                                                                     `json:"domainId,omitempty"`
		FriendlyName              *string                                                                                                                     `json:"friendlyName,omitempty"`
		ID                        *string                                                                                                                     `json:"id,omitempty"`
		Mappings                  AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice                  `json:"mappings,omitempty"`
		ModifiedDateTime          *int                                                                                                                        `json:"modifiedDateTime,omitempty"`
		ModifierID                *string                                                                                                                     `json:"modifierId,omitempty"`
		ModifierUsername          *string                                                                                                                     `json:"modifierUsername,omitempty"`
		MvnoID                    *string                                                                                                                     `json:"mvnoId,omitempty"`
		Name                      *string                                                                                                                     `json:"name,omitempty"`
		Protocol                  *string                                                                                                                     `json:"protocol,omitempty"`
		Type                      *string                                                                                                                     `json:"type,omitempty"`
		WhitelistedDomains        AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListWhitelistedDomainsSlice        `json:"whitelistedDomains,omitempty"`
		WillCollectEmail          *bool                                                                                                                       `json:"willCollectEmail,omitempty"`
	}

	AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                                   `json:"extra,omitempty"`
		FirstIndex *int                                                                                               `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                              `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                               `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost: Use this API command to retrieve a list of oAuth Authentication services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost(ctx *UserContext, requestBody *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPostRequest) (*http.Response, *AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/oauth/query")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListOauthAuthenticationServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AuthenticationServiceDeleteOauthAuthenticationServiceDelete: Use this API command to delete an OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): OAUTH Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDeleteOauthAuthenticationServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth/oauth/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseDefaultWhitelistedDomainsSlice []*string

	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappingsSlice []*AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappings

	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappingsUserRole struct {
		ID                 *string                                                                                                  `json:"id,omitempty"`
		Name               *string                                                                                                  `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappings struct {
		GroupAttr *string                                                                                `json:"groupAttr,omitempty"`
		ID        *string                                                                                `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseWhitelistedDomainsSlice []*string

	AuthenticationServiceRetrieveOauthAuthenticationServiceGet200Response struct {
		ApplicationID             *string                                                                                             `json:"applicationId,omitempty"`
		ApplicationSecret         *string                                                                                             `json:"applicationSecret,omitempty"`
		CreateDateTime            *int                                                                                                `json:"createDateTime,omitempty"`
		CreatorID                 *string                                                                                             `json:"creatorId,omitempty"`
		CreatorUsername           *string                                                                                             `json:"creatorUsername,omitempty"`
		DefaultWhitelistedDomains AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseDefaultWhitelistedDomainsSlice `json:"defaultWhitelistedDomains,omitempty"`
		Description               *string                                                                                             `json:"description,omitempty"`
		DomainID                  *string                                                                                             `json:"domainId,omitempty"`
		FriendlyName              *string                                                                                             `json:"friendlyName,omitempty"`
		ID                        *string                                                                                             `json:"id,omitempty"`
		Mappings                  AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseMappingsSlice                  `json:"mappings,omitempty"`
		ModifiedDateTime          *int                                                                                                `json:"modifiedDateTime,omitempty"`
		ModifierID                *string                                                                                             `json:"modifierId,omitempty"`
		ModifierUsername          *string                                                                                             `json:"modifierUsername,omitempty"`
		MvnoID                    *string                                                                                             `json:"mvnoId,omitempty"`
		Name                      *string                                                                                             `json:"name,omitempty"`
		Protocol                  *string                                                                                             `json:"protocol,omitempty"`
		Type                      *string                                                                                             `json:"type,omitempty"`
		WhitelistedDomains        AuthenticationServiceRetrieveOauthAuthenticationServiceGet200ResponseWhitelistedDomainsSlice        `json:"whitelistedDomains,omitempty"`
		WillCollectEmail          *bool                                                                                               `json:"willCollectEmail,omitempty"`
	}
)

// AuthenticationServiceRetrieveOauthAuthenticationServiceGet: Use this API command to retrieve a OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): OAUTH Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveOauthAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveOauthAuthenticationServiceGet(ctx *UserContext, id string) (*http.Response, *AuthenticationServiceRetrieveOauthAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/oauth/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AuthenticationServiceRetrieveOauthAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceModifyOauthAuthenticationServicePatchRequest struct {
		ApplicationID     *string `json:"applicationId,omitempty"`
		ApplicationSecret *string `json:"applicationSecret,omitempty"`
		Description       *string `json:"description,omitempty"`
		DomainID          *string `json:"domainId,omitempty"`
		FriendlyName      *string `json:"friendlyName,omitempty"`
		ID                *string `json:"id,omitempty"`
		Name              *string `json:"name,omitempty"`
		Protocol          *string `json:"protocol,omitempty"`
		Type              *string `json:"type,omitempty"`
		WillCollectEmail  *bool   `json:"willCollectEmail,omitempty"`
	}
)

// AuthenticationServiceModifyOauthAuthenticationServicePatch: Use this API command to modify the basic information of an OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): OAUTH Service ID
// - requestBody: *AuthenticationServiceModifyOauthAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyOauthAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyOauthAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/oauth/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequestSlice []*AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequest

	AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequestUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequest struct {
		GroupAttr *string                                                                          `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequestUserRole `json:"userRole,omitempty"`
	}
)

// AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatch: Use this API command to modify user traffic profile mapping of an OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): OAUTH Service ID
// - requestBody: *AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatch(ctx *UserContext, id string, requestBody AuthenticationServiceModifyUserTrafficProfileMappingOfOauthPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/oauth/{id}/mappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyWhitelistedDomainsOfOauthPatchRequest []*string
)

// AuthenticationServiceModifyWhitelistedDomainsOfOauthPatch: Use this API command to modify whitelisted domains of an OAuth authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): OAUTH Service ID
// - requestBody: *AuthenticationServiceModifyWhitelistedDomainsOfOauthPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyWhitelistedDomainsOfOauthPatch(ctx *UserContext, id string, requestBody AuthenticationServiceModifyWhitelistedDomainsOfOauthPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/oauth/{id}/whitelistedDomains")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestAttributesSlice []*string

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice []*AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraFilters

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFiltersSlice []*AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFilters

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                             `json:"type,omitempty"`
		Value  *string                                                                                                             `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                                 `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                               `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                                 `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                                 `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                               `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                                 `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                                 `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                                 `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                                 `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                                 `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                                 `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                               `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                                 `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                               `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                               `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                               `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                               `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                               `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                                 `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                                 `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                                 `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                               `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                               `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                               `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                               `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                               `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                               `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                               `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                               `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                               `json:"localUser_userSource,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequest struct {
		Attributes      AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                                        `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                          `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                           `json:"limit,omitempty"`
		Options         *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                           `json:"page,omitempty"`
		Query           *string                                                                                                        `json:"query,omitempty"`
		SortInfo        *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                           `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListSlice []*AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseList

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappings

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                                               `json:"id,omitempty"`
		Name               *string                                                                                                                               `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappings struct {
		GroupAttr *string                                                                                                             `json:"groupAttr,omitempty"`
		ID        *string                                                                                                             `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseList struct {
		CreateDateTime   *int                                                                                                            `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                                         `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                                         `json:"creatorUsername,omitempty"`
		Description      *string                                                                                                         `json:"description,omitempty"`
		DomainID         *string                                                                                                         `json:"domainId,omitempty"`
		FriendlyName     *string                                                                                                         `json:"friendlyName,omitempty"`
		ID               *string                                                                                                         `json:"id,omitempty"`
		Mappings         AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice `json:"mappings,omitempty"`
		ModifiedDateTime *int                                                                                                            `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                                         `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                                         `json:"modifierUsername,omitempty"`
		MvnoID           *string                                                                                                         `json:"mvnoId,omitempty"`
		Name             *string                                                                                                         `json:"name,omitempty"`
		Protocol         *string                                                                                                         `json:"protocol,omitempty"`
		Type             *string                                                                                                         `json:"type,omitempty"`
	}

	AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                                        `json:"extra,omitempty"`
		FirstIndex *int                                                                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                                   `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                                    `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost: Use this API command to retrieve a list of Authentication services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost(ctx *UserContext, requestBody *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPostRequest) (*http.Response, *AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/query")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListAllTypesOfAuthenticationServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseList

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappings

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                           `json:"id,omitempty"`
		Name               *string                                                                                                           `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappings struct {
		GroupAttr *string                                                                                         `json:"groupAttr,omitempty"`
		ID        *string                                                                                         `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseList struct {
		CreateDateTime          *int                                                                                             `json:"createDateTime,omitempty"`
		CreatorID               *string                                                                                          `json:"creatorId,omitempty"`
		CreatorUsername         *string                                                                                          `json:"creatorUsername,omitempty"`
		Description             *string                                                                                          `json:"description,omitempty"`
		DomainID                *string                                                                                          `json:"domainId,omitempty"`
		FriendlyName            *string                                                                                          `json:"friendlyName,omitempty"`
		HealthCheckPolicy       *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                      *string                                                                                          `json:"id,omitempty"`
		LocationDeliveryEnabled *bool                                                                                            `json:"locationDeliveryEnabled,omitempty"`
		Mappings                AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListMappingsSlice      `json:"mappings,omitempty"`
		ModifiedDateTime        *int                                                                                             `json:"modifiedDateTime,omitempty"`
		ModifierID              *string                                                                                          `json:"modifierId,omitempty"`
		ModifierUsername        *string                                                                                          `json:"modifierUsername,omitempty"`
		MvnoID                  *string                                                                                          `json:"mvnoId,omitempty"`
		Name                    *string                                                                                          `json:"name,omitempty"`
		Primary                 *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListPrimary           `json:"primary,omitempty"`
		Protocol                *string                                                                                          `json:"protocol,omitempty"`
		RateLimiting            *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary               *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListSecondary         `json:"secondary,omitempty"`
		Type                    *string                                                                                          `json:"type,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200Response struct {
		Extra      *json.RawMessage                                                                    `json:"extra,omitempty"`
		FirstIndex *int                                                                                `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                               `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet: Use this API command to retrieve a list of RADIUS authentication services.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet(ctx *UserContext) (*http.Response, *AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/radius")
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListRadiusAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestMappingsSlice []*AuthenticationServiceCreateRadiusAuthenticationServicePostRequestMappings

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestMappingsUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestMappings struct {
		GroupAttr *string                                                                            `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceCreateRadiusAuthenticationServicePostRequestMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequestSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePostRequest struct {
		Description             *string                                                                             `json:"description,omitempty"`
		DomainID                *string                                                                             `json:"domainId,omitempty"`
		FriendlyName            *string                                                                             `json:"friendlyName,omitempty"`
		HealthCheckPolicy       *AuthenticationServiceCreateRadiusAuthenticationServicePostRequestHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                      *string                                                                             `json:"id,omitempty"`
		LocationDeliveryEnabled *bool                                                                               `json:"locationDeliveryEnabled,omitempty"`
		Mappings                AuthenticationServiceCreateRadiusAuthenticationServicePostRequestMappingsSlice      `json:"mappings,omitempty"`
		Name                    *string                                                                             `json:"name,omitempty"`
		Primary                 *AuthenticationServiceCreateRadiusAuthenticationServicePostRequestPrimary           `json:"primary,omitempty"`
		RateLimiting            *AuthenticationServiceCreateRadiusAuthenticationServicePostRequestRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary               *AuthenticationServiceCreateRadiusAuthenticationServicePostRequestSecondary         `json:"secondary,omitempty"`
		Type                    *string                                                                             `json:"type,omitempty"`
	}

	AuthenticationServiceCreateRadiusAuthenticationServicePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// AuthenticationServiceCreateRadiusAuthenticationServicePost: Use this API command to create a new RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceCreateRadiusAuthenticationServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceCreateRadiusAuthenticationServicePost201Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceCreateRadiusAuthenticationServicePost(ctx *UserContext, requestBody *AuthenticationServiceCreateRadiusAuthenticationServicePostRequest) (*http.Response, *AuthenticationServiceCreateRadiusAuthenticationServicePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/radius")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceCreateRadiusAuthenticationServicePost201Response{}
	httpResponse, _, err := s.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestAttributesSlice []*string

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraFilters

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFiltersSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFilters

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice []*string

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFullTextSearch struct {
		Fields AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                                                         `json:"type,omitempty"`
		Value  *string                                                                                                         `json:"value,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestOptions struct {
		INCLUDERBACMETADATA           *bool                                                                                                             `json:"INCLUDE_RBAC_METADATA,omitempty"`
		TENANTID                      *string                                                                                                           `json:"TENANT_ID,omitempty"`
		AcctIncludeNa                 *bool                                                                                                             `json:"acct_includeNa,omitempty"`
		AcctTestableOnly              *bool                                                                                                             `json:"acct_testableOnly,omitempty"`
		AcctType                      *string                                                                                                           `json:"acct_type,omitempty"`
		AuthHostedAaaSupportedEnabled *bool                                                                                                             `json:"auth_hostedAaaSupportedEnabled,omitempty"`
		AuthIncludeAdGlobal           *bool                                                                                                             `json:"auth_includeAdGlobal,omitempty"`
		AuthIncludeGuest              *bool                                                                                                             `json:"auth_includeGuest,omitempty"`
		AuthIncludeLocalDb            *bool                                                                                                             `json:"auth_includeLocalDb,omitempty"`
		AuthIncludeNa                 *bool                                                                                                             `json:"auth_includeNa,omitempty"`
		AuthPlmnIdentifierEnabled     *bool                                                                                                             `json:"auth_plmnIdentifierEnabled,omitempty"`
		AuthRealmType                 *string                                                                                                           `json:"auth_realmType,omitempty"`
		AuthTestableOnly              *bool                                                                                                             `json:"auth_testableOnly,omitempty"`
		AuthType                      *string                                                                                                           `json:"auth_type,omitempty"`
		ForwardingType                *string                                                                                                           `json:"forwarding_type,omitempty"`
		GlobalFilterID                *string                                                                                                           `json:"globalFilterId,omitempty"`
		GuestPassDisplayName          *string                                                                                                           `json:"guestPass_displayName,omitempty"`
		GuestPassExpiration           *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                                                           `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                                                             `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                                                             `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                                                             `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
		LocalUserDisplayName          *string                                                                                                           `json:"localUser_displayName,omitempty"`
		LocalUserFirstName            *string                                                                                                           `json:"localUser_firstName,omitempty"`
		LocalUserLastName             *string                                                                                                           `json:"localUser_lastName,omitempty"`
		LocalUserMailAddress          *string                                                                                                           `json:"localUser_mailAddress,omitempty"`
		LocalUserPrimaryPhoneNumber   *string                                                                                                           `json:"localUser_primaryPhoneNumber,omitempty"`
		LocalUserStatus               *string                                                                                                           `json:"localUser_status,omitempty"`
		LocalUserSubscriberType       *string                                                                                                           `json:"localUser_subscriberType,omitempty"`
		LocalUserUserName             *string                                                                                                           `json:"localUser_userName,omitempty"`
		LocalUserUserSource           *string                                                                                                           `json:"localUser_userSource,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequest struct {
		Attributes      AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                                                                    `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                                                      `json:"expandDomains,omitempty"`
		ExtraFilters    AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *int                                                                                                       `json:"limit,omitempty"`
		Options         *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                                                       `json:"page,omitempty"`
		Query           *string                                                                                                    `json:"query,omitempty"`
		SortInfo        *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *int                                                                                                       `json:"start,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseList

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice []*AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappings

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole struct {
		ID                 *string                                                                                                                           `json:"id,omitempty"`
		Name               *string                                                                                                                           `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappings struct {
		GroupAttr *string                                                                                                         `json:"groupAttr,omitempty"`
		ID        *string                                                                                                         `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListPrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseList struct {
		CreateDateTime          *int                                                                                                             `json:"createDateTime,omitempty"`
		CreatorID               *string                                                                                                          `json:"creatorId,omitempty"`
		CreatorUsername         *string                                                                                                          `json:"creatorUsername,omitempty"`
		Description             *string                                                                                                          `json:"description,omitempty"`
		DomainID                *string                                                                                                          `json:"domainId,omitempty"`
		FriendlyName            *string                                                                                                          `json:"friendlyName,omitempty"`
		HealthCheckPolicy       *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                      *string                                                                                                          `json:"id,omitempty"`
		LocationDeliveryEnabled *bool                                                                                                            `json:"locationDeliveryEnabled,omitempty"`
		Mappings                AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListMappingsSlice      `json:"mappings,omitempty"`
		ModifiedDateTime        *int                                                                                                             `json:"modifiedDateTime,omitempty"`
		ModifierID              *string                                                                                                          `json:"modifierId,omitempty"`
		ModifierUsername        *string                                                                                                          `json:"modifierUsername,omitempty"`
		MvnoID                  *string                                                                                                          `json:"mvnoId,omitempty"`
		Name                    *string                                                                                                          `json:"name,omitempty"`
		Primary                 *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListPrimary           `json:"primary,omitempty"`
		Protocol                *string                                                                                                          `json:"protocol,omitempty"`
		RateLimiting            *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary               *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListSecondary         `json:"secondary,omitempty"`
		Type                    *string                                                                                                          `json:"type,omitempty"`
	}

	AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200Response struct {
		Extra      *json.RawMessage                                                                                    `json:"extra,omitempty"`
		FirstIndex *int                                                                                                `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                                               `json:"hasMore,omitempty"`
		List       AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                                                `json:"totalCount,omitempty"`
	}
)

// AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost: Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost(ctx *UserContext, requestBody *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPostRequest) (*http.Response, *AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/radius/query")
	request.body = requestBody
	request.authenticated = true
	out := &AuthenticationServiceRetrieveListRadiusAuthenticationServiceByQueryCritariaPost200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// AuthenticationServiceDeleteRadiusAuthenticationServiceDelete: Use this API command to delete a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDeleteRadiusAuthenticationServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth/radius/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseHealthCheckPolicy struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappingsSlice []*AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappings

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappingsUserRole struct {
		ID                 *string                                                                                                   `json:"id,omitempty"`
		Name               *string                                                                                                   `json:"name,omitempty"`
		UserTrafficProfile *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappingsUserRoleUserTrafficProfile `json:"userTrafficProfile,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappings struct {
		GroupAttr *string                                                                                 `json:"groupAttr,omitempty"`
		ID        *string                                                                                 `json:"id,omitempty"`
		UserRole  *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappingsUserRole `json:"userRole,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponsePrimary struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseRateLimiting struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseSecondary struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}

	AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200Response struct {
		CreateDateTime          *int                                                                                     `json:"createDateTime,omitempty"`
		CreatorID               *string                                                                                  `json:"creatorId,omitempty"`
		CreatorUsername         *string                                                                                  `json:"creatorUsername,omitempty"`
		Description             *string                                                                                  `json:"description,omitempty"`
		DomainID                *string                                                                                  `json:"domainId,omitempty"`
		FriendlyName            *string                                                                                  `json:"friendlyName,omitempty"`
		HealthCheckPolicy       *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
		ID                      *string                                                                                  `json:"id,omitempty"`
		LocationDeliveryEnabled *bool                                                                                    `json:"locationDeliveryEnabled,omitempty"`
		Mappings                AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseMappingsSlice      `json:"mappings,omitempty"`
		ModifiedDateTime        *int                                                                                     `json:"modifiedDateTime,omitempty"`
		ModifierID              *string                                                                                  `json:"modifierId,omitempty"`
		ModifierUsername        *string                                                                                  `json:"modifierUsername,omitempty"`
		MvnoID                  *string                                                                                  `json:"mvnoId,omitempty"`
		Name                    *string                                                                                  `json:"name,omitempty"`
		Primary                 *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponsePrimary           `json:"primary,omitempty"`
		Protocol                *string                                                                                  `json:"protocol,omitempty"`
		RateLimiting            *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseRateLimiting      `json:"rateLimiting,omitempty"`
		Secondary               *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200ResponseSecondary         `json:"secondary,omitempty"`
		Type                    *string                                                                                  `json:"type,omitempty"`
	}
)

// AuthenticationServiceRetrieveRadiusAuthenticationServiceGet: Use this API command to retrieve a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200Response
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceRetrieveRadiusAuthenticationServiceGet(ctx *UserContext, id string) (*http.Response, *AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "GET", "/v5_0/services/auth/radius/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &AuthenticationServiceRetrieveRadiusAuthenticationServiceGet200Response{}
	httpResponse, _, err := s.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	AuthenticationServiceModifyRadiusAuthenticationServicePatchRequest struct {
		Description             *string `json:"description,omitempty"`
		DomainID                *string `json:"domainId,omitempty"`
		FriendlyName            *string `json:"friendlyName,omitempty"`
		ID                      *string `json:"id,omitempty"`
		LocationDeliveryEnabled *bool   `json:"locationDeliveryEnabled,omitempty"`
		Name                    *string `json:"name,omitempty"`
		Type                    *string `json:"type,omitempty"`
	}
)

// AuthenticationServiceModifyRadiusAuthenticationServicePatch: Use this API command to modify the basic information of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
// - requestBody: *AuthenticationServiceModifyRadiusAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyRadiusAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyRadiusAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/radius/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyHealthCheckPolicyOfAuthenticationServicePatchRequest struct {
		ResponseFail   *bool    `json:"responseFail,omitempty"`
		ResponseWindow *float64 `json:"responseWindow,omitempty"`
		ReviveInterval *float64 `json:"reviveInterval,omitempty"`
		ZombiePeriod   *float64 `json:"zombiePeriod,omitempty"`
	}
)

// AuthenticationServiceModifyHealthCheckPolicyOfAuthenticationServicePatch: Use this API command to modify health check policy of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
// - requestBody: *AuthenticationServiceModifyHealthCheckPolicyOfAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyHealthCheckPolicyOfAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyHealthCheckPolicyOfAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/radius/{id}/healthCheckPolicy")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyUserTrafficProfileMappingPatchRequestSlice []*AuthenticationServiceModifyUserTrafficProfileMappingPatchRequest

	AuthenticationServiceModifyUserTrafficProfileMappingPatchRequestUserRole struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AuthenticationServiceModifyUserTrafficProfileMappingPatchRequest struct {
		GroupAttr *string                                                                   `json:"groupAttr,omitempty"`
		UserRole  *AuthenticationServiceModifyUserTrafficProfileMappingPatchRequestUserRole `json:"userRole,omitempty"`
	}
)

// AuthenticationServiceModifyUserTrafficProfileMappingPatch: Use this API command to modify user traffic profile mapping of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
// - requestBody: *AuthenticationServiceModifyUserTrafficProfileMappingPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyUserTrafficProfileMappingPatch(ctx *UserContext, id string, requestBody AuthenticationServiceModifyUserTrafficProfileMappingPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/radius/{id}/mappings")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyPrimaryRadiusServerOfAuthenticationServicePatchRequest struct {
		IP           *string `json:"ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
		SharedSecret *string `json:"sharedSecret,omitempty"`
	}
)

// AuthenticationServiceModifyPrimaryRadiusServerOfAuthenticationServicePatch: Use this API command to modify primary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
// - requestBody: *AuthenticationServiceModifyPrimaryRadiusServerOfAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyPrimaryRadiusServerOfAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyPrimaryRadiusServerOfAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/radius/{id}/primary")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifyRateLimitingOfAuthenticationServicePatchRequest struct {
		MaxOutstandingRequestsPerServer *float64 `json:"maxOutstandingRequestsPerServer,omitempty"`
		SanityTimer                     *float64 `json:"sanityTimer,omitempty"`
		Threshold                       *float64 `json:"threshold,omitempty"`
	}
)

// AuthenticationServiceModifyRateLimitingOfAuthenticationServicePatch: Use this API command to modify rate limiting of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
// - requestBody: *AuthenticationServiceModifyRateLimitingOfAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifyRateLimitingOfAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifyRateLimitingOfAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/radius/{id}/rateLimiting")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

// AuthenticationServiceDisableSecondaryRadiusServerOfAuthenticationServiceDelete: Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDisableSecondaryRadiusServerOfAuthenticationServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth/radius/{id}/secondary")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceModifySecondaryRadiusServerOfAuthenticationServicePatchRequest struct {
		AutoFallbackDisable *bool   `json:"autoFallbackDisable,omitempty"`
		IP                  *string `json:"ip,omitempty"`
		Port                *int    `json:"port,omitempty"`
		SharedSecret        *string `json:"sharedSecret,omitempty"`
	}
)

// AuthenticationServiceModifySecondaryRadiusServerOfAuthenticationServicePatch: Use this API command to modify secondary RADIUS server of a RADIUS authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): RADIUS Service ID
// - requestBody: *AuthenticationServiceModifySecondaryRadiusServerOfAuthenticationServicePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceModifySecondaryRadiusServerOfAuthenticationServicePatch(ctx *UserContext, id string, requestBody *AuthenticationServiceModifySecondaryRadiusServerOfAuthenticationServicePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "PATCH", "/v5_0/services/auth/radius/{id}/secondary")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

type (
	AuthenticationServiceTestSpecificAuthenticationServicePostRequestLoginRequest struct {
		Password          *string `json:"password,omitempty"`
		Protocol          *string `json:"protocol,omitempty"`
		TimeZoneUtcOffset *string `json:"timeZoneUtcOffset,omitempty"`
		UserName          *string `json:"userName,omitempty"`
	}

	AuthenticationServiceTestSpecificAuthenticationServicePostRequest struct {
		ID           *string                                                                        `json:"id,omitempty"`
		LoginRequest *AuthenticationServiceTestSpecificAuthenticationServicePostRequestLoginRequest `json:"loginRequest,omitempty"`
	}
)

// AuthenticationServiceTestSpecificAuthenticationServicePost: Use this API command to test an authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *AuthenticationServiceTestSpecificAuthenticationServicePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceTestSpecificAuthenticationServicePost(ctx *UserContext, id string, requestBody *AuthenticationServiceTestSpecificAuthenticationServicePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "POST", "/v5_0/services/auth/test/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}

// AuthenticationServiceDeleteAuthenticationServiceDelete: Use this API command to delete an authentication service.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (s *ServicesAPI) AuthenticationServiceDeleteAuthenticationServiceDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := s.client.newRequest(ctx, "DELETE", "/v5_0/services/auth/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return s.client.doRequest(request, 204, nil)
}
