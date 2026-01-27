package main

import (
	"fmt"

	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/pay"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
	"github.com/google/uuid"
)

func main() {
	const alipayGatewayUrl = ""
	const alipayClientId = ""
	const alipayMerchantPrivateKey = ""
	const alipayAlipayPublicKey = ""

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipayGatewayUrl,
		alipayClientId,
		alipayMerchantPrivateKey,
		alipayAlipayPublicKey)

	//doPay(client)
	//payQuery("1e6d724d-da95-407b-9802-6f2217c301d6", client)
	//refund("202408151940108001001886C0209996792", client)
	//queryRefund("ad53716a-81-4c4c-b151-216c5225e", client)
	//cancel("ad53716a-81-4c4c-b51-20916c5225e", client)
	//consult(client)
	createSession(client)
}

func doPay(body *defaultAlipayClient.DefaultAlipayClient) {
	payRequest, request := pay.NewAlipayPayRequest()

	request.PaymentRequestId = uuid.NewString()
	order := &model.Order{}
	order.OrderDescription = "example order"
	order.ReferenceOrderId = uuid.NewString()
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

	request.PaymentMethod = &model.PaymentMethod{PaymentMethodType: model.ALIPAY_HK}

	request.ProductCode = model.CASHIER_PAYMENT

	execute, err := body.Execute(payRequest)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}

func payQuery(paymentRequestId string, body *defaultAlipayClient.DefaultAlipayClient) {
	queryRequest := pay.AlipayPayQueryRequest{}
	queryRequest.PaymentRequestId = paymentRequestId
	request := queryRequest.NewRequest()
	//1. 这里接收结果
	execute, err := body.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayQueryResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)
}

func refund(paymentId string, client *defaultAlipayClient.DefaultAlipayClient) {
	refundRequest := pay.AlipayRefundRequest{}
	refundRequest.RefundRequestId = uuid.NewString()
	refundRequest.PaymentId = paymentId
	refundRequest.RefundAmount = model.NewAmount("100", "HKD")
	refundRequest.RefundReason = "example refund"
	request := refundRequest.NewRequest()
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayRefundResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)

}

func queryRefund(refundRequestId string, client *defaultAlipayClient.DefaultAlipayClient) {
	queryRefundRequest := pay.AlipayInquiryRefundRequest{}
	queryRefundRequest.RefundRequestId = refundRequestId
	request := queryRefundRequest.NewRequest()
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayInquiryRefundResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)
}

func consult(client *defaultAlipayClient.DefaultAlipayClient) {
	consultRequest := &pay.AlipayPayConsultRequest{}
	request := consultRequest.NewRequest()
	consultRequest.SettlementStrategy = &model.SettlementStrategy{
		SettlementCurrency: "USD",
	}
	consultRequest.ProductCode = model.CASHIER_PAYMENT
	consultRequest.UserRegion = "SG"
	consultRequest.AllowedPaymentMethodRegions = []string{"SG", "HK", "CN"}
	consultRequest.Env = &model.Env{
		OsType:       model.IOS,
		TerminalType: model.APP,
	}
	consultRequest.PaymentAmount = model.NewAmount("1000", "USD")
	consultRequest.PaymentFactor = &model.PaymentFactor{
		PresentmentMode: model.BUNDLE,
	}

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayConsultResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)

}

func cancel(paymentRequestId string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, cancelRequest := pay.NewAlipayPayCancelRequest()
	cancelRequest.PaymentRequestId = paymentRequestId
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayCancelResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)

}

func createSession(client *defaultAlipayClient.DefaultAlipayClient) {
	request, createSessionRequest := pay.NewAlipayPaymentSessionRequest()
	createSessionRequest.PaymentRequestId = uuid.NewString()
	createSessionRequest.Order = &model.Order{
		OrderDescription: "example order",
		ReferenceOrderId: "289473927358748",
		OrderAmount:      model.NewAmount("100", "HKD"),
		Buyer: &model.Buyer{
			ReferenceBuyerId: "111112132143434",
		},
	}
	createSessionRequest.PaymentAmount = model.NewAmount("100", "HKD")
	createSessionRequest.ProductCode = model.CASHIER_PAYMENT
	createSessionRequest.PaymentMethod = &model.PaymentMethod{
		PaymentMethodType: "",
		PaymentMethodMetaData: map[string]any{
			"is3DSAuthentication": "true",
		},
	}
	createSessionRequest.PaymentNotifyUrl = "https://www.yourNotifyUrl.com"
	createSessionRequest.PaymentRedirectUrl = "https://www.yourRedirectUrl.com"
	createSessionRequest.Env = &model.Env{
		OsType:       model.IOS,
		TerminalType: model.APP,
	}

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPaymentSessionResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)
}
