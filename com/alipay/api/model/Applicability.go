package model

type Applicability struct {
	Scope    string `json:"scope,omitempty"`
	PriceIds string `json:"priceIds,omitempty"`
}
