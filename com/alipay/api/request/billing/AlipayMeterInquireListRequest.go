package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterInquireListRequest struct {
	PageNum       int32  `json:"pageNum,omitempty"`
	PageSize      int32  `json:"pageSize,omitempty"`
	MeterName     string `json:"meterName,omitempty"`
	EventName     string `json:"eventName,omitempty"`
	Status        string `json:"status,omitempty"`
	StartDateTime string `json:"startDateTime,omitempty"`
	EndDateTime   string `json:"endDateTime,omitempty"`
}

func NewAlipayMeterInquireListRequest() (*request.AlipayRequest, *AlipayMeterInquireListRequest) {
	alipayMeterInquireListRequest := &AlipayMeterInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterInquireListRequest, "/ams/api/v1/meter/inquireList", &responseBilling.AlipayMeterInquireListResponse{})
	return alipayRequest, alipayMeterInquireListRequest
}

func (alipayMeterInquireListRequest *AlipayMeterInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterInquireListRequest, "/ams/api/v1/meter/inquireList", &responseBilling.AlipayMeterInquireListResponse{})
}
