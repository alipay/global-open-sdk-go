package model

type RetryInfo struct {
	AvailableRetries int               `json:"availableRetries,omitempty"`
	PaymentAttempts  []*PaymentAttempt `json:"paymentAttempts,omitempty"`
}
