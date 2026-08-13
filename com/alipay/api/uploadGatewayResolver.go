package defaultAlipayClient

import (
	"net/url"
	"strings"
)

var defaultUploadGateways = map[string]string{
	"open-sea-global.alipay.com": "https://open-big-sea.alipay.com",
	"open-sea.alipay.com":        "https://open-big-sea.alipay.com",
	"open-na-global.alipay.com":  "https://open-big-na.alipay.com",
	"open-na.alipay.com":         "https://open-big-na.alipay.com",
	"open-de-global.alipay.com":  "https://open-big-de-global.alipay.com",
}

func resolveUploadGateway(normalGatewayURL, explicitUploadGatewayURL string) (string, error) {
	if strings.TrimSpace(explicitUploadGatewayURL) != "" {
		return explicitUploadGatewayURL, nil
	}
	parsed, err := parseNormalGateway(normalGatewayURL)
	if err != nil {
		return "", err
	}
	mapped := defaultUploadGateways[strings.ToLower(parsed.Hostname())]
	if mapped == "" {
		return "", libraryError("no default file gateway mapping exists for " + parsed.Hostname() + "; configure uploadGatewayUrl explicitly")
	}
	return mapped, nil
}

func normalizeExplicitUploadGateway(uploadGatewayURL string) (string, error) {
	parsed, err := parseAbsoluteHTTPS(strings.TrimSpace(uploadGatewayURL))
	if err != nil {
		return "", libraryError("uploadGatewayUrl must be an absolute HTTPS base URL")
	}
	if parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || (parsed.EscapedPath() != "" && parsed.EscapedPath() != "/") {
		return "", libraryError("uploadGatewayUrl must not include user info, path, query, or fragment")
	}
	return "https://" + strings.ToLower(parsed.Host), nil
}

func parseNormalGateway(normalGatewayURL string) (*url.URL, error) {
	candidate := strings.TrimSpace(normalGatewayURL)
	if candidate == "" {
		return nil, libraryError("gatewayUrl can't be empty; configure uploadGatewayUrl explicitly")
	}
	if !strings.Contains(candidate, "://") {
		candidate = "https://" + candidate
	}
	parsed, err := parseAbsoluteHTTPS(candidate)
	if err != nil || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || (parsed.EscapedPath() != "" && parsed.EscapedPath() != "/") || (parsed.Port() != "" && parsed.Port() != "443") {
		return nil, libraryError("gatewayUrl can't be mapped to a file gateway; configure uploadGatewayUrl explicitly")
	}
	return parsed, nil
}

func parseAbsoluteHTTPS(value string) (*url.URL, error) {
	parsed, err := url.Parse(value)
	if err != nil || !parsed.IsAbs() || !strings.EqualFold(parsed.Scheme, "https") || parsed.Hostname() == "" {
		return nil, libraryError("URL must be an absolute HTTPS base URL")
	}
	return parsed, nil
}
