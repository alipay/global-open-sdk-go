package model

type Contact struct {
	Type   ContactType `json:"type,omitempty"`
	Info   string      `json:"info,omitempty"`
	Home   string      `json:"home,omitempty"`
	Work   string      `json:"work,omitempty"`
	Mobile string      `json:"mobile,omitempty"`
}
