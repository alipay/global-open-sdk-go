package model

type Transaction struct {
	TransactionResult    *Result               `json:"transactionResult,omitempty"`
	TransactionId        string                `json:"transactionId,omitempty"`
	TransactionType      TransactionType       `json:"transactionType,omitempty"`
	TransactionStatus    TransactionStatusType `json:"transactionStatus,omitempty"`
	TransactionAmount    *Amount               `json:"transactionAmount,omitempty"`
	TransactionRequestId string                `json:"transactionRequestId,omitempty"`
	TransactionTime      string                `json:"transactionTime,omitempty"`
	AcquirerInfo         *AcquirerInfo         `json:"acquirerInfo,omitempty"`
}
