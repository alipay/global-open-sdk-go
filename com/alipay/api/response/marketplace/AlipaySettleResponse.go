package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySettleResponse struct {
	response.AlipayResponse
	Result              *model.Result `json:"result,omitempty"`
	SettlementRequestId string        `json:"settlementRequestId,omitempty"`
	SettlementId        string        `json:"settlementId,omitempty"`
}
