package model

type AgenticPaymentTokenOrderGoods struct {
	ReferenceGoodsId string         `json:"referenceGoodsId,omitempty"`
	GoodsName        string         `json:"goodsName,omitempty"`
	GoodsUnitAmount  *AgenticAmount `json:"goodsUnitAmount,omitempty"`
	GoodsQuantity    string         `json:"goodsQuantity,omitempty"`
	GoodsUrl         string         `json:"goodsUrl,omitempty"`
	GoodsCategory    string         `json:"goodsCategory,omitempty"`
}
