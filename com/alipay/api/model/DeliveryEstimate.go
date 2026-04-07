package model

import (

)




type DeliveryEstimate struct {
        Minimum *DeliveryEstimateInfo `json:"minimum,omitempty"` 
        Maximum *DeliveryEstimateInfo `json:"maximum,omitempty"` 
}








