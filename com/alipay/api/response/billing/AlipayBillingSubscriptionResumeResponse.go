package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayBillingSubscriptionResumeResponse struct {
	response.AlipayResponse
	Result             *model.ResultInfo `json:"result,omitempty"`
	SubscriptionId     string            `json:"subscriptionId,omitempty"`
	Status             string            `json:"status,omitempty"`
	BillingCycleAnchor string            `json:"billingCycleAnchor,omitempty"`
	ProrationInvoiceId string            `json:"prorationInvoiceId,omitempty"`
	ProrationDate      string            `json:"prorationDate,omitempty"`
}
