package aba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAba "github.com/alipay/global-open-sdk-go/com/alipay/api/response/aba"
)

type AlipayCreateDirectPaymentRequest struct {
	PaymentRequestId string               `json:"paymentRequestId,omitempty"`
	Memo             string               `json:"memo,omitempty"`
	Remark           string               `json:"remark,omitempty"`
	Order            *model.Order         `json:"order,omitempty"`
	PaymentNotifyUrl string               `json:"paymentNotifyUrl,omitempty"`
	PayToMethod      *model.PaymentMethod `json:"payToMethod,omitempty"`
	PayToAmount      *model.Amount        `json:"payToAmount,omitempty"`
	PayFromAmount    *model.Amount        `json:"payFromAmount,omitempty"`
}

func NewAlipayCreateDirectPaymentRequest() (*request.AlipayRequest, *AlipayCreateDirectPaymentRequest) {
	alipayCreateDirectPaymentRequest := &AlipayCreateDirectPaymentRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateDirectPaymentRequest, "/ams/api/v1/aba/funds/createDirectPayment", &responseAba.AlipayCreateDirectPaymentResponse{})
	return alipayRequest, alipayCreateDirectPaymentRequest
}

func (alipayCreateDirectPaymentRequest *AlipayCreateDirectPaymentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateDirectPaymentRequest, "/ams/api/v1/aba/funds/createDirectPayment", &responseAba.AlipayCreateDirectPaymentResponse{})
}
