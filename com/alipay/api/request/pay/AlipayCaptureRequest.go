package pay

import (
responsePay  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayCaptureRequest struct {
        CaptureRequestId string `json:"captureRequestId,omitempty"`
        PaymentId string `json:"paymentId,omitempty"`
        CaptureAmount *model.Amount `json:"captureAmount,omitempty"` 
        IsLastCapture bool `json:"isLastCapture,omitempty"`
        CaptureType string `json:"captureType,omitempty"`
        Transit *model.Transit `json:"transit,omitempty"` 
}

func NewAlipayCaptureRequest() (*request.AlipayRequest, *AlipayCaptureRequest) { 
alipayCaptureRequest := &AlipayCaptureRequest{} 
alipayRequest := request.NewAlipayRequest (alipayCaptureRequest,  "null", &responsePay.AlipayCaptureResponse{}) 
return alipayRequest, alipayCaptureRequest 
} 
 
func (alipayCaptureRequest *AlipayCaptureRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayCaptureRequest,  "null", &responsePay.AlipayCaptureResponse{}) 
}







