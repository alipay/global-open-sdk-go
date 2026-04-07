package dispute

import (
responseDispute  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayAcceptDisputeRequest struct {
        DisputeId string `json:"disputeId,omitempty"`
}

func NewAlipayAcceptDisputeRequest() (*request.AlipayRequest, *AlipayAcceptDisputeRequest) { 
alipayAcceptDisputeRequest := &AlipayAcceptDisputeRequest{} 
alipayRequest := request.NewAlipayRequest (alipayAcceptDisputeRequest,  "null", &responseDispute.AlipayAcceptDisputeResponse{}) 
return alipayRequest, alipayAcceptDisputeRequest 
} 
 
func (alipayAcceptDisputeRequest *AlipayAcceptDisputeRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayAcceptDisputeRequest,  "null", &responseDispute.AlipayAcceptDisputeResponse{}) 
}







