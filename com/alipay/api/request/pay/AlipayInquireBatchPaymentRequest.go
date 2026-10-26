package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayInquireBatchPaymentRequest struct {
	BatchPaymentRequestId string `json:"batchPaymentRequestId,omitempty"`
}

func NewAlipayInquireBatchPaymentRequest() (*request.AlipayRequest, *AlipayInquireBatchPaymentRequest) {
	alipayInquireBatchPaymentRequest := &AlipayInquireBatchPaymentRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireBatchPaymentRequest, "/ams/api/v1/payments/inquireBatchPayment", &responsePay.AlipayInquireBatchPaymentResponse{})
	return alipayRequest, alipayInquireBatchPaymentRequest
}

func (alipayInquireBatchPaymentRequest *AlipayInquireBatchPaymentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireBatchPaymentRequest, "/ams/api/v1/payments/inquireBatchPayment", &responsePay.AlipayInquireBatchPaymentResponse{})
}
