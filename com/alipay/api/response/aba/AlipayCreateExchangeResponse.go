package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateExchangeResponse struct {
	response.AlipayResponse
	ExchangeRequestId string        `json:"exchangeRequestId,omitempty"`
	ExchangeTradeType string        `json:"exchangeTradeType,omitempty"`
	ExchangeMode      string        `json:"exchangeMode,omitempty"`
	Quote             *model.Quote  `json:"quote,omitempty"`
	BuyAmount         *model.Amount `json:"buyAmount,omitempty"`
	SellAmount        *model.Amount `json:"sellAmount,omitempty"`
	ExchangeStartTime string        `json:"exchangeStartTime,omitempty"`
	ExchangeEndTime   string        `json:"exchangeEndTime,omitempty"`
	Result            *model.Result `json:"result,omitempty"`
}
