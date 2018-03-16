package api

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type AuthCAS uint64

type Authenticator interface {
	// Decorate should do whatever is needed to decorate the request with whatever authentication token is desired
	// if decoration fails, Refresh will be called, being given the CAS returned by Decorate.
	Decorate(*http.Request) (AuthCAS, error)
	Refresh(*Client, AuthCAS) (AuthCAS, error)
}

type PasswordAuthenticator struct {
	username   string
	password   string
	requestTTL time.Duration

	cas       uint64
	refreshed time.Time
	cookieTTL time.Duration
	cookie    *http.Cookie
	cookieMu  sync.RWMutex
}

func NewPasswordAuthenticator(username, password string, cookieTTL time.Duration, requestTTL time.Duration) *PasswordAuthenticator {
	csa := &PasswordAuthenticator{
		username:   username,
		password:   password,
		cookieTTL:  cookieTTL,
		requestTTL: requestTTL,
	}
	return csa
}

func (csa *PasswordAuthenticator) Username() string {
	return csa.username
}

func (csa *PasswordAuthenticator) Password() string {
	return csa.password
}

func (csa *PasswordAuthenticator) Decorate(httpRequest *http.Request) (AuthCAS, error) {
	if httpRequest == nil {
		return 0, errors.New("httpRequest cannot be nil")
	}
	csa.cookieMu.RLock()
	cas := atomic.LoadUint64(&csa.cas)
	cookie := csa.cookie
	// TODO improve efficiency of this?
	if cookie != nil && !csa.refreshed.IsZero() && csa.refreshed.Add(csa.cookieTTL).After(time.Now()) {
		httpRequest.AddCookie(cookie)
		csa.cookieMu.RUnlock()
		return AuthCAS(cas), nil
	}
	csa.cookieMu.RUnlock()
	return AuthCAS(cas), errors.New("cookie requires refresh")
}

func (csa *PasswordAuthenticator) Refresh(client *Client, cas AuthCAS) (AuthCAS, error) {
	if client == nil {
		return AuthCAS(atomic.LoadUint64(&csa.cas)), errors.New("client cannot be nil")
	}
	csa.cookieMu.Lock()
	ccas := atomic.LoadUint64(&csa.cas)
	if ccas != uint64(cas) {
		// if the passed in CAS value does not match the currently stored one, assume somebody else has modified the
		// cookie and just return the current cas
		csa.cookieMu.Unlock()
		return AuthCAS(ccas), nil
	}
	username := csa.username
	password := csa.password
	ctx, cancel := context.WithTimeout(context.Background(), csa.requestTTL)
	defer cancel()
	request := NewRequest("POST", "/v5_0/session", false)
	err := request.SetBodyModel(&LoginSessionLogonPostRequest{
		Username: &username,
		Password: &password,
	})
	if err != nil {
		csa.cookie = nil
		ncas := atomic.AddUint64(&csa.cas, 1)
		csa.cookieMu.Unlock()
		return AuthCAS(ncas), err
	}
	resp, _, err := client.Ensure(ctx, request, 200, nil)
	if err != nil {
		csa.cookie = nil
		ncas := atomic.AddUint64(&csa.cas, 1)
		csa.cookieMu.Unlock()
		return AuthCAS(ncas), err
	}
	cookie := tryExtractSessionCookie(request, resp)
	if cookie == nil {
		csa.cookie = nil
		ncas := atomic.AddUint64(&csa.cas, 1)
		csa.cookieMu.Unlock()
		return AuthCAS(ncas), errors.New("unable to locate cookie in response")
	}
	csa.cookie = cookie
	csa.refreshed = time.Now()
	ncas := atomic.AddUint64(&csa.cas, 1)
	csa.cookieMu.Unlock()
	return AuthCAS(ncas), nil
}
