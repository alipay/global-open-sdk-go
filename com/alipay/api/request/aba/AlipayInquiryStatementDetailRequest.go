package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquiryStatementDetailRequest struct {
        StatementId string `json:"statementId,omitempty"`
}

func NewAlipayInquiryStatementDetailRequest() (*request.AlipayRequest, *AlipayInquiryStatementDetailRequest) { 
alipayInquiryStatementDetailRequest := &AlipayInquiryStatementDetailRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquiryStatementDetailRequest,  "null", &responseAba.AlipayInquiryStatementDetailResponse{}) 
return alipayRequest, alipayInquiryStatementDetailRequest 
} 
 
func (alipayInquiryStatementDetailRequest *AlipayInquiryStatementDetailRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquiryStatementDetailRequest,  "null", &responseAba.AlipayInquiryStatementDetailResponse{}) 
}







