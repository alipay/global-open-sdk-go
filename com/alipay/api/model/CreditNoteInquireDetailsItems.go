package model

type CreditNoteInquireDetailsItems struct {
	Data    []*LineItem `json:"data,omitempty"`
	HasMore bool        `json:"hasMore,omitempty"`
}
