package model

type Limit struct {
	RemainingLimit *Amount `json:"remainingLimit,omitempty"`
	RangeLimit     *Amount `json:"rangeLimit,omitempty"`
	UsedLimit      *Amount `json:"usedLimit,omitempty"`
}
