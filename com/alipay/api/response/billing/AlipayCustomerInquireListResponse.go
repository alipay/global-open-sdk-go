package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCustomerInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result     `json:"result,omitempty"`
	Customers  []*model.Customer `json:"customers,omitempty"`
	Total      int32             `json:"total,omitempty"`
	HasMore    bool              `json:"hasMore,omitempty"`
	NextCursor string            `json:"nextCursor,omitempty"`
}
