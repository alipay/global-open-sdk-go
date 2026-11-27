package model

type NotifyInvoiceInfo struct {
	InvoiceId     string `json:"invoiceId,omitempty"`
	InvoiceStatus string `json:"invoiceStatus,omitempty"`
	// Immutable original invoice total at finalization.
	OriginalAmount *Amount `json:"originalAmount,omitempty"`
	// Cumulative non-voided PRE_PAYMENT amount; reduces invoice total. Zero when none.
	PrePaymentCreditNotesAmount *Amount `json:"prePaymentCreditNotesAmount,omitempty"`
	// Cumulative non-voided POST_PAYMENT amount, including processing/failed refunds until void.
	// Does not change invoice total and is not the successfully-refunded amount.
	PostPaymentCreditNotesAmount *Amount `json:"postPaymentCreditNotesAmount,omitempty"`
	// Current totalAmount after PRE_PAYMENT reductions; unchanged by POST_PAYMENT refunds.
	// Zero may transition the invoice to PAID.
	AdjustedAmount *Amount `json:"adjustedAmount,omitempty"`
}
