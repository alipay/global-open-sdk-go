package model

type AgenticAmount struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}
