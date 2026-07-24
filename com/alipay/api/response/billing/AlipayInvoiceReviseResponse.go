package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceReviseResponse struct {
	response.AlipayResponse
	Result          *model.Result `json:"result,omitempty"`
	NewInvoiceId    string        `json:"newInvoiceId,omitempty"`
	VoidedInvoiceId string        `json:"voidedInvoiceId,omitempty"`
}
