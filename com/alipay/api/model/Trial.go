package model

type Trial struct {
	TrialPeriod      int32   `json:"trialPeriod,omitempty"`
	TrialAmount      *Amount `json:"trialAmount,omitempty"`
	TrialStartPeriod int32   `json:"trialStartPeriod,omitempty"`
	TrialEndPeriod   int32   `json:"trialEndPeriod,omitempty"`
}
