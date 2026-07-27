package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditNoteVoidRequest struct {
	CreditNoteId string `json:"creditNoteId,omitempty"`
	Reason       string `json:"reason,omitempty"`
}

func NewAlipayCreditNoteVoidRequest() (*request.AlipayRequest, *AlipayCreditNoteVoidRequest) {
	alipayCreditNoteVoidRequest := &AlipayCreditNoteVoidRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditNoteVoidRequest, "/ams/api/v1/billing/creditNote/void", &responseBilling.AlipayCreditNoteVoidResponse{})
	return alipayRequest, alipayCreditNoteVoidRequest
}

func (alipayCreditNoteVoidRequest *AlipayCreditNoteVoidRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditNoteVoidRequest, "/ams/api/v1/billing/creditNote/void", &responseBilling.AlipayCreditNoteVoidResponse{})
}
