package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayPayConsultRequest struct {
	ProductCode                 model.ProductCodeType     `json:"productCode,omitempty"`
	PaymentAmount               *model.Amount             `json:"paymentAmount,omitempty"`
	MerchantRegion              string                    `json:"merchantRegion,omitempty"`
	AllowedPaymentMethodRegions []string                  `json:"allowedPaymentMethodRegions,omitempty"`
	AllowedPaymentMethods       []string                  `json:"allowedPaymentMethods,omitempty"`
	BlockedPaymentMethods       []string                  `json:"blockedPaymentMethods,omitempty"`
	Region                      string                    `json:"region,omitempty"`
	CustomerId                  string                    `json:"customerId,omitempty"`
	ReferenceUserId             string                    `json:"referenceUserId,omitempty"`
	Env                         *model.Env                `json:"env,omitempty"`
	ExtendInfo                  string                    `json:"extendInfo,omitempty"`
	UserRegion                  string                    `json:"userRegion,omitempty"`
	PaymentFactor               *model.PaymentFactor      `json:"paymentFactor,omitempty"`
	SettlementStrategy          *model.SettlementStrategy `json:"settlementStrategy,omitempty"`
	Merchant                    *model.Merchant           `json:"merchant,omitempty"`
	AllowedPspRegions           []string                  `json:"allowedPspRegions,omitempty"`
	Buyer                       *model.Buyer              `json:"buyer,omitempty"`
	MerchantAccountId           string                    `json:"merchantAccountId,omitempty"`
}

func NewAlipayPayConsultRequest() (*request.AlipayRequest, *AlipayPayConsultRequest) {
	alipayPayConsultRequest := &AlipayPayConsultRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayConsultRequest, "/ams/api/v1/payments/consult", &responsePay.AlipayPayConsultResponse{})
	return alipayRequest, alipayPayConsultRequest
}

func (alipayPayConsultRequest *AlipayPayConsultRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPayConsultRequest, "/ams/api/v1/payments/consult", &responsePay.AlipayPayConsultResponse{})
}
