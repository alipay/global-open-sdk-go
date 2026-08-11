package model

type Product struct {
	ProductId        string   `json:"productId,omitempty"`
	ProductRequestId string   `json:"productRequestId,omitempty"`
	Name             string   `json:"name,omitempty"`
	Type             string   `json:"type,omitempty"`
	Description      string   `json:"description,omitempty"`
	Images           []string `json:"images,omitempty"`
	UnitLabel        string   `json:"unitLabel,omitempty"`
	Metadata         string   `json:"metadata,omitempty"`
	Active           bool     `json:"active,omitempty"`
	CreatedAt        string   `json:"createdAt,omitempty"`
	DeactivatedAt    string   `json:"deactivatedAt,omitempty"`
	UpdatedAt        string   `json:"updatedAt,omitempty"`
}
