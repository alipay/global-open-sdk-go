package model

type AgenticMandate struct {
	MandateId string                `json:"mandateId,omitempty"`
	Intent    *AgenticMandateIntent `json:"intent,omitempty"`
	Order     *Order                `json:"order,omitempty"`
}
