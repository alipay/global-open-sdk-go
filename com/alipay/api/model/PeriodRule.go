package model

type PeriodRule struct {
	PeriodType  string `json:"periodType,omitempty"`
	PeriodCount int32  `json:"periodCount,omitempty"`
}
