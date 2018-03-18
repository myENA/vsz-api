package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func tryExtractSessionCookie(request *Request, resp *http.Response) *http.Cookie {
	if debug {
		log.Printf("[request-%d] Attempting to locate \"%s\" cookie in: %+v", request.id, SessionCookieToken, resp.Cookies())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == SessionCookieToken {
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

func fetchContextTTL(ctx context.Context) time.Duration {
	if ctx.Err() == nil {
		if ctxTime, ok := ctx.Deadline(); ok {
			return ctxTime.Sub(time.Now())
		}
	}
	return 0
}
