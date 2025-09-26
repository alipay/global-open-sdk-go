package model

type SettlementInfo struct {
	SettlementCurrency    string                 `json:"settlementCurrency,omitempty"`
	SettlementBankAccount *SettlementBankAccount `json:"settlementBankAccount,omitempty"`
}
