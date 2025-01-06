package auth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseAuth "github.com/alipay/global-open-sdk-go/com/alipay/api/response/auth"
)

type AlipayAuthCreateSessionRequest struct {
	ProductCode      model.ProductCodeType `json:"productCode,omitempty"`
	AgreementInfo    *model.AgreementInfo  `json:"agreementInfo,omitempty"`
	Scopes           []model.ScopeType     `json:"scopes,omitempty"`
	PaymentMethod    *model.PaymentMethod  `json:"paymentMethod,omitempty"`
	PaymentNotifyUrl string                `json:"paymentNotifyUrl,omitempty"`
}

func (alipayAuthCreateSessionRequest *AlipayAuthCreateSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAuthCreateSessionRequest, model.CREATE_SESSION_PATH, &responseAuth.AlipayAuthCreateSessionResponse{})
}

func NewAlipayAuthCreateSessionRequest() (*request.AlipayRequest, *AlipayAuthCreateSessionRequest) {
	alipayAuthCreateSessionRequest := &AlipayAuthCreateSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthCreateSessionRequest, model.CREATE_SESSION_PATH, &responseAuth.AlipayAuthCreateSessionResponse{})
	return alipayRequest, alipayAuthCreateSessionRequest
}
