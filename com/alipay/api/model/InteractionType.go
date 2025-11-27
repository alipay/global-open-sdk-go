package model

type InteractionType string

const (
	InteractionType_QR          InteractionType = "QR"
	InteractionType_REDIRECT    InteractionType = "REDIRECT"
	InteractionType_PUSH        InteractionType = "PUSH"
	InteractionType_ATM         InteractionType = "ATM"
	InteractionType_IBANKING    InteractionType = "IBANKING"
	InteractionType_BANKCOUNTER InteractionType = "BANKCOUNTER"
	InteractionType_OTC         InteractionType = "OTC"
)
