package model

type ThreeDSResult struct {
	ThreeDSVersion  string `json:"threeDSVersion,omitempty"`
	Eci             string `json:"eci,omitempty"`
	Cavv            string `json:"cavv,omitempty"`
	DsTransactionId string `json:"dsTransactionId,omitempty"`
	Xid             string `json:"xid,omitempty"`
}
