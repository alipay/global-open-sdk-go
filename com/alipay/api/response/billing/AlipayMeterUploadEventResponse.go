package responseBilling

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipayMeterUploadEventResponse struct {
	response.AlipayResponse
	Result     *model.Result  `json:"result,omitempty"`
	RetryAfter int32          `json:"retryAfter,omitempty"`
	Errors     []*model.Error `json:"errors,omitempty"`
}
