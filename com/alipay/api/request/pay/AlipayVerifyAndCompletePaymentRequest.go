package pay

import (
responsePay  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayVerifyAndCompletePaymentRequest struct {
        VerifyMethod *model.VerifyMethod `json:"verifyMethod,omitempty"` 
        VerifyRequestId string `json:"verifyRequestId,omitempty"`
        PaymentId string `json:"paymentId,omitempty"`
}

func NewAlipayVerifyAndCompletePaymentRequest() (*request.AlipayRequest, *AlipayVerifyAndCompletePaymentRequest) { 
alipayVerifyAndCompletePaymentRequest := &AlipayVerifyAndCompletePaymentRequest{} 
alipayRequest := request.NewAlipayRequest (alipayVerifyAndCompletePaymentRequest,  "null", &responsePay.AlipayVerifyAndCompletePaymentResponse{}) 
return alipayRequest, alipayVerifyAndCompletePaymentRequest 
} 
 
func (alipayVerifyAndCompletePaymentRequest *AlipayVerifyAndCompletePaymentRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayVerifyAndCompletePaymentRequest,  "null", &responsePay.AlipayVerifyAndCompletePaymentResponse{}) 
}







