package model

type SubscriptionPlan struct {
	AllowAccumulate             bool        `json:"allowAccumulate,omitempty"`
	MaxAccumulateAmount         *Amount     `json:"maxAccumulateAmount,omitempty"`
	PeriodRule                  *PeriodRule `json:"periodRule,omitempty"`
	SubscriptionStartTime       string      `json:"subscriptionStartTime,omitempty"`
	SubscriptionNotificationUrl string      `json:"subscriptionNotificationUrl,omitempty"`
}
