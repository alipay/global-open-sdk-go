package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCouponInquireDetailsRequest struct {
	CouponId string `json:"couponId,omitempty"`
}

func NewAlipayCouponInquireDetailsRequest() (*request.AlipayRequest, *AlipayCouponInquireDetailsRequest) {
	alipayCouponInquireDetailsRequest := &AlipayCouponInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCouponInquireDetailsRequest, "/ams/api/v1/billing/coupon/inquireDetails", &responseBilling.AlipayCouponInquireDetailsResponse{})
	return alipayRequest, alipayCouponInquireDetailsRequest
}

func (alipayCouponInquireDetailsRequest *AlipayCouponInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCouponInquireDetailsRequest, "/ams/api/v1/billing/coupon/inquireDetails", &responseBilling.AlipayCouponInquireDetailsResponse{})
}
