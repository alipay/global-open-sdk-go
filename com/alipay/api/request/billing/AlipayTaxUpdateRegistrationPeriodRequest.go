package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxUpdateRegistrationPeriodRequest struct {
	RegistrationUpdatePeriodRequestId string `json:"registrationUpdatePeriodRequestId,omitempty"`
	TaxRegistrationId                 string `json:"taxRegistrationId,omitempty"`
	ActiveFrom                        string `json:"activeFrom,omitempty"`
	ExpireAt                          string `json:"expireAt,omitempty"`
}

func NewAlipayTaxUpdateRegistrationPeriodRequest() (*request.AlipayRequest, *AlipayTaxUpdateRegistrationPeriodRequest) {
	alipayTaxUpdateRegistrationPeriodRequest := &AlipayTaxUpdateRegistrationPeriodRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxUpdateRegistrationPeriodRequest, "/ams/api/v1/tax/updateRegistrationPeriod", &responseBilling.AlipayTaxUpdateRegistrationPeriodResponse{})
	return alipayRequest, alipayTaxUpdateRegistrationPeriodRequest
}

func (alipayTaxUpdateRegistrationPeriodRequest *AlipayTaxUpdateRegistrationPeriodRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxUpdateRegistrationPeriodRequest, "/ams/api/v1/tax/updateRegistrationPeriod", &responseBilling.AlipayTaxUpdateRegistrationPeriodResponse{})
}
