package model

type AcquirerInfo struct {
	AcquirerName          string `json:"acquirerName,omitempty"`
	ReferenceRequestId    string `json:"referenceRequestId,omitempty"`
	AcquirerTransactionId string `json:"acquirerTransactionId,omitempty"`
	AcquirerMerchantId    string `json:"acquirerMerchantId,omitempty"`
	AcquirerMerchantName  string `json:"acquirerMerchantName,omitempty"`
	AcquirerResultCode    string `json:"acquirerResultCode,omitempty"`
	AcquirerResultMessage string `json:"acquirerResultMessage,omitempty"`
}
