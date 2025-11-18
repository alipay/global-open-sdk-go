package model

type PaymentMethodTypeItem struct {
	PaymentMethodType  string `json:"paymentMethodType,omitempty"`
	PaymentMethodOrder int32  `json:"paymentMethodOrder,omitempty"`
	ExpressCheckout    bool   `json:"expressCheckout,omitempty"`
}
