package responseAuth

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayAuthRevokeTokenResponse struct {
	response.AlipayResponse
	ExtendInfo string `json:"extendInfo"`
}
