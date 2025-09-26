package model

type Goods struct {
	ReferenceGoodsId   string  `json:"referenceGoodsId,omitempty"`
	GoodsName          string  `json:"goodsName,omitempty"`
	GoodsCategory      string  `json:"goodsCategory,omitempty"`
	GoodsBrand         string  `json:"goodsBrand,omitempty"`
	GoodsUnitAmount    *Amount `json:"goodsUnitAmount,omitempty"`
	GoodsQuantity      string  `json:"goodsQuantity,omitempty"`
	GoodsSkuName       string  `json:"goodsSkuName,omitempty"`
	GoodsUrl           string  `json:"goodsUrl,omitempty"`
	DeliveryMethodType string  `json:"deliveryMethodType,omitempty"`
	GoodsImageUrl      string  `json:"goodsImageUrl,omitempty"`
	PriceId            string  `json:"priceId,omitempty"`
}
