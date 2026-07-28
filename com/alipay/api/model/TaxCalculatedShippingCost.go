package model

type TaxCalculatedShippingCost struct {
	Amount       string          `json:"amount,omitempty"`
	TaxAmount    string          `json:"taxAmount,omitempty"`
	TaxBreakdown []*TaxBreakdown `json:"taxBreakdown,omitempty"`
}
