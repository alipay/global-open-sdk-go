package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryStatementListResponse struct {
	response.AlipayResponse
	StatementList []*model.Statement `json:"statementList,omitempty"`
	Result        *model.Result      `json:"result,omitempty"`
	TotalCount    *model.TotalCount  `json:"totalCount,omitempty"`
}
