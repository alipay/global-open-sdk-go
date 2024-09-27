package auth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
)

type AlipayAuthRevokeTokenRequest struct {
	AccessToken string `json:"accessToken,omitempty"`
	ExtendInfo  string `json:"extendInfo,omitempty"`
}

func (alipayAuthRevokeTokenRequest *AlipayAuthRevokeTokenRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAuthRevokeTokenRequest, model.AUTH_REVOKE_PATH, &responseAuth.AlipayAuthRevokeTokenResponse{})
}

func NewAlipayAuthRevokeTokenRequest() (*request.AlipayRequest, *AlipayAuthRevokeTokenRequest) {
	alipayAuthRevokeTokenRequest := &AlipayAuthRevokeTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthRevokeTokenRequest, model.AUTH_REVOKE_PATH, &responseAuth.AlipayAuthRevokeTokenResponse{})
	return alipayRequest, alipayAuthRevokeTokenRequest
}
