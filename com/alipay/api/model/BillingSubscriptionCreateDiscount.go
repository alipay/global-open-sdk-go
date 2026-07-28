package model

type BillingSubscriptionCreateDiscount struct {
	CouponId        string `json:"couponId,omitempty"`
	PromotionCodeId string `json:"promotionCodeId,omitempty"`
}
