package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayPayResultNotify struct {
	AlipayNotify
	PaymentRequestId         string                   `json:"paymentRequestId,omitempty"`
	PaymentId                string                   `json:"paymentId,omitempty"`
	PaymentAmount            *model.Amount            `json:"paymentAmount,omitempty"`
	PaymentCreateTime        string                   `json:"paymentCreateTime,omitempty"`
	PaymentTime              string                   `json:"paymentTime,omitempty"`
	CustomsDeclarationAmount *model.Amount            `json:"customsDeclarationAmount,omitempty"`
	GrossSettlementAmount    *model.Amount            `json:"grossSettlementAmount,omitempty"`
	SettlementQuote          *model.Quote             `json:"settlementQuote,omitempty"`
	PspCustomerInfo          *model.PspCustomerInfo   `json:"pspCustomerInfo,omitempty"`
	AcquirerReferenceNo      string                   `json:"acquirerReferenceNo,omitempty"`
	PaymentResultInfo        *model.PaymentResultInfo `json:"paymentResultInfo,omitempty"`
	AcquirerInfo             *model.AcquirerInfo      `json:"acquirerInfo,omitempty"`
	PromotionResult          []*model.PromotionResult `json:"promotionResult,omitempty"`
	PaymentMethodType        string                   `json:"paymentMethodType,omitempty"`
	Metadata                 string                   `json:"metadata,omitempty"`
	SubscriptionOrderId      string                   `json:"subscriptionOrderId,omitempty"`
	RetryInfo                *model.RetryInfo         `json:"retryInfo,omitempty"`
	UpdateRequestId          string                   `json:"updateRequestId,omitempty"`
	AuthExpiryTime           string                   `json:"authExpiryTime,omitempty"`
	// TaxCalculationId identifies the tax calculation associated with the payment. Retain it for reconciliation and subsequent refunds; query tax details through inquireTransactionList. It does not indicate that tax has been posted or recorded. If absent, the payment is not subject to tax.
	TaxCalculationId string `json:"taxCalculationId,omitempty"`
}
