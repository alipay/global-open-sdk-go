package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPriceInquireDetailsRequest struct {
	PriceId string `json:"priceId,omitempty"`
}

func NewAlipayPriceInquireDetailsRequest() (*request.AlipayRequest, *AlipayPriceInquireDetailsRequest) {
	alipayPriceInquireDetailsRequest := &AlipayPriceInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPriceInquireDetailsRequest, "/ams/api/v1/billing/price/inquireDetails", &responseBilling.AlipayPriceInquireDetailsResponse{})
	return alipayRequest, alipayPriceInquireDetailsRequest
}

func (alipayPriceInquireDetailsRequest *AlipayPriceInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPriceInquireDetailsRequest, "/ams/api/v1/billing/price/inquireDetails", &responseBilling.AlipayPriceInquireDetailsResponse{})
}
