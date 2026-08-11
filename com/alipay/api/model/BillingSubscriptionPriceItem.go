package model

type BillingSubscriptionPriceItem struct {
	CurrentPeriodEnd       string  `json:"currentPeriodEnd,omitempty"`
	CurrentPeriodStart     string  `json:"currentPeriodStart,omitempty"`
	ItemAmount             *Amount `json:"itemAmount,omitempty"`
	ItemId                 string  `json:"itemId,omitempty"`
	Nickname               string  `json:"nickname,omitempty"`
	PriceId                string  `json:"priceId,omitempty"`
	PriceType              string  `json:"priceType,omitempty"`
	PricingModel           string  `json:"pricingModel,omitempty"`
	ProductId              string  `json:"productId,omitempty"`
	ProductName            string  `json:"productName,omitempty"`
	Quantity               int32   `json:"quantity,omitempty"`
	RecurringInterval      string  `json:"recurringInterval,omitempty"`
	RecurringIntervalCount int32   `json:"recurringIntervalCount,omitempty"`
	UnitAmount             *Amount `json:"unitAmount,omitempty"`
	UsageType              string  `json:"usageType,omitempty"`
}
