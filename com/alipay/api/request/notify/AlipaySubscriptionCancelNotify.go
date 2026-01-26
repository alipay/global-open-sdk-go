package notify

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
)

type AlipaySubscriptionCancelNotify struct {
	AlipayNotify
	PaymentAmount           *model.Amount                      `json:"paymentAmount"`
	PaymentCreateTime       string                             `json:"paymentCreateTime"`
	PaymentTime             string                             `json:"paymentTime"`
	PeriodEndTime           string                             `json:"periodEndTime"`
	PhaseNo                 string                             `json:"phaseNo"`
	SubscriptionId          string                             `json:"subscriptionId"`
	SubscriptionOrderId     string                             `json:"subscriptionOrderId"`
	SubscriptionOrderStatus model.SubscriptionNotificationType `json:"subscriptionOrderStatus"`
	SubscriptionRequestId   string                             `json:"subscriptionRequestId"`
	SubscriptionTransId     string                             `json:"subscriptionTransId"`
}
