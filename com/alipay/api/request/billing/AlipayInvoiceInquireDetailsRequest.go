package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceInquireDetailsRequest struct {
	InvoiceId        string `json:"invoiceId,omitempty"`
	InvoiceRequestId string `json:"invoiceRequestId,omitempty"`
}

func NewAlipayInvoiceInquireDetailsRequest() (*request.AlipayRequest, *AlipayInvoiceInquireDetailsRequest) {
	alipayInvoiceInquireDetailsRequest := &AlipayInvoiceInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceInquireDetailsRequest, "/ams/api/v1/billing/invoice/inquireDetails", &responseBilling.AlipayInvoiceInquireDetailsResponse{})
	return alipayRequest, alipayInvoiceInquireDetailsRequest
}

func (alipayInvoiceInquireDetailsRequest *AlipayInvoiceInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceInquireDetailsRequest, "/ams/api/v1/billing/invoice/inquireDetails", &responseBilling.AlipayInvoiceInquireDetailsResponse{})
}
