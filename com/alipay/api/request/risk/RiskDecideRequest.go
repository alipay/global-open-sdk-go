package risk

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseRisk "github.com/alipay/global-open-sdk-go/com/alipay/api/response/risk"
)

type RiskDecideRequest struct {
	ReferenceTransactionId string                   `json:"referenceTransactionId,omitempty"`
	AuthorizationPhase     model.AuthorizationPhase `json:"authorizationPhase,omitempty"`
	Orders                 []*model.Order           `json:"orders,omitempty"`
	Buyer                  *model.Buyer             `json:"buyer,omitempty"`
	ActualPaymentAmount    *model.Amount            `json:"actualPaymentAmount,omitempty"`
	PaymentDetails         []*model.PaymentDetail   `json:"paymentDetails,omitempty"`
	DiscountAmount         *model.Amount            `json:"discountAmount,omitempty"`
	EnvInfo                *model.Env               `json:"envInfo,omitempty"`
}

func NewRiskDecideRequest() (*request.AlipayRequest, *RiskDecideRequest) {
	riskDecideRequest := &RiskDecideRequest{}
	alipayRequest := request.NewAlipayRequest(riskDecideRequest, model.RISK_DECIDE_PATH, &responseRisk.RiskDecideResponse{})
	return alipayRequest, riskDecideRequest
}
