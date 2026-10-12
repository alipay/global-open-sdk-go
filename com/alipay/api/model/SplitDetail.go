package model

type SplitDetail struct {
	SplitTo           string  `json:"splitTo,omitempty"`
	SplitAmount       *Amount `json:"splitAmount,omitempty"`
	ActualSplitAmount *Amount `json:"actualSplitAmount,omitempty"`
	Description       string  `json:"description,omitempty"`
}
