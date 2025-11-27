package model

type Company struct {
	LegalName         string          `json:"legalName,omitempty"`
	CompanyType       CompanyType     `json:"companyType,omitempty"`
	RegisteredAddress *Address        `json:"registeredAddress,omitempty"`
	OperatingAddress  *Address        `json:"operatingAddress,omitempty"`
	IncorporationDate string          `json:"incorporationDate,omitempty"`
	StockInfo         *StockInfo      `json:"stockInfo,omitempty"`
	Certificates      *Certificate    `json:"certificates,omitempty"`
	Attachments       []*Attachment   `json:"attachments,omitempty"`
	CompanyUnit       CompanyUnitType `json:"companyUnit,omitempty"`
	Contacts          []*Contact      `json:"contacts,omitempty"`
	VatNo             string          `json:"vatNo,omitempty"`
}
