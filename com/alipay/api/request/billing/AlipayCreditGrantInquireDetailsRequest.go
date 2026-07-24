package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditGrantInquireDetailsRequest struct {
	CreditGrantId string `json:"creditGrantId,omitempty"`
}

func NewAlipayCreditGrantInquireDetailsRequest() (*request.AlipayRequest, *AlipayCreditGrantInquireDetailsRequest) {
	alipayCreditGrantInquireDetailsRequest := &AlipayCreditGrantInquireDetailsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditGrantInquireDetailsRequest, "/ams/api/v1/meter/creditGrant/inquireDetails", &responseBilling.AlipayCreditGrantInquireDetailsResponse{})
	return alipayRequest, alipayCreditGrantInquireDetailsRequest
}

func (alipayCreditGrantInquireDetailsRequest *AlipayCreditGrantInquireDetailsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditGrantInquireDetailsRequest, "/ams/api/v1/meter/creditGrant/inquireDetails", &responseBilling.AlipayCreditGrantInquireDetailsResponse{})
}
