package auth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
)

type AlipayAuthRevokeTokenRequest struct {
	AccessToken string `json:"accessToken,omitempty"`
	ExtendInfo  string `json:"extendInfo,omitempty"`
}

func NewAlipayAuthRevokeTokenRequest() (*request.AlipayRequest, *AlipayAuthRevokeTokenRequest) {
	alipayAuthRevokeTokenRequest := &AlipayAuthRevokeTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthRevokeTokenRequest, "/ams/api/v1/authorizations/revoke", &responseAuth.AlipayAuthRevokeTokenResponse{})
	return alipayRequest, alipayAuthRevokeTokenRequest
}

func (alipayAuthRevokeTokenRequest *AlipayAuthRevokeTokenRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAuthRevokeTokenRequest, "/ams/api/v1/authorizations/revoke", &responseAuth.AlipayAuthRevokeTokenResponse{})
}
