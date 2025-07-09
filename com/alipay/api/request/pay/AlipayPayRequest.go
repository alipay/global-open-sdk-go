package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayPayRequest struct {
	ProductCode             model.ProductCodeType          `json:"productCode,omitempty"`
	PaymentRequestId        string                         `json:"paymentRequestId,omitempty"`
	Order                   *model.Order                   `json:"order,omitempty"`
	PaymentAmount           *model.Amount                  `json:"paymentAmount,omitempty"`
	PaymentMethod           *model.PaymentMethod           `json:"paymentMethod,omitempty"`
	PaymentExpiryTime       string                         `json:"paymentExpiryTime,omitempty"`
	PaymentRedirectUrl      string                         `json:"paymentRedirectUrl,omitempty"`
	PaymentNotifyUrl        string                         `json:"paymentNotifyUrl,omitempty"`
	PaymentFactor           *model.PaymentFactor           `json:"paymentFactor,omitempty"`
	SettlementStrategy      *model.SettlementStrategy      `json:"settlementStrategy,omitempty"`
	CreditPayPlan           *model.CreditPayPlan           `json:"creditPayPlan,omitempty"`
	AppId                   string                         `json:"appId,omitempty"`
	MerchantRegion          string                         `json:"merchantRegion,omitempty"`
	UserRegion              string                         `json:"userRegion,omitempty"`
	Env                     *model.Env                     `json:"env,omitempty"`
	PayToMethod             *model.PaymentMethod           `json:"payToMethod,omitempty"`
	IsAuthorization         *bool                          `json:"isAuthorization,omitempty"`
	Merchant                *model.Merchant                `json:"merchant,omitempty"`
	PaymentVerificationData *model.PaymentVerificationData `json:"paymentVerificationData,omitempty"`
	ExtendInfo              string                         `json:"extendInfo,omitempty"`
	MerchantAccountId       string                         `json:"merchantAccountId,omitempty"`
	SubscriptionInfo        *model.SubscriptionInfo        `json:"subscriptionInfo,omitempty"`
}

func (alipayPayRequest *AlipayPayRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPayRequest, model.PAYMENT_PATH, &responsePay.AlipayPayResponse{})
}

func NewAlipayPayRequest() (*request.AlipayRequest, *AlipayPayRequest) {
	alipayPayRequest := &AlipayPayRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayRequest, model.PAYMENT_PATH, &responsePay.AlipayPayResponse{})
	return alipayRequest, alipayPayRequest
}
