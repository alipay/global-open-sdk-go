package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPromotionCodeInquireListRequest struct {
	CouponId      string `json:"couponId,omitempty"`
	Status        string `json:"status,omitempty"`
	StartingAfter string `json:"startingAfter,omitempty"`
	EndingBefore  string `json:"endingBefore,omitempty"`
	Limit         int32  `json:"limit,omitempty"`
	IncludeTotal  bool   `json:"includeTotal,omitempty"`
}

func NewAlipayPromotionCodeInquireListRequest() (*request.AlipayRequest, *AlipayPromotionCodeInquireListRequest) {
	alipayPromotionCodeInquireListRequest := &AlipayPromotionCodeInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPromotionCodeInquireListRequest, "/ams/api/v1/billing/promotionCode/inquireList", &responseBilling.AlipayPromotionCodeInquireListResponse{})
	return alipayRequest, alipayPromotionCodeInquireListRequest
}

func (alipayPromotionCodeInquireListRequest *AlipayPromotionCodeInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPromotionCodeInquireListRequest, "/ams/api/v1/billing/promotionCode/inquireList", &responseBilling.AlipayPromotionCodeInquireListResponse{})
}
