package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayTaxUpdateRegistrationPeriodResponse struct {
	response.AlipayResponse
	Result            *model.Result          `json:"result,omitempty"`
	TaxRegistrationId string                 `json:"taxRegistrationId,omitempty"`
	TaxType           string                 `json:"taxType,omitempty"`
	Jurisdiction      *model.TaxJurisdiction `json:"jurisdiction,omitempty"`
	RegistrationType  string                 `json:"registrationType,omitempty"`
	TaxId             string                 `json:"taxId,omitempty"`
	Status            string                 `json:"status,omitempty"`
	ActiveFrom        string                 `json:"activeFrom,omitempty"`
	ExpireAt          string                 `json:"expireAt,omitempty"`
}
