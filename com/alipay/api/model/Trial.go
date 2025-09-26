package model

type Trial struct {
	TrialStartPeriod int32   `json:"trialStartPeriod,omitempty"`
	TrialAmount      *Amount `json:"trialAmount,omitempty"`
	TrialEndPeriod   int32   `json:"trialEndPeriod,omitempty"`
}
