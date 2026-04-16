package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipayInquireSubscriptionPaymentRequest struct {
	MerchantAccountId string                 `json:"merchantAccountId,omitempty"`
	SubscriptionId    string                 `json:"subscriptionId,omitempty"`
	PaymentStatuses   []*model.PaymentStatus `json:"paymentStatuses,omitempty"`
	CurrentPage       int32                  `json:"currentPage,omitempty"`
	PageSize          int32                  `json:"pageSize,omitempty"`
}

func NewAlipayInquireSubscriptionPaymentRequest() (*request.AlipayRequest, *AlipayInquireSubscriptionPaymentRequest) {
	alipayInquireSubscriptionPaymentRequest := &AlipayInquireSubscriptionPaymentRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireSubscriptionPaymentRequest, "/ams/api/v1/subscriptions/inquireSubscriptionPayment", &responseSubscription.AlipayInquireSubscriptionPaymentResponse{})
	return alipayRequest, alipayInquireSubscriptionPaymentRequest
}

func (alipayInquireSubscriptionPaymentRequest *AlipayInquireSubscriptionPaymentRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireSubscriptionPaymentRequest, "/ams/api/v1/subscriptions/inquireSubscriptionPayment", &responseSubscription.AlipayInquireSubscriptionPaymentResponse{})
}
