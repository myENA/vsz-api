package api

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func TryExtractSessionCookie(resp *http.Response) *http.Cookie {
	if debug {
		log.Printf("Attempting to locate \"%s\" cookie in: %+v", SessionCookieName, resp.Cookies())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == SessionCookieName {
			return cookie
		}
	}
	return nil
}

func compileRequestURLString(scheme, address, pathPrefix, uri string, pathParams, queryParams map[string]string) string {
	url := fmt.Sprintf("%s://%s", scheme, path.Join(address, pathPrefix, uri))

	if nil == pathParams && nil == queryParams {
		return url
	}

	if nil != pathParams && 0 < len(pathParams) {
		for name, value := range pathParams {
			url = strings.Replace(url, fmt.Sprintf("{%s}", name), value, 1)
		}
	}

	if nil != queryParams && 0 < len(queryParams) {
		url = fmt.Sprintf("%s%s", url, buildQueryParamString(queryParams))
	}

	if debug {
		log.Printf("Build request URL: %s", url)
	}

	return url
}

func buildQueryParamString(queryParams map[string]string) string {
	// TODO: Make more efficient?
	paramParts := make([]string, 0)

	for name, value := range queryParams {
		value = strings.TrimSpace(value)
		if "" != value {
			paramParts = append(paramParts, fmt.Sprintf("%s=%s", name, value))
		}
	}

	if 0 < len(paramParts) {
		return fmt.Sprintf("?%s", strings.Join(paramParts, "&"))
	} else {
		return ""
	}
}
