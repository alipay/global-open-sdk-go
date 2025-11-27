package model

type AvailablePaymentMethod struct {
	PaymentMethodTypeList []*PaymentMethodTypeItem          `json:"paymentMethodTypeList,omitempty"`
	PaymentMethodMetaData map[string]map[string]interface{} `json:"paymentMethodMetaData,omitempty"`
}
