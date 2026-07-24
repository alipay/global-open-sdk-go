package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayReceiptInquireListRequest struct {
	StartingAfter  string `json:"startingAfter,omitempty"`
	EndingBefore   string `json:"endingBefore,omitempty"`
	Limit          int32  `json:"limit,omitempty"`
	IncludeTotal   bool   `json:"includeTotal,omitempty"`
	CustomerId     string `json:"customerId,omitempty"`
	InvoiceId      string `json:"invoiceId,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	Status         string `json:"status,omitempty"`
	ReceiptType    string `json:"receiptType,omitempty"`
	StartDate      string `json:"startDate,omitempty"`
	EndDate        string `json:"endDate,omitempty"`
}

func NewAlipayReceiptInquireListRequest() (*request.AlipayRequest, *AlipayReceiptInquireListRequest) {
	alipayReceiptInquireListRequest := &AlipayReceiptInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayReceiptInquireListRequest, "/ams/api/v1/billing/receipt/inquireList", &responseBilling.AlipayReceiptInquireListResponse{})
	return alipayRequest, alipayReceiptInquireListRequest
}

func (alipayReceiptInquireListRequest *AlipayReceiptInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayReceiptInquireListRequest, "/ams/api/v1/billing/receipt/inquireList", &responseBilling.AlipayReceiptInquireListResponse{})
}
