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

type ConfigurationAPI struct {
	client *Client
}
type (
	ConfigurationBackupAndRestoreRetrieveListGet200ResponseListSlice []*ConfigurationBackupAndRestoreRetrieveListGet200ResponseList

	ConfigurationBackupAndRestoreRetrieveListGet200ResponseList struct {
		BackupElapsed               *float64    `json:"backupElapsed,omitempty"`
		ControlPlaneSoftwareVersion *string     `json:"controlPlaneSoftwareVersion,omitempty"`
		CreatedBy                   *string     `json:"createdBy,omitempty"`
		CreatedOn                   interface{} `json:"createdOn,omitempty"`
		DataPlaneSoftwareVersion    *string     `json:"dataPlaneSoftwareVersion,omitempty"`
		FileSize                    *float64    `json:"fileSize,omitempty"`
		ID                          *string     `json:"id,omitempty"`
		ScgVersion                  *string     `json:"scgVersion,omitempty"`
		Type                        *string     `json:"type,omitempty"`
	}

	ConfigurationBackupAndRestoreRetrieveListGet200Response struct {
		FirstIndex *int                                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                                            `json:"hasMore,omitempty"`
		List       ConfigurationBackupAndRestoreRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                             `json:"totalCount,omitempty"`
	}
)

// ConfigurationBackupAndRestoreRetrieveListGet: Retrive system configuration list
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConfigurationBackupAndRestoreRetrieveListGet200Response
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreRetrieveListGet(ctx *UserContext, optionalParams map[string]string) (*http.Response, *ConfigurationBackupAndRestoreRetrieveListGet200Response, error) {
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
	request := c.client.newRequest(ctx, "GET", "/v5_0/configuration")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}
	out := &ConfigurationBackupAndRestoreRetrieveListGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ConfigurationBackupAndRestoreSystemConfigurationBackupPost: Backup system configuration
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreSystemConfigurationBackupPost(ctx *UserContext) (*http.Response, *ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/configuration/backup")
	request.authenticated = true
	out := &ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response{}
	httpResponse, _, err := c.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// ConfigurationBackupAndRestoreDownloadGet: Download system configuration file
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - backupUUID (UUIDv4): System configuration file uuid
//
// Optional Parameter Map:
// - timeZone (string): DEPRECATED Time zone
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreDownloadGet(ctx *UserContext, backupUUID string, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(backupUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"backupUUID\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "GET", "/v5_0/configuration/download")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"backupUUID": backupUUID,
		"timeZone":   optionalParams["timeZone"],
	}
	return c.client.doRequest(request, 200, nil)
}

// ConfigurationBackupAndRestoreSystemConfigurationRestorePost: Restore system configuration with specified backupUUID
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): Configuration ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreSystemConfigurationRestorePost(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/configuration/restore/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return c.client.doRequest(request, 204, nil)
}

// ConfigurationBackupAndRestoreUploadPost: Upload system configuration file
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreUploadPost(ctx *UserContext) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/configuration/upload")
	request.authenticated = true
	return c.client.doRequest(request, 204, nil)
}

// ConfigurationBackupAndRestoreDeleteDelete: Delete system configuration file
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (UUIDv4): Configuration ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreDeleteDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "DELETE", "/v5_0/configuration/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return c.client.doRequest(request, 204, nil)
}
