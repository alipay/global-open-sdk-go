package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquireCardTransactionEventResponse struct {
	response.AlipayResponse
	Result            *model.Result                 `json:"result,omitempty"`
	Events            []*model.CardTransactionEvent `json:"events,omitempty"`
	TotalCount        int32                         `json:"totalCount,omitempty"`
	TotalPageNumber   int32                         `json:"totalPageNumber,omitempty"`
	CurrentPageNumber int32                         `json:"currentPageNumber,omitempty"`
}
