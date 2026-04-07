package pay

import (
responsePay  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayPayCancelRequest struct {
        PaymentId string `json:"paymentId,omitempty"`
        PaymentRequestId string `json:"paymentRequestId,omitempty"`
        MerchantAccountId string `json:"merchantAccountId,omitempty"`
}

func NewAlipayPayCancelRequest() (*request.AlipayRequest, *AlipayPayCancelRequest) { 
alipayPayCancelRequest := &AlipayPayCancelRequest{} 
alipayRequest := request.NewAlipayRequest (alipayPayCancelRequest,  "null", &responsePay.AlipayPayCancelResponse{}) 
return alipayRequest, alipayPayCancelRequest 
} 
 
func (alipayPayCancelRequest *AlipayPayCancelRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayPayCancelRequest,  "null", &responsePay.AlipayPayCancelResponse{}) 
}







