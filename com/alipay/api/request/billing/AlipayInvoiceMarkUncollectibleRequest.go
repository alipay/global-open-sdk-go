package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceMarkUncollectibleRequest struct {
	InvoiceId   string `json:"invoiceId,omitempty"`
	InvoiceNote string `json:"invoiceNote,omitempty"`
}

func NewAlipayInvoiceMarkUncollectibleRequest() (*request.AlipayRequest, *AlipayInvoiceMarkUncollectibleRequest) {
	alipayInvoiceMarkUncollectibleRequest := &AlipayInvoiceMarkUncollectibleRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceMarkUncollectibleRequest, "/ams/api/v1/billing/invoice/markUncollectible", &responseBilling.AlipayInvoiceMarkUncollectibleResponse{})
	return alipayRequest, alipayInvoiceMarkUncollectibleRequest
}

func (alipayInvoiceMarkUncollectibleRequest *AlipayInvoiceMarkUncollectibleRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceMarkUncollectibleRequest, "/ams/api/v1/billing/invoice/markUncollectible", &responseBilling.AlipayInvoiceMarkUncollectibleResponse{})
}
