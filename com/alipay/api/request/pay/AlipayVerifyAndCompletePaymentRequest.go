package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayVerifyAndCompletePaymentRequest struct {
	VerifyMethod    *model.VerifyMethod `json:"verifyMethod,omitempty"`
	VerifyRequestId string              `json:"verifyRequestId,omitempty"`
	PaymentId       string              `json:"paymentId,omitempty"`
}

func NewAlipayVerifyAndCompletePaymentRequest() (*request.AlipayRequest, *AlipayVerifyAndCompletePaymentRequest) {
	alipayVerifyAndCompletePaymentRequest := &AlipayVerifyAndCompletePaymentRequest{}
	alipayRequest := request.NewAlipayRequest(alipayVerifyAndCompletePaymentRequest, "/ams/api/v1/payments/verifyAndCompletePayment", &responsePay.AlipayVerifyAndCompletePaymentResponse{})
	return alipayRequest, alipayVerifyAndCompletePaymentRequest
}

func (alipayVerifyAndCompletePaymentRequest *AlipayVerifyAndCompletePaymentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayVerifyAndCompletePaymentRequest, "/ams/api/v1/payments/verifyAndCompletePayment", &responsePay.AlipayVerifyAndCompletePaymentResponse{})
}
