package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type TransferToDetailResponse struct {
	response.AlipayResponse
	TransferToMethod *model.PaymentMethod `json:"transferToMethod,omitempty"`
	TransferToAmount *model.Amount        `json:"transferToAmount,omitempty"`
	FeeAmount        *model.Amount        `json:"feeAmount,omitempty"`
}
