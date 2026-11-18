package pay_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	requestPay "github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

func TestVaultingRequestMetadataAndResponseTypes(t *testing.T) {
	detail := &model.VaultingRequestPaymentMethodDetail{PaymentMethodType: "CARD", PaymentMethodMetaData: map[string]any{
		"recurringType": "UNSCHEDULED", "enabled": false, "count": 0, "items": []string{},
	}}
	body, err := json.Marshal(requestPay.AlipayVaultingPaymentMethodRequest{PaymentMethodDetail: detail})
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatal(err)
	}
	metadata := decoded["paymentMethodDetail"].(map[string]any)["paymentMethodMetaData"].(map[string]any)
	if metadata["recurringType"] != "UNSCHEDULED" || metadata["enabled"] != false || metadata["count"] != float64(0) || len(metadata["items"].([]any)) != 0 {
		t.Fatalf("metadata changed: %#v", metadata)
	}
	old := &model.PaymentMethodDetail{PaymentMethodType: "CARD"}
	response := responsePay.AlipayVaultingPaymentMethodResponse{PaymentMethodDetail: old}
	query := responsePay.AlipayVaultingQueryResponse{PaymentMethodDetail: old}
	if response.PaymentMethodDetail != old || query.PaymentMethodDetail != old {
		t.Fatal("response type changed")
	}
	if _, exists := reflect.TypeOf(*old).FieldByName("PaymentMethodMetaData"); exists {
		t.Fatal("request metadata leaked into responses")
	}
}
