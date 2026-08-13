package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayCaptureResultNotify struct {
	AlipayNotify
	CaptureRequestId    string              `json:"captureRequestId,omitempty"`
	PaymentId           string              `json:"paymentId,omitempty"`
	CaptureId           string              `json:"captureId,omitempty"`
	CaptureAmount       *model.Amount       `json:"captureAmount,omitempty"`
	CaptureTime         string              `json:"captureTime,omitempty"`
	AcquirerReferenceNo string              `json:"acquirerReferenceNo,omitempty"`
	AcquirerInfo        *model.AcquirerInfo `json:"acquirerInfo,omitempty"`
	// TaxCalculationId identifies the tax calculation associated with the payment. Retain it for reconciliation and subsequent refunds; query tax details through inquireTransactionList. It does not indicate that tax has been posted or recorded. If absent, the payment is not subject to tax.
	TaxCalculationId string `json:"taxCalculationId,omitempty"`
}
