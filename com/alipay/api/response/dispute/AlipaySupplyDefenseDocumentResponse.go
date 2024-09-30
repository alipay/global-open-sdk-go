package responseDispute

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipaySupplyDefenseDocumentResponse struct {
	response.AlipayResponse
	DisputeId             string `json:"disputeId,omitempty"`
	DisputeResolutionTime string `json:"disputeResolutionTime,omitempty"`
}
