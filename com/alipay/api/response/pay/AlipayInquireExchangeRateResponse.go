package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireExchangeRateResponse struct {
	response.AlipayResponse
	Quotes []model.Quote `json:"quotes,omitempty"`
}
