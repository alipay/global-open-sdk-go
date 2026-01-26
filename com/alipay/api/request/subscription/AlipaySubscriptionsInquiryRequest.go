package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipaySubscriptionsInquiryRequest struct {
	SubscriptionId        string `json:"subscriptionId,omitempty"`
	SubscriptionRequestId string `json:"subscriptionRequestId,omitempty"`
}

func NewAlipaySubscriptionsInquiryRequest() (*request.AlipayRequest, *AlipaySubscriptionsInquiryRequest) {
	alipaySubscriptionsInquiryRequest := &AlipaySubscriptionsInquiryRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySubscriptionsInquiryRequest, "/ams/v1/subscriptions/inquiry", &responseSubscription.AlipaySubscriptionsInquiryResponse{})
	return alipayRequest, alipaySubscriptionsInquiryRequest
}

func (alipaySubscriptionsInquiryRequest *AlipaySubscriptionsInquiryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySubscriptionsInquiryRequest, "/ams/v1/subscriptions/inquiry", &responseSubscription.AlipaySubscriptionsInquiryResponse{})
}
