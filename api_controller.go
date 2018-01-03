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

type Controller struct {
	client *Client
}
type (
	SystemSystemSummaryGet200ResponseListSlice []*SystemSystemSummaryGet200ResponseList

	SystemSystemSummaryGet200ResponseList struct {
		ApVersion      *string  `json:"apVersion,omitempty"`
		ClusterIP      *string  `json:"clusterIp,omitempty"`
		ClusterIpv6    *string  `json:"clusterIpv6,omitempty"`
		ClusterRole    *string  `json:"clusterRole,omitempty"`
		ControlIP      *string  `json:"controlIp,omitempty"`
		ControlIpv6    *string  `json:"controlIpv6,omitempty"`
		ControlNatIP   *string  `json:"controlNatIp,omitempty"`
		Description    *string  `json:"description,omitempty"`
		HostName       *string  `json:"hostName,omitempty"`
		ID             *string  `json:"id,omitempty"`
		Mac            *string  `json:"mac,omitempty"`
		ManagementIP   *string  `json:"managementIp,omitempty"`
		ManagementIpv6 *string  `json:"managementIpv6,omitempty"`
		Model          *string  `json:"model,omitempty"`
		Name           *string  `json:"name,omitempty"`
		SerialNumber   *string  `json:"serialNumber,omitempty"`
		UptimeInSec    *float64 `json:"uptimeInSec,omitempty"`
		Version        *string  `json:"version,omitempty"`
	}

	SystemSystemSummaryGet200Response struct {
		FirstIndex *int                                       `json:"firstIndex,omitempty"`
		HasMore    *bool                                      `json:"hasMore,omitempty"`
		List       SystemSystemSummaryGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                       `json:"totalCount,omitempty"`
	}
)

// SystemSystemSummaryGet: Use this API command to retrieve the system summary.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *SystemSystemSummaryGet200Response
// - error: Error seen or nil if none
func (c *Controller) SystemSystemSummaryGet(ctx context.Context) (*http.Response, *SystemSystemSummaryGet200Response, error) {
	request := c.client.newRequest(ctx, "GET", "/v5_0/controller")
	request.authenticated = true

	out := &SystemSystemSummaryGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	SystemSystemStatisticsGet200ResponseSlice []*SystemSystemStatisticsGet200Response

	SystemSystemStatisticsGet200ResponseCluster struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponseControl struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponseCPU struct {
		MaxPercent *float64 `json:"maxPercent,omitempty"`
		MinPercent *float64 `json:"minPercent,omitempty"`
		Percent    *float64 `json:"percent,omitempty"`
	}

	SystemSystemStatisticsGet200ResponseDisk struct {
		Free    *float64 `json:"free,omitempty"`
		MaxFree *float64 `json:"maxFree,omitempty"`
		MinFree *float64 `json:"minFree,omitempty"`
		Total   *float64 `json:"total,omitempty"`
	}

	SystemSystemStatisticsGet200ResponseManagement struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponseMemory struct {
		MaxPercent *float64 `json:"maxPercent,omitempty"`
		MinPercent *float64 `json:"minPercent,omitempty"`
		Percent    *float64 `json:"percent,omitempty"`
	}

	SystemSystemStatisticsGet200ResponsePort0 struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponsePort1 struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponsePort2 struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponsePort3 struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponsePort4 struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200ResponsePort5 struct {
		RxBps     *float64 `json:"rxBps,omitempty"`
		RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
		RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
		RxBytes   *int     `json:"rxBytes,omitempty"`
		RxDropped *float64 `json:"rxDropped,omitempty"`
		RxPackets *float64 `json:"rxPackets,omitempty"`
		TxBps     *float64 `json:"txBps,omitempty"`
		TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
		TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
		TxBytes   *int     `json:"txBytes,omitempty"`
		TxDropped *float64 `json:"txDropped,omitempty"`
		TxPackets *float64 `json:"txPackets,omitempty"`
	}

	SystemSystemStatisticsGet200Response struct {
		Cluster    *SystemSystemStatisticsGet200ResponseCluster    `json:"cluster,omitempty"`
		Control    *SystemSystemStatisticsGet200ResponseControl    `json:"control,omitempty"`
		CpID       *string                                         `json:"cpId,omitempty"`
		CPU        *SystemSystemStatisticsGet200ResponseCPU        `json:"cpu,omitempty"`
		Disk       *SystemSystemStatisticsGet200ResponseDisk       `json:"disk,omitempty"`
		Management *SystemSystemStatisticsGet200ResponseManagement `json:"management,omitempty"`
		Memory     *SystemSystemStatisticsGet200ResponseMemory     `json:"memory,omitempty"`
		Port0      *SystemSystemStatisticsGet200ResponsePort0      `json:"port0,omitempty"`
		Port1      *SystemSystemStatisticsGet200ResponsePort1      `json:"port1,omitempty"`
		Port2      *SystemSystemStatisticsGet200ResponsePort2      `json:"port2,omitempty"`
		Port3      *SystemSystemStatisticsGet200ResponsePort3      `json:"port3,omitempty"`
		Port4      *SystemSystemStatisticsGet200ResponsePort4      `json:"port4,omitempty"`
		Port5      *SystemSystemStatisticsGet200ResponsePort5      `json:"port5,omitempty"`
		Timestamp  *float64                                        `json:"timestamp,omitempty"`
	}
)

// SystemSystemStatisticsGet: Use this API command to retrieve the system statistics.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - id (UUIDv4): Controller ID
//
// Optional Parameter Map:
// - interval (string): Interval, only valid of (QUARTER, HOUR, DAY).
// - size (integer): Size, list size to response.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - SystemSystemStatisticsGet200ResponseSlice
// - error: Error seen or nil if none
func (c *Controller) SystemSystemStatisticsGet(ctx context.Context, id string, optionalParams map[string]string) (*http.Response, SystemSystemStatisticsGet200ResponseSlice, error) {
	var err error

	err = validators.StrIsUUIDv4(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}

	interval, ok := optionalParams["interval"]
	if !ok {
		interval = "QUARTER"
	}
	size, ok := optionalParams["size"]
	if !ok {
		size = "32"
	}

	request := c.client.newRequest(ctx, "GET", "/v5_0/controller/{id}/statistics")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	request.queryParameters = map[string]string{
		"interval": interval,
		"size":     size,
	}

	out := make(SystemSystemStatisticsGet200ResponseSlice, 0)
	httpResponse, _, err := c.client.doRequest(request, 200, &(out))
	return httpResponse, out, err
}
