package model

type PriceItem struct {
	PriceId  string `json:"priceId,omitempty"`
	Quantity int32  `json:"quantity,omitempty"`
}
