package model

type PriceItemChange struct {
	ChangeType     string `json:"changeType,omitempty"`
	CurrentPriceId string `json:"currentPriceId,omitempty"`
	NewPriceId     string `json:"newPriceId,omitempty"`
	Quantity       int32  `json:"quantity,omitempty"`
}
