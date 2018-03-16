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
// - ctx (context.Context): Context to use for this request
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
func (c *ClusterAPI) ClusterBackupAndRestoreRetrieveListGet(ctx context.Context, optionalParams map[string]string) (*http.Response, *ClusterBackupAndRestoreRetrieveListGet200Response, error) {
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
	request := NewRequest("GET", "/v5_0/cluster", true)
	request.SetQueryParameter("index", index)
	request.SetQueryParameter("listSize", listSize)
	request.SetQueryParameter("timezone", optionalParams["timezone"])
	out := &ClusterBackupAndRestoreRetrieveListGet200Response{}
	httpResponse, _, err := c.client.Ensure(ctx, request, 200, out)
	return httpResponse, out, err
}

// ClusterBackupAndRestoreClusterBackupPost: Backup cluster
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreClusterBackupPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/cluster/backup", true)
	return c.client.Ensure(ctx, request, 204, nil)
}

// ClusterBackupAndRestoreClusterRestorePost: Restore cluster backup by ID
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreClusterRestorePost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("POST", "/v5_0/cluster/restore/{id:.+}", true)
	return c.client.Ensure(ctx, request, 204, nil)
}

// ClusterBackupAndRestoreDeleteDelete: Delete cluster backup by ID
//
// Required Parameters:
// - ctx (context.Context): Context to use for this request
//
// Returns:
// - *http.Response: HTTP Response or nil on error
// - []byte: Any bytes to be found in response body
// - error: Error seen or nil if none
func (c *ClusterAPI) ClusterBackupAndRestoreDeleteDelete(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("user context cannot be nil")
	}
	request := NewRequest("DELETE", "/v5_0/cluster/{id:.+}", true)
	return c.client.Ensure(ctx, request, 204, nil)
}
