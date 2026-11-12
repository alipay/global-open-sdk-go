package responsePay

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayQueryTradeBillListResponse struct {
	response.AlipayResponse
	CurrentPage           int32                         `json:"currentPage,omitempty"`
	PageSize              int32                         `json:"pageSize,omitempty"`
	TimeZone              string                        `json:"timeZone,omitempty"`
	Total                 int64                         `json:"total,omitempty"`
	Records               []*model.TradeBillQueryInfo   `json:"records,omitempty"`
	Result                *model.Result                 `json:"result,omitempty"`
	ReferenceStoreId      string                        `json:"referenceStoreId,omitempty"`
	Tid                   string                        `json:"tid,omitempty"`
	MerchantTransactionId string                        `json:"merchantTransactionId,omitempty"`
	PaymentAmount         *model.TradeBillPaymentAmount `json:"paymentAmount,omitempty"`
	BizOrderType          string                        `json:"bizOrderType,omitempty"`
	StartTime             string                        `json:"startTime,omitempty"`
	CompletedTime         string                        `json:"completedTime,omitempty"`
	BizOrderStatus        string                        `json:"bizOrderStatus,omitempty"`
}
