package model

// CreditNoteInfo is the notification projection, excluding detail-only merchant metadata.
type CreditNoteInfo struct {
	CreditNoteId string  `json:"creditNoteId,omitempty"`
	Type         string  `json:"type,omitempty"`
	Status       string  `json:"status,omitempty"`
	TotalAmount  *Amount `json:"totalAmount,omitempty"`
	// Optional planned original-method refund for refund-bearing POST_PAYMENT; not proof of completion.
	// Absent/null for note-only POST_PAYMENT and PRE_PAYMENT.
	RefundAmount *Amount `json:"refundAmount,omitempty"`
	// Optional REFUND_INITIATED, REFUND_SUCCEEDED or REFUND_FAILED progress.
	// Absent/null before progress or for note-only notes; PRE_PAYMENT may use NOT_APPLICABLE.
	RefundStatus string `json:"refundStatus,omitempty"`
	// Optional payment-system reference when obtained; alone does not prove success.
	// Absent/null beforehand, for note-only POST_PAYMENT and PRE_PAYMENT.
	RefundId string `json:"refundId,omitempty"`
	// Optional REFUND routing supplied at creation; null when omitted.
	RefundDestination string `json:"refundDestination,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ReasonDescription string `json:"reasonDescription,omitempty"`
	Memo              string `json:"memo,omitempty"`
	EffectiveDate     string `json:"effectiveDate,omitempty"`
	IssuedAt          string `json:"issuedAt,omitempty"`
	RefundedAt        string `json:"refundedAt,omitempty"`
	VoidedAt          string `json:"voidedAt,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	// Optional, at most 100 items; the service returns [] when absent.
	Items []*CreditNoteNotificationLineItem `json:"items,omitempty"`
}

// CreditNoteNotificationLineItem excludes the shared LineItem's detail-only metadata.
type CreditNoteNotificationLineItem struct {
	CreditNoteItemId string `json:"creditNoteItemId,omitempty"`
	Type             string `json:"type,omitempty"`
	InvoiceItemId    string `json:"invoiceItemId,omitempty"`
	Description      string `json:"description,omitempty"`
	// Quantity mode only, from 1 to 999.
	Quantity *int32 `json:"quantity,omitempty"`
	// Required C (credit) or D (debit); D itemAmount is subtracted from the credit total.
	Direction  string  `json:"direction,omitempty"`
	UnitAmount *Amount `json:"unitAmount,omitempty"`
	ItemAmount *Amount `json:"itemAmount,omitempty"`
}
