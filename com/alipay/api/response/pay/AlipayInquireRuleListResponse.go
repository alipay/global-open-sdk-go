package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireRuleListResponse struct {
	response.AlipayResponse
	Result  *model.Result        `json:"result,omitempty"`
	HasMore bool                 `json:"hasMore,omitempty"`
	Rules   []*model.ReserveRule `json:"rules,omitempty"`
}
