package model

type TrialPhase struct {
	PhaseNo     int32   `json:"phaseNo,omitempty"`
	TrialAmount *Amount `json:"trialAmount,omitempty"`
}
