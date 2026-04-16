package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipayCreateTransferRequest struct {
	TransferRequestId  string                    `json:"transferRequestId,omitempty"`
	TransferFromDetail *model.TransferFromDetail `json:"transferFromDetail,omitempty"`
	TransferToDetail   *model.TransferToDetail   `json:"transferToDetail,omitempty"`
}

func NewAlipayCreateTransferRequest() (*request.AlipayRequest, *AlipayCreateTransferRequest) {
	alipayCreateTransferRequest := &AlipayCreateTransferRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreateTransferRequest, "/ams/api/v1/funds/createTransfer", &responseMarketplace.AlipayCreateTransferResponse{})
	return alipayRequest, alipayCreateTransferRequest
}

func (alipayCreateTransferRequest *AlipayCreateTransferRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreateTransferRequest, "/ams/api/v1/funds/createTransfer", &responseMarketplace.AlipayCreateTransferResponse{})
}
