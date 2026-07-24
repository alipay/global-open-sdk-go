package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCouponInquireListResponse struct {
	response.AlipayResponse
	Result     *model.ResultInfo `json:"result,omitempty"`
	Coupons    []*model.Coupon   `json:"coupons,omitempty"`
	HasMore    bool              `json:"hasMore,omitempty"`
	NextCursor string            `json:"nextCursor,omitempty"`
	PrevCursor string            `json:"prevCursor,omitempty"`
	Total      int32             `json:"total,omitempty"`
}
