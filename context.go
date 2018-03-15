package api

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type (
	CookieCAS uint64

	UserContext struct {
		mu sync.RWMutex

		ctx    context.Context
		cancel context.CancelFunc

		username string
		password string

		requestTTL time.Duration

		cookie    *http.Cookie
		cookieCAS uint64
		cookieMu  sync.RWMutex
	}
)

func NewUserContext(parentCTX context.Context, username, password string, requestTTL time.Duration) *UserContext {
	c := &UserContext{
		username: username,
		password: password,
	}

	if parentCTX == nil {
		parentCTX = context.Background()
	}

	c.ctx, c.cancel = context.WithCancel(parentCTX)

	if requestTTL == 0 {
		c.requestTTL = DefaultRequestTimeout
	} else {
		c.requestTTL = requestTTL
	}

	return c
}

func (c *UserContext) context() context.Context {
	c.mu.Lock()
	if c.ctx == nil {
		c.ctx, c.cancel = context.WithCancel(context.Background())
	}
	ctx := c.ctx
	c.mu.Unlock()
	return ctx
}

func (c *UserContext) Deadline() (time.Time, bool) {
	return c.context().Deadline()
}

func (c *UserContext) Done() <-chan struct{} {
	return c.context().Done()
}

func (c *UserContext) Err() error {
	return c.context().Err()
}

func (c *UserContext) Value(key interface{}) interface{} {
	return c.context().Value(key)
}

// Username must return the username for this user context
func (c *UserContext) Username() string {
	c.mu.RLock()
	username := c.username
	c.mu.RUnlock()
	return username
}

// Password must return the password for this user context
func (c *UserContext) Password() string {
	c.mu.RLock()
	password := c.password
	c.mu.RUnlock()
	return password
}

// RequestTTL returns the max duration a request is allowed to live for this user context
func (c *UserContext) RequestTTL() time.Duration {
	c.mu.RLock()
	ttl := c.requestTTL
	c.mu.RUnlock()
	return ttl
}

// RequestContext returns a new timerCtx and CancelFunc using the UserContext as its parent and a deadline set to now + requestTTL
func (c *UserContext) RequestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(c, c.RequestTTL())
}

// Cookie will return the session cookie to use with this user context, if one has been defined
func (c *UserContext) Cookie() (*http.Cookie, CookieCAS) {
	c.cookieMu.RLock()
	cookie := c.cookie
	cas := atomic.LoadUint64(&c.cookieCAS)
	c.cookieMu.RUnlock()
	return cookie, CookieCAS(cas)
}

// RefreshCookie will attempt to refresh the session cookie for this UserContext.
func (c *UserContext) RefreshCookie(client *Client, cas CookieCAS) (CookieCAS, error) {
	if client == nil {
		return CookieCAS(atomic.LoadUint64(&c.cookieCAS)), errors.New("client cannot be nil")
	}
	c.cookieMu.Lock()
	ccas := atomic.LoadUint64(&c.cookieCAS)
	// if the passed in CAS value does not match the currently stored one, assume somebody else has modified the cookie
	// and just return the current cas
	if ccas != uint64(cas) {
		c.cookieMu.Unlock()
		return CookieCAS(ccas), nil
	}
	username := c.Username()
	password := c.Password()
	// TODO: not sure I like this.
	request := client.newRequest(c, "POST", "/v5_0/session")
	request.body = &LoginSessionLogonPostRequest{
		Username: &username,
		Password: &password,
	}
	resp, _, err := client.doRequest(request, 200, nil)
	if err != nil {
		c.cookie = nil
		c.cookieCAS = atomic.AddUint64(&c.cookieCAS, 1)
		c.cookieMu.Unlock()
		return CookieCAS(atomic.LoadUint64(&c.cookieCAS)), err
	}
	cookie := TryExtractSessionCookie(resp)
	if cookie == nil {
		c.cookie = nil
		c.cookieCAS = atomic.AddUint64(&c.cookieCAS, 1)
		c.cookieMu.Unlock()
		return CookieCAS(atomic.LoadUint64(&c.cookieCAS)), errors.New("unable to locate cookie in response")
	}
	c.cookie = cookie
	c.cookieCAS = atomic.AddUint64(&c.cookieCAS, 1)
	c.cookieMu.Unlock()
	return CookieCAS(atomic.LoadUint64(&c.cookieCAS)), nil
}
