package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCouponInquireListRequest struct {
	Status        string `json:"status,omitempty"`
	DiscountType  string `json:"discountType,omitempty"`
	StartingAfter string `json:"startingAfter,omitempty"`
	EndingBefore  string `json:"endingBefore,omitempty"`
	Limit         int32  `json:"limit,omitempty"`
	IncludeTotal  bool   `json:"includeTotal,omitempty"`
}

func NewAlipayCouponInquireListRequest() (*request.AlipayRequest, *AlipayCouponInquireListRequest) {
	alipayCouponInquireListRequest := &AlipayCouponInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCouponInquireListRequest, "/ams/api/v1/billing/coupon/inquireList", &responseBilling.AlipayCouponInquireListResponse{})
	return alipayRequest, alipayCouponInquireListRequest
}

func (alipayCouponInquireListRequest *AlipayCouponInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCouponInquireListRequest, "/ams/api/v1/billing/coupon/inquireList", &responseBilling.AlipayCouponInquireListResponse{})
}
