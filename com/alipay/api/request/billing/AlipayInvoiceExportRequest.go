package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceExportRequest struct {
	CustomerId     string   `json:"customerId,omitempty"`
	Status         string   `json:"status,omitempty"`
	SubscriptionId string   `json:"subscriptionId,omitempty"`
	InvoiceIds     []string `json:"invoiceIds,omitempty"`
	StartDate      string   `json:"startDate,omitempty"`
	EndDate        string   `json:"endDate,omitempty"`
	FileFormat     string   `json:"fileFormat,omitempty"`
	Language       string   `json:"language,omitempty"`
	DownloadType   string   `json:"downloadType,omitempty"`
	ColumnPreset   string   `json:"columnPreset,omitempty"`
}

func NewAlipayInvoiceExportRequest() (*request.AlipayRequest, *AlipayInvoiceExportRequest) {
	alipayInvoiceExportRequest := &AlipayInvoiceExportRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceExportRequest, "/ams/api/v1/billing/invoice/export", &responseBilling.AlipayInvoiceExportResponse{})
	return alipayRequest, alipayInvoiceExportRequest
}

func (alipayInvoiceExportRequest *AlipayInvoiceExportRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceExportRequest, "/ams/api/v1/billing/invoice/export", &responseBilling.AlipayInvoiceExportResponse{})
}
