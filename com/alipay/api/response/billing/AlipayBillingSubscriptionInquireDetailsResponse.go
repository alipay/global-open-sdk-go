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
	CreateTime            string                                                      `json:"createTime,omitempty"`
	Status                string                                                      `json:"status,omitempty"`
	BillingMode           string                                                      `json:"billingMode,omitempty"`
	PaymentBehavior       string                                                      `json:"paymentBehavior,omitempty"`
	CurrentPeriodStart    string                                                      `json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd      string                                                      `json:"currentPeriodEnd,omitempty"`
	BillingCycleAnchor    string                                                      `json:"billingCycleAnchor,omitempty"`
	CancelAtPeriodEnd     bool                                                        `json:"cancelAtPeriodEnd,omitempty"`
	CanceledAt            string                                                      `json:"canceledAt,omitempty"`
	TrialStart            string                                                      `json:"trialStart,omitempty"`
	TrialEnd              string                                                      `json:"trialEnd,omitempty"`
	TrialSettings         *model.BillingSubscriptionTrialSettings                     `json:"trialSettings,omitempty"`
	CancelAt              string                                                      `json:"cancelAt,omitempty"`
	CollectionMethod      string                                                      `json:"collectionMethod,omitempty"`
	DaysUntilDue          int32                                                       `json:"daysUntilDue,omitempty"`
	CancellationDetails   *model.BillingSubscriptionInquireDetailsCancellationDetails `json:"cancellationDetails,omitempty"`
	TerminationReason     string                                                      `json:"terminationReason,omitempty"`
	Description           string                                                      `json:"description,omitempty"`
	DefaultPaymentMethod  string                                                      `json:"defaultPaymentMethod,omitempty"`
	Subtotal              *model.Amount                                               `json:"subtotal,omitempty"`
	DiscountAmount        *model.Amount                                               `json:"discountAmount,omitempty"`
	TotalAmount           *model.Amount                                               `json:"totalAmount,omitempty"`
	PriceItems            []*model.BillingSubscriptionPriceItem                       `json:"priceItems,omitempty"`
	Discounts             []*model.BillingSubscriptionDiscountInfo                    `json:"discounts,omitempty"`
	Metadata              string                                                      `json:"metadata,omitempty"`
}
