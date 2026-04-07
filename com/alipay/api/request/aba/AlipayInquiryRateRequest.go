package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquiryRateRequest struct {
        RateConditionList[] *model.InquiryRateCondition `json:"rateConditionList,omitempty"` 
}

func NewAlipayInquiryRateRequest() (*request.AlipayRequest, *AlipayInquiryRateRequest) { 
alipayInquiryRateRequest := &AlipayInquiryRateRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquiryRateRequest,  "null", &responseAba.AlipayInquiryRateResponse{}) 
return alipayRequest, alipayInquiryRateRequest 
} 
 
func (alipayInquiryRateRequest *AlipayInquiryRateRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquiryRateRequest,  "null", &responseAba.AlipayInquiryRateResponse{}) 
}







