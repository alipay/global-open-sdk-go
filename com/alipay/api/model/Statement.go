package model

type Statement struct {
	StatementId    string          `json:"statementId,omitempty"`
	FundMoveDetail *FundMoveDetail `json:"fundMoveDetail,omitempty"`
}
