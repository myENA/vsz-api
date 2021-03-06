package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/vsz-api/validators"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-04-27T15:10:38-0500
// API Version: v5

type ToolAPI struct {
	client *Client
}

// ConnectivityToolsPingGet: Use this API command to run the PING test on an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): MAC address of the AP running the PING test
// - targetIP (MAC): the IP address to PING
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (t *ToolAPI) ConnectivityToolsPingGet(ctx context.Context, apMac string, targetIP string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	err = validators.StrIsIP(targetIP)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"targetIP\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/tool/ping", true)
	request.SetQueryParameter("apMac", apMac)
	request.SetQueryParameter("targetIP", targetIP)
	return t.client.Ensure(ctx, request, 200, nil)
}

type (
	ConnectivityToolsSpeedflexPostRequest struct {
		ClientIP  *string  `json:"clientIp,omitempty"`
		ClientMac *string  `json:"clientMac,omitempty"`
		Model     *string  `json:"model,omitempty"`
		Protocol  *string  `json:"protocol,omitempty"`
		ServerIP  *string  `json:"serverIp,omitempty"`
		ServerMac *string  `json:"serverMac,omitempty"`
		Syspmtu   *float64 `json:"syspmtu,omitempty"`
		Tool      *string  `json:"tool,omitempty"`
	}

	ConnectivityToolsSpeedflexPost200Response struct {
		Downlink   *float64 `json:"downlink,omitempty"`
		Etf        *float64 `json:"etf,omitempty"`
		Latency    *float64 `json:"latency,omitempty"`
		PacketLoss *float64 `json:"packetLoss,omitempty"`
		ResultID   *float64 `json:"resultId,omitempty"`
		Uplink     *float64 `json:"uplink,omitempty"`
		Wcid       *string  `json:"wcid,omitempty"`
	}
)

// ConnectivityToolsSpeedflexPost: Use this API command to start the SpeedFlex test.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *ConnectivityToolsSpeedflexPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConnectivityToolsSpeedflexPost200Response
// - error: Error seen or nil if none
func (t *ToolAPI) ConnectivityToolsSpeedflexPost(ctx context.Context, requestBody *ConnectivityToolsSpeedflexPostRequest) (*http.Response, *ConnectivityToolsSpeedflexPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/tool/speedflex", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	out := &ConnectivityToolsSpeedflexPost200Response{}
	httpResponse, _, err := t.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	ConnectivityToolsRetrieveSppedflexResultsGet200Response struct {
		Downlink   *float64 `json:"downlink,omitempty"`
		Etf        *float64 `json:"etf,omitempty"`
		Latency    *float64 `json:"latency,omitempty"`
		PacketLoss *float64 `json:"packetLoss,omitempty"`
		ResultID   *float64 `json:"resultId,omitempty"`
		Uplink     *float64 `json:"uplink,omitempty"`
		Wcid       *string  `json:"wcid,omitempty"`
	}
)

// ConnectivityToolsRetrieveSppedflexResultsGet: Use this API command to retrieve existing SpeedFlex test results.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - wcid (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ConnectivityToolsRetrieveSppedflexResultsGet200Response
// - error: Error seen or nil if none
func (t *ToolAPI) ConnectivityToolsRetrieveSppedflexResultsGet(ctx context.Context, wcid string) (*http.Response, *ConnectivityToolsRetrieveSppedflexResultsGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(wcid)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"wcid\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/tool/speedflex/{wcid}", true)
	request.SetPathParameter("wcid", wcid)
	out := &ConnectivityToolsRetrieveSppedflexResultsGet200Response{}
	httpResponse, _, err := t.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// ConnectivityToolsTraceRouteGet: Use this API command to run the traceroute test on an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): MAC address of the AP running the traceRoute test
// - targetIP (MAC): the target IP address to traceRoute
//
// Optional Parameter Map:
// - timeoutInSec (integer): Timeout in unit of seconds (
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (t *ToolAPI) ConnectivityToolsTraceRouteGet(ctx context.Context, apMac string, targetIP string, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	err = validators.StrIsIP(targetIP)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"targetIP\" failed validation check: %s", err)
	}
	timeoutInSec, ok := optionalParams["timeoutInSec"]
	if ok {
		err = validators.StrIsPositiveInt(timeoutInSec)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"timeoutInSec\" failed validation check: %s", err)
		}
	} else {
		timeoutInSec = "30"
	}
	request := NewRequest("GET", "/v5_0/tool/traceRoute", true)
	request.SetQueryParameter("apMac", apMac)
	request.SetQueryParameter("targetIP", targetIP)
	request.SetQueryParameter("timeoutInSec", timeoutInSec)
	return t.client.Ensure(ctx, request, 200, nil)
}
