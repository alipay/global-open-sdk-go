package model

type Invoice struct {
	InvoiceId         string  `json:"invoiceId,omitempty"`
	SubscriptionId    string  `json:"subscriptionId,omitempty"`
	CustomerId        string  `json:"customerId,omitempty"`
	CustomerFirstName string  `json:"customerFirstName,omitempty"`
	CustomerLastName  string  `json:"customerLastName,omitempty"`
	CustomerEmail     string  `json:"customerEmail,omitempty"`
	Reason            string  `json:"reason,omitempty"`
	Status            string  `json:"status,omitempty"`
	TotalAmount       *Amount `json:"totalAmount,omitempty"`
	PaidAmount        *Amount `json:"paidAmount,omitempty"`
	RemainingAmount   *Amount `json:"remainingAmount,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	PaidTime          string  `json:"paidTime,omitempty"`
	VoidedTime        string  `json:"voidedTime,omitempty"`
	PeriodStart       string  `json:"periodStart,omitempty"`
	PeriodEnd         string  `json:"periodEnd,omitempty"`
	DueDate           string  `json:"dueDate,omitempty"`
	GmtCreate         string  `json:"gmtCreate,omitempty"`
	GmtUpdate         string  `json:"gmtUpdate,omitempty"`
	Description       string  `json:"description,omitempty"`
	PdfFileUrl        string  `json:"pdfFileUrl,omitempty"`
}
