package model

type TaxCalculatedCustomerDetails struct {
	BusinessDetails *TaxCalculatedBusinessDetails `json:"businessDetails,omitempty"`
	ShippingAddress *TaxCalculatedAddress         `json:"shippingAddress,omitempty"`
	BillingAddress  *TaxCalculatedAddress         `json:"billingAddress,omitempty"`
	TaxIds          []*TaxCalculatedTaxId         `json:"taxIds,omitempty"`
	TaxExemptions   []*TaxCalculatedExemption     `json:"taxExemptions,omitempty"`
}
