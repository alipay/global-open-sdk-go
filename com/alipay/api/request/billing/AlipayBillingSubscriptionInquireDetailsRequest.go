package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionInquireDetailsRequest struct {
	SubscriptionId string `json:"subscriptionId,omitempty"`
}

func NewAlipayBillingSubscriptionInquireDetailsRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionInquireDetailsRequest) {
	alipayBillingSubscriptionInquireDetailsRequest := &AlipayBillingSubscriptionInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionInquireDetailsRequest, "/ams/api/v1/billing/subscription/inquireDetails", &responseBilling.AlipayBillingSubscriptionInquireDetailsResponse{})
	return alipayRequest, alipayBillingSubscriptionInquireDetailsRequest
}

func (alipayBillingSubscriptionInquireDetailsRequest *AlipayBillingSubscriptionInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionInquireDetailsRequest, "/ams/api/v1/billing/subscription/inquireDetails", &responseBilling.AlipayBillingSubscriptionInquireDetailsResponse{})
}
