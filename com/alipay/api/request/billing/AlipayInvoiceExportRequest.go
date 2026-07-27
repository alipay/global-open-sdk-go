package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceExportRequest struct {
	Limit          int32         `json:"limit,omitempty"`
	CustomerId     string        `json:"customerId,omitempty"`
	Status         string        `json:"status,omitempty"`
	SubscriptionId string        `json:"subscriptionId,omitempty"`
	InvoiceIds     []string      `json:"invoiceIds,omitempty"`
	Reason         string        `json:"reason,omitempty"`
	StartDate      string        `json:"startDate,omitempty"`
	EndDate        string        `json:"endDate,omitempty"`
	MinAmount      *model.Amount `json:"minAmount,omitempty"`
	MaxAmount      *model.Amount `json:"maxAmount,omitempty"`
}

func NewAlipayInvoiceExportRequest() (*request.AlipayRequest, *AlipayInvoiceExportRequest) {
	alipayInvoiceExportRequest := &AlipayInvoiceExportRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceExportRequest, "/ams/api/v1/billing/invoice/export", &responseBilling.AlipayInvoiceExportResponse{})
	return alipayRequest, alipayInvoiceExportRequest
}

func (alipayInvoiceExportRequest *AlipayInvoiceExportRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceExportRequest, "/ams/api/v1/billing/invoice/export", &responseBilling.AlipayInvoiceExportResponse{})
}
