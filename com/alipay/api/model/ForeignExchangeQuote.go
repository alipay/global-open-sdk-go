package model

type ForeignExchangeQuote struct {
	ExchangeRate   string `json:"exchangeRate,omitempty"`
	SourceCurrency string `json:"sourceCurrency,omitempty"`
	TargetCurrency string `json:"targetCurrency,omitempty"`
	QuoteTime      string `json:"quoteTime,omitempty"`
	QuotePrice     string `json:"quotePrice,omitempty"`
}
