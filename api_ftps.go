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

type FTPSAPI struct {
	client *Client
}
type (
	FtpserversettingsRemoveFtpServersDeleteRequestIDListSlice []*string

	FtpserversettingsRemoveFtpServersDeleteRequest struct {
		IDList FtpserversettingsRemoveFtpServersDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// FtpserversettingsRemoveFtpServersDelete: Remove FTP servers
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *FtpserversettingsRemoveFtpServersDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsRemoveFtpServersDelete(ctx *UserContext, requestBody *FtpserversettingsRemoveFtpServersDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := f.client.newRequest(ctx, "DELETE", "/v5_0/ftps")
	request.body = requestBody
	request.authenticated = true
	return f.client.doRequest(request, 200, nil)
}

type (
	FtpserversettingsAddFtpServerPostRequest struct {
		CreateDatetime     *float64 `json:"createDatetime,omitempty"`
		CreatorUUID        *string  `json:"creatorUUID,omitempty"`
		DomainID           *string  `json:"domainId,omitempty"`
		FtpHost            *string  `json:"ftpHost,omitempty"`
		FtpName            *string  `json:"ftpName,omitempty"`
		FtpPassword        *string  `json:"ftpPassword,omitempty"`
		FtpPort            *float64 `json:"ftpPort,omitempty"`
		FtpProtocol        *string  `json:"ftpProtocol,omitempty"`
		FtpRemoteDirectory *string  `json:"ftpRemoteDirectory,omitempty"`
		FtpUserName        *string  `json:"ftpUserName,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		LastModifiedBy     *string  `json:"lastModifiedBy,omitempty"`
		LastModifiedOn     *float64 `json:"lastModifiedOn,omitempty"`
		TenantID           *string  `json:"tenantId,omitempty"`
	}
)

// FtpserversettingsAddFtpServerPost: Add FTP server
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *FtpserversettingsAddFtpServerPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsAddFtpServerPost(ctx *UserContext, requestBody *FtpserversettingsAddFtpServerPostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := f.client.newRequest(ctx, "POST", "/v5_0/ftps")
	request.body = requestBody
	request.authenticated = true
	return f.client.doRequest(request, 201, nil)
}

type (
	FtpserversettingsRetrieveFtpServerListPostRequestAttributesSlice []*string

	FtpserversettingsRetrieveFtpServerListPostRequestExtraFiltersSlice []*FtpserversettingsRetrieveFtpServerListPostRequestExtraFilters

	FtpserversettingsRetrieveFtpServerListPostRequestExtraFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestExtraNotFiltersSlice []*FtpserversettingsRetrieveFtpServerListPostRequestExtraNotFilters

	FtpserversettingsRetrieveFtpServerListPostRequestExtraNotFilters struct {
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestExtraTimeRange struct {
		End      *float64 `json:"end,omitempty"`
		Field    *string  `json:"field,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestFiltersSlice []*FtpserversettingsRetrieveFtpServerListPostRequestFilters

	FtpserversettingsRetrieveFtpServerListPostRequestFilters struct {
		Operator *string `json:"operator,omitempty"`
		Type     *string `json:"type,omitempty"`
		Value    *string `json:"value,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestFullTextSearchFieldsSlice []*string

	FtpserversettingsRetrieveFtpServerListPostRequestFullTextSearch struct {
		Fields FtpserversettingsRetrieveFtpServerListPostRequestFullTextSearchFieldsSlice `json:"fields,omitempty"`
		Type   *string                                                                    `json:"type,omitempty"`
		Value  *string                                                                    `json:"value,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestOptionsGuestPassExpiration struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestOptionsLocalUserAuditTime struct {
		End      *float64 `json:"end,omitempty"`
		Interval *float64 `json:"interval,omitempty"`
		Start    *float64 `json:"start,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequestOptions struct {
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
		GuestPassExpiration           *FtpserversettingsRetrieveFtpServerListPostRequestOptionsGuestPassExpiration `json:"guestPass_expiration,omitempty"`
		GuestPassWlan                 *string                                                                      `json:"guestPass_wlan,omitempty"`
		InMap                         *bool                                                                        `json:"inMap,omitempty"`
		IncludeSharedResources        *bool                                                                        `json:"includeSharedResources,omitempty"`
		IncludeUsers                  *bool                                                                        `json:"includeUsers,omitempty"`
		LocalUserAuditTime            *FtpserversettingsRetrieveFtpServerListPostRequestOptionsLocalUserAuditTime  `json:"localUser_auditTime,omitempty"`
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

	FtpserversettingsRetrieveFtpServerListPostRequestSortInfo struct {
		Dir        *string `json:"dir,omitempty"`
		SortColumn *string `json:"sortColumn,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPostRequest struct {
		Attributes      FtpserversettingsRetrieveFtpServerListPostRequestAttributesSlice      `json:"attributes,omitempty"`
		Criteria        *string                                                               `json:"criteria,omitempty"`
		ExpandDomains   *bool                                                                 `json:"expandDomains,omitempty"`
		ExtraFilters    FtpserversettingsRetrieveFtpServerListPostRequestExtraFiltersSlice    `json:"extraFilters,omitempty"`
		ExtraNotFilters FtpserversettingsRetrieveFtpServerListPostRequestExtraNotFiltersSlice `json:"extraNotFilters,omitempty"`
		ExtraTimeRange  *FtpserversettingsRetrieveFtpServerListPostRequestExtraTimeRange      `json:"extraTimeRange,omitempty"`
		Filters         FtpserversettingsRetrieveFtpServerListPostRequestFiltersSlice         `json:"filters,omitempty"`
		FullTextSearch  *FtpserversettingsRetrieveFtpServerListPostRequestFullTextSearch      `json:"fullTextSearch,omitempty"`
		Limit           *float64                                                              `json:"limit,omitempty"`
		Options         *FtpserversettingsRetrieveFtpServerListPostRequestOptions             `json:"options,omitempty"`
		Page            *int                                                                  `json:"page,omitempty"`
		Query           *string                                                               `json:"query,omitempty"`
		SortInfo        *FtpserversettingsRetrieveFtpServerListPostRequestSortInfo            `json:"sortInfo,omitempty"`
		Start           *float64                                                              `json:"start,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPost200ResponseListSlice []*FtpserversettingsRetrieveFtpServerListPost200ResponseList

	FtpserversettingsRetrieveFtpServerListPost200ResponseList struct {
		CreateDatetime     *float64 `json:"createDatetime,omitempty"`
		CreatorUUID        *string  `json:"creatorUUID,omitempty"`
		DomainID           *string  `json:"domainId,omitempty"`
		FtpHost            *string  `json:"ftpHost,omitempty"`
		FtpName            *string  `json:"ftpName,omitempty"`
		FtpPassword        *string  `json:"ftpPassword,omitempty"`
		FtpPort            *float64 `json:"ftpPort,omitempty"`
		FtpProtocol        *string  `json:"ftpProtocol,omitempty"`
		FtpRemoteDirectory *string  `json:"ftpRemoteDirectory,omitempty"`
		FtpUserName        *string  `json:"ftpUserName,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		LastModifiedBy     *string  `json:"lastModifiedBy,omitempty"`
		LastModifiedOn     *float64 `json:"lastModifiedOn,omitempty"`
		TenantID           *string  `json:"tenantId,omitempty"`
	}

	FtpserversettingsRetrieveFtpServerListPost200Response struct {
		Extra      *json.RawMessage                                               `json:"extra,omitempty"`
		FirstIndex *int                                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                                          `json:"hasMore,omitempty"`
		List       FtpserversettingsRetrieveFtpServerListPost200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                           `json:"totalCount,omitempty"`
	}
)

// FtpserversettingsRetrieveFtpServerListPost: Retrieve information of all FTP server
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *FtpserversettingsRetrieveFtpServerListPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *FtpserversettingsRetrieveFtpServerListPost200Response
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsRetrieveFtpServerListPost(ctx *UserContext, requestBody *FtpserversettingsRetrieveFtpServerListPostRequest) (*http.Response, *FtpserversettingsRetrieveFtpServerListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := f.client.newRequest(ctx, "POST", "/v5_0/ftps/query")
	request.body = requestBody
	request.authenticated = true
	out := &FtpserversettingsRetrieveFtpServerListPost200Response{}
	httpResponse, _, err := f.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	FtpserversettingsTestFtpServerGetRequest struct {
		CreateDatetime     *float64 `json:"createDatetime,omitempty"`
		CreatorUUID        *string  `json:"creatorUUID,omitempty"`
		DomainID           *string  `json:"domainId,omitempty"`
		FtpHost            *string  `json:"ftpHost,omitempty"`
		FtpName            *string  `json:"ftpName,omitempty"`
		FtpPassword        *string  `json:"ftpPassword,omitempty"`
		FtpPort            *float64 `json:"ftpPort,omitempty"`
		FtpProtocol        *string  `json:"ftpProtocol,omitempty"`
		FtpRemoteDirectory *string  `json:"ftpRemoteDirectory,omitempty"`
		FtpUserName        *string  `json:"ftpUserName,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		LastModifiedBy     *string  `json:"lastModifiedBy,omitempty"`
		LastModifiedOn     *float64 `json:"lastModifiedOn,omitempty"`
		TenantID           *string  `json:"tenantId,omitempty"`
	}

	FtpserversettingsTestFtpServerGet200Response struct {
		Data    *bool            `json:"data,omitempty"`
		Error   *string          `json:"error,omitempty"`
		Extra   *json.RawMessage `json:"extra,omitempty"`
		Success *bool            `json:"success,omitempty"`
	}
)

// FtpserversettingsTestFtpServerGet: Test ftp server of specific FTP server settings
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *FtpserversettingsTestFtpServerGetRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *FtpserversettingsTestFtpServerGet200Response
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsTestFtpServerGet(ctx *UserContext, requestBody *FtpserversettingsTestFtpServerGetRequest) (*http.Response, *FtpserversettingsTestFtpServerGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := f.client.newRequest(ctx, "GET", "/v5_0/ftps/test")
	request.body = requestBody
	request.authenticated = true
	out := &FtpserversettingsTestFtpServerGet200Response{}
	httpResponse, _, err := f.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	FtpserversettingsTestFtpServerGet1200Response struct {
		Data    *bool            `json:"data,omitempty"`
		Error   *string          `json:"error,omitempty"`
		Extra   *json.RawMessage `json:"extra,omitempty"`
		Success *bool            `json:"success,omitempty"`
	}
)

// FtpserversettingsTestFtpServerGet1: Test ftp server of specific FTP server id
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - ftpId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *FtpserversettingsTestFtpServerGet1200Response
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsTestFtpServerGet1(ctx *UserContext, ftpId string) (*http.Response, *FtpserversettingsTestFtpServerGet1200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(ftpId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"ftpId\" failed validation check: %s", err)
	}
	request := f.client.newRequest(ctx, "GET", "/v5_0/ftps/test/{ftpId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"ftpId": ftpId,
	}
	out := &FtpserversettingsTestFtpServerGet1200Response{}
	httpResponse, _, err := f.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// FtpserversettingsRemoveFtpServerDelete: Remove FTP server
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - ftpId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsRemoveFtpServerDelete(ctx *UserContext, ftpId string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(ftpId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"ftpId\" failed validation check: %s", err)
	}
	request := f.client.newRequest(ctx, "DELETE", "/v5_0/ftps/{ftpId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"ftpId": ftpId,
	}
	return f.client.doRequest(request, 204, nil)
}

type (
	FtpserversettingsRetrieveFtpServerInformationGet200Response struct {
		CreateDatetime     *float64 `json:"createDatetime,omitempty"`
		CreatorUUID        *string  `json:"creatorUUID,omitempty"`
		DomainID           *string  `json:"domainId,omitempty"`
		FtpHost            *string  `json:"ftpHost,omitempty"`
		FtpName            *string  `json:"ftpName,omitempty"`
		FtpPassword        *string  `json:"ftpPassword,omitempty"`
		FtpPort            *float64 `json:"ftpPort,omitempty"`
		FtpProtocol        *string  `json:"ftpProtocol,omitempty"`
		FtpRemoteDirectory *string  `json:"ftpRemoteDirectory,omitempty"`
		FtpUserName        *string  `json:"ftpUserName,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		LastModifiedBy     *string  `json:"lastModifiedBy,omitempty"`
		LastModifiedOn     *float64 `json:"lastModifiedOn,omitempty"`
		TenantID           *string  `json:"tenantId,omitempty"`
	}
)

// FtpserversettingsRetrieveFtpServerInformationGet: Retrieve information of specific FTP server
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - ftpId (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *FtpserversettingsRetrieveFtpServerInformationGet200Response
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsRetrieveFtpServerInformationGet(ctx *UserContext, ftpId string) (*http.Response, *FtpserversettingsRetrieveFtpServerInformationGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(ftpId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"ftpId\" failed validation check: %s", err)
	}
	request := f.client.newRequest(ctx, "GET", "/v5_0/ftps/{ftpId}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"ftpId": ftpId,
	}
	out := &FtpserversettingsRetrieveFtpServerInformationGet200Response{}
	httpResponse, _, err := f.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	FtpserversettingsUpdateFtpServerPatchRequest struct {
		CreateDatetime     *float64 `json:"createDatetime,omitempty"`
		CreatorUUID        *string  `json:"creatorUUID,omitempty"`
		DomainID           *string  `json:"domainId,omitempty"`
		FtpHost            *string  `json:"ftpHost,omitempty"`
		FtpName            *string  `json:"ftpName,omitempty"`
		FtpPassword        *string  `json:"ftpPassword,omitempty"`
		FtpPort            *float64 `json:"ftpPort,omitempty"`
		FtpProtocol        *string  `json:"ftpProtocol,omitempty"`
		FtpRemoteDirectory *string  `json:"ftpRemoteDirectory,omitempty"`
		FtpUserName        *string  `json:"ftpUserName,omitempty"`
		ID                 *string  `json:"id,omitempty"`
		LastModifiedBy     *string  `json:"lastModifiedBy,omitempty"`
		LastModifiedOn     *float64 `json:"lastModifiedOn,omitempty"`
		TenantID           *string  `json:"tenantId,omitempty"`
	}
)

// FtpserversettingsUpdateFtpServerPatch: Update FTP server settings
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - ftpId (string)
// - requestBody: *FtpserversettingsUpdateFtpServerPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (f *FTPSAPI) FtpserversettingsUpdateFtpServerPatch(ctx *UserContext, ftpId string, requestBody *FtpserversettingsUpdateFtpServerPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(ftpId)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"ftpId\" failed validation check: %s", err)
	}
	request := f.client.newRequest(ctx, "PATCH", "/v5_0/ftps/{ftpId}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"ftpId": ftpId,
	}
	return f.client.doRequest(request, 204, nil)
}
