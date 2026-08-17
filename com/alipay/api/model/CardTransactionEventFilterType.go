package model

type CardTransactionEventFilterType string

const (
	CardTransactionEventFilterType_AUTH        CardTransactionEventFilterType = "AUTH"
	CardTransactionEventFilterType_AUTH_CANCEL CardTransactionEventFilterType = "AUTH_CANCEL"
	CardTransactionEventFilterType_CAPTURE     CardTransactionEventFilterType = "CAPTURE"
	CardTransactionEventFilterType_REFUND      CardTransactionEventFilterType = "REFUND"
	CardTransactionEventFilterType_CHARGEBACK  CardTransactionEventFilterType = "CHARGEBACK"
	CardTransactionEventFilterType_REPAYMENT   CardTransactionEventFilterType = "REPAYMENT"
)
