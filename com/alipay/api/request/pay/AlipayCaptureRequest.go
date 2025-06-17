package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayCaptureRequest struct {
	CaptureRequestId string        `json:"captureRequestId,omitempty"`
	PaymentId        string        `json:"paymentId,omitempty"`
	CaptureAmount    *model.Amount `json:"captureAmount,omitempty"`
	IsLastCapture    bool          `json:"isLastCapture,omitempty"`
}

func NewAlipayCaptureRequest() (*request.AlipayRequest, *AlipayCaptureRequest) {
	alipayCaptureRequest := &AlipayCaptureRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCaptureRequest, "/ams/api/v1/payments/capture", &responsePay.AlipayCaptureResponse{})
	return alipayRequest, alipayCaptureRequest
}

func (alipayCaptureRequest *AlipayCaptureRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCaptureRequest, "/ams/api/v1/payments/capture", &responsePay.AlipayCaptureResponse{})
}
