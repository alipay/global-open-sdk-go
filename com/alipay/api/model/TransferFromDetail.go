package model

import (

)




type TransferFromDetail struct {
        TransferFromMethod *PaymentMethod `json:"transferFromMethod,omitempty"` 
        TransferFromAmount *Amount `json:"transferFromAmount,omitempty"` 
}








