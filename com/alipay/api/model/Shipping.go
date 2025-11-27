package model

type Shipping struct {
	ShippingName        *UserName         `json:"shippingName,omitempty"`
	ShippingAddress     *Address          `json:"shippingAddress,omitempty"`
	ShippingCarrier     string            `json:"shippingCarrier,omitempty"`
	ShippingPhoneNo     string            `json:"shippingPhoneNo,omitempty"`
	ShipToEmail         string            `json:"shipToEmail,omitempty"`
	ShippingFeeId       string            `json:"shippingFeeId,omitempty"`
	ShippingFee         *Amount           `json:"shippingFee,omitempty"`
	ShippingDescription string            `json:"shippingDescription,omitempty"`
	DeliveryEstimate    *DeliveryEstimate `json:"deliveryEstimate,omitempty"`
	ShippingNumber      string            `json:"shippingNumber,omitempty"`
	Notes               string            `json:"notes,omitempty"`
}
