package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireDirectPaymentRequest struct {
	PaymentId        string `json:"paymentId,omitempty"`
	PaymentRequestId string `json:"paymentRequestId,omitempty"`
}

func NewAlipayInquireDirectPaymentRequest() (*request.AlipayRequest, *AlipayInquireDirectPaymentRequest) {
	alipayInquireDirectPaymentRequest := &AlipayInquireDirectPaymentRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireDirectPaymentRequest, "null", &responseAba.AlipayInquireDirectPaymentResponse{})
	return alipayRequest, alipayInquireDirectPaymentRequest
}

func (alipayInquireDirectPaymentRequest *AlipayInquireDirectPaymentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireDirectPaymentRequest, "null", &responseAba.AlipayInquireDirectPaymentResponse{})
}
