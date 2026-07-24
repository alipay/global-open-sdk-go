package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInvoiceInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result    `json:"result,omitempty"`
	Invoices   []*model.Invoice `json:"invoices,omitempty"`
	Total      string           `json:"total,omitempty"`
	HasMore    bool             `json:"hasMore,omitempty"`
	NextCursor string           `json:"nextCursor,omitempty"`
}
