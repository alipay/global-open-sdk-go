package responseAuth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayAuthApplyTokenResponse struct {
	response.AlipayResponse
	Result                 *model.Result          `json:"result,omitempty"`
	AccessToken            string                 `json:"accessToken,omitempty"`
	AccessTokenExpiryTime  string                 `json:"accessTokenExpiryTime,omitempty"`
	RefreshToken           string                 `json:"refreshToken,omitempty"`
	RefreshTokenExpiryTime string                 `json:"refreshTokenExpiryTime,omitempty"`
	ExtendInfo             string                 `json:"extendInfo,omitempty"`
	UserLoginId            string                 `json:"userLoginId,omitempty"`
	PspCustomerInfo        *model.PspCustomerInfo `json:"pspCustomerInfo,omitempty"`
}
