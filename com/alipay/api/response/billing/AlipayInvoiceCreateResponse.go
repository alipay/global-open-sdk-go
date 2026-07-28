package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceCreateResponse struct {
	response.AlipayResponse
	Result           *model.Result `json:"result,omitempty"`
	InvoiceId        string        `json:"invoiceId,omitempty"`
	InvoiceRequestId string        `json:"invoiceRequestId,omitempty"`
	Status           string        `json:"status,omitempty"`
	HostedInvoiceUrl string        `json:"hostedInvoiceUrl,omitempty"`
	SendStatus       string        `json:"sendStatus,omitempty"`
}
