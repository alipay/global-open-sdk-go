package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipaySubscriptionChangeRequest struct {
	SubscriptionChangeRequestId string            `json:"subscriptionChangeRequestId,omitempty"`
	SubscriptionId              string            `json:"subscriptionId,omitempty"`
	SubscriptionDescription     string            `json:"subscriptionDescription,omitempty"`
	SubscriptionStartTime       string            `json:"subscriptionStartTime,omitempty"`
	SubscriptionEndTime         string            `json:"subscriptionEndTime,omitempty"`
	PeriodRule                  *model.PeriodRule `json:"periodRule,omitempty"`
	SubscriptionExpiryTime      string            `json:"subscriptionExpiryTime,omitempty"`
	OrderInfo                   *model.OrderInfo  `json:"orderInfo,omitempty"`
	PaymentAmount               *model.Amount     `json:"paymentAmount,omitempty"`
	PaymentAmountDifference     *model.Amount     `json:"paymentAmountDifference,omitempty"`
}

func NewAlipaySubscriptionChangeRequest() (*request.AlipayRequest, *AlipaySubscriptionChangeRequest) {
	alipaySubscriptionChangeRequest := &AlipaySubscriptionChangeRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySubscriptionChangeRequest, "/ams/api/v1/subscriptions/change", &responseSubscription.AlipaySubscriptionChangeResponse{})
	return alipayRequest, alipaySubscriptionChangeRequest
}

func (alipaySubscriptionChangeRequest *AlipaySubscriptionChangeRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySubscriptionChangeRequest, "/ams/api/v1/subscriptions/change", &responseSubscription.AlipaySubscriptionChangeResponse{})
}
