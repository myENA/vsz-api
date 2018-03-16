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
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConfigurationBackupAndRestoreRetrieveListGet200Response
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *ConfigurationBackupAndRestoreRetrieveListGet200Response, error) {
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
	request := NewRequest("GET", "/v5_0/configuration", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &ConfigurationBackupAndRestoreRetrieveListGet200Response{}
	httpResponse, _, err := c.client.Ensure(ctx, request, 200, out)
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
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreSystemConfigurationBackupPost(ctx context.Context) (*http.Response, *ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/configuration/backup", true)
	out := &ConfigurationBackupAndRestoreSystemConfigurationBackupPost201Response{}
	httpResponse, _, err := c.client.Ensure(ctx, request, 201, out)
	return httpResponse, out, err
}

// ConfigurationBackupAndRestoreDownloadGet: Download system configuration file
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - backupUUID (UUIDv4): System configuration file uuid
//
// Optional Parameter Map:
// - timeZone (string): DEPRECATED Time zone
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreDownloadGet(ctx context.Context, backupUUID string, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(backupUUID)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"backupUUID\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/configuration/download", true)
	request.SetQueryParameter("backupUUID", backupUUID)
	request.SetQueryParameter("timeZone", optionalParams["timeZone"])
	return c.client.Ensure(ctx, request, 200, nil)
}

// ConfigurationBackupAndRestoreSystemConfigurationRestorePost: Restore system configuration with specified backupUUID
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Configuration ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreSystemConfigurationRestorePost(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/configuration/restore/{id}", true)
	request.SetPathParameter("id", id)
	return c.client.Ensure(ctx, request, 204, nil)
}

// ConfigurationBackupAndRestoreUploadPost: Upload system configuration file
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreUploadPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/configuration/upload", true)
	return c.client.Ensure(ctx, request, 204, nil)
}

// ConfigurationBackupAndRestoreDeleteDelete: Delete system configuration file
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Configuration ID
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationAPI) ConfigurationBackupAndRestoreDeleteDelete(ctx context.Context, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/configuration/{id}", true)
	request.SetPathParameter("id", id)
	return c.client.Ensure(ctx, request, 204, nil)
}
