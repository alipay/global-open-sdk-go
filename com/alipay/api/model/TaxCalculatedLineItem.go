package model

type TaxCalculatedLineItem struct {
	GoodsReferenceId string          `json:"goodsReferenceId,omitempty"`
	Amount           *Amount         `json:"amount,omitempty"`
	Quantity         int32           `json:"quantity,omitempty"`
	TaxCode          string          `json:"taxCode,omitempty"`
	TaxBehavior      string          `json:"taxBehavior,omitempty"`
	TaxAmount        *Amount         `json:"taxAmount,omitempty"`
	TaxBreakdown     []*TaxBreakdown `json:"taxBreakdown,omitempty"`
}
