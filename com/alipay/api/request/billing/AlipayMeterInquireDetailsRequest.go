package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterInquireDetailsRequest struct {
	MeterId string `json:"meterId,omitempty"`
}

func NewAlipayMeterInquireDetailsRequest() (*request.AlipayRequest, *AlipayMeterInquireDetailsRequest) {
	alipayMeterInquireDetailsRequest := &AlipayMeterInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterInquireDetailsRequest, "/ams/api/v1/meter/inquireDetails", &responseBilling.AlipayMeterInquireDetailsResponse{})
	return alipayRequest, alipayMeterInquireDetailsRequest
}

func (alipayMeterInquireDetailsRequest *AlipayMeterInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterInquireDetailsRequest, "/ams/api/v1/meter/inquireDetails", &responseBilling.AlipayMeterInquireDetailsResponse{})
}
