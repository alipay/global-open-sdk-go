package model

type BillingSubscriptionUpdateTrialSettings struct {
	TrialEnd    string                                             `json:"trialEnd,omitempty"`
	EndBehavior *BillingSubscriptionUpdateTrialSettingsEndBehavior `json:"endBehavior,omitempty"`
}
