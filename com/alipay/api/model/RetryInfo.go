package model

type RetryInfo struct {
	AvailableRetries int32             `json:"availableRetries,omitempty"`
	OrderId          string            `json:"orderId,omitempty"`
	PaymentAttempts  []*PaymentAttempt `json:"paymentAttempts,omitempty"`
}
