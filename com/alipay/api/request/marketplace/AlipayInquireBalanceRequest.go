package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipayInquireBalanceRequest struct {
	ReferenceMerchantId string `json:"referenceMerchantId,omitempty"`
}

func NewAlipayInquireBalanceRequest() (*request.AlipayRequest, *AlipayInquireBalanceRequest) {
	alipayInquireBalanceRequest := &AlipayInquireBalanceRequest{}
	alipayRequest := request.NewAlipayRequest(alipayInquireBalanceRequest, model.MARKETPLACE_INQUIREBALANCE_PATH, &responseMarketplace.AlipayInquireBalanceResponse{})
	return alipayRequest, alipayInquireBalanceRequest
}
