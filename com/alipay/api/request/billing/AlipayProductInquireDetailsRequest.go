package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayProductInquireDetailsRequest struct {
	ProductId     string `json:"productId,omitempty"`
	IncludePrices bool   `json:"includePrices,omitempty"`
}

func NewAlipayProductInquireDetailsRequest() (*request.AlipayRequest, *AlipayProductInquireDetailsRequest) {
	alipayProductInquireDetailsRequest := &AlipayProductInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayProductInquireDetailsRequest, "/ams/api/v1/billing/product/inquireDetails", &responseBilling.AlipayProductInquireDetailsResponse{})
	return alipayRequest, alipayProductInquireDetailsRequest
}

func (alipayProductInquireDetailsRequest *AlipayProductInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayProductInquireDetailsRequest, "/ams/api/v1/billing/product/inquireDetails", &responseBilling.AlipayProductInquireDetailsResponse{})
}
