package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayReceiptInquireDetailsRequest struct {
	ReceiptId string `json:"receiptId,omitempty"`
	InvoiceId string `json:"invoiceId,omitempty"`
}

func NewAlipayReceiptInquireDetailsRequest() (*request.AlipayRequest, *AlipayReceiptInquireDetailsRequest) {
	alipayReceiptInquireDetailsRequest := &AlipayReceiptInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayReceiptInquireDetailsRequest, "/ams/api/v1/billing/receipt/inquireDetails", &responseBilling.AlipayReceiptInquireDetailsResponse{})
	return alipayRequest, alipayReceiptInquireDetailsRequest
}

func (alipayReceiptInquireDetailsRequest *AlipayReceiptInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayReceiptInquireDetailsRequest, "/ams/api/v1/billing/receipt/inquireDetails", &responseBilling.AlipayReceiptInquireDetailsResponse{})
}
