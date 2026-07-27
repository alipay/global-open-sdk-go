package model

type InvoiceItem struct {
	ItemId         string  `json:"itemId,omitempty"`
	PriceId        string  `json:"priceId,omitempty"`
	Description    string  `json:"description,omitempty"`
	Quantity       string  `json:"quantity,omitempty"`
	UnitAmount     *Amount `json:"unitAmount,omitempty"`
	Amount         *Amount `json:"amount,omitempty"`
	UsageAmount    *Amount `json:"usageAmount,omitempty"`
	UsageQuantity  string  `json:"usageQuantity,omitempty"`
	PeriodStart    string  `json:"periodStart,omitempty"`
	PeriodEnd      string  `json:"periodEnd,omitempty"`
	Proration      bool    `json:"proration,omitempty"`
	DiscountAmount *Amount `json:"discountAmount,omitempty"`
	TaxAmount      *Amount `json:"taxAmount,omitempty"`
	GmtCreate      string  `json:"gmtCreate,omitempty"`
	GmtUpdate      string  `json:"gmtUpdate,omitempty"`
}
