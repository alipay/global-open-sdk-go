package auth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
)

type AlipayAuthApplyTokenRequest struct {
	GrantType         model.GrantType         `json:"grantType,omitempty"`
	CustomerBelongsTo model.CustomerBelongsTo `json:"customerBelongsTo,omitempty"`
	AuthCode          string                  `json:"authCode,omitempty"`
	RefreshToken      string                  `json:"refreshToken,omitempty"`
	ExtendInfo        string                  `json:"extendInfo,omitempty"`
	MerchantRegion    string                  `json:"merchantRegion,omitempty"`
}

func (alipayAuthApplyTokenRequest *AlipayAuthApplyTokenRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAuthApplyTokenRequest, model.AUTH_APPLY_TOKEN_PATH, &responseAuth.AlipayAuthApplyTokenResponse{})
}

func NewAlipayAuthApplyTokenRequest() (*request.AlipayRequest, *AlipayAuthApplyTokenRequest) {
	alipayAuthApplyTokenRequest := &AlipayAuthApplyTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthApplyTokenRequest, model.AUTH_APPLY_TOKEN_PATH, &responseAuth.AlipayAuthApplyTokenResponse{})
	return alipayRequest, alipayAuthApplyTokenRequest
}
