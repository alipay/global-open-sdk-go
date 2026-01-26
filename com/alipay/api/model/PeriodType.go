package model

type PeriodType string

const (
	PeriodType_WEEK      PeriodType = "WEEK"
	PeriodType_MONTH     PeriodType = "MONTH"
	PeriodType_QUARTER   PeriodType = "QUARTER"
	PeriodType_HALF_YEAR PeriodType = "HALF_YEAR"
	PeriodType_YEAR      PeriodType = "YEAR"
)
