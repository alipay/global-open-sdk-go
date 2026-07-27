package model

type BillingDiscount struct {
	CouponId        string `json:"couponId,omitempty"`
	PromotionCodeId string `json:"promotionCodeId,omitempty"`
}
