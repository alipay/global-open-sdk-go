package model

type AgenticAuthenticationSessionMandate struct {
	MandateId string                                     `json:"mandateId,omitempty"`
	Intent    *AgenticAuthenticationSessionMandateIntent `json:"intent,omitempty"`
	Order     *AgenticAuthenticationSessionMandateOrder  `json:"order,omitempty"`
}
