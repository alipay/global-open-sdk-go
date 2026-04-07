package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateDirectRefundResponse struct {
	response.AlipayResponse
	PaymentId        string                  `json:"paymentId,omitempty"`
	RefundId         string                  `json:"refundId,omitempty"`
	RefundRequestId  string                  `json:"refundRequestId,omitempty"`
	RefundTime       string                  `json:"refundTime,omitempty"`
	RefundFromMethod *model.RefundFromMethod `json:"refundFromMethod,omitempty"`
	RefundFromAmount *model.Amount           `json:"refundFromAmount,omitempty"`
	RefundToAmount   *model.Amount           `json:"refundToAmount,omitempty"`
	Result           *model.Result           `json:"result,omitempty"`
}
