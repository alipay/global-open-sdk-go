package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayCreateExchangeRequest struct {
	BuyAmount         *model.Amount `json:"buyAmount,omitempty"`
	SellAmount        *model.Amount `json:"sellAmount,omitempty"`
	ExchangeTradeType string        `json:"exchangeTradeType,omitempty"`
	ExchangeRequestId string        `json:"exchangeRequestId,omitempty"`
	ExchangeMode      string        `json:"exchangeMode,omitempty"`
}

func NewAlipayCreateExchangeRequest() (*request.AlipayRequest, *AlipayCreateExchangeRequest) {
	alipayCreateExchangeRequest := &AlipayCreateExchangeRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateExchangeRequest, "/ams/v1/aba/funds/createExchange", &responseAba.AlipayCreateExchangeResponse{})
	return alipayRequest, alipayCreateExchangeRequest
}

func (alipayCreateExchangeRequest *AlipayCreateExchangeRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateExchangeRequest, "/ams/v1/aba/funds/createExchange", &responseAba.AlipayCreateExchangeResponse{})
}
