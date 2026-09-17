package defaultAlipayClient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
)

// ClientConfig selects one authentication mode when NewClient is called.
type ClientConfig struct {
	GatewayUrl         string
	ApiKey             string
	ClientId           string
	MerchantPrivateKey string
	AlipayPublicKey    string
	AgentToken         string
}

type apiKeyAuth struct {
	key      string
	gateway  string
	clientID string
	sandbox  bool
}

// NewClient creates an API Key client, or delegates a complete RSA configuration
// to the existing constructor. Authentication is fixed for the new instance.
func NewClient(config ClientConfig) (*DefaultAlipayClient, error) {
	if config.ApiKey == "" {
		if config.GatewayUrl == "" || config.ClientId == "" || config.MerchantPrivateKey == "" || config.AlipayPublicKey == "" {
			return nil, fmt.Errorf("provide API Key or complete RSA configuration")
		}
		return NewDefaultAlipayClient(config.GatewayUrl, config.ClientId, config.MerchantPrivateKey, config.AlipayPublicKey, config.AgentToken), nil
	}
	gateway, err := url.Parse(config.GatewayUrl)
	if err != nil || gateway.Scheme != "https" || gateway.Hostname() == "" || gateway.User != nil || (gateway.Path != "" && gateway.Path != "/") || gateway.RawQuery != "" || gateway.Fragment != "" || gateway.ForceQuery {
		return nil, fmt.Errorf("API Key requires an absolute HTTPS gateway without credentials, path, query or fragment")
	}
	if config.MerchantPrivateKey != "" || config.AlipayPublicKey != "" || config.AgentToken != "" {
		return nil, fmt.Errorf("API Key cannot be combined with RSA credentials or AgentToken")
	}
	parts := strings.SplitN(config.ApiKey, "_", 4)
	if len(parts) != 4 || (parts[0] != "isak" && parts[0] != "irak") || (parts[1] != "TEST" && parts[1] != "PROD") || !regexp.MustCompile(`^[A-Za-z0-9+/]+$`).MatchString(parts[2]) || !regexp.MustCompile(`^[A-Za-z0-9_-]+$`).MatchString(parts[3]) {
		return nil, fmt.Errorf("invalid API Key format; expected Standard or Restricted TEST/PROD key")
	}
	decoded, err := base64.RawStdEncoding.Strict().DecodeString(parts[2])
	if err != nil || len(decoded) == 0 || !utf8.Valid(decoded) {
		return nil, fmt.Errorf("invalid API Key client ID encoding")
	}
	if config.ClientId != "" && config.ClientId != string(decoded) {
		return nil, fmt.Errorf("ClientId does not match API Key")
	}
	auth := &apiKeyAuth{key: config.ApiKey, gateway: strings.TrimSuffix(config.GatewayUrl, "/"), clientID: string(decoded), sandbox: parts[1] == "TEST"}
	return &DefaultAlipayClient{GatewayUrl: auth.gateway, ClientId: auth.clientID, IsSandboxMode: auth.sandbox, apiKeyAuth: auth}, nil
}

func (auth *apiKeyAuth) path(path string) (string, error) {
	if strings.HasPrefix(path, "/ams/sandbox/api/") {
		path = "/ams/api/" + strings.TrimPrefix(path, "/ams/sandbox/api/")
	}
	if !strings.HasPrefix(path, "/ams/api/") || strings.ContainsAny(path, "?#") {
		return "", fmt.Errorf("API Key requires an ordinary /ams/api/ request path")
	}
	if auth.sandbox {
		path = strings.Replace(path, "/ams/api/", "/ams/sandbox/api/", 1)
	}
	return path, nil
}

func (c *DefaultAlipayClient) executeAPIKey(r *request.AlipayRequest, extra map[string]string) (any, error) {
	if r == nil {
		return nil, &exception.AlipayLibraryError{Message: "request must not be nil"}
	}
	auth := c.apiKeyAuth
	if r.ClientId != "" && r.ClientId != auth.clientID {
		return nil, &exception.AlipayLibraryError{Message: "request ClientId does not match API Key"}
	}
	path, err := auth.path(r.Path)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: strings.ReplaceAll(err.Error(), auth.key, "[REDACTED]")}
	}
	payload, err := json.Marshal(r.Param)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: strings.ReplaceAll(err.Error(), auth.key, "[REDACTED]")}
	}
	req, err := http.NewRequest(r.HttpMethod, auth.gateway+path, bytes.NewReader(payload))
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: strings.ReplaceAll(err.Error(), auth.key, "[REDACTED]")}
	}
	for key, value := range extra {
		lower := strings.ToLower(key)
		if strings.TrimSpace(key) != "" && !reservedHeaders[lower] && lower != "authorization" && lower != "key-version" && lower != "keyversion" && lower != "host" {
			req.Header.Set(key, value)
		}
	}
	req.Header.Set("Authorization", "Bearer "+auth.key)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", sdkUserAgent())
	req.Header.Set("SDK-VERSION", "global-open-sdk-go")
	transport := &http.Transport{DialContext: (&net.Dialer{Timeout: connectTimeout}).DialContext, ResponseHeaderTimeout: readTimeout}
	defer transport.CloseIdleConnections()
	client := &http.Client{Timeout: totalTimeout, Transport: transport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: strings.ReplaceAll(err.Error(), auth.key, "[REDACTED]")}
	}
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: strings.ReplaceAll(err.Error(), auth.key, "[REDACTED]")}
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, &exception.AlipayLibraryError{Message: fmt.Sprintf("API Key HTTP status %d: %s", rsp.StatusCode, strings.ReplaceAll(string(body), auth.key, "[REDACTED]"))}
	}
	if err := json.Unmarshal(body, r.AlipayResponse); err != nil {
		return nil, &exception.AlipayLibraryError{Message: "API Key response is not valid JSON for the response type"}
	}
	return r.AlipayResponse, nil
}
