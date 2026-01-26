package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateQuoteResponse struct {
	response.AlipayResponse
	ExchangeTradeType string        `json:"exchangeTradeType,omitempty"`
	Quote             *model.Quote  `json:"quote,omitempty"`
	SellAmount        *model.Amount `json:"sellAmount,omitempty"`
	BuyAmount         *model.Amount `json:"buyAmount,omitempty"`
	Result            *model.Result `json:"result,omitempty"`
}
