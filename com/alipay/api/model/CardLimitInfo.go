package model

type CardLimitInfo struct {
	Currency               string `json:"currency,omitempty"`
	DailyLimitMax          string `json:"dailyLimitMax,omitempty"`
	MonthlyLimitMax        string `json:"monthlyLimitMax,omitempty"`
	PerTransactionLimitMax string `json:"perTransactionLimitMax,omitempty"`
	PerCardLimitMax        string `json:"perCardLimitMax,omitempty"`
}
