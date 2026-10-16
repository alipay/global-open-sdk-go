package model

type AgenticBindingSessionMandate struct {
	MandateId string                              `json:"mandateId,omitempty"`
	Intent    *AgenticBindingSessionMandateIntent `json:"intent,omitempty"`
	Order     *AgenticBindingSessionMandateOrder  `json:"order,omitempty"`
}
