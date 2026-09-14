package defaultAlipayClient

import (
	"strings"
	"testing"
)

func TestPostHTTP2JSONDoesNotRejectSupportedGoRuntime(t *testing.T) {
	_, err := postHTTP2JSON(
		"http://example.invalid",
		"/ams/api/v1/meter/uploadEvent",
		"session-id",
		[]byte("{}"),
	)
	if err == nil || !strings.Contains(err.Error(), "gatewayUrl must be an HTTPS origin") {
		t.Fatalf("expected HTTPS origin validation error, got %v", err)
	}
}
