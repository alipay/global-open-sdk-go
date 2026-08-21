package model

type TaxExemption struct {
	CertificateNumber string                    `json:"certificateNumber,omitempty"`
	ExemptionType     string                    `json:"exemptionType,omitempty"`
	Jurisdiction      *TaxExemptionJurisdiction `json:"jurisdiction,omitempty"`
}
