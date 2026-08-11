package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionUpdateResponse struct {
	response.AlipayResponse
	Result                   *model.ResultInfo         `json:"result,omitempty"`
	SubscriptionId           string                    `json:"subscriptionId,omitempty"`
	Status                   string                    `json:"status,omitempty"`
	SubscriptionItems        []*model.SubscriptionItem `json:"subscriptionItems,omitempty"`
	ProrationInvoiceId       string                    `json:"prorationInvoiceId,omitempty"`
	ProrationInvoiceAmount   int64                     `json:"prorationInvoiceAmount,omitempty"`
	ProrationInvoiceCurrency string                    `json:"prorationInvoiceCurrency,omitempty"`
	CreditNoteId             string                    `json:"creditNoteId,omitempty"`
	CreditNoteAmount         int64                     `json:"creditNoteAmount,omitempty"`
	CreditNoteCurrency       string                    `json:"creditNoteCurrency,omitempty"`
	PendingUpdate            bool                      `json:"pendingUpdate,omitempty"`
}
