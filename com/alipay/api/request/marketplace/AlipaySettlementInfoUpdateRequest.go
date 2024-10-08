package marketplace

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseMarketplace "github.com/alipay/global-open-sdk-go/com/alipay/api/response/marketplace"
)

type AlipaySettlementInfoUpdateRequest struct {
	UpdateRequestId       string                       `json:"updateRequestId,omitempty"`
	ReferenceMerchantId   string                       `json:"referenceMerchantId,omitempty"`
	SettlementCurrency    string                       `json:"settlementCurrency,omitempty"`
	SettlementBankAccount *model.SettlementBankAccount `json:"settlementBankAccount,omitempty"`
}

func NewAlipaySettlementInfoUpdateRequest() (*request.AlipayRequest, *AlipaySettlementInfoUpdateRequest) {
	alipaySettlementInfoUpdateRequest := &AlipaySettlementInfoUpdateRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySettlementInfoUpdateRequest, model.MARKETPLACE_SETTLEMENTINFO_UPDATE_PATH, &responseMarketplace.AlipaySettlementInfoUpdateResponse{})
	return alipayRequest, alipaySettlementInfoUpdateRequest
}
