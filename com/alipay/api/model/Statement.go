package model

type Statement struct {
	FundMoveDetail            *FundMoveDetail       `json:"fundMoveDetail,omitempty"`
	ForeignExchangeQuote      *ForeignExchangeQuote `json:"foreignExchangeQuote,omitempty"`
	StatementId               string                `json:"statementId,omitempty"`
	TransactionTime           string                `json:"transactionTime,omitempty"`
	TransactionType           string                `json:"transactionType,omitempty"`
	OriginalTransactionAmount *Amount               `json:"originalTransactionAmount,omitempty"`
	TransactionAmount         *Amount               `json:"transactionAmount,omitempty"`
	FeeAmount                 *Amount               `json:"feeAmount,omitempty"`
	NetAmount                 *Amount               `json:"netAmount,omitempty"`
	AccountBalance            *Amount               `json:"accountBalance,omitempty"`
	TransactionId             string                `json:"transactionId,omitempty"`
	ExtTransactionId          string                `json:"extTransactionId,omitempty"`
	TransactionStatus         string                `json:"transactionStatus,omitempty"`
	BeneficiaryAssetId        string                `json:"beneficiaryAssetId,omitempty"`
}
