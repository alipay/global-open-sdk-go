package responseDispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySupplyDefenseDocumentResponse struct {
	response.AlipayResponse
	Result                *model.Result `json:"result,omitempty"`
	DisputeId             string        `json:"disputeId,omitempty"`
	DisputeResolutionTime string        `json:"disputeResolutionTime,omitempty"`
}
