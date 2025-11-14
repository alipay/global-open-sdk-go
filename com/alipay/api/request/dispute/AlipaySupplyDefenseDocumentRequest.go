package dispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseDispute "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute"
)

type AlipaySupplyDefenseDocumentRequest struct {
	DisputeId       string `json:"disputeId,omitempty"`
	DisputeEvidence string `json:"disputeEvidence,omitempty"`
}

func NewAlipaySupplyDefenseDocumentRequest() (*request.AlipayRequest, *AlipaySupplyDefenseDocumentRequest) {
	alipaySupplyDefenseDocumentRequest := &AlipaySupplyDefenseDocumentRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySupplyDefenseDocumentRequest, "/ams/api/v1/payments/supplyDefenseDocument", &responseDispute.AlipaySupplyDefenseDocumentResponse{})
	return alipayRequest, alipaySupplyDefenseDocumentRequest
}

func (alipaySupplyDefenseDocumentRequest *AlipaySupplyDefenseDocumentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySupplyDefenseDocumentRequest, "/ams/api/v1/payments/supplyDefenseDocument", &responseDispute.AlipaySupplyDefenseDocumentResponse{})
}
