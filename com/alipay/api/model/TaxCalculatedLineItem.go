package model

type TaxCalculatedLineItem struct {
	GoodsReferenceId string          `json:"goodsReferenceId,omitempty"`
	Amount           string          `json:"amount,omitempty"`
	Quantity         int32           `json:"quantity,omitempty"`
	TaxCode          string          `json:"taxCode,omitempty"`
	TaxBehavior      string          `json:"taxBehavior,omitempty"`
	TaxAmount        string          `json:"taxAmount,omitempty"`
	TaxBreakdown     []*TaxBreakdown `json:"taxBreakdown,omitempty"`
}
