package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCouponInquireDetailsResponse struct {
	response.AlipayResponse
	Result          *model.ResultInfo                    `json:"result,omitempty"`
	CouponId        string                               `json:"couponId,omitempty"`
	CouponRequestId string                               `json:"couponRequestId,omitempty"`
	CouponName      string                               `json:"couponName,omitempty"`
	DiscountType    string                               `json:"discountType,omitempty"`
	PercentOff      string                               `json:"percentOff,omitempty"`
	AmountOff       *model.Amount                        `json:"amountOff,omitempty"`
	DurationType    string                               `json:"durationType,omitempty"`
	DurationValue   int32                                `json:"durationValue,omitempty"`
	DurationUnit    string                               `json:"durationUnit,omitempty"`
	MaxRedemptions  int32                                `json:"maxRedemptions,omitempty"`
	RedeemedCount   int32                                `json:"redeemedCount,omitempty"`
	RedeemBy        string                               `json:"redeemBy,omitempty"`
	AppliesTo       *model.CouponInquireDetailsAppliesTo `json:"appliesTo,omitempty"`
	Status          string                               `json:"status,omitempty"`
	Metadata        map[string]string                    `json:"metadata,omitempty"`
	CreateTime      string                               `json:"createTime,omitempty"`
}
