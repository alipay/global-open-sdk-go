package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateTransferResponse struct {
	response.AlipayResponse
	Result             *model.Result             `json:"result,omitempty"`
	TransferId         string                    `json:"transferId,omitempty"`
	TransferRequestId  string                    `json:"transferRequestId,omitempty"`
	TransferFromDetail *model.TransferFromDetail `json:"transferFromDetail,omitempty"`
	TransferToDetail   *model.TransferToDetail   `json:"transferToDetail,omitempty"`
}
