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

type DHCPDataAPI struct {
	client *Client
}
type (
	DhcpGetApDhcpMessageStatisticUsageGet200ResponseDhcpMsgRecvdStats struct {
		DhcpDecline  *float64 `json:"dhcpDecline,omitempty"`
		DhcpDiscover *float64 `json:"dhcpDiscover,omitempty"`
		DhcpInform   *float64 `json:"dhcpInform,omitempty"`
		DhcpRelease  *float64 `json:"dhcpRelease,omitempty"`
		DhcpRequest  *float64 `json:"dhcpRequest,omitempty"`
	}

	DhcpGetApDhcpMessageStatisticUsageGet200ResponseDhcpMsgSentStats struct {
		DhcpAck   *float64 `json:"dhcpAck,omitempty"`
		DhcpNak   *float64 `json:"dhcpNak,omitempty"`
		DhcpOffer *float64 `json:"dhcpOffer,omitempty"`
	}

	DhcpGetApDhcpMessageStatisticUsageGet200Response struct {
		ApMac             *string                                                            `json:"apMac,omitempty"`
		DhcpMsgRecvdStats *DhcpGetApDhcpMessageStatisticUsageGet200ResponseDhcpMsgRecvdStats `json:"dhcpMsgRecvdStats,omitempty"`
		DhcpMsgSentStats  *DhcpGetApDhcpMessageStatisticUsageGet200ResponseDhcpMsgSentStats  `json:"dhcpMsgSentStats,omitempty"`
		DomainID          *string                                                            `json:"domainId,omitempty"`
		ID                *string                                                            `json:"id,omitempty"`
		TenantID          *string                                                            `json:"tenantId,omitempty"`
	}
)

// DhcpGetApDhcpMessageStatisticUsageGet: Use this API command to get AP DHCP Message Statistic
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpGetApDhcpMessageStatisticUsageGet200Response
// - error: Error seen or nil if none
func (d *DHCPDataAPI) DhcpGetApDhcpMessageStatisticUsageGet(ctx *UserContext, apMac string) (*http.Response, *DhcpGetApDhcpMessageStatisticUsageGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := d.client.newRequest(ctx, "GET", "/v5_0/dhcpData/dhcpMsgStats/{apMac}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"apMac": apMac,
	}
	out := &DhcpGetApDhcpMessageStatisticUsageGet200Response{}
	httpResponse, _, err := d.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListSlice []*DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoList

	DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListClientInfoListListSlice []*DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListClientInfoListList

	DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListClientInfoListList struct {
		ClientIP         *string  `json:"clientIp,omitempty"`
		ClientMac        *string  `json:"clientMac,omitempty"`
		LeaseExpirtyTime *float64 `json:"leaseExpirtyTime,omitempty"`
		LeaseExpiryTime  *float64 `json:"leaseExpiryTime,omitempty"`
		LeaseTime        *float64 `json:"leaseTime,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
	}

	DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListClientInfoList struct {
		FirstIndex *int                                                                     `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                    `json:"hasMore,omitempty"`
		List       DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListClientInfoListListSlice `json:"list,omitempty"`
		TotalCount *int                                                                     `json:"totalCount,omitempty"`
	}

	DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoList struct {
		ApIP             *string                                                          `json:"apIp,omitempty"`
		AvailableIPCount *float64                                                         `json:"availableIpCount,omitempty"`
		ClientInfoList   *DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListClientInfoList `json:"clientInfoList,omitempty"`
		PoolIndex        *string                                                          `json:"poolIndex,omitempty"`
		TotalIPCount     *float64                                                         `json:"totalIpCount,omitempty"`
		UsedIPCount      *float64                                                         `json:"usedIpCount,omitempty"`
		VlanID           *string                                                          `json:"vlanId,omitempty"`
	}

	DhcpGetApDhcpPoolsUsageGet200Response struct {
		ApMac        *string                                                `json:"apMac,omitempty"`
		DomainEntity *string                                                `json:"domainEntity,omitempty"`
		DomainID     *string                                                `json:"domainId,omitempty"`
		ID           *string                                                `json:"id,omitempty"`
		PoolInfoList DhcpGetApDhcpPoolsUsageGet200ResponsePoolInfoListSlice `json:"poolInfoList,omitempty"`
		TenantID     *string                                                `json:"tenantId,omitempty"`
	}
)

// DhcpGetApDhcpPoolsUsageGet: Use this API command to get AP DHCP Pools Usage
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpGetApDhcpPoolsUsageGet200Response
// - error: Error seen or nil if none
func (d *DHCPDataAPI) DhcpGetApDhcpPoolsUsageGet(ctx *UserContext, apMac string) (*http.Response, *DhcpGetApDhcpPoolsUsageGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := d.client.newRequest(ctx, "GET", "/v5_0/dhcpData/dhcpPools/{apMac}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"apMac": apMac,
	}
	out := &DhcpGetApDhcpPoolsUsageGet200Response{}
	httpResponse, _, err := d.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	DhcpGetApDhcpPoolUsageByPoolSIndexGet200ResponseClientInfoListListSlice []*DhcpGetApDhcpPoolUsageByPoolSIndexGet200ResponseClientInfoListList

	DhcpGetApDhcpPoolUsageByPoolSIndexGet200ResponseClientInfoListList struct {
		ClientIP         *string  `json:"clientIp,omitempty"`
		ClientMac        *string  `json:"clientMac,omitempty"`
		LeaseExpirtyTime *float64 `json:"leaseExpirtyTime,omitempty"`
		LeaseExpiryTime  *float64 `json:"leaseExpiryTime,omitempty"`
		LeaseTime        *float64 `json:"leaseTime,omitempty"`
		LeaseTimeHours   *float64 `json:"leaseTimeHours,omitempty"`
		LeaseTimeMinutes *float64 `json:"leaseTimeMinutes,omitempty"`
	}

	DhcpGetApDhcpPoolUsageByPoolSIndexGet200ResponseClientInfoList struct {
		FirstIndex *int                                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                   `json:"hasMore,omitempty"`
		List       DhcpGetApDhcpPoolUsageByPoolSIndexGet200ResponseClientInfoListListSlice `json:"list,omitempty"`
		TotalCount *int                                                                    `json:"totalCount,omitempty"`
	}

	DhcpGetApDhcpPoolUsageByPoolSIndexGet200Response struct {
		ApIP             *string                                                         `json:"apIp,omitempty"`
		AvailableIPCount *float64                                                        `json:"availableIpCount,omitempty"`
		ClientInfoList   *DhcpGetApDhcpPoolUsageByPoolSIndexGet200ResponseClientInfoList `json:"clientInfoList,omitempty"`
		PoolIndex        *float64                                                        `json:"poolIndex,omitempty"`
		TotalIPCount     *float64                                                        `json:"totalIpCount,omitempty"`
		UsedIPCount      *float64                                                        `json:"usedIpCount,omitempty"`
		VlanID           *float64                                                        `json:"vlanId,omitempty"`
	}
)

// DhcpGetApDhcpPoolUsageByPoolSIndexGet: Use this API command to get AP DHCP Pool Usage by pool's index
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - poolIndex (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *DhcpGetApDhcpPoolUsageByPoolSIndexGet200Response
// - error: Error seen or nil if none
func (d *DHCPDataAPI) DhcpGetApDhcpPoolUsageByPoolSIndexGet(ctx *UserContext, apMac string, poolIndex string) (*http.Response, *DhcpGetApDhcpPoolUsageByPoolSIndexGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	err = validators.StrNotEmpty(poolIndex)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"poolIndex\" failed validation check: %s", err)
	}
	request := d.client.newRequest(ctx, "GET", "/v5_0/dhcpData/dhcpPools/{apMac}/{poolIndex}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"apMac":     apMac,
		"poolIndex": poolIndex,
	}
	out := &DhcpGetApDhcpPoolUsageByPoolSIndexGet200Response{}
	httpResponse, _, err := d.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
