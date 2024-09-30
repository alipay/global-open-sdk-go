package dispute

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseDispute "github.com/alipay/global-open-sdk-go/com/alipay/api/response/dispute"
)

type AlipayAcceptDisputeRequest struct {
	DisputeId string `json:"disputeId,omitempty"`
}

func NewAlipayAcceptDisputeRequest() (*request.AlipayRequest, *AlipayAcceptDisputeRequest) {
	alipayAcceptDisputeRequest := &AlipayAcceptDisputeRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAcceptDisputeRequest, model.ACCEPT_DISPUTE_PATH, &responseDispute.AlipayAcceptDisputeResponse{})
	return alipayRequest, alipayAcceptDisputeRequest
}
