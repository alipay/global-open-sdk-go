package model

type CreateSubscriptionInfo struct {
	AllowRetry     bool    `json:"allowRetry,omitempty"`
	RetryMode      string  `json:"retryMode,omitempty"`
	MaxAmountFloor *Amount `json:"maxAmountFloor,omitempty"`
	FixedAmount    *Amount `json:"fixedAmount,omitempty"`
}
