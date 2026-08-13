package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionResumeRequest struct {
	SubscriptionId string `json:"subscriptionId,omitempty"`
	ReasonCode     string `json:"reasonCode,omitempty"`
}

func NewAlipayBillingSubscriptionResumeRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionResumeRequest) {
	alipayBillingSubscriptionResumeRequest := &AlipayBillingSubscriptionResumeRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionResumeRequest, "/ams/api/v1/billing/subscription/resume", &responseBilling.AlipayBillingSubscriptionResumeResponse{})
	return alipayRequest, alipayBillingSubscriptionResumeRequest
}

func (alipayBillingSubscriptionResumeRequest *AlipayBillingSubscriptionResumeRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionResumeRequest, "/ams/api/v1/billing/subscription/resume", &responseBilling.AlipayBillingSubscriptionResumeResponse{})
}
