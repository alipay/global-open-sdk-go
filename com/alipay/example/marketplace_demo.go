package main

import (
	"encoding/json"
	"fmt"
	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/marketplace"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
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

	register(client)
	//update(client, "mid_zhangtianren_ztr_20230807_180716_981")
	//queryBalance(client, "mid_zhangtianren_ztr_20230807_180716_981")
}

func register(client *defaultAlipayClient.DefaultAlipayClient) {
	request, registerRequest := marketplace.NewAlipayRegisterRequest()
	registerRequest.RegistrationRequestId = uuid.NewString()
	registerRequest.SettlementInfos = []*model.SettlementInfo{
		{
			SettlementCurrency: "BRL",
			SettlementBankAccount: &model.SettlementBankAccount{
				BranchCode:        "1231",
				RoutingNumber:     "12",
				BankRegion:        "BR",
				AccountType:       model.AccountType_CHECKING,
				AccountHolderTIN:  "12345678901",
				AccountHolderName: "Timi",
				BankAccountNo:     "12312412421",
				AccountHolderType: model.AccountHolderType_ENTERPRISE,
			},
		},
	}

	registerRequest.MerchantInfo = &model.MerchantInfo{
		LoginId:         uuid.NewString()[:8] + "wangzunj3ao.wzj@digital-engine.com",
		LegalEntityType: model.LegalEntityType_COMPANY,
		Company: &model.Company{
			LegalName: "legalName",
			RegisteredAddress: &model.Address{
				Region: "BR",
			},
			CompanyType: model.CompanyType_LTDA,
			OperatingAddress: &model.Address{
				Region:   "BR",
				Address1: "address1",
			},
			Attachments: []*model.Attachment{
				{
					AttachmentName: "1.jpg",
					FileKey:        "test",
					AttachmentType: string(model.AttachmentType_ARTICLES_OF_ASSOCIATION),
				},
				{
					AttachmentName: "2.jpg",
					FileKey:        "test",
					AttachmentType: string(model.AttachmentType_ASSOCIATION_ARTICLE),
				},
			},
			Certificates: &model.Certificate{
				CertificateNo:   "12312412",
				CertificateType: model.CertificateType_ENTERPRISE_REGISTRATION,
			},
		},
		BusinessInfo: &model.BusinessInfo{
			DoingBusinessAs: "businessName_DBA",
			Websites: []*model.WebSite{
				{
					Url: "www.global.alipay.com",
				},
			},
		},
		ReferenceMerchantId: uuid.NewString(),
		EntityAssociations: []*model.EntityAssociations{
			{
				Individual: &model.Individual{
					Certificates: []*model.Certificate{
						{
							CertificateNo:   "123214321",
							CertificateType: model.CertificateType_CPF,
						},
					},
					Name: &model.UserName{
						FirstName:  "firstName",
						LastName:   "lastName",
						MiddleName: "middleName",
						FullName:   "fullName",
					},
					DateOfBirth: "1990-01-01",
				},
				LegalEntityType: model.LegalEntityType_INDIVIDUAL,
				AssociationType: model.AssociationType_LEGAL_REPRESENTATIVE,
			},
			{
				Individual: &model.Individual{
					Certificates: []*model.Certificate{
						{
							CertificateNo:   "123214321",
							CertificateType: model.CertificateType_CPF,
							FileKeys:        []string{"wetrewqratewtewgewgewg"},
						},
					},
					Name: &model.UserName{
						FirstName:  "firstName",
						LastName:   "lastName",
						MiddleName: "middleName",
						FullName:   "fullName",
					},
					DateOfBirth: "1990-01-01",
				},
				LegalEntityType: model.LegalEntityType_INDIVIDUAL,
				AssociationType: model.AssociationType_UBO,
			},
		},
	}

	response, err := client.Execute(request)
	registerResponse := response.(*responseMarketplace.AlipayRegisterResponse)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(registerResponse.Result)
	}

}

func update(client *defaultAlipayClient.DefaultAlipayClient, referenceMerchantId string) {
	request, updateRequest := marketplace.NewAlipaySettlementInfoUpdateRequest()
	updateRequest.ReferenceMerchantId = referenceMerchantId
	updateRequest.UpdateRequestId = uuid.NewString()
	updateRequest.SettlementCurrency = "BRL"
	updateRequest.SettlementBankAccount = &model.SettlementBankAccount{
		BranchCode:        "1231",
		RoutingNumber:     "12",
		BankRegion:        "BR",
		AccountType:       model.AccountType_CHECKING,
		AccountHolderTIN:  "12345678901",
		AccountHolderName: "Timi",
		BankAccountNo:     "12312412421",
		AccountHolderType: model.AccountHolderType_ENTERPRISE,
	}

	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func queryBalance(client *defaultAlipayClient.DefaultAlipayClient, referenceMerchantId string) {
	request, queryBalanceRequest := marketplace.NewAlipayInquireBalanceRequest()
	queryBalanceRequest.ReferenceMerchantId = referenceMerchantId
	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		indent, _ := json.MarshalIndent(response.(*responseMarketplace.AlipayInquireBalanceResponse), "", "  ")
		fmt.Println(string(indent))
	}
}

func settleRequest(client *defaultAlipayClient.DefaultAlipayClient, paymentId string) {
	request, settleRequest := marketplace.NewAlipaySettleRequest()
	settleRequest.PaymentId = paymentId
	settleRequest.SettlementRequestId = uuid.NewString()
	settleRequest.SettlementDetails = []*model.SettlementDetail{
		{
			SettleTo: model.SettleToType_SELLER,
			SettlementAmount: &model.Amount{
				Currency: "BRL",
				Value:    "90",
			},
		},
		{
			SettleTo: model.SettleToType_MARKETPLACE,
			SettlementAmount: &model.Amount{
				Currency: "BRL",
				Value:    "10",
			},
		},
	}

	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func createPayout(client *defaultAlipayClient.DefaultAlipayClient) {
	request, createPayoutRequest := marketplace.NewAlipayCreatePayoutRequest()
	createPayoutRequest.TransferRequestId = uuid.NewString()
	createPayoutRequest.TransferFromDetail = &model.TransferFromDetail{
		TransferFromMethod: &model.PaymentMethod{
			PaymentMethodId:   uuid.NewString(),
			PaymentMethodType: model.BALANCE_ACCOUNT,
		},
	}
	createPayoutRequest.TransferToDetail = &model.TransferToDetail{
		TransferToMethod: &model.PaymentMethod{
			PaymentMethodId:   uuid.NewString(),
			PaymentMethodType: model.SETTLEMENT_CARD,
		},
		TransferToCurrency: "BRL",
		PurposeCode:        "GSD",
		TransferNotifyUrl:  "www.global.alipay.com",
	}

	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

}

func createTransfer(client *defaultAlipayClient.DefaultAlipayClient) {
	request, createTransferRequest := marketplace.NewAlipayCreateTransferRequest()
	createTransferRequest.TransferRequestId = uuid.NewString()

	createTransferRequest.TransferFromDetail = &model.TransferFromDetail{
		TransferFromMethod: &model.PaymentMethod{
			PaymentMethodId:   uuid.NewString(),
			PaymentMethodType: model.BALANCE_ACCOUNT,
		},
		TransferFromAmount: &model.Amount{"BRL", "100"},
	}

	createTransferRequest.TransferToDetail = &model.TransferToDetail{
		TransferToMethod: &model.PaymentMethod{
			PaymentMethodId:   uuid.NewString(),
			PaymentMethodType: model.BALANCE_ACCOUNT,
		},
		TransferToCurrency: "BRL",
		PurposeCode:        "GSD",
		TransferNotifyUrl:  "www.global.alipay.com",
	}

	response, err := client.Execute(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
