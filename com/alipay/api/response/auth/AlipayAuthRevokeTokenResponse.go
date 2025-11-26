package responseAuth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayAuthRevokeTokenResponse struct {
	response.AlipayResponse
	Result     *model.Result `json:"result,omitempty"`
	ExtendInfo string        `json:"extendInfo,omitempty"`
}
