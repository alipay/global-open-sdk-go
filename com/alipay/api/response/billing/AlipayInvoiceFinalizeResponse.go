package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceFinalizeResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	InvoiceId        string        `json:"invoiceId,omitempty"`
	Status           string        `json:"status,omitempty"`
	HostedInvoiceUrl string        `json:"hostedInvoiceUrl,omitempty"`
	FinalizedAt      string        `json:"finalizedAt,omitempty"`
	InvoiceNote      string        `json:"invoiceNote,omitempty"`
	SendStatus       string        `json:"sendStatus,omitempty"`
}
