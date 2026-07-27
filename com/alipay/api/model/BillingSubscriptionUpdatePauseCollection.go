package model

type BillingSubscriptionUpdatePauseCollection struct {
	Behavior   string `json:"behavior,omitempty"`
	ResumesAt  string `json:"resumesAt,omitempty"`
	ReasonCode string `json:"reasonCode,omitempty"`
}
