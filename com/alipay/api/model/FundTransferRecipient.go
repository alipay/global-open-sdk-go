package model

type FundTransferRecipient struct {
	RecipientName          *FundTransferRecipientName `json:"recipientName,omitempty"`
	Address                *FundTransferAddress       `json:"address,omitempty"`
	RecipientAccountNumber string                     `json:"recipientAccountNumber,omitempty"`
}
