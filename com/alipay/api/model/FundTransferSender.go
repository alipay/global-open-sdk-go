package model

type FundTransferSender struct {
	SenderName *FundTransferSenderName `json:"senderName,omitempty"`
	Address    *FundTransferAddress    `json:"address,omitempty"`
}
