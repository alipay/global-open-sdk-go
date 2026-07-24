package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceFinalizeRequest struct {
	InvoiceId   string `json:"invoiceId,omitempty"`
	AutoSend    bool   `json:"autoSend,omitempty"`
	InvoiceNote string `json:"invoiceNote,omitempty"`
}

func NewAlipayInvoiceFinalizeRequest() (*request.AlipayRequest, *AlipayInvoiceFinalizeRequest) {
	alipayInvoiceFinalizeRequest := &AlipayInvoiceFinalizeRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceFinalizeRequest, "/ams/api/v1/billing/invoice/finalize", &responseBilling.AlipayInvoiceFinalizeResponse{})
	return alipayRequest, alipayInvoiceFinalizeRequest
}

func (alipayInvoiceFinalizeRequest *AlipayInvoiceFinalizeRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceFinalizeRequest, "/ams/api/v1/billing/invoice/finalize", &responseBilling.AlipayInvoiceFinalizeResponse{})
}
