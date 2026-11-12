package model

type ExcludedPaymentMethod struct {
	PaymentMethodTypeList []*PaymentMethodTypeItem `json:"paymentMethodTypeList,omitempty"`
}
