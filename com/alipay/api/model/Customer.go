package model

type Customer struct {
	CustomerId          string `json:"customerId,omitempty"`
	CustomerRequestId   string `json:"customerRequestId,omitempty"`
	Email               string `json:"email,omitempty"`
	FirstName           string `json:"firstName,omitempty"`
	LastName            string `json:"lastName,omitempty"`
	Status              string `json:"status,omitempty"`
	PhoneNo             string `json:"phoneNo,omitempty"`
	CountryCode         string `json:"countryCode,omitempty"`
	BillingEmail        string `json:"billingEmail,omitempty"`
	ShippingFirstName   string `json:"shippingFirstName,omitempty"`
	ShippingLastName    string `json:"shippingLastName,omitempty"`
	ShippingCountryCode string `json:"shippingCountryCode,omitempty"`
}
