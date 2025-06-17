package model

type AmountLimit struct {
	MaxAmount    string `json:"maxAmount,omitempty"`
	MinAmount    string `json:"minAmount,omitempty"`
	RemainAmount string `json:"remainAmount,omitempty"`
}
