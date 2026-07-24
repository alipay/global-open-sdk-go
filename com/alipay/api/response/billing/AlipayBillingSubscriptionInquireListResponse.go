package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionInquireListResponse struct {
	response.AlipayResponse
	Result        *model.ResultInfo     `json:"result,omitempty"`
	Subscriptions []*model.Subscription `json:"subscriptions,omitempty"`
	HasMore       bool                  `json:"hasMore,omitempty"`
	NextCursor    string                `json:"nextCursor,omitempty"`
	PrevCursor    string                `json:"prevCursor,omitempty"`
	Total         int32                 `json:"total,omitempty"`
}
