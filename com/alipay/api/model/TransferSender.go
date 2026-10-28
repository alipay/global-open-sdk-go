package model

type TransferSender struct {
	SenderName *FundTransferName    `json:"senderName,omitempty"`
	Address    *FundTransferAddress `json:"address,omitempty"`
}
