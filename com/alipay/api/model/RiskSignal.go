package model

type RiskSignal struct {
	RiskCode   string `json:"riskCode,omitempty"`
	RiskReason string `json:"riskReason,omitempty"`
}
