package model

type Subscription struct {
	SubscriptionId        string                          `json:"subscriptionId,omitempty"`
	SubscriptionRequestId string                          `json:"subscriptionRequestId,omitempty"`
	Status                string                          `json:"status,omitempty"`
	CustomerId            string                          `json:"customerId,omitempty"`
	Description           string                          `json:"description,omitempty"`
	BillingMode           string                          `json:"billingMode,omitempty"`
	CurrentPeriodStart    string                          `json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd      string                          `json:"currentPeriodEnd,omitempty"`
	CancelAtPeriodEnd     bool                            `json:"cancelAtPeriodEnd,omitempty"`
	CanceledAt            string                          `json:"canceledAt,omitempty"`
	CancelAt              string                          `json:"cancelAt,omitempty"`
	BillingCycleAnchor    string                          `json:"billingCycleAnchor,omitempty"`
	TrialStart            string                          `json:"trialStart,omitempty"`
	TrialEnd              string                          `json:"trialEnd,omitempty"`
	Subtotal              *Amount                         `json:"subtotal,omitempty"`
	DiscountAmount        *Amount                         `json:"discountAmount,omitempty"`
	TotalAmount           *Amount                         `json:"totalAmount,omitempty"`
	PriceItems            []*BillingSubscriptionPriceItem `json:"priceItems,omitempty"`
	TerminationReason     string                          `json:"terminationReason,omitempty"`
	CreateTime            string                          `json:"createTime,omitempty"`
}
