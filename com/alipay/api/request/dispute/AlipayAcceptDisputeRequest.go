package dispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseDispute "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute"
)

type AlipayAcceptDisputeRequest struct {
	DisputeId string `json:"disputeId,omitempty"`
}

func NewAlipayAcceptDisputeRequest() (*request.AlipayRequest, *AlipayAcceptDisputeRequest) {
	alipayAcceptDisputeRequest := &AlipayAcceptDisputeRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAcceptDisputeRequest, "/ams/api/v1/payments/acceptDispute", &responseDispute.AlipayAcceptDisputeResponse{})
	return alipayRequest, alipayAcceptDisputeRequest
}

func (alipayAcceptDisputeRequest *AlipayAcceptDisputeRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAcceptDisputeRequest, "/ams/api/v1/payments/acceptDispute", &responseDispute.AlipayAcceptDisputeResponse{})
}
