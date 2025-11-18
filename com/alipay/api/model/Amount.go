package model

type Amount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

func NewAmount(value string, currency string) *Amount {
	return &Amount{
		Value:    value,
		Currency: currency,
	}
}
