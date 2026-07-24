package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionCancelResponse struct {
	response.AlipayResponse
	Result              *model.ResultInfo                                   `json:"result,omitempty"`
	SubscriptionId      string                                              `json:"subscriptionId,omitempty"`
	Status              string                                              `json:"status,omitempty"`
	CancellationReason  string                                              `json:"cancellationReason,omitempty"`
	CancellationDetails *model.BillingSubscriptionCancelCancellationDetails `json:"cancellationDetails,omitempty"`
	CanceledAt          string                                              `json:"canceledAt,omitempty"`
	CancelAtPeriodEnd   bool                                                `json:"cancelAtPeriodEnd,omitempty"`
	CreditNoteId        string                                              `json:"creditNoteId,omitempty"`
}
