package model

type TransactionType string

const (
	TransactionType_PAYMENT       TransactionType = "PAYMENT"
	TransactionType_REFUND        TransactionType = "REFUND"
	TransactionType_CAPTURE       TransactionType = "CAPTURE"
	TransactionType_CANCEL        TransactionType = "CANCEL"
	TransactionType_AUTHORIZATION TransactionType = "AUTHORIZATION"
	TransactionType_VOID          TransactionType = "VOID"
)
