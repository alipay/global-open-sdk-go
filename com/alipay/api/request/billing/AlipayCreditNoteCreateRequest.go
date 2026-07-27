package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditNoteCreateRequest struct {
	CreditNoteRequestId string            `json:"creditNoteRequestId,omitempty"`
	InvoiceId           string            `json:"invoiceId,omitempty"`
	Type                string            `json:"type,omitempty"`
	TotalAmount         *model.Amount     `json:"totalAmount,omitempty"`
	Items               []*model.LineItem `json:"items,omitempty"`
	RefundAmount        *model.Amount     `json:"refundAmount,omitempty"`
	RefundDestination   string            `json:"refundDestination,omitempty"`
	Reason              string            `json:"reason,omitempty"`
	ReasonDescription   string            `json:"reasonDescription,omitempty"`
	Memo                string            `json:"memo,omitempty"`
	EmailType           string            `json:"emailType,omitempty"`
	Language            string            `json:"language,omitempty"`
	EffectiveDate       string            `json:"effectiveDate,omitempty"`
	Metadata            map[string]string `json:"metadata,omitempty"`
}

func NewAlipayCreditNoteCreateRequest() (*request.AlipayRequest, *AlipayCreditNoteCreateRequest) {
	alipayCreditNoteCreateRequest := &AlipayCreditNoteCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditNoteCreateRequest, "/ams/api/v1/billing/creditNote/create", &responseBilling.AlipayCreditNoteCreateResponse{})
	return alipayRequest, alipayCreditNoteCreateRequest
}

func (alipayCreditNoteCreateRequest *AlipayCreditNoteCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditNoteCreateRequest, "/ams/api/v1/billing/creditNote/create", &responseBilling.AlipayCreditNoteCreateResponse{})
}
