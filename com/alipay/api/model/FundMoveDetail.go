package model

type FundMoveDetail struct {
	MoveType               string `json:"moveType,omitempty"`
	SourceAccount          string `json:"sourceAccount,omitempty"`
	TargetAccount          string `json:"targetAccount,omitempty"`
	MoveTime               string `json:"moveTime,omitempty"`
	Memo                   string `json:"memo,omitempty"`
	ReferenceTransactionId string `json:"referenceTransactionId,omitempty"`
	PayerAssetId           string `json:"payerAssetId,omitempty"`
	BeneficiaryAssetId     string `json:"beneficiaryAssetId,omitempty"`
}
