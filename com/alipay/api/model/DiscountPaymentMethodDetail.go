package model

type DiscountPaymentMethodDetail struct {
	DiscountId                  string  `json:"discountId,omitempty"`
	AvailableAmount             *Amount `json:"availableAmount,omitempty"`
	DiscountName                string  `json:"discountName,omitempty"`
	DiscountDescription         string  `json:"discountDescription,omitempty"`
	PaymentMethodDetailMetadata string  `json:"paymentMethodDetailMetadata,omitempty"`
}
