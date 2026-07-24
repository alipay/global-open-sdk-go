package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayCreditGrantCreateResponse struct {
	response.AlipayResponse
	Result      *model.Result      `json:"result,omitempty"`
	CreditGrant *model.CreditGrant `json:"creditGrant,omitempty"`
}
