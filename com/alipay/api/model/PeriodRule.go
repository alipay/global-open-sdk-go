package model

type PeriodRule struct {
	PeriodType  model.PeriodType `json:"periodType,omitempty"`
	Period      int32            `json:"period,omitempty"`
	Price       *Amount          `json:"price,omitempty"`
	PeriodCount int32            `json:"periodCount,omitempty"`
}
