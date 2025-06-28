package model

type RiskThreeDSResult struct {
	ThreeDSVersion         string `json:"threeDSVersion,omitempty"`
	ThreeDSInteractionMode string `json:"threeDSInteractionMode,omitempty"`
	Eci                    string `json:"eci,omitempty"`
	Cavv                   string `json:"cavv,omitempty"`
}
