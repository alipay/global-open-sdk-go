package model

type TaxCustomerDetails struct {
	BusinessDetails *TaxBusinessDetails `json:"businessDetails,omitempty"`
	Name            string              `json:"name,omitempty"`
	ShippingAddress *TaxAddress         `json:"shippingAddress,omitempty"`
	BillingAddress  *TaxAddress         `json:"billingAddress,omitempty"`
	TaxIds          []*TaxId            `json:"taxIds,omitempty"`
	TaxExemptions   []*TaxExemption     `json:"taxExemptions,omitempty"`
}
