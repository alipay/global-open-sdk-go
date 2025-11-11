package main

import (
	"fmt"
	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/subscription"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
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

	//SubscriptionsCreate(client)
	SubscriptionsChange(client, "202409141900000000000001J0000009488")
	//subscriptionCancel(client, "202409141900000000000001J0000009488")

}

func SubscriptionsCreate(client *defaultAlipayClient.DefaultAlipayClient) {

	request, alipaySubscriptionCreateRequest := subscription.NewAlipaySubscriptionCreateRequest()
	alipaySubscriptionCreateRequest.SubscriptionRequestId = uuid.NewString()
	alipaySubscriptionCreateRequest.Env = &model.Env{
		ClientIp:     "1.*.*.*",
		OsType:       model.OsType_ANDROID,
		TerminalType: model.TerminalType_APP,
	}
	alipaySubscriptionCreateRequest.PaymentAmount = &model.Amount{
		Currency: "HKD",
		Value:    "10",
	}
	alipaySubscriptionCreateRequest.PaymentNotificationUrl = "https://www.yourNotifyUrl.com"
	alipaySubscriptionCreateRequest.PeriodRule = &model.PeriodRule{
		PeriodType:  string(model.PeriodType_MONTH),
		PeriodCount: 1,
	}
	alipaySubscriptionCreateRequest.SettlementStrategy = &model.SettlementStrategy{
		SettlementCurrency: "USD",
	}
	alipaySubscriptionCreateRequest.SubscriptionDescription = "test_subscription"
	alipaySubscriptionCreateRequest.SubscriptionStartTime = "2024-09-12T12:01:01+08:00"
	alipaySubscriptionCreateRequest.SubscriptionEndTime = "2024-09-14T12:01:01+08:00"
	// The duration of subscription preparation process should be less than 48 hours
	alipaySubscriptionCreateRequest.SubscriptionExpiryTime = "2024-09-15T12:01:01+08:00"
	alipaySubscriptionCreateRequest.PaymentNotificationUrl = "https://www.yourNotifyUrl.com"

	alipaySubscriptionCreateRequest.OrderInfo = &model.OrderInfo{
		OrderAmount: &model.Amount{
			Currency: "HKD",
			Value:    "10",
		},
	}

	alipaySubscriptionCreateRequest.PaymentMethod = &model.PaymentMethod{
		PaymentMethodType: model.ALIPAY_HK,
	}

	alipaySubscriptionCreateRequest.SubscriptionRedirectUrl = "https://www.alipay.com"
	alipaySubscriptionCreateRequest.SubscriptionNotificationUrl = "https://www.alipay.com"

	execute, err := client.Execute(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(alipaySubscriptionCreateRequest.SubscriptionRequestId)
	response := execute.(*responseSubscription.AlipaySubscriptionCreateResponse)
	fmt.Println(response.Result)
	fmt.Println(response)
}

func SubscriptionsChange(client *defaultAlipayClient.DefaultAlipayClient, subscriptionId string) {
	request, changeRequest := subscription.NewAlipaySubscriptionChangeRequest()
	changeRequest.SubscriptionId = subscriptionId
	changeRequest.SubscriptionChangeRequestId = uuid.NewString()
	changeRequest.PaymentAmountDifference = &model.Amount{
		Currency: "HKD",
		Value:    "100",
	}
	changeRequest.PaymentAmount = &model.Amount{
		Currency: "HKD",
		Value:    "100",
	}
	changeRequest.PeriodRule = &model.PeriodRule{
		PeriodType:  string(model.PeriodType_MONTH),
		PeriodCount: 1,
	}
	changeRequest.SubscriptionStartTime = "2024-09-12T12:01:01+08:00"
	changeRequest.SubscriptionEndTime = "2024-09-13T12:01:01+08:00"
	changeRequest.OrderInfo = &model.OrderInfo{
		OrderAmount: &model.Amount{
			Currency: "BRL",
			Value:    "100",
		},
	}

	execute, err := client.Execute(request)
	if err != nil {
		panic(err)
	}
	response := execute.(*responseSubscription.AlipaySubscriptionChangeResponse)
	fmt.Println(response.Result)
	fmt.Println(response)
}

func subscriptionCancel(client *defaultAlipayClient.DefaultAlipayClient, subscriptionId string) {
	request, cancelRequest := subscription.NewAlipaySubscriptionCancelRequest()
	cancelRequest.SubscriptionId = subscriptionId
	cancelRequest.CancellationType = string(model.CancellationType_CANCEL)
	execute, err := client.Execute(request)
	if err != nil {
		panic(err)
	}
	response := execute.(*responseSubscription.AlipaySubscriptionCancelResponse)
	fmt.Println(response)
}
