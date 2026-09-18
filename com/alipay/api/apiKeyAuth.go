package defaultAlipayClient

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
)

var apiKeyPattern = regexp.MustCompile(`^(isak|irak)_(TEST|PROD)_[^_\s\x00-\x1f\x7f]+_[^\s\x00-\x1f\x7f]+$`)
var apiKeyPathPattern = regexp.MustCompile(`^/ams/(sandbox/)?api/[A-Za-z0-9_/-]+$`)
var apiKeyHeaderPattern = regexp.MustCompile("^[!#$%&'*+.^_`|~0-9A-Za-z-]+$")
var apiKeyReservedHeaders = map[string]bool{
	"authorization": true, "signature": true, "client-id": true, "request-time": true,
	"key-version": true, "keyversion": true, "agent-token": true, "content-type": true, "user-agent": true,
	"x-sdkversion": true, "sdk-version": true, "host": true, "content-length": true,
	"transfer-encoding": true, "connection": true, "proxy-authorization": true,
}

type apiKeyAuth struct {
	gatewayURL string
	key        string
	sandbox    bool
}

func apiKeyError(message string) error { return &exception.AlipayLibraryError{Message: message} }

func newApiKeyAuth(gatewayURL, key string) (*apiKeyAuth, error) {
	if !apiKeyPattern.MatchString(key) {
		return nil, apiKeyError("apiKey must be a Standard or Restricted TEST/PROD key")
	}
	u, err := url.Parse(gatewayURL)
	if err != nil || u.Scheme != "https" || u.Hostname() == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
		return nil, apiKeyError("gatewayUrl must be an HTTPS base URL")
	}
	return &apiKeyAuth{strings.TrimSuffix(gatewayURL, "/"), key, strings.SplitN(key, "_", 3)[1] == "TEST"}, nil
}

func (a *apiKeyAuth) path(path string) (string, error) {
	if !apiKeyPathPattern.MatchString(path) || strings.Contains(path, "//") {
		return "", apiKeyError("request path must be an ordinary /ams/api/ path")
	}
	path = strings.Replace(path, "/ams/sandbox/api/", "/ams/api/", 1)
	if a.sandbox {
		path = strings.Replace(path, "/ams/api/", "/ams/sandbox/api/", 1)
	}
	return path, nil
}

func (a *apiKeyAuth) redact(err error) error {
	return apiKeyError(strings.ReplaceAll(err.Error(), a.key, "[REDACTED]"))
}
