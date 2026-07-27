package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayReceiptSendResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	ReceiptId        string        `json:"receiptId,omitempty"`
	SendStatus       string        `json:"sendStatus,omitempty"`
	HostedReceiptUrl string        `json:"hostedReceiptUrl,omitempty"`
}
