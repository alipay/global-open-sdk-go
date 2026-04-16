package subscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseSubscription "github.com/alipay/global-open-sdk-go/com/alipay/api/response/subscription"
)

type AlipayInquireSubscriptionListRequest struct {
	MerchantAccountId  string                      `json:"merchantAccountId,omitempty"`
	StartTimeFrom      string                      `json:"startTimeFrom,omitempty"`
	StartTimeTo        string                      `json:"startTimeTo,omitempty"`
	Statuses           []*model.SubscriptionStatus `json:"statuses,omitempty"`
	PaymentMethodTypes []string                    `json:"paymentMethodTypes,omitempty"`
	Currencies         []string                    `json:"currencies,omitempty"`
	PeriodTypes        []model.PeriodType          `json:"periodTypes,omitempty"`
	CurrentPage        int32                       `json:"currentPage,omitempty"`
	PageSize           int32                       `json:"pageSize,omitempty"`
}

func NewAlipayInquireSubscriptionListRequest() (*request.AlipayRequest, *AlipayInquireSubscriptionListRequest) {
	alipayInquireSubscriptionListRequest := &AlipayInquireSubscriptionListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireSubscriptionListRequest, "/ams/api/v1/subscriptions/inquireSubscriptionList", &responseSubscription.AlipayInquireSubscriptionListResponse{})
	return alipayRequest, alipayInquireSubscriptionListRequest
}

func (alipayInquireSubscriptionListRequest *AlipayInquireSubscriptionListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireSubscriptionListRequest, "/ams/api/v1/subscriptions/inquireSubscriptionList", &responseSubscription.AlipayInquireSubscriptionListResponse{})
}
