package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

// NotifyABATransferNotify is the notification for ABA transfer
type AlipayABATransferNotify struct {
	AlipayNotify
	TransferRequestId  string                    `json:"transferRequestId,omitempty"`
	TransferId         string                    `json:"transferId,omitempty"`
	TransferResult     *model.Result             `json:"transferResult,omitempty"`
	TransferFinishTime string                    `json:"transferFinishTime,omitempty"`
	TransferFromDetail *TransferFromDetailNotify `json:"transferFromDetail,omitempty"`
	TransferToDetail   *TransferToDetailNotify   `json:"transferToDetail,omitempty"`
}

// TransferFromDetailNotify contains the transfer from detail information
type TransferFromDetailNotify struct {
	TransferFromAmount *model.Amount `json:"transferFromAmount,omitempty"`
}

// TransferToDetailNotify contains the transfer to detail information
type TransferToDetailNotify struct {
	TransferToMethod  *model.PaymentMethod `json:"transferToMethod,omitempty"`
	TransferToAmount  *model.Amount        `json:"transferToAmount,omitempty"`
	FeeAmount         *model.Amount        `json:"feeAmount,omitempty"`
	TransferNotifyUrl string               `json:"transferNotifyUrl,omitempty"`
}
