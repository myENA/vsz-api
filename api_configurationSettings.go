package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type ConfigurationSettingsAPI struct {
	client *Client
}
type (
	ConfigurationBackupAndRestoreGetAutoExportBackupGet200Response struct {
		EnableAutoExportBackup *bool   `json:"enableAutoExportBackup,omitempty"`
		FtpServer              *string `json:"ftpServer,omitempty"`
	}
)

// ConfigurationBackupAndRestoreGetAutoExportBackupGet: Get Auto Export Backup Settings
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConfigurationBackupAndRestoreGetAutoExportBackupGet200Response
// - error: Error seen or nil if none
func (c *ConfigurationSettingsAPI) ConfigurationBackupAndRestoreGetAutoExportBackupGet(ctx context.Context) (*http.Response, *ConfigurationBackupAndRestoreGetAutoExportBackupGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/configurationSettings/autoExportBackup", true)
	out := &ConfigurationBackupAndRestoreGetAutoExportBackupGet200Response{}
	httpResponse, _, err := c.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ConfigurationBackupAndRestoreModifyAutoExportBackupPatchRequest struct {
		EnableAutoExportBackup *bool   `json:"enableAutoExportBackup,omitempty"`
		FtpServer              *string `json:"ftpServer,omitempty"`
	}
)

// ConfigurationBackupAndRestoreModifyAutoExportBackupPatch: Modify Auto Export Backup Settings
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ConfigurationBackupAndRestoreModifyAutoExportBackupPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationSettingsAPI) ConfigurationBackupAndRestoreModifyAutoExportBackupPatch(ctx context.Context, requestBody *ConfigurationBackupAndRestoreModifyAutoExportBackupPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/configurationSettings/autoExportBackup", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return c.client.Ensure(ctx, request, 204, nil)
}

type (
	ConfigurationBackupAndRestoreGetScheduleBackupSettingGet200Response struct {
		DateOfMonth          *float64 `json:"dateOfMonth,omitempty"`
		DayOfWeek            *string  `json:"dayOfWeek,omitempty"`
		EnableScheduleBackup *bool    `json:"enableScheduleBackup,omitempty"`
		Hour                 *float64 `json:"hour,omitempty"`
		Interval             *string  `json:"interval,omitempty"`
		Minute               *float64 `json:"minute,omitempty"`
	}
)

// ConfigurationBackupAndRestoreGetScheduleBackupSettingGet: Get Schedule Backup Setting
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConfigurationBackupAndRestoreGetScheduleBackupSettingGet200Response
// - error: Error seen or nil if none
func (c *ConfigurationSettingsAPI) ConfigurationBackupAndRestoreGetScheduleBackupSettingGet(ctx context.Context) (*http.Response, *ConfigurationBackupAndRestoreGetScheduleBackupSettingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/configurationSettings/scheduleBackup", true)
	out := &ConfigurationBackupAndRestoreGetScheduleBackupSettingGet200Response{}
	httpResponse, _, err := c.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatchRequest struct {
		DateOfMonth          *float64 `json:"dateOfMonth,omitempty"`
		DayOfWeek            *string  `json:"dayOfWeek,omitempty"`
		EnableScheduleBackup *bool    `json:"enableScheduleBackup,omitempty"`
		Hour                 *float64 `json:"hour,omitempty"`
		Interval             *string  `json:"interval,omitempty"`
		Minute               *float64 `json:"minute,omitempty"`
	}
)

// ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatch: Modify Schedule Backup Setting
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ConfigurationSettingsAPI) ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatch(ctx context.Context, requestBody *ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/configurationSettings/scheduleBackup", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return c.client.Ensure(ctx, request, 204, nil)
}
