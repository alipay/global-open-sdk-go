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
	const alipayGatewayUrl = "https://open-sea.alipay.com"
	const alipayClientId = "SANDBOX_5YC31G2ZNMQK07357"
	const alipayMerchantPrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDhnOLYh3Tte5ELNzD6kr6TSN+F4oXaNlnKgx2aGf/xSSHIh1k+wihv6HbwPAexdtjpDQAgwEv4YXpdH3erQLuy3oBIBVsdbXWg09TRttyBeM8FzbMru6qR1TceypEPhR9W4/hP9DZEmn9XZmR7xR9KStDKJdnqSNr578IVFvp3hXUt2HfwoHbUwwOPbu8a66th9b1PyNJ5DOdSoTj52tvFMOyfCmKOn9U/bwtcT/EGEJFlIj1QjBSwlEeCBDUVwwKlo2ttMP5Omy7i4Lxi5NKAMdw+Cru49FqGEf9B/JKfovd6EcwrnUeDfXVltNrPJjdr19WzatDh0k0wE/9QT6EnAgMBAAECggEAZsbQdDFo69KRpZlT36I/3NqisNwbe+esidum/Y+Aj8tv72jxF+zc/PXaUOAX5RkuASSh/Ul8kj7dvlRacJJBr1868xQ1iLdXkZdOaOazluuQ66TkTTTlpB+MR6Oh538OYsfhU5L9sczr28XSWqvW8EIa0SvjFJ5x2tAFCxR3r0AqXFrRteHSPYI01sle9ynCq3frBQwX481N/T0YvDQNFiRw+YlzJwJsZqPooFA2H/o+AL+LQED7eevnLYvevNS4GGVkWNO7gfKFHJA3RCMJgRqsXfxs2SG2cBx6YBYCQ7JurP8veMr3NBf/OOGCZln4zY4Vl5bTXe5vxeT5gzm18QKBgQD/jCH38x0AIjx0zwpZyvcp6C/2PohVjb6B/TbAiMmaIjpei06DCXHrDiObTLoxguZfmA+ypiPTZBOwFEDo7wDJ8khQwRMx9ydPMiaWoFCvl5iSke2vs/ONxdw02zRGj2uivgqjDf+94eTS8aSTJ+7kt1KLq4ZQf80Efywv+0xVnwKBgQDiAy5MMU8oDiSun8FoSBX3SomjdOX/tg4hMZ2PKYnXTJFUJ5bkjLhgdsPO5WGcFGsdReuweTzKteIRmS0zvdekVIpWFchflyeIM3+OuI1ZJJG6t28Xg3e8VOXCD917fjLnOOmH3f7PV/rj59wVM0yPvGStlAbPP0kwrcci8Wo3eQKBgQDE/ujctGwhw0KppUVMbRtWEeiPQitlEGzQ1jtT9t661DH82hT/DNPlqLOoL2DFdCxVeup3BH5PojFPJn3XUw9fnkdDAWPju6xw768xpIouooV6T8ZUETvqiaG0mVrWHg+SmD+o7My+OxpjxuXgjwMpC201wFc9TRflpIeSwX1Z7wKBgQCpUgy7VC2jKoVctZ6ly2t5akQXSxqMKg4H3C3X9RypSVmPHGG1M59l1VP4imxIDBv7QEjEWu+qRfzphkIRA2asXBGPUJ5eztT0+u/TMnvijr0GjyoRCZMIaun+KviY7gCgrUh3W17sY1M4rpl44Ie5H0ClscIwPY9NgsMvcIFMsQKBgG/XoSq4KjB+/SFdLTH4ITLb6Q8rvWOOyu6OvTBCgfxhq4R+rBP/40bvd1Ax5Eawfwq/GDfUL5jpzoaJGXGXDI90eXdeDOHZB7rq/+un9De1jPLGc1Ty7YT3SctYAvFw8jo0K65eckL7AaiGHk39eOXrWmJVVchOVlkX8TayiTgk"
	const alipayAlipayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAq7zEydi4Q2VvUIb1Mjpm/I2R3NVWcSMddlhvHYJADZ07YOGjvlQPbH3iixhLMnk1Y0tT7Sw7B1Ov1kXDGUhnui/YmGQBDbz9vg4iPDXA8OU7TaSsIk2BbP4+uZoortx2AZu/ABTGBDvyhLyJBkNplJ7196Np+IMaxi2RlT2NEAV4vFIurUcfFl5vvMliyV1SacvIyONkurzixSLirlKBl35t1mGm44xqh7M40tcMScgdF8pIdkzVz0nAnBcGb0aTeD3YLQmYFFmbQhWIe7MAa0BPEK7sxTJ1x1PbRUCHEXiZURnPjZTD7FBsTfLlcGOlOe0DXB7mrWm+AP+fVBjbAwIDAQAB"

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipayGatewayUrl,
		alipayClientId,
		alipayMerchantPrivateKey,
		alipayAlipayPublicKey)

	doPay(client)
	//payQuery("1e6d724d-da95-407b-9802-6f2217c301d6", client)
	//refund("202408151940108001001886C0209996792", client)
	//queryRefund("ad53716a-81-4c4c-b151-216c5225e", client)
	//cancel("ad53716a-81-4c4c-b51-20916c5225e", client)
	//consult(client)
	//createSession(client)
}

func doPay(body *defaultAlipayClient.DefaultAlipayClient) {
	payRequest, request := pay.NewAlipayPayRequest()

	request.PaymentRequestId = uuid.NewString()
	order := &model.Order{}
	order.OrderDescription = "example order"
	order.ReferenceOrderId = uuid.NewString()
	order.OrderAmount = &model.Amount{"HKD", "100"}
	merchant := &model.Merchant{}
	merchant.ReferenceMerchantId = "1238rye8yr8erwer"
	merchant.MerchantMCC = "7011"
	merchant.MerchantName = "example merchant"
	merchant.Store = &model.Store{StoreMCC: "7011", ReferenceStoreId: "289285674", StoreName: "store 1111"}
	order.Merchant = merchant
	order.Env = &model.Env{OsType: model.OsType_ANDROID, TerminalType: model.TerminalType_WEB}
	request.Order = order

	request.PaymentAmount = &model.Amount{"HKD", "100"}

	request.PaymentNotifyUrl = "https://www.yourNotifyUrl.com"
	request.PaymentRedirectUrl = "https://www.yourRedirectUrl.com"

	request.PaymentMethod = &model.PaymentMethod{PaymentMethodType: model.ALIPAY_HK}

	request.ProductCode = model.ProductCodeType_CASHIER_PAYMENT

	execute, err := body.Execute(payRequest)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayResponse)
	fmt.Println("result: ", response.Result)
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
	refundRequest.RefundAmount = &model.Amount{"HKD", "100"}
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
	consultRequest.ProductCode = model.ProductCodeType_CASHIER_PAYMENT
	consultRequest.UserRegion = "SG"
	consultRequest.AllowedPaymentMethodRegions = []string{"SG", "HK", "CN"}
	consultRequest.Env = &model.Env{
		OsType:       model.OsType_IOS,
		TerminalType: model.TerminalType_APP,
	}
	consultRequest.PaymentAmount = &model.Amount{"USD", "100"}
	consultRequest.PaymentFactor = &model.PaymentFactor{
		PresentmentMode: model.PresentmentMode_UNIFIED,
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
		OrderAmount:      &model.Amount{"HKD", "100"},
		Buyer: &model.Buyer{
			ReferenceBuyerId: "111112132143434",
		},
	}
	createSessionRequest.PaymentAmount = &model.Amount{"HKD", "100"}
	createSessionRequest.ProductCode = model.ProductCodeType_CASHIER_PAYMENT
	createSessionRequest.PaymentMethod = &model.PaymentMethod{
		PaymentMethodType: model.SHOPEEPAY_SG,
	}
	createSessionRequest.PaymentNotifyUrl = "https://www.yourNotifyUrl.com"
	createSessionRequest.PaymentRedirectUrl = "https://www.yourRedirectUrl.com"
	createSessionRequest.Env = &model.Env{
		OsType:       model.OsType_IOS,
		TerminalType: model.TerminalType_WEB,
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
