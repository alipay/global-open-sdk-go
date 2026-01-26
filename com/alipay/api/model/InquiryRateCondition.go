package model

type InquiryRateCondition struct {
	BuyCurrency  string `json:"buyCurrency,omitempty"`
	SellCurrency string `json:"sellCurrency,omitempty"`
}
