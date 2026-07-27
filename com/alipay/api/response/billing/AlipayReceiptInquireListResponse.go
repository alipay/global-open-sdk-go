package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayReceiptInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result    `json:"result,omitempty"`
	Receipts   []*model.Receipt `json:"receipts,omitempty"`
	Total      int32            `json:"total,omitempty"`
	HasMore    bool             `json:"hasMore,omitempty"`
	NextCursor string           `json:"nextCursor,omitempty"`
}
