package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayReceiptInquireDetailsResponse struct {
	response.AlipayResponse
	Result                *model.Result    `json:"result,omitempty"`
	ReceiptId             string           `json:"receiptId,omitempty"`
	OriginalReceiptId     string           `json:"originalReceiptId,omitempty"`
	InvoiceId             string           `json:"invoiceId,omitempty"`
	SubscriptionId        string           `json:"subscriptionId,omitempty"`
	CustomerId            string           `json:"customerId,omitempty"`
	PaymentId             string           `json:"paymentId,omitempty"`
	RefundId              string           `json:"refundId,omitempty"`
	ReceiptType           string           `json:"receiptType,omitempty"`
	Status                string           `json:"status,omitempty"`
	Reason                string           `json:"reason,omitempty"`
	CustomerFirstName     string           `json:"customerFirstName,omitempty"`
	CustomerLastName      string           `json:"customerLastName,omitempty"`
	CustomerEmail         string           `json:"customerEmail,omitempty"`
	CollectionMethod      string           `json:"collectionMethod,omitempty"`
	TotalAmount           *model.Amount    `json:"totalAmount,omitempty"`
	Subtotal              *model.Amount    `json:"subtotal,omitempty"`
	PaidAmount            *model.Amount    `json:"paidAmount,omitempty"`
	DiscountAmount        *model.Amount    `json:"discountAmount,omitempty"`
	TaxAmount             *model.Amount    `json:"taxAmount,omitempty"`
	ShippingFeeAmount     *model.Amount    `json:"shippingFeeAmount,omitempty"`
	PaymentDeductedAmount *model.Amount    `json:"paymentDeductedAmount,omitempty"`
	RefundAmount          *model.Amount    `json:"refundAmount,omitempty"`
	RefundedAmount        *model.Amount    `json:"refundedAmount,omitempty"`
	RemainingAmount       *model.Amount    `json:"remainingAmount,omitempty"`
	SettlementAmount      *model.Amount    `json:"settlementAmount,omitempty"`
	FxRate                string           `json:"fxRate,omitempty"`
	FxRateId              string           `json:"fxRateId,omitempty"`
	PaymentMethod         string           `json:"paymentMethod,omitempty"`
	PeriodStart           string           `json:"periodStart,omitempty"`
	PeriodEnd             string           `json:"periodEnd,omitempty"`
	PaidTime              string           `json:"paidTime,omitempty"`
	DueDate               string           `json:"dueDate,omitempty"`
	PaymentRequestId      string           `json:"paymentRequestId,omitempty"`
	PayToRequestId        string           `json:"payToRequestId,omitempty"`
	PayToId               string           `json:"payToId,omitempty"`
	Description           string           `json:"description,omitempty"`
	FileUrl               string           `json:"fileUrl,omitempty"`
	Items                 []*model.Item    `json:"items,omitempty"`
	Payments              []*model.Payment `json:"payments,omitempty"`
	GmtCreate             string           `json:"gmtCreate,omitempty"`
	GmtUpdate             string           `json:"gmtUpdate,omitempty"`
	PaymentMethodType     string           `json:"paymentMethodType,omitempty"`
	Footer                string           `json:"footer,omitempty"`
}
