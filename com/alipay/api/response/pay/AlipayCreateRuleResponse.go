package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateRuleResponse struct {
	response.AlipayResponse
	Result             *model.Result            `json:"result,omitempty"`
	RuleId             string                   `json:"ruleId,omitempty"`
	MerchantId         string                   `json:"merchantId,omitempty"`
	FundsType          model.FundsType          `json:"fundsType,omitempty"`
	TakeType           model.TakeType           `json:"takeType,omitempty"`
	PaymentMethodScope model.PaymentMethodScope `json:"paymentMethodScope,omitempty"`
	Ratio              int32                    `json:"ratio,omitempty"`
	ReleaseType        model.ReleaseType        `json:"releaseType,omitempty"`
	ReleaseTime        string                   `json:"releaseTime,omitempty"`
	RetentionTime      int32                    `json:"retentionTime,omitempty"`
	RuleStatus         model.RuleStatus         `json:"ruleStatus,omitempty"`
	CreateTime         string                   `json:"createTime,omitempty"`
	UpdateTime         string                   `json:"updateTime,omitempty"`
}
