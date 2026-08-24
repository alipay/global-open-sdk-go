package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayUpdateRuleRequest struct {
	RuleId             string                   `json:"ruleId,omitempty"`
	PaymentMethodScope model.PaymentMethodScope `json:"paymentMethodScope,omitempty"`
	Ratio              int32                    `json:"ratio,omitempty"`
	ReleaseTime        string                   `json:"releaseTime,omitempty"`
	RetentionTime      int32                    `json:"retentionTime,omitempty"`
	RuleStatus         model.RuleStatus         `json:"ruleStatus,omitempty"`
}

func NewAlipayUpdateRuleRequest() (*request.AlipayRequest, *AlipayUpdateRuleRequest) {
	alipayUpdateRuleRequest := &AlipayUpdateRuleRequest{}
	alipayRequest := request.NewAlipayRequest(alipayUpdateRuleRequest, "/ams/api/v1/payments/reserve/updateRule", &responsePay.AlipayUpdateRuleResponse{})
	return alipayRequest, alipayUpdateRuleRequest
}

func (alipayUpdateRuleRequest *AlipayUpdateRuleRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayUpdateRuleRequest, "/ams/api/v1/payments/reserve/updateRule", &responsePay.AlipayUpdateRuleResponse{})
}
