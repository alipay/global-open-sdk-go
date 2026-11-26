package model

type BillingCustomerDetails struct {
	FullName        string           `json:"fullName,omitempty"`
	CustomerType    string           `json:"customerType,omitempty"`
	BusinessName    string           `json:"businessName,omitempty"`
	BillingAddress  *CustomerAddress `json:"billingAddress,omitempty"`
	ShippingAddress *CustomerAddress `json:"shippingAddress,omitempty"`
	BusinessAddress *CustomerAddress `json:"businessAddress,omitempty"`
	TaxIds          []*CustomerTaxId `json:"taxIds,omitempty"`
}
