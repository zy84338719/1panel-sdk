package onepanel

import (
	"context"
	"net/http"

	"github.com/zy84338719/1panel-sdk/client"
)

// AuthService covers /core/auth/* and the master-panel authentication flow.
type AuthService struct {
	ServiceBase
}

// LoginForm is the payload for /core/auth/login. Most fields are optional.
// When MFA is required the panel returns 401 with mfaSession; the caller should
// then call LoginByMFA with the received code.
type LoginForm struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Captcha    string `json:"captcha,omitempty"`
	CaptchaID  string `json:"captchaID,omitempty"`
	AuthMethod string `json:"authMethod,omitempty"`
	AuthSource string `json:"authSource,omitempty"`
}

// Login performs username/password authentication.
func (s *AuthService) Login(form LoginForm) (*client.LoginResult, error) {
	if form.AuthMethod == "" {
		form.AuthMethod = "session"
	}
	if form.AuthSource == "" {
		form.AuthSource = "local"
	}
	var out client.LoginResult
	if err := s.Post(context.Background(), "/core/auth/login", form, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// LoginByMFA finishes an MFA challenge given the session id returned by the first Login attempt.
func (s *AuthService) LoginByMFA(sessionID, code string) (*client.LoginResult, error) {
	var out client.LoginResult
	if err := s.Post(context.Background(), "/core/auth/mfalogin", map[string]string{
		"sessionId":  sessionID,
		"code":       code,
		"authMethod": "session",
	}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Logout clears the session.
func (s *AuthService) Logout(ctx context.Context) error {
	return s.Post(ctx, "/core/auth/logout", nil, nil)
}

// Captcha fetches a captcha image. The result carries an imagePath, captchaID, and length.
func (s *AuthService) Captcha(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/captcha")
}

// LoginSetting returns the public login-page settings.
func (s *AuthService) LoginSetting(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/setting")
}

// WelcomePage returns the welcome text shown on the login screen.
func (s *AuthService) WelcomePage(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/welcome")
}

// MFAStatus returns the MFA state of the current user.
func (s *AuthService) MFAStatus(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/mfa", nil)
}

// BindMFA binds a TOTP MFA token to the current user.
func (s *AuthService) BindMFA(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/mfa/bind", body)
}

// CloseMFA disables MFA on the current user.
func (s *AuthService) CloseMFA(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/mfa/close", body)
}

// PasskeyBeginLogin starts a passkey assertion.
func (s *AuthService) PasskeyBeginLogin(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/passkey/begin", nil)
}

// PasskeyFinishLogin finishes a passkey assertion. sessionID is the value returned by PasskeyBeginLogin.
func (s *AuthService) PasskeyFinishLogin(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/passkey/finish", body)
}

// PasskeyRegisterBegin starts passkey registration for the logged-in user.
func (s *AuthService) PasskeyRegisterBegin(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/passkey/register/begin", nil)
}

// PasskeyRegisterFinish completes passkey registration.
func (s *AuthService) PasskeyRegisterFinish(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/passkey/register/finish", body)
}

// ListPasskeys returns the registered passkeys for the current user.
func (s *AuthService) ListPasskeys(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/passkey/list")
}

// DeletePasskey removes a passkey.
func (s *AuthService) DeletePasskey(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/passkey/del", body)
}

// LDAPStatus returns whether LDAP authentication is enabled.
func (s *AuthService) LDAPStatus(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/ldap/status")
}

// OIDCStatus returns the OIDC configuration status.
func (s *AuthService) OIDCStatus(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/oidc/status")
}

// OIDCBegin starts the OIDC login flow.
func (s *AuthService) OIDCBegin(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/oidc/begin", nil)
}

// OIDCFinish completes the OIDC login flow with the OIDC ticket.
func (s *AuthService) OIDCFinish(ctx context.Context, ticket string) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/oidc/finish", map[string]string{"ticket": ticket})
}

// SAML2Status returns the SAML2 configuration status.
func (s *AuthService) SAML2Status(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/saml2/status")
}

// SAML2Begin starts the SAML2 login flow.
func (s *AuthService) SAML2Begin(ctx context.Context) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/saml2/begin", nil)
}

// SAML2Finish completes the SAML2 login flow.
func (s *AuthService) SAML2Finish(ctx context.Context, ticket string) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/saml2/finish", map[string]string{"ticket": ticket})
}

// GenerateAPIKey mints a new API key for the current user.
func (s *AuthService) GenerateAPIKey(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/api/generate", body)
}

// UpdateAPIConfig updates the API key configuration (name, scopes, expiry).
func (s *AuthService) UpdateAPIConfig(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/api/update", body)
}

// CurrentUser returns the user profile of the current session.
func (s *AuthService) CurrentUser(ctx context.Context) (map[string]any, error) {
return s.getMap(ctx, "/core/auth/current")
}

// UpdateCurrentUser updates the current user profile.
func (s *AuthService) UpdateCurrentUser(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/current/update", body)
}

// ResetPassword changes the current user's password (used after expiry).
func (s *AuthService) ResetPassword(ctx context.Context, body map[string]any) (map[string]any, error) {
return s.postMap(ctx, "/core/auth/expired/reset", body)
}

// === Wildcard ===

// Call invokes an arbitrary /core/auth/* endpoint. Useful for endpoints not yet wrapped.
func (s *AuthService) Call(ctx context.Context, method, path string, body, out any) error {
	return s.Do(ctx, method, path, body, out)
}

// (helper to silence unused import)
var _ = http.MethodGet
