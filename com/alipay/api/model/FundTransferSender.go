package model

type FundTransferSender struct {
	SenderName *FundTransferName    `json:"senderName,omitempty"`
	Address    *FundTransferAddress `json:"address,omitempty"`
}
