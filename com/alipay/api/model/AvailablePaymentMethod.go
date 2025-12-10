package model

type AvailablePaymentMethod struct {
	PaymentMethodMetaData map[string]map[string]interface{} `json:"paymentMethodMetaData,omitempty"`
	PaymentMethodTypeList []*PaymentMethodTypeItem          `json:"paymentMethodTypeList,omitempty"`
}
