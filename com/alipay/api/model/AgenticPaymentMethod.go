package model

type AgenticPaymentMethod struct {
	PaymentMethodType string `json:"paymentMethodType,omitempty"`
	PaymentMethodId   string `json:"paymentMethodId,omitempty"`
	SelectedCardBrand string `json:"selectedCardBrand,omitempty"`
}
