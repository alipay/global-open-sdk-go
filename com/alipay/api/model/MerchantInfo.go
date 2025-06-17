package model

type MerchantInfo struct {
	ReferenceMerchantId string                `json:"referenceMerchantId,omitempty"`
	LoginId             string                `json:"loginId,omitempty"`
	LegalEntityType     LegalEntityType       `json:"legalEntityType,omitempty"`
	Company             *Company              `json:"company,omitempty"`
	BusinessInfo        *BusinessInfo         `json:"businessInfo,omitempty"`
	EntityAssociations  []*EntityAssociations `json:"entityAssociations,omitempty"`
}
