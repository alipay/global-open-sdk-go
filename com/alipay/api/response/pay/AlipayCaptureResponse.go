package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCaptureResponse struct {
	response.AlipayResponse
	Result              *model.Result `json:"result,omitempty"`
	CaptureRequestId    string        `json:"captureRequestId,omitempty"`
	CaptureId           string        `json:"captureId,omitempty"`
	PaymentId           string        `json:"paymentId,omitempty"`
	CaptureAmount       *model.Amount `json:"captureAmount,omitempty"`
	CaptureTime         string        `json:"captureTime,omitempty"`
	AcquirerReferenceNo string        `json:"acquirerReferenceNo,omitempty"`
}
