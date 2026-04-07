package model

type SubscriptionTransaction struct {
	PaymentId     string         `json:"paymentId,omitempty"`
	Status        string         `json:"status,omitempty"`
	PhaseNo       int32          `json:"phaseNo,omitempty"`
	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentAmount *Amount        `json:"paymentAmount,omitempty"`
	PaymentTime   string         `json:"paymentTime,omitempty"`
	DisputeId     string         `json:"disputeId,omitempty"`
}
