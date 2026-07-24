package model

type Price struct {
	PriceId          string             `json:"priceId,omitempty"`
	ProductId        string             `json:"productId,omitempty"`
	Name             string             `json:"name,omitempty"`
	PricingModel     string             `json:"pricingModel,omitempty"`
	UsageType        string             `json:"usageType,omitempty"`
	UnitLabel        string             `json:"unitLabel,omitempty"`
	MeterId          string             `json:"meterId,omitempty"`
	UnitAmount       *Amount            `json:"unitAmount,omitempty"`
	Recurring        *RecurringSettings `json:"recurring,omitempty"`
	Active           bool               `json:"active,omitempty"`
	IncludedQuantity int32              `json:"includedQuantity,omitempty"`
	TiersMode        string             `json:"tiersMode,omitempty"`
	Tiers            []*Tier            `json:"tiers,omitempty"`
	Metadata         map[string]string  `json:"metadata,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	DeactivatedAt    string             `json:"deactivatedAt,omitempty"`
	UpdatedAt        string             `json:"updatedAt,omitempty"`
}
