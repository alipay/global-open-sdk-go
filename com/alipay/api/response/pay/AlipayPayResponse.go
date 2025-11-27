package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPayResponse struct {
	response.AlipayResponse
	Result                  *model.Result              `json:"result,omitempty"`
	ProcessingAmount        *model.Amount              `json:"processingAmount,omitempty"`
	PaymentRequestId        string                     `json:"paymentRequestId,omitempty"`
	PaymentId               string                     `json:"paymentId,omitempty"`
	PaymentAmount           *model.Amount              `json:"paymentAmount,omitempty"`
	PaymentData             string                     `json:"paymentData,omitempty"`
	ActualPaymentAmount     *model.Amount              `json:"actualPaymentAmount,omitempty"`
	PaymentQuote            *model.Quote               `json:"paymentQuote,omitempty"`
	PaymentTime             string                     `json:"paymentTime,omitempty"`
	PaymentCreateTime       string                     `json:"paymentCreateTime,omitempty"`
	AuthExpiryTime          string                     `json:"authExpiryTime,omitempty"`
	NonGuaranteeCouponValue *model.Amount              `json:"nonGuaranteeCouponValue,omitempty"`
	PaymentActionForm       string                     `json:"paymentActionForm,omitempty"`
	PspCustomerInfo         *model.PspCustomerInfo     `json:"pspCustomerInfo,omitempty"`
	ChallengeActionForm     *model.ChallengeActionForm `json:"challengeActionForm,omitempty"`
	RedirectActionForm      *model.RedirectActionForm  `json:"redirectActionForm,omitempty"`
	OrderCodeForm           *model.OrderCodeForm       `json:"orderCodeForm,omitempty"`
	GrossSettlementAmount   *model.Amount              `json:"grossSettlementAmount,omitempty"`
	SettlementQuote         *model.Quote               `json:"settlementQuote,omitempty"`
	ExtendInfo              string                     `json:"extendInfo,omitempty"`
	NormalUrl               string                     `json:"normalUrl,omitempty"`
	SchemeUrl               string                     `json:"schemeUrl,omitempty"`
	ApplinkUrl              string                     `json:"applinkUrl,omitempty"`
	AppIdentifier           string                     `json:"appIdentifier,omitempty"`
	PaymentResultInfo       *model.PaymentResultInfo   `json:"paymentResultInfo,omitempty"`
	AcquirerInfo            *model.AcquirerInfo        `json:"acquirerInfo,omitempty"`
	PromotionResult         []*model.PromotionResult   `json:"promotionResult,omitempty"`
}
