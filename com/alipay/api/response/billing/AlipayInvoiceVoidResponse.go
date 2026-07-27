package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceVoidResponse struct {
	response.AlipayResponse
	Result        *model.Result `json:"result,omitempty"`
	VoidRequestId string        `json:"voidRequestId,omitempty"`
	InvoiceId     string        `json:"invoiceId,omitempty"`
	Status        string        `json:"status,omitempty"`
	VoidedAt      string        `json:"voidedAt,omitempty"`
	InvoiceNote   string        `json:"invoiceNote,omitempty"`
}
