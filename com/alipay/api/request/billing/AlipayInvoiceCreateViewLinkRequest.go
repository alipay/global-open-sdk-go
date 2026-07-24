package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceCreateViewLinkRequest struct {
	InvoiceId        string `json:"invoiceId,omitempty"`
	InvoiceRequestId string `json:"invoiceRequestId,omitempty"`
	LinkExpiryDays   int32  `json:"linkExpiryDays,omitempty"`
}

func NewAlipayInvoiceCreateViewLinkRequest() (*request.AlipayRequest, *AlipayInvoiceCreateViewLinkRequest) {
	alipayInvoiceCreateViewLinkRequest := &AlipayInvoiceCreateViewLinkRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceCreateViewLinkRequest, "/ams/api/v1/billing/invoice/createViewLink", &responseBilling.AlipayInvoiceCreateViewLinkResponse{})
	return alipayRequest, alipayInvoiceCreateViewLinkRequest
}

func (alipayInvoiceCreateViewLinkRequest *AlipayInvoiceCreateViewLinkRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceCreateViewLinkRequest, "/ams/api/v1/billing/invoice/createViewLink", &responseBilling.AlipayInvoiceCreateViewLinkResponse{})
}
