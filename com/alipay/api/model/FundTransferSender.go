package model

type FundTransferSender struct {
	SenderName *FundTransferPartyName    `json:"senderName,omitempty"`
	Address    *FundTransferPartyAddress `json:"address,omitempty"`
}
