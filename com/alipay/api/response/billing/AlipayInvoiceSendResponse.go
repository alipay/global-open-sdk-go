package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceSendResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	InvoiceId        string        `json:"invoiceId,omitempty"`
	SendStatus       string        `json:"sendStatus,omitempty"`
	HostedInvoiceUrl string        `json:"hostedInvoiceUrl,omitempty"`
}
