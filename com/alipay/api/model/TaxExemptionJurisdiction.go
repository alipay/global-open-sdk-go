package model

type TaxExemptionJurisdiction struct {
	Country       string `json:"country,omitempty"`
	Region        string `json:"region,omitempty"`
	City          string `json:"city,omitempty"`
	EffectiveFrom string `json:"effectiveFrom,omitempty"`
}
