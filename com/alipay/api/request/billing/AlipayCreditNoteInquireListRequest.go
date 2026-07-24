package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditNoteInquireListRequest struct {
	InvoiceId     string `json:"invoiceId,omitempty"`
	CustomerId    string `json:"customerId,omitempty"`
	Status        string `json:"status,omitempty"`
	Type          string `json:"type,omitempty"`
	Reason        string `json:"reason,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	StartingAfter string `json:"startingAfter,omitempty"`
	EndingBefore  string `json:"endingBefore,omitempty"`
	Limit         int32  `json:"limit,omitempty"`
	IncludeTotal  bool   `json:"includeTotal,omitempty"`
}

func NewAlipayCreditNoteInquireListRequest() (*request.AlipayRequest, *AlipayCreditNoteInquireListRequest) {
	alipayCreditNoteInquireListRequest := &AlipayCreditNoteInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditNoteInquireListRequest, "/ams/api/v1/billing/creditNote/inquireList", &responseBilling.AlipayCreditNoteInquireListResponse{})
	return alipayRequest, alipayCreditNoteInquireListRequest
}

func (alipayCreditNoteInquireListRequest *AlipayCreditNoteInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditNoteInquireListRequest, "/ams/api/v1/billing/creditNote/inquireList", &responseBilling.AlipayCreditNoteInquireListResponse{})
}
