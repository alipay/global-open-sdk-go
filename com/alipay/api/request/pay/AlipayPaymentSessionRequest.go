package pay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responsePay "github.com/alipay/global-open-sdk-go/com/alipay/api/response/pay"
)

type AlipayPaymentSessionRequest struct {
	ProductCode                 model.ProductCodeType         `json:"productCode,omitempty"`
	PaymentRequestId            string                        `json:"paymentRequestId,omitempty"`
	Order                       *model.Order                  `json:"order,omitempty"`
	PaymentAmount               *model.Amount                 `json:"paymentAmount,omitempty"`
	PaymentMethod               *model.PaymentMethod          `json:"paymentMethod,omitempty"`
	PaymentSessionExpiryTime    string                        `json:"paymentSessionExpiryTime,omitempty"`
	PaymentRedirectUrl          string                        `json:"paymentRedirectUrl,omitempty"`
	PaymentNotifyUrl            string                        `json:"paymentNotifyUrl,omitempty"`
	PaymentFactor               *model.PaymentFactor          `json:"paymentFactor,omitempty"`
	SettlementStrategy          *model.SettlementStrategy     `json:"settlementStrategy,omitempty"`
	EnableInstallmentCollection bool                          `json:"enableInstallmentCollection,omitempty"`
	CreditPayPlan               *model.CreditPayPlan          `json:"creditPayPlan,omitempty"`
	MerchantRegion              string                        `json:"merchantRegion,omitempty"`
	Env                         *model.Env                    `json:"env,omitempty"`
	AgreementInfo               *model.AgreementInfo          `json:"agreementInfo,omitempty"`
	RiskData                    *model.RiskData               `json:"riskData,omitempty"`
	ProductScene                string                        `json:"productScene,omitempty"`
	SavedPaymentMethods         []*model.PaymentMethod        `json:"savedPaymentMethods,omitempty"`
	Locale                      string                        `json:"locale,omitempty"`
	AvailablePaymentMethod      *model.AvailablePaymentMethod `json:"availablePaymentMethod,omitempty"`
	TestFile                    string                        `json:"testFile,omitempty"`
}

func NewAlipayPaymentSessionRequest() (*request.AlipayRequest, *AlipayPaymentSessionRequest) {
	alipayPaymentSessionRequest := &AlipayPaymentSessionRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPaymentSessionRequest, "/ams/api/v1/payments/createPaymentSession", &responsePay.AlipayPaymentSessionResponse{})
	return alipayRequest, alipayPaymentSessionRequest
}

func (alipayPaymentSessionRequest *AlipayPaymentSessionRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPaymentSessionRequest, "/ams/api/v1/payments/createPaymentSession", &responsePay.AlipayPaymentSessionResponse{})
}
