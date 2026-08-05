package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayInvoiceNotify struct {
	AlipayNotify
	InvoiceRequestId  string                    `json:"invoiceRequestId,omitempty"`
	InvoiceId         string                    `json:"invoiceId,omitempty"`
	InvoiceStatus     string                    `json:"invoiceStatus,omitempty"`
	InvoiceAmount     *model.Amount             `json:"invoiceAmount,omitempty"`
	PaymentInfo       *model.InvoicePaymentInfo `json:"paymentInfo,omitempty"`
	Subscription      *model.SubscriptionInfo   `json:"subscription,omitempty"`
	CustomerId        string                    `json:"customerId,omitempty"`
	Reason            string                    `json:"reason,omitempty"`
	ReasonDescription string                    `json:"reasonDescription,omitempty"`
}
