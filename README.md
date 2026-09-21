# Antom SDK for Go

Latest release: **1.4.1**

## Installation

```sh
go get github.com/alipay/global-open-sdk-go@latest
```

Requires Go 1.22.5+. Use a supported Go release with current security fixes.

## Quick start

- **API Key:** follow the [setup guide](docs/api-key-client.md) and run the [sandbox example](com/alipay/example/api_key_payment_session/main.go).
- **RSA:** follow the [configuration](#rsa-configuration) and [inline example](#payment) below.
- Browse [more examples](com/alipay/example) and the [API documentation](https://global.alipay.com/docs/).

API Key and RSA clients share request/response models. File uploads and notification
verification still require RSA credentials.

### API Key client

Set `ANTOM_GATEWAY_URL` and `ANTOM_API_KEY` in your server environment.
This initializes the client; see the [setup guide](docs/api-key-client.md) for a
complete sandbox request and its additional configuration.

```go
package main

import (
    "os"
    antom "github.com/alipay/global-open-sdk-go/com/alipay/api"
)

func main() {
    client, err := antom.NewApiKeyAlipayClient(
        os.Getenv("ANTOM_GATEWAY_URL"), os.Getenv("ANTOM_API_KEY"))
    if err != nil { panic(err) }
    defer client.Close()
    // Use client.Execute(request) with a business request.
}
```

### RSA configuration

Before using the RSA client, prepare these values from your Antom integration:

```properties
gatewayUrl=your_regional_https_gateway
clientId=your_client_id
merchantPrivateKey=your_merchant_private_key
alipayPublicKey=your_antom_public_key
```

The examples read them from `ANTOM_GATEWAY_URL`, `ANTOM_CLIENT_ID`,
`ANTOM_MERCHANT_PRIVATE_KEY`, and `ANTOM_PUBLIC_KEY`, respectively.
Keep keys in server-side configuration. Replace sample order data and callback
URLs with your own values before sending a request.

### Payment

Save as `main.go` in your application module. Run `go mod tidy` to resolve the
example's imports, configure the four environment variables above, then use
`go run .`.

```go
package main

import (
    "fmt"
    "os"
    antom "github.com/alipay/global-open-sdk-go/com/alipay/api"
    "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
    "github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
    responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
    "github.com/google/uuid"
)

func main() {
    client := antom.NewDefaultAlipayClient(
        os.Getenv("ANTOM_GATEWAY_URL"), os.Getenv("ANTOM_CLIENT_ID"),
        os.Getenv("ANTOM_MERCHANT_PRIVATE_KEY"), os.Getenv("ANTOM_PUBLIC_KEY"))
    payRequest, request := pay.NewAlipayPayRequest()

    request.PaymentRequestId = uuid.NewString()
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
    fmt.Println(response.Result)
}
```

## Upgrade notes

Billing integrations: `availableAmount` now uses `Amount`; the `AvailableAmount`
model has been removed.

## Meter event upload

`meter/createSession` uses the regular signed AMS transport. Call
`meter/uploadEvent` with `ExecuteWithHeaders` and the returned `X-Session-Id`.
The SDK uses the gateway URL configured on the client without sandbox path
rewriting, request signing, response signature verification, or automatic
retries. This API requires HTTP/2 and supports the Go version declared in
`go.mod`. For production use, build applications with a currently supported
Go release containing the latest security fixes.

See the [Meter example](com/alipay/example/meter_upload_event_demo.go) for a complete request.

## Advanced usage

`DefaultAlipayClient.Execute` handles JSON serialization, signing, HTTP calls, and
response verification. The following is an implementation excerpt from the
[client source](com/alipay/api/defaultAlipayClient.go), not a standalone application
example: private helpers are defined in the SDK package.

```go
type DefaultAlipayClient struct {
	GatewayUrl         string
	ClientId           string
	MerchantPrivateKey string
	AlipayPublicKey    string
	AgentToken         string
	IsSandboxMode      bool
	uploadGatewayUrl   string
}

func NewDefaultAlipayClient(gatewayUrl, clientId, merchantPrivateKey, alipayPublicKey string, agentTokenOpt ...string) *DefaultAlipayClient {
	isSandboxMode := strings.HasPrefix(clientId, "SANDBOX_")

	c := &DefaultAlipayClient{
		GatewayUrl:         gatewayUrl,
		ClientId:           clientId,
		MerchantPrivateKey: merchantPrivateKey,
		AlipayPublicKey:    alipayPublicKey,
		IsSandboxMode:      isSandboxMode,
	}
	if len(agentTokenOpt) > 0 {
		c.AgentToken = agentTokenOpt[0]
	}
	return c
}

func (alipayClient *DefaultAlipayClient) Execute(alipayRequest *request.AlipayRequest) (any, error) {
	reqPayload, err := json.Marshal(alipayRequest.Param)
	if err != nil {
		return nil, &exception.AlipayLibraryError{Message: "json.Marshal is fail " + err.Error()}
	}
	AdjustSandboxUrl(alipayClient.IsSandboxMode, alipayRequest)
	path := alipayRequest.Path
	httpMethod := alipayRequest.HttpMethod
	reqTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign, err := tools.GenSign(fmt.Sprintf("%s", httpMethod), path, alipayClient.ClientId, reqTime, string(reqPayload), alipayClient.MerchantPrivateKey)
	if err != nil {
		return nil, err
	}
	header := buildBaseHeader(reqTime, alipayClient.ClientId, alipayRequest.KeyVersion, sign, alipayClient.AgentToken)
	alipayResponse, err := alipayClient.httpDo(alipayClient.GatewayUrl+path, httpMethod, map[string]string{}, header, reqPayload, alipayRequest.AlipayResponse)
	if err != nil {
		return nil, err
	}
	return alipayResponse, nil
}
```

## Support

For integration questions, contact overseas_support@service.alibaba.com.
