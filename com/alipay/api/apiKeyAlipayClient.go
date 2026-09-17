package defaultAlipayClient

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
)

// ApiKeyAlipayClient is independent of the RSA client and shares its business models.
type ApiKeyAlipayClient struct {
	auth   *apiKeyAuth
	client *http.Client
}

// String deliberately excludes authentication material, including in formatted logs.
func (c *ApiKeyAlipayClient) String() string   { return "ApiKeyAlipayClient{apiKey:[REDACTED]}" }
func (c *ApiKeyAlipayClient) GoString() string { return c.String() }

// ApiKeyClientOption configures an API Key client before its first use.
type ApiKeyClientOption func(*apiKeyClientOptions) error
type apiKeyClientOptions struct{ timeout time.Duration }

// WithApiKeyTimeout sets the total timeout for ordinary JSON requests.
func WithApiKeyTimeout(timeout time.Duration) ApiKeyClientOption {
	return func(o *apiKeyClientOptions) error {
		if timeout <= 0 {
			return apiKeyError("timeout must be positive")
		}
		o.timeout = timeout
		return nil
	}
}

// NewApiKeyAlipayClient needs only a regional HTTPS gateway and the complete key.
func NewApiKeyAlipayClient(gatewayURL, apiKey string, options ...ApiKeyClientOption) (*ApiKeyAlipayClient, error) {
	auth, err := newApiKeyAuth(gatewayURL, apiKey)
	if err != nil {
		return nil, err
	}
	config := apiKeyClientOptions{timeout: 30 * time.Second}
	for _, option := range options {
		if option == nil {
			return nil, apiKeyError("option cannot be nil")
		}
		if err := option(&config); err != nil {
			return nil, err
		}
	}
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         (&net.Dialer{Timeout: 15 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		TLSHandshakeTimeout: 15 * time.Second,
		IdleConnTimeout:     90 * time.Second,
	}
	return &ApiKeyAlipayClient{auth: auth, client: &http.Client{
		Transport: transport, Timeout: config.timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}}, nil
}

// Close releases idle connections; in-flight calls retain their normal lifecycle.
func (c *ApiKeyAlipayClient) Close() { c.client.CloseIdleConnections() }

// Execute preserves the SDK's response pointer and error convention.
func (c *ApiKeyAlipayClient) Execute(req *request.AlipayRequest) (any, error) {
	return c.ExecuteWithHeaders(req, nil)
}

// ExecuteWithHeaders delegates Session uploads before ordinary Bearer authentication.
func (c *ApiKeyAlipayClient) ExecuteWithHeaders(req *request.AlipayRequest, extraHeaders map[string]string) (any, error) {
	if req == nil {
		return nil, apiKeyError("request cannot be nil")
	}
	if requiresSessionHTTP2(req) {
		return executeSessionHTTP2(c.auth.gatewayURL, req, extraHeaders)
	}
	path, err := c.auth.path(req.Path)
	if err != nil {
		return nil, err
	}
	if !strings.EqualFold(req.HttpMethod, "POST") {
		return nil, apiKeyError("Only POST is supported for ordinary API requests")
	}
	payload, err := json.Marshal(req.Param)
	if err != nil {
		return nil, c.auth.redact(err)
	}
	// Do not provide GetBody: ordinary requests must not be transparently replayed.
	httpRequest, err := http.NewRequest("POST", c.auth.gatewayURL+path, io.NopCloser(bytes.NewReader(payload)))
	if err != nil {
		return nil, c.auth.redact(err)
	}
	httpRequest.ContentLength = int64(len(payload))
	httpRequest.Header.Set("Authorization", "Bearer "+c.auth.key)
	httpRequest.Header.Set("Content-Type", "application/json; charset=UTF-8")
	httpRequest.Header.Set("User-Agent", sdkUserAgent())
	httpRequest.Header.Set("SDK-VERSION", "global-open-sdk-go")
	for name, value := range extraHeaders {
		if !apiKeyHeaderPattern.MatchString(name) || strings.ContainsAny(value, "\r\n") {
			return nil, apiKeyError("Invalid custom header")
		}
		if !apiKeyReservedHeaders[strings.ToLower(name)] {
			httpRequest.Header.Set(name, value)
		}
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, c.auth.redact(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, c.auth.redact(err)
	}
	if response.StatusCode >= 300 && response.StatusCode < 400 {
		return nil, apiKeyError("Redirects are not supported")
	}
	// Like RSA, preserve a valid business response even on a non-200 status.
	if err := json.Unmarshal(body, req.AlipayResponse); err != nil {
		return nil, c.auth.redact(err)
	}
	return req.AlipayResponse, nil
}
