package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryBalanceResponse struct {
	response.AlipayResponse
	AccountBalances []*model.AccountBalance `json:"accountBalances,omitempty"`
	Result          *model.Result           `json:"result,omitempty"`
}
