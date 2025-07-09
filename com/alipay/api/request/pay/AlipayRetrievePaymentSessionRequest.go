package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayRetrievePaymentSessionRequest struct {
	PaymentRequestId string `json:"paymentRequestId,omitempty"`
}

func (alipayRetrievePaymentSessionRequest *AlipayRetrievePaymentSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayRetrievePaymentSessionRequest, model.RETRIEVE_PATH, &responsePay.AlipayRetrievePaymentSessionResponse{})
}

func NewAlipayRetrievePaymentSessionRequest() (*request.AlipayRequest, *AlipayRetrievePaymentSessionRequest) {
	alipayRetrievePaymentSessionRequest := &AlipayRetrievePaymentSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayRetrievePaymentSessionRequest, model.RETRIEVE_PATH, &responsePay.AlipayRetrievePaymentSessionResponse{})
	return alipayRequest, alipayRetrievePaymentSessionRequest
}
