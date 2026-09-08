package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayUpdateHoldRequest struct {
	HoldId      string `json:"holdId,omitempty"`
	ReleaseTime string `json:"releaseTime,omitempty"`
}

func NewAlipayUpdateHoldRequest() (*request.AlipayRequest, *AlipayUpdateHoldRequest) {
	alipayUpdateHoldRequest := &AlipayUpdateHoldRequest{}
	alipayRequest := request.NewAlipayRequest(alipayUpdateHoldRequest, "/ams/api/v1/payments/reserve/updateHold", &responsePay.AlipayUpdateHoldResponse{})
	return alipayRequest, alipayUpdateHoldRequest
}

func (alipayUpdateHoldRequest *AlipayUpdateHoldRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayUpdateHoldRequest, "/ams/api/v1/payments/reserve/updateHold", &responsePay.AlipayUpdateHoldResponse{})
}
