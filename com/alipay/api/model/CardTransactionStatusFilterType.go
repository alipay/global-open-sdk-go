package model

type CardTransactionStatusFilterType string

const (
	CardTransactionStatusFilterType_SUCCESS    CardTransactionStatusFilterType = "SUCCESS"
	CardTransactionStatusFilterType_FAIL       CardTransactionStatusFilterType = "FAIL"
	CardTransactionStatusFilterType_PROCESSING CardTransactionStatusFilterType = "PROCESSING"
)
