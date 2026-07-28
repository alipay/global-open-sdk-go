package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceConfirmPaymentRequest struct {
	InvoiceId        string `json:"invoiceId,omitempty"`
	ConfirmationType string `json:"confirmationType,omitempty"`
	PaymentMethod    string `json:"paymentMethod,omitempty"`
	Reference        string `json:"reference,omitempty"`
	AutoSend         bool   `json:"autoSend,omitempty"`
	InvoiceNote      string `json:"invoiceNote,omitempty"`
}

func NewAlipayInvoiceConfirmPaymentRequest() (*request.AlipayRequest, *AlipayInvoiceConfirmPaymentRequest) {
	alipayInvoiceConfirmPaymentRequest := &AlipayInvoiceConfirmPaymentRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceConfirmPaymentRequest, "/ams/api/v1/billing/invoice/confirmPayment", &responseBilling.AlipayInvoiceConfirmPaymentResponse{})
	return alipayRequest, alipayInvoiceConfirmPaymentRequest
}

func (alipayInvoiceConfirmPaymentRequest *AlipayInvoiceConfirmPaymentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceConfirmPaymentRequest, "/ams/api/v1/billing/invoice/confirmPayment", &responseBilling.AlipayInvoiceConfirmPaymentResponse{})
}
