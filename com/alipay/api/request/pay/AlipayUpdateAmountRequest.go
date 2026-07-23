package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayUpdateAmountRequest struct {
	UpdateRequestId string        `json:"updateRequestId,omitempty"`
	PaymentId       string        `json:"paymentId,omitempty"`
	Amount          *model.Amount `json:"amount,omitempty"`
}

func NewAlipayUpdateAmountRequest() (*request.AlipayRequest, *AlipayUpdateAmountRequest) {
	alipayUpdateAmountRequest := &AlipayUpdateAmountRequest{}
	alipayRequest := request.NewAlipayRequest(alipayUpdateAmountRequest, "/ams/api/v1/payments/updateAmount", &responsePay.AlipayUpdateAmountResponse{})
	return alipayRequest, alipayUpdateAmountRequest
}

func (alipayUpdateAmountRequest *AlipayUpdateAmountRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayUpdateAmountRequest, "/ams/api/v1/payments/updateAmount", &responsePay.AlipayUpdateAmountResponse{})
}
