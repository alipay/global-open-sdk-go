package model

type PromotionInfo struct {
	PromotionType PromotionType `json:"promotionType,omitempty"`
	Discount      *Discount     `json:"discount,omitempty"`
	InterestFree  *InterestFree `json:"interestFree,omitempty"`
}
