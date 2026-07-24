package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPromotionCodeInquireDetailsRequest struct {
	PromotionCodeId string `json:"promotionCodeId,omitempty"`
}

func NewAlipayPromotionCodeInquireDetailsRequest() (*request.AlipayRequest, *AlipayPromotionCodeInquireDetailsRequest) {
	alipayPromotionCodeInquireDetailsRequest := &AlipayPromotionCodeInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPromotionCodeInquireDetailsRequest, "/ams/api/v1/billing/promotionCode/inquireDetails", &responseBilling.AlipayPromotionCodeInquireDetailsResponse{})
	return alipayRequest, alipayPromotionCodeInquireDetailsRequest
}

func (alipayPromotionCodeInquireDetailsRequest *AlipayPromotionCodeInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPromotionCodeInquireDetailsRequest, "/ams/api/v1/billing/promotionCode/inquireDetails", &responseBilling.AlipayPromotionCodeInquireDetailsResponse{})
}
