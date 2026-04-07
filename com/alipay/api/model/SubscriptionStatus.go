package model

type SubscriptionStatus string

const (
	SubscriptionStatus_ACTIVE     SubscriptionStatus = "ACTIVE"
	SubscriptionStatus_CANCELLED  SubscriptionStatus = "CANCELLED"
	SubscriptionStatus_TERMINATED SubscriptionStatus = "TERMINATED"
)
