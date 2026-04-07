package dispute

import (
responseDispute  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipaySupplyDefenseDocumentRequest struct {
        DisputeId string `json:"disputeId,omitempty"`
        DisputeEvidence string `json:"disputeEvidence,omitempty"`
}

func NewAlipaySupplyDefenseDocumentRequest() (*request.AlipayRequest, *AlipaySupplyDefenseDocumentRequest) { 
alipaySupplyDefenseDocumentRequest := &AlipaySupplyDefenseDocumentRequest{} 
alipayRequest := request.NewAlipayRequest (alipaySupplyDefenseDocumentRequest,  "null", &responseDispute.AlipaySupplyDefenseDocumentResponse{}) 
return alipayRequest, alipaySupplyDefenseDocumentRequest 
} 
 
func (alipaySupplyDefenseDocumentRequest *AlipaySupplyDefenseDocumentRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipaySupplyDefenseDocumentRequest,  "null", &responseDispute.AlipaySupplyDefenseDocumentResponse{}) 
}







