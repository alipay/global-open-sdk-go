package model

type LineItem struct {
	CreditNoteItemId string            `json:"creditNoteItemId,omitempty"`
	Type             string            `json:"type,omitempty"`
	InvoiceItemId    string            `json:"invoiceItemId,omitempty"`
	Description      string            `json:"description,omitempty"`
	Quantity         string            `json:"quantity,omitempty"`
	UnitAmount       *Amount           `json:"unitAmount,omitempty"`
	ItemAmount       *Amount           `json:"itemAmount,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}
