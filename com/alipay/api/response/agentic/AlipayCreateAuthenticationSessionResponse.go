package responseAgentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateAuthenticationSessionResponse struct {
	response.AlipayResponse
	Result                          *model.Result `json:"result,omitempty"`
	NormalUrl                       string        `json:"normalUrl,omitempty"`
	SdHash                          string        `json:"sdHash,omitempty"`
	AuthenticationSessionExpiryTime string        `json:"authenticationSessionExpiryTime,omitempty"`
}
