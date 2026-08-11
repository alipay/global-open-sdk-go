package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPriceUpdateRequest struct {
	PriceId      string `json:"priceId,omitempty"`
	Name         string `json:"name,omitempty"`
	Metadata     string `json:"metadata,omitempty"`
	Active       bool   `json:"active,omitempty"`
	DefaultPrice bool   `json:"defaultPrice,omitempty"`
}

func NewAlipayPriceUpdateRequest() (*request.AlipayRequest, *AlipayPriceUpdateRequest) {
	alipayPriceUpdateRequest := &AlipayPriceUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPriceUpdateRequest, "/ams/api/v1/billing/price/update", &responseBilling.AlipayPriceUpdateResponse{})
	return alipayRequest, alipayPriceUpdateRequest
}

func (alipayPriceUpdateRequest *AlipayPriceUpdateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPriceUpdateRequest, "/ams/api/v1/billing/price/update", &responseBilling.AlipayPriceUpdateResponse{})
}
