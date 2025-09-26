package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryRefundResponse struct {
	response.AlipayResponse
	Result                *model.Result               `json:"result,omitempty"`
	RefundId              string                      `json:"refundId,omitempty"`
	RefundRequestId       string                      `json:"refundRequestId,omitempty"`
	RefundAmount          *model.Amount               `json:"refundAmount,omitempty"`
	RefundStatus          model.TransactionStatusType `json:"refundStatus,omitempty"`
	RefundTime            string                      `json:"refundTime,omitempty"`
	GrossSettlementAmount *model.Amount               `json:"grossSettlementAmount,omitempty"`
	SettlementQuote       *model.Quote                `json:"settlementQuote,omitempty"`
	AcquirerInfo          *model.AcquirerInfo         `json:"acquirerInfo,omitempty"`
}
