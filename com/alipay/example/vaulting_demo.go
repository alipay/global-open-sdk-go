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
	const alipayGatewayUrl = "https://open-sea.alipay.com"
	const alipayClientId = "SANDBOX_5YC31G2ZNMQK07357"
	const alipayMerchantPrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDhnOLYh3Tte5ELNzD6kr6TSN+F4oXaNlnKgx2aGf/xSSHIh1k+wihv6HbwPAexdtjpDQAgwEv4YXpdH3erQLuy3oBIBVsdbXWg09TRttyBeM8FzbMru6qR1TceypEPhR9W4/hP9DZEmn9XZmR7xR9KStDKJdnqSNr578IVFvp3hXUt2HfwoHbUwwOPbu8a66th9b1PyNJ5DOdSoTj52tvFMOyfCmKOn9U/bwtcT/EGEJFlIj1QjBSwlEeCBDUVwwKlo2ttMP5Omy7i4Lxi5NKAMdw+Cru49FqGEf9B/JKfovd6EcwrnUeDfXVltNrPJjdr19WzatDh0k0wE/9QT6EnAgMBAAECggEAZsbQdDFo69KRpZlT36I/3NqisNwbe+esidum/Y+Aj8tv72jxF+zc/PXaUOAX5RkuASSh/Ul8kj7dvlRacJJBr1868xQ1iLdXkZdOaOazluuQ66TkTTTlpB+MR6Oh538OYsfhU5L9sczr28XSWqvW8EIa0SvjFJ5x2tAFCxR3r0AqXFrRteHSPYI01sle9ynCq3frBQwX481N/T0YvDQNFiRw+YlzJwJsZqPooFA2H/o+AL+LQED7eevnLYvevNS4GGVkWNO7gfKFHJA3RCMJgRqsXfxs2SG2cBx6YBYCQ7JurP8veMr3NBf/OOGCZln4zY4Vl5bTXe5vxeT5gzm18QKBgQD/jCH38x0AIjx0zwpZyvcp6C/2PohVjb6B/TbAiMmaIjpei06DCXHrDiObTLoxguZfmA+ypiPTZBOwFEDo7wDJ8khQwRMx9ydPMiaWoFCvl5iSke2vs/ONxdw02zRGj2uivgqjDf+94eTS8aSTJ+7kt1KLq4ZQf80Efywv+0xVnwKBgQDiAy5MMU8oDiSun8FoSBX3SomjdOX/tg4hMZ2PKYnXTJFUJ5bkjLhgdsPO5WGcFGsdReuweTzKteIRmS0zvdekVIpWFchflyeIM3+OuI1ZJJG6t28Xg3e8VOXCD917fjLnOOmH3f7PV/rj59wVM0yPvGStlAbPP0kwrcci8Wo3eQKBgQDE/ujctGwhw0KppUVMbRtWEeiPQitlEGzQ1jtT9t661DH82hT/DNPlqLOoL2DFdCxVeup3BH5PojFPJn3XUw9fnkdDAWPju6xw768xpIouooV6T8ZUETvqiaG0mVrWHg+SmD+o7My+OxpjxuXgjwMpC201wFc9TRflpIeSwX1Z7wKBgQCpUgy7VC2jKoVctZ6ly2t5akQXSxqMKg4H3C3X9RypSVmPHGG1M59l1VP4imxIDBv7QEjEWu+qRfzphkIRA2asXBGPUJ5eztT0+u/TMnvijr0GjyoRCZMIaun+KviY7gCgrUh3W17sY1M4rpl44Ie5H0ClscIwPY9NgsMvcIFMsQKBgG/XoSq4KjB+/SFdLTH4ITLb6Q8rvWOOyu6OvTBCgfxhq4R+rBP/40bvd1Ax5Eawfwq/GDfUL5jpzoaJGXGXDI90eXdeDOHZB7rq/+un9De1jPLGc1Ty7YT3SctYAvFw8jo0K65eckL7AaiGHk39eOXrWmJVVchOVlkX8TayiTgk"
	const alipayAlipayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAq7zEydi4Q2VvUIb1Mjpm/I2R3NVWcSMddlhvHYJADZ07YOGjvlQPbH3iixhLMnk1Y0tT7Sw7B1Ov1kXDGUhnui/YmGQBDbz9vg4iPDXA8OU7TaSsIk2BbP4+uZoortx2AZu/ABTGBDvyhLyJBkNplJ7196Np+IMaxi2RlT2NEAV4vFIurUcfFl5vvMliyV1SacvIyONkurzixSLirlKBl35t1mGm44xqh7M40tcMScgdF8pIdkzVz0nAnBcGb0aTeD3YLQmYFFmbQhWIe7MAa0BPEK7sxTJ1x1PbRUCHEXiZURnPjZTD7FBsTfLlcGOlOe0DXB7mrWm+AP+fVBjbAwIDAQAB"

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
