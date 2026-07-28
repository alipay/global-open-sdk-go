package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxUpdateSettingsRequest struct {
	SettingsUpdateRequestId string               `json:"settingsUpdateRequestId,omitempty"`
	DefaultTaxCode          string               `json:"defaultTaxCode,omitempty"`
	DefaultTaxBehavior      string               `json:"defaultTaxBehavior,omitempty"`
	HeadOffice              *model.TaxHeadOffice `json:"headOffice,omitempty"`
}

func NewAlipayTaxUpdateSettingsRequest() (*request.AlipayRequest, *AlipayTaxUpdateSettingsRequest) {
	alipayTaxUpdateSettingsRequest := &AlipayTaxUpdateSettingsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxUpdateSettingsRequest, "/ams/api/v1/tax/updateSettings", &responseBilling.AlipayTaxUpdateSettingsResponse{})
	return alipayRequest, alipayTaxUpdateSettingsRequest
}

func (alipayTaxUpdateSettingsRequest *AlipayTaxUpdateSettingsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxUpdateSettingsRequest, "/ams/api/v1/tax/updateSettings", &responseBilling.AlipayTaxUpdateSettingsResponse{})
}
