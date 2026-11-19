package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquirePaymentMethodResponse struct {
	response.AlipayResponse
	Result               *model.Result                       `json:"result,omitempty"`
	PaymentMethodDetails []*model.InquirePaymentMethodDetail `json:"paymentMethodDetails,omitempty"`
}
