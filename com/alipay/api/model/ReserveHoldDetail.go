package model

type ReserveHoldDetail struct {
	HoldRequestId    string  `json:"holdRequestId,omitempty"`
	HoldId           string  `json:"holdId,omitempty"`
	HoldAmount       *Amount `json:"holdAmount,omitempty"`
	ReleasableAmount *Amount `json:"releasableAmount,omitempty"`
	ReleaseTime      string  `json:"releaseTime,omitempty"`
	HoldStatus       string  `json:"holdStatus,omitempty"`
	ReleaseStatus    string  `json:"releaseStatus,omitempty"`
}
