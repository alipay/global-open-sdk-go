package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayPriceInquireListRequest struct {
	ProductId     string `json:"productId,omitempty"`
	PricingModel  string `json:"pricingModel,omitempty"`
	Active        bool   `json:"active,omitempty"`
	StartingAfter string `json:"startingAfter,omitempty"`
	EndingBefore  string `json:"endingBefore,omitempty"`
	List          int32  `json:"list,omitempty"`
	IncludeTotal  bool   `json:"includeTotal,omitempty"`
}

func NewAlipayPriceInquireListRequest() (*request.AlipayRequest, *AlipayPriceInquireListRequest) {
	alipayPriceInquireListRequest := &AlipayPriceInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPriceInquireListRequest, "/ams/api/v1/billing/price/inquireList", &responseBilling.AlipayPriceInquireListResponse{})
	return alipayRequest, alipayPriceInquireListRequest
}

func (alipayPriceInquireListRequest *AlipayPriceInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPriceInquireListRequest, "/ams/api/v1/billing/price/inquireList", &responseBilling.AlipayPriceInquireListResponse{})
}
