package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceCreateRequest struct {
	InvoiceRequestId string               `json:"invoiceRequestId,omitempty"`
	CustomerId       string               `json:"customerId,omitempty"`
	SubscriptionId   string               `json:"subscriptionId,omitempty"`
	Currency         string               `json:"currency,omitempty"`
	Items            []*model.Item        `json:"items,omitempty"`
	Status           string               `json:"status,omitempty"`
	AutoSend         bool                 `json:"autoSend,omitempty"`
	CcEmails         []string             `json:"ccEmails,omitempty"`
	Description      string               `json:"description,omitempty"`
	DueDate          string               `json:"dueDate,omitempty"`
	CollectionMethod string               `json:"collectionMethod,omitempty"`
	PaymentMethod    *model.PaymentMethod `json:"paymentMethod,omitempty"`
	Shipping         *model.Shipping      `json:"shipping,omitempty"`
	Discounts        []*model.Discount    `json:"discounts,omitempty"`
}

func NewAlipayInvoiceCreateRequest() (*request.AlipayRequest, *AlipayInvoiceCreateRequest) {
	alipayInvoiceCreateRequest := &AlipayInvoiceCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceCreateRequest, "/ams/api/v1/billing/invoice/create", &responseBilling.AlipayInvoiceCreateResponse{})
	return alipayRequest, alipayInvoiceCreateRequest
}

func (alipayInvoiceCreateRequest *AlipayInvoiceCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceCreateRequest, "/ams/api/v1/billing/invoice/create", &responseBilling.AlipayInvoiceCreateResponse{})
}
