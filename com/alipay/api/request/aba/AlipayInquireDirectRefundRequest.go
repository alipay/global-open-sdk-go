package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireDirectRefundRequest struct {
	RefundId        string `json:"refundId,omitempty"`
	RefundRequestId string `json:"refundRequestId,omitempty"`
}

func NewAlipayInquireDirectRefundRequest() (*request.AlipayRequest, *AlipayInquireDirectRefundRequest) {
	alipayInquireDirectRefundRequest := &AlipayInquireDirectRefundRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireDirectRefundRequest, "/ams/aba/funds/inquireDirectRefund", &responseAba.AlipayInquireDirectRefundResponse{})
	return alipayRequest, alipayInquireDirectRefundRequest
}

func (alipayInquireDirectRefundRequest *AlipayInquireDirectRefundRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireDirectRefundRequest, "/ams/aba/funds/inquireDirectRefund", &responseAba.AlipayInquireDirectRefundResponse{})
}
