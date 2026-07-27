package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayReceiptSendRequest struct {
	ReceiptId     string `json:"receiptId,omitempty"`
	SendRequestId string `json:"sendRequestId,omitempty"`
}

func NewAlipayReceiptSendRequest() (*request.AlipayRequest, *AlipayReceiptSendRequest) {
	alipayReceiptSendRequest := &AlipayReceiptSendRequest{}
	alipayRequest := request.NewAlipayRequest(alipayReceiptSendRequest, "/ams/api/v1/billing/receipt/send", &responseBilling.AlipayReceiptSendResponse{})
	return alipayRequest, alipayReceiptSendRequest
}

func (alipayReceiptSendRequest *AlipayReceiptSendRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayReceiptSendRequest, "/ams/api/v1/billing/receipt/send", &responseBilling.AlipayReceiptSendResponse{})
}
