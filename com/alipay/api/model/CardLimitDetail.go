package model

type CardLimitDetail struct {
	PerTransactionLimit    *Amount `json:"perTransactionLimit,omitempty"`
	DailyLimit             *Limit  `json:"dailyLimit,omitempty"`
	MonthlyLimit           *Limit  `json:"monthlyLimit,omitempty"`
	PerCardLimit           *Limit  `json:"perCardLimit,omitempty"`
	DailyLimitMax          string  `json:"dailyLimitMax,omitempty"`
	MonthlyLimitMax        string  `json:"monthlyLimitMax,omitempty"`
	PerTransactionLimitMax string  `json:"perTransactionLimitMax,omitempty"`
	PerCardLimitMax        string  `json:"perCardLimitMax,omitempty"`
}
