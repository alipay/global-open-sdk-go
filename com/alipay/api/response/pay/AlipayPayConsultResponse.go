package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPayConsultResponse struct {
	response.AlipayResponse
	Result             *model.Result              `json:"result,omitempty"`
	PaymentOptions     []*model.PaymentOption     `json:"paymentOptions,omitempty"`
	PaymentMethodInfos []*model.PaymentMethodInfo `json:"paymentMethodInfos,omitempty"`
	ExtendInfo         string                     `json:"extendInfo,omitempty"`
}
