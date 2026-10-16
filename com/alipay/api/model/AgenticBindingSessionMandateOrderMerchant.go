package model

type AgenticBindingSessionMandateOrderMerchant struct {
	ReferenceMerchantId string `json:"referenceMerchantId,omitempty"`
	MerchantName        string `json:"merchantName,omitempty"`
	MerchantMCC         string `json:"merchantMCC,omitempty"`
	MerchantUrl         string `json:"merchantUrl,omitempty"`
	MerchantCountryCode string `json:"merchantCountryCode,omitempty"`
}
