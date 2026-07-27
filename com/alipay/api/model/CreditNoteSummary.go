package model

type CreditNoteSummary struct {
	CreditNoteId  string  `json:"creditNoteId,omitempty"`
	CustomerId    string  `json:"customerId,omitempty"`
	InvoiceId     string  `json:"invoiceId,omitempty"`
	Type          string  `json:"type,omitempty"`
	Status        string  `json:"status,omitempty"`
	TotalAmount   *Amount `json:"totalAmount,omitempty"`
	Reason        string  `json:"reason,omitempty"`
	EffectiveDate string  `json:"effectiveDate,omitempty"`
	RefundStatus  string  `json:"refundStatus,omitempty"`
	VoidedAt      string  `json:"voidedAt,omitempty"`
	RefundedAt    string  `json:"refundedAt,omitempty"`
	CreatedAt     string  `json:"createdAt,omitempty"`
}
