package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreateTransferResponse struct {
	response.AlipayResponse
	Result             *model.Result                     `json:"result,omitempty"`
	TransferRequestId  string                            `json:"transferRequestId,omitempty"`
	TransferId         string                            `json:"transferId,omitempty"`
	TransferFromDetail *model.TransferFromDetailResponse `json:"transferFromDetail,omitempty"`
	TransferToDetail   *model.TransferToDetailResponse   `json:"transferToDetail,omitempty"`
}
