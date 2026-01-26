package model

type SubscriptionInfo struct {
	SubscriptionDescription string      `json:"subscriptionDescription,omitempty"`
	SubscriptionStartTime   string      `json:"subscriptionStartTime,omitempty"`
	SubscriptionEndTime     string      `json:"subscriptionEndTime,omitempty"`
	PeriodRule              *PeriodRule `json:"periodRule,omitempty"`
	Trials                  []*Trial    `json:"trials,omitempty"`
	SubscriptionNotifyUrl   string      `json:"subscriptionNotifyUrl,omitempty"`
	SubscriptionExpiryTime  string      `json:"subscriptionExpiryTime,omitempty"`
	AllowRetry              bool        `json:"allowRetry,omitempty"`
	MaxAmountFloor          *Amount     `json:"maxAmountFloor,omitempty"`
}
