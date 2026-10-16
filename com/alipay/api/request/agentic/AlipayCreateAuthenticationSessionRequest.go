package agentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAgentic "github.com/alipay/global-open-sdk-go/com/alipay/api/response/agentic"
)

type AlipayCreateAuthenticationSessionRequest struct {
	AuthenticationRequestId string                                     `json:"authenticationRequestId,omitempty"`
	UserLoginId             string                                     `json:"userLoginId,omitempty"`
	Agent                   *model.AgenticAuthenticationSessionAgent   `json:"agent,omitempty"`
	BindOnAuthorize         *bool                                      `json:"bindOnAuthorize,omitempty"`
	PaymentMethod           *model.AgenticPaymentMethod                `json:"paymentMethod,omitempty"`
	PurchaseType            string                                     `json:"purchaseType,omitempty"`
	Mandate                 *model.AgenticAuthenticationSessionMandate `json:"mandate,omitempty"`
	CountryCode             string                                     `json:"countryCode,omitempty"`
	RedirectUrl             string                                     `json:"redirectUrl,omitempty"`
	Env                     *model.AgenticSessionEnv                   `json:"env,omitempty"`
}

func NewAlipayCreateAuthenticationSessionRequest() (*request.AlipayRequest, *AlipayCreateAuthenticationSessionRequest) {
	alipayCreateAuthenticationSessionRequest := &AlipayCreateAuthenticationSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateAuthenticationSessionRequest, "/ams/api/v1/agenticPayment/createAuthenticationSession", &responseAgentic.AlipayCreateAuthenticationSessionResponse{})
	return alipayRequest, alipayCreateAuthenticationSessionRequest
}

func (alipayCreateAuthenticationSessionRequest *AlipayCreateAuthenticationSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateAuthenticationSessionRequest, "/ams/api/v1/agenticPayment/createAuthenticationSession", &responseAgentic.AlipayCreateAuthenticationSessionResponse{})
}
