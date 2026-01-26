package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayCreateQuoteRequest struct {
	BuyAmount         *model.Amount `json:"buyAmount,omitempty"`
	SellAmount        *model.Amount `json:"sellAmount,omitempty"`
	ExchangeTradeType string        `json:"exchangeTradeType,omitempty"`
}

func NewAlipayCreateQuoteRequest() (*request.AlipayRequest, *AlipayCreateQuoteRequest) {
	alipayCreateQuoteRequest := &AlipayCreateQuoteRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateQuoteRequest, "/ams/v1/aba/funds/createQuote", &responseAba.AlipayCreateQuoteResponse{})
	return alipayRequest, alipayCreateQuoteRequest
}

func (alipayCreateQuoteRequest *AlipayCreateQuoteRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateQuoteRequest, "/ams/v1/aba/funds/createQuote", &responseAba.AlipayCreateQuoteResponse{})
}
