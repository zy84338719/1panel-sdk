package client

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// API key sign methods supported by 1Panel.
//
// When APIKey is configured on the client every outgoing request gets two
// additional headers:
//
//	1Panel-Timestamp: <unix seconds>
//	1Panel-Token:      <hex(hex(hex))... of the chosen method>
//
// The server accepts both methods so that legacy clients keep working; new
// integrations should use HMAC-SHA256.
const (
	SignMethodMD5        = "md5"
	SignMethodHMACSHA256 = "hmac-sha256"
)

// SignMethod is the algorithm used to compute the 1Panel-Token header.
type SignMethod string

// Sign returns the value sent in the 1Panel-Token header for the given
// API key and Unix timestamp (seconds), using the configured method.
//
// Exposed for callers that want to compute the same signature themselves
// (e.g. for proxying or verification) without going through the HTTP client.
func (c *Client) Sign(apiKey, ts string) string {
	return SignToken(string(c.signMethod()), apiKey, ts)
}

func (c *Client) signMethod() SignMethod {
	if c.cfg.APISignMethod == "" {
		return SignMethodHMACSHA256
	}
	return SignMethod(c.cfg.APISignMethod)
}

// SignToken computes the 1Panel-Token value. Exported so callers can produce
// identical signatures for proxies, custom transports, or for verifying a
// signature the server echoed back.
func SignToken(method, apiKey, unixSec string) string {
	switch SignMethod(method) {
	case "", SignMethodHMACSHA256:
		return signHMACSHA256(apiKey, unixSec)
	case SignMethodMD5:
		return signMD5(apiKey, unixSec)
	default:
		// Unknown method: fall back to the recommended HMAC-SHA256 to avoid
		// silently generating an invalid signature.
		return signHMACSHA256(apiKey, unixSec)
	}
}

// signMD5 mirrors the legacy 1Panel v2 algorithm:
// Token = md5("1panel" + API_KEY + UnixTimestamp), hex encoded.
func signMD5(apiKey, ts string) string {
	sum := md5.New()
	sum.Write([]byte("1panel" + apiKey + ts))
	return hex.EncodeToString(sum.Sum(nil))
}

// signHMACSHA256 mirrors the recommended 1Panel v2 algorithm:
// Token = hmac_sha256(API_KEY, "1panel:" + UnixTimestamp), hex encoded.
func signHMACSHA256(apiKey, ts string) string {
	mac := hmac.New(sha256.New, []byte(apiKey))
	mac.Write([]byte("1panel:" + ts))
	return hex.EncodeToString(mac.Sum(nil))
}

// SetAPIKey enables (or replaces) API key signing for subsequent requests.
// Pass an empty key to disable. The sign method is taken from the original
// Config (defaults to HMAC-SHA256).
func (c *Client) SetAPIKey(key string) {
	c.cfg.APIKey = key
}

// SetAPISignMethod switches the signature algorithm. Use SignMethodMD5 for
// legacy 1Panel installations and SignMethodHMACSHA256 (the default) for new
// ones. Empty string resets to the default.
func (c *Client) SetAPISignMethod(method string) {
	if method == "" {
		c.cfg.APISignMethod = SignMethodHMACSHA256
		return
	}
	c.cfg.APISignMethod = method
}

// HasAPIKey reports whether the client is configured to use API key auth.
func (c *Client) HasAPIKey() bool { return c.cfg.APIKey != "" }

// APIKey returns the configured API key (empty when unset).
func (c *Client) APIKey() string { return c.cfg.APIKey }

// currentTimestamp returns the Unix-second timestamp used in the
// 1Panel-Timestamp header. Exposed for tests and proxies.
func currentTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// Ensure the import block is exercised even if the file ever drops a usage.
var _ = fmt.Sprintf
