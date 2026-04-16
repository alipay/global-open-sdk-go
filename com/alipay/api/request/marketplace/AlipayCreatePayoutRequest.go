package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipayCreatePayoutRequest struct {
	TransferRequestId  string                    `json:"transferRequestId,omitempty"`
	TransferFromDetail *model.TransferFromDetail `json:"transferFromDetail,omitempty"`
	TransferToDetail   *model.TransferToDetail   `json:"transferToDetail,omitempty"`
}

func NewAlipayCreatePayoutRequest() (*request.AlipayRequest, *AlipayCreatePayoutRequest) {
	alipayCreatePayoutRequest := &AlipayCreatePayoutRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreatePayoutRequest, "/ams/api/v1/funds/createPayout", &responseMarketplace.AlipayCreatePayoutResponse{})
	return alipayRequest, alipayCreatePayoutRequest
}

func (alipayCreatePayoutRequest *AlipayCreatePayoutRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreatePayoutRequest, "/ams/api/v1/funds/createPayout", &responseMarketplace.AlipayCreatePayoutResponse{})
}
