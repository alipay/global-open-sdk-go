package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxRegisterRequest struct {
	RegistrationRequestId string                 `json:"registrationRequestId,omitempty"`
	TaxType               string                 `json:"taxType,omitempty"`
	Jurisdiction          *model.TaxJurisdiction `json:"jurisdiction,omitempty"`
	RegistrationType      string                 `json:"registrationType,omitempty"`
	TaxId                 string                 `json:"taxId,omitempty"`
	ActiveFrom            string                 `json:"activeFrom,omitempty"`
	ExpireAt              string                 `json:"expireAt,omitempty"`
}

func NewAlipayTaxRegisterRequest() (*request.AlipayRequest, *AlipayTaxRegisterRequest) {
	alipayTaxRegisterRequest := &AlipayTaxRegisterRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxRegisterRequest, "/ams/api/v1/tax/register", &responseBilling.AlipayTaxRegisterResponse{})
	return alipayRequest, alipayTaxRegisterRequest
}

func (alipayTaxRegisterRequest *AlipayTaxRegisterRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxRegisterRequest, "/ams/api/v1/tax/register", &responseBilling.AlipayTaxRegisterResponse{})
}
