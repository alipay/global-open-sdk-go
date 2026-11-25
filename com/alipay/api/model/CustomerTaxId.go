package model

type CustomerTaxId struct {
	Value   string `json:"value,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	Type    string `json:"type,omitempty"`
}
