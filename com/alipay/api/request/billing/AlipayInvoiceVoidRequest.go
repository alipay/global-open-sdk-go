package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceVoidRequest struct {
	VoidRequestId string `json:"voidRequestId,omitempty"`
	InvoiceId     string `json:"invoiceId,omitempty"`
	InvoiceNote   string `json:"invoiceNote,omitempty"`
}

func NewAlipayInvoiceVoidRequest() (*request.AlipayRequest, *AlipayInvoiceVoidRequest) {
	alipayInvoiceVoidRequest := &AlipayInvoiceVoidRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceVoidRequest, "/ams/api/v1/billing/invoice/void", &responseBilling.AlipayInvoiceVoidResponse{})
	return alipayRequest, alipayInvoiceVoidRequest
}

func (alipayInvoiceVoidRequest *AlipayInvoiceVoidRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceVoidRequest, "/ams/api/v1/billing/invoice/void", &responseBilling.AlipayInvoiceVoidResponse{})
}
