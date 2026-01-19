package model

type PaymentQuote struct {
	BuyCurrency  string  `json:"buyCurrency,omitempty"`
	SellCurrency string  `json:"sellCurrency,omitempty"`
	QuoteId      string  `json:"quoteId,omitempty"`
	ExchangeRate float32 `json:"exchangeRate,omitempty"`
}
