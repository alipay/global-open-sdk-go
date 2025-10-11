package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayInquireExchangeRateRequest struct {
	MerchantAccountId string                `json:"merchantAccountId,omitempty"`
	PaymentCurrency   string                `json:"paymentCurrency,omitempty"`
	CurrencyPairs     []*model.CurrencyPair `json:"currencyPairs,omitempty"`
	SellCurrency      string                `json:"sellCurrency,omitempty"`
	BuyCurrency       string                `json:"buyCurrency,omitempty"`
}

func (alipayInquireExchangeRateRequest *AlipayInquireExchangeRateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireExchangeRateRequest, model.PAYMENT_INQUIRE_EXCHANGE_RATE_PATH, &responsePay.AlipayInquireExchangeRateResponse{})
}

func NewAlipayInquireExchangeRateRequest() (*request.AlipayRequest, *AlipayInquireExchangeRateRequest) {
	alipayInquireExchangeRateRequest := &AlipayInquireExchangeRateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireExchangeRateRequest, model.PAYMENT_INQUIRE_EXCHANGE_RATE_PATH, &responsePay.AlipayInquireExchangeRateResponse{})
	return alipayRequest, alipayInquireExchangeRateRequest
}
