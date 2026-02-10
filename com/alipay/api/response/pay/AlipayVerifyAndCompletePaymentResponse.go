package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayVerifyAndCompletePaymentResponse struct {
	response.AlipayResponse
	Result            *model.Result       `json:"result,omitempty"`
	PaymentRequestId  string              `json:"paymentRequestId,omitempty"`
	VerifyRequestId   string              `json:"verifyRequestId,omitempty"`
	PaymentId         string              `json:"paymentId,omitempty"`
	PaymentAmount     *model.Amount       `json:"paymentAmount,omitempty"`
	PaymentCreateTime string              `json:"paymentCreateTime,omitempty"`
	AcquirerInfo      *model.AcquirerInfo `json:"acquirerInfo,omitempty"`
}
