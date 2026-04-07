package pay

import (
responsePay  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayInquiryRefundRequest struct {
        RefundRequestId string `json:"refundRequestId,omitempty"`
        RefundId string `json:"refundId,omitempty"`
        MerchantAccountId string `json:"merchantAccountId,omitempty"`
}

func NewAlipayInquiryRefundRequest() (*request.AlipayRequest, *AlipayInquiryRefundRequest) { 
alipayInquiryRefundRequest := &AlipayInquiryRefundRequest{} 
alipayRequest := request.NewAlipayRequest (alipayInquiryRefundRequest,  "null", &responsePay.AlipayInquiryRefundResponse{}) 
return alipayRequest, alipayInquiryRefundRequest 
} 
 
func (alipayInquiryRefundRequest *AlipayInquiryRefundRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayInquiryRefundRequest,  "null", &responsePay.AlipayInquiryRefundResponse{}) 
}







