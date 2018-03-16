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

type APsAPI struct {
	client *Client
}
type (
	AccessPointConfigurationRetrieveListGet200ResponseListSlice []*AccessPointConfigurationRetrieveListGet200ResponseList

	AccessPointConfigurationRetrieveListGet200ResponseList struct {
		ApGroupID *string `json:"apGroupId,omitempty"`
		Mac       *string `json:"mac,omitempty"`
		Name      *string `json:"name,omitempty"`
		ZoneID    *string `json:"zoneId,omitempty"`
	}

	AccessPointConfigurationRetrieveListGet200Response struct {
		FirstIndex *int                                                        `json:"firstIndex,omitempty"`
		HasMore    *bool                                                       `json:"hasMore,omitempty"`
		List       AccessPointConfigurationRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                        `json:"totalCount,omitempty"`
	}
)

// AccessPointConfigurationRetrieveListGet: Use this API command to retrieve the list of APs that belong to a zone or a domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - zoneId (UUIDv4): filter AP list by zone
// - domainId (UUIDv4): filter AP list by domain.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointConfigurationRetrieveListGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *AccessPointConfigurationRetrieveListGet200Response, error) {
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
	zoneId, ok := optionalParams["zoneId"]
	if ok {
		err = validators.StrIsUUIDv4(zoneId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
		}
	}
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/aps", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("zoneId", zoneId)
	request.SetQueryParameter("domainId", domainId)
	out := &AccessPointConfigurationRetrieveListGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	AccessPointConfigurationCreatePostRequest struct {
		AdministrativeState *string  `json:"administrativeState,omitempty"`
		ApGroupID           *string  `json:"apGroupId,omitempty"`
		AwsVenue            *string  `json:"awsVenue,omitempty"`
		Description         *string  `json:"description,omitempty"`
		GpsSource           *string  `json:"gpsSource,omitempty"`
		Latitude            *float64 `json:"latitude,omitempty"`
		Location            *string  `json:"location,omitempty"`
		Longitude           *float64 `json:"longitude,omitempty"`
		Mac                 *string  `json:"mac,omitempty"`
		Model               *string  `json:"model,omitempty"`
		Name                *string  `json:"name,omitempty"`
		ProvisionChecklist  *string  `json:"provisionChecklist,omitempty"`
		Serial              *string  `json:"serial,omitempty"`
		ZoneID              *string  `json:"zoneId,omitempty"`
	}
)

// AccessPointConfigurationCreatePost: Use this API command to create a new access point.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *AccessPointConfigurationCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationCreatePost(ctx context.Context, requestBody *AccessPointConfigurationCreatePostRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/v5_0/aps", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return a.client.Ensure(ctx, request, 201, nil)
}

type (
	ApAppRetrieveApSummaryGet200ResponseListSlice []*ApAppRetrieveApSummaryGet200ResponseList

	ApAppRetrieveApSummaryGet200ResponseListAlarms struct {
		CriticalCount *int `json:"criticalCount,omitempty"`
		MajorCount    *int `json:"majorCount,omitempty"`
		MinorCount    *int `json:"minorCount,omitempty"`
		WarningCount  *int `json:"warningCount,omitempty"`
	}

	ApAppRetrieveApSummaryGet200ResponseList struct {
		Alarms      *ApAppRetrieveApSummaryGet200ResponseListAlarms `json:"alarms,omitempty"`
		ConfigState *string                                         `json:"configState,omitempty"`
		Latitude    *float64                                        `json:"latitude,omitempty"`
		Location    *string                                         `json:"location,omitempty"`
		Longitude   *float64                                        `json:"longitude,omitempty"`
		Mac         *string                                         `json:"mac,omitempty"`
		Name        *string                                         `json:"name,omitempty"`
	}

	ApAppRetrieveApSummaryGet200Response struct {
		FirstIndex *int                                          `json:"firstIndex,omitempty"`
		HasMore    *bool                                         `json:"hasMore,omitempty"`
		List       ApAppRetrieveApSummaryGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                          `json:"totalCount,omitempty"`
	}
)

// ApAppRetrieveApSummaryGet: Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - zoneId (UUIDv4): filter AP lineman by zone.
// - domainId (UUIDv4): filter AP lineman by domain.
// - showAlarm (bool): indicate to show alarm counter.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApAppRetrieveApSummaryGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) ApAppRetrieveApSummaryGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *ApAppRetrieveApSummaryGet200Response, error) {
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
	zoneId, ok := optionalParams["zoneId"]
	if ok {
		err = validators.StrIsUUIDv4(zoneId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
		}
	}
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	showAlarm, ok := optionalParams["showAlarm"]
	if ok {
		err = validators.StrIsBool(showAlarm)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"showAlarm\" failed validation check: %s", err)
		}
	} else {
		showAlarm = "true"
	}
	request := NewRequest("GET", "/v5_0/aps/lineman", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("zoneId", zoneId)
	request.SetQueryParameter("domainId", domainId)
	request.SetQueryParameter("showAlarm", showAlarm)
	out := &ApAppRetrieveApSummaryGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// ApAppRetrieveTotalApCountGet: Use this API command to retrieve the total AP count within a zone or a domain.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - zoneId (UUIDv4): filter AP total count by zone.
// - domainId (UUIDv4): filter AP total count by domain.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) ApAppRetrieveTotalApCountGet(ctx context.Context, optionalParams map[string]string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	zoneId, ok := optionalParams["zoneId"]
	if ok {
		err = validators.StrIsUUIDv4(zoneId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"zoneId\" failed validation check: %s", err)
		}
	}
	domainId, ok := optionalParams["domainId"]
	if ok {
		err = validators.StrIsUUIDv4(domainId)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"domainId\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/aps/totalCount", true)
	request.SetQueryParameter("zoneId", zoneId)
	request.SetQueryParameter("domainId", domainId)
	return a.client.Ensure(ctx, request, 200, nil)
}

// AccessPointConfigurationDeleteDelete: Use this API command to delete an access point.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDeleteDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationRetrieveGet200ResponseAltitude struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseApMgmtVlan struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseAutoChannelSelection24 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseAutoChannelSelection50 struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseBonjourGateway struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseClientAdmissionControl24 struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseClientAdmissionControl50 struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseLogin struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseMeshOptionsMeshUplinkEntryListSlice []*string

	AccessPointConfigurationRetrieveGet200ResponseMeshOptions struct {
		MeshMode            *string                                                                           `json:"meshMode,omitempty"`
		MeshUplinkEntryList AccessPointConfigurationRetrieveGet200ResponseMeshOptionsMeshUplinkEntryListSlice `json:"meshUplinkEntryList,omitempty"`
		UplinkSelection     *string                                                                           `json:"uplinkSelection,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseNetwork struct {
		Gateway      *string `json:"gateway,omitempty"`
		IP           *string `json:"ip,omitempty"`
		IPType       *string `json:"ipType,omitempty"`
		Netmask      *string `json:"netmask,omitempty"`
		PrimaryDNS   *string `json:"primaryDns,omitempty"`
		SecondaryDNS *string `json:"secondaryDns,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSmartMonitor struct {
		Enabled        *bool `json:"enabled,omitempty"`
		IntervalInSec  *int  `json:"intervalInSec,omitempty"`
		RetryThreshold *int  `json:"retryThreshold,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificExternalAntenna24 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificExternalAntenna50 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsSlice []*AccessPointConfigurationRetrieveGet200ResponseSpecificLanPorts

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XAuthenticatorAccounting struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XAuthenticatorAuthentication struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XAuthenticator struct {
		Accounting           *AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XAuthenticatorAccounting     `json:"accounting,omitempty"`
		Authentication       *AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XAuthenticatorAuthentication `json:"authentication,omitempty"`
		MacAuthByPassEnabled *bool                                                                                            `json:"macAuthByPassEnabled,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XSupplicant struct {
		Password *string `json:"password,omitempty"`
		Type     *string `json:"type,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021X struct {
		Authenticator *AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XAuthenticator `json:"authenticator,omitempty"`
		Supplicant    *AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021XSupplicant    `json:"supplicant,omitempty"`
		Type          *string                                                                            `json:"type,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsEthPortProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLanPorts struct {
		X8021X               *AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsX8021X         `json:"_8021X,omitempty"`
		Enabled              *bool                                                                         `json:"enabled,omitempty"`
		EthPortProfile       *AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsEthPortProfile `json:"ethPortProfile,omitempty"`
		Members              *string                                                                       `json:"members,omitempty"`
		OverwriteVlanEnabled *bool                                                                         `json:"overwriteVlanEnabled,omitempty"`
		PortName             *string                                                                       `json:"portName,omitempty"`
		Type                 *string                                                                       `json:"type,omitempty"`
		VlanUntagID          *float64                                                                      `json:"vlanUntagId,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecificLldp struct {
		AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
		Enabled                *bool `json:"enabled,omitempty"`
		HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
		ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSpecific struct {
		ExternalAntenna24     *AccessPointConfigurationRetrieveGet200ResponseSpecificExternalAntenna24 `json:"externalAntenna24,omitempty"`
		ExternalAntenna50     *AccessPointConfigurationRetrieveGet200ResponseSpecificExternalAntenna50 `json:"externalAntenna50,omitempty"`
		InternalHeaterEnabled *bool                                                                    `json:"internalHeaterEnabled,omitempty"`
		LanPorts              AccessPointConfigurationRetrieveGet200ResponseSpecificLanPortsSlice      `json:"lanPorts,omitempty"`
		LedMode               interface{}                                                              `json:"ledMode,omitempty"`
		LedStatusEnabled      *bool                                                                    `json:"ledStatusEnabled,omitempty"`
		Lldp                  *AccessPointConfigurationRetrieveGet200ResponseSpecificLldp              `json:"lldp,omitempty"`
		PoeModeSetting        *string                                                                  `json:"poeModeSetting,omitempty"`
		PoeOutPortEnabled     *bool                                                                    `json:"poeOutPortEnabled,omitempty"`
		PoeTxChain            *int                                                                     `json:"poeTxChain,omitempty"`
		RadioBand             interface{}                                                              `json:"radioBand,omitempty"`
		UsbPowerEnable        *bool                                                                    `json:"usbPowerEnable,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseSyslog struct {
		Address  *string `json:"address,omitempty"`
		Enabled  *bool   `json:"enabled,omitempty"`
		Facility *string `json:"facility,omitempty"`
		Port     *int    `json:"port,omitempty"`
		Priority *string `json:"priority,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseVenueProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseWifi24AvailableChannelRangeSlice []*float64

	AccessPointConfigurationRetrieveGet200ResponseWifi24ChannelRangeSlice []*float64

	AccessPointConfigurationRetrieveGet200ResponseWifi24 struct {
		AvailableChannelRange AccessPointConfigurationRetrieveGet200ResponseWifi24AvailableChannelRangeSlice `json:"availableChannelRange,omitempty"`
		Channel               *int                                                                           `json:"channel,omitempty"`
		ChannelRange          AccessPointConfigurationRetrieveGet200ResponseWifi24ChannelRangeSlice          `json:"channelRange,omitempty"`
		ChannelWidth          *int                                                                           `json:"channelWidth,omitempty"`
		TxPower               *string                                                                        `json:"txPower,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseWifi50AvailableChannelRangeSlice []*float64

	AccessPointConfigurationRetrieveGet200ResponseWifi50ChannelRangeSlice []*float64

	AccessPointConfigurationRetrieveGet200ResponseWifi50 struct {
		AvailableChannelRange AccessPointConfigurationRetrieveGet200ResponseWifi50AvailableChannelRangeSlice `json:"availableChannelRange,omitempty"`
		Channel               *int                                                                           `json:"channel,omitempty"`
		ChannelRange          AccessPointConfigurationRetrieveGet200ResponseWifi50ChannelRangeSlice          `json:"channelRange,omitempty"`
		ChannelWidth          *int                                                                           `json:"channelWidth,omitempty"`
		SecondaryChannel      *float64                                                                       `json:"secondaryChannel,omitempty"`
		TxPower               *string                                                                        `json:"txPower,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseWlanGroup24 struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200ResponseWlanGroup50 struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationRetrieveGet200Response struct {
		AdministrativeState       *string                                                                 `json:"administrativeState,omitempty"`
		Altitude                  *AccessPointConfigurationRetrieveGet200ResponseAltitude                 `json:"altitude,omitempty"`
		ApGroupID                 *string                                                                 `json:"apGroupId,omitempty"`
		ApMgmtVlan                *AccessPointConfigurationRetrieveGet200ResponseApMgmtVlan               `json:"apMgmtVlan,omitempty"`
		AutoChannelSelection24    *AccessPointConfigurationRetrieveGet200ResponseAutoChannelSelection24   `json:"autoChannelSelection24,omitempty"`
		AutoChannelSelection50    *AccessPointConfigurationRetrieveGet200ResponseAutoChannelSelection50   `json:"autoChannelSelection50,omitempty"`
		AwsVenue                  *string                                                                 `json:"awsVenue,omitempty"`
		BonjourGateway            *AccessPointConfigurationRetrieveGet200ResponseBonjourGateway           `json:"bonjourGateway,omitempty"`
		ChannelEvaluationInterval *float64                                                                `json:"channelEvaluationInterval,omitempty"`
		ClientAdmissionControl24  *AccessPointConfigurationRetrieveGet200ResponseClientAdmissionControl24 `json:"clientAdmissionControl24,omitempty"`
		ClientAdmissionControl50  *AccessPointConfigurationRetrieveGet200ResponseClientAdmissionControl50 `json:"clientAdmissionControl50,omitempty"`
		Description               *string                                                                 `json:"description,omitempty"`
		GpsSource                 *string                                                                 `json:"gpsSource,omitempty"`
		Latitude                  *float64                                                                `json:"latitude,omitempty"`
		Location                  *string                                                                 `json:"location,omitempty"`
		LocationAdditionalInfo    *string                                                                 `json:"locationAdditionalInfo,omitempty"`
		Login                     *AccessPointConfigurationRetrieveGet200ResponseLogin                    `json:"login,omitempty"`
		Longitude                 *float64                                                                `json:"longitude,omitempty"`
		Mac                       *string                                                                 `json:"mac,omitempty"`
		MeshOptions               *AccessPointConfigurationRetrieveGet200ResponseMeshOptions              `json:"meshOptions,omitempty"`
		Model                     *string                                                                 `json:"model,omitempty"`
		Name                      *string                                                                 `json:"name,omitempty"`
		Network                   *AccessPointConfigurationRetrieveGet200ResponseNetwork                  `json:"network,omitempty"`
		ProvisionChecklist        *string                                                                 `json:"provisionChecklist,omitempty"`
		Serial                    *string                                                                 `json:"serial,omitempty"`
		SmartMonitor              *AccessPointConfigurationRetrieveGet200ResponseSmartMonitor             `json:"smartMonitor,omitempty"`
		Specific                  *AccessPointConfigurationRetrieveGet200ResponseSpecific                 `json:"specific,omitempty"`
		Syslog                    *AccessPointConfigurationRetrieveGet200ResponseSyslog                   `json:"syslog,omitempty"`
		VenueProfile              *AccessPointConfigurationRetrieveGet200ResponseVenueProfile             `json:"venueProfile,omitempty"`
		Wifi24                    *AccessPointConfigurationRetrieveGet200ResponseWifi24                   `json:"wifi24,omitempty"`
		Wifi50                    *AccessPointConfigurationRetrieveGet200ResponseWifi50                   `json:"wifi50,omitempty"`
		WlanGroup24               *AccessPointConfigurationRetrieveGet200ResponseWlanGroup24              `json:"wlanGroup24,omitempty"`
		WlanGroup50               *AccessPointConfigurationRetrieveGet200ResponseWlanGroup50              `json:"wlanGroup50,omitempty"`
		WlanService24Enabled      *bool                                                                   `json:"wlanService24Enabled,omitempty"`
		WlanService50Enabled      *bool                                                                   `json:"wlanService50Enabled,omitempty"`
		ZoneID                    *string                                                                 `json:"zoneId,omitempty"`
	}
)

// AccessPointConfigurationRetrieveGet: Use this API command to retrieve the configuration of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointConfigurationRetrieveGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationRetrieveGet(ctx context.Context, apMac string) (*http.Response, *AccessPointConfigurationRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}", true)
	request.SetPathParameter("apMac", apMac)
	out := &AccessPointConfigurationRetrieveGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	AccessPointConfigurationModifyBasicPatchRequest struct {
		AdministrativeState       *string  `json:"administrativeState,omitempty"`
		ApGroupID                 *string  `json:"apGroupId,omitempty"`
		AwsVenue                  *string  `json:"awsVenue,omitempty"`
		ChannelEvaluationInterval *float64 `json:"channelEvaluationInterval,omitempty"`
		Description               *string  `json:"description,omitempty"`
		GpsSource                 *string  `json:"gpsSource,omitempty"`
		Latitude                  *float64 `json:"latitude,omitempty"`
		Location                  *string  `json:"location,omitempty"`
		LocationAdditionalInfo    *string  `json:"locationAdditionalInfo,omitempty"`
		Longitude                 *float64 `json:"longitude,omitempty"`
		Model                     *string  `json:"model,omitempty"`
		Name                      *string  `json:"name,omitempty"`
		ProvisionChecklist        *string  `json:"provisionChecklist,omitempty"`
		Serial                    *string  `json:"serial,omitempty"`
		WlanService24Enabled      *bool    `json:"wlanService24Enabled,omitempty"`
		WlanService50Enabled      *bool    `json:"wlanService50Enabled,omitempty"`
		ZoneID                    *string  `json:"zoneId,omitempty"`
	}
)

// AccessPointConfigurationModifyBasicPatch: Use this API command to modify the basic information of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyBasicPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableAltitudeOverrideDelete: Use this API command to disable AP level override of altitude. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableAltitudeOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/altitude", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyAltitudeOverridePatchRequest struct {
		AltitudeUnit  *string  `json:"altitudeUnit,omitempty"`
		AltitudeValue *float64 `json:"altitudeValue,omitempty"`
	}
)

// AccessPointConfigurationModifyAltitudeOverridePatch: Use this API command to modify the altitude of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyAltitudeOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyAltitudeOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyAltitudeOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/altitude", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableApManagementVlanOverrideDelete: Disable AP Management Vlan Override of an AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableApManagementVlanOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/apMgmtVlan", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyApManagementVlanPatchRequest struct {
		ID   *float64 `json:"id,omitempty"`
		Mode *string  `json:"mode,omitempty"`
	}
)

// AccessPointConfigurationModifyApManagementVlanPatch: Modify AP Management Vlan of an AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyApManagementVlanPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyApManagementVlanPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyApManagementVlanPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/apMgmtVlan", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio24gAutoChannelSelectionDelete: Use this API command to disable the AP level override of auto channel selection on the 2.4GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio24gAutoChannelSelectionDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/autoChannelSelection24", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyRadio24gAutoChannelselectmodeOverridePatchRequest struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}
)

// AccessPointConfigurationModifyRadio24gAutoChannelselectmodeOverridePatch: Override Radio 2.4G Auto ChannelSelectMode and ChannelFly MTBC of an AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyRadio24gAutoChannelselectmodeOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyRadio24gAutoChannelselectmodeOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyRadio24gAutoChannelselectmodeOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/autoChannelSelection24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio5gAutoChannelSelectionDelete: Use this API command to disable the AP level override of auto channel selection on the 5GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio5gAutoChannelSelectionDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/autoChannelSelection50", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyRadio5gAutoChannelselectmodeOverridePatchRequest struct {
		ChannelFlyMtbc    *float64 `json:"channelFlyMtbc,omitempty"`
		ChannelSelectMode *string  `json:"channelSelectMode,omitempty"`
	}
)

// AccessPointConfigurationModifyRadio5gAutoChannelselectmodeOverridePatch: Override Radio 5G Auto ChannelSelectMode and ChannelFly MTBC of an AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyRadio5gAutoChannelselectmodeOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyRadio5gAutoChannelselectmodeOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyRadio5gAutoChannelselectmodeOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/autoChannelSelection50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	BonjourFencingPolicyGetBonjourFencingStatisticGet200ResponseServiceListSlice []*BonjourFencingPolicyGetBonjourFencingStatisticGet200ResponseServiceList

	BonjourFencingPolicyGetBonjourFencingStatisticGet200ResponseServiceList struct {
		NeighborApMac *string `json:"neighborApMac,omitempty"`
		ServiceType   *string `json:"serviceType,omitempty"`
		SourceType    *string `json:"sourceType,omitempty"`
	}

	BonjourFencingPolicyGetBonjourFencingStatisticGet200Response struct {
		ApMac                          *string                                                                      `json:"apMac,omitempty"`
		DroppedPacketsDueToNeighbor    *float64                                                                     `json:"droppedPacketsDueToNeighbor,omitempty"`
		DroppedPacketsDueToServiceType *float64                                                                     `json:"droppedPacketsDueToServiceType,omitempty"`
		ForwardedPackets               *float64                                                                     `json:"forwardedPackets,omitempty"`
		ServiceList                    BonjourFencingPolicyGetBonjourFencingStatisticGet200ResponseServiceListSlice `json:"serviceList,omitempty"`
	}
)

// BonjourFencingPolicyGetBonjourFencingStatisticGet: Use this API command to get Bonjour Fencing Statistic per AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *BonjourFencingPolicyGetBonjourFencingStatisticGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) BonjourFencingPolicyGetBonjourFencingStatisticGet(ctx context.Context, apMac string) (*http.Response, *BonjourFencingPolicyGetBonjourFencingStatisticGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/bonjourFencingStatistic", true)
	request.SetPathParameter("apMac", apMac)
	out := &BonjourFencingPolicyGetBonjourFencingStatisticGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// AccessPointConfigurationDisableBonjourGatewayOverrideDelete: Use this API command to disable AP level override of bonjour gateway. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableBonjourGatewayOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/bonjourGateway", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyBonjourGatewayPatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// AccessPointConfigurationModifyBonjourGatewayPatch: Use this API command to modify bonjour gateway of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyBonjourGatewayPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyBonjourGatewayPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyBonjourGatewayPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/bonjourGateway", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableApChannelEvaluationIntervalDelete: Disable AP lChannel Evaluation Interval. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableApChannelEvaluationIntervalDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/channelEvaluationInterval", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableClientAdmissionControl24gOverrideDelete: Use this API command to disable AP level override of client admission control 2.4GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableClientAdmissionControl24gOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/clientAdmissionControl24", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyClientAdmissionControl24gPatchRequest struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}
)

// AccessPointConfigurationModifyClientAdmissionControl24gPatch: Use this API command to modify client admission control 2.4GHz radio configuration of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyClientAdmissionControl24gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyClientAdmissionControl24gPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyClientAdmissionControl24gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/clientAdmissionControl24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableClientAdmissionControl5gOverrideDelete: Use this API command to disable AP level override of client admission control 5GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableClientAdmissionControl5gOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/clientAdmissionControl50", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyClientAdmissionControl5gPatchRequest struct {
		Enabled                 *bool    `json:"enabled,omitempty"`
		MaxRadioLoadPercent     *float64 `json:"maxRadioLoadPercent,omitempty"`
		MinClientCount          *float64 `json:"minClientCount,omitempty"`
		MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
	}
)

// AccessPointConfigurationModifyClientAdmissionControl5gPatch: Use this API command to modify client admission control 5GHz radio configuration of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyClientAdmissionControl5gPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyClientAdmissionControl5gPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyClientAdmissionControl5gPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/clientAdmissionControl50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableApManagementGpsCooordinatesDelete: Disable AP Management GPS Cooordinates of an AP
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableApManagementGpsCooordinatesDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/gpsCoordinates", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableLocationOverrideDelete: Use this API command to disable AP level override of location. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableLocationOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/location", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableLocationAdditionalinfoOverrideDelete: Use this API command to disable AP level override of location additionalInfo. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableLocationAdditionalinfoOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/locationAdditionalInfo", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableLoginOverrideDelete: Use this API command to disable the AP-level logon override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableLoginOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/login", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyLoginOverridePatchRequest struct {
		ApLoginName     *string `json:"apLoginName,omitempty"`
		ApLoginPassword *string `json:"apLoginPassword,omitempty"`
	}
)

// AccessPointConfigurationModifyLoginOverridePatch: Use this API command to enable or modify the AP-level logon override settings.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyLoginOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyLoginOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyLoginOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/login", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableMeshOptionsDelete: Use this API command to disable mesh options.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableMeshOptionsDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/meshOptions", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyMeshOptionsPatchRequestMeshUplinkEntryListSlice []*string

	AccessPointConfigurationModifyMeshOptionsPatchRequest struct {
		MeshMode            *string                                                                       `json:"meshMode,omitempty"`
		MeshUplinkEntryList AccessPointConfigurationModifyMeshOptionsPatchRequestMeshUplinkEntryListSlice `json:"meshUplinkEntryList,omitempty"`
		UplinkSelection     *string                                                                       `json:"uplinkSelection,omitempty"`
	}
)

// AccessPointConfigurationModifyMeshOptionsPatch: Use this API command to modify mesh options of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyMeshOptionsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyMeshOptionsPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyMeshOptionsPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/meshOptions", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyNetworkSettingsPatchRequest struct {
		Gateway      *string `json:"gateway,omitempty"`
		IP           *string `json:"ip,omitempty"`
		IPType       *string `json:"ipType,omitempty"`
		Netmask      *string `json:"netmask,omitempty"`
		PrimaryDNS   *string `json:"primaryDns,omitempty"`
		SecondaryDNS *string `json:"secondaryDns,omitempty"`
	}
)

// AccessPointConfigurationModifyNetworkSettingsPatch: Use this API command to modify the network settings of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyNetworkSettingsPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyNetworkSettingsPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyNetworkSettingsPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/network", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointOperationalRetrieveAlarmListGet200ResponseListSlice []*AccessPointOperationalRetrieveAlarmListGet200ResponseList

	AccessPointOperationalRetrieveAlarmListGet200ResponseList struct {
		AcknowledgedTime *float64 `json:"acknowledgedTime,omitempty"`
		Category         *string  `json:"category,omitempty"`
		ClearedTime      *float64 `json:"clearedTime,omitempty"`
		Code             *float64 `json:"code,omitempty"`
		Description      *string  `json:"description,omitempty"`
		ID               *string  `json:"id,omitempty"`
		Severity         *string  `json:"severity,omitempty"`
		Status           *string  `json:"status,omitempty"`
		Time             *float64 `json:"time,omitempty"`
		Type             *string  `json:"type,omitempty"`
	}

	AccessPointOperationalRetrieveAlarmListGet200Response struct {
		FirstIndex *int                                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                                          `json:"hasMore,omitempty"`
		List       AccessPointOperationalRetrieveAlarmListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                           `json:"totalCount,omitempty"`
	}
)

// AccessPointOperationalRetrieveAlarmListGet: Use this API command to retrieve the list of alarms on an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - severity (string): filter by severity
// - category (string): filter by category
// - code (number): filter by alarm code
// - startTime (string): filter by trigger time later than startTime (milliseconds in UTC time)
// - endTime (string): filter by trigger time earlier than endTime (milliseconds in UTC time)
// - status (string): filter by alarm status
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveAlarmListGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalRetrieveAlarmListGet(ctx context.Context, apMac string, optionalParams map[string]string) (*http.Response, *AccessPointOperationalRetrieveAlarmListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
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
	severity, ok := optionalParams["severity"]
	if ok {
		err = validators.InRange([]string{"Critical", "Major", "Minor", "Warning", "Informational", "Debug"}, severity)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"severity\" failed validation check: %s", err)
		}
	}
	status, ok := optionalParams["status"]
	if ok {
		err = validators.InRange([]string{"Outstanding", "Acknowledged", "Cleared"}, status)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"status\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/alarms", true)
	request.SetPathParameter("apMac", apMac)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("severity", severity)
	request.SetQueryParameter("category", optionalParams["category"])
	request.SetQueryParameter("code", optionalParams["code"])
	request.SetQueryParameter("startTime", optionalParams["startTime"])
	request.SetQueryParameter("endTime", optionalParams["endTime"])
	request.SetQueryParameter("status", status)
	out := &AccessPointOperationalRetrieveAlarmListGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	AccessPointOperationalRetrieveAlarmSummaryGet200Response struct {
		CriticalCount *int `json:"criticalCount,omitempty"`
		MajorCount    *int `json:"majorCount,omitempty"`
		MinorCount    *int `json:"minorCount,omitempty"`
		WarningCount  *int `json:"warningCount,omitempty"`
	}
)

// AccessPointOperationalRetrieveAlarmSummaryGet: Use this API command to retrieve the alarm summary of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveAlarmSummaryGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalRetrieveAlarmSummaryGet(ctx context.Context, apMac string) (*http.Response, *AccessPointOperationalRetrieveAlarmSummaryGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/alarmSummary", true)
	request.SetPathParameter("apMac", apMac)
	out := &AccessPointOperationalRetrieveAlarmSummaryGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// AccessPointOperationalApBlinkLedPost: use this API to make ap blink its led to show its position
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalApBlinkLedPost(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("POST", "/v5_0/aps/{apMac}/operational/blinkLed", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 200, nil)
}

type (
	WirelessClientRetrieveClientListGet200ResponseListSlice []*WirelessClientRetrieveClientListGet200ResponseList

	WirelessClientRetrieveClientListGet200ResponseList struct {
		ApRxSignal          *string  `json:"apRxSignal,omitempty"`
		ApTxDataRate        *string  `json:"apTxDataRate,omitempty"`
		Channel             *string  `json:"channel,omitempty"`
		ConnectedSince      *float64 `json:"connectedSince,omitempty"`
		FromClientBytes     *float64 `json:"fromClientBytes,omitempty"`
		FromClientPkts      *float64 `json:"fromClientPkts,omitempty"`
		HostName            *string  `json:"hostName,omitempty"`
		IPAddress           *string  `json:"ipAddress,omitempty"`
		Ipv6Address         *string  `json:"ipv6Address,omitempty"`
		Mac                 *string  `json:"mac,omitempty"`
		OsType              *string  `json:"osType,omitempty"`
		RadioID             *string  `json:"radioId,omitempty"`
		RadioMode           *string  `json:"radioMode,omitempty"`
		Rssi                *string  `json:"rssi,omitempty"`
		RxAvgByteRate       *float64 `json:"rxAvgByteRate,omitempty"`
		RxByteRate          *float64 `json:"rxByteRate,omitempty"`
		Ssid                *string  `json:"ssid,omitempty"`
		Status              *string  `json:"status,omitempty"`
		ToClientBytes       *float64 `json:"toClientBytes,omitempty"`
		ToClientDroppedPkts *float64 `json:"toClientDroppedPkts,omitempty"`
		ToClientPkts        *float64 `json:"toClientPkts,omitempty"`
		TxAvgByteRate       *float64 `json:"txAvgByteRate,omitempty"`
		TxByteRate          *float64 `json:"txByteRate,omitempty"`
		User                *string  `json:"user,omitempty"`
		Vlan                *string  `json:"vlan,omitempty"`
		WlanID              *string  `json:"wlanId,omitempty"`
	}

	WirelessClientRetrieveClientListGet200Response struct {
		FirstIndex *int                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                   `json:"hasMore,omitempty"`
		List       WirelessClientRetrieveClientListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                    `json:"totalCount,omitempty"`
	}
)

// WirelessClientRetrieveClientListGet: Use this API command to retrieve the client list per AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *WirelessClientRetrieveClientListGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) WirelessClientRetrieveClientListGet(ctx context.Context, apMac string, optionalParams map[string]string) (*http.Response, *WirelessClientRetrieveClientListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
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
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/client", true)
	request.SetPathParameter("apMac", apMac)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &WirelessClientRetrieveClientListGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// WirelessClientRetrieveTotalClientCountGet: Use this API command to retrieve the total client count per AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) WirelessClientRetrieveTotalClientCountGet(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/client/totalCount", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 200, nil)
}

type (
	AccessPointOperationalRetrieveEventListGet200ResponseListSlice []*AccessPointOperationalRetrieveEventListGet200ResponseList

	AccessPointOperationalRetrieveEventListGet200ResponseList struct {
		Category    *string  `json:"category,omitempty"`
		Code        *float64 `json:"code,omitempty"`
		Description *string  `json:"description,omitempty"`
		ID          *string  `json:"id,omitempty"`
		Severity    *string  `json:"severity,omitempty"`
		Time        *float64 `json:"time,omitempty"`
		Type        *string  `json:"type,omitempty"`
	}

	AccessPointOperationalRetrieveEventListGet200Response struct {
		FirstIndex *int                                                           `json:"firstIndex,omitempty"`
		HasMore    *bool                                                          `json:"hasMore,omitempty"`
		List       AccessPointOperationalRetrieveEventListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                           `json:"totalCount,omitempty"`
	}
)

// AccessPointOperationalRetrieveEventListGet: Use this API command to retrieve the list of events from an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - severity (string): filter by severity
// - category (string): filter by category
// - code (number): filter by event code
// - startTime (string): filter by trigger time later than startTime (milliseconds in UTC time)
// - endTime (string): filter by trigger time earlier than endTime (milliseconds in UTC time)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveEventListGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalRetrieveEventListGet(ctx context.Context, apMac string, optionalParams map[string]string) (*http.Response, *AccessPointOperationalRetrieveEventListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
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
	severity, ok := optionalParams["severity"]
	if ok {
		err = validators.InRange([]string{"Critical", "Major", "Minor", "Warning", "Informational", "Debug"}, severity)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"severity\" failed validation check: %s", err)
		}
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/events", true)
	request.SetPathParameter("apMac", apMac)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("severity", severity)
	request.SetQueryParameter("category", optionalParams["category"])
	request.SetQueryParameter("code", optionalParams["code"])
	request.SetQueryParameter("startTime", optionalParams["startTime"])
	request.SetQueryParameter("endTime", optionalParams["endTime"])
	out := &AccessPointOperationalRetrieveEventListGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	AccessPointOperationalRetrieveEventSummaryGet200Response struct {
		CriticalCount      *int `json:"criticalCount,omitempty"`
		DebugCount         *int `json:"debugCount,omitempty"`
		InformationalCount *int `json:"informationalCount,omitempty"`
		MajorCount         *int `json:"majorCount,omitempty"`
		MinorCount         *int `json:"minorCount,omitempty"`
		WarningCount       *int `json:"warningCount,omitempty"`
	}
)

// AccessPointOperationalRetrieveEventSummaryGet: Use this API command to retrieve the event summary of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveEventSummaryGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalRetrieveEventSummaryGet(ctx context.Context, apMac string) (*http.Response, *AccessPointOperationalRetrieveEventSummaryGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/eventSummary", true)
	request.SetPathParameter("apMac", apMac)
	out := &AccessPointOperationalRetrieveEventSummaryGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	AccessPointOperationalRetrieveMeshNeighborApListGet200ResponseListSlice []*AccessPointOperationalRetrieveMeshNeighborApListGet200ResponseList

	AccessPointOperationalRetrieveMeshNeighborApListGet200ResponseList struct {
		Channel         *string `json:"channel,omitempty"`
		ConnectionState *string `json:"connectionState,omitempty"`
		ExternalIP      *string `json:"externalIp,omitempty"`
		ExternalPort    *string `json:"externalPort,omitempty"`
		IP              *string `json:"ip,omitempty"`
		Mac             *string `json:"mac,omitempty"`
		Model           *string `json:"model,omitempty"`
		Name            *string `json:"name,omitempty"`
		Signal          *string `json:"signal,omitempty"`
		Version         *string `json:"version,omitempty"`
		ZoneName        *string `json:"zoneName,omitempty"`
	}

	AccessPointOperationalRetrieveMeshNeighborApListGet200Response struct {
		FirstIndex *int                                                                    `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                   `json:"hasMore,omitempty"`
		List       AccessPointOperationalRetrieveMeshNeighborApListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                    `json:"totalCount,omitempty"`
	}
)

// AccessPointOperationalRetrieveMeshNeighborApListGet: Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveMeshNeighborApListGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalRetrieveMeshNeighborApListGet(ctx context.Context, apMac string, optionalParams map[string]string) (*http.Response, *AccessPointOperationalRetrieveMeshNeighborApListGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
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
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/neighbor", true)
	request.SetPathParameter("apMac", apMac)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	out := &AccessPointOperationalRetrieveMeshNeighborApListGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	AccessPointOperationalRetrieveOperationalInformationGet200Response struct {
		AdministrativeState *string     `json:"administrativeState,omitempty"`
		ApGroupID           *string     `json:"apGroupId,omitempty"`
		ApprovedTime        *float64    `json:"approvedTime,omitempty"`
		ClientCount         *float64    `json:"clientCount,omitempty"`
		ConfigState         *string     `json:"configState,omitempty"`
		ConnectionState     *string     `json:"connectionState,omitempty"`
		CountryCode         *string     `json:"countryCode,omitempty"`
		CpID                *string     `json:"cpId,omitempty"`
		Description         *string     `json:"description,omitempty"`
		DpID                *string     `json:"dpId,omitempty"`
		ExternalIP          *string     `json:"externalIp,omitempty"`
		ExternalPort        *float64    `json:"externalPort,omitempty"`
		IP                  *string     `json:"ip,omitempty"`
		IPType              *string     `json:"ipType,omitempty"`
		Ipv6                interface{} `json:"ipv6,omitempty"`
		Ipv6Type            interface{} `json:"ipv6Type,omitempty"`
		IsCriticalAP        *bool       `json:"isCriticalAP,omitempty"`
		LastSeenTime        *float64    `json:"lastSeenTime,omitempty"`
		Latitude            *float64    `json:"latitude,omitempty"`
		Location            *string     `json:"location,omitempty"`
		Longitude           *float64    `json:"longitude,omitempty"`
		Mac                 *string     `json:"mac,omitempty"`
		ManagementVlan      *float64    `json:"managementVlan,omitempty"`
		MeshHop             *float64    `json:"meshHop,omitempty"`
		MeshRole            *string     `json:"meshRole,omitempty"`
		Model               *string     `json:"model,omitempty"`
		Name                *string     `json:"name,omitempty"`
		ProvisionMethod     *string     `json:"provisionMethod,omitempty"`
		ProvisionStage      *string     `json:"provisionStage,omitempty"`
		RegistrationState   *string     `json:"registrationState,omitempty"`
		Serial              *string     `json:"serial,omitempty"`
		Uptime              *float64    `json:"uptime,omitempty"`
		Version             *string     `json:"version,omitempty"`
		Wifi24Channel       *string     `json:"wifi24Channel,omitempty"`
		Wifi50Channel       *string     `json:"wifi50Channel,omitempty"`
		ZoneID              *string     `json:"zoneId,omitempty"`
	}
)

// AccessPointOperationalRetrieveOperationalInformationGet: Use this API command to retrieve the operational information of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveOperationalInformationGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointOperationalRetrieveOperationalInformationGet(ctx context.Context, apMac string) (*http.Response, *AccessPointOperationalRetrieveOperationalInformationGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/operational/summary", true)
	request.SetPathParameter("apMac", apMac)
	out := &AccessPointOperationalRetrieveOperationalInformationGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// AccessPointConfigurationRebootPut: reboot an access point
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationRebootPut(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/aps/{apMac}/reboot", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableSmartMonitorOverrideDelete: Use this API command to disable AP level override of smart monitor. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableSmartMonitorOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/smartMonitor", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifySmartMonitorPatchRequest struct {
		Enabled        *bool `json:"enabled,omitempty"`
		IntervalInSec  *int  `json:"intervalInSec,omitempty"`
		RetryThreshold *int  `json:"retryThreshold,omitempty"`
	}
)

// AccessPointConfigurationModifySmartMonitorPatch: Use this API command to modify smart monitor of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifySmartMonitorPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifySmartMonitorPatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifySmartMonitorPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/smartMonitor", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableModelSpecificOverrideDelete: Use this API command to disable model specific configuration override from AP group or zone.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableModelSpecificOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/specific", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyModelSpecificPutRequestExternalAntenna24 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	AccessPointConfigurationModifyModelSpecificPutRequestExternalAntenna50 struct {
		ChainMask interface{} `json:"chainMask,omitempty"`
		Dbi       *float64    `json:"dbi,omitempty"`
		Enabled   *bool       `json:"enabled,omitempty"`
	}

	AccessPointConfigurationModifyModelSpecificPutRequestLanPortsSlice []*AccessPointConfigurationModifyModelSpecificPutRequestLanPorts

	AccessPointConfigurationModifyModelSpecificPutRequestLanPortsEthPortProfile struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationModifyModelSpecificPutRequestLanPorts struct {
		Enabled              *bool                                                                        `json:"enabled,omitempty"`
		EthPortProfile       *AccessPointConfigurationModifyModelSpecificPutRequestLanPortsEthPortProfile `json:"ethPortProfile,omitempty"`
		Members              *string                                                                      `json:"members,omitempty"`
		OverwriteVlanEnabled *bool                                                                        `json:"overwriteVlanEnabled,omitempty"`
		PortName             *string                                                                      `json:"portName,omitempty"`
		VlanUntagID          *float64                                                                     `json:"vlanUntagId,omitempty"`
	}

	AccessPointConfigurationModifyModelSpecificPutRequestLldp struct {
		AdvertiseIntervalInSec *int  `json:"advertiseIntervalInSec,omitempty"`
		Enabled                *bool `json:"enabled,omitempty"`
		HoldTimeInSec          *int  `json:"holdTimeInSec,omitempty"`
		ManagementIPTLVEnabled *bool `json:"managementIPTLVEnabled,omitempty"`
	}

	AccessPointConfigurationModifyModelSpecificPutRequest struct {
		ExternalAntenna24     *AccessPointConfigurationModifyModelSpecificPutRequestExternalAntenna24 `json:"externalAntenna24,omitempty"`
		ExternalAntenna50     *AccessPointConfigurationModifyModelSpecificPutRequestExternalAntenna50 `json:"externalAntenna50,omitempty"`
		InternalHeaterEnabled *bool                                                                   `json:"internalHeaterEnabled,omitempty"`
		LanPorts              AccessPointConfigurationModifyModelSpecificPutRequestLanPortsSlice      `json:"lanPorts,omitempty"`
		LedMode               interface{}                                                             `json:"ledMode,omitempty"`
		LedStatusEnabled      *bool                                                                   `json:"ledStatusEnabled,omitempty"`
		Lldp                  *AccessPointConfigurationModifyModelSpecificPutRequestLldp              `json:"lldp,omitempty"`
		PoeModeSetting        *string                                                                 `json:"poeModeSetting,omitempty"`
		PoeOutPortEnabled     *bool                                                                   `json:"poeOutPortEnabled,omitempty"`
		PoeTxChain            *int                                                                    `json:"poeTxChain,omitempty"`
		RadioBand             interface{}                                                             `json:"radioBand,omitempty"`
		UsbPowerEnable        *bool                                                                   `json:"usbPowerEnable,omitempty"`
	}
)

// AccessPointConfigurationModifyModelSpecificPut: Use this API command to modify model specific configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyModelSpecificPutRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyModelSpecificPut(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyModelSpecificPutRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PUT", "/v5_0/aps/{apMac}/specific", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDownloadApSupportLogGet: Use this API command to download AP support log.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDownloadApSupportLogGet(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/supportLog", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 200, nil)
}

// AccessPointConfigurationDisableSyslogOverrideDelete: Use this API command to disable the AP level syslog override. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableSyslogOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/syslog", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifySyslogOverridePatchRequest struct {
		Address  *string `json:"address,omitempty"`
		Enabled  *bool   `json:"enabled,omitempty"`
		Facility *string `json:"facility,omitempty"`
		Port     *int    `json:"port,omitempty"`
		Priority *string `json:"priority,omitempty"`
	}
)

// AccessPointConfigurationModifySyslogOverridePatch: Use this API command to enable or modify the AP-level syslog override settings.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifySyslogOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifySyslogOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifySyslogOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/syslog", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	ApUsbSoftwarePackageGetApAssociateGet200Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApUsbSoftwarePackageGetApAssociateGet: Get APUsbSoftwarePackage associate AP by model name
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApUsbSoftwarePackageGetApAssociateGet200Response
// - error: Error seen or nil if none
func (a *APsAPI) ApUsbSoftwarePackageGetApAssociateGet(ctx context.Context, apMac string) (*http.Response, *ApUsbSoftwarePackageGetApAssociateGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("GET", "/v5_0/aps/{apMac}/usbsoftware", true)
	request.SetPathParameter("apMac", apMac)
	out := &ApUsbSoftwarePackageGetApAssociateGet200Response{}
	httpResponse, _, err := a.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// AccessPointConfigurationDisableApUsbSoftwarePackageDelete: Disable AP level Usb Software Package. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableApUsbSoftwarePackageDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/usbSoftwarePackage", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequestApplyModelSlice []*string

	AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequestUsbSoftware struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequest struct {
		ApplyModel  AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequestApplyModelSlice `json:"applyModel,omitempty"`
		UsbSoftware *AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequestUsbSoftware    `json:"usbSoftware,omitempty"`
	}
)

// AccessPointConfigurationModifyApUsbSoftwarePackagePatch: Modify AP Usb Software Package of an access point
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyApUsbSoftwarePackagePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyApUsbSoftwarePackagePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/usbSoftwarePackage", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableVenueProfileOverrideDelete: Use this API command to disable AP level override of venue profile. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableVenueProfileOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/venueProfile", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyVenueProfilePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// AccessPointConfigurationModifyVenueProfilePatch: Use this API command to modify venue profile of an AP.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyVenueProfilePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyVenueProfilePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyVenueProfilePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/venueProfile", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio24gOverrideDelete: Use this API command to disable the AP level override of the 2.4GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio24gOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi24", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyRadio24gOverridePatchRequestChannelRangeSlice []*float64

	AccessPointConfigurationModifyRadio24gOverridePatchRequest struct {
		Channel      *int                                                                        `json:"channel,omitempty"`
		ChannelRange AccessPointConfigurationModifyRadio24gOverridePatchRequestChannelRangeSlice `json:"channelRange,omitempty"`
		ChannelWidth *int                                                                        `json:"channelWidth,omitempty"`
		TxPower      *string                                                                     `json:"txPower,omitempty"`
	}
)

// AccessPointConfigurationModifyRadio24gOverridePatch: Use this API command to modify the AP level override of the 2.4GHz radio configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyRadio24gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyRadio24gOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyRadio24gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/wifi24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio24gChannelOverrideDelete: Use this API command to disable the AP level override of the 2.4GHz radio channel. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio24gChannelOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi24/channel", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio24gChannelrangeOverrideDelete: Use this API command to disable the AP level override of the 2.4GHz radio channelRange. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio24gChannelrangeOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi24/channelRange", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio24gChannelwidthOverrideDelete: Use this API command to disable the AP level override of the 2.4GHz radio channelWidth. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio24gChannelwidthOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi24/channelWidth", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio24gTxpowerOverrideDelete: Use this API command to disable the AP level override of the 2.4GHz radio txPower. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio24gTxpowerOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi24/txPower", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio5gOverrideDelete: Use this API command to disable the AP level override of 5GHz radio configuration. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio5gOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi50", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyRadio5gOverridePatchRequestChannelRangeSlice []*float64

	AccessPointConfigurationModifyRadio5gOverridePatchRequest struct {
		Channel          *int                                                                       `json:"channel,omitempty"`
		ChannelRange     AccessPointConfigurationModifyRadio5gOverridePatchRequestChannelRangeSlice `json:"channelRange,omitempty"`
		ChannelWidth     *int                                                                       `json:"channelWidth,omitempty"`
		SecondaryChannel *float64                                                                   `json:"secondaryChannel,omitempty"`
		TxPower          *string                                                                    `json:"txPower,omitempty"`
	}
)

// AccessPointConfigurationModifyRadio5gOverridePatch: Use this API command to Modify the AP level override of the 5GHz radio configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyRadio5gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyRadio5gOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyRadio5gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/wifi50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio5gChannelOverrideDelete: Use this API command to disable the AP level override of 5GHz radio channel. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio5gChannelOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi50/channel", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio5gChannelrangeOverrideDelete: Use this API command to disable the AP level override of 5GHz radio channelRange. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio5gChannelrangeOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi50/channelRange", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio5gChannelwidthOverrideDelete: Use this API command to disable the AP level override of 5GHz radio channelWidth. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio5gChannelwidthOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi50/channelWidth", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableRadio5gTxpowerOverrideDelete: Use this API command to disable the AP level override of 5GHz radio txPower. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableRadio5gTxpowerOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wifi50/txPower", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableWlanGroup24gOverrideDelete: Use this API command to disable the AP level override of WLAN group configuration on 2.4GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableWlanGroup24gOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wlanGroup24", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyWlanGroup24gOverridePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// AccessPointConfigurationModifyWlanGroup24gOverridePatch: Use this API command to enable or modify the AP level override of the WLAN group configuration on the 2.4GHz radio.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyWlanGroup24gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyWlanGroup24gOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyWlanGroup24gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/wlanGroup24", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

// AccessPointConfigurationDisableWlanGroup5gOverrideDelete: Use this API command to disable the AP level override of WLAN group on the 5GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationDisableWlanGroup5gOverrideDelete(ctx context.Context, apMac string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("DELETE", "/v5_0/aps/{apMac}/wlanGroup50", true)
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}

type (
	AccessPointConfigurationModifyWlanGroup5gOverridePatchRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// AccessPointConfigurationModifyWlanGroup5gOverridePatch: Use this API command to enable or modify the AP level override of the WLAN group configuration on the 5GHz radio.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - apMac (MAC): Access Point MAC Address
// - requestBody: *AccessPointConfigurationModifyWlanGroup5gOverridePatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *APsAPI) AccessPointConfigurationModifyWlanGroup5gOverridePatch(ctx context.Context, apMac string, requestBody *AccessPointConfigurationModifyWlanGroup5gOverridePatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrIsMAC(apMac)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"apMac\" failed validation check: %s", err)
	}
	request := NewRequest("PATCH", "/v5_0/aps/{apMac}/wlanGroup50", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	request.SetPathParameter("apMac", apMac)
	return a.client.Ensure(ctx, request, 204, nil)
}
