package model

type PopRiskDecisionResultInfo struct {
	RiskAuthDecision string `json:"riskAuthDecision,omitempty"`
	RiskDecision     string `json:"riskDecision,omitempty"`
	PostRiskDecision string `json:"postRiskDecision,omitempty"`
}
