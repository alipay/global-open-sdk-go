package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipaySubscriptionUpdateRequest struct {
	SubscriptionUpdateRequestId string                   `json:"subscriptionUpdateRequestId,omitempty"`
	SubscriptionId              string                   `json:"subscriptionId,omitempty"`
	SubscriptionDescription     string                   `json:"subscriptionDescription,omitempty"`
	PeriodRule                  *model.PeriodRule        `json:"periodRule,omitempty"`
	PaymentAmount               *model.Amount            `json:"paymentAmount,omitempty"`
	SubscriptionEndTime         string                   `json:"subscriptionEndTime,omitempty"`
	OrderInfo                   *model.OrderInfo         `json:"orderInfo,omitempty"`
	ProrationSettings           *model.ProrationSettings `json:"prorationSettings,omitempty"`
	NextSubscriptionDate        string                   `json:"nextSubscriptionDate,omitempty"`
}

func NewAlipaySubscriptionUpdateRequest() (*request.AlipayRequest, *AlipaySubscriptionUpdateRequest) {
	alipaySubscriptionUpdateRequest := &AlipaySubscriptionUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySubscriptionUpdateRequest, "null", &responseSubscription.AlipaySubscriptionUpdateResponse{})
	return alipayRequest, alipaySubscriptionUpdateRequest
}

func (alipaySubscriptionUpdateRequest *AlipaySubscriptionUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySubscriptionUpdateRequest, "null", &responseSubscription.AlipaySubscriptionUpdateResponse{})
}
