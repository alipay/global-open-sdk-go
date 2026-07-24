package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceUpdateResponse struct {
	response.AlipayResponse
	Result    *model.Result `json:"result,omitempty"`
	InvoiceId string        `json:"invoiceId,omitempty"`
	Status    string        `json:"status,omitempty"`
}
