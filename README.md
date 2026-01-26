```
Language：GO  
GO version：1.22.5+  
Tags：v1.2.37
Copyright：Ant financial services group  
```

#### 1 Please use the latest version

https://mvnrepository.com/artifact/com.alipay.global.sdk/global-open-sdk-go

```  
 go get github.com/alipay/global-open-sdk-go
```

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

