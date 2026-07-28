package model

type TaxBreakdown struct {
	TaxType          string `json:"taxType,omitempty"`
	AuthorityName    string `json:"authorityName,omitempty"`
	TaxRate          string `json:"taxRate,omitempty"`
	TaxAmount        string `json:"taxAmount,omitempty"`
	TaxableAmount    string `json:"taxableAmount,omitempty"`
	TaxabilityReason string `json:"taxabilityReason,omitempty"`
	Inclusive        bool   `json:"inclusive,omitempty"`
}
