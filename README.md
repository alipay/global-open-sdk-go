```
Language：GO  
GO version：1.22.5+  
Tags：v1.4.0
Copyright：Ant financial services group  
```

#### 1 Please use the latest version

https://mvnrepository.com/artifact/com.alipay.global.sdk/global-open-sdk-go

```  
 go get github.com/alipay/global-open-sdk-go
```

#### Meter event upload

`meter/createSession` uses the regular signed AMS transport. Call
`meter/uploadEvent` with `ExecuteWithHeaders` and the returned `X-Session-Id`.
The SDK uses the gateway URL configured on the client without sandbox path
rewriting, request signing, response signature verification, or automatic
retries. This API requires HTTP/2 and supports the Go version declared in
`go.mod`. For production use, build applications with a currently supported
Go release containing the latest security fixes.

See `com/alipay/example/meter_upload_event_demo.go` for a complete request.

#### 2 The demo code for create payment
```
   	payRequest, request := pay.NewAlipayPayRequest()

	request.PaymentRequestId = "ad716a-81-4c4c-b51-20916c5225e"
	order := &model.Order{}
	order.OrderDescription = "example order"
	order.ReferenceOrderId = "28947397358748"
	order.OrderAmount = model.NewAmount("100", "HKD")
	merchant := &model.Merchant{}
	merchant.ReferenceMerchantId = "1238rye8yr8erwer"
	merchant.MerchantMCC = "7011"
	merchant.MerchantName = "example merchant"
	merchant.Store = &model.Store{StoreMCC: "7011", ReferenceStoreId: "289285674", StoreName: "store 1111"}
	order.Merchant = merchant
	order.Env = &model.Env{OsType: model.ANDROID, TerminalType: model.WEB}
	request.Order = order

	request.PaymentAmount = model.NewAmount("100", "HKD")

	request.PaymentNotifyUrl = "https://www.yourNotifyUrl.com"
	request.PaymentRedirectUrl = "https://www.yourRedirectUrl.com"

	request.PaymentMethod = &model.PaymentMethod{PaymentMethodType: model.ALIPAY_HK, PaymentMethodId: "1234567890"}

	request.ProductCode = model.CASHIER_PAYMENT

	execute, err := client.Execute(payRequest)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayResponse)
```

The execute method contains the HTTP request to the gateway.
```
type DefaultAlipayClient struct {
	GatewayUrl         string
	ClientId           string
	MerchantPrivateKey string
	AlipayPublicKey    string
	IsSandboxMode      bool
}

func NewDefaultAlipayClient(gatewayUrl string, clientId string, merchantPrivateKey string, alipayPublicKey string) *DefaultAlipayClient {
	isSandboxMode := false
	if strings.HasPrefix(clientId, "SANDBOX_") {
		isSandboxMode = true
	}

	return &DefaultAlipayClient{
		GatewayUrl:         gatewayUrl,
		ClientId:           clientId,
		MerchantPrivateKey: merchantPrivateKey,
		AlipayPublicKey:    alipayPublicKey,
		IsSandboxMode:      isSandboxMode,
	}
}

func (alipayClient *DefaultAlipayClient) Execute(alipayRequest *request.AlipayRequest) (any, error) {
	reqPayload, err := json.Marshal(alipayRequest.Param)
	if err != nil {
		return nil, &exception.AlipaySDKError{Message: "json.Marshal is fail " + err.Error()}
	}
	path := alipayRequest.Path
	httpMethod := alipayRequest.HttpMethod
	reqTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	sign, err := genSign(fmt.Sprintf("%s", httpMethod), path, alipayClient.ClientId, reqTime, string(reqPayload), getPkcsKeu(alipayClient.MerchantPrivateKey))
	if err != nil {
		return nil, err
	}
	header := buildBaseHeader(reqTime, alipayClient.ClientId, alipayRequest.KeyVersion, sign)
	alipayResponse, err := alipayClient.httpDo(alipayClient.GatewayUrl+path, httpMethod, map[string]string{}, header, reqPayload, alipayRequest.AlipayResponse)
	if err != nil {
		return nil, err
	}
	return alipayResponse, nil
}

```

## API Key authentication

See the [complete API Key example](com/alipay/example/apikey/main.go). From the repository root, run `go run ./com/alipay/example/apikey`.

Initialize the client with your gateway URL and API Key; existing RSA usage remains supported.
This feature is available in the current source branch and has not been published yet.

Set `ANTOM_GATEWAY_URL` to your regional HTTPS gateway (for example,
`https://open-sea-global.alipay.com` for Asia), `ANTOM_API_KEY` to your key,
`ANTOM_REDIRECT_URL` to your checkout return URL, and `ANTOM_NOTIFY_URL` to your
notification endpoint. The application reads these variables; the SDK does not load them automatically.

The example creates a CARD payment session for USD 1.00 (`100` minor units),
with USD settlement. Use a merchant configured for this combination and a key
with createPaymentSession permission. Replace the example client IP with the
buyer's IP in your application. Exceptions propagate to the caller; a normal
response must still be checked for business success.

```go
package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "os"
    client "github.com/alipay/global-open-sdk-go/com/alipay/api"
    "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
    pay "github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
    responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

func main() {
    if err := createSession(); err != nil { panic(err) }
}

func createSession() error {
    c, err := client.NewClient(client.ClientConfig{
        GatewayUrl: os.Getenv("ANTOM_GATEWAY_URL"),
        ApiKey: os.Getenv("ANTOM_API_KEY"),
    })
    if err != nil { return err }
    id := make([]byte, 16)
    if _, err := rand.Read(id); err != nil { return err }
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

    raw, err := c.Execute(request)
    if err != nil { return err }
    response, ok := raw.(*responsePay.AlipayPaymentSessionResponse)
    if !ok || response == nil || response.Result == nil {
        return fmt.Errorf("missing payment session response/result")
    }
    if response.Result.ResultStatus != "S" || response.Result.ResultCode != "SUCCESS" {
        return fmt.Errorf("session creation was not successful: %s", response.Result.ResultCode)
    }
    if response.PaymentSessionId == "" { return fmt.Errorf("missing paymentSessionId") }
    // Use response.PaymentSessionData or the returned URL with your checkout.
    fmt.Println("Payment session created")
    return nil
}
```

- Standard and Restricted keys use the same client. TEST/PROD in the key selects
  the request environment; do not add a sandbox path to the gateway URL.
- Creating a session does not mean payment is complete. Notifications still use
  the existing signature verification mechanism.
- File upload is not supported with API Key authentication.
