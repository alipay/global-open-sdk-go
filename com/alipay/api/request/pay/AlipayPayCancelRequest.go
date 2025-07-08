package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayPayCancelRequest struct {
	PaymentId         string `json:"paymentId,omitempty"`
	PaymentRequestId  string `json:"paymentRequestId,omitempty"`
	MerchantAccountId string `json:"merchantAccountId,omitempty"`
}

func NewAlipayPayCancelRequest() (*request.AlipayRequest, *AlipayPayCancelRequest) {
	alipayPayCancelRequest := &AlipayPayCancelRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayCancelRequest, "/ams/api/v1/payments/cancel", &responsePay.AlipayPayCancelResponse{})
	return alipayRequest, alipayPayCancelRequest
}

func (alipayPayCancelRequest *AlipayPayCancelRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPayCancelRequest, "/ams/api/v1/payments/cancel", &responsePay.AlipayPayCancelResponse{})
}
