package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCouponUpdateResponse struct {
	response.AlipayResponse
	Result         *model.CouponUpdateResult `json:"result,omitempty"`
	CouponId       string                    `json:"couponId,omitempty"`
	CouponName     string                    `json:"couponName,omitempty"`
	Status         string                    `json:"status,omitempty"`
	MaxRedemptions int32                     `json:"maxRedemptions,omitempty"`
	RedeemedCount  int32                     `json:"redeemedCount,omitempty"`
	RedeemBy       string                    `json:"redeemBy,omitempty"`
	Metadata       map[string]string         `json:"metadata,omitempty"`
	UpdateTime     string                    `json:"updateTime,omitempty"`
}
