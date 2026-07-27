package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreditGrantInquireListResponse struct {
	response.AlipayResponse
	Result       *model.Result      `json:"result,omitempty"`
	CreditGrants *model.CreditGrant `json:"creditGrants,omitempty"`
}
