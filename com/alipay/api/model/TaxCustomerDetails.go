package model

type TaxCustomerDetails struct {
	Name            string              `json:"name,omitempty"`
	BusinessDetails *TaxBusinessDetails `json:"businessDetails,omitempty"`
	ShippingAddress *TaxAddress         `json:"shippingAddress,omitempty"`
	BillingAddress  *TaxAddress         `json:"billingAddress,omitempty"`
	TaxIds          []*TaxId            `json:"taxIds,omitempty"`
}
