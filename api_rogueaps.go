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

type RogueAPs struct {
	client *Client
}
type (
	AccessPointOperationalRetrieveRogueApListGet200ResponseListSlice []*AccessPointOperationalRetrieveRogueApListGet200ResponseList

	AccessPointOperationalRetrieveRogueApListGet200ResponseList struct {
		ApMac      *string `json:"apMac,omitempty"`
		ApName     *string `json:"apName,omitempty"`
		Channel    *string `json:"channel,omitempty"`
		Encryption *string `json:"encryption,omitempty"`
		ID         *string `json:"id,omitempty"`
		Radio      *string `json:"radio,omitempty"`
		RogueAPMac *string `json:"rogueAPMac,omitempty"`
		RogueMac   *string `json:"rogueMac,omitempty"`
		Rssi       *string `json:"rssi,omitempty"`
		Ssid       *string `json:"ssid,omitempty"`
		Timestamp  *string `json:"timestamp,omitempty"`
		Type       *string `json:"type,omitempty"`
	}

	AccessPointOperationalRetrieveRogueApListGet200Response struct {
		FirstIndex *int                                                             `json:"firstIndex,omitempty"`
		HasMore    *bool                                                            `json:"hasMore,omitempty"`
		List       AccessPointOperationalRetrieveRogueApListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                             `json:"totalCount,omitempty"`
	}
)

// AccessPointOperationalRetrieveRogueApListGet: Use this API command to retrieve a list of rogue access points.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - type (string): Filter rogue ap list by type. Valid values are: (ROGUE, MALICIOUS_AP_SSID_SPOOF, MALICIOUS_AP_MAC_SPOOF, MALICIOUS_AP_SAME_NETWORK, AD_HOC).
// - rogueMac (MAC): Filter rogue ap list by Rogue MAC.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *AccessPointOperationalRetrieveRogueApListGet200Response
// - error: Error seen or nil if none
func (r *RogueAPs) AccessPointOperationalRetrieveRogueApListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *AccessPointOperationalRetrieveRogueApListGet200Response, error) {
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
	xtype, ok := optionalParams["type"]
	if ok {
		err = validators.InRange([]string{"ROGUE", "MALICIOUS_AP_SSID_SPOOF", "MALICIOUS_AP_MAC_SPOOF", "MALICIOUS_AP_SAME_NETWORK", "AD_HOC"}, xtype)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"type\" failed validation check: %s", err)
		}
	}
	rogueMac, ok := optionalParams["rogueMac"]
	if ok {
		err = validators.StrIsMAC(rogueMac)
		if nil != err {
			return nil, nil, fmt.Errorf("parameter \"rogueMac\" failed validation check: %s", err)
		}
	}

	request := r.client.newRequest(ctx, "GET", "/v5_0/rogueaps")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
		"type":     xtype,
		"rogueMac": rogueMac,
	}

	out := &AccessPointOperationalRetrieveRogueApListGet200Response{}
	httpResponse, _, err := r.client.doRequest(request, 200, out)
	return httpResponse, out, err
}
