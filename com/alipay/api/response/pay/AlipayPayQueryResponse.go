package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayPayQueryResponse struct {
	response.AlipayResponse
	Metadata                 string                      `json:"metadata,omitempty"`
	Result                   *model.Result               `json:"result,omitempty"`
	CustomizedInfo           *model.CustomizedInfo       `json:"customizedInfo,omitempty"`
	ProcessingAmount         *model.Amount               `json:"processingAmount,omitempty"`
	PaymentStatus            model.TransactionStatusType `json:"paymentStatus,omitempty"`
	PaymentResultCode        string                      `json:"paymentResultCode,omitempty"`
	PaymentResultMessage     string                      `json:"paymentResultMessage,omitempty"`
	PaymentRequestId         string                      `json:"paymentRequestId,omitempty"`
	PaymentId                string                      `json:"paymentId,omitempty"`
	AuthPaymentId            string                      `json:"authPaymentId,omitempty"`
	PaymentAmount            *model.Amount               `json:"paymentAmount,omitempty"`
	ActualPaymentAmount      *model.Amount               `json:"actualPaymentAmount,omitempty"`
	PaymentQuote             *model.Quote                `json:"paymentQuote,omitempty"`
	AuthExpiryTime           string                      `json:"authExpiryTime,omitempty"`
	PaymentCreateTime        string                      `json:"paymentCreateTime,omitempty"`
	PaymentTime              string                      `json:"paymentTime,omitempty"`
	NonGuaranteeCouponAmount *model.Amount               `json:"nonGuaranteeCouponAmount,omitempty"`
	PspCustomerInfo          *model.PspCustomerInfo      `json:"pspCustomerInfo,omitempty"`
	RedirectActionForm       *model.RedirectActionForm   `json:"redirectActionForm,omitempty"`
	CardInfo                 *model.CardInfo             `json:"cardInfo,omitempty"`
	AcquirerReferenceNo      string                      `json:"acquirerReferenceNo,omitempty"`
	ExtendInfo               string                      `json:"extendInfo,omitempty"`
	Transactions             []*model.Transaction        `json:"transactions,omitempty"`
	CustomsDeclarationAmount *model.Amount               `json:"customsDeclarationAmount,omitempty"`
	GrossSettlementAmount    *model.Amount               `json:"grossSettlementAmount,omitempty"`
	SettlementQuote          *model.Quote                `json:"settlementQuote,omitempty"`
	PaymentResultInfo        *model.PaymentResultInfo    `json:"paymentResultInfo,omitempty"`
	AcquirerInfo             *model.AcquirerInfo         `json:"acquirerInfo,omitempty"`
	MerchantAccountId        string                      `json:"merchantAccountId,omitempty"`
	PromotionResults         []*model.PromotionResult    `json:"promotionResults,omitempty"`
	EarliestSettlementTime   string                      `json:"earliestSettlementTime,omitempty"`
	PaymentMethodType        string                      `json:"paymentMethodType,omitempty"`
}
