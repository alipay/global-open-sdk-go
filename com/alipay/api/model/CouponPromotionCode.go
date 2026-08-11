package model

type CouponPromotionCode struct {
	PromotionCodeId string `json:"promotionCodeId,omitempty"`
	Code            string `json:"code,omitempty"`
	Status          string `json:"status,omitempty"`
}
