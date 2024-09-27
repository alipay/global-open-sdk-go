package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCaptureResponse struct {
	response.AlipayResponse
	CaptureRequestId    string       `json:"captureRequestId"`
	CaptureId           string       `json:"captureId"`
	PaymentId           string       `json:"paymentId"`
	CaptureAmount       model.Amount `json:"captureAmount"`
	CaptureTime         string       `json:"captureTime"`
	AcquirerReferenceNo string       `json:"acquirerReferenceNo"`
}
