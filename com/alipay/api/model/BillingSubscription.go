package model

type BillingSubscription struct {
	CustomerId            string                         `json:"customerId,omitempty"`
	TrialSettings         *BillingTrialSettings          `json:"trialSettings,omitempty"`
	PaymentBehavior       string                         `json:"paymentBehavior,omitempty"`
	CollectionMethod      string                         `json:"collectionMethod,omitempty"`
	DaysUntilDue          int32                          `json:"daysUntilDue,omitempty"`
	BillingCycleAnchor    string                         `json:"billingCycleAnchor,omitempty"`
	CancelAt              string                         `json:"cancelAt,omitempty"`
	CancelAtPeriodEnd     bool                           `json:"cancelAtPeriodEnd,omitempty"`
	Description           string                         `json:"description,omitempty"`
	Discounts             []*BillingSubscriptionDiscount `json:"discounts,omitempty"`
	AllowPromotionCode    bool                           `json:"allowPromotionCode,omitempty"`
	SubscriptionNotifyUrl string                         `json:"subscriptionNotifyUrl,omitempty"`
}
