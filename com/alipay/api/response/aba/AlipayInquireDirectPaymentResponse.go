package responseAba

import (
"github.com/alipay/global-open-sdk-go/com/alipay/api/model" 
"github.com/alipay/global-open-sdk-go/com/alipay/api/response" 

)




type AlipayInquireDirectPaymentResponse struct {
response.AlipayResponse 
        PaymentStatus model.TransactionStatusType `json:"paymentStatus,omitempty"` 
        PaymentResultMessage string `json:"paymentResultMessage,omitempty"`
        PaymentResultCode string `json:"paymentResultCode,omitempty"`
        PaymentId string `json:"paymentId,omitempty"`
        PaymentRequestId string `json:"paymentRequestId,omitempty"`
        PayToMethod *model.PaymentMethod `json:"payToMethod,omitempty"` 
        PaymentAmount *model.Amount `json:"paymentAmount,omitempty"` 
        PayToAmount *model.Amount `json:"payToAmount,omitempty"` 
        PaymentTime string `json:"paymentTime,omitempty"`
        Result *model.Result `json:"result,omitempty"` 
}








