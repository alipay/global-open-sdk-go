package model

type CustomerDetails struct {
	FullName        string           `json:"fullName,omitempty"`
	BillingAddress  *CustomerAddress `json:"billingAddress,omitempty"`
	ShippingAddress *CustomerAddress `json:"shippingAddress,omitempty"`
	CustomerType    string           `json:"customerType,omitempty"`
	BusinessName    string           `json:"businessName,omitempty"`
	BusinessAddress *CustomerAddress `json:"businessAddress,omitempty"`
	TaxIds          []*CustomerTaxId `json:"taxIds,omitempty"`
}
