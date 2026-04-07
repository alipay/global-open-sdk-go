package dispute

import (
responseDispute  "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute" 
 "github.com/alipay/global-open-sdk-go/com/alipay/api/model"
 "github.com/alipay/global-open-sdk-go/com/alipay/api/request"

)




type AlipayDownloadDisputeEvidenceRequest struct {
        DisputeId string `json:"disputeId,omitempty"`
        DisputeEvidenceType model.DisputeEvidenceType `json:"disputeEvidenceType,omitempty"` 
}

func NewAlipayDownloadDisputeEvidenceRequest() (*request.AlipayRequest, *AlipayDownloadDisputeEvidenceRequest) { 
alipayDownloadDisputeEvidenceRequest := &AlipayDownloadDisputeEvidenceRequest{} 
alipayRequest := request.NewAlipayRequest (alipayDownloadDisputeEvidenceRequest,  "null", &responseDispute.AlipayDownloadDisputeEvidenceResponse{}) 
return alipayRequest, alipayDownloadDisputeEvidenceRequest 
} 
 
func (alipayDownloadDisputeEvidenceRequest *AlipayDownloadDisputeEvidenceRequest) NewRequest() *request.AlipayRequest { 
return request.NewAlipayRequest(&alipayDownloadDisputeEvidenceRequest,  "null", &responseDispute.AlipayDownloadDisputeEvidenceResponse{}) 
}







