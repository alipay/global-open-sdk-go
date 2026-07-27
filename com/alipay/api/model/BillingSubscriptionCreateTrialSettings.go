package model

type BillingSubscriptionCreateTrialSettings struct {
	TrialPeriodDays int32  `json:"trialPeriodDays,omitempty"`
	TrialEnd        string `json:"trialEnd,omitempty"`
}
