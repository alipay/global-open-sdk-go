package model

type RefundFromType string

const (
	RefundFromType_SELLER          RefundFromType = "SELLER"
	RefundFromType_MARKETPLACE     RefundFromType = "MARKETPLACE"
	RefundFromType_UNSETTLED_FUNDS RefundFromType = "UNSETTLED_FUNDS"
)
