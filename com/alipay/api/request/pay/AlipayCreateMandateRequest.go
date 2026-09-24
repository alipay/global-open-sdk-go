package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayCreateMandateRequest struct {
	MandateRequestId string        `json:"mandateRequestId,omitempty"`
	AccessToken      string        `json:"accessToken,omitempty"`
	PaymentAmount    *model.Amount `json:"paymentAmount,omitempty"`
}

func NewAlipayCreateMandateRequest() (*request.AlipayRequest, *AlipayCreateMandateRequest) {
	alipayCreateMandateRequest := &AlipayCreateMandateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateMandateRequest, "/ams/api/v1/payments/createMandate", &responsePay.AlipayCreateMandateResponse{})
	return alipayRequest, alipayCreateMandateRequest
}

func (alipayCreateMandateRequest *AlipayCreateMandateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateMandateRequest, "/ams/api/v1/payments/createMandate", &responsePay.AlipayCreateMandateResponse{})
}
