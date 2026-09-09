package model

type BuyerTaxId struct {
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	Value   string `json:"value,omitempty"`
}
