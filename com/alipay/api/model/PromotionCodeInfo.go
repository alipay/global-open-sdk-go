package model

type PromotionCodeInfo struct {
	PromotionCodeId string `json:"promotionCodeId,omitempty"`
	Code            string `json:"code,omitempty"`
	Status          string `json:"status,omitempty"`
	MaxRedemptions  int32  `json:"maxRedemptions,omitempty"`
	RedeemedCount   int32  `json:"redeemedCount,omitempty"`
	ExpiryTime      string `json:"expiryTime,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
}
