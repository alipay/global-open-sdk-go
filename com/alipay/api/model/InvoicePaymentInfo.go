package model

type InvoicePaymentInfo struct {
	Result        Result  `json:"result,omitempty"`
	PaymentId     string  `json:"paymentId,omitempty"`
	PaymentAmount *Amount `json:"paymentAmount,omitempty"`
	PaymentTime   string  `json:"paymentTime,omitempty"`
}
