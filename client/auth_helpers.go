package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// LoginForm is sent to /core/auth/login.
type LoginForm struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Captcha    string `json:"captcha,omitempty"`
	CaptchaID  string `json:"captchaID,omitempty"`
	AuthMethod string `json:"authMethod,omitempty"` // "session" | "api" ...
	AuthSource string `json:"authSource,omitempty"` // "local" | "ldap"
}

// LoginResult is the standard login response (see frontend/src/api/interface/auth.ts).
type LoginResult struct {
	Name       string `json:"name"`
	Role       string `json:"role"`
	Token      string `json:"token"`
	MFAStatus  string `json:"mfaStatus"`
	MFASession string `json:"mfaSession"`
}

// Login performs username/password authentication against /core/auth/login.
//
// After a successful call the session cookies (and "pcsrftoken") are stored in the
// underlying cookie jar, ready for subsequent authenticated requests.
func (c *Client) Login(ctx context.Context, form LoginForm) (*LoginResult, error) {
	if form.AuthMethod == "" {
		form.AuthMethod = "session"
	}
	if form.AuthSource == "" {
		form.AuthSource = "local"
	}
	var out LoginResult
	if err := c.Post(ctx, "/core/auth/login", form, &out); err != nil {
		return nil, err
	}
	c.SyncCSRF()
	return &out, nil
}

// Logout posts to /core/auth/logout. The cookie jar is cleared so future calls fail auth.
func (c *Client) Logout(ctx context.Context) error {
	// 1Panel returns success with no body. We just POST and let SyncCSRF refresh.
	if err := c.Post(ctx, "/core/auth/logout", nil, nil); err != nil {
		// Even on error clear CSRF to be safe.
		c.csrfToken = ""
		return err
	}
	c.csrfToken = ""
	return nil
}

// SetCookieForURL manually inserts a cookie into the jar. Useful when the caller obtained
// the session through a non-standard flow (e.g. API key login) and wants to reuse the jar.
func (c *Client) SetCookieForURL(rawURL string, name, value string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	c.hc.Jar.SetCookies(u, []*http.Cookie{{Name: name, Value: value, Path: "/"}})
	return nil
}

// BuildLoginURL composes the public URL used by human browsers to reach the panel.
// Convenience for printing or sharing.
func (c *Client) BuildLoginURL() string {
	base := strings.TrimRight(c.Endpoint(), "/")
	if c.cfg.Entrance == "" {
		return base
	}
	return fmt.Sprintf("%s/%s", base, c.cfg.Entrance)
}
