package model

type TaxRegistration struct {
	TaxRegistrationId string           `json:"taxRegistrationId,omitempty"`
	TaxType           string           `json:"taxType,omitempty"`
	Jurisdiction      *TaxJurisdiction `json:"jurisdiction,omitempty"`
	RegistrationType  string           `json:"registrationType,omitempty"`
	TaxId             string           `json:"taxId,omitempty"`
	Status            string           `json:"status,omitempty"`
	ActiveFrom        string           `json:"activeFrom,omitempty"`
	ExpireAt          string           `json:"expireAt,omitempty"`
}
