package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayCreateQuoteRequest struct {
        BuyAmount *model.Amount `json:"buyAmount,omitempty"` 
        SellAmount *model.Amount `json:"sellAmount,omitempty"` 
        ExchangeTradeType string `json:"exchangeTradeType,omitempty"`
}

func NewAlipayCreateQuoteRequest() (*request.AlipayRequest, *AlipayCreateQuoteRequest) { 
alipayCreateQuoteRequest := &AlipayCreateQuoteRequest{} 
alipayRequest := request.NewAlipayRequest (alipayCreateQuoteRequest,  "null", &responseAba.AlipayCreateQuoteResponse{}) 
return alipayRequest, alipayCreateQuoteRequest 
} 
 
func (alipayCreateQuoteRequest *AlipayCreateQuoteRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayCreateQuoteRequest,  "null", &responseAba.AlipayCreateQuoteResponse{}) 
}







