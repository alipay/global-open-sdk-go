package model

import (

)




type ProrationSettings struct {
        ProrationMode string `json:"prorationMode,omitempty"`
        CustomAmount *Amount `json:"customAmount,omitempty"` 
}








