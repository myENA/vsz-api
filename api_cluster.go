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

type ClusterAPI struct {
	client *Client
}
type (
	ClusterBackupAndRestoreRetrieveListGet200ResponseListSlice []*ClusterBackupAndRestoreRetrieveListGet200ResponseList

	ClusterBackupAndRestoreRetrieveListGet200ResponseList struct {
		CreatedOn *string  `json:"createdOn,omitempty"`
		Filesize  *float64 `json:"filesize,omitempty"`
		ID        *string  `json:"id,omitempty"`
		Version   *string  `json:"version,omitempty"`
	}

	ClusterBackupAndRestoreRetrieveListGet200Response struct {
		FirstIndex *int                                                       `json:"firstIndex,omitempty"`
		HasMore    *bool                                                      `json:"hasMore,omitempty"`
		List       ClusterBackupAndRestoreRetrieveListGet200ResponseListSlice `json:"list,omitempty"`
		TotalCount *int                                                       `json:"totalCount,omitempty"`
	}
)

// ClusterBackupAndRestoreRetrieveListGet: Retrive cluster backup list
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Optional Parameter Map:
// - index (integer): The index of the first entry to be retrieved.
// - listSize (integer): The maximum number of entries to be retrieved.
// - timezone (string): Change display timezone of created date and time, like timezone=Asia/Taipei,
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - *ClusterBackupAndRestoreRetrieveListGet200Response
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreRetrieveListGet(ctx *UserContext, optionalParams map[string]string) (*http.Response, *ClusterBackupAndRestoreRetrieveListGet200Response, error) {
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
	request := c.client.newRequest(ctx, "GET", "/v5_0/cluster")
	request.authenticated = true
	request.queryParameters = map[string]string{
		"index":    index,
		"listSize": listSize,
		"timezone": optionalParams["timezone"],
	}
	out := &ClusterBackupAndRestoreRetrieveListGet200Response{}
	httpResponse, _, err := c.client.doRequest(request, 200, out)
	return httpResponse, out, err
}

// ClusterBackupAndRestoreClusterBackupPost: Backup cluster
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreClusterBackupPost(ctx *UserContext) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/cluster/backup")
	request.authenticated = true
	return c.client.doRequest(request, 204, nil)
}

// ClusterBackupAndRestoreClusterRestorePost: Restore cluster backup by ID
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreClusterRestorePost(ctx *UserContext) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "POST", "/v5_0/cluster/restore/{id:.+}")
	request.authenticated = true
	return c.client.doRequest(request, 204, nil)
}

// ClusterBackupAndRestoreDeleteDelete: Delete cluster backup by ID
//
// Required Parameters:
// - ctx (*UserContext): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreDeleteDelete(ctx *UserContext) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := c.client.newRequest(ctx, "DELETE", "/v5_0/cluster/{id:.+}")
	request.authenticated = true
	return c.client.doRequest(request, 204, nil)
}
