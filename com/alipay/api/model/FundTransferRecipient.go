package model

type FundTransferRecipient struct {
	RecipientName          *FundTransferName    `json:"recipientName,omitempty"`
	Address                *FundTransferAddress `json:"address,omitempty"`
	RecipientAccountNumber string               `json:"recipientAccountNumber,omitempty"`
}
