package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireBalanceResponse struct {
	response.AlipayResponse
	Result          *model.Result           `json:"result,omitempty"`
	AccountBalances []*model.AccountBalance `json:"accountBalances,omitempty"`
}
