package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceConfirmPaymentResponse struct {
	response.AlipayResponse
	Result      *model.Result `json:"result,omitempty"`
	InvoiceId   string        `json:"invoiceId,omitempty"`
	Status      string        `json:"status,omitempty"`
	ReceiptId   string        `json:"receiptId,omitempty"`
	InvoiceNote string        `json:"invoiceNote,omitempty"`
	PaidAt      string        `json:"paidAt,omitempty"`
	SendStatus  string        `json:"sendStatus,omitempty"`
}
