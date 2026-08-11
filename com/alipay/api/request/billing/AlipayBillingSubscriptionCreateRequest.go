package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionCreateRequest struct {
	SubscriptionRequestId string                      `json:"subscriptionRequestId,omitempty"`
	CustomerId            string                      `json:"customerId,omitempty"`
	CustomerEmail         string                      `json:"customerEmail,omitempty"`
	PriceItems            []*model.PriceItem          `json:"priceItems,omitempty"`
	TrialSettings         *model.BillingTrialSettings `json:"trialSettings,omitempty"`
	Discounts             []*model.BillingDiscount    `json:"discounts,omitempty"`
	PaymentBehavior       string                      `json:"paymentBehavior,omitempty"`
	CollectionMethod      string                      `json:"collectionMethod,omitempty"`
	DaysUntilDue          int32                       `json:"daysUntilDue,omitempty"`
	CancelAt              string                      `json:"cancelAt,omitempty"`
	CancelAtPeriodEnd     bool                        `json:"cancelAtPeriodEnd,omitempty"`
	Description           string                      `json:"description,omitempty"`
	SubscriptionNotifyUrl string                      `json:"subscriptionNotifyUrl,omitempty"`
	Metadata              string                      `json:"metadata,omitempty"`
}

func NewAlipayBillingSubscriptionCreateRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionCreateRequest) {
	alipayBillingSubscriptionCreateRequest := &AlipayBillingSubscriptionCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionCreateRequest, "/ams/api/v1/billing/subscription/create", &responseBilling.AlipayBillingSubscriptionCreateResponse{})
	return alipayRequest, alipayBillingSubscriptionCreateRequest
}

func (alipayBillingSubscriptionCreateRequest *AlipayBillingSubscriptionCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionCreateRequest, "/ams/api/v1/billing/subscription/create", &responseBilling.AlipayBillingSubscriptionCreateResponse{})
}
