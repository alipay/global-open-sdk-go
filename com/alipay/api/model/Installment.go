package model

import (

)




type Installment struct {
        SupportCardBrands [] *SupportCardBrand `json:"supportCardBrands,omitempty"` 
        Plans [] *Plan `json:"plans,omitempty"` 
}








