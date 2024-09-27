package auth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
)

type AlipayAuthQueryTokenRequest struct {
	AccessToken string `json:"accessToken"`
}

func NewAlipayAuthQueryTokenRequest() (*request.AlipayRequest, *AlipayAuthQueryTokenRequest) {
	alipayAuthQueryTokenRequest := &AlipayAuthQueryTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthQueryTokenRequest, model.AUTH_QUERY_PATH, &responseAuth.AlipayAuthQueryTokenResponse{})
	return alipayRequest, alipayAuthQueryTokenRequest
}
