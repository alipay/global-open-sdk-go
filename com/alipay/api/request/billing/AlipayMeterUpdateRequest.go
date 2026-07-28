package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterUpdateRequest struct {
	MeterId   string `json:"meterId,omitempty"`
	MeterName string `json:"meterName,omitempty"`
	Status    string `json:"status,omitempty"`
}

func NewAlipayMeterUpdateRequest() (*request.AlipayRequest, *AlipayMeterUpdateRequest) {
	alipayMeterUpdateRequest := &AlipayMeterUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterUpdateRequest, "/ams/api/v1/meter/update", &responseBilling.AlipayMeterUpdateResponse{})
	return alipayRequest, alipayMeterUpdateRequest
}

func (alipayMeterUpdateRequest *AlipayMeterUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterUpdateRequest, "/ams/api/v1/meter/update", &responseBilling.AlipayMeterUpdateResponse{})
}
