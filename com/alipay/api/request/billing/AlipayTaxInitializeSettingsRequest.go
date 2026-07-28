package billing

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/request"
	responseBilling "github.com/alipay/global-open-sdk-go/com/alipay/api/response/billing"
)

type AlipayTaxInitializeSettingsRequest struct {
	SettingsRequestId  string               `json:"settingsRequestId,omitempty"`
	DefaultTaxCode     string               `json:"defaultTaxCode,omitempty"`
	DefaultTaxBehavior string               `json:"defaultTaxBehavior,omitempty"`
	HeadOffice         *model.TaxHeadOffice `json:"headOffice,omitempty"`
}

func NewAlipayTaxInitializeSettingsRequest() (*request.AlipayRequest, *AlipayTaxInitializeSettingsRequest) {
	alipayTaxInitializeSettingsRequest := &AlipayTaxInitializeSettingsRequest{}
	alipayRequest := request.NewAlipayRequest(alipayTaxInitializeSettingsRequest, "/ams/api/v1/tax/initializeSettings", &responseBilling.AlipayTaxInitializeSettingsResponse{})
	return alipayRequest, alipayTaxInitializeSettingsRequest
}

func (alipayTaxInitializeSettingsRequest *AlipayTaxInitializeSettingsRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayTaxInitializeSettingsRequest, "/ams/api/v1/tax/initializeSettings", &responseBilling.AlipayTaxInitializeSettingsResponse{})
}
