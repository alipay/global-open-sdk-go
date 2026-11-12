package model

type TradeBillPaymentAmount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}
