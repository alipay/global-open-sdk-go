package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceUpdateRequest struct {
	InvoiceId        string               `json:"invoiceId,omitempty"`
	Description      string               `json:"description,omitempty"`
	DueDate          string               `json:"dueDate,omitempty"`
	CollectionMethod string               `json:"collectionMethod,omitempty"`
	PaymentMethod    *model.PaymentMethod `json:"paymentMethod,omitempty"`
	Shipping         *model.Shipping      `json:"shipping,omitempty"`
}

func NewAlipayInvoiceUpdateRequest() (*request.AlipayRequest, *AlipayInvoiceUpdateRequest) {
	alipayInvoiceUpdateRequest := &AlipayInvoiceUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceUpdateRequest, "/ams/api/v1/billing/invoice/update", &responseBilling.AlipayInvoiceUpdateResponse{})
	return alipayRequest, alipayInvoiceUpdateRequest
}

func (alipayInvoiceUpdateRequest *AlipayInvoiceUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceUpdateRequest, "/ams/api/v1/billing/invoice/update", &responseBilling.AlipayInvoiceUpdateResponse{})
}
