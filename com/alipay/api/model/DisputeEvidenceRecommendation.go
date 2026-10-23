package model

type DisputeEvidenceRecommendation struct {
	// Evidence tag.
	EvidenceTag string `json:"evidenceTag,omitempty"`
	// Whether Antom has already obtained this evidence.
	Obtained *bool `json:"obtained,omitempty"`
	// Optional contribution level: HIGH, MEDIUM or LOW.
	WinRateContribution string `json:"winRateContribution,omitempty"`
}
