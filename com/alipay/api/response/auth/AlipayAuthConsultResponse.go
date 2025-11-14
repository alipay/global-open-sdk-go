package responseAuth

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayAuthConsultResponse struct {
	response.AlipayResponse
	Result        *model.Result       `json:"result,omitempty"`
	AuthUrl       string              `json:"authUrl,omitempty"`
	ExtendInfo    string              `json:"extendInfo,omitempty"`
	NormalUrl     string              `json:"normalUrl,omitempty"`
	SchemeUrl     string              `json:"schemeUrl,omitempty"`
	ApplinkUrl    string              `json:"applinkUrl,omitempty"`
	AppIdentifier string              `json:"appIdentifier,omitempty"`
	AuthCodeForm  *model.AuthCodeForm `json:"authCodeForm,omitempty"`
}
