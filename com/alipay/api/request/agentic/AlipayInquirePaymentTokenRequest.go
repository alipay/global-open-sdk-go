package agentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAgentic "github.com/alipay/global-open-sdk-go/com/alipay/api/response/agentic"
)

type AlipayInquirePaymentTokenRequest struct {
	PaymentTokenRequestId string                          `json:"paymentTokenRequestId,omitempty"`
	UserLoginId           string                          `json:"userLoginId,omitempty"`
	Agent                 *model.AgenticPaymentTokenAgent `json:"agent,omitempty"`
	PaymentMethod         *model.AgenticPaymentMethod     `json:"paymentMethod,omitempty"`
	MandateId             string                          `json:"mandateId,omitempty"`
	Metadata              string                          `json:"metadata,omitempty"`
	Order                 *model.AgenticPaymentTokenOrder `json:"order,omitempty"`
}

func NewAlipayInquirePaymentTokenRequest() (*request.AlipayRequest, *AlipayInquirePaymentTokenRequest) {
	alipayInquirePaymentTokenRequest := &AlipayInquirePaymentTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquirePaymentTokenRequest, "/ams/api/v1/agenticPayment/inquirePaymentToken", &responseAgentic.AlipayInquirePaymentTokenResponse{})
	return alipayRequest, alipayInquirePaymentTokenRequest
}

func (alipayInquirePaymentTokenRequest *AlipayInquirePaymentTokenRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquirePaymentTokenRequest, "/ams/api/v1/agenticPayment/inquirePaymentToken", &responseAgentic.AlipayInquirePaymentTokenResponse{})
}
