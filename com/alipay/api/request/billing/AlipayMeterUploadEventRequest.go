package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterUploadEventRequest struct {
	Meters []*model.Meter `json:"meters,omitempty"`
}

func NewAlipayMeterUploadEventRequest() (*request.AlipayRequest, *AlipayMeterUploadEventRequest) {
	alipayMeterUploadEventRequest := &AlipayMeterUploadEventRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterUploadEventRequest, "/ams/api/v1/meter/uploadEvent", &responseBilling.AlipayMeterUploadEventResponse{})
	return alipayRequest, alipayMeterUploadEventRequest
}

func (alipayMeterUploadEventRequest *AlipayMeterUploadEventRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterUploadEventRequest, "/ams/api/v1/meter/uploadEvent", &responseBilling.AlipayMeterUploadEventResponse{})
}
