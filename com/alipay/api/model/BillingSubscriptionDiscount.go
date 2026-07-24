package model

type BillingSubscriptionDiscount struct {
	CouponId        string `json:"couponId,omitempty"`
	PromotionCodeId string `json:"promotionCodeId,omitempty"`
}
