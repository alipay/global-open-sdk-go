package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireHoldListResponse struct {
	response.AlipayResponse
	Result      *model.Result              `json:"result,omitempty"`
	HoldDetails []*model.ReserveHoldDetail `json:"holdDetails,omitempty"`
	HasMore     bool                       `json:"hasMore,omitempty"`
}
