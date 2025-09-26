package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayRefundResponse struct {
	response.AlipayResponse
	Result                         *model.Result       `json:"result,omitempty"`
	RefundRequestId                string              `json:"refundRequestId,omitempty"`
	RefundId                       string              `json:"refundId,omitempty"`
	PaymentId                      string              `json:"paymentId,omitempty"`
	RefundAmount                   *model.Amount       `json:"refundAmount,omitempty"`
	RefundTime                     string              `json:"refundTime,omitempty"`
	RefundNonGuaranteeCouponAmount *model.Amount       `json:"refundNonGuaranteeCouponAmount,omitempty"`
	GrossSettlementAmount          *model.Amount       `json:"grossSettlementAmount,omitempty"`
	SettlementQuote                *model.Quote        `json:"settlementQuote,omitempty"`
	AcquirerInfo                   *model.AcquirerInfo `json:"acquirerInfo,omitempty"`
	AcquirerReferenceNo            string              `json:"acquirerReferenceNo,omitempty"`
}
