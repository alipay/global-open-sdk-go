package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreditNoteInquireListResponse struct {
	response.AlipayResponse
	Result     *model.Result              `json:"result,omitempty"`
	HasMore    bool                       `json:"hasMore,omitempty"`
	TotalCount int32                      `json:"totalCount,omitempty"`
	List       []*model.CreditNoteSummary `json:"list,omitempty"`
}
