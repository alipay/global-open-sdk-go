package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayInquirePaymentMethodRequest struct {
	CardToken string `json:"cardToken,omitempty"`
}

func NewAlipayInquirePaymentMethodRequest() (*request.AlipayRequest, *AlipayInquirePaymentMethodRequest) {
	alipayInquirePaymentMethodRequest := &AlipayInquirePaymentMethodRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquirePaymentMethodRequest, "/ams/api/v1/paymentMethods/inquirePaymentMethod", &responsePay.AlipayInquirePaymentMethodResponse{})
	return alipayRequest, alipayInquirePaymentMethodRequest
}

func (alipayInquirePaymentMethodRequest *AlipayInquirePaymentMethodRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquirePaymentMethodRequest, "/ams/api/v1/paymentMethods/inquirePaymentMethod", &responsePay.AlipayInquirePaymentMethodResponse{})
}
