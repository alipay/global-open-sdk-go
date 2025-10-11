package model

type CurrencyPair struct {
	SellCurrency string `json:"sellCurrency,omitempty"`
	BuyCurrency  string `json:"buyCurrency,omitempty"`
}
