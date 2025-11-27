package model

type RefundDetail struct {
	RefundAmount *Amount        `json:"refundAmount,omitempty"`
	RefundFrom   RefundFromType `json:"refundFrom,omitempty"`
}
