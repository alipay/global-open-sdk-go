package model

type TaxCalculationLineItem struct {
	GoodsReferenceId string `json:"goodsReferenceId,omitempty"`
	UnitAmount       string `json:"unitAmount,omitempty"`
	Quantity         int32  `json:"quantity,omitempty"`
	TaxCode          string `json:"taxCode,omitempty"`
	ProductId        string `json:"productId,omitempty"`
	TaxBehavior      string `json:"taxBehavior,omitempty"`
}
