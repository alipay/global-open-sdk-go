package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceReviseRequest struct {
	InvoiceId                string `json:"invoiceId,omitempty"`
	InvoiceRequestId         string `json:"invoiceRequestId,omitempty"`
	InvoiceRevisionRequestId string `json:"invoiceRevisionRequestId,omitempty"`
	Void                     bool   `json:"void,omitempty"`
}

func NewAlipayInvoiceReviseRequest() (*request.AlipayRequest, *AlipayInvoiceReviseRequest) {
	alipayInvoiceReviseRequest := &AlipayInvoiceReviseRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceReviseRequest, "/ams/api/v1/billing/invoice/revise", &responseBilling.AlipayInvoiceReviseResponse{})
	return alipayRequest, alipayInvoiceReviseRequest
}

func (alipayInvoiceReviseRequest *AlipayInvoiceReviseRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceReviseRequest, "/ams/api/v1/billing/invoice/revise", &responseBilling.AlipayInvoiceReviseResponse{})
}
