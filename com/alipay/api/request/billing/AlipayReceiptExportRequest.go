package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayReceiptExportRequest struct {
	Limit          int32    `json:"limit,omitempty"`
	Status         string   `json:"status,omitempty"`
	ReceiptType    string   `json:"receiptType,omitempty"`
	InvoiceId      string   `json:"invoiceId,omitempty"`
	SubscriptionId string   `json:"subscriptionId,omitempty"`
	CustomerId     string   `json:"customerId,omitempty"`
	StartDate      string   `json:"startDate,omitempty"`
	EndDate        string   `json:"endDate,omitempty"`
	ReceiptIds     []string `json:"receiptIds,omitempty"`
}

func NewAlipayReceiptExportRequest() (*request.AlipayRequest, *AlipayReceiptExportRequest) {
	alipayReceiptExportRequest := &AlipayReceiptExportRequest{}
	alipayRequest := request.NewAlipayRequest(alipayReceiptExportRequest, "/ams/api/v1/billing/receipt/export", &responseBilling.AlipayReceiptExportResponse{})
	return alipayRequest, alipayReceiptExportRequest
}

func (alipayReceiptExportRequest *AlipayReceiptExportRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayReceiptExportRequest, "/ams/api/v1/billing/receipt/export", &responseBilling.AlipayReceiptExportResponse{})
}
