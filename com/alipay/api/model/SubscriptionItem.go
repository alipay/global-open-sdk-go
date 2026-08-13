package model

type SubscriptionItem struct {
	ItemId                 string `json:"itemId,omitempty"`
	PriceId                string `json:"priceId,omitempty"`
	Quantity               int32  `json:"quantity,omitempty"`
	CurrentPeriodStart     string `json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd       string `json:"currentPeriodEnd,omitempty"`
	RecurringIntervalCount int32  `json:"recurringIntervalCount,omitempty"`
}
