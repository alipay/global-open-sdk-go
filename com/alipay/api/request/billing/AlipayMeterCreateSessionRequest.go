package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterCreateSessionRequest struct {
	RequestId string `json:"requestId,omitempty"`
}

func NewAlipayMeterCreateSessionRequest() (*request.AlipayRequest, *AlipayMeterCreateSessionRequest) {
	alipayMeterCreateSessionRequest := &AlipayMeterCreateSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterCreateSessionRequest, "/ams/api/v1/meter/createSession", &responseBilling.AlipayMeterCreateSessionResponse{})
	return alipayRequest, alipayMeterCreateSessionRequest
}

func (alipayMeterCreateSessionRequest *AlipayMeterCreateSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterCreateSessionRequest, "/ams/api/v1/meter/createSession", &responseBilling.AlipayMeterCreateSessionResponse{})
}
