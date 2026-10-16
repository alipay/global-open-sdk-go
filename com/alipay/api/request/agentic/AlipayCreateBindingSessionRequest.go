package agentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAgentic "github.com/alipay/global-open-sdk-go/com/alipay/api/response/agentic"
)

type AlipayCreateBindingSessionRequest struct {
	BindingRequestId string                              `json:"bindingRequestId,omitempty"`
	UserLoginId      string                              `json:"userLoginId,omitempty"`
	Agent            *model.AgenticBindingSessionAgent   `json:"agent,omitempty"`
	AuthorizeOnBind  *bool                               `json:"authorizeOnBind,omitempty"`
	PaymentMethod    *model.AgenticPaymentMethod         `json:"paymentMethod,omitempty"`
	PurchaseType     string                              `json:"purchaseType,omitempty"`
	Mandate          *model.AgenticBindingSessionMandate `json:"mandate,omitempty"`
	CountryCode      string                              `json:"countryCode,omitempty"`
	RedirectUrl      string                              `json:"redirectUrl,omitempty"`
	Env              *model.AgenticSessionEnv            `json:"env,omitempty"`
}

func NewAlipayCreateBindingSessionRequest() (*request.AlipayRequest, *AlipayCreateBindingSessionRequest) {
	alipayCreateBindingSessionRequest := &AlipayCreateBindingSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateBindingSessionRequest, "/ams/api/v1/agenticPayment/createBindingSession", &responseAgentic.AlipayCreateBindingSessionResponse{})
	return alipayRequest, alipayCreateBindingSessionRequest
}

func (alipayCreateBindingSessionRequest *AlipayCreateBindingSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateBindingSessionRequest, "/ams/api/v1/agenticPayment/createBindingSession", &responseAgentic.AlipayCreateBindingSessionResponse{})
}
