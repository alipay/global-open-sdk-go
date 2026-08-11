package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCouponUpdateRequest struct {
	CouponId       string `json:"couponId,omitempty"`
	CouponName     string `json:"couponName,omitempty"`
	Status         string `json:"status,omitempty"`
	RedeemBy       string `json:"redeemBy,omitempty"`
	Metadata       string `json:"metadata,omitempty"`
	MaxRedemptions int32  `json:"maxRedemptions,omitempty"`
}

func NewAlipayCouponUpdateRequest() (*request.AlipayRequest, *AlipayCouponUpdateRequest) {
	alipayCouponUpdateRequest := &AlipayCouponUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCouponUpdateRequest, "/ams/api/v1/billing/coupon/update", &responseBilling.AlipayCouponUpdateResponse{})
	return alipayRequest, alipayCouponUpdateRequest
}

func (alipayCouponUpdateRequest *AlipayCouponUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCouponUpdateRequest, "/ams/api/v1/billing/coupon/update", &responseBilling.AlipayCouponUpdateResponse{})
}
