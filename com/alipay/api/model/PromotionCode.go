package model

type PromotionCode struct {
	PromotionCodeRequestId string  `json:"promotionCodeRequestId,omitempty"`
	Code                   string  `json:"code,omitempty"`
	ExpiryTime             string  `json:"expiryTime,omitempty"`
	MinAmount              *Amount `json:"minAmount,omitempty"`
	OneTimeOnly            bool    `json:"oneTimeOnly,omitempty"`
	CustomerId             string  `json:"customerId,omitempty"`
	Metadata               string  `json:"metadata,omitempty"`
	MaxRedemptions         int32   `json:"maxRedemptions,omitempty"`
}
