package model

type TransferToDetail struct {
	TransferToMethod  *PaymentMethod `json:"transferToMethod,omitempty"`
	TransferToAmount  *Amount        `json:"transferToAmount,omitempty"`
	TransferNotifyUrl string         `json:"transferNotifyUrl,omitempty"`
	TransferRemark    string         `json:"transferRemark,omitempty"`
	TransferMemo      string         `json:"transferMemo,omitempty"`
}
