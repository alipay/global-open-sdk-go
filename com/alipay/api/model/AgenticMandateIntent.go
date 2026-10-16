package model

type AgenticMandateIntent struct {
	DeclineThreshold   *AgenticDeclineThreshold `json:"declineThreshold,omitempty"`
	EffectiveFromTime  string                   `json:"effectiveFromTime,omitempty"`
	EffectiveUntilTime string                   `json:"effectiveUntilTime,omitempty"`
	PreferredMerchant  *Merchant                `json:"preferredMerchant,omitempty"`
	AllowedMerchant    []*Merchant              `json:"allowedMerchant,omitempty"`
	ConsumerPrompt     string                   `json:"consumerPrompt,omitempty"`
	IntentDesc         string                   `json:"intentDesc,omitempty"`
}
