package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceInquireDetailsResponse struct {
	response.AlipayResponse
	Result                *model.Result        `json:"result,omitempty"`
	InvoiceId             string               `json:"invoiceId,omitempty"`
	InvoiceRequestId      string               `json:"invoiceRequestId,omitempty"`
	SubscriptionId        string               `json:"subscriptionId,omitempty"`
	CustomerId            string               `json:"customerId,omitempty"`
	InvoiceNumber         string               `json:"invoiceNumber,omitempty"`
	CustomerFirstName     string               `json:"customerFirstName,omitempty"`
	CustomerLastName      string               `json:"customerLastName,omitempty"`
	CustomerEmail         string               `json:"customerEmail,omitempty"`
	Reason                string               `json:"reason,omitempty"`
	PhaseNo               string               `json:"phaseNo,omitempty"`
	Status                string               `json:"status,omitempty"`
	Currency              string               `json:"currency,omitempty"`
	Subtotal              *model.Amount        `json:"subtotal,omitempty"`
	DiscountAmount        *model.Amount        `json:"discountAmount,omitempty"`
	TotalAmount           *model.Amount        `json:"totalAmount,omitempty"`
	PaidAmount            *model.Amount        `json:"paidAmount,omitempty"`
	RemainingAmount       *model.Amount        `json:"remainingAmount,omitempty"`
	TaxAmount             *model.Amount        `json:"taxAmount,omitempty"`
	PaymentDeductedAmount *model.Amount        `json:"paymentDeductedAmount,omitempty"`
	CollectionMethod      string               `json:"collectionMethod,omitempty"`
	PaymentMethod         *model.PaymentMethod `json:"paymentMethod,omitempty"`
	Shipping              *model.Shipping      `json:"shipping,omitempty"`
	HostedInvoiceUrl      string               `json:"hostedInvoiceUrl,omitempty"`
	PeriodStart           string               `json:"periodStart,omitempty"`
	PeriodEnd             string               `json:"periodEnd,omitempty"`
	DueDate               string               `json:"dueDate,omitempty"`
	PaidTime              string               `json:"paidTime,omitempty"`
	Description           string               `json:"description,omitempty"`
	Items                 []*model.Item        `json:"items,omitempty"`
	Payments              []*model.Payment     `json:"payments,omitempty"`
	InvoiceNotes          []*model.InvoiceNote `json:"invoiceNotes,omitempty"`
	GmtCreate             string               `json:"gmtCreate,omitempty"`
	GmtUpdate             string               `json:"gmtUpdate,omitempty"`
	Discounts             []*model.Discount    `json:"discounts,omitempty"`
}
