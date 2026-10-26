package model

type AmsInquiryPaymentResult struct {
	PaymentRequestId     string                     `json:"paymentRequestId,omitempty"`
	PaymentId            string                     `json:"paymentId,omitempty"`
	PaymentAmount        *Amount                    `json:"paymentAmount,omitempty"`
	PaymentStatus        string                     `json:"paymentStatus,omitempty"`
	PaymentResultCode    string                     `json:"paymentResultCode,omitempty"`
	PaymentResultMessage string                     `json:"paymentResultMessage,omitempty"`
	PaymentResultInfo    *PaymentResultInfo         `json:"paymentResultInfo,omitempty"`
	Transactions         []*BatchPaymentTransaction `json:"transactions,omitempty"`
}
