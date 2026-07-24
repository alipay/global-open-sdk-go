package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionCreateResponse struct {
	response.AlipayResponse
	Result                *model.ResultInfo         `json:"result,omitempty"`
	SubscriptionRequestId string                    `json:"subscriptionRequestId,omitempty"`
	SubscriptionId        string                    `json:"subscriptionId,omitempty"`
	CustomerId            string                    `json:"customerId,omitempty"`
	InvoiceId             string                    `json:"invoiceId,omitempty"`
	Status                string                    `json:"status,omitempty"`
	CurrentPeriodStart    string                    `json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd      string                    `json:"currentPeriodEnd,omitempty"`
	BillingCycleAnchor    string                    `json:"billingCycleAnchor,omitempty"`
	TrialStart            string                    `json:"trialStart,omitempty"`
	TrialEnd              string                    `json:"trialEnd,omitempty"`
	CancelAt              string                    `json:"cancelAt,omitempty"`
	CancelAtPeriodEnd     bool                      `json:"cancelAtPeriodEnd,omitempty"`
	Description           string                    `json:"description,omitempty"`
	CollectionMethod      string                    `json:"collectionMethod,omitempty"`
	DaysUntilDue          int32                     `json:"daysUntilDue,omitempty"`
	SubscriptionItems     []*model.SubscriptionItem `json:"subscriptionItems,omitempty"`
	Discounts             []*model.Discount         `json:"discounts,omitempty"`
	SubscriptionNotifyUrl string                    `json:"subscriptionNotifyUrl,omitempty"`
}
