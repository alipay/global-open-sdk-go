package model

type BillingCustomerDetails struct {
	FullName        string                  `json:"fullName,omitempty"`
	CustomerType    string                  `json:"customerType,omitempty"`
	BusinessName    string                  `json:"businessName,omitempty"`
	BillingAddress  *BillingCustomerAddress `json:"billingAddress,omitempty"`
	ShippingAddress *BillingCustomerAddress `json:"shippingAddress,omitempty"`
	BusinessAddress *BillingCustomerAddress `json:"businessAddress,omitempty"`
	TaxIds          []*BillingCustomerTaxId `json:"taxIds,omitempty"`
}
