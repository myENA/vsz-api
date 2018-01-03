package api

import (
	"context"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2017-11-28T11:37:31-0600
// API Version: v5

type Upgrade struct {
	client *Client
}
type (
	SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgressBladeProgresssSlice []*SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgressBladeProgresss

	SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgressBladeProgresss struct {
		BladeUUID     *string  `json:"bladeUUID,omitempty"`
		HostName      *string  `json:"hostName,omitempty"`
		IterationName *string  `json:"iterationName,omitempty"`
		Progress      *float64 `json:"progress,omitempty"`
		State         *string  `json:"state,omitempty"`
	}

	SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgressPreviousOperationRecord struct {
		ErrorMsg  *string `json:"errorMsg,omitempty"`
		Operation *string `json:"operation,omitempty"`
		Success   *bool   `json:"success,omitempty"`
	}

	SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgress struct {
		BladeProgresss             SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgressBladeProgresssSlice      `json:"bladeProgresss,omitempty"`
		ClusterOperationBlockUI    *bool                                                                                     `json:"clusterOperationBlockUI,omitempty"`
		ClusterOperationDisplayMsg *string                                                                                   `json:"clusterOperationDisplayMsg,omitempty"`
		ClusterSubTaskState        *string                                                                                   `json:"clusterSubTaskState,omitempty"`
		IsSelfBladeRebooting       *bool                                                                                     `json:"isSelfBladeRebooting,omitempty"`
		Operation                  *string                                                                                   `json:"operation,omitempty"`
		OverallProgress            *float64                                                                                  `json:"overallProgress,omitempty"`
		PreviousOperationRecord    *SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgressPreviousOperationRecord `json:"previousOperationRecord,omitempty"`
	}

	SystemUpgradeSystemUpgradePost200Response struct {
		ClusterOperationProgress *SystemUpgradeSystemUpgradePost200ResponseClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
	}
)

// SystemUpgradeSystemUpgradePost: Use this API command to do system upgrade
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemUpgradeSystemUpgradePost200Response
// - error: Error seen or nil if none
func (u *Upgrade) SystemUpgradeSystemUpgradePost(ctx context.Context) (*http.Response, *SystemUpgradeSystemUpgradePost200Response, error) {
	request := u.client.newRequest(ctx, "POST", "/v5_0/upgrade")
	request.authenticated = true

	out := &SystemUpgradeSystemUpgradePost200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SystemUpgradeRetriveUpgradeHistoryGet200ResponseListSlice []*SystemUpgradeRetriveUpgradeHistoryGet200ResponseList

	SystemUpgradeRetriveUpgradeHistoryGet200ResponseList struct {
		ApFwVersion    *string  `json:"apFwVersion,omitempty"`
		CbVersion      *string  `json:"cbVersion,omitempty"`
		DpVersion      *string  `json:"dpVersion,omitempty"`
		ElapsedSeconds *float64 `json:"elapsedSeconds,omitempty"`
		FileName       *string  `json:"fileName,omitempty"`
		OldApFwVersion *string  `json:"oldApFwVersion,omitempty"`
		OldCbVersion   *string  `json:"oldCbVersion,omitempty"`
		OldDpVersion   *string  `json:"oldDpVersion,omitempty"`
		OldVersion     *string  `json:"oldVersion,omitempty"`
		StartTime      *string  `json:"startTime,omitempty"`
		Version        *string  `json:"version,omitempty"`
	}

	SystemUpgradeRetriveUpgradeHistoryGet200Response struct {
		FirstIndex *int                                                      `json:"firstIndex,omitempty"`
		HasMore    *bool                                                     `json:"hasMore,omitempty"`
		List       SystemUpgradeRetriveUpgradeHistoryGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                      `json:"totalCount,omitempty"`
	}
)

// SystemUpgradeRetriveUpgradeHistoryGet: Use this API command to retrive upgrade history
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - timezone (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemUpgradeRetriveUpgradeHistoryGet200Response
// - error: Error seen or nil if none
func (u *Upgrade) SystemUpgradeRetriveUpgradeHistoryGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *SystemUpgradeRetriveUpgradeHistoryGet200Response, error) {
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

	request := u.client.newRequest(ctx, "GET", "/v5_0/upgrade/history")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
		"timezone": optionalParams["timezone"],
	}

	out := &SystemUpgradeRetriveUpgradeHistoryGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgressBladeProgresssSlice []*SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgressBladeProgresss

	SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgressBladeProgresss struct {
		BladeUUID     *string  `json:"bladeUUID,omitempty"`
		HostName      *string  `json:"hostName,omitempty"`
		IterationName *string  `json:"iterationName,omitempty"`
		Progress      *float64 `json:"progress,omitempty"`
		State         *string  `json:"state,omitempty"`
	}

	SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgressPreviousOperationRecord struct {
		ErrorMsg  *string `json:"errorMsg,omitempty"`
		Operation *string `json:"operation,omitempty"`
		Success   *bool   `json:"success,omitempty"`
	}

	SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgress struct {
		BladeProgresss             SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgressBladeProgresssSlice      `json:"bladeProgresss,omitempty"`
		ClusterOperationBlockUI    *bool                                                                                             `json:"clusterOperationBlockUI,omitempty"`
		ClusterOperationDisplayMsg *string                                                                                           `json:"clusterOperationDisplayMsg,omitempty"`
		ClusterSubTaskState        *string                                                                                           `json:"clusterSubTaskState,omitempty"`
		IsSelfBladeRebooting       *bool                                                                                             `json:"isSelfBladeRebooting,omitempty"`
		Operation                  *string                                                                                           `json:"operation,omitempty"`
		OverallProgress            *float64                                                                                          `json:"overallProgress,omitempty"`
		PreviousOperationRecord    *SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgressPreviousOperationRecord `json:"previousOperationRecord,omitempty"`
	}

	SystemUpgradeRetriveUploadPatchInfoGet200ResponseUploadPatchInfoAllowVersionsSlice []*string

	SystemUpgradeRetriveUploadPatchInfoGet200ResponseUploadPatchInfo struct {
		AllowVersions       SystemUpgradeRetriveUploadPatchInfoGet200ResponseUploadPatchInfoAllowVersionsSlice `json:"allowVersions,omitempty"`
		ApVersion           *string                                                                            `json:"apVersion,omitempty"`
		ControlbladeVersion *string                                                                            `json:"controlbladeVersion,omitempty"`
		DatabladeVersion    *string                                                                            `json:"databladeVersion,omitempty"`
		FileName            *string                                                                            `json:"fileName,omitempty"`
		FileSize            *float64                                                                           `json:"fileSize,omitempty"`
		FileUploadPath      *string                                                                            `json:"fileUploadPath,omitempty"`
		Version             *string                                                                            `json:"version,omitempty"`
	}

	SystemUpgradeRetriveUploadPatchInfoGet200Response struct {
		ClusterOperationProgress *SystemUpgradeRetriveUploadPatchInfoGet200ResponseClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
		UploadPatchInfo          *SystemUpgradeRetriveUploadPatchInfoGet200ResponseUploadPatchInfo          `json:"uploadPatchInfo,omitempty"`
	}
)

// SystemUpgradeRetriveUploadPatchInfoGet: Use this API command to retrive upload file Info
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemUpgradeRetriveUploadPatchInfoGet200Response
// - error: Error seen or nil if none
func (u *Upgrade) SystemUpgradeRetriveUploadPatchInfoGet(ctx context.Context) (*http.Response, *SystemUpgradeRetriveUploadPatchInfoGet200Response, error) {
	request := u.client.newRequest(ctx, "GET", "/v5_0/upgrade/patch")
	request.authenticated = true

	out := &SystemUpgradeRetriveUploadPatchInfoGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgressBladeProgresssSlice []*SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgressBladeProgresss

	SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgressBladeProgresss struct {
		BladeUUID     *string  `json:"bladeUUID,omitempty"`
		HostName      *string  `json:"hostName,omitempty"`
		IterationName *string  `json:"iterationName,omitempty"`
		Progress      *float64 `json:"progress,omitempty"`
		State         *string  `json:"state,omitempty"`
	}

	SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgressPreviousOperationRecord struct {
		ErrorMsg  *string `json:"errorMsg,omitempty"`
		Operation *string `json:"operation,omitempty"`
		Success   *bool   `json:"success,omitempty"`
	}

	SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgress struct {
		BladeProgresss             SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgressBladeProgresssSlice      `json:"bladeProgresss,omitempty"`
		ClusterOperationBlockUI    *bool                                                                                                   `json:"clusterOperationBlockUI,omitempty"`
		ClusterOperationDisplayMsg *string                                                                                                 `json:"clusterOperationDisplayMsg,omitempty"`
		ClusterSubTaskState        *string                                                                                                 `json:"clusterSubTaskState,omitempty"`
		IsSelfBladeRebooting       *bool                                                                                                   `json:"isSelfBladeRebooting,omitempty"`
		Operation                  *string                                                                                                 `json:"operation,omitempty"`
		OverallProgress            *float64                                                                                                `json:"overallProgress,omitempty"`
		PreviousOperationRecord    *SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgressPreviousOperationRecord `json:"previousOperationRecord,omitempty"`
	}

	SystemUpgradeRetriveClusterProgressStatusGet200Response struct {
		ClusterOperationProgress *SystemUpgradeRetriveClusterProgressStatusGet200ResponseClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
	}
)

// SystemUpgradeRetriveClusterProgressStatusGet: Use this API command to retrive cluster progress status
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemUpgradeRetriveClusterProgressStatusGet200Response
// - error: Error seen or nil if none
func (u *Upgrade) SystemUpgradeRetriveClusterProgressStatusGet(ctx context.Context) (*http.Response, *SystemUpgradeRetriveClusterProgressStatusGet200Response, error) {
	request := u.client.newRequest(ctx, "GET", "/v5_0/upgrade/status")
	request.authenticated = true

	out := &SystemUpgradeRetriveClusterProgressStatusGet200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SystemUpgradeUploadFilePost200ResponseClusterOperationProgressBladeProgresssSlice []*SystemUpgradeUploadFilePost200ResponseClusterOperationProgressBladeProgresss

	SystemUpgradeUploadFilePost200ResponseClusterOperationProgressBladeProgresss struct {
		BladeUUID     *string  `json:"bladeUUID,omitempty"`
		HostName      *string  `json:"hostName,omitempty"`
		IterationName *string  `json:"iterationName,omitempty"`
		Progress      *float64 `json:"progress,omitempty"`
		State         *string  `json:"state,omitempty"`
	}

	SystemUpgradeUploadFilePost200ResponseClusterOperationProgressPreviousOperationRecord struct {
		ErrorMsg  *string `json:"errorMsg,omitempty"`
		Operation *string `json:"operation,omitempty"`
		Success   *bool   `json:"success,omitempty"`
	}

	SystemUpgradeUploadFilePost200ResponseClusterOperationProgress struct {
		BladeProgresss             SystemUpgradeUploadFilePost200ResponseClusterOperationProgressBladeProgresssSlice      `json:"bladeProgresss,omitempty"`
		ClusterOperationBlockUI    *bool                                                                                  `json:"clusterOperationBlockUI,omitempty"`
		ClusterOperationDisplayMsg *string                                                                                `json:"clusterOperationDisplayMsg,omitempty"`
		ClusterSubTaskState        *string                                                                                `json:"clusterSubTaskState,omitempty"`
		IsSelfBladeRebooting       *bool                                                                                  `json:"isSelfBladeRebooting,omitempty"`
		Operation                  *string                                                                                `json:"operation,omitempty"`
		OverallProgress            *float64                                                                               `json:"overallProgress,omitempty"`
		PreviousOperationRecord    *SystemUpgradeUploadFilePost200ResponseClusterOperationProgressPreviousOperationRecord `json:"previousOperationRecord,omitempty"`
	}

	SystemUpgradeUploadFilePost200Response struct {
		ClusterOperationProgress *SystemUpgradeUploadFilePost200ResponseClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
	}
)

// SystemUpgradeUploadFilePost: Use this API command to upload patch file
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemUpgradeUploadFilePost200Response
// - error: Error seen or nil if none
func (u *Upgrade) SystemUpgradeUploadFilePost(ctx context.Context) (*http.Response, *SystemUpgradeUploadFilePost200Response, error) {
	request := u.client.newRequest(ctx, "POST", "/v5_0/upgrade/upload")
	request.authenticated = true

	out := &SystemUpgradeUploadFilePost200Response{}
	httpResponse, _, err := u.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
