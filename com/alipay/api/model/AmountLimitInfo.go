package model

type AmountLimitInfo struct {
	SingleLimit *AmountLimit `json:"singleLimit,omitempty"`
	DayLimit    *AmountLimit `json:"dayLimit,omitempty"`
	MonthLimit  *AmountLimit `json:"monthLimit,omitempty"`
}
