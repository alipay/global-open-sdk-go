package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionCancelRequest struct {
	SubscriptionId      string                                              `json:"subscriptionId,omitempty"`
	CancellationType    model.CancellationType                              `json:"cancellationType,omitempty"`
	CancellationReason  string                                              `json:"cancellationReason,omitempty"`
	CancellationDetails *model.BillingSubscriptionCancelCancellationDetails `json:"cancellationDetails,omitempty"`
	ProrationBehavior   string                                              `json:"prorationBehavior,omitempty"`
}

func NewAlipayBillingSubscriptionCancelRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionCancelRequest) {
	alipayBillingSubscriptionCancelRequest := &AlipayBillingSubscriptionCancelRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionCancelRequest, "/ams/api/v1/billing/subscription/cancel", &responseBilling.AlipayBillingSubscriptionCancelResponse{})
	return alipayRequest, alipayBillingSubscriptionCancelRequest
}

func (alipayBillingSubscriptionCancelRequest *AlipayBillingSubscriptionCancelRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionCancelRequest, "/ams/api/v1/billing/subscription/cancel", &responseBilling.AlipayBillingSubscriptionCancelResponse{})
}
