package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPromotionCodeInquireListResponse struct {
	response.AlipayResponse
	Result         *model.Result              `json:"result,omitempty"`
	PromotionCodes []*model.PromotionCodeInfo `json:"promotionCodes,omitempty"`
	HasMore        bool                       `json:"hasMore,omitempty"`
	NextCursor     string                     `json:"nextCursor,omitempty"`
	Total          int32                      `json:"total,omitempty"`
	PreviousCursor string                     `json:"previousCursor,omitempty"`
}
