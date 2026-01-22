package model

type FundMoveDetail struct {
	PayerName                 string `json:"payerName,omitempty"`
	PayerAccountNo            string `json:"payerAccountNo,omitempty"`
	PayerChaneelAccountnumber string `json:"payerChaneelAccountnumber,omitempty"`
	PayerAccountType          string `json:"payerAccountType,omitempty"`
	PayerAssetId              string `json:"payerAssetId,omitempty"`
	BeneficiaryName           string `json:"beneficiaryName,omitempty"`
	BeneficiaryAccountType    string `json:"beneficiaryAccountType,omitempty"`
	BeneficiaryBankCountry    string `json:"beneficiaryBankCountry,omitempty"`
	BeneficiaryBankName       string `json:"beneficiaryBankName,omitempty"`
	BeneficiaryAssetId        string `json:"beneficiaryAssetId,omitempty"`
	Remarks                   string `json:"remarks,omitempty"`
	Description               string `json:"description,omitempty"`
	Memo                      string `json:"memo,omitempty"`
	ReferenceTransactionId    string `json:"referenceTransactionId,omitempty"`
}
