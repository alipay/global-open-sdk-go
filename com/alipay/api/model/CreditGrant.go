package model

type CreditGrant struct {
	CreditGrantId     string           `json:"creditGrantId,omitempty"`
	CreditGrantName   string           `json:"creditGrantName,omitempty"`
	CustomerId        string           `json:"customerId,omitempty"`
	Status            string           `json:"status,omitempty"`
	Amount            *Amount          `json:"amount,omitempty"`
	AvailableAmount   *AvailableAmount `json:"availableAmount,omitempty"`
	Applicability     *Applicability   `json:"applicability,omitempty"`
	Priority          int32            `json:"priority,omitempty"`
	Category          string           `json:"category,omitempty"`
	EffectiveDateTime string           `json:"effectiveDateTime,omitempty"`
	ExpiryDateTime    string           `json:"expiryDateTime,omitempty"`
	CreatedDateTime   string           `json:"createdDateTime,omitempty"`
	UpdatedDateTime   string           `json:"updatedDateTime,omitempty"`
}
