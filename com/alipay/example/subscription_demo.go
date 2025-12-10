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
	const alipayGatewayUrl = ""
	const alipayClientId = ""
	const alipayMerchantPrivateKey = ""
	const alipayAlipayPublicKey = ""

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
		OsType:       model.ANDROID,
		TerminalType: model.WEB,
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
	fmt.Println(response)
}

func subscriptionCancel(client *defaultAlipayClient.DefaultAlipayClient, subscriptionId string) {
	request, cancelRequest := subscription.NewAlipaySubscriptionCancelRequest()
	cancelRequest.SubscriptionId = subscriptionId
	cancelRequest.CancellationType = model.CancellationType_CANCEL
	execute, err := client.Execute(request)
	if err != nil {
		panic(err)
	}
	response := execute.(*responseSubscription.AlipaySubscriptionCancelResponse)
	fmt.Println(response)
}
