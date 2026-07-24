package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditGrantCreateRequest struct {
	CustomerId        string               `json:"customerId,omitempty"`
	CreditGrantName   string               `json:"creditGrantName,omitempty"`
	Amount            *model.Amount        `json:"amount,omitempty"`
	Applicability     *model.Applicability `json:"applicability,omitempty"`
	Priority          int32                `json:"priority,omitempty"`
	Category          string               `json:"category,omitempty"`
	EffectiveDateTime string               `json:"effectiveDateTime,omitempty"`
	ExpiryDateTime    string               `json:"expiryDateTime,omitempty"`
}

func NewAlipayCreditGrantCreateRequest() (*request.AlipayRequest, *AlipayCreditGrantCreateRequest) {
	alipayCreditGrantCreateRequest := &AlipayCreditGrantCreateRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditGrantCreateRequest, "/ams/api/v1/meter/creditGrant/create", &responseBilling.AlipayCreditGrantCreateResponse{})
	return alipayRequest, alipayCreditGrantCreateRequest
}

func (alipayCreditGrantCreateRequest *AlipayCreditGrantCreateRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditGrantCreateRequest, "/ams/api/v1/meter/creditGrant/create", &responseBilling.AlipayCreditGrantCreateResponse{})
}
