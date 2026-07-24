package model

type Item struct {
	ItemId         string  `json:"itemId,omitempty"`
	Description    string  `json:"description,omitempty"`
	Quantity       int32   `json:"quantity,omitempty"`
	UnitAmount     *Amount `json:"unitAmount,omitempty"`
	Amount         *Amount `json:"amount,omitempty"`
	UsageAmount    *Amount `json:"usageAmount,omitempty"`
	UsageQuantity  string  `json:"usageQuantity,omitempty"`
	UsageUnit      string  `json:"usageUnit,omitempty"`
	DiscountAmount *Amount `json:"discountAmount,omitempty"`
	TaxAmount      *Amount `json:"taxAmount,omitempty"`
	PeriodStart    string  `json:"periodStart,omitempty"`
	PeriodEnd      string  `json:"periodEnd,omitempty"`
	Proration      bool    `json:"proration,omitempty"`
	GmtCreate      string  `json:"gmtCreate,omitempty"`
	GmtUpdate      string  `json:"gmtUpdate,omitempty"`
}
