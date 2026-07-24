package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayCreditGrantInquireListRequest struct {
	CustomerId string `json:"customerId,omitempty"`
	Status     string `json:"status,omitempty"`
}

func NewAlipayCreditGrantInquireListRequest() (*request.AlipayRequest, *AlipayCreditGrantInquireListRequest) {
	alipayCreditGrantInquireListRequest := &AlipayCreditGrantInquireListRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCreditGrantInquireListRequest, "/ams/api/v1/meter/creditGrant/inquireList", &responseBilling.AlipayCreditGrantInquireListResponse{})
	return alipayRequest, alipayCreditGrantInquireListRequest
}

func (alipayCreditGrantInquireListRequest *AlipayCreditGrantInquireListRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCreditGrantInquireListRequest, "/ams/api/v1/meter/creditGrant/inquireList", &responseBilling.AlipayCreditGrantInquireListResponse{})
}
