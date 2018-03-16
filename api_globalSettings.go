package api

import (
	"context"
	"errors"
	"net/http"
)

// This file is auto-generated
// Generation Date: 2018-03-16T16:29:52-0500
// API Version: v5

type GlobalSettingsAPI struct {
	client *Client
}
type (
	GlobalReferenceGetFriendlyNameOfUsableLanguageGet200ResponseListSlice []*GlobalReferenceGetFriendlyNameOfUsableLanguageGet200ResponseList

	GlobalReferenceGetFriendlyNameOfUsableLanguageGet200ResponseList struct {
		Display *string `json:"display,omitempty"`
		Value   *string `json:"value,omitempty"`
	}

	GlobalReferenceGetFriendlyNameOfUsableLanguageGet200Response struct {
		FirstIndex        *int                                                                  `json:"firstIndex,omitempty"`
		HasMore           *bool                                                                 `json:"hasMore,omitempty"`
		List              GlobalReferenceGetFriendlyNameOfUsableLanguageGet200ResponseListSlice `json:"list,omitempty"`
		RawDataTotalCount *int                                                                  `json:"rawDataTotalCount,omitempty"`
		TotalCount        *int                                                                  `json:"totalCount,omitempty"`
	}
)

// GlobalReferenceGetFriendlyNameOfUsableLanguageGet: Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *GlobalReferenceGetFriendlyNameOfUsableLanguageGet200Response
// - error: Error seen or nil if none
func (g *GlobalSettingsAPI) GlobalReferenceGetFriendlyNameOfUsableLanguageGet(ctx context.Context) (*http.Response, *GlobalReferenceGetFriendlyNameOfUsableLanguageGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/globalSettings/friendlyNameLang", true)
	out := &GlobalReferenceGetFriendlyNameOfUsableLanguageGet200Response{}
	httpResponse, _, err := g.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200ResponseListSlice []*GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200ResponseList

	GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200ResponseList struct {
		Display *string `json:"display,omitempty"`
		Value   *string `json:"value,omitempty"`
	}

	GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200Response struct {
		FirstIndex        *int                                                                   `json:"firstIndex,omitempty"`
		HasMore           *bool                                                                  `json:"hasMore,omitempty"`
		List              GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200ResponseListSlice `json:"list,omitempty"`
		RawDataTotalCount *int                                                                   `json:"rawDataTotalCount,omitempty"`
		TotalCount        *int                                                                   `json:"totalCount,omitempty"`
	}
)

// GlobalReferenceGetFriendlyNameOfUsableLanguageGet1: Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200Response
// - error: Error seen or nil if none
func (g *GlobalSettingsAPI) GlobalReferenceGetFriendlyNameOfUsableLanguageGet1(ctx context.Context) (*http.Response, *GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/globalSettings/portalLang", true)
	out := &GlobalReferenceGetFriendlyNameOfUsableLanguageGet1200Response{}
	httpResponse, _, err := g.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	UploadStatisticsToFtpRetrieveGet200Response struct {
		Enabled     *bool   `json:"enabled,omitempty"`
		FtpID       *string `json:"ftpId,omitempty"`
		FtpInterval *string `json:"ftpInterval,omitempty"`
	}
)

// UploadStatisticsToFtpRetrieveGet: Use this API command to retrieve the uploading statistical data to FTP server setting.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *UploadStatisticsToFtpRetrieveGet200Response
// - error: Error seen or nil if none
func (g *GlobalSettingsAPI) UploadStatisticsToFtpRetrieveGet(ctx context.Context) (*http.Response, *UploadStatisticsToFtpRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("GET", "/v5_0/globalSettings/statsFtp", true)
	out := &UploadStatisticsToFtpRetrieveGet200Response{}
	httpResponse, _, err := g.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

type (
	UploadStatisticsToFtpModifyBasicPatchRequest struct {
		Enabled     *bool   `json:"enabled,omitempty"`
		FtpID       *string `json:"ftpId,omitempty"`
		FtpInterval *string `json:"ftpInterval,omitempty"`
	}
)

// UploadStatisticsToFtpModifyBasicPatch: Use this API command to modify the setting of uploading statistical data to FTP server.
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
// - requestBody: *UploadStatisticsToFtpModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (g *GlobalSettingsAPI) UploadStatisticsToFtpModifyBasicPatch(ctx context.Context, requestBody *UploadStatisticsToFtpModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	request := NewRequest("PATCH", "/v5_0/globalSettings/statsFtp", true)
	err = request.SetBodyModel(requestBody)
	if err != nil {
		return nil, nil, err
	}
	return g.client.Ensure(ctx, request, 204, nil)
}
