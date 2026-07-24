package model

type Product struct {
	ProductId     string            `json:"productId,omitempty"`
	Name          string            `json:"name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Description   string            `json:"description,omitempty"`
	Images        []string          `json:"images,omitempty"`
	UnitLabel     string            `json:"unitLabel,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	Active        bool              `json:"active,omitempty"`
	CreatedAt     string            `json:"createdAt,omitempty"`
	DeactivatedAt string            `json:"deactivatedAt,omitempty"`
	UpdatedAt     string            `json:"updatedAt,omitempty"`
}
