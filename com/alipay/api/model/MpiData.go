package model

type MpiData struct {
	ThreeDSVersion  string `json:"threeDSVersion,omitempty"`
	Eci             string `json:"eci,omitempty"`
	Cavv            string `json:"cavv,omitempty"`
	DsTransactionId string `json:"dsTransactionId,omitempty"`
	CredentialType  string `json:"credentialType,omitempty"`
}
