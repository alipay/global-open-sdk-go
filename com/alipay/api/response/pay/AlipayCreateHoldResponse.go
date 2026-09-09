package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateHoldResponse struct {
	response.AlipayResponse
	Result        *model.Result `json:"result,omitempty"`
	HoldRequestId string        `json:"holdRequestId,omitempty"`
	HoldId        string        `json:"holdId,omitempty"`
	HoldAmount    *model.Amount `json:"holdAmount,omitempty"`
	ReleaseTime   string        `json:"releaseTime,omitempty"`
	HoldStatus    string        `json:"holdStatus,omitempty"`
}
