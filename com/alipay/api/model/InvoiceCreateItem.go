package model

type InvoiceCreateItem struct {
	Description string  `json:"description,omitempty"`
	ItemAmount  *Amount `json:"itemAmount,omitempty"`
	UnitAmount  *Amount `json:"unitAmount,omitempty"`
	PriceId     string  `json:"priceId,omitempty"`
	ProductId   string  `json:"productId,omitempty"`
	Quantity    int32   `json:"quantity,omitempty"`
	ItemId      string  `json:"itemId,omitempty"`
	SupplyStart string  `json:"supplyStart,omitempty"`
	SupplyEnd   string  `json:"supplyEnd,omitempty"`
}
