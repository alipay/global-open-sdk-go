package notify

import "github.com/alipay/global-open-sdk-go/com/alipay/api/model"

type AlipayDirectRefundNotify struct {
	AlipayNotify
	RefundStatus     string                  `json:"refundStatus,omitempty"`
	RefundId         string                  `json:"refundId,omitempty"`
	PaymentId        string                  `json:"paymentId,omitempty"`
	RefundRequestId  string                  `json:"refundRequestId,omitempty"`
	RefundTime       string                  `json:"refundTime,omitempty"`
	RefundFromMethod *model.RefundFromMethod `json:"refundFromMethod,omitempty"`
	RefundToAmount   *model.Amount           `json:"refundToAmount,omitempty"`
	RefundFromAmount *model.Amount           `json:"refundFromAmount,omitempty"`
}
