package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayInquireExchangeRequest struct {
	ExchangeRequestId string `json:"exchangeRequestId,omitempty"`
}

func NewAlipayInquireExchangeRequest() (*request.AlipayRequest, *AlipayInquireExchangeRequest) {
	alipayInquireExchangeRequest := &AlipayInquireExchangeRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireExchangeRequest, "/ams/v1/aba/funds/inquireExchange", &responseAba.AlipayInquireExchangeResponse{})
	return alipayRequest, alipayInquireExchangeRequest
}

func (alipayInquireExchangeRequest *AlipayInquireExchangeRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireExchangeRequest, "/ams/v1/aba/funds/inquireExchange", &responseAba.AlipayInquireExchangeResponse{})
}
