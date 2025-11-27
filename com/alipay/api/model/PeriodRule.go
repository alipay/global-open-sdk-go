package model

type PeriodRule struct {
	PeriodType  string  `json:"periodType,omitempty"`
	Period      int32   `json:"period,omitempty"`
	Price       *Amount `json:"price,omitempty"`
	PeriodCount int32   `json:"periodCount,omitempty"`
}
