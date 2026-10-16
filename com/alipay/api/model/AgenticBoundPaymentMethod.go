package model

type AgenticBoundPaymentMethod struct {
	PaymentMethodType   string                           `json:"paymentMethodType,omitempty"`
	PaymentMethodId     string                           `json:"paymentMethodId,omitempty"`
	PaymentMethodDetail *AgenticBoundPaymentMethodDetail `json:"paymentMethodDetail,omitempty"`
}
