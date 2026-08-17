package model

type EventMerchantInfo struct {
	Name   string `json:"name,omitempty"`
	Region string `json:"region,omitempty"`
	Mcc    string `json:"mcc,omitempty"`
}
