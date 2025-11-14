package model

type TransitType string

const (
	TransitType_FLIGHT TransitType = "FLIGHT"
	TransitType_TRAIN  TransitType = "TRAIN"
	TransitType_CRUISE TransitType = "CRUISE"
	TransitType_COACH  TransitType = "COACH"
)
