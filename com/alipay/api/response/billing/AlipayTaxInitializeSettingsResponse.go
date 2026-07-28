package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayTaxInitializeSettingsResponse struct {
	response.AlipayResponse
	Result             *model.Result        `json:"result,omitempty"`
	DefaultTaxCode     string               `json:"defaultTaxCode,omitempty"`
	DefaultTaxBehavior string               `json:"defaultTaxBehavior,omitempty"`
	HeadOffice         *model.TaxHeadOffice `json:"headOffice,omitempty"`
	Status             string               `json:"status,omitempty"`
}
