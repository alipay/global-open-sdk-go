package model

type InvoiceNote struct {
	NoteId   string `json:"noteId,omitempty"`
	Note     string `json:"note,omitempty"`
	Action   string `json:"action,omitempty"`
	NoteTime string `json:"noteTime,omitempty"`
}
