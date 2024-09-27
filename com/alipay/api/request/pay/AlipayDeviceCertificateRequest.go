package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayDeviceCertificateRequest struct {
	DevicePublicKey string `json:"devicePublicKey,omitempty"`
	DeviceRequestId string `json:"deviceRequestId,omitempty"`
}

func (alipayCaptureRequest *AlipayDeviceCertificateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCaptureRequest, model.CREATE_DEVICE_CERTIFICATE_PATH, &responsePay.AlipayDeviceCertificateResponse{})
}

func NewAlipayDeviceCertificateRequest() (*request.AlipayRequest, *AlipayDeviceCertificateRequest) {
	alipayDeviceCertificateRequest := &AlipayDeviceCertificateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayDeviceCertificateRequest, model.CREATE_DEVICE_CERTIFICATE_PATH, &responsePay.AlipayDeviceCertificateResponse{})
	return alipayRequest, alipayDeviceCertificateRequest
}
