package model

type PaymentOptionDetail struct {
	SupportCardBrands []*SupportCardBrand `json:"supportCardBrands,omitempty"`
	Funding           []string            `json:"funding,omitempty"`
	SupportBanks      []*SupportBank      `json:"supportBanks,omitempty"`
	InteractionTypes  []string            `json:"interactionTypes,omitempty"`
}
