package model

type TaxCalculatedTaxId struct {
	Value   string `json:"value,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	Name    string `json:"name,omitempty"`
}
