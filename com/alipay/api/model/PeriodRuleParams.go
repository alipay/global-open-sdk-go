package model

type PeriodRuleParams struct {
	PeriodType   string `json:"periodType,omitempty"`
	Period       string `json:"period,omitempty"`
	SingleAmount string `json:"singleAmount,omitempty"`
}
