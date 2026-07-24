package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPromotionCodeUpdateResponse struct {
	response.AlipayResponse
	Result          *model.Result `json:"result,omitempty"`
	PromotionCodeId string        `json:"promotionCodeId,omitempty"`
	Status          string        `json:"status,omitempty"`
}
