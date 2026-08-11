package model

type CreditNoteItem struct {
	CreditNoteItemId string  `json:"creditNoteItemId,omitempty"`
	Type             string  `json:"type,omitempty"`
	InvoiceItemId    string  `json:"invoiceItemId,omitempty"`
	Description      string  `json:"description,omitempty"`
	Quantity         int32   `json:"quantity,omitempty"`
	UnitAmount       *Amount `json:"unitAmount,omitempty"`
	ItemAmount       *Amount `json:"itemAmount,omitempty"`
}
