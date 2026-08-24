package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayCancelRuleRequest struct {
	RuleId string `json:"ruleId,omitempty"`
}

func NewAlipayCancelRuleRequest() (*request.AlipayRequest, *AlipayCancelRuleRequest) {
	alipayCancelRuleRequest := &AlipayCancelRuleRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCancelRuleRequest, "/ams/api/v1/payments/reserve/cancelRule", &responsePay.AlipayCancelRuleResponse{})
	return alipayRequest, alipayCancelRuleRequest
}

func (alipayCancelRuleRequest *AlipayCancelRuleRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCancelRuleRequest, "/ams/api/v1/payments/reserve/cancelRule", &responsePay.AlipayCancelRuleResponse{})
}
