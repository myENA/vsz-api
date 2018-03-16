package api

import (
	"fmt"
	"net/http"
	"strings"
)

func tryExtractSessionCookie(request *request, resp *http.Response) *http.Cookie {
	if debug {
		log.Printf("[request-%d] Attempting to locate \"%s\" cookie in: %+v", request.id, SessionCookieName, resp.Cookies())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == SessionCookieName {
			return cookie
		}
	}
	return nil
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
