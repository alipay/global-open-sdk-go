package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionUpdateRequest struct {
	SubscriptionId          string                                        `json:"subscriptionId,omitempty"`
	PriceItemChanges        []*model.PriceItemChange                      `json:"priceItemChanges,omitempty"`
	ProrationBehavior       string                                        `json:"prorationBehavior,omitempty"`
	ResetBillingCycleAnchor bool                                          `json:"resetBillingCycleAnchor,omitempty"`
	TrialSettings           *model.BillingTrialSettings                   `json:"trialSettings,omitempty"`
	CancelAtPeriodEnd       bool                                          `json:"cancelAtPeriodEnd,omitempty"`
	CancelAt                string                                        `json:"cancelAt,omitempty"`
	CancellationDetails     *model.BillingSubscriptionCancellationDetails `json:"cancellationDetails,omitempty"`
	CollectionMethod        string                                        `json:"collectionMethod,omitempty"`
	DaysUntilDue            int32                                         `json:"daysUntilDue,omitempty"`
	Description             string                                        `json:"description,omitempty"`
	Metadata                string                                        `json:"metadata,omitempty"`
}

func NewAlipayBillingSubscriptionUpdateRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionUpdateRequest) {
	alipayBillingSubscriptionUpdateRequest := &AlipayBillingSubscriptionUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionUpdateRequest, "/ams/api/v1/billing/subscription/update", &responseBilling.AlipayBillingSubscriptionUpdateResponse{})
	return alipayRequest, alipayBillingSubscriptionUpdateRequest
}

func (alipayBillingSubscriptionUpdateRequest *AlipayBillingSubscriptionUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionUpdateRequest, "/ams/api/v1/billing/subscription/update", &responseBilling.AlipayBillingSubscriptionUpdateResponse{})
}
