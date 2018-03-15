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

type AVCAPI struct {
	client *Client
}
type (
	ApplicationVisibilityControlApplicationPolicyMultipleDeleteDeleteRequestIDListSlice []*string

	ApplicationVisibilityControlApplicationPolicyMultipleDeleteDeleteRequest struct {
		IDList ApplicationVisibilityControlApplicationPolicyMultipleDeleteDeleteRequestIDListSlice `json:"idList,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationPolicyMultipleDeleteDelete: Use this API command to delete a AVC Application Policy Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ApplicationVisibilityControlApplicationPolicyMultipleDeleteDeleteRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationPolicyMultipleDeleteDelete(ctx *UserContext, requestBody *ApplicationVisibilityControlApplicationPolicyMultipleDeleteDeleteRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "DELETE", "/v5_0/avc/applicationPolicy")
	request.body = requestBody
	request.authenticated = true
	return a.client.doRequest(request, 200, nil)
}

type (
	ApplicationVisibilityControlApplicationPolicyCreatePostRequestApplicationRulesSlice []*ApplicationVisibilityControlApplicationPolicyCreatePostRequestApplicationRules

	ApplicationVisibilityControlApplicationPolicyCreatePostRequestApplicationRules struct {
		AppID              *string  `json:"appId,omitempty"`
		AppName            *string  `json:"appName,omitempty"`
		ApplicationType    *string  `json:"applicationType,omitempty"`
		CatID              *string  `json:"catId,omitempty"`
		CatName            *string  `json:"catName,omitempty"`
		ClassificationType *string  `json:"classificationType,omitempty"`
		Downlink           *float64 `json:"downlink,omitempty"`
		MarkingPriority    *string  `json:"markingPriority,omitempty"`
		MarkingType        *string  `json:"markingType,omitempty"`
		Priority           *float64 `json:"priority,omitempty"`
		RuleType           *string  `json:"ruleType,omitempty"`
		Uplink             *float64 `json:"uplink,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyCreatePostRequest struct {
		ApplicationRules ApplicationVisibilityControlApplicationPolicyCreatePostRequestApplicationRulesSlice `json:"applicationRules,omitempty"`
		Description      *string                                                                             `json:"description,omitempty"`
		DomainID         *string                                                                             `json:"domainId,omitempty"`
		Name             *string                                                                             `json:"name,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationPolicyCreatePost: Use this API command to create a new AVC Application Policy profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ApplicationVisibilityControlApplicationPolicyCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationPolicyCreatePost201Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationPolicyCreatePost(ctx *UserContext, requestBody *ApplicationVisibilityControlApplicationPolicyCreatePostRequest) (*http.Response, *ApplicationVisibilityControlApplicationPolicyCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "POST", "/v5_0/avc/applicationPolicy")
	request.body = requestBody
	request.authenticated = true
	out := &ApplicationVisibilityControlApplicationPolicyCreatePost201Response{}
	httpResponse, _, err := a.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// ApplicationVisibilityControlApplicationPolicyDeleteDelete: Use this API command to delete a AVC Application Policy Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationPolicyDeleteDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "DELETE", "/v5_0/avc/applicationPolicy/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return a.client.doRequest(request, 204, nil)
}

type (
	ApplicationVisibilityControlApplicationPolicyRetrieveGet200ResponseApplicationRulesSlice []*ApplicationVisibilityControlApplicationPolicyRetrieveGet200ResponseApplicationRules

	ApplicationVisibilityControlApplicationPolicyRetrieveGet200ResponseApplicationRules struct {
		AppID              *string  `json:"appId,omitempty"`
		AppName            *string  `json:"appName,omitempty"`
		ApplicationType    *string  `json:"applicationType,omitempty"`
		CatID              *string  `json:"catId,omitempty"`
		CatName            *string  `json:"catName,omitempty"`
		ClassificationType *string  `json:"classificationType,omitempty"`
		Downlink           *float64 `json:"downlink,omitempty"`
		MarkingPriority    *string  `json:"markingPriority,omitempty"`
		MarkingType        *string  `json:"markingType,omitempty"`
		Priority           *float64 `json:"priority,omitempty"`
		RuleType           *string  `json:"ruleType,omitempty"`
		Uplink             *float64 `json:"uplink,omitempty"`
	}

	ApplicationVisibilityControlApplicationPolicyRetrieveGet200Response struct {
		ApplicationRules ApplicationVisibilityControlApplicationPolicyRetrieveGet200ResponseApplicationRulesSlice `json:"applicationRules,omitempty"`
		CreateDateTime   *int                                                                                     `json:"createDateTime,omitempty"`
		CreatorID        *string                                                                                  `json:"creatorId,omitempty"`
		CreatorUsername  *string                                                                                  `json:"creatorUsername,omitempty"`
		Description      *string                                                                                  `json:"description,omitempty"`
		DomainID         *string                                                                                  `json:"domainId,omitempty"`
		ID               *string                                                                                  `json:"id,omitempty"`
		ModifiedDateTime *int                                                                                     `json:"modifiedDateTime,omitempty"`
		ModifierID       *string                                                                                  `json:"modifierId,omitempty"`
		ModifierUsername *string                                                                                  `json:"modifierUsername,omitempty"`
		Name             *string                                                                                  `json:"name,omitempty"`
		TenantID         *string                                                                                  `json:"tenantId,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationPolicyRetrieveGet: Use this API command to retrieve a AVC Application Policy profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationPolicyRetrieveGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationPolicyRetrieveGet(ctx *UserContext, id string) (*http.Response, *ApplicationVisibilityControlApplicationPolicyRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/applicationPolicy/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &ApplicationVisibilityControlApplicationPolicyRetrieveGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlApplicationPolicyModifyBasicPatchRequest struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationPolicyModifyBasicPatch: Use this API command to modify the basic information on AVC Application Policy profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *ApplicationVisibilityControlApplicationPolicyModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationPolicyModifyBasicPatch(ctx *UserContext, id string, requestBody *ApplicationVisibilityControlApplicationPolicyModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "PATCH", "/v5_0/avc/applicationPolicy/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return a.client.doRequest(request, 204, nil)
}

type (
	ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatchRequestSlice []*ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatchRequest

	ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatchRequest struct {
		AppID              *string  `json:"appId,omitempty"`
		AppName            *string  `json:"appName,omitempty"`
		ApplicationType    *string  `json:"applicationType,omitempty"`
		CatID              *string  `json:"catId,omitempty"`
		CatName            *string  `json:"catName,omitempty"`
		ClassificationType *string  `json:"classificationType,omitempty"`
		Downlink           *float64 `json:"downlink,omitempty"`
		MarkingPriority    *string  `json:"markingPriority,omitempty"`
		MarkingType        *string  `json:"markingType,omitempty"`
		Priority           *float64 `json:"priority,omitempty"`
		RuleType           *string  `json:"ruleType,omitempty"`
		Uplink             *float64 `json:"uplink,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatch: Use this API command to modify the applicationRules information on AVC Application Policy profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatchRequestSlice
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatch(ctx *UserContext, id string, requestBody ApplicationVisibilityControlApplicationPolicyModifyApplicationrulesPatchRequestSlice) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "PATCH", "/v5_0/avc/applicationPolicy/{id}/applicationRules")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return a.client.doRequest(request, 204, nil)
}

type (
	ApplicationVisibilityControlCurrentSignaturePackageInfoGet200Response struct {
		FileName *string  `json:"fileName,omitempty"`
		ID       *string  `json:"id,omitempty"`
		Size     *float64 `json:"size,omitempty"`
		Version  *string  `json:"version,omitempty"`
	}
)

// ApplicationVisibilityControlCurrentSignaturePackageInfoGet: Get current Signature Package info
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlCurrentSignaturePackageInfoGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlCurrentSignaturePackageInfoGet(ctx *UserContext) (*http.Response, *ApplicationVisibilityControlCurrentSignaturePackageInfoGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/signaturePackage")
	request.authenticated = true
	out := &ApplicationVisibilityControlCurrentSignaturePackageInfoGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlApplicationsGet200ResponseListSlice []*ApplicationVisibilityControlApplicationsGet200ResponseList

	ApplicationVisibilityControlApplicationsGet200ResponseList struct {
		AppID *string `json:"appId,omitempty"`
		CatID *string `json:"catId,omitempty"`
		Name  *string `json:"name,omitempty"`
	}

	ApplicationVisibilityControlApplicationsGet200Response struct {
		FirstIndex *int                                                            `json:"firstIndex,omitempty"`
		HasMore    *bool                                                           `json:"hasMore,omitempty"`
		List       ApplicationVisibilityControlApplicationsGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                            `json:"totalCount,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationsGet: Get Application list from current Signature Package
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationsGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationsGet(ctx *UserContext) (*http.Response, *ApplicationVisibilityControlApplicationsGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/signaturePackage/applications")
	request.authenticated = true
	out := &ApplicationVisibilityControlApplicationsGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlApplicationGet200Response struct {
		AppID *string `json:"appId,omitempty"`
		CatID *string `json:"catId,omitempty"`
		Name  *string `json:"name,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationGet: Get Application info (catId, appId and name) by application name
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - applicationName (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationGet(ctx *UserContext, applicationName string) (*http.Response, *ApplicationVisibilityControlApplicationGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(applicationName)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"applicationName\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/signaturePackage/application/{applicationName}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"applicationName": applicationName,
	}
	out := &ApplicationVisibilityControlApplicationGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlApplicationCategoriesGet200ResponseListSlice []*ApplicationVisibilityControlApplicationCategoriesGet200ResponseList

	ApplicationVisibilityControlApplicationCategoriesGet200ResponseList struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	ApplicationVisibilityControlApplicationCategoriesGet200Response struct {
		FirstIndex *int                                                                     `json:"firstIndex,omitempty"`
		HasMore    *bool                                                                    `json:"hasMore,omitempty"`
		List       ApplicationVisibilityControlApplicationCategoriesGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                                     `json:"totalCount,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationCategoriesGet: Get Application Category list from current Signature Package
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationCategoriesGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationCategoriesGet(ctx *UserContext) (*http.Response, *ApplicationVisibilityControlApplicationCategoriesGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/signaturePackage/categories")
	request.authenticated = true
	out := &ApplicationVisibilityControlApplicationCategoriesGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlApplicationCategoryGet200Response struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}
)

// ApplicationVisibilityControlApplicationCategoryGet: Get Application Category info (catId and name) by category name
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - categoryName (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlApplicationCategoryGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlApplicationCategoryGet(ctx *UserContext, categoryName string) (*http.Response, *ApplicationVisibilityControlApplicationCategoryGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(categoryName)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"categoryName\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/signaturePackage/category/{categoryName}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"categoryName": categoryName,
	}
	out := &ApplicationVisibilityControlApplicationCategoryGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlUploadFilePost200Response struct {
		FileName *string  `json:"fileName,omitempty"`
		ID       *string  `json:"id,omitempty"`
		Size     *float64 `json:"size,omitempty"`
		Version  *string  `json:"version,omitempty"`
	}
)

// ApplicationVisibilityControlUploadFilePost: Update AVC Signature Package by upload file
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlUploadFilePost200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlUploadFilePost(ctx *UserContext) (*http.Response, *ApplicationVisibilityControlUploadFilePost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "POST", "/v5_0/avc/signaturePackage/upload")
	request.authenticated = true
	out := &ApplicationVisibilityControlUploadFilePost200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlUserDefinedMultipleDeleteDelete200ResponseIDListSlice []*string

	ApplicationVisibilityControlUserDefinedMultipleDeleteDelete200Response struct {
		IDList ApplicationVisibilityControlUserDefinedMultipleDeleteDelete200ResponseIDListSlice `json:"idList,omitempty"`
	}
)

// ApplicationVisibilityControlUserDefinedMultipleDeleteDelete: Use this API command to delete a AVC User Defined Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlUserDefinedMultipleDeleteDelete200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlUserDefinedMultipleDeleteDelete(ctx *UserContext) (*http.Response, *ApplicationVisibilityControlUserDefinedMultipleDeleteDelete200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "DELETE", "/v5_0/avc/userDefined")
	request.authenticated = true
	out := &ApplicationVisibilityControlUserDefinedMultipleDeleteDelete200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlUserDefinedCreatePostRequest struct {
		DestIP   *string  `json:"destIp,omitempty"`
		DestPort *float64 `json:"destPort,omitempty"`
		DomainID *string  `json:"domainId,omitempty"`
		Name     *string  `json:"name,omitempty"`
		Netmask  *string  `json:"netmask,omitempty"`
		Protocol *string  `json:"protocol,omitempty"`
		Type     *string  `json:"type,omitempty"`
	}

	ApplicationVisibilityControlUserDefinedCreatePost201Response struct {
		ID *string `json:"id,omitempty"`
	}
)

// ApplicationVisibilityControlUserDefinedCreatePost: Use this API command to create a new AVC User Defined profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - requestBody: *ApplicationVisibilityControlUserDefinedCreatePostRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlUserDefinedCreatePost201Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlUserDefinedCreatePost(ctx *UserContext, requestBody *ApplicationVisibilityControlUserDefinedCreatePostRequest) (*http.Response, *ApplicationVisibilityControlUserDefinedCreatePost201Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := a.client.newRequest(ctx, "POST", "/v5_0/avc/userDefined")
	request.body = requestBody
	request.authenticated = true
	out := &ApplicationVisibilityControlUserDefinedCreatePost201Response{}
	httpResponse, _, err := a.client.doRequest(request, 201, out)
	return httpResponse, out, err
}

// ApplicationVisibilityControlUserDefinedDeleteDelete: Use this API command to delete a AVC User Defined Profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlUserDefinedDeleteDelete(ctx *UserContext, id string) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "DELETE", "/v5_0/avc/userDefined/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return a.client.doRequest(request, 204, nil)
}

type (
	ApplicationVisibilityControlUserDefinedRetrieveGet200Response struct {
		AppID            *float64 `json:"appId,omitempty"`
		CreateDateTime   *int     `json:"createDateTime,omitempty"`
		CreatorID        *string  `json:"creatorId,omitempty"`
		CreatorUsername  *string  `json:"creatorUsername,omitempty"`
		DestIP           *string  `json:"destIp,omitempty"`
		DestPort         *float64 `json:"destPort,omitempty"`
		DomainID         *string  `json:"domainId,omitempty"`
		ID               *string  `json:"id,omitempty"`
		ModifiedDateTime *int     `json:"modifiedDateTime,omitempty"`
		ModifierID       *string  `json:"modifierId,omitempty"`
		ModifierUsername *string  `json:"modifierUsername,omitempty"`
		Name             *string  `json:"name,omitempty"`
		Netmask          *string  `json:"netmask,omitempty"`
		Protocol         *string  `json:"protocol,omitempty"`
		TenantID         *string  `json:"tenantId,omitempty"`
		Type             *string  `json:"type,omitempty"`
	}
)

// ApplicationVisibilityControlUserDefinedRetrieveGet: Use this API command to retrieve a AVC User Defined profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ApplicationVisibilityControlUserDefinedRetrieveGet200Response
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlUserDefinedRetrieveGet(ctx *UserContext, id string) (*http.Response, *ApplicationVisibilityControlUserDefinedRetrieveGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "GET", "/v5_0/avc/userDefined/{id}")
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	out := &ApplicationVisibilityControlUserDefinedRetrieveGet200Response{}
	httpResponse, _, err := a.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

type (
	ApplicationVisibilityControlUserDefinedModifyBasicPatchRequest struct {
		DestIP   *string  `json:"destIp,omitempty"`
		DestPort *float64 `json:"destPort,omitempty"`
		Name     *string  `json:"name,omitempty"`
		Netmask  *string  `json:"netmask,omitempty"`
		Protocol *string  `json:"protocol,omitempty"`
		Type     *string  `json:"type,omitempty"`
	}
)

// ApplicationVisibilityControlUserDefinedModifyBasicPatch: Use this API command to modify the basic information on AVC User Defined profile.
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
// - id (string)
// - requestBody: *ApplicationVisibilityControlUserDefinedModifyBasicPatchRequest
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (a *AVCAPI) ApplicationVisibilityControlUserDefinedModifyBasicPatch(ctx *UserContext, id string, requestBody *ApplicationVisibilityControlUserDefinedModifyBasicPatchRequest) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	var err error
	err = validators.StrNotEmpty(id)
	if nil != err {
		return nil, nil, fmt.Errorf("parameter \"id\" failed validation check: %s", err)
	}
	request := a.client.newRequest(ctx, "PATCH", "/v5_0/avc/userDefined/{id}")
	request.body = requestBody
	request.authenticated = true
	request.pathParameters = map[string]string{
		"id": id,
	}
	return a.client.doRequest(request, 204, nil)
}
