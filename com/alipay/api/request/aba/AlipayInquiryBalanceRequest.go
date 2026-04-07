package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquiryBalanceRequest struct {
        CurrencyList []string `json:"currencyList,omitempty"`
        Accesstoken string `json:"accesstoken,omitempty"`
        CustomerId string `json:"customerId,omitempty"`
}

func NewAlipayInquiryBalanceRequest() (*request.AlipayRequest, *AlipayInquiryBalanceRequest) { 
alipayInquiryBalanceRequest := &AlipayInquiryBalanceRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquiryBalanceRequest,  "null", &responseAba.AlipayInquiryBalanceResponse{}) 
return alipayRequest, alipayInquiryBalanceRequest 
} 
 
func (alipayInquiryBalanceRequest *AlipayInquiryBalanceRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquiryBalanceRequest,  "null", &responseAba.AlipayInquiryBalanceResponse{}) 
}







