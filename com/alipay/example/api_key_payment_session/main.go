// Creates one sandbox payment session using an authorized Restricted TEST key.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	antom "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	for _, name := range []string{"ANTOM_GATEWAY_URL", "ANTOM_API_KEY", "ANTOM_NOTIFY_URL"} {
		if strings.TrimSpace(os.Getenv(name)) == "" {
			return fmt.Errorf("missing environment variable: %s", name)
		}
	}
	key := os.Getenv("ANTOM_API_KEY")
	if !strings.HasPrefix(key, "irak_TEST_") {
		return fmt.Errorf("this example requires a Restricted TEST key")
	}
	client, err := antom.NewApiKeyAlipayClient(os.Getenv("ANTOM_GATEWAY_URL"), key)
	if err != nil {
		return err
	}
	defer client.Close()
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	amount := model.NewAmount("100", "USD")
	request, business := pay.NewAlipayPaymentSessionRequest()
	business.ProductCode = model.CASHIER_PAYMENT
	business.ProductScene = "CHECKOUT_PAYMENT"
	business.PaymentRequestId = "example-session-" + id
	business.Order = &model.Order{ReferenceOrderId: "example-order-" + id,
		OrderDescription: "API Key sandbox example", OrderAmount: amount}
	business.PaymentAmount = amount
	business.PaymentMethod = &model.PaymentMethod{PaymentMethodType: "CARD"}
	business.PaymentFactor = &model.PaymentFactor{IsAuthorization: false}
	business.SettlementStrategy = &model.SettlementStrategy{SettlementCurrency: "USD"}
	business.Env = &model.Env{TerminalType: model.WEB, ClientIp: "127.0.0.1"}
	// Use a separate redirect page in a real integration.
	business.PaymentRedirectUrl = os.Getenv("ANTOM_NOTIFY_URL")
	business.PaymentNotifyUrl = os.Getenv("ANTOM_NOTIFY_URL")
	result, err := client.Execute(request)
	if err != nil {
		return err
	}
	response := result.(*responsePay.AlipayPaymentSessionResponse)
	// Local debugging only: the response contains payment-session credentials.
	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	if response.Result == nil || response.Result.ResultCode != "SUCCESS" || response.Result.ResultStatus != "S" {
		return fmt.Errorf("createPaymentSession failed; see the response result")
	}
	return nil
}
