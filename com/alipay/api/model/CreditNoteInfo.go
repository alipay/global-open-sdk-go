package model

type CreditNoteInfo struct {
	CreditNoteId      string      `json:"creditNoteId,omitempty"`
	Type              string      `json:"type,omitempty"`
	Status            string      `json:"status,omitempty"`
	TotalAmount       *Amount     `json:"totalAmount,omitempty"`
	RefundAmount      *Amount     `json:"refundAmount,omitempty"`
	RefundStatus      string      `json:"refundStatus,omitempty"`
	RefundId          string      `json:"refundId,omitempty"`
	RefundDestination string      `json:"refundDestination,omitempty"`
	Reason            string      `json:"reason,omitempty"`
	ReasonDescription string      `json:"reasonDescription,omitempty"`
	Memo              string      `json:"memo,omitempty"`
	EffectiveDate     string      `json:"effectiveDate,omitempty"`
	IssuedAt          string      `json:"issuedAt,omitempty"`
	RefundedAt        string      `json:"refundedAt,omitempty"`
	VoidedAt          string      `json:"voidedAt,omitempty"`
	CreatedAt         string      `json:"createdAt,omitempty"`
	Items             []*LineItem `json:"items,omitempty"`
}
