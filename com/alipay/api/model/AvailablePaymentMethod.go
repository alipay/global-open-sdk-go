package model

type AvailablePaymentMethod struct {
	PaymentMethodTypeList []*PaymentMethodTypeItem `json:"paymentMethodTypeList,omitempty"`
}
