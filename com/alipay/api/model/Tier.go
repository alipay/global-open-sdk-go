package model

type Tier struct {
	UpTo       int32   `json:"upTo,omitempty"`
	UnitAmount *Amount `json:"unitAmount,omitempty"`
	FlatAmount *Amount `json:"flatAmount,omitempty"`
}
