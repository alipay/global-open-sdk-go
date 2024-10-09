package responseRisk

import "github.com/alipay/global-open-sdk-go/com/alipay/api/response"

type RiskDecideResponse struct {
	response.AlipayResponse
	Decision               string `json:"decision,omitempty"`
	AuthenticationDecision string `json:"authenticationDecision,omitempty"`
}
