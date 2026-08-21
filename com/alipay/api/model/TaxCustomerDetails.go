package model

type TaxCustomerDetails struct {
	BusinessDetails *TaxBusinessDetails `json:"businessDetails,omitempty"`
	ShippingAddress *TaxAddress         `json:"shippingAddress,omitempty"`
	BillingAddress  *TaxAddress         `json:"billingAddress,omitempty"`
	TaxIds          []*TaxId            `json:"taxIds,omitempty"`
	TaxExemptions   []*TaxExemption     `json:"taxExemptions,omitempty"`
}
