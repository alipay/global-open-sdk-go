package model

type FundMoveDetail struct {
	Memo                   string `json:"memo,omitempty"`
	ReferenceTransactionId string `json:"referenceTransactionId,omitempty"`
	PayerAssetId           string `json:"payerAssetId,omitempty"`
	BeneficiaryAssetId     string `json:"beneficiaryAssetId,omitempty"`
}
