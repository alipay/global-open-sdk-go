package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayMeterUpdateResponse struct {
	response.AlipayResponse
	Result *model.Result `json:"result,omitempty"`
	Meter  *model.Meter  `json:"meter,omitempty"`
}
