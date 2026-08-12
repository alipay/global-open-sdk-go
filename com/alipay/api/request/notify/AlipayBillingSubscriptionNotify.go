package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayBillingSubscriptionNotify struct {
	AlipayNotify
	// SubscriptionRequestId identifies the original subscription request and serves as the merchant-side idempotency key for notification deduplication.
	SubscriptionRequestId string        `json:"subscriptionRequestId,omitempty"`
	EventTime             string        `json:"eventTime,omitempty"`
	SubscriptionId        string        `json:"subscriptionId,omitempty"`
	InvoiceId             string        `json:"invoiceId,omitempty"`
	Status                string        `json:"status,omitempty"`
	Reason                string        `json:"reason,omitempty"`
	ReasonDescription     string        `json:"reasonDescription,omitempty"`
	PreviousStatus        string        `json:"previousStatus,omitempty"`
	FixedAmount           *model.Amount `json:"fixedAmount,omitempty"`
}
