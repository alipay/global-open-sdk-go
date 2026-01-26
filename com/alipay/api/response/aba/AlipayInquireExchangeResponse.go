package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireExchangeResponse struct {
	response.AlipayResponse
	Result            *model.Result `json:"result,omitempty"`
	ExchangeResult    *model.Result `json:"exchangeResult,omitempty"`
	ExchangeTradeType string        `json:"exchangeTradeType,omitempty"`
	ExchangeMode      string        `json:"exchangeMode,omitempty"`
	ExchangeRequestId string        `json:"exchangeRequestId,omitempty"`
	Quote             *model.Quote  `json:"quote,omitempty"`
	SellAmount        *model.Amount `json:"sellAmount,omitempty"`
	BuyAmount         *model.Amount `json:"buyAmount,omitempty"`
	ExchangeId        string        `json:"exchangeId,omitempty"`
	ExchangeStartTime string        `json:"exchangeStartTime,omitempty"`
	ExchangeEndTime   string        `json:"exchangeEndTime,omitempty"`
}
