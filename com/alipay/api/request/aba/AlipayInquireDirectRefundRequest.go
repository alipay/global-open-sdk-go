package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquireDirectRefundRequest struct {
        RefundId string `json:"refundId,omitempty"`
        RefundRequestId string `json:"refundRequestId,omitempty"`
}

func NewAlipayInquireDirectRefundRequest() (*request.AlipayRequest, *AlipayInquireDirectRefundRequest) { 
alipayInquireDirectRefundRequest := &AlipayInquireDirectRefundRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquireDirectRefundRequest,  "null", &responseAba.AlipayInquireDirectRefundResponse{}) 
return alipayRequest, alipayInquireDirectRefundRequest 
} 
 
func (alipayInquireDirectRefundRequest *AlipayInquireDirectRefundRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquireDirectRefundRequest,  "null", &responseAba.AlipayInquireDirectRefundResponse{}) 
}







