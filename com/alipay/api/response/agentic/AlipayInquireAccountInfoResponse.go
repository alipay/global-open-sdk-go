package responseAgentic

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireAccountInfoResponse struct {
	response.AlipayResponse
	Result         *model.Result                      `json:"result,omitempty"`
	PaymentMethods []*model.AgenticBoundPaymentMethod `json:"paymentMethods,omitempty"`
}
