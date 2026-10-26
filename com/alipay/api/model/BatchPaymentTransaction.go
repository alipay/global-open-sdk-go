package model

type BatchPaymentTransaction struct {
	TransactionResult    *Result `json:"transactionResult,omitempty"`
	TransactionId        string  `json:"transactionId,omitempty"`
	TransactionType      string  `json:"transactionType,omitempty"`
	TransactionStatus    string  `json:"transactionStatus,omitempty"`
	TransactionAmount    *Amount `json:"transactionAmount,omitempty"`
	TransactionRequestId string  `json:"transactionRequestId,omitempty"`
}
