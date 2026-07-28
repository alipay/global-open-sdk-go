package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayInvoiceInquireListRequest struct {
	StartingAfter  string        `json:"startingAfter,omitempty"`
	EndingBefore   string        `json:"endingBefore,omitempty"`
	Limit          int32         `json:"limit,omitempty"`
	IncludeTotal   bool          `json:"includeTotal,omitempty"`
	SubscriptionId string        `json:"subscriptionId,omitempty"`
	CustomerId     string        `json:"customerId,omitempty"`
	InvoiceId      string        `json:"invoiceId,omitempty"`
	Status         string        `json:"status,omitempty"`
	Reason         string        `json:"reason,omitempty"`
	StartDate      string        `json:"startDate,omitempty"`
	EndDate        string        `json:"endDate,omitempty"`
	MinAmount      *model.Amount `json:"minAmount,omitempty"`
	MaxAmount      *model.Amount `json:"maxAmount,omitempty"`
}

func NewAlipayInvoiceInquireListRequest() (*request.AlipayRequest, *AlipayInvoiceInquireListRequest) {
	alipayInvoiceInquireListRequest := &AlipayInvoiceInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInvoiceInquireListRequest, "/ams/api/v1/billing/invoice/inquireList", &responseBilling.AlipayInvoiceInquireListResponse{})
	return alipayRequest, alipayInvoiceInquireListRequest
}

func (alipayInvoiceInquireListRequest *AlipayInvoiceInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInvoiceInquireListRequest, "/ams/api/v1/billing/invoice/inquireList", &responseBilling.AlipayInvoiceInquireListResponse{})
}
