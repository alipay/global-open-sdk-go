package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxCancelRegistrationRequest struct {
	RegistrationCancelRequestId string `json:"registrationCancelRequestId,omitempty"`
	TaxRegistrationId           string `json:"taxRegistrationId,omitempty"`
}

func NewAlipayTaxCancelRegistrationRequest() (*request.AlipayRequest, *AlipayTaxCancelRegistrationRequest) {
	alipayTaxCancelRegistrationRequest := &AlipayTaxCancelRegistrationRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxCancelRegistrationRequest, "/ams/api/v1/tax/cancelRegistration", &responseBilling.AlipayTaxCancelRegistrationResponse{})
	return alipayRequest, alipayTaxCancelRegistrationRequest
}

func (alipayTaxCancelRegistrationRequest *AlipayTaxCancelRegistrationRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxCancelRegistrationRequest, "/ams/api/v1/tax/cancelRegistration", &responseBilling.AlipayTaxCancelRegistrationResponse{})
}
