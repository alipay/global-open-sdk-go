// Read ANTOM_GATEWAY_URL, ANTOM_API_KEY, ANTOM_REDIRECT_URL and ANTOM_NOTIFY_URL from the environment.
// Use a TEST key with createPaymentSession permission for sandbox testing.
// This example creates a CARD session for USD 1.00; it does not complete a payment.
// Replace 127.0.0.1 with the buyer's IP. Notifications require a reachable endpoint.

package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	client "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	pay "github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
	"os"
)

func main() {
	if err := createSession(); err != nil {
		panic(err)
	}
}

func createSession() error {
	c, err := client.NewClient(client.ClientConfig{
		GatewayUrl: os.Getenv("ANTOM_GATEWAY_URL"),
		ApiKey:     os.Getenv("ANTOM_API_KEY"),
	})
	if err != nil {
		return err
	}
	id := make([]byte, 16)
	if _, err := rand.Read(id); err != nil {
		return err
	}
	requestID := hex.EncodeToString(id)
	request, body := pay.NewAlipayPaymentSessionRequest()
	amount := model.NewAmount("100", "USD")
	body.ProductCode = model.ProductCodeType_CASHIER_PAYMENT
	body.ProductScene = "CHECKOUT_PAYMENT"
	body.PaymentRequestId = requestID
	body.Order = &model.Order{ReferenceOrderId: requestID,
		OrderDescription: "API Key example", OrderAmount: amount}
	body.PaymentAmount = amount
	body.PaymentMethod = &model.PaymentMethod{PaymentMethodType: "CARD"}
	body.PaymentFactor = &model.PaymentFactor{IsAuthorization: false}
	body.SettlementStrategy = &model.SettlementStrategy{SettlementCurrency: "USD"}
	body.Env = &model.Env{TerminalType: model.TerminalType_WEB, ClientIp: "127.0.0.1"}
	body.PaymentRedirectUrl = os.Getenv("ANTOM_REDIRECT_URL")
	body.PaymentNotifyUrl = os.Getenv("ANTOM_NOTIFY_URL")

	// Transport errors return as error; also check the business result.
	raw, err := c.Execute(request)
	if err != nil {
		return err
	}
	response, ok := raw.(*responsePay.AlipayPaymentSessionResponse)
	if !ok || response == nil || response.Result == nil {
		return fmt.Errorf("missing payment session response/result")
	}
	if response.Result.ResultStatus != "S" || response.Result.ResultCode != "SUCCESS" {
		return fmt.Errorf("session creation was not successful: %s", response.Result.ResultCode)
	}
	if response.PaymentSessionId == "" {
		return fmt.Errorf("missing paymentSessionId")
	}
	// Use response.PaymentSessionData or the returned URL with your checkout.
	fmt.Println("Payment session created")
	return nil
}
