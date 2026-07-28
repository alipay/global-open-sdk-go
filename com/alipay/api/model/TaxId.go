package model

type TaxId struct {
	Value   string `json:"value,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
}
