package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipayInquireSubscriptionRequest struct {
	MerchantAccountId string `json:"merchantAccountId,omitempty"`
	SubscriptionId    string `json:"subscriptionId,omitempty"`
}

func NewAlipayInquireSubscriptionRequest() (*request.AlipayRequest, *AlipayInquireSubscriptionRequest) {
	alipayInquireSubscriptionRequest := &AlipayInquireSubscriptionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireSubscriptionRequest, "null", &responseSubscription.AlipayInquireSubscriptionResponse{})
	return alipayRequest, alipayInquireSubscriptionRequest
}

func (alipayInquireSubscriptionRequest *AlipayInquireSubscriptionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireSubscriptionRequest, "null", &responseSubscription.AlipayInquireSubscriptionResponse{})
}
