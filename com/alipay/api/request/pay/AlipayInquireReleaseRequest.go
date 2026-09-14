package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayInquireReleaseRequest struct {
	ReleaseId        string `json:"releaseId,omitempty"`
	ReleaseRequestId string `json:"releaseRequestId,omitempty"`
}

func NewAlipayInquireReleaseRequest() (*request.AlipayRequest, *AlipayInquireReleaseRequest) {
	alipayInquireReleaseRequest := &AlipayInquireReleaseRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireReleaseRequest, "/ams/api/v1/payments/reserve/inquireRelease", &responsePay.AlipayInquireReleaseResponse{})
	return alipayRequest, alipayInquireReleaseRequest
}

func (alipayInquireReleaseRequest *AlipayInquireReleaseRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireReleaseRequest, "/ams/api/v1/payments/reserve/inquireRelease", &responsePay.AlipayInquireReleaseResponse{})
}
