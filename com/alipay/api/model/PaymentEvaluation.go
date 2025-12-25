package model

type PaymentEvaluation struct {
	PaymentMethods []*PaymentMethod `json:"paymentMethods,omitempty"`
}
