package model

type CreditNoteCreateItems struct {
	Data    []*LineItem `json:"data,omitempty"`
	HasMore bool        `json:"hasMore,omitempty"`
}
