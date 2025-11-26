package model

type RiskOrder struct {
	OrderType     string `json:"orderType,omitempty"`
	ReferringSite string `json:"referringSite,omitempty"`
}
