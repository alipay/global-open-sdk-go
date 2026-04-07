package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireAvailableQuotaRequest struct {
        Currency string `json:"currency,omitempty"`
}

func NewAlipayInquireAvailableQuotaRequest() (*request.AlipayRequest, *AlipayInquireAvailableQuotaRequest) { 
alipayInquireAvailableQuotaRequest := &AlipayInquireAvailableQuotaRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireAvailableQuotaRequest,  "null", &responseAba.AlipayInquireAvailableQuotaResponse{}) 
return alipayRequest, alipayInquireAvailableQuotaRequest 
} 
 
func (alipayInquireAvailableQuotaRequest *AlipayInquireAvailableQuotaRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireAvailableQuotaRequest,  "null", &responseAba.AlipayInquireAvailableQuotaResponse{}) 
}







