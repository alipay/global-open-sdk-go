package model

type PaymentMethodTypeItem struct {
	PaymentMethodType  string `json:"paymentMethodType,omitempty"`
	PaymentMethodOrder string `json:"paymentMethodOrder,omitempty"`
	ExpressCheckout    string `json:"expressCheckout,omitempty"`
}
