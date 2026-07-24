package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterCreateRequest struct {
	MeterName         string `json:"meterName,omitempty"`
	EventName         string `json:"eventName,omitempty"`
	AggregationMethod string `json:"aggregationMethod,omitempty"`
	EventTimeWindow   string `json:"eventTimeWindow,omitempty"`
	ValueKeyOverride  string `json:"valueKeyOverride,omitempty"`
}

func NewAlipayMeterCreateRequest() (*request.AlipayRequest, *AlipayMeterCreateRequest) {
	alipayMeterCreateRequest := &AlipayMeterCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterCreateRequest, "/ams/api/v1/meter/create", &responseBilling.AlipayMeterCreateResponse{})
	return alipayRequest, alipayMeterCreateRequest
}

func (alipayMeterCreateRequest *AlipayMeterCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterCreateRequest, "/ams/api/v1/meter/create", &responseBilling.AlipayMeterCreateResponse{})
}
