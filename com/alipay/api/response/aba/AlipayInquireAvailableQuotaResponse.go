package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireAvailableQuotaResponse struct {
	response.AlipayResponse
	AvailableQuota *model.Amount `json:"availableQuota,omitempty"`
	Result         *model.Result `json:"result,omitempty"`
}
