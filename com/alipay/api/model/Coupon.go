package model

type Coupon struct {
	CouponId     string  `json:"couponId,omitempty"`
	CouponName   string  `json:"couponName,omitempty"`
	DiscountType string  `json:"discountType,omitempty"`
	PercentOff   string  `json:"percentOff,omitempty"`
	AmountOff    *Amount `json:"amountOff,omitempty"`
	RedeemBy     string  `json:"redeemBy,omitempty"`
	Status       string  `json:"status,omitempty"`
	CreateTime   string  `json:"createTime,omitempty"`
}
