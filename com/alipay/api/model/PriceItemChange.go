package model

type PriceItemChange struct {
	Type           string `json:"type,omitempty"`
	ItemId         string `json:"itemId,omitempty"`
	CurrentPriceId string `json:"currentPriceId,omitempty"`
	NewPriceId     string `json:"newPriceId,omitempty"`
	NewQuantity    int32  `json:"newQuantity,omitempty"`
}
