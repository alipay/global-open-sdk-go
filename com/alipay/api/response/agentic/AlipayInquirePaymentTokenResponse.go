package responseAgentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquirePaymentTokenResponse struct {
	response.AlipayResponse
	Result       *model.Result              `json:"result,omitempty"`
	PaymentToken *model.AgenticPaymentToken `json:"paymentToken,omitempty"`
}
