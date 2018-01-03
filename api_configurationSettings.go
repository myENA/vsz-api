package api

import (
	"context"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type ConfigurationSettings struct {
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
func (c *ConfigurationSettings) ConfigurationBackupAndRestoreGetAutoExportBackupGet(ctx context.Context) (*http.Response, *ConfigurationBackupAndRestoreGetAutoExportBackupGet200Response, error) {
	request := c.client.newRequest(ctx, "GET", "/v5_0/configurationSettings/autoExportBackup")
	request.authenticated = true

	out := &ConfigurationBackupAndRestoreGetAutoExportBackupGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
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
func (c *ConfigurationSettings) ConfigurationBackupAndRestoreModifyAutoExportBackupPatch(ctx context.Context, requestBody *ConfigurationBackupAndRestoreModifyAutoExportBackupPatchRequest) (*http.Response, []byte, error) {
	request := c.client.newRequest(ctx, "PATCH", "/v5_0/configurationSettings/autoExportBackup")
	request.body = requestBody
	request.authenticated = true

	return c.client.doRequest(request, 204, nil)
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
func (c *ConfigurationSettings) ConfigurationBackupAndRestoreGetScheduleBackupSettingGet(ctx context.Context) (*http.Response, *ConfigurationBackupAndRestoreGetScheduleBackupSettingGet200Response, error) {
	request := c.client.newRequest(ctx, "GET", "/v5_0/configurationSettings/scheduleBackup")
	request.authenticated = true

	out := &ConfigurationBackupAndRestoreGetScheduleBackupSettingGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
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
func (c *ConfigurationSettings) ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatch(ctx context.Context, requestBody *ConfigurationBackupAndRestoreModifyScheduleBackupSettingPatchRequest) (*http.Response, []byte, error) {
	request := c.client.newRequest(ctx, "PATCH", "/v5_0/configurationSettings/scheduleBackup")
	request.body = requestBody
	request.authenticated = true

	return c.client.doRequest(request, 204, nil)
}
