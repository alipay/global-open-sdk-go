package model

type PeriodRule struct {
	PeriodType  model.PeriodType `json:"periodType,omitempty"`
	PeriodCount int32            `json:"periodCount,omitempty"`
}
