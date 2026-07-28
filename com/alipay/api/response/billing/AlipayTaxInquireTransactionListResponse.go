package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayTaxInquireTransactionListResponse struct {
	response.AlipayResponse
	Result       *model.Result           `json:"result,omitempty"`
	Transactions []*model.TaxTransaction `json:"transactions,omitempty"`
	Paginator    *model.Paginator        `json:"paginator,omitempty"`
}
