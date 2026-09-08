package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireReleaseResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	ReleaseRequestId string        `json:"releaseRequestId,omitempty"`
	ReleaseId        string        `json:"releaseId,omitempty"`
	HoldId           string        `json:"holdId,omitempty"`
	ReleaseAmount    *model.Amount `json:"releaseAmount,omitempty"`
	ReleaseStatus    string        `json:"releaseStatus,omitempty"`
}
