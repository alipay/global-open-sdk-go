package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayCreditNoteNotify struct {
	AlipayNotify
	NotifyId                   string                   `json:"notifyId,omitempty"`
	CreditNoteRequestId        string                   `json:"creditNoteRequestId,omitempty"`
	CreditNoteNotificationType string                   `json:"creditNoteNotificationType,omitempty"`
	CustomerId                 string                   `json:"customerId,omitempty"`
	CreditNote                 *model.CreditNoteInfo    `json:"creditNote,omitempty"`
	Invoice                    *model.NotifyInvoiceInfo `json:"invoice,omitempty"`
}
