package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreditNoteInquireDetailsResponse struct {
	response.AlipayResponse
	Result              *model.Result                        `json:"result,omitempty"`
	CreditNoteId        string                               `json:"creditNoteId,omitempty"`
	CreditNoteRequestId string                               `json:"creditNoteRequestId,omitempty"`
	InvoiceId           string                               `json:"invoiceId,omitempty"`
	CustomerId          string                               `json:"customerId,omitempty"`
	Type                string                               `json:"type,omitempty"`
	Status              string                               `json:"status,omitempty"`
	TotalAmount         *model.Amount                        `json:"totalAmount,omitempty"`
	RefundStatus        string                               `json:"refundStatus,omitempty"`
	RefundId            string                               `json:"refundId,omitempty"`
	Reason              string                               `json:"reason,omitempty"`
	ReasonDescription   string                               `json:"reasonDescription,omitempty"`
	RefundDestination   string                               `json:"refundDestination,omitempty"`
	EffectiveDate       string                               `json:"effectiveDate,omitempty"`
	Memo                string                               `json:"memo,omitempty"`
	Items               *model.CreditNoteInquireDetailsItems `json:"items,omitempty"`
	IssuedAt            string                               `json:"issuedAt,omitempty"`
	RefundedAt          string                               `json:"refundedAt,omitempty"`
	VoidedAt            string                               `json:"voidedAt,omitempty"`
	CreatedAt           string                               `json:"createdAt,omitempty"`
}
