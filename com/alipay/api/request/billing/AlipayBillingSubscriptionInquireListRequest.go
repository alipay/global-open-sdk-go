package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayBillingSubscriptionInquireListRequest struct {
	Status         string `json:"status,omitempty"`
	CustomerId     string `json:"customerId,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	GmtCreateStart string `json:"gmtCreateStart,omitempty"`
	GmtCreateEnd   string `json:"gmtCreateEnd,omitempty"`
	SortOrder      string `json:"sortOrder,omitempty"`
	StartingAfter  string `json:"startingAfter,omitempty"`
	EndingBefore   string `json:"endingBefore,omitempty"`
	Limit          int32  `json:"limit,omitempty"`
}

func NewAlipayBillingSubscriptionInquireListRequest() (*request.AlipayRequest, *AlipayBillingSubscriptionInquireListRequest) {
	alipayBillingSubscriptionInquireListRequest := &AlipayBillingSubscriptionInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayBillingSubscriptionInquireListRequest, "/ams/api/v1/billing/subscription/inquireList", &responseBilling.AlipayBillingSubscriptionInquireListResponse{})
	return alipayRequest, alipayBillingSubscriptionInquireListRequest
}

func (alipayBillingSubscriptionInquireListRequest *AlipayBillingSubscriptionInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayBillingSubscriptionInquireListRequest, "/ams/api/v1/billing/subscription/inquireList", &responseBilling.AlipayBillingSubscriptionInquireListResponse{})
}
