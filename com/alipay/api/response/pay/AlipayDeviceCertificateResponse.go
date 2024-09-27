package responsePay

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipayDeviceCertificateResponse struct {
	response.AlipayResponse
	DeviceCertificate string `json:"deviceCertificate,omitempty"`
}
