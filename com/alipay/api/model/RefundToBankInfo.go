package model

type RefundToBankInfo struct {
	BankCode          string    `json:"bankCode,omitempty"`
	AccountHolderName *UserName `json:"accountHolderName,omitempty"`
	AccountNo         string    `json:"accountNo,omitempty"`
}
