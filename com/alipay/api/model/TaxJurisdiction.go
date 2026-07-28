package model

type TaxJurisdiction struct {
	Country  string `json:"country,omitempty"`
	Region   string `json:"region,omitempty"`
	County   string `json:"county,omitempty"`
	City     string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
}
