package model

type DisputeEvidenceResult struct {
	EvidenceTag     string `json:"evidenceTag,omitempty"`
	Success         string `json:"success,omitempty"`
	ResultCode      string `json:"resultCode,omitempty"`
	FailedFieldPath string `json:"failedFieldPath,omitempty"`
}
