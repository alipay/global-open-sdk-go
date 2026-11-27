package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

// AlipayCreditNoteNotify is the CREDIT_NOTE_RESULT payload. ACK uses a separate standard Result envelope.
type AlipayCreditNoteNotify struct {
	AlipayNotify
	// Required creation request ID shared by all events; correlation, not delivery deduplication.
	// Delivery deduplication uses outer notify_id, which is not a payload member.
	CreditNoteRequestId        string                   `json:"creditNoteRequestId,omitempty"`
	CreditNoteNotificationType string                   `json:"creditNoteNotificationType,omitempty"`
	CustomerId                 string                   `json:"customerId,omitempty"`
	CreditNote                 *model.CreditNoteInfo    `json:"creditNote,omitempty"`
	Invoice                    *model.NotifyInvoiceInfo `json:"invoice,omitempty"`
}
