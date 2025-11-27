package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipayInquireBalanceRequest struct {
	ReferenceMerchantId string `json:"referenceMerchantId,omitempty"`
}

func NewAlipayInquireBalanceRequest() (*request.AlipayRequest, *AlipayInquireBalanceRequest) {
	alipayInquireBalanceRequest := &AlipayInquireBalanceRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireBalanceRequest, "/ams/api/v1/accounts/inquireBalance", &responseMarketplace.AlipayInquireBalanceResponse{})
	return alipayRequest, alipayInquireBalanceRequest
}

func (alipayInquireBalanceRequest *AlipayInquireBalanceRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayInquireBalanceRequest, "/ams/api/v1/accounts/inquireBalance", &responseMarketplace.AlipayInquireBalanceResponse{})
}
