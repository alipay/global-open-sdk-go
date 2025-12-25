package responseAba

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayInquiryStatementDetailResponse struct {
	response.AlipayResponse
	Result    *model.Result     `json:"result,omitempty"`
	Statement *model.Statement  `json:"statement,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}
