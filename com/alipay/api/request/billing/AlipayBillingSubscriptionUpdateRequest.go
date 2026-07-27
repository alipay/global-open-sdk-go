package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionUpdateRequest struct {
	SubscriptionId       string                                          `json:"subscriptionId,omitempty"`
	PriceItemChanges     []*model.PriceItemChange                        `json:"priceItemChanges,omitempty"`
	ProrationBehavior    string                                          `json:"prorationBehavior,omitempty"`
	ProrationDate        string                                          `json:"prorationDate,omitempty"`
	PauseCollection      *model.BillingSubscriptionUpdatePauseCollection `json:"pauseCollection,omitempty"`
	BillingCycleAnchor   string                                          `json:"billingCycleAnchor,omitempty"`
	TrialSettings        *model.BillingSubscriptionUpdateTrialSettings   `json:"trialSettings,omitempty"`
	CancelAtPeriodEnd    bool                                            `json:"cancelAtPeriodEnd,omitempty"`
	CollectionMethod     string                                          `json:"collectionMethod,omitempty"`
	DaysUntilDue         int32                                           `json:"daysUntilDue,omitempty"`
	DefaultPaymentMethod string                                          `json:"defaultPaymentMethod,omitempty"`
	Description          string                                          `json:"description,omitempty"`
	Metadata             map[string]string                               `json:"metadata,omitempty"`
}

func NewAlipayBillingSubscriptionUpdateRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionUpdateRequest) {
	alipayBillingSubscriptionUpdateRequest := &AlipayBillingSubscriptionUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionUpdateRequest, "/ams/api/v1/billing/subscription/update", &responseBilling.AlipayBillingSubscriptionUpdateResponse{})
	return alipayRequest, alipayBillingSubscriptionUpdateRequest
}

func (alipayBillingSubscriptionUpdateRequest *AlipayBillingSubscriptionUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionUpdateRequest, "/ams/api/v1/billing/subscription/update", &responseBilling.AlipayBillingSubscriptionUpdateResponse{})
}
