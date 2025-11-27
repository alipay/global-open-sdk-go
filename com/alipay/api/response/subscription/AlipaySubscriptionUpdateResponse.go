package responseSubscription

import (
	"github.com/alipay/global-open-sdk-go/com/alipay/api/model"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/response"
)

type AlipaySubscriptionUpdateResponse struct {
	response.AlipayResponse
	Result *model.Result `json:"result,omitempty"`
}
