package model

type TaxCalculatedExemption struct {
	CertificateNumber string                              `json:"certificateNumber,omitempty"`
	ExemptionType     string                              `json:"exemptionType,omitempty"`
	Jurisdiction      *TaxCalculatedExemptionJurisdiction `json:"jurisdiction,omitempty"`
}
