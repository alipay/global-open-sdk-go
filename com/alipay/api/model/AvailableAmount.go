package model

type AvailableAmount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}
