package main

import (
	"fmt"
	defaultAlipayClient "github.com/alipay/global-open-sdk-go/com/alipay/api"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request/auth"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
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

	authConsult(client)
	//applyToken("281001139639787089651362", client)
	//revokeToken("28288803001291161724296551000BgIrDiWzU0171000529", client)
}

func authConsult(client *defaultAlipayClient.DefaultAlipayClient) {
	request, authConsultRequest := auth.NewAlipayAuthConsultRequest()
	authConsultRequest.AuthRedirectUrl = "https://www.alipay.com"
	authConsultRequest.AuthState = uuid.NewString()
	authConsultRequest.CustomerBelongsTo = model.ALIPAY_CN
	authConsultRequest.OsType = model.ANDROID
	authConsultRequest.OsVersion = "1.0.0"
	authConsultRequest.Scopes = []model.ScopeType{model.ScopeTypeAgreementPay}
	authConsultRequest.TerminalType = model.APP

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthConsultResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}

func applyToken(authCode string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, authApplyTokenRequest := auth.NewAlipayAuthApplyTokenRequest()
	authApplyTokenRequest.GrantType = model.GrantTypeAUTHORIZATION_CODE
	authApplyTokenRequest.CustomerBelongsTo = model.ALIPAY_CN
	authApplyTokenRequest.AuthCode = authCode

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthApplyTokenResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}

func revokeToken(accessToken string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, authRevokeTokenRequest := auth.NewAlipayAuthRevokeTokenRequest()
	authRevokeTokenRequest.AccessToken = accessToken
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthRevokeTokenResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}
