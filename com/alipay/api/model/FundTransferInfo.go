package model

type FundTransferInfo struct {
	FundTransferType  string                 `json:"fundTransferType,omitempty"`
	SellerMcc         string                 `json:"sellerMcc,omitempty"`
	SellerName        string                 `json:"sellerName,omitempty"`
	TransferSender    *FundTransferSender    `json:"transferSender,omitempty"`
	TransferRecipient *FundTransferRecipient `json:"transferRecipient,omitempty"`
}
