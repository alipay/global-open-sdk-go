package aba

import (
responseAba  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayCreateDirectPaymentRequest struct {
        PaymentRequestId string `json:"paymentRequestId,omitempty"`
        Memo string `json:"memo,omitempty"`
        Remark string `json:"remark,omitempty"`
        Order *model.Order `json:"order,omitempty"` 
        PaymentNotifyUrl string `json:"paymentNotifyUrl,omitempty"`
        PayToMethod *model.PaymentMethod `json:"payToMethod,omitempty"` 
        PayToAmount *model.Amount `json:"payToAmount,omitempty"` 
        PayFromAmount *model.Amount `json:"payFromAmount,omitempty"` 
}

func NewAlipayCreateDirectPaymentRequest() (*request.AlipayRequest, *AlipayCreateDirectPaymentRequest) { 
alipayCreateDirectPaymentRequest := &AlipayCreateDirectPaymentRequest{} 
alipayRequest := request.NewAlipayRequest (alipayCreateDirectPaymentRequest,  "null", &responseAba.AlipayCreateDirectPaymentResponse{}) 
return alipayRequest, alipayCreateDirectPaymentRequest 
} 
 
func (alipayCreateDirectPaymentRequest *AlipayCreateDirectPaymentRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayCreateDirectPaymentRequest,  "null", &responseAba.AlipayCreateDirectPaymentResponse{}) 
}







