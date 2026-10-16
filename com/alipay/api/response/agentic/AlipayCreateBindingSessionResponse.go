package responseAgentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateBindingSessionResponse struct {
	response.AlipayResponse
	Result                   *model.Result `json:"result,omitempty"`
	NormalUrl                string        `json:"normalUrl,omitempty"`
	SdHash                   string        `json:"sdHash,omitempty"`
	BindingSessionData       string        `json:"bindingSessionData,omitempty"`
	BindingSessionExpiryTime string        `json:"bindingSessionExpiryTime,omitempty"`
}
