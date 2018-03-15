package api

import (
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-15T14:33:32-0500
// API Version: v5

type ApplicationsAPI struct {
	client *Client
}
type (
	ApplicationLogAndStatusModifyLogLevelPatchRequest struct {
		ApplicationName *string `json:"applicationName,omitempty"`
		LogLevel        *string `json:"logLevel,omitempty"`
	}
)

// ApplicationLogAndStatusModifyLogLevelPatch: Use this API command to modify log level of specified application.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ApplicationLogAndStatusModifyLogLevelPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *ApplicationsAPI) ApplicationLogAndStatusModifyLogLevelPatch(ctx *UserContext, requestBody *ApplicationLogAndStatusModifyLogLevelPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "PATCH", "/v5_0/applications")
	request.body = requestBody
	request.authenticated = true
	return a.client.doRequest(request, 204, nil)
}

type (
	ApplicationLogAndStatusGetControlPlaneListGet200ResponseListSlice []*ApplicationLogAndStatusGetControlPlaneListGet200ResponseList

	ApplicationLogAndStatusGetControlPlaneListGet200ResponseList struct {
		CpName *string `json:"cpName,omitempty"`
		CpUUID *string `json:"cpUUID,omitempty"`
	}

	ApplicationLogAndStatusGetControlPlaneListGet200Response struct {
		FirstIndex *int                                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                                             `json:"hasMore,omitempty"`
		List       ApplicationLogAndStatusGetControlPlaneListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                              `json:"totalCount,omitempty"`
	}
)

// ApplicationLogAndStatusGetControlPlaneListGet: Use this API command to retrieve a list of control plane.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationLogAndStatusGetControlPlaneListGet200Response
// - error: Error seen or nil if none
func (a *ApplicationsAPI) ApplicationLogAndStatusGetControlPlaneListGet(ctx *UserContext) (*http.Response, *ApplicationLogAndStatusGetControlPlaneListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/applications/controlplane")
	request.authenticated = true
	out := &ApplicationLogAndStatusGetControlPlaneListGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// ApplicationLogAndStatusDownloadSnapshotLogGet: Use this API command to download snapshot logs.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *ApplicationsAPI) ApplicationLogAndStatusDownloadSnapshotLogGet(ctx *UserContext, bladeUUID string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/applications/downloadsnap/{bladeUUID}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}
	return a.client.doRequest(request, 200, nil)
}

// ApplicationLogAndStatusDownloadLogGet: Use this API command to download logs of the application.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - bladeUUID (UUIDv4)
// - appName (string): Download all logs of the specified application name.
//
// Optional Parameter Map:
// - logFileName (string): Download log of the specified log file name.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *ApplicationsAPI) ApplicationLogAndStatusDownloadLogGet(ctx *UserContext, bladeUUID string, appName string, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(appName)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"appName\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/applications/download/{bladeUUID}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}
	request.queryParameters = map[string]string{
		"appName":     appName,
		"logFileName": optionalParams["logFileName"],
	}
	return a.client.doRequest(request, 200, nil)
}

type (
	ApplicationLogAndStatusRetrieveListGet200ResponseListSlice []*ApplicationLogAndStatusRetrieveListGet200ResponseList

	ApplicationLogAndStatusRetrieveListGet200ResponseListLogFileNamesSlice []*string

	ApplicationLogAndStatusRetrieveListGet200ResponseList struct {
		ApplicationName *string                                                                `json:"applicationName,omitempty"`
		HealthStatus    *string                                                                `json:"healthStatus,omitempty"`
		LogFileNames    ApplicationLogAndStatusRetrieveListGet200ResponseListLogFileNamesSlice `json:"logFileNames,omitempty"`
		LogLevel        *string                                                                `json:"logLevel,omitempty"`
		NumOfLogs       *float64                                                               `json:"numOfLogs,omitempty"`
	}

	ApplicationLogAndStatusRetrieveListGet200Response struct {
		FirstIndex *int                                                       `json:"firstIndex,omitempty"`
		HasMore    *bool                                                      `json:"hasMore,omitempty"`
		List       ApplicationLogAndStatusRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                       `json:"totalCount,omitempty"`
	}
)

// ApplicationLogAndStatusRetrieveListGet: Use this API command to retrieve a list of application log and status.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - bladeUUID (UUIDv4)
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationLogAndStatusRetrieveListGet200Response
// - error: Error seen or nil if none
func (a *ApplicationsAPI) ApplicationLogAndStatusRetrieveListGet(ctx *UserContext, bladeUUID string, optionalParams map[string]string) (*http.Response, *ApplicationLogAndStatusRetrieveListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(bladeUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"bladeUUID\" failed validation check: %s", err)
	}
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
	request := a.client.newRequest(ctx, "GET", "/v5_0/applications/{bladeUUID}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"bladeUUID": bladeUUID,
	}
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}
	out := &ApplicationLogAndStatusRetrieveListGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
