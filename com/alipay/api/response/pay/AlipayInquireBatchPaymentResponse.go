package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireBatchPaymentResponse struct {
	response.AlipayResponse
	Result                   *model.Result                    `json:"result,omitempty"`
	BatchPaymentRequestId    string                           `json:"batchPaymentRequestId,omitempty"`
	PaymentAmount            *model.Amount                    `json:"paymentAmount,omitempty"`
	PaymentCreateTime        string                           `json:"paymentCreateTime,omitempty"`
	PaymentMethodType        string                           `json:"paymentMethodType,omitempty"`
	RedirectActionForm       *model.RedirectActionForm        `json:"redirectActionForm,omitempty"`
	PspCustomerInfo          *model.PspCustomerInfo           `json:"pspCustomerInfo,omitempty"`
	PaymentTime              string                           `json:"paymentTime,omitempty"`
	AuthExpirytime           string                           `json:"authExpirytime,omitempty"`
	CustomsDeclarationAmount *model.Amount                    `json:"customsDeclarationAmount,omitempty"`
	PaymentResults           []*model.AmsInquiryPaymentResult `json:"paymentResults,omitempty"`
}
