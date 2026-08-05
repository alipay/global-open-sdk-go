package model

type NotifyInvoiceInfo struct {
	InvoiceId                    string  `json:"invoiceId,omitempty"`
	InvoiceStatus                string  `json:"invoiceStatus,omitempty"`
	OriginalAmount               *Amount `json:"originalAmount,omitempty"`
	PrePaymentCreditNotesAmount  *Amount `json:"prePaymentCreditNotesAmount,omitempty"`
	PostPaymentCreditNotesAmount *Amount `json:"postPaymentCreditNotesAmount,omitempty"`
	AdjustedAmount               *Amount `json:"adjustedAmount,omitempty"`
}
