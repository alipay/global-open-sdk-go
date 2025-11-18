package model

type Statement struct {
	StatementId    string         `json:"statementId,omitempty"`
	FundMoveDetail FundMoveDetail `json:"fundMoveDetail,omitempty"`
}

type FundMoveDetail struct {
	Memo                   string `json:"memo,omitempty"`
	ReferenceTransactionId string `json:"referenceTransactionId,omitempty"`
}
