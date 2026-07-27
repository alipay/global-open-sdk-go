package notify

type AlipayInvoiceNotify struct {
	AlipayNotify
	MerchantRequestId string `json:"merchantRequestId,omitempty"`
	EventTime         string `json:"eventTime,omitempty"`
	InvoiceId         string `json:"invoiceId,omitempty"`
	SubscriptionId    string `json:"subscriptionId,omitempty"`
	CustomerId        string `json:"customerId,omitempty"`
	Status            string `json:"status,omitempty"`
	PreviousStatus    string `json:"previousStatus,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ReasonDescription string `json:"reasonDescription,omitempty"`
}
