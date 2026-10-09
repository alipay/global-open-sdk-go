package model

type InvoiceCustomerDetails struct {
	Email            string                   `json:"email,omitempty"`
	CustomerType     string                   `json:"customerType,omitempty"`
	BusinessName     string                   `json:"businessName,omitempty"`
	FirstName        string                   `json:"firstName,omitempty"`
	LastName         string                   `json:"lastName,omitempty"`
	BusinessAddress  *CustomerBusinessAddress `json:"businessAddress,omitempty"`
	PreferredLocales string                   `json:"preferredLocales,omitempty"`
	TaxIds           *[]*BuyerTaxId           `json:"taxIds,omitempty"`
}
