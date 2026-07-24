package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditGrantVoidRequest struct {
	CreditGrantId string `json:"creditGrantId,omitempty"`
	VoidReason    string `json:"voidReason,omitempty"`
}

func NewAlipayCreditGrantVoidRequest() (*request.AlipayRequest, *AlipayCreditGrantVoidRequest) {
	alipayCreditGrantVoidRequest := &AlipayCreditGrantVoidRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditGrantVoidRequest, "/ams/api/v1/meter/creditGrant/void", &responseBilling.AlipayCreditGrantVoidResponse{})
	return alipayRequest, alipayCreditGrantVoidRequest
}

func (alipayCreditGrantVoidRequest *AlipayCreditGrantVoidRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditGrantVoidRequest, "/ams/api/v1/meter/creditGrant/void", &responseBilling.AlipayCreditGrantVoidResponse{})
}
