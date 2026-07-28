package model

type InvoiceShipping struct {
	ShippingName            *UserName         `json:"shippingName,omitempty"`
	ShippingAddress         *Address          `json:"shippingAddress,omitempty"`
	ShippingCarrier         string            `json:"shippingCarrier,omitempty"`
	ShippingNumber          string            `json:"shippingNumber,omitempty"`
	ShippingPhoneNo         string            `json:"shippingPhoneNo,omitempty"`
	ShipToEmail             string            `json:"shipToEmail,omitempty"`
	Notes                   string            `json:"notes,omitempty"`
	ShippingFeeId           string            `json:"shippingFeeId,omitempty"`
	ShippingFee             *Amount           `json:"shippingFee,omitempty"`
	ShippingDescription     string            `json:"shippingDescription,omitempty"`
	DeliveryEstimate        *DeliveryEstimate `json:"deliveryEstimate,omitempty"`
	TrackingUrl             string            `json:"trackingUrl,omitempty"`
	ShippingMethodIndicator string            `json:"shippingMethodIndicator,omitempty"`
}
