package responseDispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayDownloadDisputeEvidenceResponse struct {
	response.AlipayResponse
	DisputeEvidence       string                          `json:"disputeEvidence,omitempty"`
	DisputeEvidenceFormat model.DisputeEvidenceFormatType `json:"disputeEvidenceFormat,omitempty"`
}
