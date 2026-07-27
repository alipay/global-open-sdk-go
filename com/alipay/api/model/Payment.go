package model

type Payment struct {
	InvoicePaymentId   string  `json:"invoicePaymentId,omitempty"`
	AttemptNo          int32   `json:"attemptNo,omitempty"`
	PaymentRequestId   string  `json:"paymentRequestId,omitempty"`
	PaymentId          string  `json:"paymentId,omitempty"`
	PayToRequestId     string  `json:"payToRequestId,omitempty"`
	PayToId            string  `json:"payToId,omitempty"`
	PaymentAmount      *Amount `json:"paymentAmount,omitempty"`
	PaymentOrderStatus string  `json:"paymentOrderStatus,omitempty"`
	PaymentMethod      string  `json:"paymentMethod,omitempty"`
	ErrorCode          string  `json:"errorCode,omitempty"`
	ErrorMessage       string  `json:"errorMessage,omitempty"`
	RetryReason        string  `json:"retryReason,omitempty"`
	PaymentTime        string  `json:"paymentTime,omitempty"`
	NextRetryAt        string  `json:"nextRetryAt,omitempty"`
	GmtCreate          string  `json:"gmtCreate,omitempty"`
	GmtUpdate          string  `json:"gmtUpdate,omitempty"`
}
