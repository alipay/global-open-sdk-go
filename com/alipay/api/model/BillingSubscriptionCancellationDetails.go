package model

type BillingSubscriptionCancellationDetails struct {
	Feedback string `json:"feedback,omitempty"`
	Comment  string `json:"comment,omitempty"`
}
