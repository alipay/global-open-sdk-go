package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionInquireDetailsResponse struct {
	response.AlipayResponse
	Result                *model.ResultInfo                                           `json:"result,omitempty"`
	SubscriptionId        string                                                      `json:"subscriptionId,omitempty"`
	SubscriptionRequestId string                                                      `json:"subscriptionRequestId,omitempty"`
	CreatedAt             string                                                      `json:"createdAt,omitempty"`
	Status                string                                                      `json:"status,omitempty"`
	CurrentPeriodStart    string                                                      `json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd      string                                                      `json:"currentPeriodEnd,omitempty"`
	BillingCycleAnchor    string                                                      `json:"billingCycleAnchor,omitempty"`
	CancelAtPeriodEnd     bool                                                        `json:"cancelAtPeriodEnd,omitempty"`
	CanceledAt            string                                                      `json:"canceledAt,omitempty"`
	TrialStart            string                                                      `json:"trialStart,omitempty"`
	TrialEnd              string                                                      `json:"trialEnd,omitempty"`
	TrialSettings         *model.BillingSubscriptionInquireDetailsTrialSettings       `json:"trialSettings,omitempty"`
	PauseCollection       *model.BillingSubscriptionInquireDetailsPauseCollection     `json:"pauseCollection,omitempty"`
	CancelAt              string                                                      `json:"cancelAt,omitempty"`
	CollectionMethod      string                                                      `json:"collectionMethod,omitempty"`
	DaysUntilDue          int32                                                       `json:"daysUntilDue,omitempty"`
	CancellationDetails   *model.BillingSubscriptionInquireDetailsCancellationDetails `json:"cancellationDetails,omitempty"`
	TerminationReason     string                                                      `json:"terminationReason,omitempty"`
	Description           string                                                      `json:"description,omitempty"`
	DefaultPaymentMethod  string                                                      `json:"defaultPaymentMethod,omitempty"`
	SubscriptionItems     []*model.SubscriptionItem                                   `json:"subscriptionItems,omitempty"`
	Discounts             []*model.Discount                                           `json:"discounts,omitempty"`
	Metadata              map[string]string                                           `json:"metadata,omitempty"`
}
