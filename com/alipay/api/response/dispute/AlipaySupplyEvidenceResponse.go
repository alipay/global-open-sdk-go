package responseDispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySupplyEvidenceResponse struct {
	response.AlipayResponse
	Result          *model.Result                  `json:"result,omitempty"`
	DisputeId       string                         `json:"disputeId,omitempty"`
	EvidenceResults []*model.DisputeEvidenceResult `json:"evidenceResults,omitempty"`
}
