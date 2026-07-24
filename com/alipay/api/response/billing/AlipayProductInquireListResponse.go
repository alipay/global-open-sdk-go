package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayProductInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result    `json:"result,omitempty"`
	Products   []*model.Product `json:"products,omitempty"`
	HasMore    bool             `json:"hasMore,omitempty"`
	TotalCount int32            `json:"totalCount,omitempty"`
}
