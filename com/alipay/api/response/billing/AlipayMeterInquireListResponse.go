package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayMeterInquireListResponse struct {
	response.AlipayResponse
	Result *model.Result  `json:"result,omitempty"`
	Meters []*model.Meter `json:"meters,omitempty"`
}
