package model

type AccountBalance struct {
	Currency         string  `json:"currency,omitempty"`
	AvailableBalance *Amount `json:"availableBalance,omitempty"`
	FrozenBalance    *Amount `json:"frozenBalance,omitempty"`
	TotalBalance     *Amount `json:"totalBalance,omitempty"`
}
