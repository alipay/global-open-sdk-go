package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayUpdateHoldResponse struct {
	response.AlipayResponse
	Result      *model.Result `json:"result,omitempty"`
	HoldId      string        `json:"holdId,omitempty"`
	ReleaseTime string        `json:"releaseTime,omitempty"`
}
