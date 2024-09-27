package responseSubscription

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipaySubscriptionCreateResponse struct {
	response.AlipayResponse
	SchemeUrl     string `json:"schemeUrl"`
	ApplinkUrl    string `json:"applinkUrl"`
	NormalUrl     string `json:"normalUrl"`
	AppIdentifier string `json:"appIdentifier"`
}
