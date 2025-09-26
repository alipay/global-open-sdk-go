package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipaySubscriptionCreateRequest struct {
	SubscriptionRequestId       string                    `json:"subscriptionRequestId,omitempty"`
	SubscriptionDescription     string                    `json:"subscriptionDescription,omitempty"`
	SubscriptionRedirectUrl     string                    `json:"subscriptionRedirectUrl,omitempty"`
	SubscriptionStartTime       string                    `json:"subscriptionStartTime,omitempty"`
	SubscriptionEndTime         string                    `json:"subscriptionEndTime,omitempty"`
	PeriodRule                  *model.PeriodRule         `json:"periodRule,omitempty"`
	SubscriptionExpiryTime      string                    `json:"subscriptionExpiryTime,omitempty"`
	PaymentMethod               *model.PaymentMethod      `json:"paymentMethod,omitempty"`
	SubscriptionNotificationUrl string                    `json:"subscriptionNotificationUrl,omitempty"`
	PaymentNotificationUrl      string                    `json:"paymentNotificationUrl,omitempty"`
	OrderInfo                   *model.OrderInfo          `json:"orderInfo,omitempty"`
	PaymentAmount               *model.Amount             `json:"paymentAmount,omitempty"`
	SettlementStrategy          *model.SettlementStrategy `json:"settlementStrategy,omitempty"`
	Env                         *model.Env                `json:"env,omitempty"`
	Trials                      []*model.Trial            `json:"trials,omitempty"`
}

func NewAlipaySubscriptionCreateRequest() (*request.AlipayRequest, *AlipaySubscriptionCreateRequest) {
	alipaySubscriptionCreateRequest := &AlipaySubscriptionCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySubscriptionCreateRequest, "/ams/api/v1/subscriptions/create", &responseSubscription.AlipaySubscriptionCreateResponse{})
	return alipayRequest, alipaySubscriptionCreateRequest
}

func (alipaySubscriptionCreateRequest *AlipaySubscriptionCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySubscriptionCreateRequest, "/ams/api/v1/subscriptions/create", &responseSubscription.AlipaySubscriptionCreateResponse{})
}
