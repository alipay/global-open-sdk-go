package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceMarkUncollectibleResponse struct {
	response.AlipayResponse
	Result                *model.Result `json:"result,omitempty"`
	InvoiceId             string        `json:"invoiceId,omitempty"`
	Status                string        `json:"status,omitempty"`
	MarkedUncollectibleAt string        `json:"markedUncollectibleAt,omitempty"`
	InvoiceNote           string        `json:"invoiceNote,omitempty"`
}
