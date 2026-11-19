package responseMarketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type TransferFromDetailResponse struct {
	response.AlipayResponse
	TransferFromMethod *model.PaymentMethodResponse `json:"transferFromMethod,omitempty"`
	TransferFromAmount *model.Amount                `json:"transferFromAmount,omitempty"`
	FeeAmount          *model.Amount                `json:"feeAmount,omitempty"`
}
