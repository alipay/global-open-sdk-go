package model

type BillingSubscriptionInquireDetailsPauseCollection struct {
	Behavior   string `json:"behavior,omitempty"`
	ResumesAt  string `json:"resumesAt,omitempty"`
	PausedAt   string `json:"pausedAt,omitempty"`
	ReasonCode string `json:"reasonCode,omitempty"`
}
