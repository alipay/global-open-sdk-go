package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceSendRequest struct {
	InvoiceId          string   `json:"invoiceId,omitempty"`
	CcEmails           []string `json:"ccEmails,omitempty"`
	InternalNote       string   `json:"internalNote,omitempty"`
	IncludePaymentLink bool     `json:"includePaymentLink,omitempty"`
}

func NewAlipayInvoiceSendRequest() (*request.AlipayRequest, *AlipayInvoiceSendRequest) {
	alipayInvoiceSendRequest := &AlipayInvoiceSendRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceSendRequest, "/ams/api/v1/billing/invoice/send", &responseBilling.AlipayInvoiceSendResponse{})
	return alipayRequest, alipayInvoiceSendRequest
}

func (alipayInvoiceSendRequest *AlipayInvoiceSendRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceSendRequest, "/ams/api/v1/billing/invoice/send", &responseBilling.AlipayInvoiceSendResponse{})
}
