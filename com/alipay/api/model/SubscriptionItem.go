package model

type SubscriptionItem struct {
	ItemId   string `json:"itemId,omitempty"`
	PriceId  string `json:"priceId,omitempty"`
	Quantity int32  `json:"quantity,omitempty"`
}
