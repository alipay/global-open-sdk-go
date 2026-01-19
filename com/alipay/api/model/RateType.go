package model

type RateType string

const (
	RateType_TRADE           RateType = "TRADE"
	RateType_REFERENCE       RateType = "REFERENCE"
	RateType_REFERENCE_TRADE RateType = "REFERENCE_TRADE"
)
