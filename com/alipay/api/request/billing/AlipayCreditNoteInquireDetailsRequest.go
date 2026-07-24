package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditNoteInquireDetailsRequest struct {
	CreditNoteId string `json:"creditNoteId,omitempty"`
}

func NewAlipayCreditNoteInquireDetailsRequest() (*request.AlipayRequest, *AlipayCreditNoteInquireDetailsRequest) {
	alipayCreditNoteInquireDetailsRequest := &AlipayCreditNoteInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditNoteInquireDetailsRequest, "/ams/api/v1/billing/creditNote/inquireDetails", &responseBilling.AlipayCreditNoteInquireDetailsResponse{})
	return alipayRequest, alipayCreditNoteInquireDetailsRequest
}

func (alipayCreditNoteInquireDetailsRequest *AlipayCreditNoteInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditNoteInquireDetailsRequest, "/ams/api/v1/billing/creditNote/inquireDetails", &responseBilling.AlipayCreditNoteInquireDetailsResponse{})
}
