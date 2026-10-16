package model

type AgenticBoundPaymentMethodDetail struct {
	CardBrand string `json:"cardBrand,omitempty"`
	LastFour  string `json:"lastFour,omitempty"`
}
