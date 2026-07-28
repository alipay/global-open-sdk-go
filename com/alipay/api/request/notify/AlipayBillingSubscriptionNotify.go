package notify

type AlipayBillingSubscriptionNotify struct {
	AlipayNotify
	MerchantRequestId string `json:"merchantRequestId,omitempty"`
	EventTime         string `json:"eventTime,omitempty"`
	SubscriptionId    string `json:"subscriptionId,omitempty"`
	InvoiceId         string `json:"invoiceId,omitempty"`
	Status            string `json:"status,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ReasonDescription string `json:"reasonDescription,omitempty"`
	PreviousStatus    string `json:"previousStatus,omitempty"`
}
