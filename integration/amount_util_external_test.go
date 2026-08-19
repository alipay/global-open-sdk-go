package integration_test

import (
	"testing"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/tools"
)

func TestAmountUtilFromExternalPackage(t *testing.T) {
	value, err := tools.ToAmount("10.25", "USD")
	if err != nil || value != "1025" {
		t.Fatalf("value=%s error=%v", value, err)
	}
	major, err := tools.FromAmount(value, "USD")
	if err != nil || major != "10.25" {
		t.Fatalf("major=%s error=%v", major, err)
	}
}
