package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayMeterUploadEventRequest struct {
	Meters []*model.MeterEventBatch `json:"meters,omitempty"`
}

// NewAlipayMeterUploadEventRequest creates a meter uploadEvent request.
//
// Execute this request with DefaultAlipayClient.ExecuteWithHeaders and provide
// X-Session-Id from meter/createSession. This API supports the Go version
// declared in go.mod. Production applications should use a currently
// supported Go release containing the latest security fixes.
func NewAlipayMeterUploadEventRequest() (*request.AlipayRequest, *AlipayMeterUploadEventRequest) {
	alipayMeterUploadEventRequest := &AlipayMeterUploadEventRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterUploadEventRequest, "/ams/api/v1/meter/uploadEvent", &responseBilling.AlipayMeterUploadEventResponse{})
	return alipayRequest, alipayMeterUploadEventRequest
}

// NewRequest creates a meter uploadEvent request from the current parameters.
// See NewAlipayMeterUploadEventRequest for runtime guidance.
func (alipayMeterUploadEventRequest *AlipayMeterUploadEventRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterUploadEventRequest, "/ams/api/v1/meter/uploadEvent", &responseBilling.AlipayMeterUploadEventResponse{})
}
