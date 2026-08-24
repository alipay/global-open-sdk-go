package model

type TaxCalculationLineItem struct {
	GoodsReferenceId string `json:"goodsReferenceId,omitempty"`
	Amount           string `json:"amount,omitempty"`
	Quantity         int32  `json:"quantity,omitempty"`
	TaxCode          string `json:"taxCode,omitempty"`
	ProductId        string `json:"productId,omitempty"`
	TaxBehavior      string `json:"taxBehavior,omitempty"`
}
