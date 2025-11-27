package model

type AmountLimit struct {
	MaxAmount    *Amount `json:"maxAmount,omitempty"`
	MinAmount    *Amount `json:"minAmount,omitempty"`
	RemainAmount *Amount `json:"remainAmount,omitempty"`
}
