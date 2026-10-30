package model

type FundTransferRecipient struct {
	RecipientName          *FundTransferPartyName    `json:"recipientName,omitempty"`
	Address                *FundTransferPartyAddress `json:"address,omitempty"`
	RecipientAccountNumber string                    `json:"recipientAccountNumber,omitempty"`
}
