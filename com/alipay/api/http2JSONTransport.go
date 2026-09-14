package defaultAlipayClient

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/exception"
)

var sessionHTTP2Transport = &http.Transport{
	TLSClientConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
		NextProtos: []string{"h2"},
	},
	DisableCompression: true,
	DialTLSContext: func(ctx context.Context, network, address string) (net.Conn, error) {
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return nil, err
		}
		tlsConfig := &tls.Config{
			MinVersion: tls.VersionTLS12,
			NextProtos: []string{"h2"},
			ServerName: host,
		}
		dialer := &tls.Dialer{
			NetDialer: &net.Dialer{Timeout: connectTimeout},
			Config:    tlsConfig,
		}
		connection, err := dialer.DialContext(ctx, network, address)
		if err != nil {
			return nil, err
		}
		tlsConnection, ok := connection.(*tls.Conn)
		if !ok || tlsConnection.ConnectionState().NegotiatedProtocol != "h2" {
			connection.Close()
			return nil, fmt.Errorf("server did not negotiate HTTP/2")
		}
		return connection, nil
	},
	ForceAttemptHTTP2: true,
	IdleConnTimeout:   90 * time.Second,
}

func postHTTP2JSON(gatewayURL, path, sessionID string, requestBody []byte) ([]byte, error) {
	requestURL, err := buildSessionRequestURL(gatewayURL, path)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), totalTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "http.NewRequest failed: " + err.Error()}
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", sdkUserAgent())
	req.Header.Set(sessionHeader, sessionID)

	response, err := sessionHTTP2Transport.RoundTrip(req)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "HTTP/2 request failed: " + err.Error()}
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "failed to read HTTP/2 response: " + err.Error()}
	}
	if response.ProtoMajor != 2 {
		return nil, &exception.AlipayLibraryError{Message: "this API requires HTTP/2, but negotiated protocol was " + response.Proto}
	}
	if response.StatusCode != http.StatusOK {
		return nil, &exception.AlipayLibraryError{Message: fmt.Sprintf("response data error, HTTP status=%d, responseBody=%s", response.StatusCode, responseBody)}
	}
	return responseBody, nil
}

func buildSessionRequestURL(gatewayURL, path string) (string, error) {
	parsed, err := url.Parse(strings.TrimSpace(gatewayURL))
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || (parsed.Path != "" && parsed.Path != "/") {
		return "", &exception.AlipayLibraryError{Message: "gatewayUrl must be an HTTPS origin without path, query, fragment, or user info"}
	}
	if !strings.HasPrefix(path, "/") {
		return "", &exception.AlipayLibraryError{Message: "path must start with /"}
	}
	parsed.Path = path
	return parsed.String(), nil
}
