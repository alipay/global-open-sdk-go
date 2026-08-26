package model

type TaxCalculatedShippingCost struct {
	Amount       *Amount         `json:"amount,omitempty"`
	TaxAmount    *Amount         `json:"taxAmount,omitempty"`
	TaxBreakdown []*TaxBreakdown `json:"taxBreakdown,omitempty"`
}
