package model

type InquirePaymentMethodDetail struct {
	PaymentMethodDetailType string                    `json:"paymentMethodDetailType,omitempty"`
	Card                    *InquirePaymentMethodCard `json:"card,omitempty"`
}
