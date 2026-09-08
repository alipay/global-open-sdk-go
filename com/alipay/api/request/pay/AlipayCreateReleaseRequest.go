package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayCreateReleaseRequest struct {
	ReleaseRequestId string        `json:"releaseRequestId,omitempty"`
	HoldId           string        `json:"holdId,omitempty"`
	ReleaseAmount    *model.Amount `json:"releaseAmount,omitempty"`
}

func NewAlipayCreateReleaseRequest() (*request.AlipayRequest, *AlipayCreateReleaseRequest) {
	alipayCreateReleaseRequest := &AlipayCreateReleaseRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateReleaseRequest, "/ams/api/v1/payments/reserve/createRelease", &responsePay.AlipayCreateReleaseResponse{})
	return alipayRequest, alipayCreateReleaseRequest
}

func (alipayCreateReleaseRequest *AlipayCreateReleaseRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateReleaseRequest, "/ams/api/v1/payments/reserve/createRelease", &responsePay.AlipayCreateReleaseResponse{})
}
