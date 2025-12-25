package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryStatementResponse struct {
	response.AlipayResponse
	StatementList     []*model.Statement  `json:"statementList,omitempty"`
	Result            *model.ResultResult `json:"result,omitempty"`
	TotalCount        int32               `json:"totalCount,omitempty"`
	TotalPageNumber   int32               `json:"totalPageNumber,omitempty"`
	CurrentPageNumber int32               `json:"currentPageNumber,omitempty"`
}
