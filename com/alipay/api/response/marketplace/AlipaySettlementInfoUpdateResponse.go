package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySettlementInfoUpdateResponse struct {
	response.AlipayResponse
	Result       *model.Result `json:"result,omitempty"`
	UpdateStatus string        `json:"updateStatus,omitempty"`
}
