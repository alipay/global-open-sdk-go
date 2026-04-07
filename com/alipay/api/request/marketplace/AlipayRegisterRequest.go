package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipayRegisterRequest struct {
	RegistrationRequestId string                  `json:"registrationRequestId,omitempty"`
	SettlementInfos       []*model.SettlementInfo `json:"settlementInfos,omitempty"`
	MerchantInfo          *model.MerchantInfo     `json:"merchantInfo,omitempty"`
	PaymentMethods        []*model.PaymentMethod  `json:"paymentMethods,omitempty"`
}

func NewAlipayRegisterRequest() (*request.AlipayRequest, *AlipayRegisterRequest) {
	alipayRegisterRequest := &AlipayRegisterRequest{}
	alipayRequest := request.NewAlipayRequest(alipayRegisterRequest, "null", &responseMarketplace.AlipayRegisterResponse{})
	return alipayRequest, alipayRegisterRequest
}

func (alipayRegisterRequest *AlipayRegisterRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayRegisterRequest, "null", &responseMarketplace.AlipayRegisterResponse{})
}
