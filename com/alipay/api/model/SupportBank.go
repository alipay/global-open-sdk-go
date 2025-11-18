package model

type SupportBank struct {
	BankIdentifierCode string `json:"bankIdentifierCode,omitempty"`
	BankShortName      string `json:"bankShortName,omitempty"`
	BankLogo           *Logo  `json:"bankLogo,omitempty"`
}
