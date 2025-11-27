package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipaySettleRequest struct {
	SettlementRequestId string                    `json:"settlementRequestId,omitempty"`
	PaymentId           string                    `json:"paymentId,omitempty"`
	SettlementDetails   []*model.SettlementDetail `json:"settlementDetails,omitempty"`
}

func NewAlipaySettleRequest() (*request.AlipayRequest, *AlipaySettleRequest) {
	alipaySettleRequest := &AlipaySettleRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySettleRequest, "/ams/api/v1/payments/settle", &responseMarketplace.AlipaySettleResponse{})
	return alipayRequest, alipaySettleRequest
}

func (alipaySettleRequest *AlipaySettleRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySettleRequest, "/ams/api/v1/payments/settle", &responseMarketplace.AlipaySettleResponse{})
}
