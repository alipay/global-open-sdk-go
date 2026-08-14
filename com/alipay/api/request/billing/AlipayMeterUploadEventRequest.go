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
// X-Session-Id from meter/createSession. Applications using this API must be
// built with Go 1.25.13+ in the Go 1.25 series, Go 1.26.6+ in the Go 1.26
// series, or a later stable Go release. This requirement applies only to
// meter/uploadEvent; other SDK APIs remain compatible with the Go version
// declared in go.mod.
func NewAlipayMeterUploadEventRequest() (*request.AlipayRequest, *AlipayMeterUploadEventRequest) {
	alipayMeterUploadEventRequest := &AlipayMeterUploadEventRequest{}
	alipayRequest := request.NewAlipayRequest(alipayMeterUploadEventRequest, "/ams/api/v1/meter/uploadEvent", &responseBilling.AlipayMeterUploadEventResponse{})
	return alipayRequest, alipayMeterUploadEventRequest
}

// NewRequest creates a meter uploadEvent request from the current parameters.
// See NewAlipayMeterUploadEventRequest for the required Go runtime version.
func (alipayMeterUploadEventRequest *AlipayMeterUploadEventRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayMeterUploadEventRequest, "/ams/api/v1/meter/uploadEvent", &responseBilling.AlipayMeterUploadEventResponse{})
}
