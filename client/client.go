// Package client provides the low-level HTTP client for 1Panel.
//
// 1Panel v2 panel exposes two API families under one HTTPS port:
//
//   - /core/...      Main panel APIs (auth, settings, scripts, command library, logs, groups, backups, ...).
//                     Served by the "core" process, reverse-proxied behind the public entrance path.
//   - /<resource>... Node-facing APIs (hosts, containers, apps, websites, databases, files, ...).
//                     Each method is invoked on a particular node, selected via the "CurrentNode" header.
//
// Authentication state:
//   - After /core/auth/login the server sets two cookies: a session cookie and "pcsrftoken".
//   - Non-GET requests must echo the csrf token in the "X-CSRF-Token" header.
//   - The "entrance" sub-path is base64-encoded and sent on the "EntranceCode" header for login.
//
// This client keeps a cookie jar, manages CSRF, and injects the entrance + node headers automatically.
package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Default timeout and base configuration.
const (
	DefaultTimeout = 60 * time.Second
	UserAgent      = "1panel-sdk/1.0 (+https://github.com/zy84338719/1panel-sdk)"
)

// Result is the standard 1Panel response envelope.
type Result struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
	TraceID string          `json:"trace_id,omitempty"`
}

// Common codes observed in the panel (see core/buserr and middleware).
const (
	CodeSuccess          = 0
	CodeAuthFail         = 401
	CodeForbidden        = 403
	CodePasswordExpired  = 313
	CodeNodeUnbind       = 433
	CodeLoading          = 434
	CodeXPackUnavailable = 408
	CodeEnterprise       = 410
)

// Error represents a non-success API response.
type Error struct {
	StatusCode int
	APIError   *Result
	URL        string
	Method     string
}

func (e *Error) Error() string {
	if e.APIError != nil {
		return fmt.Sprintf("1panel %s %s: code=%d message=%q trace_id=%s",
			e.Method, e.URL, e.APIError.Code, e.APIError.Message, e.APIError.TraceID)
	}
	return fmt.Sprintf("1panel %s %s: http=%d", e.Method, e.URL, e.StatusCode)
}

// IsCode returns true when the error carries the given 1Panel business code.
func (e *Error) IsCode(code int) bool {
	return e.APIError != nil && e.APIError.Code == code
}

// Config controls how a Client is constructed.
type Config struct {
	// BaseURL is the public panel URL, e.g. "https://1panel.example.com/1panel_entrance".
	// If empty, Endpoint is used directly.
	BaseURL string

	// Endpoint is the host[:port] of the core API. Mutually exclusive with BaseURL.
	// Example: "https://1panel.example.com:9999".
	Endpoint string

	// Entrance is the entrance sub-path used to obscure the panel.
	// Will be base64-encoded into the EntranceCode header on login. Optional.
	Entrance string

	// HTTPClient is an optional underlying http.Client. A cookie jar is always installed
	// unless the caller already provided one with a Jar set.
	HTTPClient *http.Client

	// Timeout is the per-request timeout. Defaults to DefaultTimeout.
	Timeout time.Duration

	// Language is sent as the "Accept-Language" header. Defaults to "zh-CN".
	Language string

	// NodeID is the default node id sent in the "CurrentNode" header for node-facing APIs.
	// Empty means "local" (the master node).
	NodeID string
}

// Client is the low-level HTTP client used by the high-level SDK.
type Client struct {
	cfg       Config
	hc        *http.Client
	csrfToken string
}

// New returns a configured Client. It does not perform any network I/O.
func New(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" && cfg.Endpoint == "" {
		return nil, errors.New("client: BaseURL or Endpoint is required")
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = DefaultTimeout
	}
	if cfg.Language == "" {
		cfg.Language = "zh-CN"
	}
	hc := cfg.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: cfg.Timeout}
	}
	if hc.Jar == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, fmt.Errorf("client: build cookie jar: %w", err)
		}
		hc.Jar = jar
	}
	if hc.Timeout <= 0 {
		hc.Timeout = cfg.Timeout
	}
	return &Client{cfg: cfg, hc: hc}, nil
}

// BaseURL returns the public base URL the client is talking to.
func (c *Client) BaseURL() string { return c.cfg.BaseURL }

// Endpoint returns the resolved endpoint (BaseURL if set, otherwise Endpoint).
func (c *Client) Endpoint() string {
	if c.cfg.BaseURL != "" {
		return c.cfg.BaseURL
	}
	return c.cfg.Endpoint
}

// SetNode switches the default node id used in CurrentNode headers.
func (c *Client) SetNode(id string) { c.cfg.NodeID = id }

// SetLanguage switches the Accept-Language header.
func (c *Client) SetLanguage(lang string) { c.cfg.Language = lang }

// Cookies returns the cookies associated with the configured endpoint URL.
func (c *Client) Cookies() []*http.Cookie { return c.CookiesFor(c.Endpoint()) }

// CSRFToken returns the current CSRF token (sourced from the "pcsrftoken" cookie).
func (c *Client) CSRFToken() string { return c.csrfToken }

// SyncCSRF reads the "pcsrftoken" cookie (set by the panel) and stores it for later use.
func (c *Client) SyncCSRF() {
	u := c.endpointURL()
	for _, ck := range c.hc.Jar.Cookies(u) {
		if ck.Name == "pcsrftoken" {
			c.csrfToken = ck.Value
			return
		}
	}
}

// SetCSRF overrides the CSRF token used for non-GET requests.
func (c *Client) SetCSRF(t string) { c.csrfToken = t }

// CookiesFor returns the cookies that the jar would send to the given absolute URL.
func (c *Client) CookiesFor(rawURL string) []*http.Cookie {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil
	}
	return c.hc.Jar.Cookies(u)
}

// do executes a single HTTP request and decodes the response.
func (c *Client) do(ctx context.Context, method, path string, body any, out any) error {
	return c.doWithNode(ctx, method, path, body, c.cfg.NodeID, out)
}

func (c *Client) doWithNode(ctx context.Context, method, path string, body any, node string, out any) error {
	full := c.resolve(path)
	var reader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("client: marshal body: %w", err)
		}
		reader = bytes.NewReader(buf)
	}
	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), full, reader)
	if err != nil {
		return fmt.Errorf("client: new request: %w", err)
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", c.cfg.Language)

	// EntranceCode only matters for login endpoints. Set when entrance is configured.
	if c.cfg.Entrance != "" {
		req.Header.Set("EntranceCode", base64.StdEncoding.EncodeToString([]byte(c.cfg.Entrance)))
	}
	if node != "" {
		req.Header.Set("CurrentNode", url.QueryEscape(node))
	}
	if c.cfg.BaseURL == "" && body == nil {
		// When talking directly to core, also hint node if set.
	}

	// Attach payload headers.
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// CSRF token for non-safe methods.
	if !isSafeMethod(method) {
		if c.csrfToken == "" {
			c.SyncCSRF()
		}
		if c.csrfToken != "" {
			req.Header.Set("X-CSRF-Token", c.csrfToken)
		}
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return fmt.Errorf("client: do request: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("client: read body: %w", err)
	}

	if resp.StatusCode >= 400 && !isJSONResponse(resp.Header) {
		return &Error{
			StatusCode: resp.StatusCode,
			URL:        full,
			Method:     method,
		}
	}

	// Decode the standard envelope. 1Panel always returns it (even for streamed logs).
	var env Result
	if err := json.Unmarshal(raw, &env); err != nil {
		// Some endpoints (file download, captcha image) return raw bytes. Surface them.
		if out != nil {
			// Best-effort: try a direct decode, otherwise wrap the raw bytes.
			if err2 := json.Unmarshal(raw, out); err2 != nil {
				return &Error{StatusCode: resp.StatusCode, URL: full, Method: method}
			}
		}
		return nil
	}

	// Make sure the CSRF cookie is loaded from the jar after the first call.
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusForbidden {
		c.SyncCSRF()
	}

	if env.Code != CodeSuccess {
		return &Error{
			StatusCode: resp.StatusCode,
			APIError:   &env,
			URL:        full,
			Method:     method,
		}
	}

	if out == nil || len(env.Data) == 0 {
		return nil
	}
	if err := json.Unmarshal(env.Data, out); err != nil {
		return fmt.Errorf("client: decode data: %w (raw=%s)", err, string(env.Data))
	}
	return nil
}

// Get issues a GET request and decodes the data field into out.
func (c *Client) Get(ctx context.Context, path string, out any) error {
	return c.do(ctx, http.MethodGet, path, nil, out)
}

// Post issues a POST request with a JSON body.
func (c *Client) Post(ctx context.Context, path string, body any, out any) error {
	return c.do(ctx, http.MethodPost, path, body, out)
}

// Put issues a PUT request with a JSON body.
func (c *Client) Put(ctx context.Context, path string, body any, out any) error {
	return c.do(ctx, http.MethodPut, path, body, out)
}

// Delete issues a DELETE request (body is sent as raw JSON if non-nil).
func (c *Client) Delete(ctx context.Context, path string, body any, out any) error {
	return c.do(ctx, http.MethodDelete, path, body, out)
}

// OnNode is a per-request override of the CurrentNode header.
func (c *Client) OnNode(nodeID string) *NodeClient {
	return &NodeClient{c: c, node: nodeID}
}

// NodeClient scopes every subsequent call to a specific node.
type NodeClient struct {
	c    *Client
	node string
}

// Get / Post / Put / Delete on NodeClient behave like the Client methods but always
// send the CurrentNode header.
func (n *NodeClient) Get(ctx context.Context, path string, out any) error {
	return n.c.doWithNode(ctx, http.MethodGet, path, nil, n.node, out)
}
func (n *NodeClient) Post(ctx context.Context, path string, body, out any) error {
	return n.c.doWithNode(ctx, http.MethodPost, path, body, n.node, out)
}
func (n *NodeClient) Put(ctx context.Context, path string, body, out any) error {
	return n.c.doWithNode(ctx, http.MethodPut, path, body, n.node, out)
}
func (n *NodeClient) Delete(ctx context.Context, path string, body, out any) error {
	return n.c.doWithNode(ctx, http.MethodDelete, path, body, n.node, out)
}

// Do executes an arbitrary HTTP method scoped to this node.
func (n *NodeClient) Do(ctx context.Context, method, path string, body, out any) error {
	return n.c.doWithNode(ctx, method, path, body, n.node, out)
}

// Raw exposes the underlying http.Client. Use for streaming endpoints (terminal ws, log download).
func (c *Client) Raw() *http.Client { return c.hc }

// Do executes an arbitrary HTTP method. Used by SDK sub-services to keep their
// own type signatures compact. Most callers should use Get/Post/Put/Delete.
func (c *Client) Do(ctx context.Context, method, path string, body, out any) error {
	return c.do(ctx, method, path, body, out)
}

// EndpointURL returns a *url.URL pointing at the endpoint root.
func (c *Client) endpointURL() *url.URL {
	u, _ := url.Parse(c.Endpoint())
	return u
}

// resolve joins Endpoint() and path. It ensures the path carries the /api/v2
// prefix used by 1Panel v2 (the "BasePath" declared in core/cmd/server/docs/swagger.json).
// If the caller already includes /api/v2 in either the BaseURL or the path, no
// additional prefix is added.
func (c *Client) resolve(path string) string {
	base := strings.TrimRight(c.Endpoint(), "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !hasAPIPrefix(base) && !hasAPIPrefix(path) {
		path = "/api/v2" + path
	}
	return base + path
}

func hasAPIPrefix(s string) bool {
	return strings.Contains(s, "/api/v2")
}

// SetEntrance re-encodes the entrance path used for the EntranceCode header.
func (c *Client) SetEntrance(e string) { c.cfg.Entrance = e }

// helpers ----------------------------------------------------------------

func isSafeMethod(m string) bool {
	switch strings.ToUpper(m) {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
		return true
	}
	return false
}

func isJSONResponse(h http.Header) bool {
	ct := h.Get("Content-Type")
	return strings.Contains(strings.ToLower(ct), "json")
}

// ParseCode extracts the 1Panel business code from a previously serialized Result.
func ParseCode(raw []byte) (int, error) {
	var env Result
	if err := json.Unmarshal(raw, &env); err != nil {
		return 0, err
	}
	return env.Code, nil
}

// AsInt converts anything returned by the panel into an int64 when possible.
func AsInt(v any) (int64, bool) {
	switch x := v.(type) {
	case int:
		return int64(x), true
	case int32:
		return int64(x), true
	case int64:
		return x, true
	case float64:
		return int64(x), true
	case string:
		n, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return 0, false
		}
		return n, true
	}
	return 0, false
}
