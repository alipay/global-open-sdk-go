package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayCreateRuleRequest struct {
	FundsType          model.FundsType          `json:"fundsType,omitempty"`
	TakeType           model.TakeType           `json:"takeType,omitempty"`
	PaymentMethodScope model.PaymentMethodScope `json:"paymentMethodScope,omitempty"`
	Ratio              int32                    `json:"ratio,omitempty"`
	ReleaseType        model.ReleaseType        `json:"releaseType,omitempty"`
	ReleaseTime        string                   `json:"releaseTime,omitempty"`
	RetentionTime      int32                    `json:"retentionTime,omitempty"`
}

func NewAlipayCreateRuleRequest() (*request.AlipayRequest, *AlipayCreateRuleRequest) {
	alipayCreateRuleRequest := &AlipayCreateRuleRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateRuleRequest, "/ams/api/v1/payments/reserve/createRule", &responsePay.AlipayCreateRuleResponse{})
	return alipayRequest, alipayCreateRuleRequest
}

func (alipayCreateRuleRequest *AlipayCreateRuleRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateRuleRequest, "/ams/api/v1/payments/reserve/createRule", &responsePay.AlipayCreateRuleResponse{})
}
