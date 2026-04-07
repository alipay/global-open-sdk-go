package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireExchangeRequest struct {
        ExchangeRequestId string `json:"exchangeRequestId,omitempty"`
}

func NewAlipayInquireExchangeRequest() (*request.AlipayRequest, *AlipayInquireExchangeRequest) { 
alipayInquireExchangeRequest := &AlipayInquireExchangeRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireExchangeRequest,  "null", &responseAba.AlipayInquireExchangeResponse{}) 
return alipayRequest, alipayInquireExchangeRequest 
} 
 
func (alipayInquireExchangeRequest *AlipayInquireExchangeRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireExchangeRequest,  "null", &responseAba.AlipayInquireExchangeResponse{}) 
}







