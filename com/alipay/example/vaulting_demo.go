package main

import (
	"fmt"
	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/vaulting"
	responseVaulting "github.com/alipay/global-open-sdk-go/com/alipay/api/response/vaulting"
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

	//createVaultingSession(client)
	//vaultPaymentMethod(client)
	inquireVaulting(client, "9116fffd-58d0-49ee-9fa1-4fec2c43c83d")

}

func createVaultingSession(client *defaultAlipayClient.DefaultAlipayClient) {
	request, vaultingRequest := vaulting.NewAlipayVaultingSessionRequest()
	vaultingRequest.VaultingRequestId = uuid.NewString()
	vaultingRequest.PaymentMethodType = "CARD"
	vaultingRequest.VaultingNotificationUrl = "https://www.yourNotifyUrl.com"
	vaultingRequest.RedirectUrl = "https://www.yourRedirectUrl.com"
	vaultingRequest.MerchantRegion = "BR"

	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vaultingRequest.VaultingRequestId)
		fmt.Println(response.(*responseVaulting.AlipayVaultingSessionResponse))
	}
}

func vaultPaymentMethod(client *defaultAlipayClient.DefaultAlipayClient) {
	request, vaultPaymentMethodRequest := vaulting.NewAlipayVaultingPaymentMethodRequest()
	vaultPaymentMethodRequest.VaultingRequestId = uuid.NewString()
	vaultPaymentMethodRequest.VaultingNotificationUrl = "https://www.yourNotifyUrl.com"
	vaultPaymentMethodRequest.RedirectUrl = "https://www.yourRedirectUrl.com"
	vaultPaymentMethodRequest.MerchantRegion = "BR"

	vaultPaymentMethodRequest.PaymentMethodDetail = &model.PaymentMethodDetail{
		PaymentMethodType: "CARD",
		Card: &model.CardPaymentMethodDetail{
			CardNo: "4111111111111111",
			Brand:  model.CardBrand_VISA,
			BillingAddress: &model.Address{
				Address1: "address1",
				Address2: "address2",
				City:     "city",
				State:    "state",
				ZipCode:  "zipcode",
			},
			CardholderName: &model.UserName{
				FirstName: "firstname",
				LastName:  "lastname",
			},
			ExpiryYear:  "2026",
			ExpiryMonth: "06",
			Cvv:         "123",
		},
	}

	vaultPaymentMethodRequest.Env = &model.Env{
		TerminalType: model.TerminalType_APP,
	}

	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.(*responseVaulting.AlipayVaultingPaymentMethodResponse))
	}
}

func inquireVaulting(client *defaultAlipayClient.DefaultAlipayClient, vaultingRequestId string) {
	request, inquireVaultingRequest := vaulting.NewAlipayVaultingQueryRequest()
	inquireVaultingRequest.VaultingRequestId = vaultingRequestId
	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.(*responseVaulting.AlipayVaultingQueryResponse).Result)
	}
}
