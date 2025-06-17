package dispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseDispute "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute"
)

type AlipayDownloadDisputeEvidenceRequest struct {
	DisputeId           string                    `json:"disputeId,omitempty"`
	DisputeEvidenceType model.DisputeEvidenceType `json:"disputeEvidenceType,omitempty"`
}

func NewAlipayDownloadDisputeEvidenceRequest() (*request.AlipayRequest, *AlipayDownloadDisputeEvidenceRequest) {
	alipayDownloadDisputeEvidenceRequest := &AlipayDownloadDisputeEvidenceRequest{}
	alipayRequest := request.NewAlipayRequest(alipayDownloadDisputeEvidenceRequest, "/ams/api/v1/payments/downloadDisputeEvidence", &responseDispute.AlipayDownloadDisputeEvidenceResponse{})
	return alipayRequest, alipayDownloadDisputeEvidenceRequest
}

func (alipayDownloadDisputeEvidenceRequest *AlipayDownloadDisputeEvidenceRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayDownloadDisputeEvidenceRequest, "/ams/api/v1/payments/downloadDisputeEvidence", &responseDispute.AlipayDownloadDisputeEvidenceResponse{})
}
