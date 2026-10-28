package model

type FundTransferInfo struct {
	FundTransferType  string             `json:"fundTransferType,omitempty"`
	SellerMcc         string             `json:"sellerMcc,omitempty"`
	SellerName        string             `json:"sellerName,omitempty"`
	TransferSender    *TransferSender    `json:"transferSender,omitempty"`
	TransferRecipient *TransferRecipient `json:"transferRecipient,omitempty"`
}
