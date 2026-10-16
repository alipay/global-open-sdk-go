package model

type AgenticAuthenticationSessionAgent struct {
	AgentId   string `json:"agentId,omitempty"`
	AgentName string `json:"agentName,omitempty"`
	AgentUrl  string `json:"agentUrl,omitempty"`
	AgentJwk  string `json:"agentJwk,omitempty"`
}
