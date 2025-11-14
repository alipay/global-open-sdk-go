package model

type EntityAssociations struct {
	AssociationType   AssociationType `json:"associationType,omitempty"`
	LegalEntityType   LegalEntityType `json:"legalEntityType,omitempty"`
	Company           *Company        `json:"company,omitempty"`
	Individual        *Individual     `json:"individual,omitempty"`
	ShareholdingRatio string          `json:"shareholdingRatio,omitempty"`
}
