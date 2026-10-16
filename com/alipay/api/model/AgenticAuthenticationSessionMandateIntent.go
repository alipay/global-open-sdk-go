package model

type AgenticAuthenticationSessionMandateIntent struct {
	DeclineThreshold   *AgenticDeclineThreshold  `json:"declineThreshold,omitempty"`
	EffectiveFromTime  string                    `json:"effectiveFromTime,omitempty"`
	EffectiveUntilTime string                    `json:"effectiveUntilTime,omitempty"`
	PreferredMerchant  *AgenticPreferredMerchant `json:"preferredMerchant,omitempty"`
	AllowedMerchant    []*AgenticAllowedMerchant `json:"allowedMerchant,omitempty"`
	ConsumerPrompt     string                    `json:"consumerPrompt,omitempty"`
	IntentDesc         string                    `json:"intentDesc,omitempty"`
}
