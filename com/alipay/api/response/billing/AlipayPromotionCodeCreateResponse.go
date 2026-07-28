package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPromotionCodeCreateResponse struct {
	response.AlipayResponse
	Result          *model.Result `json:"result,omitempty"`
	PromotionCodeId string        `json:"promotionCodeId,omitempty"`
	Code            string        `json:"code,omitempty"`
	Status          string        `json:"status,omitempty"`
}
