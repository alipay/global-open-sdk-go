package responseSubscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySubscriptionCreateResponse struct {
	response.AlipayResponse
	Result        *model.Result `json:"result,omitempty"`
	SchemeUrl     string        `json:"schemeUrl,omitempty"`
	ApplinkUrl    string        `json:"applinkUrl,omitempty"`
	NormalUrl     string        `json:"normalUrl,omitempty"`
	AppIdentifier string        `json:"appIdentifier,omitempty"`
}
