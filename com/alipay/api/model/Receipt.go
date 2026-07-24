package model

type Receipt struct {
	ReceiptId             string  `json:"receiptId,omitempty"`
	InvoiceId             string  `json:"invoiceId,omitempty"`
	CustomerId            string  `json:"customerId,omitempty"`
	SubscriptionId        string  `json:"subscriptionId,omitempty"`
	OriginalReceiptId     string  `json:"originalReceiptId,omitempty"`
	ReceiptType           string  `json:"receiptType,omitempty"`
	Status                string  `json:"status,omitempty"`
	Reason                string  `json:"reason,omitempty"`
	CollectionMethod      string  `json:"collectionMethod,omitempty"`
	PaymentMethod         string  `json:"paymentMethod,omitempty"`
	Subtotal              *Amount `json:"subtotal,omitempty"`
	TotalAmount           *Amount `json:"totalAmount,omitempty"`
	PaidAmount            *Amount `json:"paidAmount,omitempty"`
	RemainingAmount       *Amount `json:"remainingAmount,omitempty"`
	RefundAmount          *Amount `json:"refundAmount,omitempty"`
	RefundedAmount        *Amount `json:"refundedAmount,omitempty"`
	PaymentDeductedAmount *Amount `json:"paymentDeductedAmount,omitempty"`
	PeriodStart           string  `json:"periodStart,omitempty"`
	PeriodEnd             string  `json:"periodEnd,omitempty"`
	Description           string  `json:"description,omitempty"`
	GmtCreate             string  `json:"gmtCreate,omitempty"`
	GmtUpdate             string  `json:"gmtUpdate,omitempty"`
	CustomerFirstName     string  `json:"customerFirstName,omitempty"`
	CustomerLastName      string  `json:"customerLastName,omitempty"`
	CustomerEmail         string  `json:"customerEmail,omitempty"`
	PaymentMethodType     string  `json:"paymentMethodType,omitempty"`
	DiscountAmount        *Amount `json:"discountAmount,omitempty"`
	TaxAmount             *Amount `json:"taxAmount,omitempty"`
	ShippingFeeAmount     *Amount `json:"shippingFeeAmount,omitempty"`
	SettlementAmount      *Amount `json:"settlementAmount,omitempty"`
	FxRate                string  `json:"fxRate,omitempty"`
	FxRateId              string  `json:"fxRateId,omitempty"`
	DueDate               string  `json:"dueDate,omitempty"`
	PaidTime              string  `json:"paidTime,omitempty"`
	PaymentRequestId      string  `json:"paymentRequestId,omitempty"`
	PayToRequestId        string  `json:"payToRequestId,omitempty"`
	PayToId               string  `json:"payToId,omitempty"`
	Footer                string  `json:"footer,omitempty"`
	FileUrl               string  `json:"fileUrl,omitempty"`
}
