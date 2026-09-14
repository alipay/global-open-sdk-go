package model

type TaxCalculatedAddress struct {
	Country    string `json:"country,omitempty"`
	Region     string `json:"region,omitempty"`
	County     string `json:"county,omitempty"`
	City       string `json:"city,omitempty"`
	District   string `json:"district,omitempty"`
	Line1      string `json:"line1,omitempty"`
	Line2      string `json:"line2,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}
