package model

type RiskData struct {
	Order                  *RiskOrder              `json:"order,omitempty"`
	Buyer                  *RiskBuyer              `json:"buyer,omitempty"`
	Env                    *RiskEnv                `json:"env,omitempty"`
	RiskSignal             *RiskSignal             `json:"riskSignal,omitempty"`
	Address                *RiskAddress            `json:"address,omitempty"`
	CardVerificationResult *CardVerificationResult `json:"cardVerificationResult,omitempty"`
}
