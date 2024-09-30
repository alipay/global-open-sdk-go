package responseMarketplace

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type AlipaySettlementInfoUpdateResponse struct {
	response.AlipayResponse
	UpdateStatus string `json:"updateStatus,omitempty"`
}
