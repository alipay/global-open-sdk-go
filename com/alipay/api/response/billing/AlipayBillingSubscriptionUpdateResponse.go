package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionUpdateResponse struct {
	response.AlipayResponse
	Result             *model.ResultInfo                               `json:"result,omitempty"`
	SubscriptionId     string                                          `json:"subscriptionId,omitempty"`
	Status             string                                          `json:"status,omitempty"`
	BillingCycleAnchor string                                          `json:"billingCycleAnchor,omitempty"`
	PauseCollection    *model.BillingSubscriptionUpdatePauseCollection `json:"pauseCollection,omitempty"`
	TrialSettings      *model.BillingSubscriptionUpdateTrialSettings   `json:"trialSettings,omitempty"`
	SubscriptionItems  []*model.SubscriptionItem                       `json:"subscriptionItems,omitempty"`
	ProrationInvoiceId string                                          `json:"prorationInvoiceId,omitempty"`
	CreditNoteId       string                                          `json:"creditNoteId,omitempty"`
	CancelAtPeriodEnd  bool                                            `json:"cancelAtPeriodEnd,omitempty"`
	CanceledAt         string                                          `json:"canceledAt,omitempty"`
	ProrationDate      string                                          `json:"prorationDate,omitempty"`
}
