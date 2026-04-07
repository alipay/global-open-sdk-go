package responseAba

import (
"github.com/alipay/global-open-sdk-go/com/alipay/api/model" 
"github.com/alipay/global-open-sdk-go/com/alipay/api/response" 

)




type AlipayInquireDirectRefundResponse struct {
response.AlipayResponse 
        RefundStatus model.TransactionStatusType `json:"refundStatus,omitempty"` 
        RefundResultMessage string `json:"refundResultMessage,omitempty"`
        RefundResultCode string `json:"refundResultCode,omitempty"`
        RefundId string `json:"refundId,omitempty"`
        PaymentId string `json:"paymentId,omitempty"`
        RefundRequestId string `json:"refundRequestId,omitempty"`
        RefundTime string `json:"refundTime,omitempty"`
        RefundFromMethod *model.RefundFromMethod `json:"refundFromMethod,omitempty"` 
        RefundToAmount *model.Amount `json:"refundToAmount,omitempty"` 
        RefundFromAmount *model.Amount `json:"refundFromAmount,omitempty"` 
        Result *model.Result `json:"result,omitempty"` 
}








