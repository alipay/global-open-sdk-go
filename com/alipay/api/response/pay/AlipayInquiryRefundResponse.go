package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryRefundResponse struct {
	response.AlipayResponse
	RefundId              string                      `json:"refundId"`
	RefundRequestId       string                      `json:"refundRequestId"`
	RefundAmount          model.Amount                `json:"refundAmount"`
	RefundStatus          model.TransactionStatusType `json:"refundStatus"`
	RefundTime            string                      `json:"refundTime"`
	GrossSettlementAmount model.Amount                `json:"grossSettlementAmount"`
	SettlementQuote       model.Quote                 `json:"settlementQuote"`
	AcquirerInfo          model.AcquirerInfo          `json:"acquirerInfo"`
}
