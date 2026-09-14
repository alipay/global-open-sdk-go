package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayCreateHoldRequest struct {
	HoldRequestId string        `json:"holdRequestId,omitempty"`
	HoldAmount    *model.Amount `json:"holdAmount,omitempty"`
	ReleaseTime   string        `json:"releaseTime,omitempty"`
}

func NewAlipayCreateHoldRequest() (*request.AlipayRequest, *AlipayCreateHoldRequest) {
	alipayCreateHoldRequest := &AlipayCreateHoldRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateHoldRequest, "/ams/api/v1/payments/reserve/createHold", &responsePay.AlipayCreateHoldResponse{})
	return alipayRequest, alipayCreateHoldRequest
}

func (alipayCreateHoldRequest *AlipayCreateHoldRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateHoldRequest, "/ams/api/v1/payments/reserve/createHold", &responsePay.AlipayCreateHoldResponse{})
}
