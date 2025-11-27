package model

type SettlementDetail struct {
	SettleTo         SettleToType `json:"settleTo,omitempty"`
	SettlementAmount *Amount      `json:"settlementAmount,omitempty"`
}
