package model

type RiskAddress struct {
	ShippingPhoneType             string `json:"shippingPhoneType,omitempty"`
	IsBillShipStateSame           bool   `json:"isBillShipStateSame,omitempty"`
	IsPreviousStateSame           bool   `json:"isPreviousStateSame,omitempty"`
	LocToShipDistance             int32  `json:"locToShipDistance,omitempty"`
	MinPreviousShipToBillDistance int32  `json:"minPreviousShipToBillDistance,omitempty"`
}
