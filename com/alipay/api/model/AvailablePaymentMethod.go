package model

import (

)




type AvailablePaymentMethod struct {
        PaymentMethodMetaData map[string]any `json:"paymentMethodMetaData,omitempty"`
        PaymentMethodTypeList [] *PaymentMethodTypeItem `json:"paymentMethodTypeList,omitempty"` 
}








