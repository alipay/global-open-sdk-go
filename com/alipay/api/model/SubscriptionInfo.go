package model

type SubscriptionInfo struct {
	SubscriptionId          string             `json:"subscriptionId,omitempty"`
	Description             string             `json:"description,omitempty"`
	SubscriptionDescription string             `json:"subscriptionDescription,omitempty"`
	Status                  SubscriptionStatus `json:"status,omitempty"`
	SubscriptionStartTime   string             `json:"subscriptionStartTime,omitempty"`
	SubscriptionEndTime     string             `json:"subscriptionEndTime,omitempty"`
	PeriodType              PeriodType         `json:"periodType,omitempty"`
	PeriodCount             int32              `json:"periodCount,omitempty"`
	PeriodRule              *PeriodRule        `json:"periodRule,omitempty"`
	CurrentPeriodStart      string             `json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd        string             `json:"currentPeriodEnd,omitempty"`
	CurrentPhaseNo          int32              `json:"currentPhaseNo,omitempty"`
	PaymentMethod           *PaymentMethod     `json:"paymentMethod,omitempty"`
	PaymentAmount           *Amount            `json:"paymentAmount,omitempty"`
	LastUpdateTime          string             `json:"lastUpdateTime,omitempty"`
	RelatedSubscriptionId   string             `json:"relatedSubscriptionId,omitempty"`
	TrialPlan               *TrialPlan         `json:"trialPlan,omitempty"`
	Trials                  []*Trial           `json:"trials,omitempty"`
	SubscriptionNotifyUrl   string             `json:"subscriptionNotifyUrl,omitempty"`
	SubscriptionExpiryTime  string             `json:"subscriptionExpiryTime,omitempty"`
	AllowRetry              bool               `json:"allowRetry,omitempty"`
	MaxAmountFloor          *Amount            `json:"maxAmountFloor,omitempty"`
}
