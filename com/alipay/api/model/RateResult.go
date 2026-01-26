package model

type RateResult struct {
	BuyCurrency  string `json:"buyCurrency,omitempty"`
	SellCurrency string `json:"sellCurrency,omitempty"`
	ExchangeRate string `json:"exchangeRate,omitempty"`
}
