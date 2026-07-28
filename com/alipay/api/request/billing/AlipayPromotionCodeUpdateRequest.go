package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPromotionCodeUpdateRequest struct {
	PromotionCodeId string            `json:"promotionCodeId,omitempty"`
	Status          string            `json:"status,omitempty"`
	MaxRedemptions  int32             `json:"maxRedemptions,omitempty"`
	ExpiryTime      string            `json:"expiryTime,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}

func NewAlipayPromotionCodeUpdateRequest() (*request.AlipayRequest, *AlipayPromotionCodeUpdateRequest) {
	alipayPromotionCodeUpdateRequest := &AlipayPromotionCodeUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPromotionCodeUpdateRequest, "/ams/api/v1/billing/promotionCode/update", &responseBilling.AlipayPromotionCodeUpdateResponse{})
	return alipayRequest, alipayPromotionCodeUpdateRequest
}

func (alipayPromotionCodeUpdateRequest *AlipayPromotionCodeUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPromotionCodeUpdateRequest, "/ams/api/v1/billing/promotionCode/update", &responseBilling.AlipayPromotionCodeUpdateResponse{})
}
