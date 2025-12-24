package model

type Statement struct {
	StatementId        string          `json:"statementId,omitempty"`
	FundMoveDetail     *FundMoveDetail `json:"fundMoveDetail,omitempty"`
	TransactionType    string          `json:"transactionType,omitempty"`
	BeneficiaryAssetId string          `json:"beneficiaryAssetId,omitempty"`
}
