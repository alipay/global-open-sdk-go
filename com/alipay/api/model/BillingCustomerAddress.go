package model

type BillingCustomerAddress struct {
	Country       string `json:"country,omitempty"`
	State         string `json:"state,omitempty"`
	City          string `json:"city,omitempty"`
	Address       string `json:"address,omitempty"`
	AddressDetail string `json:"addressDetail,omitempty"`
	Zipcode       string `json:"zipcode,omitempty"`
}
