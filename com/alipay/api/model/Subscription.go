package model

type Subscription struct {
	SubscriptionId   string `json:"subscriptionId,omitempty"`
	Status           string `json:"status,omitempty"`
	CustomerId       string `json:"customerId,omitempty"`
	CurrentPeriodEnd string `json:"currentPeriodEnd,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
}
