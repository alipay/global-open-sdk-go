package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayPayQueryRequest struct {
	PaymentRequestId  string `json:"paymentRequestId,omitempty"`
	PaymentId         string `json:"paymentId,omitempty"`
	MerchantAccountId string `json:"merchantAccountId,omitempty"`
}

func NewAlipayPayQueryRequest() (*request.AlipayRequest, *AlipayPayQueryRequest) {
	alipayPayQueryRequest := &AlipayPayQueryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayQueryRequest, "/ams/api/v1/payments/inquiryPayment", &responsePay.AlipayPayQueryResponse{})
	return alipayRequest, alipayPayQueryRequest
}

func (alipayPayQueryRequest *AlipayPayQueryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPayQueryRequest, "/ams/api/v1/payments/inquiryPayment", &responsePay.AlipayPayQueryResponse{})
}
