package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditGrantUpdateExpiryRequest struct {
	CreditGrantId  string `json:"creditGrantId,omitempty"`
	ExpiryMode     string `json:"expiryMode,omitempty"`
	ExpiryDateTime string `json:"expiryDateTime,omitempty"`
}

func NewAlipayCreditGrantUpdateExpiryRequest() (*request.AlipayRequest, *AlipayCreditGrantUpdateExpiryRequest) {
	alipayCreditGrantUpdateExpiryRequest := &AlipayCreditGrantUpdateExpiryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditGrantUpdateExpiryRequest, "/ams/api/v1/meter/creditGrant/updateExpiry", &responseBilling.AlipayCreditGrantUpdateExpiryResponse{})
	return alipayRequest, alipayCreditGrantUpdateExpiryRequest
}

func (alipayCreditGrantUpdateExpiryRequest *AlipayCreditGrantUpdateExpiryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditGrantUpdateExpiryRequest, "/ams/api/v1/meter/creditGrant/updateExpiry", &responseBilling.AlipayCreditGrantUpdateExpiryResponse{})
}
