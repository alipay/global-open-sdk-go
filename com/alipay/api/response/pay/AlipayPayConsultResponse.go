package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPayConsultResponse struct {
	response.AlipayResponse
	PaymentOptions     []model.PaymentOption     `json:"paymentOptions"`
	PaymentMethodInfos []model.PaymentMethodInfo `json:"paymentMethodInfos"`
	ExtendInfo         string                    `json:"extendInfo"`
}
