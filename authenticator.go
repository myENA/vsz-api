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
	// Decorate must do one of two things:
	//
	// If the internal state of this authenticator is such that it currently has whatever is needed to modify a given
	// request with the appropriate authentication cookie / token / header / etc., then it must do so, returning the current
	// CAS and a nil error
	//
	// If the internal state is such that decoration CANNOT happen, it must return the current CAS and an error, describing
	// the reason it cannot decorate the request.  If the error is not nil, Refresh will be called with the CAS value
	// returned by this method.
	//
	// In all cases, the current CAS must be returned.
	Decorate(*http.Request) (AuthCAS, error)

	// Refresh must doe one of two things:
	//
	// If the provided CAS value is current, it must assume that its current state is no longer valid and try to do what
	// i sneeded to get back to a state that Decorate is able to do what it needs to do.
	//
	// If the provided CAS value is NOT equal to the current state, it must assume that a refresh attempt has already
	// been made in another process, and just return the current CAS value with no error.
	//
	// The client will only attempt a refresh once per execution, so if Decorate STILL returns an error at this point,
	// the client will return the Decorate error as the error for the entire request, as it was unable to perform its
	// task.
	Refresh(*Client, AuthCAS) (AuthCAS, error)
}

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
	if client == nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), errors.New("client cannot be nil")
	}
	pa.cookieMu.Lock()
	ccas := atomic.LoadUint64(&pa.cas)
	if ccas != uint64(cas) {
		// if the passed in CAS value does not match the currently stored one, assume somebody else has modified the
		// cookie and just return the current cas
		pa.cookieMu.Unlock()
		return AuthCAS(ccas), nil
	}
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
