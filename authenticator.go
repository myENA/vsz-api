package api

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	SessionCookieToken = "JSESSIONID"
)

type (
	AuthCAS uint64

	Authenticator interface {
		// Decorate must do one of two things:
		//
		// If the internal state of this authenticator is such that it currently has whatever is needed to modify a
		// given request with the appropriate authentication cookie / token / header / etc., then it must do so,
		// returning the current CAS and a nil error
		//
		// If the internal state is such that decoration CANNOT happen, it must return the current CAS and an error,
		// describing the reason it cannot decorate the request.  If the error is not nil, Refresh will be called with
		// the CAS value returned by this method.
		//
		// In all cases, the current CAS must be returned.
		Decorate(*http.Request) (AuthCAS, error)

		// Refresh must do one of two things:
		//
		// If the provided CAS value is current, it must assume that its current state is no longer valid and try to do
		// what is needed to get back to a state that Decorate is able to do what it needs to do.
		//
		// If the provided CAS value is NOT equal to the current state, it must assume that a refresh attempt has
		// already been made in another process, and just return the current CAS value with no error.
		//
		// The client will only attempt a maximum of 2 times per execution:
		//
		// 1. If the FIRST Decorate call fails
		// 2. If initial Decorate did not fail but VSZ returned a 401, causing an Invalidate -> Refresh -> Decorate loop
		// that will execute exactly 1 times.
		Refresh(*Client, AuthCAS) (AuthCAS, error)

		// Invalidate will only be called if a 401 is seen after a refresh has been attempted, and should indicate to
		// the implementor that whatever decoration the authenticator is current using is no longer considered valid by
		// the VSZ being queried
		Invalidate(AuthCAS) (AuthCAS, error)
	}
)

// PasswordAuthenticator is a simple example implementation of an Authenticator that will decorate a given request
// with a session id bearing cookie if one exists, and attempt to create one if it doesn't.
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
	pa := &PasswordAuthenticator{
		username:   username,
		password:   password,
		cookieTTL:  cookieTTL,
		requestTTL: requestTTL,
	}
	return pa
}

func (pa *PasswordAuthenticator) Username() string {
	return pa.username
}

func (pa *PasswordAuthenticator) Password() string {
	return pa.password
}

func (pa *PasswordAuthenticator) Decorate(httpRequest *http.Request) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Decorate called for request \"%s\"", pa.username, httpRequest.URL)
	}
	if httpRequest == nil {
		// TODO: yell a bit more if the request is nil?
		return AuthCAS(atomic.LoadUint64(&pa.cas)), errors.New("httpRequest cannot be nil")
	}
	pa.cookieMu.RLock()
	cas := atomic.LoadUint64(&pa.cas)
	cookie := pa.cookie
	// TODO improve efficiency of this?
	if cookie != nil && !pa.refreshed.IsZero() && pa.refreshed.Add(pa.cookieTTL).After(time.Now()) {
		httpRequest.AddCookie(cookie)
		pa.cookieMu.RUnlock()
		return AuthCAS(cas), nil
	}
	pa.cookieMu.RUnlock()
	return AuthCAS(cas), errors.New("cookie requires refresh")
}

func (pa *PasswordAuthenticator) Refresh(client *Client, cas AuthCAS) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Refresh called", pa.username)
	}
	if client == nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), errors.New("client cannot be nil")
	}
	pa.cookieMu.Lock()
	ccas := atomic.LoadUint64(&pa.cas)
	// if the passed cas value is greater than the internal CAS, assume weirdness and return current CAS and an error
	if ccas < uint64(cas) {
		pa.cookieMu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	// if the passed in CAS value is less than the currently stored one, assume another routine called either Refresh
	// or Invalidate and just return current cas
	if ccas > uint64(cas) {
		pa.cookieMu.Unlock()
		return AuthCAS(ccas), nil
	}

	// if cas matches internal...

	username := pa.username
	password := pa.password
	ctx, cancel := context.WithTimeout(context.Background(), pa.requestTTL)
	defer cancel()
	request := NewRequest("POST", "/v5_0/session", false)
	err := request.SetBodyModel(&LoginSessionLogonPostRequest{
		Username: &username,
		Password: &password,
	})
	if err != nil {
		pa.cookie = nil
		ncas := atomic.AddUint64(&pa.cas, 1)
		pa.cookieMu.Unlock()
		return AuthCAS(ncas), err
	}
	resp, _, err := client.Ensure(ctx, request, 200, nil)
	if err != nil {
		pa.cookie = nil
		ncas := atomic.AddUint64(&pa.cas, 1)
		pa.cookieMu.Unlock()
		return AuthCAS(ncas), err
	}
	cookie := tryExtractSessionCookie(request, resp)
	if cookie == nil {
		pa.cookie = nil
		ncas := atomic.AddUint64(&pa.cas, 1)
		pa.cookieMu.Unlock()
		return AuthCAS(ncas), errors.New("unable to locate cookie in response")
	}
	pa.cookie = cookie
	pa.refreshed = time.Now()
	ncas := atomic.AddUint64(&pa.cas, 1)
	pa.cookieMu.Unlock()
	return AuthCAS(ncas), nil
}

func (pa *PasswordAuthenticator) Invalidate(cas AuthCAS) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Invalidate called", pa.username)
	}
	pa.cookieMu.Lock()
	ccas := atomic.LoadUint64(&pa.cas)
	// if current cas is less than provided, assume insanity.
	if ccas < uint64(cas) {
		pa.cookieMu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value greater than possible")
	}
	// if current cas is greater than provided, assume Refresh or Invalidate has already been called.
	if ccas > uint64(cas) {
		pa.cookieMu.Unlock()
		return AuthCAS(ccas), nil
	}
	ncas := atomic.AddUint64(&pa.cas, 1)
	pa.cookie = nil
	pa.refreshed = time.Now()
	pa.cookieMu.Unlock()
	return AuthCAS(ncas), nil
}
