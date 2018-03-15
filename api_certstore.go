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

type CertStoreAPI struct {
	client *Client
}
type (
	CertificateRetrieveCertificateListGet200ResponseListSlice []*CertificateRetrieveCertificateListGet200ResponseList

	CertificateRetrieveCertificateListGet200ResponseList struct {
		Description *string `json:"description,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	CertificateRetrieveCertificateListGet200Response struct {
		FirstIndex *int                                                      `json:"firstIndex,omitempty"`
		HasMore    *bool                                                     `json:"hasMore,omitempty"`
		List       CertificateRetrieveCertificateListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                      `json:"totalCount,omitempty"`
	}
)

// CertificateRetrieveCertificateListGet: Use this API command to retrieve list of installed certificates.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveCertificateListGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveCertificateListGet(ctx *UserContext, optionalParams map[string]string) (*http.Response, *CertificateRetrieveCertificateListGet200Response, error) {
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
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/certificate")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}
	out := &CertificateRetrieveCertificateListGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificateCreateCertificatePostRequestCertificasSigningRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	CertificateCreateCertificatePostRequestIntermediateDataSlice []*string

	CertificateCreateCertificatePostRequest struct {
		CertificasSigningRequest *CertificateCreateCertificatePostRequestCertificasSigningRequest `json:"certificasSigningRequest,omitempty"`
		Data                     *string                                                          `json:"data,omitempty"`
		Description              *string                                                          `json:"description,omitempty"`
		IntermediateData         CertificateCreateCertificatePostRequestIntermediateDataSlice     `json:"intermediateData,omitempty"`
		Name                     *string                                                          `json:"name,omitempty"`
		Passphrase               *string                                                          `json:"passphrase,omitempty"`
		PrivateKeyData           *string                                                          `json:"privateKeyData,omitempty"`
		RootData                 *string                                                          `json:"rootData,omitempty"`
	}

	CertificateCreateCertificatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// CertificateCreateCertificatePost: Use this API command to create an installed certificate.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *CertificateCreateCertificatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateCreateCertificatePost201Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateCreateCertificatePost(ctx *UserContext, requestBody *CertificateCreateCertificatePostRequest) (*http.Response, *CertificateCreateCertificatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/certstore/certificate")
	request.body = requestBody
	request.authenticated = true
	out := &CertificateCreateCertificatePost201Response{}
	httpResponse, _, err := c.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// CertificateDeleteCertificateDelete: Use this API command to delete an installed certificate.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateDeleteCertificateDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "DELETE", "/v5_0/certstore/certificate/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return c.client.doRequest(request, 204, nil)
}

type (
	CertificateRetrieveCertificateGet200ResponseCertificasSigningRequest struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	CertificateRetrieveCertificateGet200ResponseIntermediateDataSlice []*string

	CertificateRetrieveCertificateGet200Response struct {
		CertificasSigningRequest *CertificateRetrieveCertificateGet200ResponseCertificasSigningRequest `json:"certificasSigningRequest,omitempty"`
		Data                     *string                                                               `json:"data,omitempty"`
		Description              *string                                                               `json:"description,omitempty"`
		ID                       *string                                                               `json:"id,omitempty"`
		Information              *string                                                               `json:"information,omitempty"`
		IntermediateData         CertificateRetrieveCertificateGet200ResponseIntermediateDataSlice     `json:"intermediateData,omitempty"`
		Name                     *string                                                               `json:"name,omitempty"`
		Passphrase               *string                                                               `json:"passphrase,omitempty"`
		PrivateKeyData           *string                                                               `json:"privateKeyData,omitempty"`
		RootData                 *string                                                               `json:"rootData,omitempty"`
	}
)

// CertificateRetrieveCertificateGet: Use this API command to retrieve an installed certificate.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveCertificateGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveCertificateGet(ctx *UserContext, id string) (*http.Response, *CertificateRetrieveCertificateGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/certificate/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &CertificateRetrieveCertificateGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificateRetrieveCsrListGet200ResponseListSlice []*CertificateRetrieveCsrListGet200ResponseList

	CertificateRetrieveCsrListGet200ResponseList struct {
		Description *string `json:"description,omitempty"`
		ID          *string `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
	}

	CertificateRetrieveCsrListGet200Response struct {
		FirstIndex *int                                              `json:"firstIndex,omitempty"`
		HasMore    *bool                                             `json:"hasMore,omitempty"`
		List       CertificateRetrieveCsrListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                              `json:"totalCount,omitempty"`
	}
)

// CertificateRetrieveCsrListGet: Use this API command to retrieve list of certificates signing request.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveCsrListGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveCsrListGet(ctx *UserContext, optionalParams map[string]string) (*http.Response, *CertificateRetrieveCsrListGet200Response, error) {
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
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/csr")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}
	out := &CertificateRetrieveCsrListGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificateCreateCsrPostRequest struct {
		City             *string `json:"city,omitempty"`
		CommonName       *string `json:"commonName,omitempty"`
		CountryCode      *string `json:"countryCode,omitempty"`
		Description      *string `json:"description,omitempty"`
		Email            *string `json:"email,omitempty"`
		Name             *string `json:"name,omitempty"`
		Organization     *string `json:"organization,omitempty"`
		OrganizationUnit *string `json:"organizationUnit,omitempty"`
		State            *string `json:"state,omitempty"`
	}

	CertificateCreateCsrPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// CertificateCreateCsrPost: Use this API command to create a certificates signing request.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *CertificateCreateCsrPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateCreateCsrPost201Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateCreateCsrPost(ctx *UserContext, requestBody *CertificateCreateCsrPostRequest) (*http.Response, *CertificateCreateCsrPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/certstore/csr")
	request.body = requestBody
	request.authenticated = true
	out := &CertificateCreateCsrPost201Response{}
	httpResponse, _, err := c.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// CertificateDeleteCsrDelete: Use this API command to delete a certificates signing request.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateDeleteCsrDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "DELETE", "/v5_0/certstore/csr/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return c.client.doRequest(request, 204, nil)
}

type (
	CertificateRetrieveCsrGet200Response struct {
		City             *string `json:"city,omitempty"`
		CommonName       *string `json:"commonName,omitempty"`
		CountryCode      *string `json:"countryCode,omitempty"`
		Description      *string `json:"description,omitempty"`
		Email            *string `json:"email,omitempty"`
		ID               *string `json:"id,omitempty"`
		Name             *string `json:"name,omitempty"`
		Organization     *string `json:"organization,omitempty"`
		OrganizationUnit *string `json:"organizationUnit,omitempty"`
		State            *string `json:"state,omitempty"`
	}
)

// CertificateRetrieveCsrGet: Use this API command to retrieve a certificates signing request.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveCsrGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveCsrGet(ctx *UserContext, id string) (*http.Response, *CertificateRetrieveCsrGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/csr/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &CertificateRetrieveCsrGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificateRetrieveCertificateSettingGet200ResponseServiceCertificatesSlice []*CertificateRetrieveCertificateSettingGet200ResponseServiceCertificates

	CertificateRetrieveCertificateSettingGet200ResponseServiceCertificatesCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	CertificateRetrieveCertificateSettingGet200ResponseServiceCertificates struct {
		Certificate *CertificateRetrieveCertificateSettingGet200ResponseServiceCertificatesCertificate `json:"certificate,omitempty"`
		Service     *string                                                                            `json:"service,omitempty"`
	}

	CertificateRetrieveCertificateSettingGet200Response struct {
		ServiceCertificates CertificateRetrieveCertificateSettingGet200ResponseServiceCertificatesSlice `json:"serviceCertificates,omitempty"`
	}
)

// CertificateRetrieveCertificateSettingGet: Use this API command to retrieve certificate setting.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveCertificateSettingGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveCertificateSettingGet(ctx *UserContext) (*http.Response, *CertificateRetrieveCertificateSettingGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/setting")
	request.authenticated = true
	out := &CertificateRetrieveCertificateSettingGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificateModifyServicecertificatesPatchRequestSlice []*CertificateModifyServicecertificatesPatchRequest

	CertificateModifyServicecertificatesPatchRequestCertificate struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	CertificateModifyServicecertificatesPatchRequest struct {
		Certificate *CertificateModifyServicecertificatesPatchRequestCertificate `json:"certificate,omitempty"`
		Service     *string                                                      `json:"service,omitempty"`
	}
)

// CertificateModifyServicecertificatesPatch: Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *CertificateModifyServicecertificatesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateModifyServicecertificatesPatch(ctx *UserContext, requestBody CertificateModifyServicecertificatesPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "PATCH", "/v5_0/certstore/setting/serviceCertificates")
	request.body = requestBody
	request.authenticated = true
	return c.client.doRequest(request, 204, nil)
}

type (
	CertificateDeleteBulkTrustedCaChainCertificatesDeleteRequestIDListSlice []*string

	CertificateDeleteBulkTrustedCaChainCertificatesDeleteRequest struct {
		IDList CertificateDeleteBulkTrustedCaChainCertificatesDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// CertificateDeleteBulkTrustedCaChainCertificatesDelete: Use this API command to delete bulk trusted CA chain certificates.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *CertificateDeleteBulkTrustedCaChainCertificatesDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateDeleteBulkTrustedCaChainCertificatesDelete(ctx *UserContext, requestBody *CertificateDeleteBulkTrustedCaChainCertificatesDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "DELETE", "/v5_0/certstore/trustedCAChainCert")
	request.body = requestBody
	request.authenticated = true
	return c.client.doRequest(request, 204, nil)
}

type (
	CertificateRetrieveTrustedCaChainCertificatesListGet200ResponseListSlice []*CertificateRetrieveTrustedCaChainCertificatesListGet200ResponseList

	CertificateRetrieveTrustedCaChainCertificatesListGet200ResponseListInterCertDataSlice []*string

	CertificateRetrieveTrustedCaChainCertificatesListGet200ResponseList struct {
		Description      *string                                                                               `json:"description,omitempty"`
		ID               *string                                                                               `json:"id,omitempty"`
		Information      *string                                                                               `json:"information,omitempty"`
		InterCertData    CertificateRetrieveTrustedCaChainCertificatesListGet200ResponseListInterCertDataSlice `json:"interCertData,omitempty"`
		ModifiedDateTime *string                                                                               `json:"modifiedDateTime,omitempty"`
		ModifierUsername *string                                                                               `json:"modifierUsername,omitempty"`
		Name             *string                                                                               `json:"name,omitempty"`
		RootCertData     *string                                                                               `json:"rootCertData,omitempty"`
	}

	CertificateRetrieveTrustedCaChainCertificatesListGet200Response struct {
		FirstIndex *int                                                                     `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                    `json:"hasMore,omitempty"`
		List       CertificateRetrieveTrustedCaChainCertificatesListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                     `json:"totalCount,omitempty"`
	}
)

// CertificateRetrieveTrustedCaChainCertificatesListGet: Use this API command to retrieve list of installed trusted CA chain certificates.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveTrustedCaChainCertificatesListGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveTrustedCaChainCertificatesListGet(ctx *UserContext, optionalParams map[string]string) (*http.Response, *CertificateRetrieveTrustedCaChainCertificatesListGet200Response, error) {
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
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/trustedCAChainCert")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
	}
	out := &CertificateRetrieveTrustedCaChainCertificatesListGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificateCreateTrustedCaChainCertificatesPostRequestInterCertDataSlice []*string

	CertificateCreateTrustedCaChainCertificatesPostRequest struct {
		Description   *string                                                                  `json:"description,omitempty"`
		InterCertData CertificateCreateTrustedCaChainCertificatesPostRequestInterCertDataSlice `json:"interCertData,omitempty"`
		Name          *string                                                                  `json:"name,omitempty"`
		RootCertData  *string                                                                  `json:"rootCertData,omitempty"`
	}

	CertificateCreateTrustedCaChainCertificatesPost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// CertificateCreateTrustedCaChainCertificatesPost: Use this API command to create trusted CA chain certificates.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *CertificateCreateTrustedCaChainCertificatesPostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateCreateTrustedCaChainCertificatesPost201Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateCreateTrustedCaChainCertificatesPost(ctx *UserContext, requestBody *CertificateCreateTrustedCaChainCertificatesPostRequest) (*http.Response, *CertificateCreateTrustedCaChainCertificatesPost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/certstore/trustedCAChainCert")
	request.body = requestBody
	request.authenticated = true
	out := &CertificateCreateTrustedCaChainCertificatesPost201Response{}
	httpResponse, _, err := c.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// CertificateDeleteTrustedCaChainCertificatesDelete: Use this API command to delete a trusted CA chain certificate.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateDeleteTrustedCaChainCertificatesDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "DELETE", "/v5_0/certstore/trustedCAChainCert/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return c.client.doRequest(request, 204, nil)
}

type (
	CertificateRetrieveTrustedCaChainCertificatesGet200ResponseInterCertDataSlice []*string

	CertificateRetrieveTrustedCaChainCertificatesGet200Response struct {
		Description   *string                                                                       `json:"description,omitempty"`
		ID            *string                                                                       `json:"id,omitempty"`
		InterCertData CertificateRetrieveTrustedCaChainCertificatesGet200ResponseInterCertDataSlice `json:"interCertData,omitempty"`
		Name          *string                                                                       `json:"name,omitempty"`
		RootCertData  *string                                                                       `json:"rootCertData,omitempty"`
	}
)

// CertificateRetrieveTrustedCaChainCertificatesGet: Use this API command to retrieve an installed trusted CA chain certificates.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *CertificateRetrieveTrustedCaChainCertificatesGet200Response
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificateRetrieveTrustedCaChainCertificatesGet(ctx *UserContext, id string) (*http.Response, *CertificateRetrieveTrustedCaChainCertificatesGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "GET", "/v5_0/certstore/trustedCAChainCert/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &CertificateRetrieveTrustedCaChainCertificatesGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	CertificatePatchTrustedCaChainCertificatesPatchRequest struct {
		Description  *string `json:"description,omitempty"`
		Information  *string `json:"information,omitempty"`
		Name         *string `json:"name,omitempty"`
		RootCertData *string `json:"rootCertData,omitempty"`
	}
)

// CertificatePatchTrustedCaChainCertificatesPatch: Use this API command to patch a trusted CA chain certificates.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *CertificatePatchTrustedCaChainCertificatesPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *CertStoreAPI) CertificatePatchTrustedCaChainCertificatesPatch(ctx *UserContext, id string, requestBody *CertificatePatchTrustedCaChainCertificatesPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := c.client.newRequest(ctx, "PATCH", "/v5_0/certstore/trustedCAChainCert/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return c.client.doRequest(request, 204, nil)
}
