package responseSubscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySubscriptionsInquiryResponse struct {
	response.AlipayResponse
	AllowRetry            string            `json:"allowRetry,omitempty"`
	MaxAmountFloor        *model.Amount     `json:"maxAmountFloor,omitempty"`
	PaymentAmount         *model.Amount     `json:"paymentAmount,omitempty"`
	PeriodRule            *model.PeriodRule `json:"periodRule,omitempty"`
	SubscriptionEndTime   string            `json:"subscriptionEndTime,omitempty"`
	SubscriptionRequestId string            `json:"subscriptionRequestId,omitempty"`
	SubscriptionStartTime string            `json:"subscriptionStartTime,omitempty"`
	SubscriptionStatus    string            `json:"subscriptionStatus,omitempty"`
	Result                *model.Result     `json:"result,omitempty"`
}
