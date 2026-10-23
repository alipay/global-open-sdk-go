package dispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseDispute "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute"
)

type AlipaySupplyEvidenceRequest struct {
	DisputeId       string                            `json:"disputeId,omitempty"`
	EvidenceContent map[string]map[string]interface{} `json:"evidenceContent,omitempty"`
}

func NewAlipaySupplyEvidenceRequest() (*request.AlipayRequest, *AlipaySupplyEvidenceRequest) {
	alipaySupplyEvidenceRequest := &AlipaySupplyEvidenceRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySupplyEvidenceRequest, "/ams/api/v1/payments/supplyEvidence", &responseDispute.AlipaySupplyEvidenceResponse{})
	return alipayRequest, alipaySupplyEvidenceRequest
}

func (alipaySupplyEvidenceRequest *AlipaySupplyEvidenceRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySupplyEvidenceRequest, "/ams/api/v1/payments/supplyEvidence", &responseDispute.AlipaySupplyEvidenceResponse{})
}
