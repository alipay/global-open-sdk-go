package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPriceInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result  `json:"result,omitempty"`
	Prices     []*model.Price `json:"prices,omitempty"`
	HasMore    bool           `json:"hasMore,omitempty"`
	TotalCount int32          `json:"totalCount,omitempty"`
}
