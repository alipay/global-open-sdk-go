package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxInquireTransactionListRequest struct {
	TaxCalculationId string `json:"taxCalculationId,omitempty"`
	PaymentId        string `json:"paymentId,omitempty"`
	RefundId         string `json:"refundId,omitempty"`
	CurrentPage      int32  `json:"currentPage,omitempty"`
	PageSize         int32  `json:"pageSize,omitempty"`
}

func NewAlipayTaxInquireTransactionListRequest() (*request.AlipayRequest, *AlipayTaxInquireTransactionListRequest) {
	alipayTaxInquireTransactionListRequest := &AlipayTaxInquireTransactionListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxInquireTransactionListRequest, "/ams/api/v1/tax/inquireTransactionList", &responseBilling.AlipayTaxInquireTransactionListResponse{})
	return alipayRequest, alipayTaxInquireTransactionListRequest
}

func (alipayTaxInquireTransactionListRequest *AlipayTaxInquireTransactionListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxInquireTransactionListRequest, "/ams/api/v1/tax/inquireTransactionList", &responseBilling.AlipayTaxInquireTransactionListResponse{})
}
