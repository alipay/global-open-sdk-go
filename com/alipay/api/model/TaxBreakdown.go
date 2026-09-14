package model

type TaxBreakdown struct {
	TaxType          string  `json:"taxType,omitempty"`
	TaxRate          string  `json:"taxRate,omitempty"`
	TaxAmount        *Amount `json:"taxAmount,omitempty"`
	TaxableAmount    *Amount `json:"taxableAmount,omitempty"`
	TaxabilityReason string  `json:"taxabilityReason,omitempty"`
	Inclusive        bool    `json:"inclusive,omitempty"`
}
